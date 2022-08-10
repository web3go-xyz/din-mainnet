import fs from 'fs'
import path from 'path'
import assert from 'assert'

import { OptimismGenesis, State } from '@eth-optimism/core-utils'
import 'hardhat-deploy'
import '@eth-optimism/hardhat-deploy-config'
import { ethers, utils, BigNumber } from 'ethers'
import { task } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import {
  CompilerOutputSource,
  CompilerOutputContract,
  BuildInfo,
} from 'hardhat/types/artifacts'

import { predeploys } from '../src'

const { hexZeroPad, hexConcat, hexDataSlice, getAddress } = utils

const prefix = '0x420000000000000000000000000000000000'
const implementationSlot =
  '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'
const adminSlot =
  '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103'

const toCodeAddr = (addr: string) => {
  const address = hexConcat([
    '0xc0d3c0d3c0d3c0d3c0d3c0d3c0d3c0d3c0d3',
    '0x' + addr.slice(prefix.length),
  ])
  return getAddress(address)
}

const toBytes32 = (num: number): string => {
  const big = BigNumber.from(num)
  return hexZeroPad(big.toHexString(), 32)
}

const assertEvenLength = (str: string) => {
  assert(str.length % 2 === 0, str)
}

// TODO: fix this to be compatible with smock's version
const getStorageLayout = async (
  hre: HardhatRuntimeEnvironment,
  name: string
) => {
  const buildInfo = await hre.artifacts.getBuildInfo(name)
  const keys = Object.keys(buildInfo.output.contracts)
  for (const key of keys) {
    if (name === path.basename(key, '.sol')) {
      const contract = buildInfo.output.contracts[key]
      const storageLayout = (contract[name] as any).storageLayout
      return storageLayout || { storage: [], types: {} }
    }
  }
  throw new Error(`Cannot locate storageLayout for ${name}`)
}

// Find the contract and source from the build info
const findContractAndSource = (name: string, buildInfo: BuildInfo) => {
  const sources = buildInfo.output.sources
  const contracts = buildInfo.output.contracts

  const compilerOutputContracts: CompilerOutputContract[] = []
  for (const [contractName, contract] of Object.entries(contracts)) {
    if (path.basename(contractName, '.sol') === name) {
      compilerOutputContracts.push(contract[name])
    }
  }
  if (compilerOutputContracts.length === 0) {
    throw new Error(`Cannot find compiler output contract for ${name}`)
  }
  if (compilerOutputContracts.length !== 1) {
    console.log(`Unexpected number of contracts for ${name}`)
  }
  const outputContract = compilerOutputContracts[0]

  const compilerOutputSources: CompilerOutputSource[] = []
  for (const [contractName, source] of Object.entries(sources)) {
    if (path.basename(contractName, '.sol') === name) {
      compilerOutputSources.push(source as CompilerOutputSource)
    }
  }
  if (compilerOutputSources.length === 0) {
    throw new Error(`Cannot find compiler output source for ${name}`)
  }
  if (compilerOutputSources.length !== 1) {
    console.log(`Unexpected number of sources for ${name}`)
  }
  const outputSource = compilerOutputSources[0]

  return { outputContract, outputSource }
}

const replaceImmutables = async (
  hre: HardhatRuntimeEnvironment,
  name: string,
  immutables: object
): Promise<string> => {
  const artifact = await hre.artifacts.readArtifact(name)
  const buildInfo = await hre.artifacts.getBuildInfo(name)

  const { outputContract, outputSource } = findContractAndSource(
    name,
    buildInfo
  )

  // Get the immutable references. They look like this:
  // { ast-id: [ {start, length} ] }
  const immutableReferences =
    outputContract.evm.deployedBytecode.immutableReferences

  const names = {}
  // Recursively find all of the immutables by traversing the solc output ast
  const findNames = (ast: any) => {
    // Add the name of the variable if it is an immutable
    const isImmutable = ast.mutability === 'immutable'
    const isASTNode = typeof ast.name === 'string' && typeof ast.id === 'number'
    if (isASTNode && isImmutable) {
      names[ast.name] = ast.id
    }
    // Iterate over each node
    if (Array.isArray(ast.nodes)) {
      for (const node of ast.nodes) {
        findNames(node)
      }
    }
    // Handle contracts that are inherited from
    if (Array.isArray(ast.baseContracts)) {
      for (const baseContract of ast.baseContracts) {
        if (baseContract.baseName) {
          const base = findContractAndSource(
            baseContract.baseName.name,
            buildInfo
          )
          findNames(base.outputSource.ast)
        }
      }
    }
  }

  findNames(outputSource.ast)

  let deployedBytecode = artifact.deployedBytecode
  const presize = deployedBytecode.length

  // For each of the immutables, put the value into the bytecode
  for (const [key, value] of Object.entries(immutables)) {
    const astId = names[key]
    if (!astId) {
      throw new Error(`Unknown immutable ${key} in contract ${name}`)
    }
    const offsets = immutableReferences[astId]
    if (!offsets) {
      throw new Error(`Unknown AST id ${astId} in contract ${name}`)
    }

    // Insert the value at each one
    for (const offset of offsets) {
      if (offset.length !== 32) {
        throw new Error(
          `Immutable slicing must be updated to handle arbitrary size immutables`
        )
      }

      // Ensure that the value being sliced out is 0
      const val = hexDataSlice(
        deployedBytecode,
        offset.start,
        offset.start + offset.length
      )
      if (!BigNumber.from(val).eq(0)) {
        throw new Error(`Unexpected value in immutable bytecode ${val}`)
      }

      deployedBytecode = ethers.utils.hexConcat([
        hexDataSlice(deployedBytecode, 0, offset.start),
        hexZeroPad(value, 32),
        hexDataSlice(deployedBytecode, offset.start + offset.length),
      ])
    }
  }

  // Ensure that the bytecode is the same size
  if (presize !== deployedBytecode.length) {
    throw new Error(
      `Size mismatch! Before ${presize}, after ${deployedBytecode.length}`
    )
  }

  return deployedBytecode
}

task('genesis-l2', 'create a genesis config')
  .addOptionalParam(
    'outfile',
    'The file to write the output JSON to',
    'genesis.json'
  )
  .addOptionalParam('l1RpcUrl', 'The L1 RPC URL', 'http://127.0.0.1:8545')
  .setAction(async (args, hre) => {
    const {
      computeStorageSlots,
      // eslint-disable-next-line @typescript-eslint/no-var-requires
    } = require('@defi-wonderland/smock/dist/src/utils')

    const { deployConfig } = hre

    const l1 = new ethers.providers.StaticJsonRpcProvider(args.l1RpcUrl)

    const l1StartingBlock = await l1.getBlock(deployConfig.l1StartingBlockTag)
    if (l1StartingBlock === null) {
      throw new Error(
        `Cannot fetch block tag ${deployConfig.l1StartingBlockTag}`
      )
    }

    if (l1StartingBlock === null) {
      console.log(`Unable to fetch L1 block that rollup starts at`)
    }

    // Use the addresses of the proxies here instead of the implementations
    // Be backwards compatible
    let ProxyL1CrossDomainMessenger = await hre.deployments.getOrNull(
      'Proxy__OVM_L1CrossDomainMessenger'
    )
    if (ProxyL1CrossDomainMessenger === undefined) {
      ProxyL1CrossDomainMessenger = await hre.deployments.get(
        'L1CrossDomainMessengerProxy'
      )
    }
    // Be backwards compatible
    let ProxyL1StandardBridge = await hre.deployments.getOrNull(
      'Proxy__OVM_L1StandardBridge'
    )
    if (ProxyL1StandardBridge === undefined) {
      ProxyL1StandardBridge = await hre.deployments.get('L1StandardBridgeProxy')
    }

    const variables = {
      L2ToL1MessagePasser: {
        nonce: 0,
      },
      L2CrossDomainMessenger: {
        _initialized: 1,
        _owner: deployConfig.l2CrossDomainMessengerOwner,
        xDomainMsgSender: '0x000000000000000000000000000000000000dEaD',
        msgNonce: 0,
        otherMessenger: ProxyL1CrossDomainMessenger.address,
        // TODO: handle blockedSystemAddresses mapping
        // blockedSystemAddresses: [{key: '', value: ''}],
      },
      GasPriceOracle: {
        _owner: deployConfig.gasPriceOracleOwner,
        overhead: deployConfig.gasPriceOracleOverhead,
        scalar: deployConfig.gasPriceOracleScalar,
        decimals: deployConfig.gasPriceOracleDecimals,
      },
      L2StandardBridge: {
        messenger: predeploys.L2CrossDomainMessenger,
        otherBridge: ProxyL1StandardBridge.address,
      },
      SequencerFeeVault: {
        l1FeeWallet: ethers.constants.AddressZero,
      },
      L1Block: {
        number: l1StartingBlock.number,
        timestamp: l1StartingBlock.timestamp,
        basefee: l1StartingBlock.baseFeePerGas,
        hash: l1StartingBlock.hash,
        sequenceNumber: 0,
      },
      LegacyERC20ETH: {
        bridge: predeploys.L2StandardBridge,
        remoteToken: ethers.constants.AddressZero,
        _name: 'Ether',
        _symbol: 'ETH',
      },
      WETH9: {
        name: 'Wrapped Ether',
        symbol: 'WETH',
        decimals: 18,
      },
      GovernanceToken: {
        name: 'Optimism',
        symbol: 'OP',
        _owner: deployConfig.proxyAdmin,
      },
    }

    assertEvenLength(implementationSlot)
    assertEvenLength(adminSlot)
    assertEvenLength(deployConfig.proxyAdmin)

    const predeployAddrs = new Set()
    for (const addr of Object.values(predeploys)) {
      predeployAddrs.add(getAddress(addr))
    }

    const alloc: State = {}

    // Set a proxy at each predeploy address
    const proxy = await hre.artifacts.readArtifact('Proxy')
    for (let i = 0; i <= 2048; i++) {
      const num = hexZeroPad('0x' + i.toString(16), 2)
      const addr = getAddress(ethers.utils.hexConcat([prefix, num]))

      // There is no proxy at LegacyERC20ETH or the GovernanceToken
      if (
        addr === getAddress(predeploys.LegacyERC20ETH) ||
        addr === getAddress(predeploys.GovernanceToken)
      ) {
        continue
      }

      alloc[addr] = {
        nonce: '0x0',
        balance: '0x0',
        code: proxy.deployedBytecode,
        storage: {
          [adminSlot]: deployConfig.proxyAdmin,
        },
      }

      if (predeployAddrs.has(getAddress(addr))) {
        const predeploy = Object.entries(predeploys).find(([, address]) => {
          return getAddress(address) === addr
        })

        // Really shouldn't happen, since predeployAddrs is a set generated from predeploys.
        if (predeploy === undefined) {
          throw new Error('could not find address')
        }

        const name = predeploy[0]
        if (variables[name]) {
          const storageLayout = await getStorageLayout(hre, name)
          if (storageLayout === undefined) {
            throw new Error(`cannot find storage layout for ${name}`)
          }
          const slots = computeStorageSlots(storageLayout, variables[name])

          for (const slot of slots) {
            alloc[addr].storage[slot.key] = slot.val
          }
        }

        alloc[addr].storage[implementationSlot] = toCodeAddr(addr)
      }
    }

    // Set the GovernanceToken in the state
    // Cannot easily set storage due to no easy access to compiler
    // output
    const governanceToken = await hre.deployments.getArtifact('GovernanceToken')
    alloc[predeploys.GovernanceToken] = {
      nonce: '0x0',
      balance: '0x0',
      code: governanceToken.deployedBytecode,
    }

    // Give each predeploy a single wei
    for (let i = 0; i <= 0xff; i++) {
      const buf = Buffer.alloc(2)
      buf.writeUInt16BE(i, 0)
      const addr = ethers.utils.hexConcat([
        '0x000000000000000000000000000000000000',
        hexZeroPad(buf, 2),
      ])
      alloc[addr] = {
        balance: '0x1',
      }
    }

    if (deployConfig.fundDevAccounts) {
      const accounts = [
        '0x14dC79964da2C08b23698B3D3cc7Ca32193d9955',
        '0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65',
        '0x1CBd3b2770909D4e10f157cABC84C7264073C9Ec',
        '0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f',
        '0x2546BcD3c84621e976D8185a91A922aE77ECEc30',
        '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC',
        '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
        '0x71bE63f3384f5fb98995898A86B02Fb2426c5788',
        '0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199',
        '0x90F79bf6EB2c4f870365E785982E1f101E93b906',
        '0x976EA74026E726554dB657fA54763abd0C3a0aa9',
        '0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc',
        '0xBcd4042DE499D14e55001CcbB24a551F3b954096',
        '0xFABB0ac9d68B0B445fB7357272Ff202C5651694a',
        '0xa0Ee7A142d267C1f36714E4a8F75612F20a79720',
        '0xbDA5747bFD65F08deb54cb465eB87D40e51B197E',
        '0xcd3B766CCDd6AE721141F452C550Ca635964ce71',
        '0xdD2FD4581271e230360230F9337D5c0430Bf44C0',
        '0xdF3e18d64BC6A983f673Ab319CCaE4f1a57C7097',
        '0xde3829a23df1479438622a08a116e8eb3f620bb5',
        '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
      ]

      for (const account of accounts) {
        alloc[account] = {
          balance:
            '0x200000000000000000000000000000000000000000000000000000000000000',
        }
      }
    }

    // Note: this currently only supports up to 32 byte values.
    // Things less than 32 bytes will be left padded with 0 bytes
    const immutables = {
      OptimismMintableERC20Factory: {
        bridge: predeploys.L2StandardBridge,
      },
      GasPriceOracle: {
        MAJOR_VERSION: toBytes32(0),
        MINOR_VERSION: toBytes32(0),
        PATCH_VERSION: toBytes32(1),
      },
      L1Block: {
        MAJOR_VERSION: toBytes32(0),
        MINOR_VERSION: toBytes32(0),
        PATCH_VERSION: toBytes32(1),
      },
      L2CrossDomainMessenger: {
        MAJOR_VERSION: toBytes32(0),
        MINOR_VERSION: toBytes32(0),
        PATCH_VERSION: toBytes32(1),
      },
      L2StandardBridge: {
        MAJOR_VERSION: toBytes32(0),
        MINOR_VERSION: toBytes32(0),
        PATCH_VERSION: toBytes32(1),
      },
      L2ToL1MessagePasser: {
        MAJOR_VERSION: toBytes32(0),
        MINOR_VERSION: toBytes32(0),
        PATCH_VERSION: toBytes32(1),
      },
      SequencerFeeVault: {
        MAJOR_VERSION: toBytes32(0),
        MINOR_VERSION: toBytes32(0),
        PATCH_VERSION: toBytes32(1),
      },
    }

    // Set the predeploys in the state
    for (const [name, addr] of Object.entries(predeploys)) {
      if (name === 'GovernanceToken') {
        continue
      }
      const artifact = await hre.artifacts.readArtifact(name)
      assertEvenLength(artifact.deployedBytecode)

      const allocAddr = name === 'LegacyERC20ETH' ? addr : toCodeAddr(addr)
      assertEvenLength(allocAddr)

      const immutableConfig = immutables[name]
      const deployedBytecode = immutableConfig
        ? await replaceImmutables(hre, name, immutableConfig)
        : artifact.deployedBytecode

      assertEvenLength(deployedBytecode)

      // TODO(tynes): initialize contracts that should be initialized
      // in the implementations here
      alloc[allocAddr] = {
        nonce: '0x00',
        balance: '0x00',
        code: deployedBytecode,
        storage: {},
      }
    }

    const startingTimestamp = l1StartingBlock?.timestamp || 0

    const genesis: OptimismGenesis = {
      config: {
        chainId: deployConfig.l2ChainID,
        homesteadBlock: 0,
        eip150Block: 0,
        eip150Hash: ethers.constants.HashZero,
        eip155Block: 0,
        eip158Block: 0,
        byzantiumBlock: 0,
        constantinopleBlock: 0,
        petersburgBlock: 0,
        istanbulBlock: 0,
        muirGlacierBlock: 0,
        berlinBlock: 0,
        londonBlock: 0,
        mergeNetsplitBlock: 0,
        terminalTotalDifficulty: 0,
        optimism: {
          baseFeeRecipient: deployConfig.optimismBaseFeeRecipient,
          l1FeeRecipient: deployConfig.optimismL1FeeRecipient,
        },
      },
      nonce: deployConfig.l2GenesisBlockNonce,
      timestamp: ethers.BigNumber.from(startingTimestamp).toHexString(),
      extraData: deployConfig.l2GenesisBlockExtraData,
      gasLimit: deployConfig.l2GenesisBlockGasLimit,
      difficulty: deployConfig.l2GenesisBlockDifficulty,
      mixHash: deployConfig.l2GenesisBlockMixHash,
      coinbase: deployConfig.l2GenesisBlockCoinbase,
      number: deployConfig.l2GenesisBlockNumber,
      gasUsed: deployConfig.l2GenesisBlockGasUsed,
      parentHash: deployConfig.l2GenesisBlockParentHash,
      baseFeePerGas: deployConfig.l2GenesisBlockBaseFeePerGas,
      alloc,
    }

    fs.writeFileSync(args.outfile, JSON.stringify(genesis, null, 2))
  })
