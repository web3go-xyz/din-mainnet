package bridge

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/indexer/database"
	"github.com/ethereum-optimism/optimism/indexer/processors/contracts"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"

	"github.com/ethereum/go-ethereum/log"
)

// L2ProcessInitiatedBridgeEvents will query the database for new bridge events that have been initiated between
// the specified block range. This covers every part of the multi-layered stack:
//  1. OptimismPortal
//  2. L2CrossDomainMessenger
//  3. L2StandardBridge
func L2ProcessInitiatedBridgeEvents(log log.Logger, db *database.DB, fromHeight *big.Int, toHeight *big.Int) error {
	// (1) L2ToL1MessagePasser
	l2ToL1MPMessagesPassed, err := contracts.L2ToL1MessagePasserMessagePassedEvents(predeploys.L2ToL1MessagePasserAddr, db, fromHeight, toHeight)
	if err != nil {
		return err
	}

	messagesPassed := make(map[logKey]*contracts.L2ToL1MessagePasserMessagePassed, len(l2ToL1MPMessagesPassed))
	transactionWithdrawals := make([]database.L2TransactionWithdrawal, len(l2ToL1MPMessagesPassed))
	for i := range l2ToL1MPMessagesPassed {
		messagePassed := l2ToL1MPMessagesPassed[i]
		messagesPassed[logKey{messagePassed.Event.BlockHash, messagePassed.Event.LogIndex}] = &messagePassed
		transactionWithdrawals[i] = database.L2TransactionWithdrawal{
			WithdrawalHash:       messagePassed.WithdrawalHash,
			InitiatedL2EventGUID: messagePassed.Event.GUID,
			Nonce:                messagePassed.Nonce,
			GasLimit:             messagePassed.GasLimit,
			Tx:                   messagePassed.Tx,
		}
	}

	if len(messagesPassed) > 0 {
		log.Info("detected transaction withdrawals", "size", len(transactionWithdrawals))
		if err := db.BridgeTransactions.StoreL2TransactionWithdrawals(transactionWithdrawals); err != nil {
			return err
		}
	}

	// (2) L2CrossDomainMessenger
	crossDomainSentMessages, err := contracts.CrossDomainMessengerSentMessageEvents("l2", predeploys.L2CrossDomainMessengerAddr, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainSentMessages) > len(messagesPassed) {
		return fmt.Errorf("missing L2ToL1MP withdrawal for each cross-domain message. withdrawals: %d, messages: %d", len(messagesPassed), len(crossDomainSentMessages))
	}

	sentMessages := make(map[logKey]*contracts.CrossDomainMessengerSentMessageEvent, len(crossDomainSentMessages))
	l2BridgeMessages := make([]database.L2BridgeMessage, len(crossDomainSentMessages))
	for i := range crossDomainSentMessages {
		sentMessage := crossDomainSentMessages[i]
		sentMessages[logKey{sentMessage.Event.BlockHash, sentMessage.Event.LogIndex}] = &sentMessage

		// extract the withdrawal hash from the previous MessagePassed event
		messagePassed, ok := messagesPassed[logKey{sentMessage.Event.BlockHash, sentMessage.Event.LogIndex - 1}]
		if !ok {
			return fmt.Errorf("expected MessagePassedEvent preceding SentMessage. tx_hash = %s", sentMessage.Event.TransactionHash)
		}

		l2BridgeMessages[i] = database.L2BridgeMessage{TransactionWithdrawalHash: messagePassed.WithdrawalHash, BridgeMessage: sentMessage.BridgeMessage}
	}

	if len(l2BridgeMessages) > 0 {
		log.Info("detected L2CrossDomainMessenger messages", "size", len(l2BridgeMessages))
		if err := db.BridgeMessages.StoreL2BridgeMessages(l2BridgeMessages); err != nil {
			return err
		}
	}

	// (3) L2StandardBridge
	initiatedBridges, err := contracts.StandardBridgeInitiatedEvents("l2", predeploys.L2StandardBridgeAddr, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(initiatedBridges) > len(crossDomainSentMessages) {
		return fmt.Errorf("missing cross-domain message for each initiated bridge event. messages: %d, bridges: %d", len(crossDomainSentMessages), len(initiatedBridges))
	}

	l2BridgeWithdrawals := make([]database.L2BridgeWithdrawal, len(initiatedBridges))
	for i := range initiatedBridges {
		initiatedBridge := initiatedBridges[i]

		// extract the cross domain message hash & deposit source hash from the following events
		messagePassed, ok := messagesPassed[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 1}]
		if !ok {
			return fmt.Errorf("expected MessagePassed following BridgeInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash)
		}
		sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 2}]
		if !ok {
			return fmt.Errorf("expected SentMessage following MessagePassed event. tx_hash = %s", initiatedBridge.Event.TransactionHash)
		}

		initiatedBridge.BridgeTransfer.CrossDomainMessageHash = &sentMessage.BridgeMessage.MessageHash
		l2BridgeWithdrawals[i] = database.L2BridgeWithdrawal{TransactionWithdrawalHash: messagePassed.WithdrawalHash, BridgeTransfer: initiatedBridge.BridgeTransfer}
	}

	if len(l2BridgeWithdrawals) > 0 {
		log.Info("detected L2StandardBridge withdrawals", "size", len(l2BridgeWithdrawals))
		if err := db.BridgeTransfers.StoreL2BridgeWithdrawals(l2BridgeWithdrawals); err != nil {
			return err
		}
	}

	// a-ok!
	return nil
}

// L2ProcessFinalizedBridgeEvent will query the database for all the finalization markers for all initiated
// bridge events. This covers every part of the multi-layered stack:
//  1. L2CrossDomainMessenger (relayMessage marker)
//  2. L2StandardBridge (no-op, since this is simply a wrapper over the L2CrossDomainMEssenger)
//
// NOTE: Unlike L1, there's no L2ToL1MessagePasser stage since transaction deposits are apart of the block derivation process.
func L2ProcessFinalizedBridgeEvents(log log.Logger, db *database.DB, fromHeight *big.Int, toHeight *big.Int) error {
	// (1) L2CrossDomainMessenger relayedMessage
	crossDomainRelayedMessages, err := contracts.CrossDomainMessengerRelayedMessageEvents("l2", predeploys.L2CrossDomainMessengerAddr, db, fromHeight, toHeight)
	if err != nil {
		return err
	}

	relayedMessages := make(map[logKey]*contracts.CrossDomainMessengerRelayedMessageEvent, len(crossDomainRelayedMessages))
	for i := range crossDomainRelayedMessages {
		relayed := crossDomainRelayedMessages[i]
		relayedMessages[logKey{BlockHash: relayed.Event.BlockHash, LogIndex: relayed.Event.LogIndex}] = &relayed
		message, err := db.BridgeMessages.L1BridgeMessage(relayed.MessageHash)
		if err != nil {
			return err
		} else if message == nil {
			log.Error("missing indexed L1CrossDomainMessenger message", "message_hash", relayed.MessageHash, "tx_hash", relayed.Event.TransactionHash)
			return fmt.Errorf("missing indexed L1CrossDomainMessager message")
		}

		if err := db.BridgeMessages.MarkRelayedL1BridgeMessage(relayed.MessageHash, relayed.Event.GUID); err != nil {
			return err
		}
	}

	if len(crossDomainRelayedMessages) > 0 {
		log.Info("relayed L1CrossDomainMessenger messages", "size", len(crossDomainRelayedMessages))
	}

	// (2) L2StandardBridge BridgeFinalized
	finalizedBridges, err := contracts.StandardBridgeFinalizedEvents("l2", predeploys.L2StandardBridgeAddr, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(finalizedBridges) > len(crossDomainRelayedMessages) {
		return fmt.Errorf("missing cross-domain message for each finalized bridge event. messages: %d, bridges: %d", len(crossDomainRelayedMessages), len(finalizedBridges))
	}

	for i := range finalizedBridges {
		// Nothing actionable on the database. However, we can treat the relayed message
		// as an invariant by ensuring we can query for a deposit by the same hash
		finalizedBridge := finalizedBridges[i]
		relayedMessage, ok := relayedMessages[logKey{finalizedBridge.Event.BlockHash, finalizedBridge.Event.LogIndex + 1}]
		if !ok {
			return fmt.Errorf("expected RelayedMessage following BridgeFinalized event. tx_hash = %s", finalizedBridge.Event.TransactionHash)
		}

		// Since the message hash is computed from the relayed message, this ensures the withdrawal fields must match. For good measure,
		// we may choose to make sure `deposit.BridgeTransfer` matches with the finalized bridge
		deposit, err := db.BridgeTransfers.L1BridgeDepositWithFilter(database.BridgeTransfer{CrossDomainMessageHash: &relayedMessage.MessageHash})
		if err != nil {
			return err
		} else if deposit == nil {
			log.Error("missing L1StandardBridge deposit on L2 finalization", "tx_hash", finalizedBridge.Event.TransactionHash)
			return errors.New("missing L1StandardBridge deposit on L2 finalization")
		}
	}

	// a-ok!
	return nil
}
