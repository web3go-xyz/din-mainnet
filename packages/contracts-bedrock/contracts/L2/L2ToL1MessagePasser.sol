// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { WithdrawalVerifier } from "../libraries/Lib_WithdrawalVerifier.sol";
import { Burn } from "../libraries/Burn.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000000
 * @title L2ToL1MessagePasser
 * @notice The L2ToL1MessagePasser is a dedicated contract where messages that are being sent from
 *         L2 to L1 can be stored. The storage root of this contract is pulled up to the top level
 *         of the L2 output to reduce the cost of proving the existence of sent messages.
 */
contract L2ToL1MessagePasser {
    /**
     * @notice Emitted any time a withdrawal is initiated.
     *
     * @param nonce    Unique value corresponding to each withdrawal.
     * @param sender   The L2 account address which initiated the withdrawal.
     * @param target   The L1 account address the call will be send to.
     * @param value    The ETH value submitted for withdrawal, to be forwarded to the target.
     * @param gasLimit The minimum amount of gas that must be provided when withdrawing on L1.
     * @param data     The data to be forwarded to the target on L1.
     */
    event WithdrawalInitiated(
        uint256 indexed nonce,
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 gasLimit,
        bytes data
    );

    /**
     * @notice Emitted when the balance of this contract is burned.
     *
     * @param amount Amount of ETh that was burned.
     */
    event WithdrawerBalanceBurnt(uint256 indexed amount);

    /*************
     * Constants *
     *************/

    /**
     * @notice The L1 gas limit set when eth is withdrawn using the receive() function.
     */
    uint256 internal constant RECEIVE_DEFAULT_GAS_LIMIT = 100_000;

    /*************
     * Variables *
     *************/

    /**
     * @notice Includes the message hashes for all withdrawals
     */
    mapping(bytes32 => bool) public sentMessages;

    /**
     * @notice A unique value hashed with each withdrawal.
     */
    uint256 public nonce;

    /**
     * @notice Allows users to withdraw ETH by sending directly to this contract.
     */
    receive() external payable {
        initiateWithdrawal(msg.sender, RECEIVE_DEFAULT_GAS_LIMIT, bytes(""));
    }

    /**
     * @notice Sends a message from L2 to L1.
     *
     * @param _target   Address to call on L1 execution.
     * @param _gasLimit Minimum gas limit for executing the message on L1.
     * @param _data     Data to forward to L1 target.
     */
    function initiateWithdrawal(
        address _target,
        uint256 _gasLimit,
        bytes memory _data
    ) public payable {
        bytes32 withdrawalHash = WithdrawalVerifier.withdrawalHash(
            nonce,
            msg.sender,
            _target,
            msg.value,
            _gasLimit,
            _data
        );

        sentMessages[withdrawalHash] = true;

        emit WithdrawalInitiated(nonce, msg.sender, _target, msg.value, _gasLimit, _data);
        unchecked {
            ++nonce;
        }
    }

    /**
     * @notice Removes all ETH held by this contract from the state. Used to prevent the amount of
     *         ETH on L2 inflating when ETH is withdrawn. Currently only way to do this is to
     *         create a contract and self-destruct it to itself. Anyone can call this function. Not
     *         incentivized since this function is very cheap.
     */
    function burn() external {
        uint256 balance = address(this).balance;
        Burn.eth(balance);
        emit WithdrawerBalanceBurnt(balance);
    }
}
