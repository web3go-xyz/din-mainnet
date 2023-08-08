// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { Semver } from "../universal/Semver.sol";
import { ResourceMetering } from "./ResourceMetering.sol";

/// @title SystemConfig
/// @notice The SystemConfig contract is used to manage configuration of an Optimism network.
///         All configuration is stored on L1 and picked up by L2 as part of the derviation of
///         the L2 chain.
contract SystemConfig is OwnableUpgradeable, Semver {
    /// @notice Enum representing different types of updates.
    /// @custom:value BATCHER              Represents an update to the batcher hash.
    /// @custom:value GAS_CONFIG           Represents an update to txn fee config on L2.
    /// @custom:value GAS_LIMIT            Represents an update to gas limit on L2.
    /// @custom:value UNSAFE_BLOCK_SIGNER  Represents an update to the signer key for unsafe
    ///                                    block distrubution.
    enum UpdateType {
        BATCHER,
        GAS_CONFIG,
        GAS_LIMIT,
        UNSAFE_BLOCK_SIGNER
    }

    struct Addresses {
        address l1CrossDomainMessenger;
        address l1ERC721Bridge;
        address l1StandardBridge;
        address l2OutputOracle;
        address optimismPortal;
    }

    /// @notice Version identifier, used for upgrades.
    uint256 public constant VERSION = 0;

    /// @notice Storage slot that the unsafe block signer is stored at.
    ///         Storing it at this deterministic storage slot allows for decoupling the storage
    ///         layout from the way that `solc` lays out storage. The `op-node` uses a storage
    ///         proof to fetch this value.
    bytes32 public constant UNSAFE_BLOCK_SIGNER_SLOT = keccak256("systemconfig.unsafeblocksigner");

    /// @notice
    bytes32 public constant L1_CROSS_DOMAIN_MESSENGER_SLOT = keccak256("systemconfig.l1crossdomainmessenger");

    /// @notice
    bytes32 public constant L1_ERC_721_BRIDGE_SLOT = keccak256("systemconfig.l1erc721bridge");

    /// @notice
    bytes32 public constant L1_STANDARD_BRIDGE_SLOT = keccak256("systemconfig.l1standardbridge");

    /// @notice
    bytes32 public constant L2_OUTPUT_ORACLE_SLOT = keccak256("systemconfig.l2outputoracle");

    /// @notice
    bytes32 public constant OPTIMISM_PORTAL_SLOT = keccak256("systemconfig.optimismportal");

    /// @notice
    bytes32 public constant BATCH_INBOX_SLOT = keccak256("systemconfig.batchinbox");

    /// @notice Fixed L2 gas overhead. Used as part of the L2 fee calculation.
    uint256 public overhead;

    /// @notice Dynamic L2 gas overhead. Used as part of the L2 fee calculation.
    uint256 public scalar;

    /// @notice Identifier for the batcher.
    ///         For version 1 of this configuration, this is represented as an address left-padded
    ///         with zeros to 32 bytes.
    bytes32 public batcherHash;

    /// @notice L2 block gas limit.
    uint64 public gasLimit;

    /// @notice The configuration for the deposit fee market.
    ///         Used by the OptimismPortal to meter the cost of buying L2 gas on L1.
    ///         Set as internal with a getter so that the struct is returned instead of a tuple.
    ResourceMetering.ResourceConfig internal _resourceConfig;

    /// @notice Emitted when configuration is updated.
    /// @param version    SystemConfig version.
    /// @param updateType Type of update.
    /// @param data       Encoded update data.
    event ConfigUpdate(uint256 indexed version, UpdateType indexed updateType, bytes data);

    /// @notice The block at which the op-node can start searching for
    ///         logs from.
    uint256 public startBlock;

    /// @custom:semver 1.4.0
    /// @notice Constructs the SystemConfig contract.
    constructor() Semver(1, 4, 0) {
        initialize({
            _owner: address(0),
            _overhead: 0,
            _scalar: 0,
            _batcherHash: bytes32(0),
            _gasLimit: 0,
            _unsafeBlockSigner: address(0),
            _config: ResourceMetering.ResourceConfig({
                maxResourceLimit: 0,
                elasticityMultiplier: 0,
                baseFeeMaxChangeDenominator: 0,
                minimumBaseFee: 0,
                systemTxMaxGas: 0,
                maximumBaseFee: 0
            }),
            _startBlock: 0,
            _batchInbox: address(0),
            _addresses: SystemConfig.Addresses({
                l1CrossDomainMessenger: address(0),
                l1ERC721Bridge: address(0),
                l1StandardBridge: address(0),
                l2OutputOracle: address(0),
                optimismPortal: address(0)
            })
        });
    }

    /// @notice Initializer.
    ///         The resource config must be set before the require check.
    /// @param _owner             Initial owner of the contract.
    /// @param _overhead          Initial overhead value.
    /// @param _scalar            Initial scalar value.
    /// @param _batcherHash       Initial batcher hash.
    /// @param _gasLimit          Initial gas limit.
    /// @param _unsafeBlockSigner Initial unsafe block signer address.
    /// @param _config            Initial ResourceConfig.
    /// @param _startBlock        Starting block for the op-node to search for logs from.
    ///                           Contracts that were deployed before this field existed
    ///                           need to have this field set manually via an override.
    ///                           Newly deployed contracts should set this value to uint256(0).
    /// @param _batchInbox        Batch inbox address. An identifier for the op-node to find
    ///                           canonical data.
    /// @param _addresses         Set of L1 contract addresses. These should be the proxies.
    function initialize(
        address _owner,
        uint256 _overhead,
        uint256 _scalar,
        bytes32 _batcherHash,
        uint64 _gasLimit,
        address _unsafeBlockSigner,
        ResourceMetering.ResourceConfig memory _config,
        uint256 _startBlock,
        address _batchInbox,
        SystemConfig.Addresses memory _addresses
    ) public reinitializer(2) {
        __Ownable_init();
        transferOwnership(_owner);

        overhead = _overhead;
        scalar = _scalar;
        batcherHash = _batcherHash;
        gasLimit = _gasLimit;

        _setAddress(_unsafeBlockSigner, UNSAFE_BLOCK_SIGNER_SLOT);
        _setAddress(_batchInbox, BATCH_INBOX_SLOT);
        _setAddress(_addresses.l1CrossDomainMessenger, L1_CROSS_DOMAIN_MESSENGER_SLOT);
        _setAddress(_addresses.l1ERC721Bridge, L1_ERC_721_BRIDGE_SLOT);
        _setAddress(_addresses.l1StandardBridge, L1_STANDARD_BRIDGE_SLOT);
        _setAddress(_addresses.l2OutputOracle, L2_OUTPUT_ORACLE_SLOT);
        _setAddress(_addresses.optimismPortal, OPTIMISM_PORTAL_SLOT);

        // The start block for the op-node to start searching for logs
        // needs to be set in a backwards compatible way. Only allow setting
        // the start block with an override if it has previously never been set.
        if (_startBlock != 0) {
            require(startBlock == 0, "SystemConfig: cannot override an already set start block");
            startBlock = _startBlock;
        } else if (startBlock == 0) {
            startBlock = block.number;
        }

        _setResourceConfig(_config);
        require(_gasLimit >= minimumGasLimit(), "SystemConfig: gas limit too low");
    }

    /// @notice Returns the minimum L2 gas limit that can be safely set for the system to
    ///         operate. The L2 gas limit must be larger than or equal to the amount of
    ///         gas that is allocated for deposits per block plus the amount of gas that
    ///         is allocated for the system transaction.
    ///         This function is used to determine if changes to parameters are safe.
    /// @return uint64 Minimum gas limit.
    function minimumGasLimit() public view returns (uint64) {
        return uint64(_resourceConfig.maxResourceLimit) + uint64(_resourceConfig.systemTxMaxGas);
    }

    /// @notice High level getter for the unsafe block signer address.
    ///         Unsafe blocks can be propagated across the p2p network if they are signed by the
    ///         key corresponding to this address.
    /// @return Address of the unsafe block signer.
    // solhint-disable-next-line ordering
    function unsafeBlockSigner() external view returns (address) {
        address addr;
        bytes32 slot = UNSAFE_BLOCK_SIGNER_SLOT;
        assembly {
            addr := sload(slot)
        }
        return addr;
    }

    /// @notice Stores an address in an arbitrary storage slot, `_slot`.
    /// @param _addr The address to store
    /// @param _slot The storage slot to store the address in.
    /// @dev WARNING! This function must be used cautiously, as it allows for overwriting values
    ///      in arbitrary storage slots. Solc will add checks that the data passed as `_addr`
    ///      is 20 bytes or less.
    function _setAddress(address _addr, bytes32 _slot) internal {
        bytes32 slot = _slot;
        assembly {
            sstore(slot, _addr)
        }
    }

    /// @notice
    function _getAddress(bytes32 _slot) internal view returns (address) {
        address addr;
        bytes32 slot = _slot;
        assembly {
            addr := sload(slot)
        }
        return addr;
    }

    /// @notice Getter for the L1CrossDomainMessenger address.
    function l1CrossDomainMessenger() external view returns (address) {
        return _getAddress(L1_CROSS_DOMAIN_MESSENGER_SLOT);
    }

    /// @notice Getter for the L1ERC721Bridge address.
    function l1ERC721Bridge() external view returns (address) {
        return _getAddress(L1_ERC_721_BRIDGE_SLOT);
    }

    /// @notice Getter for the L1StandardBridge address.
    function l1StandardBridge() external view returns (address) {
        return _getAddress(L1_STANDARD_BRIDGE_SLOT);
    }

    /// @notice Getter for the L2OutputOracle address.
    function l2OutputOracle() external view returns (address) {
        return _getAddress(L2_OUTPUT_ORACLE_SLOT);
    }

    /// @notice Getter for the OptimismPortal address.
    function optimismPortal() external view returns (address) {
        return _getAddress(OPTIMISM_PORTAL_SLOT);
    }

    /// @notice Getter for the BatchInbox address.
    function batchInbox() external view returns (address) {
        return _getAddress(BATCH_INBOX_SLOT);
    }

    /// @notice Updates the unsafe block signer address.
    /// @param _unsafeBlockSigner New unsafe block signer address.
    function setUnsafeBlockSigner(address _unsafeBlockSigner) external onlyOwner {
        _setAddress(_unsafeBlockSigner, UNSAFE_BLOCK_SIGNER_SLOT);

        bytes memory data = abi.encode(_unsafeBlockSigner);
        emit ConfigUpdate(VERSION, UpdateType.UNSAFE_BLOCK_SIGNER, data);
    }

    /// @notice Updates the batcher hash.
    /// @param _batcherHash New batcher hash.
    function setBatcherHash(bytes32 _batcherHash) external onlyOwner {
        batcherHash = _batcherHash;

        bytes memory data = abi.encode(_batcherHash);
        emit ConfigUpdate(VERSION, UpdateType.BATCHER, data);
    }

    /// @notice Updates gas config.
    /// @param _overhead New overhead value.
    /// @param _scalar   New scalar value.
    function setGasConfig(uint256 _overhead, uint256 _scalar) external onlyOwner {
        overhead = _overhead;
        scalar = _scalar;

        bytes memory data = abi.encode(_overhead, _scalar);
        emit ConfigUpdate(VERSION, UpdateType.GAS_CONFIG, data);
    }

    /// @notice Updates the L2 gas limit.
    /// @param _gasLimit New gas limit.
    function setGasLimit(uint64 _gasLimit) external onlyOwner {
        require(_gasLimit >= minimumGasLimit(), "SystemConfig: gas limit too low");
        gasLimit = _gasLimit;

        bytes memory data = abi.encode(_gasLimit);
        emit ConfigUpdate(VERSION, UpdateType.GAS_LIMIT, data);
    }

    /// @notice A getter for the resource config.
    ///         Ensures that the struct is returned instead of a tuple.
    /// @return ResourceConfig
    function resourceConfig() external view returns (ResourceMetering.ResourceConfig memory) {
        return _resourceConfig;
    }

    /// @notice An external setter for the resource config.
    ///         In the future, this method may emit an event that the `op-node` picks up
    ///         for when the resource config is changed.
    /// @param _config The new resource config values.
    function setResourceConfig(ResourceMetering.ResourceConfig memory _config) external onlyOwner {
        _setResourceConfig(_config);
    }

    /// @notice An internal setter for the resource config.
    ///         Ensures that the config is sane before storing it by checking for invariants.
    /// @param _config The new resource config.
    function _setResourceConfig(ResourceMetering.ResourceConfig memory _config) internal {
        // Min base fee must be less than or equal to max base fee.
        require(
            _config.minimumBaseFee <= _config.maximumBaseFee,
            "SystemConfig: min base fee must be less than max base"
        );
        // Base fee change denominator must be greater than 1.
        require(
            _config.baseFeeMaxChangeDenominator > 1,
            "SystemConfig: denominator must be larger than 1"
        );
        // Max resource limit plus system tx gas must be less than or equal to the L2 gas limit.
        // The gas limit must be increased before these values can be increased.
        require(
            _config.maxResourceLimit + _config.systemTxMaxGas <= gasLimit,
            "SystemConfig: gas limit too low"
        );
        // Elasticity multiplier must be greater than 0.
        require(
            _config.elasticityMultiplier > 0,
            "SystemConfig: elasticity multiplier cannot be 0"
        );
        // No precision loss when computing target resource limit.
        require(
            ((_config.maxResourceLimit / _config.elasticityMultiplier) *
                _config.elasticityMultiplier) == _config.maxResourceLimit,
            "SystemConfig: precision loss with target resource limit"
        );

        _resourceConfig = _config;
    }
}
