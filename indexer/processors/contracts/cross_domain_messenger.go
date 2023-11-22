package contracts

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/indexer/bigint"
	"github.com/ethereum-optimism/optimism/indexer/database"
	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain"
)

var (
	// Standard ABI types copied from golang ABI tests
	addressType, _ = abi.NewType("address", "", nil)
	bytesType, _   = abi.NewType("bytes", "", nil)
	uint256Type, _ = abi.NewType("uint256", "", nil)

	CrossDomainMessengerLegacyRelayMessageEncoding = abi.NewMethod(
		"relayMessage",
		"relayMessage",
		abi.Function,
		"external", // mutability
		false,      // isConst
		true,       // payable
		abi.Arguments{ // inputs
			{Name: "target", Type: addressType},
			{Name: "sender", Type: addressType},
			{Name: "data", Type: bytesType},
			{Name: "nonce", Type: uint256Type},
			// The actual transaction on the legacy L1CrossDomainMessenger has a trailing
			// proof argument but is ignored for the `XDomainCallData` encoding
		},
		abi.Arguments{}, // outputs
	)
)

type CrossDomainMessengerSentMessageEvent struct {
	Event         *database.ContractEvent
	BridgeMessage database.BridgeMessage
	Version       uint16
}

type CrossDomainMessengerRelayedMessageEvent struct {
	Event       *database.ContractEvent
	MessageHash common.Hash
}

func CrossDomainMessengerSentMessageEvents(chainSelector string, contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]CrossDomainMessengerSentMessageEvent, error) {
	crossDomainMessengerAbi, err := bindings.CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	sentMessageEventAbi := crossDomainMessengerAbi.Events["SentMessage"]
	contractEventFilter := database.ContractEvent{ContractAddress: contractAddress, EventSignature: sentMessageEventAbi.ID}
	sentMessageEvents, err := db.ContractEvents.ContractEventsWithFilter(contractEventFilter, chainSelector, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	if len(sentMessageEvents) == 0 {
		return nil, nil
	}

	sentMessageExtensionEventAbi := crossDomainMessengerAbi.Events["SentMessageExtension1"]
	contractEventFilter = database.ContractEvent{ContractAddress: contractAddress, EventSignature: sentMessageExtensionEventAbi.ID}
	sentMessageExtensionEvents, err := db.ContractEvents.ContractEventsWithFilter(contractEventFilter, chainSelector, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	if len(sentMessageExtensionEvents) > len(sentMessageEvents) {
		return nil, fmt.Errorf("mismatch. %d sent messages & %d sent message extensions", len(sentMessageEvents), len(sentMessageExtensionEvents))
	}

	// We handle version zero messages uniquely since the first version of cross domain messages
	// do not have the SentMessageExtension1 event emitted, introduced in version 1.
	numVersionZeroMessages := len(sentMessageEvents) - len(sentMessageExtensionEvents)
	crossDomainSentMessages := make([]CrossDomainMessengerSentMessageEvent, len(sentMessageEvents))
	for i := range sentMessageEvents {
		sentMessage := bindings.CrossDomainMessengerSentMessage{Raw: *sentMessageEvents[i].RLPLog}
		err = UnpackLog(&sentMessage, sentMessageEvents[i].RLPLog, sentMessageEventAbi.Name, crossDomainMessengerAbi)
		if err != nil {
			return nil, err
		}

		_, versionBig := crossdomain.DecodeVersionedNonce(sentMessage.MessageNonce)
		version := uint16(versionBig.Uint64())
		if i < numVersionZeroMessages && version != 0 {
			return nil, fmt.Errorf("expected version zero nonce: nonce %d, tx_hash %s", sentMessage.MessageNonce, sentMessage.Raw.TxHash)
		}

		value, messageHash := bigint.Zero, common.Hash{}
		switch version {
		case 0:
			messageHash, err = crossdomain.HashCrossDomainMessageV0(sentMessage.Target, sentMessage.Sender, sentMessage.Message, sentMessage.MessageNonce)
			if err != nil {
				return nil, err
			}

		case 1:
			sentMessageExtension := bindings.CrossDomainMessengerSentMessageExtension1{Raw: *sentMessageExtensionEvents[i].RLPLog}
			err = UnpackLog(&sentMessageExtension, sentMessageExtensionEvents[i].RLPLog, sentMessageExtensionEventAbi.Name, crossDomainMessengerAbi)
			if err != nil {
				return nil, err
			}

			value = sentMessageExtension.Value
			messageHash, err = crossdomain.HashCrossDomainMessageV1(sentMessage.MessageNonce, sentMessage.Sender, sentMessage.Target, value, sentMessage.GasLimit, sentMessage.Message)
			if err != nil {
				return nil, err
			}

		default:
			// NOTE: We explicitly fail here since the presence of a new version means finalization
			// logic needs to be updated to ensure L1 finalization can run from genesis and handle
			// the changing version formats. This is meant to be a serving indicator of this.
			return nil, fmt.Errorf("expected cross domain version 0 or version 1: %d", version)
		}

		crossDomainSentMessages[i] = CrossDomainMessengerSentMessageEvent{
			Event:   &sentMessageEvents[i],
			Version: version,
			BridgeMessage: database.BridgeMessage{
				MessageHash:          messageHash,
				Nonce:                sentMessage.MessageNonce,
				SentMessageEventGUID: sentMessageEvents[i].GUID,
				GasLimit:             sentMessage.GasLimit,
				Tx: database.Transaction{
					FromAddress: sentMessage.Sender,
					ToAddress:   sentMessage.Target,
					Amount:      value,
					Data:        sentMessage.Message,
					Timestamp:   sentMessageEvents[i].Timestamp,
				},
			},
		}
	}

	return crossDomainSentMessages, nil
}

func CrossDomainMessengerRelayedMessageEvents(chainSelector string, contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]CrossDomainMessengerRelayedMessageEvent, error) {
	crossDomainMessengerAbi, err := bindings.CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	relayedMessageEventAbi := crossDomainMessengerAbi.Events["RelayedMessage"]
	contractEventFilter := database.ContractEvent{ContractAddress: contractAddress, EventSignature: relayedMessageEventAbi.ID}
	relayedMessageEvents, err := db.ContractEvents.ContractEventsWithFilter(contractEventFilter, chainSelector, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	crossDomainRelayedMessages := make([]CrossDomainMessengerRelayedMessageEvent, len(relayedMessageEvents))
	for i := range relayedMessageEvents {
		relayedMessage := bindings.CrossDomainMessengerRelayedMessage{Raw: *relayedMessageEvents[i].RLPLog}
		err = UnpackLog(&relayedMessage, relayedMessageEvents[i].RLPLog, relayedMessageEventAbi.Name, crossDomainMessengerAbi)
		if err != nil {
			return nil, err
		}

		crossDomainRelayedMessages[i] = CrossDomainMessengerRelayedMessageEvent{
			Event:       &relayedMessageEvents[i],
			MessageHash: relayedMessage.MsgHash,
		}
	}

	return crossDomainRelayedMessages, nil
}
