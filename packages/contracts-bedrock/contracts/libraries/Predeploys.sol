// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title Predeploys
 * @notice Contains constant addresses for contracts that are pre-deployed to the L2 system.
 */
library Predeploys {
    /**
     * @notice Address of the L2ToL1MessagePasser predeploy.
     */
    address internal constant L2_TO_L1_MESSAGE_PASSER = 0x4200000000000000000000000000000000000000;

    /**
     * @notice Address of the L2CrossDomainMessenger predeploy.
     */
    address internal constant L2_CROSS_DOMAIN_MESSENGER =
        0x4200000000000000000000000000000000000007;

    /**
     * @notice Address of the L2StandardBridge predeploy.
     */
    address internal constant L2_STANDARD_BRIDGE = 0x4200000000000000000000000000000000000010;

    /**
     * @notice Address of the SequencerFeeWallet predeploy.
     */
    address internal constant SEQUENCER_FEE_WALLET = 0x4200000000000000000000000000000000000011;

    /**
     * @notice Address of the OptimismMintableERC20Factory predeploy.
     */
    address internal constant OPTIMISM_MINTABLE_ERC20_FACTORY =
        0x4200000000000000000000000000000000000012;

    /**
     * @notice Address of the L1Block predeploy.
     */
    address internal constant L1_BLOCK_ATTRIBUTES = 0x4200000000000000000000000000000000000015;

    /**
     * @custom:legacy
     * @notice Address of the L1MessageSender predeploy. Deprecated. Use L2CrossDomainMessenger
     *         or access tx.origin (or msg.sender) in a L1 to L2 transaction instead.
     */
    address internal constant L1_MESSAGE_SENDER = 0x4200000000000000000000000000000000000001;

    /**
     * @custom:legacy
     * @notice Address of the DeployerWhitelist predeploy. No longer active.
     */
    address internal constant DEPLOYER_WHITELIST = 0x4200000000000000000000000000000000000002;

    /**
     * @custom:legacy
     * @notice Address of the LegacyERC20ETH predeploy. Deprecated. Balances are migrated to the
     *         state trie as of the Bedrock upgrade. Contract has been locked and write functions
     *         can no longer be accessed.
     */
    address internal constant LEGACY_ERC20_ETH = 0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000;

    /**
     * @custom:legacy
     * @notice Address of the L1BlockNumber predeploy. Deprecated. Use the L1Block predeploy
     *         instead, which exposes more information about the L1 state.
     */
    address internal constant L1_BLOCK_NUMBER = 0x4200000000000000000000000000000000000013;
}
