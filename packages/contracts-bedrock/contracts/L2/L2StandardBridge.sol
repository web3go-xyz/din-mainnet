// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Lib_PredeployAddresses } from "../libraries/Lib_PredeployAddresses.sol";
import { StandardBridge } from "../universal/StandardBridge.sol";
import { OptimismMintableERC20 } from "../universal/OptimismMintableERC20.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000010
 * @title L2StandardBridge
 * @notice The L2StandardBridge is responsible for transfering ETH and ERC20 tokens between L1 and
 *         L2. ERC20 tokens sent to L1 are escrowed within this contract.
 *         Note that this contract is not intended to support all variations of ERC20 tokens.
 *         Examples of some token types that may not be properly supported by this contract include,
 *         but are not limited to: tokens with transfer fees, rebasing tokens, and
 *         tokens with blocklists.
 *         TODO: ensure that this has 1:1 backwards compatibility
 */
contract L2StandardBridge is StandardBridge {
    /**
     * @custom:legacy
     * @notice Emitted whenever a withdrawal from L2 to L1 is initiated.
     *
     * @param _l1Token   Address of the token on L1.
     * @param _l2Token   Address of the corresponding token on L2.
     * @param _from      Address of the withdrawer.
     * @param _to        Address of the recipient on L1.
     * @param _amount    Amount of the ERC20 withdrawn.
     * @param _extraData Extra data attached to the withdrawal.
     */
    event WithdrawalInitiated(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _extraData
    );

    /**
     * @custom:legacy
     * @notice Emitted whenever an ERC20 deposit is finalized.
     *
     * @param _l1Token   Address of the token on L1.
     * @param _l2Token   Address of the corresponding token on L2.
     * @param _from      Address of the depositor.
     * @param _to        Address of the recipient on L2.
     * @param _amount    Amount of the ERC20 deposited.
     * @param _extraData Extra data attached to the deposit.
     */
    event DepositFinalized(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _extraData
    );

    /**
     * @custom:legacy
     * @notice Emitted whenever a deposit fails.
     *
     * @param _l1Token   Address of the token on L1.
     * @param _l2Token   Address of the corresponding token on L2.
     * @param _from      Address of the depositor.
     * @param _to        Address of the recipient on L2.
     * @param _amount    Amount of the ERC20 deposited.
     * @param _extraData Extra data attached to the deposit.
     */
    event DepositFailed(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _extraData
    );

    /**
     * @notice Initializes the L2StandardBridge.
     *
     * @param _otherBridge Address of the L1StandardBridge.
     */
    function initialize(address payable _otherBridge) public {
        _initialize(payable(Lib_PredeployAddresses.L2_CROSS_DOMAIN_MESSENGER), _otherBridge);
    }

    /**
     * @custom:legacy
     * @notice Initiates a withdrawal from L2 to L1.
     *
     * @param _l2Token     Address of the L2 token to withdraw.
     * @param _amount      Amount of the L2 token to withdraw.
     * @param _minGasLimit Minimum gas limit to use for the transaction.
     * @param _extraData   Extra data attached to the withdrawal.
     */
    function withdraw(
        address _l2Token,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes calldata _extraData
    ) external payable virtual onlyEOA {
        _initiateWithdrawal(_l2Token, msg.sender, msg.sender, _amount, _minGasLimit, _extraData);
    }

    /**
     * @custom:legacy
     * @notice Initiates a withdrawal from L2 to L1 to a target account on L1.
     *         Note that if ETH is sent to a contract on L1 and the call fails, then that ETH will
     *         be locked in the L1StandardBridge. ETH may be recoverable if the call can be
     *         successfully replayed by increasing the amount of gas supplied to the call. If the
     *         call will fail for any amount of gas, then the ETH will be locked permanently.
     *
     * @param _l2Token     Address of the L2 token to withdraw.
     * @param _to          Recipient account on L1.
     * @param _amount      Amount of the L2 token to withdraw.
     * @param _minGasLimit Minimum gas limit to use for the transaction.
     * @param _extraData   Extra data attached to the withdrawal.
     */
    function withdrawTo(
        address _l2Token,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes calldata _extraData
    ) external payable virtual {
        _initiateWithdrawal(_l2Token, msg.sender, _to, _amount, _minGasLimit, _extraData);
    }

    /**
     * @custom:legacy
     * @notice Finalizes a deposit from L1 to L2.
     *
     * @param _l1Token   Address of the L1 token to deposit.
     * @param _l2Token   Address of the corresponding L2 token.
     * @param _from      Address of the depositor.
     * @param _to        Address of the recipient.
     * @param _amount    Amount of the tokens being deposited.
     * @param _extraData Extra data attached to the deposit.
     */
    function finalizeDeposit(
        address _l1Token,
        address _l2Token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _extraData
    ) external payable virtual {
        if (_l1Token == address(0) && _l2Token == Lib_PredeployAddresses.OVM_ETH) {
            finalizeBridgeETH(_from, _to, _amount, _extraData);
        } else {
            finalizeBridgeERC20(_l2Token, _l1Token, _from, _to, _amount, _extraData);
        }
        emit DepositFinalized(_l1Token, _l2Token, _from, _to, _amount, _extraData);
    }

    /**
     * @custom:legacy
     * @notice Internal function to a withdrawal from L2 to L1 to a target account on L1.
     *
     * @param _l2Token     Address of the L2 token to withdraw.
     * @param _from        Address of the withdrawer.
     * @param _to          Recipient account on L1.
     * @param _amount      Amount of the L2 token to withdraw.
     * @param _minGasLimit Minimum gas limit to use for the transaction.
     * @param _extraData   Extra data attached to the withdrawal.
     */
    function _initiateWithdrawal(
        address _l2Token,
        address _from,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes calldata _extraData
    ) internal {
        address l1Token = OptimismMintableERC20(_l2Token).l1Token();
        if (_l2Token == Lib_PredeployAddresses.OVM_ETH) {
            require(msg.value == _amount, "ETH withdrawals must include sufficient ETH value.");
            _initiateBridgeETH(_from, _to, _amount, _minGasLimit, _extraData);
        } else {
            _initiateBridgeERC20(_l2Token, l1Token, _from, _to, _amount, _minGasLimit, _extraData);
        }
        emit WithdrawalInitiated(l1Token, _l2Token, _from, _to, _amount, _extraData);
    }
}
