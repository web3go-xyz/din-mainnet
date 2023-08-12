// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { OptimismMintableERC20 } from "../src/universal/OptimismMintableERC20.sol";
import { Bridge_Initializer } from "./CommonTest.t.sol";

contract OptimismMintableTokenFactory_Test is Bridge_Initializer {
    event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken);
    event OptimismMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer);

    function setUp() public override {
        super.setUp();
    }

    function test_bridge_succeeds() external {
        assertEq(address(L2TokenFactory.BRIDGE()), address(L2Bridge));
    }

    function test_createStandardL2Token_succeeds() external {
        address remote = address(4);
        address local = calculateTokenAddress(remote, "Beep", "BOOP");

        vm.expectEmit(true, true, true, true);
        emit StandardL2TokenCreated(remote, local);

        vm.expectEmit(true, true, true, true);
        emit OptimismMintableERC20Created(local, remote, alice);

        vm.prank(alice);
        L2TokenFactory.createStandardL2Token(remote, "Beep", "BOOP");
    }

    function test_createStandardL2Token_sameTwice_reverts() external {
        address remote = address(4);

        vm.prank(alice);
        L2TokenFactory.createStandardL2Token(remote, "Beep", "BOOP");

        vm.expectRevert();

        vm.prank(alice);
        L2TokenFactory.createStandardL2Token(remote, "Beep", "BOOP");
    }

    function test_createStandardL2Token_remoteIsZero_reverts() external {
        address remote = address(0);
        vm.expectRevert("OptimismMintableERC20Factory: must provide remote token address");
        L2TokenFactory.createStandardL2Token(remote, "Beep", "BOOP");
    }

    function calculateTokenAddress(
        address _remote,
        string memory _name,
        string memory _symbol
    )
        internal
        view
        returns (address)
    {
        bytes memory constructorArgs = abi.encode(address(L2Bridge), _remote, _name, _symbol);
        bytes memory bytecode = abi.encodePacked(type(OptimismMintableERC20).creationCode, constructorArgs);
        bytes32 salt = keccak256(abi.encode(_remote, _name, _symbol));
        bytes32 hash = keccak256(abi.encodePacked(bytes1(0xff), address(L2TokenFactory), salt, keccak256(bytecode)));
        return address(uint160(uint256(hash)));
    }
}
