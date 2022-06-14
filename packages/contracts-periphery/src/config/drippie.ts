import assert from 'assert'

import { BigNumberish, ethers } from 'ethers'
import { Interface } from 'ethers/lib/utils'
import { HardhatRuntimeEnvironment } from 'hardhat/types'

import { Etherscan } from '../etherscan'

export interface DripConfig {
  interval: BigNumberish
  dripcheck: string
  checkparams?: any
  actions: Array<{
    target: string
    value?: BigNumberish
    data?:
      | string
      | {
          fn: string
          args?: any[]
        }
  }>
}

export interface DrippieConfig {
  [name: string]: DripConfig
}

export const getDrippieConfig = async (
  hre: HardhatRuntimeEnvironment
): Promise<Required<DrippieConfig>> => {
  let config: DrippieConfig
  try {
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    config = require(`../../config/drippie/${hre.network.name}.ts`).default
  } catch (err) {
    throw new Error(
      `error while loading drippie config for network: ${hre.network.name}, ${err}`
    )
  }

  return parseDrippieConfig(hre, config)
}

export const encodeDripCheckParams = (
  iface: Interface,
  params: any
): string => {
  return ethers.utils.defaultAbiCoder.encode(
    [iface.getEvent('_EventToExposeStructInABI__Params').inputs[0]],
    [params]
  )
}

export const parseDrippieConfig = async (
  hre: HardhatRuntimeEnvironment,
  config: DrippieConfig
): Promise<Required<DrippieConfig>> => {
  // Create a clone of the config object. Shallow clone is fine because none of the input options
  // are expected to be objects or functions etc.
  const parsed = { ...config }

  const etherscan = new Etherscan(
    hre.network.config.verify.etherscan.apiKey,
    hre.network.config.chainId
  )

  for (const dripConfig of Object.values(parsed)) {
    for (const action of dripConfig.actions) {
      assert(ethers.utils.isAddress(action.target), 'target is not an address')

      if (action.data === undefined) {
        action.data = '0x'
      } else if (typeof action.data === 'string') {
        assert(
          ethers.utils.isHexString(action.data),
          'action is not a hex string'
        )
      } else {
        const abi = await etherscan.getContractABI(action.target)
        const iface = new ethers.utils.Interface(abi)
        action.data = iface.encodeFunctionData(
          action.data.fn,
          action.data.args || []
        )
      }

      if (action.value === undefined) {
        action.value = ethers.BigNumber.from(0)
      } else {
        action.value = ethers.BigNumber.from(action.value)
      }
    }

    const dripcheck = await hre.deployments.get(dripConfig.dripcheck)
    dripConfig.dripcheck = dripcheck.address

    if (dripConfig.checkparams === undefined) {
      dripConfig.checkparams = '0x'
    } else {
      dripConfig.checkparams = encodeDripCheckParams(
        new ethers.utils.Interface(dripcheck.abi),
        dripConfig.checkparams
      )
    }

    dripConfig.interval = ethers.BigNumber.from(dripConfig.interval)
  }

  return parsed as Required<DrippieConfig>
}
