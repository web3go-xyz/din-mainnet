import { expect } from 'chai'

import { Direction, waitForXDomainTransaction } from './watcher-utils'

import {
  getContractFactory,
  getContractInterface,
  predeploys,
} from '@eth-optimism/contracts'
import { remove0x, Watcher } from '@eth-optimism/core-utils'
import {
  Contract,
  Wallet,
  constants,
  providers,
  BigNumberish,
  BigNumber,
  utils,
} from 'ethers'
import { cleanEnv, str, num } from 'envalid'

export const GWEI = BigNumber.from(1e9)

const env = cleanEnv(process.env, {
  L1_URL: str({ default: 'http://localhost:9545' }),
  L2_URL: str({ default: 'http://localhost:8545' }),
  VERIFIER_URL: str({ default: 'http://localhost:8547' }),
  L1_POLLING_INTERVAL: num({ default: 10 }),
  L2_POLLING_INTERVAL: num({ default: 10 }),
  VERIFIER_POLLING_INTERVAL: num({ default: 10 }),
  PRIVATE_KEY: str({
    default:
      '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
  }),
  ADDRESS_MANAGER: str({
    default: '0x5FbDB2315678afecb367f032d93F642f64180aa3',
  }),
})

// The hardhat instance
export const l1Provider = new providers.JsonRpcProvider(env.L1_URL)
l1Provider.pollingInterval = env.L1_POLLING_INTERVAL

export const l2Provider = new providers.JsonRpcProvider(env.L2_URL)
l2Provider.pollingInterval = env.L2_POLLING_INTERVAL

export const verifierProvider = new providers.JsonRpcProvider(env.VERIFIER_URL)
verifierProvider.pollingInterval = env.VERIFIER_POLLING_INTERVAL

// The sequencer private key which is funded on L1
export const l1Wallet = new Wallet(env.PRIVATE_KEY, l1Provider)

// A random private key which should always be funded with deposits from L1 -> L2
// if it's using non-0 gas price
export const l2Wallet = l1Wallet.connect(l2Provider)

// Predeploys
export const PROXY_SEQUENCER_ENTRYPOINT_ADDRESS =
  '0x4200000000000000000000000000000000000004'
export const OVM_ETH_ADDRESS = predeploys.OVM_ETH

export const getAddressManager = (provider: any) => {
  return getContractFactory('Lib_AddressManager')
    .connect(provider)
    .attach(env.ADDRESS_MANAGER)
}

// Gets the bridge contract
export const getL1Bridge = async (wallet: Wallet, AddressManager: Contract) => {
  const l1BridgeInterface = getContractInterface('OVM_L1StandardBridge')
  const ProxyBridgeAddress = await AddressManager.getAddress(
    'Proxy__OVM_L1StandardBridge'
  )

  if (
    !utils.isAddress(ProxyBridgeAddress) ||
    ProxyBridgeAddress === constants.AddressZero
  ) {
    throw new Error('Proxy__OVM_L1StandardBridge not found')
  }

  const OVM_L1StandardBridge = new Contract(
    ProxyBridgeAddress,
    l1BridgeInterface,
    wallet
  )
  return OVM_L1StandardBridge
}

export const getL2Bridge = async (wallet: Wallet) => {
  const L2BridgeInterface = getContractInterface('OVM_L2StandardBridge')

  const OVM_L2StandardBridge = new Contract(
    predeploys.OVM_L2StandardBridge,
    L2BridgeInterface,
    wallet
  )
  return OVM_L2StandardBridge
}

export const getOvmEth = (wallet: Wallet) => {
  const OVM_ETH = new Contract(
    OVM_ETH_ADDRESS,
    getContractInterface('OVM_ETH'),
    wallet
  )

  return OVM_ETH
}

export const fundUser = async (
  watcher: Watcher,
  bridge: Contract,
  amount: BigNumberish,
  recipient?: string
) => {
  const value = BigNumber.from(amount)
  const tx = recipient
    ? bridge.depositETHTo(recipient, 1_300_000, '0x', { value })
    : bridge.depositETH(1_300_000, '0x', { value })

  await waitForXDomainTransaction(watcher, tx, Direction.L1ToL2)
}

export const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms))

const abiCoder = new utils.AbiCoder()
export const encodeSolidityRevertMessage = (_reason: string): string => {
  return '0x08c379a0' + remove0x(abiCoder.encode(['string'], [_reason]))
}

export const DEFAULT_TRANSACTION = {
  to: '0x' + '1234'.repeat(10),
  gasLimit: 33600000000001,
  gasPrice: 0,
  data: '0x',
  value: 0,
}

export const expectApprox = (
  actual: BigNumber | number,
  target: BigNumber | number,
  upperDeviation: number,
  lowerDeviation: number = 100
) => {
  actual = BigNumber.from(actual)
  target = BigNumber.from(target)

  const validDeviations =
    upperDeviation >= 0 &&
    upperDeviation <= 100 &&
    lowerDeviation >= 0 &&
    lowerDeviation <= 100
  if (!validDeviations) {
    throw new Error(
      'Upper and lower deviation percentage arguments should be between 0 and 100'
    )
  }
  const upper = target.mul(100 + upperDeviation).div(100)
  const lower = target.mul(100 - lowerDeviation).div(100)

  expect(
    actual.lte(upper),
    `Actual value is more than ${upperDeviation}% greater than target`
  ).to.be.true
  expect(
    actual.gte(lower),
    `Actual value is more than ${lowerDeviation}% less than target`
  ).to.be.true
}
