// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { ISemver } from "src/universal/ISemver.sol";
import { L2ToL1MessagePasser } from "src/L2/L2ToL1MessagePasser.sol";

/// @custom:proxied
/// @custom:predeploy 0x4200000000000000000000000000000000000007
/// @title L2CrossDomainMessenger
/// @notice The L2CrossDomainMessenger is a high-level interface for message passing between L1 and
///         L2 on the L2 side. Users are generally encouraged to use this contract instead of lower
///         level message passing contracts.
contract L2CrossDomainMessenger is CrossDomainMessenger, ISemver {
    /// @custom:semver 1.6.0
    string public constant version = "1.6.0";

    /// @notice Constructs the L2CrossDomainMessenger contract.
    /// @param _l1CrossDomainMessenger Address of the L1CrossDomainMessenger contract.
    constructor(address _l1CrossDomainMessenger) CrossDomainMessenger(_l1CrossDomainMessenger) {
        initialize();
    }

    /// @notice Initializer.
    function initialize() public reinitializer(2) {
        __CrossDomainMessenger_init();
    }

    /// @custom:legacy
    /// @notice Legacy getter for the remote messenger.
    ///         Use otherMessenger going forward.
    /// @return Address of the L1CrossDomainMessenger contract.
    function l1CrossDomainMessenger() public view returns (address) {
        return OTHER_MESSENGER;
    }

    /// @inheritdoc CrossDomainMessenger
    function _sendMessage(address _to, uint64 _gasLimit, uint256 _value, bytes memory _data) internal override {
        L2ToL1MessagePasser(payable(Predeploys.L2_TO_L1_MESSAGE_PASSER)).initiateWithdrawal{ value: _value }(
            _to, _gasLimit, _data
        );
    }

    /// @inheritdoc CrossDomainMessenger
    function _isOtherMessenger() internal view override returns (bool) {
        return AddressAliasHelper.undoL1ToL2Alias(msg.sender) == OTHER_MESSENGER;
    }

    /// @inheritdoc CrossDomainMessenger
    function _isUnsafeTarget(address _target) internal view override returns (bool) {
        return _target == address(this) || _target == address(Predeploys.L2_TO_L1_MESSAGE_PASSER);
    }
}
