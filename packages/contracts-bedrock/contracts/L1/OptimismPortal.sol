//SPDX-License-Identifier: MIT
pragma solidity 0.8.10;

import { L2OutputOracle } from "./L2OutputOracle.sol";
import { WithdrawalVerifier } from "../libraries/Lib_WithdrawalVerifier.sol";
import { AddressAliasHelper } from "../libraries/AddressAliasHelper.sol";
import { ExcessivelySafeCall } from "../libraries/ExcessivelySafeCall.sol";
import { ResourceMetering } from "./ResourceMetering.sol";

/**
 * @custom:proxied
 * @title OptimismPortal
 * @notice The OptimismPortal is a low-level contract responsible for passing messages between L1
 *         and L2. Messages sent directly to the OptimismPortal have no form of replayability.
 *         Users are encouraged to use the L1CrossDomainMessenger for a higher-level interface.
 */
contract OptimismPortal is ResourceMetering {
    /**
     * @notice Emitted when a transaction is deposited from L1 to L2. The parameters of this event
     *         are read by the rollup node and used to derive deposit transactions on L2.
     *
     * @param from       Address that triggered the deposit transaction.
     * @param to         Address that the deposit transaction is directed to.
     * @param mint       Amount of ETH to mint to the sender on L2.
     * @param value      Amount of ETH to send to the recipient.
     * @param gasLimit   Minimum gas limit that the message can be executed with.
     * @param isCreation Whether the message is a contract creation.
     * @param data       Data to attach to the message and call the recipient with.
     */
    event TransactionDeposited(
        address indexed from,
        address indexed to,
        uint256 mint,
        uint256 value,
        uint64 gasLimit,
        bool isCreation,
        bytes data
    );

    /**
     * @notice Emitted when a withdrawal transaction is finalized.
     *
     * @param withdrawalHash Hash of the withdrawal transaction.
     * @param success        Whether the withdrawal transaction was successful.
     */
    event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success);

    /**
     * @notice Value used to reset the l2Sender, this is more efficient than setting it to zero.
     */
    address internal constant DEFAULT_L2_SENDER = 0x000000000000000000000000000000000000dEaD;

    /**
     * @notice Minimum time (in seconds) that must elapse before a withdrawal can be finalized.
     */
    uint256 public immutable FINALIZATION_PERIOD_SECONDS;

    /**
     * @notice Address of the L2OutputOracle.
     */
    L2OutputOracle public immutable L2_ORACLE;

    /**
     * @notice Address of the L2 account which initiated a withdrawal in this transaction. If the
     *         of this variable is the default L2 sender address, then we are NOT inside of a call
     *         to finalizeWithdrawalTransaction.
     */
    address public l2Sender = DEFAULT_L2_SENDER;

    /**
     * @notice The L2 gas limit set when eth is deposited using the receive() function.
     */
    uint64 internal constant RECEIVE_DEFAULT_GAS_LIMIT = 100_000;

    /**
     * @notice Additional gas reserved for clean up after finalizing a transaction withdrawal.
     */
    uint256 internal constant FINALIZE_GAS_BUFFER = 20_000;

    /**
     * @notice A list of withdrawal hashes which have been successfully finalized.
     */
    mapping(bytes32 => bool) public finalizedWithdrawals;

    /**
     * @param _l2Oracle                  Address of the L2OutputOracle.
     * @param _finalizationPeriodSeconds Finalization time in seconds.
     */
    constructor(L2OutputOracle _l2Oracle, uint256 _finalizationPeriodSeconds) {
        L2_ORACLE = _l2Oracle;
        FINALIZATION_PERIOD_SECONDS = _finalizationPeriodSeconds;
    }

    /**
     * @notice Accepts value so that users can send ETH directly to this contract and have the
     *         funds be deposited to their address on L2. This is intended as a convenience
     *         function for EOAs. Contracts should call the depositTransaction() function directly
     *         otherwise any deposited funds will be lost due to address aliasing.
     */
    receive() external payable {
        depositTransaction(msg.sender, msg.value, RECEIVE_DEFAULT_GAS_LIMIT, false, bytes(""));
    }

    /**
     * @notice Accepts deposits of ETH and data, and emits a TransactionDeposited event for use in
     *         deriving deposit transactions. Note that if a deposit is made by a contract, its
     *         address will be aliased when retrieved using `tx.origin` or `msg.sender`. Consider
     *         using the CrossDomainMessenger contracts for a simpler developer experience.
     *
     * @param _to         Target address on L2.
     * @param _value      ETH value to send to the recipient.
     * @param _gasLimit   Minimum L2 gas limit (can be greater than or equal to this value).
     * @param _isCreation Whether or not the transaction is a contract creation.
     * @param _data       Data to trigger the recipient with.
     */
    function depositTransaction(
        address _to,
        uint256 _value,
        uint64 _gasLimit,
        bool _isCreation,
        bytes memory _data
    ) public payable metered(_gasLimit) {
        // Just to be safe, make sure that people specify address(0) as the target when doing
        // contract creations.
        // TODO: Do we really need this? Prevents some user error, but adds gas.
        if (_isCreation) {
            require(
                _to == address(0),
                "OptimismPortal: must send to address(0) when creating a contract"
            );
        }

        // Transform the from-address to its alias if the caller is a contract.
        address from = msg.sender;
        if (msg.sender != tx.origin) {
            from = AddressAliasHelper.applyL1ToL2Alias(msg.sender);
        }

        // Emit a TransactionDeposited event so that the rollup node can derive a deposit
        // transaction for this deposit.
        emit TransactionDeposited(from, _to, msg.value, _value, _gasLimit, _isCreation, _data);
    }

    /**
     * @notice Finalizes a withdrawal transaction.
     *
     * @param _nonce           Nonce for the provided message.
     * @param _sender          Message sender address on L2.
     * @param _target          Target address on L1.
     * @param _value           ETH to send to the target.
     * @param _gasLimit        Minumum gas to be forwarded to the target.
     * @param _data            Data to send to the target.
     * @param _l2BlockNumber   L2 block number of the outputRoot.
     * @param _outputRootProof Inclusion proof of the withdrawer contracts storage root.
     * @param _withdrawalProof Inclusion proof for the given withdrawal in the withdrawer contract.
     */
    function finalizeWithdrawalTransaction(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes calldata _data,
        uint256 _l2BlockNumber,
        WithdrawalVerifier.OutputRootProof calldata _outputRootProof,
        bytes calldata _withdrawalProof
    ) external payable {
        // Prevent nested withdrawals within withdrawals.
        require(
            l2Sender == DEFAULT_L2_SENDER,
            "OptimismPortal: can only trigger one withdrawal per transaction"
        );

        // Prevent users from creating a deposit transaction where this address is the message
        // sender on L2.
        require(
            _target != address(this),
            "OptimismPortal: you cannot send messages to the portal contract"
        );

        // Get the output root.
        L2OutputOracle.OutputProposal memory proposal = L2_ORACLE.getL2Output(_l2BlockNumber);

        // Ensure that enough time has passed since the proposal was submitted before allowing a
        // withdrawal. Under the assumption that the fault proof mechanism is operating correctly,
        // we can infer that any withdrawal that has passed the finalization period must be valid
        // and can therefore be operated on.
        require(
            block.timestamp > proposal.timestamp + FINALIZATION_PERIOD_SECONDS,
            "OptimismPortal: proposal is not yet finalized"
        );

        // Verify that the output root can be generated with the elements in the proof.
        require(
            proposal.outputRoot == WithdrawalVerifier._deriveOutputRoot(_outputRootProof),
            "OptimismPortal: invalid output root proof"
        );

        // All withdrawals have a unique hash, we'll use this as the identifier for the withdrawal
        // and to prevent replay attacks.
        bytes32 withdrawalHash = WithdrawalVerifier.withdrawalHash(
            _nonce,
            _sender,
            _target,
            _value,
            _gasLimit,
            _data
        );

        // Verify that the hash of this withdrawal was stored in the withdrawal contract on L2. If
        // this is true, then we know that this withdrawal was actually triggered on L2 can can
        // therefore be relayed on L1.
        require(
            WithdrawalVerifier._verifyWithdrawalInclusion(
                withdrawalHash,
                _outputRootProof.withdrawerStorageRoot,
                _withdrawalProof
            ),
            "OptimismPortal: invalid withdrawal inclusion proof"
        );

        // Check that this withdrawal has not already been finalized, this is replay protection.
        require(
            finalizedWithdrawals[withdrawalHash] == false,
            "OptimismPortal: withdrawal has already been finalized"
        );

        // Mark the withdrawal as finalized so it can't be replayed.
        finalizedWithdrawals[withdrawalHash] = true;

        // We want to maintain the property that the amount of gas supplied to the call to the
        // target contract is at least the gas limit specified by the user. We can do this by
        // enforcing that, at this point in time, we still have gaslimit + buffer gas available.
        require(
            gasleft() >= _gasLimit + FINALIZE_GAS_BUFFER,
            "OptimismPortal: insufficient gas to finalize withdrawal"
        );

        // Set the l2Sender so contracts know who triggered this withdrawal on L2.
        l2Sender = _sender;

        // Trigger the call to the target contract. We use excessivelySafeCall because we don't
        // care about the returndata and we don't want target contracts to be able to force this
        // call to run out of gas.
        (bool success, ) = ExcessivelySafeCall.excessivelySafeCall(
            _target,
            _gasLimit,
            _value,
            0,
            _data
        );

        // Reset the l2Sender back to the default value.
        l2Sender = DEFAULT_L2_SENDER;

        // All withdrawals are immediately finalized. Replayability can
        // be achieved through contracts built on top of this contract
        emit WithdrawalFinalized(withdrawalHash, success);
    }
}
