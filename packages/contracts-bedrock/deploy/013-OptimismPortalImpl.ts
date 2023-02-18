import { DeployFunction } from 'hardhat-deploy/dist/types'
import '@eth-optimism/hardhat-deploy-config'

import {
  assertContractVariable,
  deploy,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const L2OutputOracleProxy = await getContractFromArtifact(
    hre,
    'L2OutputOracleProxy'
  )

  const finalSystemOwner = hre.deployConfig.finalSystemOwner
  const finalSystemOwnerCode = await hre.ethers.provider.getCode(
    finalSystemOwner
  )
  if (finalSystemOwnerCode === '0x') {
    console.log(
      `WARNING: setting OptimismPortal.GUARDIAN to ${finalSystemOwner} and it has no code`
    )
  }

  // Deploy the OptimismPortal implementation as paused to
  // ensure that users do not interact with it and instead
  // interact with the proxied contract.
  // The `finalSystemOwner` is set at the GUARDIAN.
  await deploy({
    hre,
    name: 'OptimismPortal',
    args: [
      L2OutputOracleProxy.address,
      finalSystemOwner,
      hre.deployConfig.finalizationPeriodSeconds,
      true, // paused
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'L2_ORACLE',
        L2OutputOracleProxy.address
      )
      await assertContractVariable(
        contract,
        'FINALIZATION_PERIOD_SECONDS',
        hre.deployConfig.finalizationPeriodSeconds
      )
      await assertContractVariable(
        contract,
        'GUARDIAN',
        hre.deployConfig.finalSystemOwner
      )
    },
  })
}

deployFn.tags = ['OptimismPortalImpl', 'setup']

export default deployFn
