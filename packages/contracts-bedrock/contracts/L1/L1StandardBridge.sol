// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Lib_PredeployAddresses } from "../libraries/Lib_PredeployAddresses.sol";
import { StandardBridge } from "../universal/StandardBridge.sol";

/**
 * @custom:proxied
 * @title L1StandardBridge
 * @notice The L1StandardBridge is responsible for transfering ETH and ERC20 tokens between L1 and
 *         L2. ERC20 tokens deposited into L2 are escrowed within this contract until withdrawal.
 *         ETH is transferred to and escrowed within the OptimismPortal contract.
 */
contract L1StandardBridge is StandardBridge {
    /**
     * @custom:legacy
     * @notice Emitted whenever a deposit of ETH from L1 into L2 is initiated.
     *
     * @param _from      Address of the depositor.
     * @param _to        Address of the recipient on L2.
     * @param _amount    Amount of ETH deposited.
     * @param _extraData Extra data attached to the deposit.
     */
    event ETHDepositInitiated(
        address indexed _from,
        address indexed _to,
        uint256 _amount,
        bytes _extraData
    );

    /**
     * @custom:legacy
     * @notice Emitted whenever a withdrawal of ETH from L2 to L1 is finalized.
     *
     * @param _from      Address of the withdrawer.
     * @param _to        Address of the recipient on L1.
     * @param _amount    Amount of ETH withdrawn.
     * @param _extraData Extra data attached to the withdrawal.
     */
    event ETHWithdrawalFinalized(
        address indexed _from,
        address indexed _to,
        uint256 _amount,
        bytes _extraData
    );

    /**
     * @custom:legacy
     * @notice Emitted whenever an ERC20 deposit is initiated.
     *
     * @param _l1Token   Address of the token on L1.
     * @param _l2Token   Address of the corresponding token on L2.
     * @param _from      Address of the depositor.
     * @param _to        Address of the recipient on L2.
     * @param _amount    Amount of the ERC20 deposited.
     * @param _extraData Extra data attached to the deposit.
     */
    event ERC20DepositInitiated(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _extraData
    );

    /**
     * @custom:legacy
     * @notice Emitted whenever an ERC20 withdrawal is finalized.
     *
     * @param _l1Token   Address of the token on L1.
     * @param _l2Token   Address of the corresponding token on L2.
     * @param _from      Address of the withdrawer.
     * @param _to        Address of the recipient on L1.
     * @param _amount    Amount of the ERC20 withdrawn.
     * @param _extraData Extra data attached to the withdrawal.
     */
    event ERC20WithdrawalFinalized(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _extraData
    );

    /**
     * @notice Initializes the L1StandardBridge.
     *
     * @param _messenger Address of the L1CrossDomainMessenger.
     */
    function initialize(address payable _messenger) public {
        _initialize(_messenger, payable(Lib_PredeployAddresses.L2_STANDARD_BRIDGE));
    }

    /**
     * @custom:legacy
     * @notice Retrieves the access of the corresponding L2 bridge contract.
     *
     * @return Address of the corresponding L2 bridge contract.
     */
    function l2TokenBridge() external returns (address) {
        return address(otherBridge);
    }

    /**
     * @custom:legacy
     * @notice Deposits some amount of ETH into the sender's account on L2.
     *
     * @param _minGasLimit Minimum gas limit for the deposit message on L2.
     * @param _extraData   Optional data to forward to L2. Data supplied here will not be used to
     *                     execute any code on L2 and is only emitted as extra data for the
     *                     convenience of off-chain tooling.
     */
    function depositETH(uint32 _minGasLimit, bytes calldata _extraData) external payable onlyEOA {
        _initiateETHDeposit(msg.sender, msg.sender, _minGasLimit, _extraData);
    }

    /**
     * @custom:legacy
     * @notice Deposits some amount of ETH into a target account on L2.
     *         Note that if ETH is sent to a contract on L2 and the call fails, then that ETH will
     *         be locked in the L2StandardBridge. ETH may be recoverable if the call can be
     *         successfully replayed by increasing the amount of gas supplied to the call. If the
     *         call will fail for any amount of gas, then the ETH will be locked permanently.
     *
     * @param _to          Address of the recipient on L2.
     * @param _minGasLimit Minimum gas limit for the deposit message on L2.
     * @param _extraData   Optional data to forward to L2. Data supplied here will not be used to
     *                     execute any code on L2 and is only emitted as extra data for the
     *                     convenience of off-chain tooling.
     */
    function depositETHTo(
        address _to,
        uint32 _minGasLimit,
        bytes calldata _extraData
    ) external payable {
        _initiateETHDeposit(msg.sender, _to, _minGasLimit, _extraData);
    }

    /**
     * @custom:legacy
     * @notice Deposits some amount of ERC20 tokens into the sender's account on L2.
     *
     * @param _l1Token     Address of the L1 token being deposited.
     * @param _l2Token     Address of the corresponding token on L2.
     * @param _amount      Amount of the ERC20 to deposit.
     * @param _minGasLimit Minimum gas limit for the deposit message on L2.
     * @param _extraData   Optional data to forward to L2. Data supplied here will not be used to
     *                     execute any code on L2 and is only emitted as extra data for the
     *                     convenience of off-chain tooling.
     */
    function depositERC20(
        address _l1Token,
        address _l2Token,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes calldata _extraData
    ) external virtual onlyEOA {
        _initiateERC20Deposit(
            _l1Token,
            _l2Token,
            msg.sender,
            msg.sender,
            _amount,
            _minGasLimit,
            _extraData
        );
    }

    /**
     * @custom:legacy
     * @notice Deposits some amount of ERC20 tokens into a target account on L2.
     *
     * @param _l1Token     Address of the L1 token being deposited.
     * @param _l2Token     Address of the corresponding token on L2.
     * @param _to          Address of the recipient on L2.
     * @param _amount      Amount of the ERC20 to deposit.
     * @param _minGasLimit Minimum gas limit for the deposit message on L2.
     * @param _extraData   Optional data to forward to L2. Data supplied here will not be used to
     *                     execute any code on L2 and is only emitted as extra data for the
     *                     convenience of off-chain tooling.
     */
    function depositERC20To(
        address _l1Token,
        address _l2Token,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes calldata _extraData
    ) external virtual {
        _initiateERC20Deposit(
            _l1Token,
            _l2Token,
            msg.sender,
            _to,
            _amount,
            _minGasLimit,
            _extraData
        );
    }

    /**
     * @custom:legacy
     * @notice Finalizes a withdrawal of ETH from L2.
     *
     * @param _from      Address of the withdrawer on L2.
     * @param _to        Address of the recipient on L1.
     * @param _amount    Amount of ETH to withdraw.
     * @param _extraData Optional data forwarded from L2.
     */
    function finalizeETHWithdrawal(
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _extraData
    ) external payable onlyOtherBridge {
        emit ETHWithdrawalFinalized(_from, _to, _amount, _extraData);
        finalizeBridgeETH(_from, _to, _amount, _extraData);
    }

    /**
     * @custom:legacy
     * @notice Finalizes a withdrawal of ERC20 tokens from L2.
     *
     * @param _l1Token   Address of the token on L1.
     * @param _l2Token   Address of the corresponding token on L2.
     * @param _from      Address of the withdrawer on L2.
     * @param _to        Address of the recipient on L1.
     * @param _amount    Amount of ETH to withdraw.
     * @param _extraData Optional data forwarded from L2.
     */
    function finalizeERC20Withdrawal(
        address _l1Token,
        address _l2Token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _extraData
    ) external onlyOtherBridge {
        emit ERC20WithdrawalFinalized(_l1Token, _l2Token, _from, _to, _amount, _extraData);
        finalizeBridgeERC20(_l1Token, _l2Token, _from, _to, _amount, _extraData);
    }

    /**
     * @notice Internal function for initiating an ETH deposit.
     *
     * @param _from        Address of the sender on L1.
     * @param _to          Address of the recipient on L2.
     * @param _minGasLimit Minimum gas limit for the deposit message on L2.
     * @param _extraData   Optional data to forward to L2.
     */
    function _initiateETHDeposit(
        address _from,
        address _to,
        uint32 _minGasLimit,
        bytes memory _extraData
    ) internal {
        emit ETHDepositInitiated(_from, _to, msg.value, _extraData);
        _initiateBridgeETH(_from, _to, msg.value, _minGasLimit, _extraData);
    }

    /**
     * @notice Internal function for initiating an ERC20 deposit.
     *
     * @param _l1Token     Address of the L1 token being deposited.
     * @param _l2Token     Address of the corresponding token on L2.
     * @param _from        Address of the sender on L1.
     * @param _to          Address of the recipient on L2.
     * @param _amount      Amount of the ERC20 to deposit.
     * @param _minGasLimit Minimum gas limit for the deposit message on L2.
     * @param _extraData   Optional data to forward to L2.
     */
    function _initiateERC20Deposit(
        address _l1Token,
        address _l2Token,
        address _from,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes calldata _extraData
    ) internal {
        emit ERC20DepositInitiated(_l1Token, _l2Token, _from, _to, _amount, _extraData);
        _initiateBridgeERC20(_l1Token, _l2Token, _from, _to, _amount, _minGasLimit, _extraData);
    }
}
