// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { console } from "forge-std/console.sol";
import { SafeBuilder } from "../universal/SafeBuilder.sol";
import { IMulticall3 } from "forge-std/interfaces/IMulticall3.sol";
import { IGnosisSafe, Enum } from "../interfaces/IGnosisSafe.sol";
import { LibSort } from "../libraries/LibSort.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { Constants } from "src/libraries/Constants.sol";
import { SystemConfig } from "src/L1/SystemConfig.sol";
import { ResourceMetering } from "src/L1/ResourceMetering.sol";
import { Semver } from "src/universal/Semver.sol";

import { DeployConfig } from "scripts/DeployConfig.s.sol";
import { SystemConfig } from "src/L1/SystemConfig.sol";
import { L1StandardBridge } from "src/L1/L1StandardBridge.sol";
import { L2OutputOracle } from "src/L1/L2OutputOracle.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { L1CrossDomainMessenger } from "src/L1/L1CrossDomainMessenger.sol";
import { L1ERC721Bridge } from "src/L1/L1ERC721Bridge.sol";

/// @title Multichain
/// @notice Upgrade script for upgrading the L1 contracts after the multichain upgrade.
///         This upgrade involves all of the contracts having the same implementations.
///         The implementation addresses are known ahead of time.
///         Configure the network using the `NETWORK` env var, the options are:
///         `goerli-prod`, `chaosnet` and `devnet`.
contract Multichain is SafeBuilder {
    /// @notice Address of the ProxyAdmin, passed in via constructor of `run`.
    ProxyAdmin internal PROXY_ADMIN;

    /// @notice
    DeployConfig internal cfg;

    /// @notice Represents a set of L1 contracts. Used to represent a set of proxies.
    struct ContractSet {
        address L1CrossDomainMessenger;
        address L1StandardBridge;
        address L2OutputOracle;
        address OptimismMintableERC20Factory;
        address OptimismPortal;
        address SystemConfig;
        address L1ERC721Bridge;
    }

    string internal constant GOERLI_PROD = "goerli-prod";
    string internal constant CHAOSNET = "chaosnet";
    string internal constant DEVNET = "devnet";

    /// @notice
    address internal constant L1CrossDomainMessengerImplementation = 0xb5df97bB67f5AA7254d40E1B7034bBFF7F183a38;
    /// @notice
    address internal constant L1StandardBridgeImplementation = 0xd9aA10f75a2a93Bfc73AaDD41ae777e900CEdBc9;
    /// @notice
    address internal constant L2OutputOracleImplementation = 0xaBd96C062c6B640d5670455E9d1cD98383Dd23CA;
    /// @notice
    address internal constant OptimismMintableERC20FactoryImplementation = 0xE220F7D7fF39837003A1835fCefFa8bCA4098582;
    /// @notice
    address internal constant OptimismPortalImplementation = 0x345D27c7B6C90fef5beA9631037C36119f4bF93e;
    /// @notice
    address internal constant SystemConfigImplementation = 0x00CB689221540dEd0FA5247dbE7Fc66621F431a6;
    /// @notice
    address internal constant L1ERC721BridgeImplementation = 0x53C115eD8D9902f4999fDBd8B93Ea79BF37cb588;

    /// @notice A mapping of deployment name to ContractSet of proxy addresses.
    ContractSet internal proxies;

    /// @notice The expected versions for the contracts to be upgraded to.
    string internal constant L1CrossDomainMessengerVersion = "1.5.1";
    string internal constant L1StandardBridgeVersion = "1.2.1";
    string internal constant L2OutputOracleVersion = "1.4.1";
    string internal constant OptimismMintableERC20FactoryVersion = "1.1.2";
    string internal constant OptimismPortalVersion = "1.8.1";
    string internal constant SystemConfigVersion = "1.5.0";
    string internal constant L1ERC721BridgeVersion = "1.2.1";

    /// @notice
    string internal NETWORK;

    // @notice
    uint256 internal l2OutputOracleStartingTimestamp;

    /// @notice Place the contract addresses in storage so they can be used when building calldata.
    function setUp() external {
        // Set the network in storage
        NETWORK = vm.envOr("NETWORK", GOERLI_PROD);

        // For simple comparisons of dynamic types
        bytes32 network = keccak256(bytes(NETWORK));
        bytes32 goerli = keccak256(bytes(GOERLI_PROD));
        bytes32 chaosnet = keccak256(bytes(CHAOSNET));
        bytes32 devnet = keccak256(bytes(DEVNET));

        string memory deployConfigPath;
        if (network == goerli) {
            deployConfigPath = string.concat(vm.projectRoot(), "/deploy-config/goerli.json");
            proxies = ContractSet({
                L1CrossDomainMessenger: 0x5086d1eEF304eb5284A0f6720f79403b4e9bE294,
                L1StandardBridge: 0x636Af16bf2f682dD3109e60102b8E1A089FedAa8,
                L2OutputOracle: 0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0,
                OptimismMintableERC20Factory: 0x883dcF8B05364083D849D8bD226bC8Cb4c42F9C5,
                OptimismPortal: 0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383,
                SystemConfig: 0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60,
                L1ERC721Bridge: 0x8DD330DdE8D9898d43b4dc840Da27A07dF91b3c9
            });
        } else if (network == chaosnet) {
            // TODO: migrate deploy config into repo
            deployConfigPath = string.concat(vm.projectRoot(), "/deploy-config/chaosnet.json");
        } else if (network == devnet) {
            deployConfigPath = string.concat(vm.projectRoot(), "/deploy-config/internal-devnet.json");
            proxies = ContractSet({
                L1CrossDomainMessenger: 0x71A046D793C71af209960DCb8bD5388d2c5D2a78,
                L1StandardBridge: 0x791590936abB3531c9d54CD10CEC4B14415B0Ba7,
                L2OutputOracle: 0xA57B9f15AA204b9D68DF58849b02Df16c80C0999,
                OptimismMintableERC20Factory: 0x2998dDaF6AaA00Fa8A7104214688d29Bc749B78F,
                OptimismPortal: 0x83a242F8481D4F8a605Aa82BA9D42BA574054258,
                SystemConfig: 0x9673B370b773D9FFc4C282B3aF315ab05f17E20C,
                L1ERC721Bridge: 0x60e0f047FfA2Fad2AF197c9CDb404bDF47277Ed4
            });
        } else {
            revert("Invalid network");
        }

        cfg = new DeployConfig(deployConfigPath);

        l2OutputOracleStartingTimestamp = cfg.l2OutputOracleStartingTimestamp();
    }

    /// @notice Follow up assertions to ensure that the script ran to completion.
    function _postCheck() internal view override {
        ContractSet memory prox = getProxies();
        require(
            _versionHash(prox.L1CrossDomainMessenger) == keccak256(bytes(L1CrossDomainMessengerVersion)),
            "L1CrossDomainMessenger"
        );
        require(_versionHash(prox.L1StandardBridge) == keccak256(bytes(L1StandardBridgeVersion)), "L1StandardBridge");
        require(_versionHash(prox.L2OutputOracle) == keccak256(bytes(L2OutputOracleVersion)), "L2OutputOracle");
        require(
            _versionHash(prox.OptimismMintableERC20Factory) == keccak256(bytes(OptimismMintableERC20FactoryVersion)),
            "OptimismMintableERC20Factory"
        );
        require(_versionHash(prox.OptimismPortal) == keccak256(bytes(OptimismPortalVersion)), "OptimismPortal");
        require(_versionHash(prox.SystemConfig) == keccak256(bytes(SystemConfigVersion)), "SystemConfig");
        require(_versionHash(prox.L1ERC721Bridge) == keccak256(bytes(L1ERC721BridgeVersion)), "L1ERC721Bridge");

        ResourceMetering.ResourceConfig memory rcfg = SystemConfig(prox.SystemConfig).resourceConfig();
        ResourceMetering.ResourceConfig memory dflt = Constants.DEFAULT_RESOURCE_CONFIG();
        require(keccak256(abi.encode(rcfg)) == keccak256(abi.encode(dflt)));

        // Check that the codehashes of all implementations match the proxies set implementations.
        require(
            PROXY_ADMIN.getProxyImplementation(prox.L1CrossDomainMessenger).codehash
                == L1CrossDomainMessengerImplementation.codehash,
            "L1CrossDomainMessenger codehash"
        );
        require(
            PROXY_ADMIN.getProxyImplementation(prox.L1StandardBridge).codehash == L1StandardBridgeImplementation.codehash,
            "L1StandardBridge codehash"
        );
        require(
            PROXY_ADMIN.getProxyImplementation(prox.L2OutputOracle).codehash == L2OutputOracleImplementation.codehash,
            "L2OutputOracle codehash"
        );
        require(
            PROXY_ADMIN.getProxyImplementation(prox.OptimismMintableERC20Factory).codehash
                == OptimismMintableERC20FactoryImplementation.codehash,
            "OptimismMintableERC20Factory codehash"
        );
        require(
            PROXY_ADMIN.getProxyImplementation(prox.OptimismPortal).codehash == OptimismPortalImplementation.codehash,
            "OptimismPortal codehash"
        );
        require(
            PROXY_ADMIN.getProxyImplementation(prox.SystemConfig).codehash == SystemConfigImplementation.codehash,
            "SystemConfig codehash"
        );
        require(
            PROXY_ADMIN.getProxyImplementation(prox.L1ERC721Bridge).codehash == L1ERC721BridgeImplementation.codehash,
            "L1ERC721Bridge codehash"
        );
    }

    /// @notice Test coverage of the logic. Should only run on goerli but other chains
    ///         could be added.
    function test_script_succeeds() external skipWhenNotForking {
        address _safe;
        address _proxyAdmin;

        if (block.chainid == GOERLI) {
            _safe = 0xBc1233d0C3e6B5d53Ab455cF65A6623F6dCd7e4f;
            _proxyAdmin = 0x01d3670863c3F4b24D7b107900f0b75d4BbC6e0d;
            // Set the proxy admin for the `_postCheck` function
            PROXY_ADMIN = ProxyAdmin(_proxyAdmin);
        }

        require(_safe != address(0) && _proxyAdmin != address(0));

        address[] memory owners = IGnosisSafe(payable(_safe)).getOwners();

        for (uint256 i; i < owners.length; i++) {
            address owner = owners[i];
            vm.startBroadcast(owner);
            bool success = _run(_safe, _proxyAdmin);
            vm.stopBroadcast();

            if (success) {
                console.log("tx success");
                break;
            }
        }

        _postCheck();
    }

    /// @notice Builds the calldata that the multisig needs to make for the upgrade to happen.
    ///         A total of 8 calls are made, 7 upgrade implementations and 1 sets the resource
    ///         config to the default value in the SystemConfig contract.
    function buildCalldata(address _proxyAdmin) internal view override returns (bytes memory) {
        IMulticall3.Call3[] memory calls = new IMulticall3.Call3[](7);

        ContractSet memory prox = getProxies();

        // Upgrade the L1CrossDomainMessenger
        calls[0] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgradeAndCall,
                (
                    payable(prox.L1CrossDomainMessenger), // proxy
                    L1CrossDomainMessengerImplementation, // implementation
                    abi.encodeCall(                       // data
                        L1CrossDomainMessenger.initialize,
                        (
                            OptimismPortal(payable(prox.OptimismPortal))
                        )
                    )
                )
            )
        });

        // Upgrade the L1StandardBridge
        calls[1] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgradeAndCall,
                (
                    payable(prox.L1StandardBridge),  // proxy
                    L1StandardBridgeImplementation,  // implementation
                    abi.encodeCall(
                        L1StandardBridge.initialize,
                        (
                            L1CrossDomainMessenger(prox.L1CrossDomainMessenger)
                        )
                    )
                )
            )
        });

        // Upgrade the L2OutputOracle
        calls[2] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgradeAndCall,
                (
                    payable(prox.L2OutputOracle),   // proxy
                    L2OutputOracleImplementation,   // implementation
                    abi.encodeCall(                 // data
                        L2OutputOracle.initialize,
                        (
                            cfg.l2OutputOracleStartingBlockNumber(),
                            l2OutputOracleStartingTimestamp,
                            cfg.l2OutputOracleProposer(),
                            cfg.l2OutputOracleChallenger()
                        )
                    )
                )
            )
        });

        // Upgrade the OptimismMintableERC20Factory. No initialize function.
        calls[3] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgrade, (payable(prox.OptimismMintableERC20Factory), OptimismMintableERC20FactoryImplementation)
            )
        });

        // Upgrade the OptimismPortal
        calls[4] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgradeAndCall,
                (
                    payable(prox.OptimismPortal),   // proxy
                    OptimismPortalImplementation,   // implementation
                    abi.encodeCall(                 // data
                        OptimismPortal.initialize,
                        (
                            L2OutputOracle(prox.L2OutputOracle),
                            cfg.portalGuardian(),
                            SystemConfig(prox.SystemConfig),
                            false
                        )
                    )
                )
            )
        });

        // Upgrade the SystemConfig
        calls[5] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgradeAndCall,
                (
                    payable(prox.SystemConfig),   // proxy
                    SystemConfigImplementation,   // implementation
                    abi.encodeCall(               // data
                        SystemConfig.initialize,
                        (
                            cfg.finalSystemOwner(),
                            cfg.gasPriceOracleOverhead(),
                            cfg.gasPriceOracleScalar(),
                            bytes32(uint256(uint160(cfg.batchSenderAddress()))),
                            uint64(cfg.l2GenesisBlockGasLimit()),
                            cfg.p2pSequencerAddress(),
                            Constants.DEFAULT_RESOURCE_CONFIG(),
                            cfg.systemConfigStartBlock(),
                            cfg.batchInboxAddress(),
                            SystemConfig.Addresses({
                                l1CrossDomainMessenger: prox.L1CrossDomainMessenger,
                                l1ERC721Bridge: prox.L1ERC721Bridge,
                                l1StandardBridge: prox.L1StandardBridge,
                                l2OutputOracle: prox.L2OutputOracle,
                                optimismPortal: prox.OptimismPortal,
                                optimismMintableERC20Factory: prox.OptimismMintableERC20Factory
                            })
                        )
                    )
                )
            )
        });

        // Upgrade the L1ERC721Bridge
        calls[6] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgradeAndCall,
                (
                    payable(prox.L1ERC721Bridge),
                    L1ERC721BridgeImplementation,
                    abi.encodeCall(
                        L1ERC721Bridge.initialize,
                        (
                            L1CrossDomainMessenger(prox.L1CrossDomainMessenger)
                        )
                    )
                )
            )
        });

        return abi.encodeCall(IMulticall3.aggregate3, (calls));
    }

    /// @notice Returns the ContractSet that represents the proxies for a given network.
    ///         Configure the network with the NETWORK env var.
    function getProxies() internal view returns (ContractSet memory) {
        require(proxies.L1CrossDomainMessenger != address(0), "no proxies for this network");
        return proxies;
    }
}
