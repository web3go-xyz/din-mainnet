/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { hexStringEquals } from '@eth-optimism/core-utils'

/* Imports: Internal */
import {
  deployAndPostDeploy,
  getLiveContract,
} from '../src/hardhat-deploy-ethers'
import { predeploys } from '../src/predeploys'

const deployFn: DeployFunction = async (hre) => {
  const Lib_AddressManager = await getLiveContract(hre, 'Lib_AddressManager')

  // ToDo: Clean up the method of mapping names to addresses esp.
  // There's probably a more functional way to generate an object or something.
  const allContractNames = [
    'ChainStorageContainer-CTC-batches',
    'ChainStorageContainer-SCC-batches',
    'CanonicalTransactionChain',
    'StateCommitmentChain',
    'BondManager',
    'OVM_L1CrossDomainMessenger',
    'Proxy__OVM_L1CrossDomainMessenger',
    'Proxy__OVM_L1StandardBridge',
  ]

  let namesAndAddresses: {
    name: string
    address: string
  }[] = await Promise.all(
    allContractNames.map(async (name) => {
      return {
        name,
        address: (await getLiveContract(hre, name)).address,
      }
    })
  )

  // Add non-deployed addresses to the Address Setter arguments.
  namesAndAddresses = [
    ...namesAndAddresses,
    // L2CrossDomainMessenger is the address of the predeploy on L2. We can refactor off-chain
    // services such that we can remove the need to set this address, but for now it's easier
    // to simply keep setting the address.
    {
      name: 'L2CrossDomainMessenger',
      address: predeploys.L2CrossDomainMessenger,
    },
    // OVM_Sequencer is the address allowed to submit "Sequencer" blocks to the
    // CanonicalTransactionChain.
    {
      name: 'OVM_Sequencer',
      address: (hre as any).deployConfig.ovmSequencerAddress,
    },
    // OVM_Proposer is the address allowed to submit state roots (transaction results) to the
    // StateCommitmentChain.
    {
      name: 'OVM_Proposer',
      address: (hre as any).deployConfig.ovmProposerAddress,
    },
  ]

  // Filter out all addresses that will not change, so that the log statement is maximally
  // verifiable and readable.
  namesAndAddresses = namesAndAddresses.filter(async ({ name, address }) => {
    const existingAddress = await Lib_AddressManager.getAddress(name)
    return !hexStringEquals(existingAddress, address)
  })

  await deployAndPostDeploy({
    hre,
    name: 'AddressSetter',
    contract: 'AddressSetter',
    args: [
      Lib_AddressManager.address,
      (hre as any).deployConfig.ovmAddressManagerOwner,
      namesAndAddresses.map((pair) => {
        return pair.name
      }),
      namesAndAddresses.map((pair) => {
        return pair.address
      }),
    ],
  })
}

deployFn.tags = ['fresh', 'upgrade', 'AddressSetter']

export default deployFn
