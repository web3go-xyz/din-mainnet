import { DeployFunction } from 'hardhat-deploy/dist/types'
import 'hardhat-deploy'
import '@nomiclabs/hardhat-ethers'
import '@eth-optimism/hardhat-deploy-config'

const deployFn: DeployFunction = async (hre) => {
  const { deploy } = hre.deployments
  const { deployer } = await hre.getNamedAccounts()
  const { deployConfig } = hre

  await deploy('ProxyAdmin', {
    from: deployer,
    args: [deployer],
    log: true,
    waitConfirmations: deployConfig.deploymentWaitConfirmations,
  })

  const admin = await hre.deployments.get('ProxyAdmin')
  const ProxyAdmin = await hre.ethers.getContractAt('ProxyAdmin', admin.address)

  // This is set up for fresh networks only
  const proxies = [
    'L2OutputOracleProxy',
    'L1CrossDomainMessengerProxy',
    'L1StandardBridgeProxy',
    'OptimismPortalProxy',
    'OptimismMintableERC20FactoryProxy',
  ]

  // Wait on all the txs in parallel so that the deployment goes faster
  const txs = []
  for (const proxy of proxies) {
    const deployment = await hre.deployments.get(proxy)
    const Proxy = await hre.ethers.getContractAt('Proxy', deployment.address)
    const tx = await Proxy.changeAdmin(admin.address)
    txs.push(tx)
  }
  await Promise.all(txs.map((tx) => tx.wait()))

  const addressManager = await hre.deployments.get('AddressManager')
  const AddressManager = await hre.ethers.getContractAt(
    'AddressManager',
    addressManager.address
  )

  const postConfig = [
    await AddressManager.transferOwnership(admin.address),
    await ProxyAdmin.setAddressManager(addressManager.address),
  ]
  await Promise.all(postConfig.map((tx) => tx.wait()))
}

deployFn.tags = ['ProxyAdmin']

export default deployFn
