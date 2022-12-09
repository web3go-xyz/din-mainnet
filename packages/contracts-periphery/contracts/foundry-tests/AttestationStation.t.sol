//SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/* Testing utilities */
import { Test } from "forge-std/Test.sol";
import { TestERC20 } from "../testing/helpers/TestERC20.sol";
import { TestERC721 } from "../testing/helpers/TestERC721.sol";
import { AssetReceiver } from "../universal/AssetReceiver.sol";
import { AttestationStation } from "../universal/op-nft/AttestationStation.sol";

contract AssetReceiver_Initializer is Test {
    address alice_attestor = address(128);
    address bob = address(256);
    address sally = address(512);

    function _setUp() public {
        // Give alice and bob some ETH
        vm.deal(alice_attestor, 1 ether);

        vm.label(alice_attestor, "alice_attestor");
        vm.label(bob, "bob");
        vm.label(sally, "sally");
    }
}

contract AssetReceiverTest is AssetReceiver_Initializer {
    function setUp() public {
        super._setUp();
    }

    function test_attest_single() external {
        AttestationStation attestationStation = new AttestationStation();

        // alice is going to attest about bob
        AttestationStation.AttestationData memory attestationData = AttestationStation
            .AttestationData({
                about: bob,
                key: bytes32("test-key:string"),
                val: bytes("test-value")
            });

        // assert the attestation starts empty
        assertEq(attestationStation.attestations(address(this), address(this), "test"), "");

        // make attestation
        vm.prank(alice_attestor);
        attestationStation.attest(attestationData.about, attestationData.key, attestationData.val);

        // assert the attestation is there
        assertEq(
            attestationStation.readAttestation(
                alice_attestor,
                attestationData.about,
                attestationData.key
            ),
            attestationData.val
        );

        bytes memory new_val = bytes("new updated value");
        // make a new attestations to same about and key
        attestationData = AttestationStation.AttestationData({
            about: attestationData.about,
            key: attestationData.key,
            val: new_val
        });

        vm.prank(alice_attestor);
        attestationStation.attest(attestationData.about, attestationData.key, attestationData.val);

        // assert the attestation is updated
        assertEq(
            attestationStation.readAttestation(
                alice_attestor,
                attestationData.about,
                attestationData.key
            ),
            attestationData.val
        );
    }

    function test_attest_bulk() external {
        AttestationStation attestationStation = new AttestationStation();

        vm.prank(alice_attestor);

        AttestationStation.AttestationData[]
            memory attestationData = new AttestationStation.AttestationData[](3);
        attestationData[0] = AttestationStation.AttestationData({
            about: bob,
            key: bytes32("test-key:string"),
            val: bytes("test-value")
        });

        attestationData[1] = AttestationStation.AttestationData({
            about: bob,
            key: bytes32("test-key2"),
            val: bytes("test-value2")
        });

        attestationData[2] = AttestationStation.AttestationData({
            about: sally,
            key: bytes32("test-key:string"),
            val: bytes("test-value3")
        });

        attestationStation.attestBulk(attestationData);

        // assert the attestations are there
        assertEq(
            attestationStation.readAttestation(
                alice_attestor,
                attestationData[0].about,
                attestationData[0].key
            ),
            attestationData[0].val
        );
        assertEq(
            attestationStation.readAttestation(
                alice_attestor,
                attestationData[1].about,
                attestationData[1].key
            ),
            attestationData[1].val
        );
        assertEq(
            attestationStation.readAttestation(
                alice_attestor,
                attestationData[2].about,
                attestationData[2].key
            ),
            attestationData[2].val
        );
    }
}
