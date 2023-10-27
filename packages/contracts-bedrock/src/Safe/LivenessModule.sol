// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Safe, OwnerManager } from "safe-contracts/Safe.sol";
import { Enum } from "safe-contracts/common/Enum.sol";
import { OwnerManager } from "safe-contracts/base/OwnerManager.sol";
import { LivenessGuard } from "src/Safe/LivenessGuard.sol";
import { ISemver } from "src/universal/ISemver.sol";

/// @title LivenessModule
/// @notice This module is intended to be used in conjunction with the LivenessGuard. In the event
///         that an owner of the safe is not recorded by the guard during the liveness interval,
///         the owner will be considered inactive and will be removed from the list of owners.
///         If the number of owners falls below the minimum number of owners, the ownership of the
///         safe will be transferred to the fallback owner.
contract LivenessModule is ISemver {
    /// @notice The Safe contract instance
    Safe internal immutable SAFE;

    /// @notice The LivenessGuard contract instance
    ///         This can be updated by replacing with a new module and switching out the guard.
    LivenessGuard internal immutable LIVENESS_GUARD;

    /// @notice The interval, in seconds, during which an owner must have demonstrated liveness
    ///         This can be updated by replacing with a new module.
    uint256 internal immutable LIVENESS_INTERVAL;

    /// @notice The minimum number of owners before ownership of the safe is transferred to the fallback owner.
    ///         This can be updated by replacing with a new module.
    uint256 internal immutable MIN_OWNERS;

    /// @notice The fallback owner of the Safe
    ///         This can be updated by replacing with a new module.
    address internal immutable FALLBACK_OWNER;

    /// @notice The storage slot used in the safe to store the guard address
    ///         keccak256("guard_manager.guard.address")
    uint256 internal constant GUARD_STORAGE_SLOT = 0x4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c8;

    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    // Constructor to initialize the Safe and baseModule instances
    constructor(
        Safe _safe,
        LivenessGuard _livenessGuard,
        uint256 _livenessInterval,
        uint256 _minOwners,
        address _fallbackOwner
    ) {
        SAFE = _safe;
        LIVENESS_GUARD = _livenessGuard;
        LIVENESS_INTERVAL = _livenessInterval;
        FALLBACK_OWNER = _fallbackOwner;
        MIN_OWNERS = _minOwners;
        address[] memory owners = _safe.getOwners();
        require(_minOwners <= owners.length, "LivenessModule: minOwners must be less than the number of owners");
        require(
            _safe.getThreshold() == get75PercentThreshold(owners.length),
            "LivenessModule: Safe must have a threshold of 75% of the number of owners"
        );
    }

    /// @notice This function can be called by anyone to remove a set of owners that have not signed a transaction
    ///         during the liveness interval. If the number of owners drops below the minimum, then all owners
    ///         must be removed.
    /// @param _previousOwners The previous owners in the linked list of owners
    /// @param _ownersToRemove The owners to remove
    function removeOwners(address[] memory _previousOwners, address[] memory _ownersToRemove) external {
        require(_previousOwners.length == _ownersToRemove.length, "LivenessModule: arrays must be the same length");

        // We will remove at least one owner, so we'll initialize the ownersCount count to the current number of owners
        uint256 ownersCount = SAFE.getOwners().length;
        for (uint256 i = 0; i < _previousOwners.length; i++) {
            ownersCount--;
            if (ownersCount >= MIN_OWNERS) {
                require(_canRemove(_ownersToRemove[i]), "LivenessModule: the owner to remove has signed recently");
            }

            _removeOwner({
                _prevOwner: _previousOwners[i],
                _ownerToRemove: _ownersToRemove[i],
                _newOwnersCount: ownersCount
            });
        }
        _verifyFinalState();
    }

    /// @notice Removes an owner from the Safe and updates the threshold.
    /// @param _prevOwner Owner that pointed to the owner to be removed in the linked list
    /// @param _ownerToRemove Owner address to be removed.
    /// @param _newOwnersCount New number of owners after removal.
    function _removeOwner(address _prevOwner, address _ownerToRemove, uint256 _newOwnersCount) internal {
        if (_newOwnersCount > 0) {
            uint256 newThreshold = get75PercentThreshold(_newOwnersCount);
            // Remove the owner and update the threshold
            _removeOwnerSafeCall({ _prevOwner: _prevOwner, _owner: _ownerToRemove, _threshold: newThreshold });
        } else {
            // There is only one owner left. The Safe will not allow a safe with no owners, so we will
            // need to swap owners instead.
            _swapToFallbackOwnerSafeCall({ _prevOwner: _prevOwner, _oldOwner: _ownerToRemove });
        }
    }

    /// @notice Sets the fallback owner as the sole owner of the Safe with a threshold of 1
    /// @param _prevOwner Owner that pointed to the owner to be replaced in the linked list
    /// @param _oldOwner Owner address to be replaced.
    function _swapToFallbackOwnerSafeCall(address _prevOwner, address _oldOwner) internal {
        require(
            SAFE.execTransactionFromModule({
                to: address(SAFE),
                value: 0,
                operation: Enum.Operation.Call,
                data: abi.encodeCall(OwnerManager.swapOwner, (_prevOwner, _oldOwner, FALLBACK_OWNER))
            }),
            "LivenessModule: failed to swap to fallback owner"
        );
    }

    /// @notice Removes the owner `owner` from the Safe and updates the threshold to `_threshold`.
    /// @param _prevOwner Owner that pointed to the owner to be removed in the linked list
    /// @param _owner Owner address to be removed.
    /// @param _threshold New threshold.
    function _removeOwnerSafeCall(address _prevOwner, address _owner, uint256 _threshold) internal {
        require(
            SAFE.execTransactionFromModule({
                to: address(SAFE),
                value: 0,
                operation: Enum.Operation.Call,
                data: abi.encodeCall(OwnerManager.removeOwner, (_prevOwner, _owner, _threshold))
            }),
            "LivenessModule: failed to remove owner"
        );
    }

    /// @notice Checks if the owner can be removed
    /// @param _owner The owner to be removed
    /// @return canRemove_ bool indicating if the owner can be removed
    function _canRemove(address _owner) internal view returns (bool canRemove_) {
        canRemove_ = LIVENESS_GUARD.lastLive(_owner) + LIVENESS_INTERVAL < block.timestamp;
    }

    /// @notice A FREI-PI invariant check enforcing requirements on number of owners and threshold.
    function _verifyFinalState() internal view {
        address[] memory owners = SAFE.getOwners();
        uint256 numOwners = owners.length;
        // Ensure that the safe is not being left in an unsafe state with too few owners.
        if (numOwners == 1) {
            require(owners[0] == FALLBACK_OWNER, "LivenessModule: must transfer ownership to fallback owner");
        } else {
            require(
                numOwners >= MIN_OWNERS,
                "LivenessModule: must remove all owners and transfer to fallback owner if numOwners < minOwners"
            );
        }

        // Check that the threshold is correct. This check is also correct when there is a single
        // owner, because get75PercentThreshold(1) returns 1.
        uint256 threshold = SAFE.getThreshold();
        require(
            threshold == get75PercentThreshold(numOwners),
            "LivenessModule: Safe must have a threshold of 75% of the number of owners"
        );

        // Check that the guard has not been changed
        require(
            address(LIVENESS_GUARD) == address(uint160(uint256(bytes32(SAFE.getStorageAt(GUARD_STORAGE_SLOT, 1))))),
            "LivenessModule: guard has been changed"
        );
    }

    /// @notice For a given number of owners, return the lowest threshold which is greater than 75.
    ///         Note: this function returns 1 for numOwners == 1.
    function get75PercentThreshold(uint256 _numOwners) public pure returns (uint256 threshold_) {
        threshold_ = (_numOwners * 75 + 99) / 100;
    }

    /// @notice Getter function for the Safe contract instance
    /// @return safe_ The Safe contract instance
    function safe() public view returns (Safe safe_) {
        safe_ = SAFE;
    }

    /// @notice Getter function for the LivenessGuard contract instance
    /// @return livenessGuard_ The LivenessGuard contract instance
    function livenessGuard() public view returns (LivenessGuard livenessGuard_) {
        livenessGuard_ = LIVENESS_GUARD;
    }

    /// @notice Getter function for the liveness interval
    /// @return livenessInterval_ The liveness interval, in seconds
    function livenessInterval() public view returns (uint256 livenessInterval_) {
        livenessInterval_ = LIVENESS_INTERVAL;
    }

    /// @notice Getter function for the minimum number of owners
    /// @return minOwners_ The minimum number of owners
    function minOwners() public view returns (uint256 minOwners_) {
        minOwners_ = MIN_OWNERS;
    }

    /// @notice Getter function for the fallback owner
    /// @return fallbackOwner_ The fallback owner of the Safe
    function fallbackOwner() public view returns (address fallbackOwner_) {
        fallbackOwner_ = FALLBACK_OWNER;
    }
}
