import {
  getContract,
  GetContractArgs,
  readContract,
  ReadContractConfig,
  writeContract,
  WriteContractMode,
  WriteContractArgs,
  WriteContractPreparedArgs,
  WriteContractUnpreparedArgs,
  prepareWriteContract,
  PrepareWriteContractConfig,
  watchContractEvent,
  WatchContractEventConfig,
  WatchContractEventCallback,
} from 'wagmi/actions'

/* eslint-disable */

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// AddressManager
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdE1FCfB0851916CA5101820A69b13a4E276bd81F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xa6f73589243a6A7a9023b1Fa0651b1d89c177111)
 */
export const addressManagerABI = [
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: '_name', internalType: 'string', type: 'string', indexed: true },
      {
        name: '_newAddress',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
      {
        name: '_oldAddress',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'AddressSet',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_name', internalType: 'string', type: 'string' }],
    name: 'getAddress',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_address', internalType: 'address', type: 'address' },
    ],
    name: 'setAddress',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
  },
] as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdE1FCfB0851916CA5101820A69b13a4E276bd81F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xa6f73589243a6A7a9023b1Fa0651b1d89c177111)
 */
export const addressManagerAddress = {
  1: '0xdE1FCfB0851916CA5101820A69b13a4E276bd81F',
  5: '0xa6f73589243a6A7a9023b1Fa0651b1d89c177111',
} as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdE1FCfB0851916CA5101820A69b13a4E276bd81F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xa6f73589243a6A7a9023b1Fa0651b1d89c177111)
 */
export const addressManagerConfig = {
  address: addressManagerAddress,
  abi: addressManagerABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// AssetReceiver
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 */
export const assetReceiverABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_owner', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'user', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnerUpdated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'ReceivedETH',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'WithdrewERC20',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'id', internalType: 'uint256', type: 'uint256', indexed: false },
    ],
    name: 'WithdrewERC721',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'WithdrewETH',
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
      { name: '_gas', internalType: 'uint256', type: 'uint256' },
      { name: '_value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'CALL',
    outputs: [
      { name: '', internalType: 'bool', type: 'bool' },
      { name: '', internalType: 'bytes', type: 'bytes' },
    ],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
      { name: '_gas', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'DELEGATECALL',
    outputs: [
      { name: '', internalType: 'bool', type: 'bool' },
      { name: '', internalType: 'bytes', type: 'bytes' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'setOwner',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC20', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC20', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
    ],
    name: 'withdrawERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC721', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_id', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawERC721',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address payable', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawETH',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_to', internalType: 'address payable', type: 'address' }],
    name: 'withdrawETH',
    outputs: [],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 */
export const assetReceiverAddress = {
  1: '0x15DdA60616Ffca20371ED1659dBB78E888f65556',
  10: '0x15DdA60616Ffca20371ED1659dBB78E888f65556',
} as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 */
export const assetReceiverConfig = {
  address: assetReceiverAddress,
  abi: assetReceiverABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// AttestationStation
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 */
export const attestationStationABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'creator',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'about',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'key', internalType: 'bytes32', type: 'bytes32', indexed: true },
      { name: 'val', internalType: 'bytes', type: 'bytes', indexed: false },
    ],
    name: 'AttestationCreated',
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_attestations',
        internalType: 'struct AttestationStation.AttestationData[]',
        type: 'tuple[]',
        components: [
          { name: 'about', internalType: 'address', type: 'address' },
          { name: 'key', internalType: 'bytes32', type: 'bytes32' },
          { name: 'val', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'attest',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_about', internalType: 'address', type: 'address' },
      { name: '_key', internalType: 'bytes32', type: 'bytes32' },
      { name: '_val', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'attest',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'attestations',
    outputs: [{ name: '', internalType: 'bytes', type: 'bytes' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 */
export const attestationStationAddress = {
  10: '0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77',
  420: '0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77',
} as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 */
export const attestationStationConfig = {
  address: attestationStationAddress,
  abi: attestationStationABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// BaseFeeVault
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000019)
 */
export const baseFeeVaultABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_recipient', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'from',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'Withdrawal',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_WITHDRAWAL_AMOUNT',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RECIPIENT',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'totalProcessed',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'withdraw',
    outputs: [],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000019)
 */
export const baseFeeVaultAddress = {
  420: '0x4200000000000000000000000000000000000019',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000019)
 */
export const baseFeeVaultConfig = {
  address: baseFeeVaultAddress,
  abi: baseFeeVaultABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// CheckBalanceHigh
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5d7103853f12109A7d27F118e54BbC654ad847E9)
 */
export const checkBalanceHighABI = [
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'params',
        internalType: 'struct CheckBalanceHigh.Params',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'threshold', internalType: 'uint256', type: 'uint256' },
        ],
        indexed: false,
      },
    ],
    name: '_EventToExposeStructInABI__Params',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_params', internalType: 'bytes', type: 'bytes' }],
    name: 'check',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
] as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5d7103853f12109A7d27F118e54BbC654ad847E9)
 */
export const checkBalanceHighAddress = {
  1: '0x7eC64a8a591bFf829ff6C8be76074D540ACb813F',
  5: '0x7eC64a8a591bFf829ff6C8be76074D540ACb813F',
  420: '0x5d7103853f12109A7d27F118e54BbC654ad847E9',
} as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5d7103853f12109A7d27F118e54BbC654ad847E9)
 */
export const checkBalanceHighConfig = {
  address: checkBalanceHighAddress,
  abi: checkBalanceHighABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// CheckBalanceLow
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x7Ce13D154FAEE5C8B3E6b19d4Add16f21d884474)
 */
export const checkBalanceLowABI = [
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'params',
        internalType: 'struct CheckBalanceLow.Params',
        type: 'tuple',
        components: [
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'threshold', internalType: 'uint256', type: 'uint256' },
        ],
        indexed: false,
      },
    ],
    name: '_EventToExposeStructInABI__Params',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_params', internalType: 'bytes', type: 'bytes' }],
    name: 'check',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
] as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x7Ce13D154FAEE5C8B3E6b19d4Add16f21d884474)
 */
export const checkBalanceLowAddress = {
  1: '0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640',
  5: '0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640',
  420: '0x7Ce13D154FAEE5C8B3E6b19d4Add16f21d884474',
} as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x7Ce13D154FAEE5C8B3E6b19d4Add16f21d884474)
 */
export const checkBalanceLowConfig = {
  address: checkBalanceLowAddress,
  abi: checkBalanceLowABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// CheckGelatoLow
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xF9c8a4Cb4021f57F9f6d69799cA9BefF64524862)
 */
export const checkGelatoLowABI = [
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'params',
        internalType: 'struct CheckGelatoLow.Params',
        type: 'tuple',
        components: [
          { name: 'treasury', internalType: 'address', type: 'address' },
          { name: 'threshold', internalType: 'uint256', type: 'uint256' },
          { name: 'recipient', internalType: 'address', type: 'address' },
        ],
        indexed: false,
      },
    ],
    name: '_EventToExposeStructInABI__Params',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_params', internalType: 'bytes', type: 'bytes' }],
    name: 'check',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
] as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xF9c8a4Cb4021f57F9f6d69799cA9BefF64524862)
 */
export const checkGelatoLowAddress = {
  1: '0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa',
  5: '0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa',
  420: '0xF9c8a4Cb4021f57F9f6d69799cA9BefF64524862',
} as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xF9c8a4Cb4021f57F9f6d69799cA9BefF64524862)
 */
export const checkGelatoLowConfig = {
  address: checkGelatoLowAddress,
  abi: checkGelatoLowABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// CheckTrue
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5c741a38cb11424711231777D71689C458eE835D)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5c741a38cb11424711231777D71689C458eE835D)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x47443D0C184e022F19BD1578F5bca6B8a9F58E32)
 */
export const checkTrueABI = [
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes', type: 'bytes' }],
    name: 'check',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
] as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5c741a38cb11424711231777D71689C458eE835D)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5c741a38cb11424711231777D71689C458eE835D)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x47443D0C184e022F19BD1578F5bca6B8a9F58E32)
 */
export const checkTrueAddress = {
  1: '0x5c741a38cb11424711231777D71689C458eE835D',
  5: '0x5c741a38cb11424711231777D71689C458eE835D',
  420: '0x47443D0C184e022F19BD1578F5bca6B8a9F58E32',
} as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5c741a38cb11424711231777D71689C458eE835D)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5c741a38cb11424711231777D71689C458eE835D)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x47443D0C184e022F19BD1578F5bca6B8a9F58E32)
 */
export const checkTrueConfig = {
  address: checkTrueAddress,
  abi: checkTrueABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Drippie
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export const drippieABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_owner', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'nameref',
        internalType: 'string',
        type: 'string',
        indexed: true,
      },
      { name: 'name', internalType: 'string', type: 'string', indexed: false },
      {
        name: 'config',
        internalType: 'struct Drippie.DripConfig',
        type: 'tuple',
        components: [
          { name: 'reentrant', internalType: 'bool', type: 'bool' },
          { name: 'interval', internalType: 'uint256', type: 'uint256' },
          {
            name: 'dripcheck',
            internalType: 'contract IDripCheck',
            type: 'address',
          },
          { name: 'checkparams', internalType: 'bytes', type: 'bytes' },
          {
            name: 'actions',
            internalType: 'struct Drippie.DripAction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'target',
                internalType: 'address payable',
                type: 'address',
              },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
        indexed: false,
      },
    ],
    name: 'DripCreated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'nameref',
        internalType: 'string',
        type: 'string',
        indexed: true,
      },
      { name: 'name', internalType: 'string', type: 'string', indexed: false },
      {
        name: 'executor',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
      {
        name: 'timestamp',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'DripExecuted',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'nameref',
        internalType: 'string',
        type: 'string',
        indexed: true,
      },
      { name: 'name', internalType: 'string', type: 'string', indexed: false },
      {
        name: 'status',
        internalType: 'enum Drippie.DripStatus',
        type: 'uint8',
        indexed: false,
      },
    ],
    name: 'DripStatusUpdated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'user', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnerUpdated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'ReceivedETH',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'WithdrewERC20',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'id', internalType: 'uint256', type: 'uint256', indexed: false },
    ],
    name: 'WithdrewERC721',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'WithdrewETH',
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
      { name: '_value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'CALL',
    outputs: [
      { name: '', internalType: 'bool', type: 'bool' },
      { name: '', internalType: 'bytes', type: 'bytes' },
    ],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'DELEGATECALL',
    outputs: [
      { name: '', internalType: 'bool', type: 'bool' },
      { name: '', internalType: 'bytes', type: 'bytes' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_name', internalType: 'string', type: 'string' },
      {
        name: '_config',
        internalType: 'struct Drippie.DripConfig',
        type: 'tuple',
        components: [
          { name: 'reentrant', internalType: 'bool', type: 'bool' },
          { name: 'interval', internalType: 'uint256', type: 'uint256' },
          {
            name: 'dripcheck',
            internalType: 'contract IDripCheck',
            type: 'address',
          },
          { name: 'checkparams', internalType: 'bytes', type: 'bytes' },
          {
            name: 'actions',
            internalType: 'struct Drippie.DripAction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'target',
                internalType: 'address payable',
                type: 'address',
              },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
      },
    ],
    name: 'create',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_name', internalType: 'string', type: 'string' }],
    name: 'drip',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'string', type: 'string' }],
    name: 'drips',
    outputs: [
      {
        name: 'status',
        internalType: 'enum Drippie.DripStatus',
        type: 'uint8',
      },
      {
        name: 'config',
        internalType: 'struct Drippie.DripConfig',
        type: 'tuple',
        components: [
          { name: 'reentrant', internalType: 'bool', type: 'bool' },
          { name: 'interval', internalType: 'uint256', type: 'uint256' },
          {
            name: 'dripcheck',
            internalType: 'contract IDripCheck',
            type: 'address',
          },
          { name: 'checkparams', internalType: 'bytes', type: 'bytes' },
          {
            name: 'actions',
            internalType: 'struct Drippie.DripAction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'target',
                internalType: 'address payable',
                type: 'address',
              },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
      },
      { name: 'last', internalType: 'uint256', type: 'uint256' },
      { name: 'count', internalType: 'uint256', type: 'uint256' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_name', internalType: 'string', type: 'string' }],
    name: 'executable',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'setOwner',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_name', internalType: 'string', type: 'string' },
      {
        name: '_status',
        internalType: 'enum Drippie.DripStatus',
        type: 'uint8',
      },
    ],
    name: 'status',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC20', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC20', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
    ],
    name: 'withdrawERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC721', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_id', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawERC721',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address payable', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawETH',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_to', internalType: 'address payable', type: 'address' }],
    name: 'withdrawETH',
    outputs: [],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export const drippieAddress = {
  1: '0x44b3A2a040057eBafC601A78647e805fd58B1f50',
} as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export const drippieConfig = {
  address: drippieAddress,
  abi: drippieABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Drippie_goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export const drippieGoerliABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_owner', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'nameref',
        internalType: 'string',
        type: 'string',
        indexed: true,
      },
      { name: 'name', internalType: 'string', type: 'string', indexed: false },
      {
        name: 'config',
        internalType: 'struct Drippie.DripConfig',
        type: 'tuple',
        components: [
          { name: 'reentrant', internalType: 'bool', type: 'bool' },
          { name: 'interval', internalType: 'uint256', type: 'uint256' },
          {
            name: 'dripcheck',
            internalType: 'contract IDripCheck',
            type: 'address',
          },
          { name: 'checkparams', internalType: 'bytes', type: 'bytes' },
          {
            name: 'actions',
            internalType: 'struct Drippie.DripAction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'target',
                internalType: 'address payable',
                type: 'address',
              },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
        indexed: false,
      },
    ],
    name: 'DripCreated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'nameref',
        internalType: 'string',
        type: 'string',
        indexed: true,
      },
      { name: 'name', internalType: 'string', type: 'string', indexed: false },
      {
        name: 'executor',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
      {
        name: 'timestamp',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'DripExecuted',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'nameref',
        internalType: 'string',
        type: 'string',
        indexed: true,
      },
      { name: 'name', internalType: 'string', type: 'string', indexed: false },
      {
        name: 'status',
        internalType: 'enum Drippie.DripStatus',
        type: 'uint8',
        indexed: false,
      },
    ],
    name: 'DripStatusUpdated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'user', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnerUpdated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'ReceivedETH',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'WithdrewERC20',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'id', internalType: 'uint256', type: 'uint256', indexed: false },
    ],
    name: 'WithdrewERC721',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'WithdrewETH',
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
      { name: '_value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'CALL',
    outputs: [
      { name: '', internalType: 'bool', type: 'bool' },
      { name: '', internalType: 'bytes', type: 'bytes' },
    ],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'DELEGATECALL',
    outputs: [
      { name: '', internalType: 'bool', type: 'bool' },
      { name: '', internalType: 'bytes', type: 'bytes' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_name', internalType: 'string', type: 'string' },
      {
        name: '_config',
        internalType: 'struct Drippie.DripConfig',
        type: 'tuple',
        components: [
          { name: 'reentrant', internalType: 'bool', type: 'bool' },
          { name: 'interval', internalType: 'uint256', type: 'uint256' },
          {
            name: 'dripcheck',
            internalType: 'contract IDripCheck',
            type: 'address',
          },
          { name: 'checkparams', internalType: 'bytes', type: 'bytes' },
          {
            name: 'actions',
            internalType: 'struct Drippie.DripAction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'target',
                internalType: 'address payable',
                type: 'address',
              },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
      },
    ],
    name: 'create',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_name', internalType: 'string', type: 'string' }],
    name: 'drip',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'string', type: 'string' }],
    name: 'drips',
    outputs: [
      {
        name: 'status',
        internalType: 'enum Drippie.DripStatus',
        type: 'uint8',
      },
      {
        name: 'config',
        internalType: 'struct Drippie.DripConfig',
        type: 'tuple',
        components: [
          { name: 'reentrant', internalType: 'bool', type: 'bool' },
          { name: 'interval', internalType: 'uint256', type: 'uint256' },
          {
            name: 'dripcheck',
            internalType: 'contract IDripCheck',
            type: 'address',
          },
          { name: 'checkparams', internalType: 'bytes', type: 'bytes' },
          {
            name: 'actions',
            internalType: 'struct Drippie.DripAction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'target',
                internalType: 'address payable',
                type: 'address',
              },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
      },
      { name: 'last', internalType: 'uint256', type: 'uint256' },
      { name: 'count', internalType: 'uint256', type: 'uint256' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_name', internalType: 'string', type: 'string' }],
    name: 'executable',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'setOwner',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_name', internalType: 'string', type: 'string' },
      {
        name: '_status',
        internalType: 'enum Drippie.DripStatus',
        type: 'uint8',
      },
    ],
    name: 'status',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC20', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC20', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
    ],
    name: 'withdrawERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC721', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_id', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawERC721',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address payable', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawETH',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_to', internalType: 'address payable', type: 'address' }],
    name: 'withdrawETH',
    outputs: [],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export const drippieGoerliAddress = {
  5: '0x44b3A2a040057eBafC601A78647e805fd58B1f50',
} as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export const drippieGoerliConfig = {
  address: drippieGoerliAddress,
  abi: drippieGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Drippie_optimism-goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x8D8d533C16D23847EB04EEB0925be8900Dd3af86)
 */
export const drippieOptimismGoerliABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_owner', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'nameref',
        internalType: 'string',
        type: 'string',
        indexed: true,
      },
      { name: 'name', internalType: 'string', type: 'string', indexed: false },
      {
        name: 'config',
        internalType: 'struct Drippie.DripConfig',
        type: 'tuple',
        components: [
          { name: 'interval', internalType: 'uint256', type: 'uint256' },
          {
            name: 'dripcheck',
            internalType: 'contract IDripCheck',
            type: 'address',
          },
          { name: 'checkparams', internalType: 'bytes', type: 'bytes' },
          {
            name: 'actions',
            internalType: 'struct Drippie.DripAction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'target',
                internalType: 'address payable',
                type: 'address',
              },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
        indexed: false,
      },
    ],
    name: 'DripCreated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'nameref',
        internalType: 'string',
        type: 'string',
        indexed: true,
      },
      { name: 'name', internalType: 'string', type: 'string', indexed: false },
      {
        name: 'executor',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
      {
        name: 'timestamp',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'DripExecuted',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'nameref',
        internalType: 'string',
        type: 'string',
        indexed: true,
      },
      { name: 'name', internalType: 'string', type: 'string', indexed: false },
      {
        name: 'status',
        internalType: 'enum Drippie.DripStatus',
        type: 'uint8',
        indexed: false,
      },
    ],
    name: 'DripStatusUpdated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'user', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnerUpdated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'ReceivedETH',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'WithdrewERC20',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'id', internalType: 'uint256', type: 'uint256', indexed: false },
    ],
    name: 'WithdrewERC721',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'WithdrewETH',
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
      { name: '_gas', internalType: 'uint256', type: 'uint256' },
      { name: '_value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'CALL',
    outputs: [
      { name: '', internalType: 'bool', type: 'bool' },
      { name: '', internalType: 'bytes', type: 'bytes' },
    ],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
      { name: '_gas', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'DELEGATECALL',
    outputs: [
      { name: '', internalType: 'bool', type: 'bool' },
      { name: '', internalType: 'bytes', type: 'bytes' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_name', internalType: 'string', type: 'string' },
      {
        name: '_config',
        internalType: 'struct Drippie.DripConfig',
        type: 'tuple',
        components: [
          { name: 'interval', internalType: 'uint256', type: 'uint256' },
          {
            name: 'dripcheck',
            internalType: 'contract IDripCheck',
            type: 'address',
          },
          { name: 'checkparams', internalType: 'bytes', type: 'bytes' },
          {
            name: 'actions',
            internalType: 'struct Drippie.DripAction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'target',
                internalType: 'address payable',
                type: 'address',
              },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
      },
    ],
    name: 'create',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_name', internalType: 'string', type: 'string' }],
    name: 'drip',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'string', type: 'string' }],
    name: 'drips',
    outputs: [
      {
        name: 'status',
        internalType: 'enum Drippie.DripStatus',
        type: 'uint8',
      },
      {
        name: 'config',
        internalType: 'struct Drippie.DripConfig',
        type: 'tuple',
        components: [
          { name: 'interval', internalType: 'uint256', type: 'uint256' },
          {
            name: 'dripcheck',
            internalType: 'contract IDripCheck',
            type: 'address',
          },
          { name: 'checkparams', internalType: 'bytes', type: 'bytes' },
          {
            name: 'actions',
            internalType: 'struct Drippie.DripAction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'target',
                internalType: 'address payable',
                type: 'address',
              },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
      },
      { name: 'last', internalType: 'uint256', type: 'uint256' },
      { name: 'count', internalType: 'uint256', type: 'uint256' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_name', internalType: 'string', type: 'string' }],
    name: 'executable',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'setOwner',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_name', internalType: 'string', type: 'string' },
      {
        name: '_status',
        internalType: 'enum Drippie.DripStatus',
        type: 'uint8',
      },
    ],
    name: 'status',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC20', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC20', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
    ],
    name: 'withdrawERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC721', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_id', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawERC721',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address payable', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawETH',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_to', internalType: 'address payable', type: 'address' }],
    name: 'withdrawETH',
    outputs: [],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x8D8d533C16D23847EB04EEB0925be8900Dd3af86)
 */
export const drippieOptimismGoerliAddress = {
  420: '0x8D8d533C16D23847EB04EEB0925be8900Dd3af86',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x8D8d533C16D23847EB04EEB0925be8900Dd3af86)
 */
export const drippieOptimismGoerliConfig = {
  address: drippieOptimismGoerliAddress,
  abi: drippieOptimismGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// EAS
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4E0275Ea5a89e7a3c1B58411379D1a0eDdc5b088)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5A633F1cc84B03F7588486CF2F386c102061E6e1)
 */
export const easABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  { type: 'error', inputs: [], name: 'AccessDenied' },
  { type: 'error', inputs: [], name: 'AlreadyRevoked' },
  { type: 'error', inputs: [], name: 'AlreadyRevokedOffchain' },
  { type: 'error', inputs: [], name: 'AlreadyTimestamped' },
  { type: 'error', inputs: [], name: 'InsufficientValue' },
  { type: 'error', inputs: [], name: 'InvalidAttestation' },
  { type: 'error', inputs: [], name: 'InvalidAttestations' },
  { type: 'error', inputs: [], name: 'InvalidExpirationTime' },
  { type: 'error', inputs: [], name: 'InvalidLength' },
  { type: 'error', inputs: [], name: 'InvalidOffset' },
  { type: 'error', inputs: [], name: 'InvalidRegistry' },
  { type: 'error', inputs: [], name: 'InvalidRevocation' },
  { type: 'error', inputs: [], name: 'InvalidRevocations' },
  { type: 'error', inputs: [], name: 'InvalidSchema' },
  { type: 'error', inputs: [], name: 'InvalidSignature' },
  { type: 'error', inputs: [], name: 'InvalidVerifier' },
  { type: 'error', inputs: [], name: 'Irrevocable' },
  { type: 'error', inputs: [], name: 'NotFound' },
  { type: 'error', inputs: [], name: 'NotPayable' },
  { type: 'error', inputs: [], name: 'WrongSchema' },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'attester',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'uid', internalType: 'bytes32', type: 'bytes32', indexed: false },
      {
        name: 'schema',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
    ],
    name: 'Attested',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'attester',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'uid', internalType: 'bytes32', type: 'bytes32', indexed: false },
      {
        name: 'schema',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
    ],
    name: 'Revoked',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'revoker',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'data', internalType: 'bytes32', type: 'bytes32', indexed: true },
      {
        name: 'timestamp',
        internalType: 'uint64',
        type: 'uint64',
        indexed: true,
      },
    ],
    name: 'RevokedOffchain',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'data', internalType: 'bytes32', type: 'bytes32', indexed: true },
      {
        name: 'timestamp',
        internalType: 'uint64',
        type: 'uint64',
        indexed: true,
      },
    ],
    name: 'Timestamped',
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      {
        name: 'request',
        internalType: 'struct AttestationRequest',
        type: 'tuple',
        components: [
          { name: 'schema', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'data',
            internalType: 'struct AttestationRequestData',
            type: 'tuple',
            components: [
              { name: 'recipient', internalType: 'address', type: 'address' },
              {
                name: 'expirationTime',
                internalType: 'uint64',
                type: 'uint64',
              },
              { name: 'revocable', internalType: 'bool', type: 'bool' },
              { name: 'refUID', internalType: 'bytes32', type: 'bytes32' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
      },
    ],
    name: 'attest',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      {
        name: 'delegatedRequest',
        internalType: 'struct DelegatedAttestationRequest',
        type: 'tuple',
        components: [
          { name: 'schema', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'data',
            internalType: 'struct AttestationRequestData',
            type: 'tuple',
            components: [
              { name: 'recipient', internalType: 'address', type: 'address' },
              {
                name: 'expirationTime',
                internalType: 'uint64',
                type: 'uint64',
              },
              { name: 'revocable', internalType: 'bool', type: 'bool' },
              { name: 'refUID', internalType: 'bytes32', type: 'bytes32' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
          {
            name: 'signature',
            internalType: 'struct EIP712Signature',
            type: 'tuple',
            components: [
              { name: 'v', internalType: 'uint8', type: 'uint8' },
              { name: 'r', internalType: 'bytes32', type: 'bytes32' },
              { name: 's', internalType: 'bytes32', type: 'bytes32' },
            ],
          },
          { name: 'attester', internalType: 'address', type: 'address' },
        ],
      },
    ],
    name: 'attestByDelegation',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [],
    name: 'getAttestTypeHash',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: 'uid', internalType: 'bytes32', type: 'bytes32' }],
    name: 'getAttestation',
    outputs: [
      {
        name: '',
        internalType: 'struct Attestation',
        type: 'tuple',
        components: [
          { name: 'uid', internalType: 'bytes32', type: 'bytes32' },
          { name: 'schema', internalType: 'bytes32', type: 'bytes32' },
          { name: 'time', internalType: 'uint64', type: 'uint64' },
          { name: 'expirationTime', internalType: 'uint64', type: 'uint64' },
          { name: 'revocationTime', internalType: 'uint64', type: 'uint64' },
          { name: 'refUID', internalType: 'bytes32', type: 'bytes32' },
          { name: 'recipient', internalType: 'address', type: 'address' },
          { name: 'attester', internalType: 'address', type: 'address' },
          { name: 'revocable', internalType: 'bool', type: 'bool' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'getDomainSeparator',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'getName',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: 'account', internalType: 'address', type: 'address' }],
    name: 'getNonce',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: 'revoker', internalType: 'address', type: 'address' },
      { name: 'data', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'getRevokeOffchain',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [],
    name: 'getRevokeTypeHash',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [],
    name: 'getSchemaRegistry',
    outputs: [
      { name: '', internalType: 'contract ISchemaRegistry', type: 'address' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: 'data', internalType: 'bytes32', type: 'bytes32' }],
    name: 'getTimestamp',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: 'uid', internalType: 'bytes32', type: 'bytes32' }],
    name: 'isAttestationValid',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      {
        name: 'multiRequests',
        internalType: 'struct MultiAttestationRequest[]',
        type: 'tuple[]',
        components: [
          { name: 'schema', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'data',
            internalType: 'struct AttestationRequestData[]',
            type: 'tuple[]',
            components: [
              { name: 'recipient', internalType: 'address', type: 'address' },
              {
                name: 'expirationTime',
                internalType: 'uint64',
                type: 'uint64',
              },
              { name: 'revocable', internalType: 'bool', type: 'bool' },
              { name: 'refUID', internalType: 'bytes32', type: 'bytes32' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
      },
    ],
    name: 'multiAttest',
    outputs: [{ name: '', internalType: 'bytes32[]', type: 'bytes32[]' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      {
        name: 'multiDelegatedRequests',
        internalType: 'struct MultiDelegatedAttestationRequest[]',
        type: 'tuple[]',
        components: [
          { name: 'schema', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'data',
            internalType: 'struct AttestationRequestData[]',
            type: 'tuple[]',
            components: [
              { name: 'recipient', internalType: 'address', type: 'address' },
              {
                name: 'expirationTime',
                internalType: 'uint64',
                type: 'uint64',
              },
              { name: 'revocable', internalType: 'bool', type: 'bool' },
              { name: 'refUID', internalType: 'bytes32', type: 'bytes32' },
              { name: 'data', internalType: 'bytes', type: 'bytes' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
          {
            name: 'signatures',
            internalType: 'struct EIP712Signature[]',
            type: 'tuple[]',
            components: [
              { name: 'v', internalType: 'uint8', type: 'uint8' },
              { name: 'r', internalType: 'bytes32', type: 'bytes32' },
              { name: 's', internalType: 'bytes32', type: 'bytes32' },
            ],
          },
          { name: 'attester', internalType: 'address', type: 'address' },
        ],
      },
    ],
    name: 'multiAttestByDelegation',
    outputs: [{ name: '', internalType: 'bytes32[]', type: 'bytes32[]' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      {
        name: 'multiRequests',
        internalType: 'struct MultiRevocationRequest[]',
        type: 'tuple[]',
        components: [
          { name: 'schema', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'data',
            internalType: 'struct RevocationRequestData[]',
            type: 'tuple[]',
            components: [
              { name: 'uid', internalType: 'bytes32', type: 'bytes32' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
      },
    ],
    name: 'multiRevoke',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      {
        name: 'multiDelegatedRequests',
        internalType: 'struct MultiDelegatedRevocationRequest[]',
        type: 'tuple[]',
        components: [
          { name: 'schema', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'data',
            internalType: 'struct RevocationRequestData[]',
            type: 'tuple[]',
            components: [
              { name: 'uid', internalType: 'bytes32', type: 'bytes32' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
          {
            name: 'signatures',
            internalType: 'struct EIP712Signature[]',
            type: 'tuple[]',
            components: [
              { name: 'v', internalType: 'uint8', type: 'uint8' },
              { name: 'r', internalType: 'bytes32', type: 'bytes32' },
              { name: 's', internalType: 'bytes32', type: 'bytes32' },
            ],
          },
          { name: 'revoker', internalType: 'address', type: 'address' },
        ],
      },
    ],
    name: 'multiRevokeByDelegation',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'data', internalType: 'bytes32[]', type: 'bytes32[]' }],
    name: 'multiRevokeOffchain',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'data', internalType: 'bytes32[]', type: 'bytes32[]' }],
    name: 'multiTimestamp',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      {
        name: 'request',
        internalType: 'struct RevocationRequest',
        type: 'tuple',
        components: [
          { name: 'schema', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'data',
            internalType: 'struct RevocationRequestData',
            type: 'tuple',
            components: [
              { name: 'uid', internalType: 'bytes32', type: 'bytes32' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
        ],
      },
    ],
    name: 'revoke',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      {
        name: 'delegatedRequest',
        internalType: 'struct DelegatedRevocationRequest',
        type: 'tuple',
        components: [
          { name: 'schema', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'data',
            internalType: 'struct RevocationRequestData',
            type: 'tuple',
            components: [
              { name: 'uid', internalType: 'bytes32', type: 'bytes32' },
              { name: 'value', internalType: 'uint256', type: 'uint256' },
            ],
          },
          {
            name: 'signature',
            internalType: 'struct EIP712Signature',
            type: 'tuple',
            components: [
              { name: 'v', internalType: 'uint8', type: 'uint8' },
              { name: 'r', internalType: 'bytes32', type: 'bytes32' },
              { name: 's', internalType: 'bytes32', type: 'bytes32' },
            ],
          },
          { name: 'revoker', internalType: 'address', type: 'address' },
        ],
      },
    ],
    name: 'revokeByDelegation',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'data', internalType: 'bytes32', type: 'bytes32' }],
    name: 'revokeOffchain',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'data', internalType: 'bytes32', type: 'bytes32' }],
    name: 'timestamp',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4E0275Ea5a89e7a3c1B58411379D1a0eDdc5b088)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5A633F1cc84B03F7588486CF2F386c102061E6e1)
 */
export const easAddress = {
  10: '0x4E0275Ea5a89e7a3c1B58411379D1a0eDdc5b088',
  420: '0x5A633F1cc84B03F7588486CF2F386c102061E6e1',
} as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4E0275Ea5a89e7a3c1B58411379D1a0eDdc5b088)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5A633F1cc84B03F7588486CF2F386c102061E6e1)
 */
export const easConfig = { address: easAddress, abi: easABI } as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// GasPriceOracle
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000000F)
 */
export const gasPriceOracleABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'baseFee',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [],
    name: 'decimals',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'gasPrice',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_data', internalType: 'bytes', type: 'bytes' }],
    name: 'getL1Fee',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_data', internalType: 'bytes', type: 'bytes' }],
    name: 'getL1GasUsed',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l1BaseFee',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'overhead',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'scalar',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000000F)
 */
export const gasPriceOracleAddress = {
  420: '0x420000000000000000000000000000000000000F',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000000F)
 */
export const gasPriceOracleConfig = {
  address: gasPriceOracleAddress,
  abi: gasPriceOracleABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L1Block
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000015)
 */
export const l1BlockABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'DEPOSITOR_ACCOUNT',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'basefee',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'batcherHash',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'hash',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l1FeeOverhead',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l1FeeScalar',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'number',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'sequenceNumber',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_number', internalType: 'uint64', type: 'uint64' },
      { name: '_timestamp', internalType: 'uint64', type: 'uint64' },
      { name: '_basefee', internalType: 'uint256', type: 'uint256' },
      { name: '_hash', internalType: 'bytes32', type: 'bytes32' },
      { name: '_sequenceNumber', internalType: 'uint64', type: 'uint64' },
      { name: '_batcherHash', internalType: 'bytes32', type: 'bytes32' },
      { name: '_l1FeeOverhead', internalType: 'uint256', type: 'uint256' },
      { name: '_l1FeeScalar', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'setL1BlockValues',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'timestamp',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000015)
 */
export const l1BlockAddress = {
  420: '0x4200000000000000000000000000000000000015',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000015)
 */
export const l1BlockConfig = {
  address: l1BlockAddress,
  abi: l1BlockABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L1CrossDomainMessenger
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1)
 */
export const l1CrossDomainMessengerABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      {
        name: '_portal',
        internalType: 'contract OptimismPortal',
        type: 'address',
      },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'msgHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
    ],
    name: 'FailedRelayedMessage',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'msgHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
    ],
    name: 'RelayedMessage',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'target',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'sender',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
      { name: 'message', internalType: 'bytes', type: 'bytes', indexed: false },
      {
        name: 'messageNonce',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'gasLimit',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'SentMessage',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'sender',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'SentMessageExtension1',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MESSAGE_VERSION',
    outputs: [{ name: '', internalType: 'uint16', type: 'uint16' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_GAS_CALLDATA_OVERHEAD',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'OTHER_MESSENGER',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'PORTAL',
    outputs: [
      { name: '', internalType: 'contract OptimismPortal', type: 'address' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_CALL_OVERHEAD',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_CONSTANT_OVERHEAD',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_GAS_CHECK_BUFFER',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_RESERVED_GAS',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [
      { name: '_message', internalType: 'bytes', type: 'bytes' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
    ],
    name: 'baseGas',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'failedMessages',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'messageNonce',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_nonce', internalType: 'uint256', type: 'uint256' },
      { name: '_sender', internalType: 'address', type: 'address' },
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_value', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint256', type: 'uint256' },
      { name: '_message', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'relayMessage',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_message', internalType: 'bytes', type: 'bytes' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
    ],
    name: 'sendMessage',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'successfulMessages',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'xDomainMessageSender',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
] as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1)
 */
export const l1CrossDomainMessengerAddress = {
  1: '0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1',
} as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1)
 */
export const l1CrossDomainMessengerConfig = {
  address: l1CrossDomainMessengerAddress,
  abi: l1CrossDomainMessengerABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L1CrossDomainMessenger_goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5086d1eEF304eb5284A0f6720f79403b4e9bE294)
 */
export const l1CrossDomainMessengerGoerliABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'msgHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
    ],
    name: 'FailedRelayedMessage',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'msgHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
    ],
    name: 'RelayedMessage',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'target',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'sender',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
      { name: 'message', internalType: 'bytes', type: 'bytes', indexed: false },
      {
        name: 'messageNonce',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'gasLimit',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'SentMessage',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'sender',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'SentMessageExtension1',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MESSAGE_VERSION',
    outputs: [{ name: '', internalType: 'uint16', type: 'uint16' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_GAS_CALLDATA_OVERHEAD',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'OTHER_MESSENGER',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_CALL_OVERHEAD',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_CONSTANT_OVERHEAD',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_GAS_CHECK_BUFFER',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_RESERVED_GAS',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [
      { name: '_message', internalType: 'bytes', type: 'bytes' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
    ],
    name: 'baseGas',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'failedMessages',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_portal',
        internalType: 'contract OptimismPortal',
        type: 'address',
      },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'messageNonce',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'portal',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_nonce', internalType: 'uint256', type: 'uint256' },
      { name: '_sender', internalType: 'address', type: 'address' },
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_value', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint256', type: 'uint256' },
      { name: '_message', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'relayMessage',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_message', internalType: 'bytes', type: 'bytes' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
    ],
    name: 'sendMessage',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'successfulMessages',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'xDomainMessageSender',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
] as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5086d1eEF304eb5284A0f6720f79403b4e9bE294)
 */
export const l1CrossDomainMessengerGoerliAddress = {
  5: '0x5086d1eEF304eb5284A0f6720f79403b4e9bE294',
} as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5086d1eEF304eb5284A0f6720f79403b4e9bE294)
 */
export const l1CrossDomainMessengerGoerliConfig = {
  address: l1CrossDomainMessengerGoerliAddress,
  abi: l1CrossDomainMessengerGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L1ERC721Bridge
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5a7749f83b81B301cAb5f48EB8516B986DAef23D)
 */
export const l1Erc721BridgeABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_messenger', internalType: 'address', type: 'address' },
      { name: '_otherBridge', internalType: 'address', type: 'address' },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'tokenId',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC721BridgeFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'tokenId',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC721BridgeInitiated',
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_tokenId', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC721',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_tokenId', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC721To',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'deposits',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_tokenId', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeBridgeERC721',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'messenger',
    outputs: [
      {
        name: '',
        internalType: 'contract CrossDomainMessenger',
        type: 'address',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'otherBridge',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5a7749f83b81B301cAb5f48EB8516B986DAef23D)
 */
export const l1Erc721BridgeAddress = {
  1: '0x5a7749f83b81B301cAb5f48EB8516B986DAef23D',
} as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5a7749f83b81B301cAb5f48EB8516B986DAef23D)
 */
export const l1Erc721BridgeConfig = {
  address: l1Erc721BridgeAddress,
  abi: l1Erc721BridgeABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L1ERC721Bridge_goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x8DD330DdE8D9898d43b4dc840Da27A07dF91b3c9)
 */
export const l1Erc721BridgeGoerliABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'tokenId',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC721BridgeFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'tokenId',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC721BridgeInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_tokenId', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC721',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_tokenId', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC721To',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'deposits',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_tokenId', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeBridgeERC721',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_messenger',
        internalType: 'contract CrossDomainMessenger',
        type: 'address',
      },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'messenger',
    outputs: [
      {
        name: '',
        internalType: 'contract CrossDomainMessenger',
        type: 'address',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'otherBridge',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x8DD330DdE8D9898d43b4dc840Da27A07dF91b3c9)
 */
export const l1Erc721BridgeGoerliAddress = {
  5: '0x8DD330DdE8D9898d43b4dc840Da27A07dF91b3c9',
} as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x8DD330DdE8D9898d43b4dc840Da27A07dF91b3c9)
 */
export const l1Erc721BridgeGoerliConfig = {
  address: l1Erc721BridgeGoerliAddress,
  abi: l1Erc721BridgeGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L1FeeVault
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000001a)
 */
export const l1FeeVaultABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_recipient', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'from',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'Withdrawal',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_WITHDRAWAL_AMOUNT',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RECIPIENT',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'totalProcessed',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'withdraw',
    outputs: [],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000001a)
 */
export const l1FeeVaultAddress = {
  420: '0x420000000000000000000000000000000000001A',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000001a)
 */
export const l1FeeVaultConfig = {
  address: l1FeeVaultAddress,
  abi: l1FeeVaultABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L1StandardBridge
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1)
 */
export const l1StandardBridgeABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_messenger', internalType: 'address payable', type: 'address' },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC20BridgeFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC20BridgeInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'l1Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'l2Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC20DepositInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'l1Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'l2Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC20WithdrawalFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ETHBridgeFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ETHBridgeInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ETHDepositInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ETHWithdrawalFinalized',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'OTHER_BRIDGE',
    outputs: [
      { name: '', internalType: 'contract StandardBridge', type: 'address' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC20To',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeETH',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeETHTo',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_l1Token', internalType: 'address', type: 'address' },
      { name: '_l2Token', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'depositERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_l1Token', internalType: 'address', type: 'address' },
      { name: '_l2Token', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'depositERC20To',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'depositETH',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'depositETHTo',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'address', type: 'address' },
    ],
    name: 'deposits',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeBridgeERC20',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeBridgeETH',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_l1Token', internalType: 'address', type: 'address' },
      { name: '_l2Token', internalType: 'address', type: 'address' },
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeERC20Withdrawal',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeETHWithdrawal',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l2TokenBridge',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'messenger',
    outputs: [
      {
        name: '',
        internalType: 'contract CrossDomainMessenger',
        type: 'address',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1)
 */
export const l1StandardBridgeAddress = {
  1: '0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1',
} as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1)
 */
export const l1StandardBridgeConfig = {
  address: l1StandardBridgeAddress,
  abi: l1StandardBridgeABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L1StandardBridge_goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x636Af16bf2f682dD3109e60102b8E1A089FedAa8)
 */
export const l1StandardBridgeGoerliABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC20BridgeFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC20BridgeInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'l1Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'l2Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC20DepositInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'l1Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'l2Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC20WithdrawalFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ETHBridgeFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ETHBridgeInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ETHDepositInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ETHWithdrawalFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC20To',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeETH',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeETHTo',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_l1Token', internalType: 'address', type: 'address' },
      { name: '_l2Token', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'depositERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_l1Token', internalType: 'address', type: 'address' },
      { name: '_l2Token', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'depositERC20To',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'depositETH',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'depositETHTo',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'address', type: 'address' },
    ],
    name: 'deposits',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeBridgeERC20',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeBridgeETH',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_l1Token', internalType: 'address', type: 'address' },
      { name: '_l2Token', internalType: 'address', type: 'address' },
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeERC20Withdrawal',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeETHWithdrawal',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_messenger',
        internalType: 'contract CrossDomainMessenger',
        type: 'address',
      },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l2TokenBridge',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'messenger',
    outputs: [
      {
        name: '',
        internalType: 'contract CrossDomainMessenger',
        type: 'address',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'otherBridge',
    outputs: [
      { name: '', internalType: 'contract StandardBridge', type: 'address' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x636Af16bf2f682dD3109e60102b8E1A089FedAa8)
 */
export const l1StandardBridgeGoerliAddress = {
  5: '0x636Af16bf2f682dD3109e60102b8E1A089FedAa8',
} as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x636Af16bf2f682dD3109e60102b8E1A089FedAa8)
 */
export const l1StandardBridgeGoerliConfig = {
  address: l1StandardBridgeGoerliAddress,
  abi: l1StandardBridgeGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L2CrossDomainMessenger
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000007)
 */
export const l2CrossDomainMessengerABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      {
        name: '_l1CrossDomainMessenger',
        internalType: 'address',
        type: 'address',
      },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'msgHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
    ],
    name: 'FailedRelayedMessage',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'msgHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
    ],
    name: 'RelayedMessage',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'target',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'sender',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
      { name: 'message', internalType: 'bytes', type: 'bytes', indexed: false },
      {
        name: 'messageNonce',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'gasLimit',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'SentMessage',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'sender',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'SentMessageExtension1',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MESSAGE_VERSION',
    outputs: [{ name: '', internalType: 'uint16', type: 'uint16' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_GAS_CALLDATA_OVERHEAD',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'OTHER_MESSENGER',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_CALL_OVERHEAD',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_CONSTANT_OVERHEAD',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_GAS_CHECK_BUFFER',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RELAY_RESERVED_GAS',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [
      { name: '_message', internalType: 'bytes', type: 'bytes' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
    ],
    name: 'baseGas',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'failedMessages',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l1CrossDomainMessenger',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'messageNonce',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_nonce', internalType: 'uint256', type: 'uint256' },
      { name: '_sender', internalType: 'address', type: 'address' },
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_value', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint256', type: 'uint256' },
      { name: '_message', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'relayMessage',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_message', internalType: 'bytes', type: 'bytes' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
    ],
    name: 'sendMessage',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'successfulMessages',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'xDomainMessageSender',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000007)
 */
export const l2CrossDomainMessengerAddress = {
  420: '0x4200000000000000000000000000000000000007',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000007)
 */
export const l2CrossDomainMessengerConfig = {
  address: l2CrossDomainMessengerAddress,
  abi: l2CrossDomainMessengerABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L2ERC721Bridge
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000014)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000014)
 */
export const l2Erc721BridgeABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_messenger', internalType: 'address', type: 'address' },
      { name: '_otherBridge', internalType: 'address', type: 'address' },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'tokenId',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC721BridgeFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'tokenId',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC721BridgeInitiated',
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_tokenId', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC721',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_tokenId', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC721To',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_tokenId', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeBridgeERC721',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'messenger',
    outputs: [
      {
        name: '',
        internalType: 'contract CrossDomainMessenger',
        type: 'address',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'otherBridge',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000014)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000014)
 */
export const l2Erc721BridgeAddress = {
  10: '0x4200000000000000000000000000000000000014',
  420: '0x4200000000000000000000000000000000000014',
} as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000014)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000014)
 */
export const l2Erc721BridgeConfig = {
  address: l2Erc721BridgeAddress,
  abi: l2Erc721BridgeABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L2OutputOracle
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdfe97868233d1aa22e815a266982f2cf17685a27)
 */
export const l2OutputOracleABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_submissionInterval', internalType: 'uint256', type: 'uint256' },
      { name: '_l2BlockTime', internalType: 'uint256', type: 'uint256' },
      {
        name: '_startingBlockNumber',
        internalType: 'uint256',
        type: 'uint256',
      },
      { name: '_startingTimestamp', internalType: 'uint256', type: 'uint256' },
      { name: '_proposer', internalType: 'address', type: 'address' },
      { name: '_challenger', internalType: 'address', type: 'address' },
      {
        name: '_finalizationPeriodSeconds',
        internalType: 'uint256',
        type: 'uint256',
      },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'outputRoot',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
      {
        name: 'l2OutputIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'l2BlockNumber',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'l1Timestamp',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'OutputProposed',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'prevNextOutputIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'newNextOutputIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
    ],
    name: 'OutputsDeleted',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'CHALLENGER',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'FINALIZATION_PERIOD_SECONDS',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'L2_BLOCK_TIME',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'PROPOSER',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'SUBMISSION_INTERVAL',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_l2BlockNumber', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'computeL2Timestamp',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_l2OutputIndex', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'deleteL2Outputs',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_l2OutputIndex', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'getL2Output',
    outputs: [
      {
        name: '',
        internalType: 'struct Types.OutputProposal',
        type: 'tuple',
        components: [
          { name: 'outputRoot', internalType: 'bytes32', type: 'bytes32' },
          { name: 'timestamp', internalType: 'uint128', type: 'uint128' },
          { name: 'l2BlockNumber', internalType: 'uint128', type: 'uint128' },
        ],
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_l2BlockNumber', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'getL2OutputAfter',
    outputs: [
      {
        name: '',
        internalType: 'struct Types.OutputProposal',
        type: 'tuple',
        components: [
          { name: 'outputRoot', internalType: 'bytes32', type: 'bytes32' },
          { name: 'timestamp', internalType: 'uint128', type: 'uint128' },
          { name: 'l2BlockNumber', internalType: 'uint128', type: 'uint128' },
        ],
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_l2BlockNumber', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'getL2OutputIndexAfter',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_startingBlockNumber',
        internalType: 'uint256',
        type: 'uint256',
      },
      { name: '_startingTimestamp', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'latestBlockNumber',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'latestOutputIndex',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'nextBlockNumber',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'nextOutputIndex',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_outputRoot', internalType: 'bytes32', type: 'bytes32' },
      { name: '_l2BlockNumber', internalType: 'uint256', type: 'uint256' },
      { name: '_l1BlockHash', internalType: 'bytes32', type: 'bytes32' },
      { name: '_l1BlockNumber', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'proposeL2Output',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'startingBlockNumber',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'startingTimestamp',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdfe97868233d1aa22e815a266982f2cf17685a27)
 */
export const l2OutputOracleAddress = {
  1: '0xdfe97868233d1aa22e815a266982f2cf17685a27',
} as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdfe97868233d1aa22e815a266982f2cf17685a27)
 */
export const l2OutputOracleConfig = {
  address: l2OutputOracleAddress,
  abi: l2OutputOracleABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L2OutputOracle_goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0)
 */
export const l2OutputOracleGoerliABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_submissionInterval', internalType: 'uint256', type: 'uint256' },
      { name: '_l2BlockTime', internalType: 'uint256', type: 'uint256' },
      {
        name: '_finalizationPeriodSeconds',
        internalType: 'uint256',
        type: 'uint256',
      },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'outputRoot',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
      {
        name: 'l2OutputIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'l2BlockNumber',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'l1Timestamp',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'OutputProposed',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'prevNextOutputIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'newNextOutputIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
    ],
    name: 'OutputsDeleted',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'challenger',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_l2BlockNumber', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'computeL2Timestamp',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_l2OutputIndex', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'deleteL2Outputs',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'finalizationPeriodSeconds',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_l2OutputIndex', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'getL2Output',
    outputs: [
      {
        name: '',
        internalType: 'struct Types.OutputProposal',
        type: 'tuple',
        components: [
          { name: 'outputRoot', internalType: 'bytes32', type: 'bytes32' },
          { name: 'timestamp', internalType: 'uint128', type: 'uint128' },
          { name: 'l2BlockNumber', internalType: 'uint128', type: 'uint128' },
        ],
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_l2BlockNumber', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'getL2OutputAfter',
    outputs: [
      {
        name: '',
        internalType: 'struct Types.OutputProposal',
        type: 'tuple',
        components: [
          { name: 'outputRoot', internalType: 'bytes32', type: 'bytes32' },
          { name: 'timestamp', internalType: 'uint128', type: 'uint128' },
          { name: 'l2BlockNumber', internalType: 'uint128', type: 'uint128' },
        ],
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_l2BlockNumber', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'getL2OutputIndexAfter',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_startingBlockNumber',
        internalType: 'uint256',
        type: 'uint256',
      },
      { name: '_startingTimestamp', internalType: 'uint256', type: 'uint256' },
      { name: '_proposer', internalType: 'address', type: 'address' },
      { name: '_challenger', internalType: 'address', type: 'address' },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l2BlockTime',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'latestBlockNumber',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'latestOutputIndex',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'nextBlockNumber',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'nextOutputIndex',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_outputRoot', internalType: 'bytes32', type: 'bytes32' },
      { name: '_l2BlockNumber', internalType: 'uint256', type: 'uint256' },
      { name: '_l1BlockHash', internalType: 'bytes32', type: 'bytes32' },
      { name: '_l1BlockNumber', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'proposeL2Output',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'proposer',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'startingBlockNumber',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'startingTimestamp',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'submissionInterval',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0)
 */
export const l2OutputOracleGoerliAddress = {
  5: '0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0',
} as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0)
 */
export const l2OutputOracleGoerliConfig = {
  address: l2OutputOracleGoerliAddress,
  abi: l2OutputOracleGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L2StandardBridge
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000010)
 */
export const l2StandardBridgeABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      {
        name: '_otherBridge',
        internalType: 'address payable',
        type: 'address',
      },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'l1Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'l2Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'DepositFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC20BridgeFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ERC20BridgeInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ETHBridgeFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'ETHBridgeInitiated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'l1Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'l2Token',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'extraData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'WithdrawalInitiated',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'OTHER_BRIDGE',
    outputs: [
      { name: '', internalType: 'contract StandardBridge', type: 'address' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeERC20To',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeETH',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'bridgeETHTo',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'address', type: 'address' },
    ],
    name: 'deposits',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_localToken', internalType: 'address', type: 'address' },
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeBridgeERC20',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeBridgeETH',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_l1Token', internalType: 'address', type: 'address' },
      { name: '_l2Token', internalType: 'address', type: 'address' },
      { name: '_from', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'finalizeDeposit',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l1TokenBridge',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'messenger',
    outputs: [
      {
        name: '',
        internalType: 'contract CrossDomainMessenger',
        type: 'address',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_l2Token', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'withdraw',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_l2Token', internalType: 'address', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
      { name: '_minGasLimit', internalType: 'uint32', type: 'uint32' },
      { name: '_extraData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'withdrawTo',
    outputs: [],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000010)
 */
export const l2StandardBridgeAddress = {
  420: '0x4200000000000000000000000000000000000010',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000010)
 */
export const l2StandardBridgeConfig = {
  address: l2StandardBridgeAddress,
  abi: l2StandardBridgeABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// L2ToL1MessagePasser
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000016)
 */
export const l2ToL1MessagePasserABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'nonce',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'sender',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'target',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'gasLimit',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      { name: 'data', internalType: 'bytes', type: 'bytes', indexed: false },
      {
        name: 'withdrawalHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: false,
      },
    ],
    name: 'MessagePassed',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
    ],
    name: 'WithdrawerBalanceBurnt',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MESSAGE_VERSION',
    outputs: [{ name: '', internalType: 'uint16', type: 'uint16' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'burn',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_gasLimit', internalType: 'uint256', type: 'uint256' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'initiateWithdrawal',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'messageNonce',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'sentMessages',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000016)
 */
export const l2ToL1MessagePasserAddress = {
  420: '0x4200000000000000000000000000000000000016',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000016)
 */
export const l2ToL1MessagePasserConfig = {
  address: l2ToL1MessagePasserAddress,
  abi: l2ToL1MessagePasserABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// MintManager
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x5C4e7Ba1E219E47948e6e3F55019A647bA501005)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x038a8825A3C3B0c08d52Cc76E5E361953Cf6Dc76)
 */
export const mintManagerABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_upgrader', internalType: 'address', type: 'address' },
      { name: '_governanceToken', internalType: 'address', type: 'address' },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'DENOMINATOR',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MINT_CAP',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MINT_PERIOD',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'governanceToken',
    outputs: [
      { name: '', internalType: 'contract GovernanceToken', type: 'address' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_account', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'mint',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'mintPermittedAfter',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_newMintManager', internalType: 'address', type: 'address' },
    ],
    name: 'upgrade',
    outputs: [],
  },
] as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x5C4e7Ba1E219E47948e6e3F55019A647bA501005)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x038a8825A3C3B0c08d52Cc76E5E361953Cf6Dc76)
 */
export const mintManagerAddress = {
  10: '0x5C4e7Ba1E219E47948e6e3F55019A647bA501005',
  420: '0x038a8825A3C3B0c08d52Cc76E5E361953Cf6Dc76',
} as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x5C4e7Ba1E219E47948e6e3F55019A647bA501005)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x038a8825A3C3B0c08d52Cc76E5E361953Cf6Dc76)
 */
export const mintManagerConfig = {
  address: mintManagerAddress,
  abi: mintManagerABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// OptimismMintableERC20Factory
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export const optimismMintableErc20FactoryABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_bridge', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'deployer',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'OptimismMintableERC20Created',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'StandardL2TokenCreated',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'BRIDGE',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
    ],
    name: 'createOptimismMintableERC20',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
    ],
    name: 'createStandardL2Token',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export const optimismMintableErc20FactoryAddress = {
  1: '0x4200000000000000000000000000000000000012',
} as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export const optimismMintableErc20FactoryConfig = {
  address: optimismMintableErc20FactoryAddress,
  abi: optimismMintableErc20FactoryABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// OptimismMintableERC20Factory_goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export const optimismMintableErc20FactoryGoerliABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'deployer',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'OptimismMintableERC20Created',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'StandardL2TokenCreated',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'bridge',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
    ],
    name: 'createOptimismMintableERC20',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
      { name: '_decimals', internalType: 'uint8', type: 'uint8' },
    ],
    name: 'createOptimismMintableERC20WithDecimals',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
    ],
    name: 'createStandardL2Token',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_bridge', internalType: 'address', type: 'address' }],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export const optimismMintableErc20FactoryGoerliAddress = {
  5: '0x4200000000000000000000000000000000000012',
} as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export const optimismMintableErc20FactoryGoerliConfig = {
  address: optimismMintableErc20FactoryGoerliAddress,
  abi: optimismMintableErc20FactoryGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// OptimismMintableERC20Factory_optimism-goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export const optimismMintableErc20FactoryOptimismGoerliABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_bridge', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'deployer',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'OptimismMintableERC20Created',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'StandardL2TokenCreated',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'BRIDGE',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
    ],
    name: 'createOptimismMintableERC20',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
    ],
    name: 'createStandardL2Token',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export const optimismMintableErc20FactoryOptimismGoerliAddress = {
  420: '0x4200000000000000000000000000000000000012',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export const optimismMintableErc20FactoryOptimismGoerliConfig = {
  address: optimismMintableErc20FactoryOptimismGoerliAddress,
  abi: optimismMintableErc20FactoryOptimismGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// OptimismMintableERC721Factory
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000017)
 */
export const optimismMintableErc721FactoryABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_bridge', internalType: 'address', type: 'address' },
      { name: '_remoteChainId', internalType: 'uint256', type: 'uint256' },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'deployer',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'OptimismMintableERC721Created',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'bridge',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
    ],
    name: 'createOptimismMintableERC721',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'address', type: 'address' }],
    name: 'isOptimismMintableERC721',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'remoteChainId',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000017)
 */
export const optimismMintableErc721FactoryAddress = {
  10: '0x4200000000000000000000000000000000000017',
} as const

/**
 * [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000017)
 */
export const optimismMintableErc721FactoryConfig = {
  address: optimismMintableErc721FactoryAddress,
  abi: optimismMintableErc721FactoryABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// OptimismMintableERC721Factory_optimism-goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000017)
 */
export const optimismMintableErc721FactoryOptimismGoerliABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_bridge', internalType: 'address', type: 'address' },
      { name: '_remoteChainId', internalType: 'uint256', type: 'uint256' },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'localToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'remoteToken',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'deployer',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'OptimismMintableERC721Created',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'BRIDGE',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'REMOTE_CHAIN_ID',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_remoteToken', internalType: 'address', type: 'address' },
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
    ],
    name: 'createOptimismMintableERC721',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'address', type: 'address' }],
    name: 'isOptimismMintableERC721',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000017)
 */
export const optimismMintableErc721FactoryOptimismGoerliAddress = {
  420: '0x4200000000000000000000000000000000000017',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000017)
 */
export const optimismMintableErc721FactoryOptimismGoerliConfig = {
  address: optimismMintableErc721FactoryOptimismGoerliAddress,
  abi: optimismMintableErc721FactoryOptimismGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// OptimismPortal
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xbEb5Fc579115071764c7423A4f12eDde41f106Ed)
 */
export const optimismPortalABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      {
        name: '_l2Oracle',
        internalType: 'contract L2OutputOracle',
        type: 'address',
      },
      { name: '_guardian', internalType: 'address', type: 'address' },
      { name: '_paused', internalType: 'bool', type: 'bool' },
      {
        name: '_config',
        internalType: 'contract SystemConfig',
        type: 'address',
      },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'account',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'Paused',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'version',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'opaqueData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'TransactionDeposited',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'account',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'Unpaused',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawalHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
      { name: 'success', internalType: 'bool', type: 'bool', indexed: false },
    ],
    name: 'WithdrawalFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawalHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
    ],
    name: 'WithdrawalProven',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'GUARDIAN',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'L2_ORACLE',
    outputs: [
      { name: '', internalType: 'contract L2OutputOracle', type: 'address' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'SYSTEM_CONFIG',
    outputs: [
      { name: '', internalType: 'contract SystemConfig', type: 'address' },
    ],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_value', internalType: 'uint256', type: 'uint256' },
      { name: '_gasLimit', internalType: 'uint64', type: 'uint64' },
      { name: '_isCreation', internalType: 'bool', type: 'bool' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'depositTransaction',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [],
    name: 'donateETH',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_tx',
        internalType: 'struct Types.WithdrawalTransaction',
        type: 'tuple',
        components: [
          { name: 'nonce', internalType: 'uint256', type: 'uint256' },
          { name: 'sender', internalType: 'address', type: 'address' },
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'gasLimit', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'finalizeWithdrawalTransaction',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'finalizedWithdrawals',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_paused', internalType: 'bool', type: 'bool' }],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_l2OutputIndex', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'isOutputFinalized',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l2Sender',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [{ name: '_byteCount', internalType: 'uint64', type: 'uint64' }],
    name: 'minimumGasLimit',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'params',
    outputs: [
      { name: 'prevBaseFee', internalType: 'uint128', type: 'uint128' },
      { name: 'prevBoughtGas', internalType: 'uint64', type: 'uint64' },
      { name: 'prevBlockNum', internalType: 'uint64', type: 'uint64' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'pause',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'paused',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_tx',
        internalType: 'struct Types.WithdrawalTransaction',
        type: 'tuple',
        components: [
          { name: 'nonce', internalType: 'uint256', type: 'uint256' },
          { name: 'sender', internalType: 'address', type: 'address' },
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'gasLimit', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '_l2OutputIndex', internalType: 'uint256', type: 'uint256' },
      {
        name: '_outputRootProof',
        internalType: 'struct Types.OutputRootProof',
        type: 'tuple',
        components: [
          { name: 'version', internalType: 'bytes32', type: 'bytes32' },
          { name: 'stateRoot', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'messagePasserStorageRoot',
            internalType: 'bytes32',
            type: 'bytes32',
          },
          { name: 'latestBlockhash', internalType: 'bytes32', type: 'bytes32' },
        ],
      },
      { name: '_withdrawalProof', internalType: 'bytes[]', type: 'bytes[]' },
    ],
    name: 'proveWithdrawalTransaction',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'provenWithdrawals',
    outputs: [
      { name: 'outputRoot', internalType: 'bytes32', type: 'bytes32' },
      { name: 'timestamp', internalType: 'uint128', type: 'uint128' },
      { name: 'l2OutputIndex', internalType: 'uint128', type: 'uint128' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'unpause',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xbEb5Fc579115071764c7423A4f12eDde41f106Ed)
 */
export const optimismPortalAddress = {
  1: '0xbEb5Fc579115071764c7423A4f12eDde41f106Ed',
} as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xbEb5Fc579115071764c7423A4f12eDde41f106Ed)
 */
export const optimismPortalConfig = {
  address: optimismPortalAddress,
  abi: optimismPortalABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// OptimismPortal_goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383)
 */
export const optimismPortalGoerliABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'account',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'Paused',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'version',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'opaqueData',
        internalType: 'bytes',
        type: 'bytes',
        indexed: false,
      },
    ],
    name: 'TransactionDeposited',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'account',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'Unpaused',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawalHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
      { name: 'success', internalType: 'bool', type: 'bool', indexed: false },
    ],
    name: 'WithdrawalFinalized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawalHash',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
    ],
    name: 'WithdrawalProven',
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_value', internalType: 'uint256', type: 'uint256' },
      { name: '_gasLimit', internalType: 'uint64', type: 'uint64' },
      { name: '_isCreation', internalType: 'bool', type: 'bool' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'depositTransaction',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [],
    name: 'donateETH',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_tx',
        internalType: 'struct Types.WithdrawalTransaction',
        type: 'tuple',
        components: [
          { name: 'nonce', internalType: 'uint256', type: 'uint256' },
          { name: 'sender', internalType: 'address', type: 'address' },
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'gasLimit', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'finalizeWithdrawalTransaction',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'finalizedWithdrawals',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'guardian',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_l2Oracle',
        internalType: 'contract L2OutputOracle',
        type: 'address',
      },
      { name: '_guardian', internalType: 'address', type: 'address' },
      {
        name: '_systemConfig',
        internalType: 'contract SystemConfig',
        type: 'address',
      },
      { name: '_paused', internalType: 'bool', type: 'bool' },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_l2OutputIndex', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'isOutputFinalized',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l2Oracle',
    outputs: [
      { name: '', internalType: 'contract L2OutputOracle', type: 'address' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l2Sender',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [{ name: '_byteCount', internalType: 'uint64', type: 'uint64' }],
    name: 'minimumGasLimit',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'params',
    outputs: [
      { name: 'prevBaseFee', internalType: 'uint128', type: 'uint128' },
      { name: 'prevBoughtGas', internalType: 'uint64', type: 'uint64' },
      { name: 'prevBlockNum', internalType: 'uint64', type: 'uint64' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'pause',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'paused',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_tx',
        internalType: 'struct Types.WithdrawalTransaction',
        type: 'tuple',
        components: [
          { name: 'nonce', internalType: 'uint256', type: 'uint256' },
          { name: 'sender', internalType: 'address', type: 'address' },
          { name: 'target', internalType: 'address', type: 'address' },
          { name: 'value', internalType: 'uint256', type: 'uint256' },
          { name: 'gasLimit', internalType: 'uint256', type: 'uint256' },
          { name: 'data', internalType: 'bytes', type: 'bytes' },
        ],
      },
      { name: '_l2OutputIndex', internalType: 'uint256', type: 'uint256' },
      {
        name: '_outputRootProof',
        internalType: 'struct Types.OutputRootProof',
        type: 'tuple',
        components: [
          { name: 'version', internalType: 'bytes32', type: 'bytes32' },
          { name: 'stateRoot', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'messagePasserStorageRoot',
            internalType: 'bytes32',
            type: 'bytes32',
          },
          { name: 'latestBlockhash', internalType: 'bytes32', type: 'bytes32' },
        ],
      },
      { name: '_withdrawalProof', internalType: 'bytes[]', type: 'bytes[]' },
    ],
    name: 'proveWithdrawalTransaction',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'provenWithdrawals',
    outputs: [
      { name: 'outputRoot', internalType: 'bytes32', type: 'bytes32' },
      { name: 'timestamp', internalType: 'uint128', type: 'uint128' },
      { name: 'l2OutputIndex', internalType: 'uint128', type: 'uint128' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'systemConfig',
    outputs: [
      { name: '', internalType: 'contract SystemConfig', type: 'address' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'unpause',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383)
 */
export const optimismPortalGoerliAddress = {
  5: '0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383',
} as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383)
 */
export const optimismPortalGoerliConfig = {
  address: optimismPortalGoerliAddress,
  abi: optimismPortalGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Optimist
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 */
export const optimistABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
      { name: '_baseURIAttestor', internalType: 'address', type: 'address' },
      {
        name: '_attestationStation',
        internalType: 'contract AttestationStation',
        type: 'address',
      },
      {
        name: '_optimistAllowlist',
        internalType: 'contract OptimistAllowlist',
        type: 'address',
      },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'owner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'approved',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'tokenId',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
    ],
    name: 'Approval',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'owner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'operator',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'approved', internalType: 'bool', type: 'bool', indexed: false },
    ],
    name: 'ApprovalForAll',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'tokenId',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
    ],
    name: 'Transfer',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'ATTESTATION_STATION',
    outputs: [
      {
        name: '',
        internalType: 'contract AttestationStation',
        type: 'address',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'BASE_URI_ATTESTATION_KEY',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'BASE_URI_ATTESTOR',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'OPTIMIST_ALLOWLIST',
    outputs: [
      { name: '', internalType: 'contract OptimistAllowlist', type: 'address' },
    ],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'approve',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: 'owner', internalType: 'address', type: 'address' }],
    name: 'balanceOf',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'baseURI',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'tokenId', internalType: 'uint256', type: 'uint256' }],
    name: 'burn',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: 'tokenId', internalType: 'uint256', type: 'uint256' }],
    name: 'getApproved',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_symbol', internalType: 'string', type: 'string' },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: 'owner', internalType: 'address', type: 'address' },
      { name: 'operator', internalType: 'address', type: 'address' },
    ],
    name: 'isApprovedForAll',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_recipient', internalType: 'address', type: 'address' }],
    name: 'isOnAllowList',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_recipient', internalType: 'address', type: 'address' }],
    name: 'mint',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'name',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: 'tokenId', internalType: 'uint256', type: 'uint256' }],
    name: 'ownerOf',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: 'from', internalType: 'address', type: 'address' },
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'tokenId', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'safeTransferFrom',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: 'from', internalType: 'address', type: 'address' },
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'tokenId', internalType: 'uint256', type: 'uint256' },
      { name: 'data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'safeTransferFrom',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'bool', type: 'bool' },
    ],
    name: 'setApprovalForAll',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: 'interfaceId', internalType: 'bytes4', type: 'bytes4' }],
    name: 'supportsInterface',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'symbol',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'pure',
    type: 'function',
    inputs: [{ name: '_owner', internalType: 'address', type: 'address' }],
    name: 'tokenIdOfAddress',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_tokenId', internalType: 'uint256', type: 'uint256' }],
    name: 'tokenURI',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: 'from', internalType: 'address', type: 'address' },
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'tokenId', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'transferFrom',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 */
export const optimistAddress = {
  10: '0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5',
  420: '0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5',
} as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 */
export const optimistConfig = {
  address: optimistAddress,
  abi: optimistABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// OptimistAllowlist
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180)
 */
export const optimistAllowlistABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      {
        name: '_attestationStation',
        internalType: 'contract AttestationStation',
        type: 'address',
      },
      { name: '_allowlistAttestor', internalType: 'address', type: 'address' },
      {
        name: '_coinbaseQuestAttestor',
        internalType: 'address',
        type: 'address',
      },
      { name: '_optimistInviter', internalType: 'address', type: 'address' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'ALLOWLIST_ATTESTOR',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'ATTESTATION_STATION',
    outputs: [
      {
        name: '',
        internalType: 'contract AttestationStation',
        type: 'address',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'COINBASE_QUEST_ATTESTOR',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'COINBASE_QUEST_ELIGIBLE_ATTESTATION_KEY',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'OPTIMIST_CAN_MINT_ATTESTATION_KEY',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'OPTIMIST_INVITER',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_claimer', internalType: 'address', type: 'address' }],
    name: 'isAllowedToMint',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180)
 */
export const optimistAllowlistAddress = {
  10: '0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180',
  420: '0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180',
} as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180)
 */
export const optimistAllowlistConfig = {
  address: optimistAllowlistAddress,
  abi: optimistAllowlistABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// OptimistInviter
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 */
export const optimistInviterABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_inviteGranter', internalType: 'address', type: 'address' },
      {
        name: '_attestationStation',
        internalType: 'contract AttestationStation',
        type: 'address',
      },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'issuer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'claimer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'InviteClaimed',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'ATTESTATION_STATION',
    outputs: [
      {
        name: '',
        internalType: 'contract AttestationStation',
        type: 'address',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'CAN_INVITE_ATTESTATION_KEY',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'CLAIMABLE_INVITE_TYPEHASH',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'EIP712_VERSION',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'INVITE_GRANTER',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_COMMITMENT_PERIOD',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_claimer', internalType: 'address', type: 'address' },
      {
        name: '_claimableInvite',
        internalType: 'struct OptimistInviter.ClaimableInvite',
        type: 'tuple',
        components: [
          { name: 'issuer', internalType: 'address', type: 'address' },
          { name: 'nonce', internalType: 'bytes32', type: 'bytes32' },
        ],
      },
      { name: '_signature', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'claimInvite',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_commitment', internalType: 'bytes32', type: 'bytes32' }],
    name: 'commitInvite',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
    name: 'commitmentTimestamps',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_name', internalType: 'string', type: 'string' }],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'address', type: 'address' }],
    name: 'inviteCounts',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_accounts', internalType: 'address[]', type: 'address[]' },
      { name: '_inviteCount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'setInviteCounts',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '', internalType: 'address', type: 'address' },
      { name: '', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'usedNonces',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 */
export const optimistInviterAddress = {
  10: '0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929',
  420: '0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929',
} as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 */
export const optimistInviterConfig = {
  address: optimistInviterAddress,
  abi: optimistInviterABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PortalSender
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x0A893d9576b9cFD9EF78595963dc973238E78210)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xe7FACd39531ee3C313330E93B4d7a8B8A3c84Aa4)
 */
export const portalSenderABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      {
        name: '_portal',
        internalType: 'contract OptimismPortal',
        type: 'address',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'PORTAL',
    outputs: [
      { name: '', internalType: 'contract OptimismPortal', type: 'address' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'donate',
    outputs: [],
  },
] as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x0A893d9576b9cFD9EF78595963dc973238E78210)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xe7FACd39531ee3C313330E93B4d7a8B8A3c84Aa4)
 */
export const portalSenderAddress = {
  1: '0x0A893d9576b9cFD9EF78595963dc973238E78210',
  5: '0xe7FACd39531ee3C313330E93B4d7a8B8A3c84Aa4',
} as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x0A893d9576b9cFD9EF78595963dc973238E78210)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xe7FACd39531ee3C313330E93B4d7a8B8A3c84Aa4)
 */
export const portalSenderConfig = {
  address: portalSenderAddress,
  abi: portalSenderABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ProtocolVersions
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x0C24F5098774aA366827D667494e9F889f7cFc08)
 */
export const protocolVersionsABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'version',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'updateType',
        internalType: 'enum ProtocolVersions.UpdateType',
        type: 'uint8',
        indexed: true,
      },
      { name: 'data', internalType: 'bytes', type: 'bytes', indexed: false },
    ],
    name: 'ConfigUpdate',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RECOMMENDED_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'REQUIRED_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_owner', internalType: 'address', type: 'address' },
      { name: '_required', internalType: 'ProtocolVersion', type: 'uint256' },
      {
        name: '_recommended',
        internalType: 'ProtocolVersion',
        type: 'uint256',
      },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'recommended',
    outputs: [
      { name: 'out_', internalType: 'ProtocolVersion', type: 'uint256' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'required',
    outputs: [
      { name: 'out_', internalType: 'ProtocolVersion', type: 'uint256' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_recommended',
        internalType: 'ProtocolVersion',
        type: 'uint256',
      },
    ],
    name: 'setRecommended',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_required', internalType: 'ProtocolVersion', type: 'uint256' },
    ],
    name: 'setRequired',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x0C24F5098774aA366827D667494e9F889f7cFc08)
 */
export const protocolVersionsAddress = {
  5: '0x0C24F5098774aA366827D667494e9F889f7cFc08',
} as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x0C24F5098774aA366827D667494e9F889f7cFc08)
 */
export const protocolVersionsConfig = {
  address: protocolVersionsAddress,
  abi: protocolVersionsABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ProxyAdmin
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000018)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000018)
 */
export const proxyAdminABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_owner', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'addressManager',
    outputs: [
      { name: '', internalType: 'contract AddressManager', type: 'address' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_proxy', internalType: 'address payable', type: 'address' },
      { name: '_newAdmin', internalType: 'address', type: 'address' },
    ],
    name: 'changeProxyAdmin',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [
      { name: '_proxy', internalType: 'address payable', type: 'address' },
    ],
    name: 'getProxyAdmin',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '_proxy', internalType: 'address', type: 'address' }],
    name: 'getProxyImplementation',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'address', type: 'address' }],
    name: 'implementationName',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'isUpgrading',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: '', internalType: 'address', type: 'address' }],
    name: 'proxyType',
    outputs: [
      { name: '', internalType: 'enum ProxyAdmin.ProxyType', type: 'uint8' },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_name', internalType: 'string', type: 'string' },
      { name: '_address', internalType: 'address', type: 'address' },
    ],
    name: 'setAddress',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_address',
        internalType: 'contract AddressManager',
        type: 'address',
      },
    ],
    name: 'setAddressManager',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_address', internalType: 'address', type: 'address' },
      { name: '_name', internalType: 'string', type: 'string' },
    ],
    name: 'setImplementationName',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_address', internalType: 'address', type: 'address' },
      {
        name: '_type',
        internalType: 'enum ProxyAdmin.ProxyType',
        type: 'uint8',
      },
    ],
    name: 'setProxyType',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_upgrading', internalType: 'bool', type: 'bool' }],
    name: 'setUpgrading',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_proxy', internalType: 'address payable', type: 'address' },
      { name: '_implementation', internalType: 'address', type: 'address' },
    ],
    name: 'upgrade',
    outputs: [],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_proxy', internalType: 'address payable', type: 'address' },
      { name: '_implementation', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'upgradeAndCall',
    outputs: [],
  },
] as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000018)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000018)
 */
export const proxyAdminAddress = {
  1: '0x4200000000000000000000000000000000000018',
  5: '0x4200000000000000000000000000000000000018',
} as const

/**
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000018)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000018)
 */
export const proxyAdminConfig = {
  address: proxyAdminAddress,
  abi: proxyAdminABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// SchemaRegistry
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x6232208d66bAc2305b46b4Cb6BCB3857B298DF13)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2545fa928d5d278cA75Fd47306e4a89096ff6403)
 */
export const schemaRegistryABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  { type: 'error', inputs: [], name: 'AlreadyExists' },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'uid', internalType: 'bytes32', type: 'bytes32', indexed: true },
      {
        name: 'registerer',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'Registered',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [{ name: 'uid', internalType: 'bytes32', type: 'bytes32' }],
    name: 'getSchema',
    outputs: [
      {
        name: '',
        internalType: 'struct SchemaRecord',
        type: 'tuple',
        components: [
          { name: 'uid', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'resolver',
            internalType: 'contract ISchemaResolver',
            type: 'address',
          },
          { name: 'revocable', internalType: 'bool', type: 'bool' },
          { name: 'schema', internalType: 'string', type: 'string' },
        ],
      },
    ],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: 'schema', internalType: 'string', type: 'string' },
      {
        name: 'resolver',
        internalType: 'contract ISchemaResolver',
        type: 'address',
      },
      { name: 'revocable', internalType: 'bool', type: 'bool' },
    ],
    name: 'register',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x6232208d66bAc2305b46b4Cb6BCB3857B298DF13)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2545fa928d5d278cA75Fd47306e4a89096ff6403)
 */
export const schemaRegistryAddress = {
  10: '0x6232208d66bAc2305b46b4Cb6BCB3857B298DF13',
  420: '0x2545fa928d5d278cA75Fd47306e4a89096ff6403',
} as const

/**
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x6232208d66bAc2305b46b4Cb6BCB3857B298DF13)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2545fa928d5d278cA75Fd47306e4a89096ff6403)
 */
export const schemaRegistryConfig = {
  address: schemaRegistryAddress,
  abi: schemaRegistryABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// SequencerFeeVault
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000011)
 */
export const sequencerFeeVaultABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_recipient', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      { name: 'to', internalType: 'address', type: 'address', indexed: false },
      {
        name: 'from',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'Withdrawal',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'MIN_WITHDRAWAL_AMOUNT',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'RECIPIENT',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l1FeeWallet',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'totalProcessed',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'withdraw',
    outputs: [],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000011)
 */
export const sequencerFeeVaultAddress = {
  420: '0x4200000000000000000000000000000000000011',
} as const

/**
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000011)
 */
export const sequencerFeeVaultConfig = {
  address: sequencerFeeVaultAddress,
  abi: sequencerFeeVaultABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// SystemConfig
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x229047fed2591dbec1eF1118d64F7aF3dB9EB290)
 */
export const systemConfigABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [
      { name: '_owner', internalType: 'address', type: 'address' },
      { name: '_overhead', internalType: 'uint256', type: 'uint256' },
      { name: '_scalar', internalType: 'uint256', type: 'uint256' },
      { name: '_batcherHash', internalType: 'bytes32', type: 'bytes32' },
      { name: '_gasLimit', internalType: 'uint64', type: 'uint64' },
      { name: '_unsafeBlockSigner', internalType: 'address', type: 'address' },
      {
        name: '_config',
        internalType: 'struct ResourceMetering.ResourceConfig',
        type: 'tuple',
        components: [
          { name: 'maxResourceLimit', internalType: 'uint32', type: 'uint32' },
          {
            name: 'elasticityMultiplier',
            internalType: 'uint8',
            type: 'uint8',
          },
          {
            name: 'baseFeeMaxChangeDenominator',
            internalType: 'uint8',
            type: 'uint8',
          },
          { name: 'minimumBaseFee', internalType: 'uint32', type: 'uint32' },
          { name: 'systemTxMaxGas', internalType: 'uint32', type: 'uint32' },
          { name: 'maximumBaseFee', internalType: 'uint128', type: 'uint128' },
        ],
      },
    ],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'version',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'updateType',
        internalType: 'enum SystemConfig.UpdateType',
        type: 'uint8',
        indexed: true,
      },
      { name: 'data', internalType: 'bytes', type: 'bytes', indexed: false },
    ],
    name: 'ConfigUpdate',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'UNSAFE_BLOCK_SIGNER_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'batcherHash',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'gasLimit',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_owner', internalType: 'address', type: 'address' },
      { name: '_overhead', internalType: 'uint256', type: 'uint256' },
      { name: '_scalar', internalType: 'uint256', type: 'uint256' },
      { name: '_batcherHash', internalType: 'bytes32', type: 'bytes32' },
      { name: '_gasLimit', internalType: 'uint64', type: 'uint64' },
      { name: '_unsafeBlockSigner', internalType: 'address', type: 'address' },
      {
        name: '_config',
        internalType: 'struct ResourceMetering.ResourceConfig',
        type: 'tuple',
        components: [
          { name: 'maxResourceLimit', internalType: 'uint32', type: 'uint32' },
          {
            name: 'elasticityMultiplier',
            internalType: 'uint8',
            type: 'uint8',
          },
          {
            name: 'baseFeeMaxChangeDenominator',
            internalType: 'uint8',
            type: 'uint8',
          },
          { name: 'minimumBaseFee', internalType: 'uint32', type: 'uint32' },
          { name: 'systemTxMaxGas', internalType: 'uint32', type: 'uint32' },
          { name: 'maximumBaseFee', internalType: 'uint128', type: 'uint128' },
        ],
      },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'minimumGasLimit',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'overhead',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'resourceConfig',
    outputs: [
      {
        name: '',
        internalType: 'struct ResourceMetering.ResourceConfig',
        type: 'tuple',
        components: [
          { name: 'maxResourceLimit', internalType: 'uint32', type: 'uint32' },
          {
            name: 'elasticityMultiplier',
            internalType: 'uint8',
            type: 'uint8',
          },
          {
            name: 'baseFeeMaxChangeDenominator',
            internalType: 'uint8',
            type: 'uint8',
          },
          { name: 'minimumBaseFee', internalType: 'uint32', type: 'uint32' },
          { name: 'systemTxMaxGas', internalType: 'uint32', type: 'uint32' },
          { name: 'maximumBaseFee', internalType: 'uint128', type: 'uint128' },
        ],
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'scalar',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_batcherHash', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'setBatcherHash',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_overhead', internalType: 'uint256', type: 'uint256' },
      { name: '_scalar', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'setGasConfig',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_gasLimit', internalType: 'uint64', type: 'uint64' }],
    name: 'setGasLimit',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_config',
        internalType: 'struct ResourceMetering.ResourceConfig',
        type: 'tuple',
        components: [
          { name: 'maxResourceLimit', internalType: 'uint32', type: 'uint32' },
          {
            name: 'elasticityMultiplier',
            internalType: 'uint8',
            type: 'uint8',
          },
          {
            name: 'baseFeeMaxChangeDenominator',
            internalType: 'uint8',
            type: 'uint8',
          },
          { name: 'minimumBaseFee', internalType: 'uint32', type: 'uint32' },
          { name: 'systemTxMaxGas', internalType: 'uint32', type: 'uint32' },
          { name: 'maximumBaseFee', internalType: 'uint128', type: 'uint128' },
        ],
      },
    ],
    name: 'setResourceConfig',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_unsafeBlockSigner', internalType: 'address', type: 'address' },
    ],
    name: 'setUnsafeBlockSigner',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'unsafeBlockSigner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x229047fed2591dbec1eF1118d64F7aF3dB9EB290)
 */
export const systemConfigAddress = {
  1: '0x229047fed2591dbec1eF1118d64F7aF3dB9EB290',
} as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x229047fed2591dbec1eF1118d64F7aF3dB9EB290)
 */
export const systemConfigConfig = {
  address: systemConfigAddress,
  abi: systemConfigABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// SystemConfig_goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60)
 */
export const systemConfigGoerliABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'version',
        internalType: 'uint256',
        type: 'uint256',
        indexed: true,
      },
      {
        name: 'updateType',
        internalType: 'enum SystemConfig.UpdateType',
        type: 'uint8',
        indexed: true,
      },
      { name: 'data', internalType: 'bytes', type: 'bytes', indexed: false },
    ],
    name: 'ConfigUpdate',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'BATCH_INBOX_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'L1_CROSS_DOMAIN_MESSENGER_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'L1_ERC_721_BRIDGE_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'L1_STANDARD_BRIDGE_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'L2_OUTPUT_ORACLE_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'OPTIMISM_MINTABLE_ERC20_FACTORY_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'OPTIMISM_PORTAL_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'UNSAFE_BLOCK_SIGNER_SLOT',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'batchInbox',
    outputs: [{ name: 'addr_', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'batcherHash',
    outputs: [{ name: '', internalType: 'bytes32', type: 'bytes32' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'gasLimit',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_owner', internalType: 'address', type: 'address' },
      { name: '_overhead', internalType: 'uint256', type: 'uint256' },
      { name: '_scalar', internalType: 'uint256', type: 'uint256' },
      { name: '_batcherHash', internalType: 'bytes32', type: 'bytes32' },
      { name: '_gasLimit', internalType: 'uint64', type: 'uint64' },
      { name: '_unsafeBlockSigner', internalType: 'address', type: 'address' },
      {
        name: '_config',
        internalType: 'struct ResourceMetering.ResourceConfig',
        type: 'tuple',
        components: [
          { name: 'maxResourceLimit', internalType: 'uint32', type: 'uint32' },
          {
            name: 'elasticityMultiplier',
            internalType: 'uint8',
            type: 'uint8',
          },
          {
            name: 'baseFeeMaxChangeDenominator',
            internalType: 'uint8',
            type: 'uint8',
          },
          { name: 'minimumBaseFee', internalType: 'uint32', type: 'uint32' },
          { name: 'systemTxMaxGas', internalType: 'uint32', type: 'uint32' },
          { name: 'maximumBaseFee', internalType: 'uint128', type: 'uint128' },
        ],
      },
      { name: '_startBlock', internalType: 'uint256', type: 'uint256' },
      { name: '_batchInbox', internalType: 'address', type: 'address' },
      {
        name: '_addresses',
        internalType: 'struct SystemConfig.Addresses',
        type: 'tuple',
        components: [
          {
            name: 'l1CrossDomainMessenger',
            internalType: 'address',
            type: 'address',
          },
          { name: 'l1ERC721Bridge', internalType: 'address', type: 'address' },
          {
            name: 'l1StandardBridge',
            internalType: 'address',
            type: 'address',
          },
          { name: 'l2OutputOracle', internalType: 'address', type: 'address' },
          { name: 'optimismPortal', internalType: 'address', type: 'address' },
          {
            name: 'optimismMintableERC20Factory',
            internalType: 'address',
            type: 'address',
          },
        ],
      },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l1CrossDomainMessenger',
    outputs: [{ name: 'addr_', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l1ERC721Bridge',
    outputs: [{ name: 'addr_', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l1StandardBridge',
    outputs: [{ name: 'addr_', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l2OutputOracle',
    outputs: [{ name: 'addr_', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'minimumGasLimit',
    outputs: [{ name: '', internalType: 'uint64', type: 'uint64' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'optimismMintableERC20Factory',
    outputs: [{ name: 'addr_', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'optimismPortal',
    outputs: [{ name: 'addr_', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'overhead',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'resourceConfig',
    outputs: [
      {
        name: '',
        internalType: 'struct ResourceMetering.ResourceConfig',
        type: 'tuple',
        components: [
          { name: 'maxResourceLimit', internalType: 'uint32', type: 'uint32' },
          {
            name: 'elasticityMultiplier',
            internalType: 'uint8',
            type: 'uint8',
          },
          {
            name: 'baseFeeMaxChangeDenominator',
            internalType: 'uint8',
            type: 'uint8',
          },
          { name: 'minimumBaseFee', internalType: 'uint32', type: 'uint32' },
          { name: 'systemTxMaxGas', internalType: 'uint32', type: 'uint32' },
          { name: 'maximumBaseFee', internalType: 'uint128', type: 'uint128' },
        ],
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'scalar',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_batcherHash', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'setBatcherHash',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_overhead', internalType: 'uint256', type: 'uint256' },
      { name: '_scalar', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'setGasConfig',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_gasLimit', internalType: 'uint64', type: 'uint64' }],
    name: 'setGasLimit',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_config',
        internalType: 'struct ResourceMetering.ResourceConfig',
        type: 'tuple',
        components: [
          { name: 'maxResourceLimit', internalType: 'uint32', type: 'uint32' },
          {
            name: 'elasticityMultiplier',
            internalType: 'uint8',
            type: 'uint8',
          },
          {
            name: 'baseFeeMaxChangeDenominator',
            internalType: 'uint8',
            type: 'uint8',
          },
          { name: 'minimumBaseFee', internalType: 'uint32', type: 'uint32' },
          { name: 'systemTxMaxGas', internalType: 'uint32', type: 'uint32' },
          { name: 'maximumBaseFee', internalType: 'uint128', type: 'uint128' },
        ],
      },
    ],
    name: 'setResourceConfig',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_unsafeBlockSigner', internalType: 'address', type: 'address' },
    ],
    name: 'setUnsafeBlockSigner',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'startBlock',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'unsafeBlockSigner',
    outputs: [{ name: 'addr_', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
  },
] as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60)
 */
export const systemConfigGoerliAddress = {
  5: '0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60',
} as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60)
 */
export const systemConfigGoerliConfig = {
  address: systemConfigGoerliAddress,
  abi: systemConfigGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// SystemDictator
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xB4453CEb33d2e67FA244A24acf2E50CEF31F53cB)
 */
export const systemDictatorABI = [
  { stateMutability: 'nonpayable', type: 'constructor', inputs: [] },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'EXIT_1_NO_RETURN_STEP',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'PROXY_TRANSFER_STEP',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'config',
    outputs: [
      {
        name: 'globalConfig',
        internalType: 'struct SystemDictator.GlobalConfig',
        type: 'tuple',
        components: [
          {
            name: 'addressManager',
            internalType: 'contract AddressManager',
            type: 'address',
          },
          {
            name: 'proxyAdmin',
            internalType: 'contract ProxyAdmin',
            type: 'address',
          },
          { name: 'controller', internalType: 'address', type: 'address' },
          { name: 'finalOwner', internalType: 'address', type: 'address' },
        ],
      },
      {
        name: 'proxyAddressConfig',
        internalType: 'struct SystemDictator.ProxyAddressConfig',
        type: 'tuple',
        components: [
          {
            name: 'l2OutputOracleProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'optimismPortalProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'l1CrossDomainMessengerProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'l1StandardBridgeProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'optimismMintableERC20FactoryProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'l1ERC721BridgeProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'systemConfigProxy',
            internalType: 'address',
            type: 'address',
          },
        ],
      },
      {
        name: 'implementationAddressConfig',
        internalType: 'struct SystemDictator.ImplementationAddressConfig',
        type: 'tuple',
        components: [
          {
            name: 'l2OutputOracleImpl',
            internalType: 'contract L2OutputOracle',
            type: 'address',
          },
          {
            name: 'optimismPortalImpl',
            internalType: 'contract OptimismPortal',
            type: 'address',
          },
          {
            name: 'l1CrossDomainMessengerImpl',
            internalType: 'contract L1CrossDomainMessenger',
            type: 'address',
          },
          {
            name: 'l1StandardBridgeImpl',
            internalType: 'contract L1StandardBridge',
            type: 'address',
          },
          {
            name: 'optimismMintableERC20FactoryImpl',
            internalType: 'contract OptimismMintableERC20Factory',
            type: 'address',
          },
          {
            name: 'l1ERC721BridgeImpl',
            internalType: 'contract L1ERC721Bridge',
            type: 'address',
          },
          {
            name: 'portalSenderImpl',
            internalType: 'contract PortalSender',
            type: 'address',
          },
          {
            name: 'systemConfigImpl',
            internalType: 'contract SystemConfig',
            type: 'address',
          },
        ],
      },
      {
        name: 'systemConfigConfig',
        internalType: 'struct SystemDictator.SystemConfigConfig',
        type: 'tuple',
        components: [
          { name: 'owner', internalType: 'address', type: 'address' },
          { name: 'overhead', internalType: 'uint256', type: 'uint256' },
          { name: 'scalar', internalType: 'uint256', type: 'uint256' },
          { name: 'batcherHash', internalType: 'bytes32', type: 'bytes32' },
          { name: 'gasLimit', internalType: 'uint64', type: 'uint64' },
          {
            name: 'unsafeBlockSigner',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'resourceConfig',
            internalType: 'struct ResourceMetering.ResourceConfig',
            type: 'tuple',
            components: [
              {
                name: 'maxResourceLimit',
                internalType: 'uint32',
                type: 'uint32',
              },
              {
                name: 'elasticityMultiplier',
                internalType: 'uint8',
                type: 'uint8',
              },
              {
                name: 'baseFeeMaxChangeDenominator',
                internalType: 'uint8',
                type: 'uint8',
              },
              {
                name: 'minimumBaseFee',
                internalType: 'uint32',
                type: 'uint32',
              },
              {
                name: 'systemTxMaxGas',
                internalType: 'uint32',
                type: 'uint32',
              },
              {
                name: 'maximumBaseFee',
                internalType: 'uint128',
                type: 'uint128',
              },
            ],
          },
        ],
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'currentStep',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'dynamicConfigSet',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'exit1',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'exited',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'finalize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'finalized',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_config',
        internalType: 'struct SystemDictator.DeployConfig',
        type: 'tuple',
        components: [
          {
            name: 'globalConfig',
            internalType: 'struct SystemDictator.GlobalConfig',
            type: 'tuple',
            components: [
              {
                name: 'addressManager',
                internalType: 'contract AddressManager',
                type: 'address',
              },
              {
                name: 'proxyAdmin',
                internalType: 'contract ProxyAdmin',
                type: 'address',
              },
              { name: 'controller', internalType: 'address', type: 'address' },
              { name: 'finalOwner', internalType: 'address', type: 'address' },
            ],
          },
          {
            name: 'proxyAddressConfig',
            internalType: 'struct SystemDictator.ProxyAddressConfig',
            type: 'tuple',
            components: [
              {
                name: 'l2OutputOracleProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'optimismPortalProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'l1CrossDomainMessengerProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'l1StandardBridgeProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'optimismMintableERC20FactoryProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'l1ERC721BridgeProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'systemConfigProxy',
                internalType: 'address',
                type: 'address',
              },
            ],
          },
          {
            name: 'implementationAddressConfig',
            internalType: 'struct SystemDictator.ImplementationAddressConfig',
            type: 'tuple',
            components: [
              {
                name: 'l2OutputOracleImpl',
                internalType: 'contract L2OutputOracle',
                type: 'address',
              },
              {
                name: 'optimismPortalImpl',
                internalType: 'contract OptimismPortal',
                type: 'address',
              },
              {
                name: 'l1CrossDomainMessengerImpl',
                internalType: 'contract L1CrossDomainMessenger',
                type: 'address',
              },
              {
                name: 'l1StandardBridgeImpl',
                internalType: 'contract L1StandardBridge',
                type: 'address',
              },
              {
                name: 'optimismMintableERC20FactoryImpl',
                internalType: 'contract OptimismMintableERC20Factory',
                type: 'address',
              },
              {
                name: 'l1ERC721BridgeImpl',
                internalType: 'contract L1ERC721Bridge',
                type: 'address',
              },
              {
                name: 'portalSenderImpl',
                internalType: 'contract PortalSender',
                type: 'address',
              },
              {
                name: 'systemConfigImpl',
                internalType: 'contract SystemConfig',
                type: 'address',
              },
            ],
          },
          {
            name: 'systemConfigConfig',
            internalType: 'struct SystemDictator.SystemConfigConfig',
            type: 'tuple',
            components: [
              { name: 'owner', internalType: 'address', type: 'address' },
              { name: 'overhead', internalType: 'uint256', type: 'uint256' },
              { name: 'scalar', internalType: 'uint256', type: 'uint256' },
              { name: 'batcherHash', internalType: 'bytes32', type: 'bytes32' },
              { name: 'gasLimit', internalType: 'uint64', type: 'uint64' },
              {
                name: 'unsafeBlockSigner',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'resourceConfig',
                internalType: 'struct ResourceMetering.ResourceConfig',
                type: 'tuple',
                components: [
                  {
                    name: 'maxResourceLimit',
                    internalType: 'uint32',
                    type: 'uint32',
                  },
                  {
                    name: 'elasticityMultiplier',
                    internalType: 'uint8',
                    type: 'uint8',
                  },
                  {
                    name: 'baseFeeMaxChangeDenominator',
                    internalType: 'uint8',
                    type: 'uint8',
                  },
                  {
                    name: 'minimumBaseFee',
                    internalType: 'uint32',
                    type: 'uint32',
                  },
                  {
                    name: 'systemTxMaxGas',
                    internalType: 'uint32',
                    type: 'uint32',
                  },
                  {
                    name: 'maximumBaseFee',
                    internalType: 'uint128',
                    type: 'uint128',
                  },
                ],
              },
            ],
          },
        ],
      },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l2OutputOracleDynamicConfig',
    outputs: [
      {
        name: 'l2OutputOracleStartingBlockNumber',
        internalType: 'uint256',
        type: 'uint256',
      },
      {
        name: 'l2OutputOracleStartingTimestamp',
        internalType: 'uint256',
        type: 'uint256',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'oldL1CrossDomainMessenger',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'optimismPortalDynamicConfig',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'phase1',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'phase2',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step1',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step2',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step3',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step4',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step5',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_l2OutputOracleDynamicConfig',
        internalType: 'struct SystemDictator.L2OutputOracleDynamicConfig',
        type: 'tuple',
        components: [
          {
            name: 'l2OutputOracleStartingBlockNumber',
            internalType: 'uint256',
            type: 'uint256',
          },
          {
            name: 'l2OutputOracleStartingTimestamp',
            internalType: 'uint256',
            type: 'uint256',
          },
        ],
      },
      {
        name: '_optimismPortalDynamicConfig',
        internalType: 'bool',
        type: 'bool',
      },
    ],
    name: 'updateDynamicConfig',
    outputs: [],
  },
] as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xB4453CEb33d2e67FA244A24acf2E50CEF31F53cB)
 */
export const systemDictatorAddress = {
  1: '0xB4453CEb33d2e67FA244A24acf2E50CEF31F53cB',
} as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xB4453CEb33d2e67FA244A24acf2E50CEF31F53cB)
 */
export const systemDictatorConfig = {
  address: systemDictatorAddress,
  abi: systemDictatorABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// SystemDictator_goerli
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x1f0613A44c9a8ECE7B3A2e0CdBdF0F5B47A50971)
 */
export const systemDictatorGoerliABI = [
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'version', internalType: 'uint8', type: 'uint8', indexed: false },
    ],
    name: 'Initialized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'EXIT_1_NO_RETURN_STEP',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'PROXY_TRANSFER_STEP',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'config',
    outputs: [
      {
        name: 'globalConfig',
        internalType: 'struct SystemDictator.GlobalConfig',
        type: 'tuple',
        components: [
          {
            name: 'addressManager',
            internalType: 'contract AddressManager',
            type: 'address',
          },
          {
            name: 'proxyAdmin',
            internalType: 'contract ProxyAdmin',
            type: 'address',
          },
          { name: 'controller', internalType: 'address', type: 'address' },
          { name: 'finalOwner', internalType: 'address', type: 'address' },
        ],
      },
      {
        name: 'proxyAddressConfig',
        internalType: 'struct SystemDictator.ProxyAddressConfig',
        type: 'tuple',
        components: [
          {
            name: 'l2OutputOracleProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'optimismPortalProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'l1CrossDomainMessengerProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'l1StandardBridgeProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'optimismMintableERC20FactoryProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'l1ERC721BridgeProxy',
            internalType: 'address',
            type: 'address',
          },
          {
            name: 'systemConfigProxy',
            internalType: 'address',
            type: 'address',
          },
        ],
      },
      {
        name: 'implementationAddressConfig',
        internalType: 'struct SystemDictator.ImplementationAddressConfig',
        type: 'tuple',
        components: [
          {
            name: 'l2OutputOracleImpl',
            internalType: 'contract L2OutputOracle',
            type: 'address',
          },
          {
            name: 'optimismPortalImpl',
            internalType: 'contract OptimismPortal',
            type: 'address',
          },
          {
            name: 'l1CrossDomainMessengerImpl',
            internalType: 'contract L1CrossDomainMessenger',
            type: 'address',
          },
          {
            name: 'l1StandardBridgeImpl',
            internalType: 'contract L1StandardBridge',
            type: 'address',
          },
          {
            name: 'optimismMintableERC20FactoryImpl',
            internalType: 'contract OptimismMintableERC20Factory',
            type: 'address',
          },
          {
            name: 'l1ERC721BridgeImpl',
            internalType: 'contract L1ERC721Bridge',
            type: 'address',
          },
          {
            name: 'portalSenderImpl',
            internalType: 'contract PortalSender',
            type: 'address',
          },
          {
            name: 'systemConfigImpl',
            internalType: 'contract SystemConfig',
            type: 'address',
          },
        ],
      },
      {
        name: 'systemConfigConfig',
        internalType: 'struct SystemDictator.SystemConfigConfig',
        type: 'tuple',
        components: [
          { name: 'owner', internalType: 'address', type: 'address' },
          { name: 'overhead', internalType: 'uint256', type: 'uint256' },
          { name: 'scalar', internalType: 'uint256', type: 'uint256' },
          { name: 'batcherHash', internalType: 'bytes32', type: 'bytes32' },
          { name: 'gasLimit', internalType: 'uint64', type: 'uint64' },
          {
            name: 'unsafeBlockSigner',
            internalType: 'address',
            type: 'address',
          },
        ],
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'currentStep',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'dynamicConfigSet',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'exit1',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'finalize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'finalized',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_config',
        internalType: 'struct SystemDictator.DeployConfig',
        type: 'tuple',
        components: [
          {
            name: 'globalConfig',
            internalType: 'struct SystemDictator.GlobalConfig',
            type: 'tuple',
            components: [
              {
                name: 'addressManager',
                internalType: 'contract AddressManager',
                type: 'address',
              },
              {
                name: 'proxyAdmin',
                internalType: 'contract ProxyAdmin',
                type: 'address',
              },
              { name: 'controller', internalType: 'address', type: 'address' },
              { name: 'finalOwner', internalType: 'address', type: 'address' },
            ],
          },
          {
            name: 'proxyAddressConfig',
            internalType: 'struct SystemDictator.ProxyAddressConfig',
            type: 'tuple',
            components: [
              {
                name: 'l2OutputOracleProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'optimismPortalProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'l1CrossDomainMessengerProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'l1StandardBridgeProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'optimismMintableERC20FactoryProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'l1ERC721BridgeProxy',
                internalType: 'address',
                type: 'address',
              },
              {
                name: 'systemConfigProxy',
                internalType: 'address',
                type: 'address',
              },
            ],
          },
          {
            name: 'implementationAddressConfig',
            internalType: 'struct SystemDictator.ImplementationAddressConfig',
            type: 'tuple',
            components: [
              {
                name: 'l2OutputOracleImpl',
                internalType: 'contract L2OutputOracle',
                type: 'address',
              },
              {
                name: 'optimismPortalImpl',
                internalType: 'contract OptimismPortal',
                type: 'address',
              },
              {
                name: 'l1CrossDomainMessengerImpl',
                internalType: 'contract L1CrossDomainMessenger',
                type: 'address',
              },
              {
                name: 'l1StandardBridgeImpl',
                internalType: 'contract L1StandardBridge',
                type: 'address',
              },
              {
                name: 'optimismMintableERC20FactoryImpl',
                internalType: 'contract OptimismMintableERC20Factory',
                type: 'address',
              },
              {
                name: 'l1ERC721BridgeImpl',
                internalType: 'contract L1ERC721Bridge',
                type: 'address',
              },
              {
                name: 'portalSenderImpl',
                internalType: 'contract PortalSender',
                type: 'address',
              },
              {
                name: 'systemConfigImpl',
                internalType: 'contract SystemConfig',
                type: 'address',
              },
            ],
          },
          {
            name: 'systemConfigConfig',
            internalType: 'struct SystemDictator.SystemConfigConfig',
            type: 'tuple',
            components: [
              { name: 'owner', internalType: 'address', type: 'address' },
              { name: 'overhead', internalType: 'uint256', type: 'uint256' },
              { name: 'scalar', internalType: 'uint256', type: 'uint256' },
              { name: 'batcherHash', internalType: 'bytes32', type: 'bytes32' },
              { name: 'gasLimit', internalType: 'uint64', type: 'uint64' },
              {
                name: 'unsafeBlockSigner',
                internalType: 'address',
                type: 'address',
              },
            ],
          },
        ],
      },
    ],
    name: 'initialize',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'l2OutputOracleDynamicConfig',
    outputs: [
      {
        name: 'l2OutputOracleStartingBlockNumber',
        internalType: 'uint256',
        type: 'uint256',
      },
      {
        name: 'l2OutputOracleStartingTimestamp',
        internalType: 'uint256',
        type: 'uint256',
      },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'oldL1CrossDomainMessenger',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step1',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step2',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step3',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step4',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step5',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'step6',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      {
        name: '_l2OutputOracleDynamicConfig',
        internalType: 'struct SystemDictator.L2OutputOracleDynamicConfig',
        type: 'tuple',
        components: [
          {
            name: 'l2OutputOracleStartingBlockNumber',
            internalType: 'uint256',
            type: 'uint256',
          },
          {
            name: 'l2OutputOracleStartingTimestamp',
            internalType: 'uint256',
            type: 'uint256',
          },
        ],
      },
    ],
    name: 'updateL2OutputOracleDynamicConfig',
    outputs: [],
  },
] as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x1f0613A44c9a8ECE7B3A2e0CdBdF0F5B47A50971)
 */
export const systemDictatorGoerliAddress = {
  5: '0x1f0613A44c9a8ECE7B3A2e0CdBdF0F5B47A50971',
} as const

/**
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x1f0613A44c9a8ECE7B3A2e0CdBdF0F5B47A50971)
 */
export const systemDictatorGoerliConfig = {
  address: systemDictatorGoerliAddress,
  abi: systemDictatorGoerliABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// TeleportrWithdrawer
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x78A25524D90E3D0596558fb43789bD800a5c3007)
 */
export const teleportrWithdrawerABI = [
  {
    stateMutability: 'nonpayable',
    type: 'constructor',
    inputs: [{ name: '_owner', internalType: 'address', type: 'address' }],
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'user', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnerUpdated',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'ReceivedETH',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'WithdrewERC20',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      { name: 'id', internalType: 'uint256', type: 'uint256', indexed: false },
    ],
    name: 'WithdrewERC721',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'WithdrewETH',
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
      { name: '_gas', internalType: 'uint256', type: 'uint256' },
      { name: '_value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'CALL',
    outputs: [
      { name: '', internalType: 'bool', type: 'bool' },
      { name: '', internalType: 'bytes', type: 'bytes' },
    ],
  },
  {
    stateMutability: 'payable',
    type: 'function',
    inputs: [
      { name: '_target', internalType: 'address', type: 'address' },
      { name: '_data', internalType: 'bytes', type: 'bytes' },
      { name: '_gas', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'DELEGATECALL',
    outputs: [
      { name: '', internalType: 'bool', type: 'bool' },
      { name: '', internalType: 'bytes', type: 'bytes' },
    ],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'data',
    outputs: [{ name: '', internalType: 'bytes', type: 'bytes' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'recipient',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_data', internalType: 'bytes', type: 'bytes' }],
    name: 'setData',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'setOwner',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_recipient', internalType: 'address', type: 'address' }],
    name: 'setRecipient',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_teleportr', internalType: 'address', type: 'address' }],
    name: 'setTeleportr',
    outputs: [],
  },
  {
    stateMutability: 'view',
    type: 'function',
    inputs: [],
    name: 'teleportr',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC20', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC20', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
    ],
    name: 'withdrawERC20',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_asset', internalType: 'contract ERC721', type: 'address' },
      { name: '_to', internalType: 'address', type: 'address' },
      { name: '_id', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawERC721',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [
      { name: '_to', internalType: 'address payable', type: 'address' },
      { name: '_amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdrawETH',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [{ name: '_to', internalType: 'address payable', type: 'address' }],
    name: 'withdrawETH',
    outputs: [],
  },
  {
    stateMutability: 'nonpayable',
    type: 'function',
    inputs: [],
    name: 'withdrawFromTeleportr',
    outputs: [],
  },
  { stateMutability: 'payable', type: 'receive' },
] as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x78A25524D90E3D0596558fb43789bD800a5c3007)
 */
export const teleportrWithdrawerAddress = {
  1: '0x78A25524D90E3D0596558fb43789bD800a5c3007',
} as const

/**
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x78A25524D90E3D0596558fb43789bD800a5c3007)
 */
export const teleportrWithdrawerConfig = {
  address: teleportrWithdrawerAddress,
  abi: teleportrWithdrawerABI,
} as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// eslintIgnore
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Core
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link addressManagerABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdE1FCfB0851916CA5101820A69b13a4E276bd81F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xa6f73589243a6A7a9023b1Fa0651b1d89c177111)
 */
export function getAddressManager(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof addressManagerAddress
  }
) {
  return getContract({
    abi: addressManagerABI,
    address:
      addressManagerAddress[
        config.chainId as keyof typeof addressManagerAddress
      ],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link addressManagerABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdE1FCfB0851916CA5101820A69b13a4E276bd81F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xa6f73589243a6A7a9023b1Fa0651b1d89c177111)
 */
export function readAddressManager<
  TAbi extends readonly unknown[] = typeof addressManagerABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof addressManagerAddress
  }
) {
  return readContract({
    abi: addressManagerABI,
    address:
      addressManagerAddress[
        config.chainId as keyof typeof addressManagerAddress
      ],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link addressManagerABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdE1FCfB0851916CA5101820A69b13a4E276bd81F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xa6f73589243a6A7a9023b1Fa0651b1d89c177111)
 */
export function writeAddressManager<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof addressManagerAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof addressManagerABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof addressManagerAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof addressManagerABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof addressManagerAddress
      })
) {
  return writeContract({
    abi: addressManagerABI,
    address:
      addressManagerAddress[
        config.chainId as keyof typeof addressManagerAddress
      ],
    ...config,
  } as unknown as WriteContractArgs<typeof addressManagerABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link addressManagerABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdE1FCfB0851916CA5101820A69b13a4E276bd81F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xa6f73589243a6A7a9023b1Fa0651b1d89c177111)
 */
export function prepareWriteAddressManager<
  TAbi extends readonly unknown[] = typeof addressManagerABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof addressManagerAddress }
) {
  return prepareWriteContract({
    abi: addressManagerABI,
    address:
      addressManagerAddress[
        config.chainId as keyof typeof addressManagerAddress
      ],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link addressManagerABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdE1FCfB0851916CA5101820A69b13a4E276bd81F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xa6f73589243a6A7a9023b1Fa0651b1d89c177111)
 */
export function watchAddressManagerEvent<
  TAbi extends readonly unknown[] = typeof addressManagerABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof addressManagerAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: addressManagerABI,
      address:
        addressManagerAddress[
          config.chainId as keyof typeof addressManagerAddress
        ],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link assetReceiverABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 */
export function getAssetReceiver(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof assetReceiverAddress
  }
) {
  return getContract({
    abi: assetReceiverABI,
    address:
      assetReceiverAddress[config.chainId as keyof typeof assetReceiverAddress],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link assetReceiverABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 */
export function readAssetReceiver<
  TAbi extends readonly unknown[] = typeof assetReceiverABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof assetReceiverAddress
  }
) {
  return readContract({
    abi: assetReceiverABI,
    address:
      assetReceiverAddress[config.chainId as keyof typeof assetReceiverAddress],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link assetReceiverABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 */
export function writeAssetReceiver<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof assetReceiverAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof assetReceiverABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof assetReceiverAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof assetReceiverABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof assetReceiverAddress
      })
) {
  return writeContract({
    abi: assetReceiverABI,
    address:
      assetReceiverAddress[config.chainId as keyof typeof assetReceiverAddress],
    ...config,
  } as unknown as WriteContractArgs<typeof assetReceiverABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link assetReceiverABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 */
export function prepareWriteAssetReceiver<
  TAbi extends readonly unknown[] = typeof assetReceiverABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof assetReceiverAddress }
) {
  return prepareWriteContract({
    abi: assetReceiverABI,
    address:
      assetReceiverAddress[config.chainId as keyof typeof assetReceiverAddress],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link assetReceiverABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x15DdA60616Ffca20371ED1659dBB78E888f65556)
 */
export function watchAssetReceiverEvent<
  TAbi extends readonly unknown[] = typeof assetReceiverABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof assetReceiverAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: assetReceiverABI,
      address:
        assetReceiverAddress[
          config.chainId as keyof typeof assetReceiverAddress
        ],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link attestationStationABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 */
export function getAttestationStation(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof attestationStationAddress
  }
) {
  return getContract({
    abi: attestationStationABI,
    address:
      attestationStationAddress[
        config.chainId as keyof typeof attestationStationAddress
      ],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link attestationStationABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 */
export function readAttestationStation<
  TAbi extends readonly unknown[] = typeof attestationStationABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof attestationStationAddress
  }
) {
  return readContract({
    abi: attestationStationABI,
    address:
      attestationStationAddress[
        config.chainId as keyof typeof attestationStationAddress
      ],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link attestationStationABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 */
export function writeAttestationStation<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof attestationStationAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof attestationStationABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof attestationStationAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof attestationStationABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof attestationStationAddress
      })
) {
  return writeContract({
    abi: attestationStationABI,
    address:
      attestationStationAddress[
        config.chainId as keyof typeof attestationStationAddress
      ],
    ...config,
  } as unknown as WriteContractArgs<typeof attestationStationABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link attestationStationABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 */
export function prepareWriteAttestationStation<
  TAbi extends readonly unknown[] = typeof attestationStationABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof attestationStationAddress }
) {
  return prepareWriteContract({
    abi: attestationStationABI,
    address:
      attestationStationAddress[
        config.chainId as keyof typeof attestationStationAddress
      ],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link attestationStationABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)
 */
export function watchAttestationStationEvent<
  TAbi extends readonly unknown[] = typeof attestationStationABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof attestationStationAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: attestationStationABI,
      address:
        attestationStationAddress[
          config.chainId as keyof typeof attestationStationAddress
        ],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link baseFeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000019)
 */
export function getBaseFeeVault(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof baseFeeVaultAddress
  }
) {
  return getContract({
    abi: baseFeeVaultABI,
    address: baseFeeVaultAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link baseFeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000019)
 */
export function readBaseFeeVault<
  TAbi extends readonly unknown[] = typeof baseFeeVaultABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof baseFeeVaultAddress
  }
) {
  return readContract({
    abi: baseFeeVaultABI,
    address: baseFeeVaultAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link baseFeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000019)
 */
export function writeBaseFeeVault<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof baseFeeVaultAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof baseFeeVaultABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof baseFeeVaultAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof baseFeeVaultABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof baseFeeVaultAddress
      })
) {
  return writeContract({
    abi: baseFeeVaultABI,
    address: baseFeeVaultAddress[420],
    ...config,
  } as unknown as WriteContractArgs<typeof baseFeeVaultABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link baseFeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000019)
 */
export function prepareWriteBaseFeeVault<
  TAbi extends readonly unknown[] = typeof baseFeeVaultABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof baseFeeVaultAddress }
) {
  return prepareWriteContract({
    abi: baseFeeVaultABI,
    address: baseFeeVaultAddress[420],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link baseFeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000019)
 */
export function watchBaseFeeVaultEvent<
  TAbi extends readonly unknown[] = typeof baseFeeVaultABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof baseFeeVaultAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: baseFeeVaultABI,
      address: baseFeeVaultAddress[420],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link checkBalanceHighABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5d7103853f12109A7d27F118e54BbC654ad847E9)
 */
export function getCheckBalanceHigh(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof checkBalanceHighAddress
  }
) {
  return getContract({
    abi: checkBalanceHighABI,
    address:
      checkBalanceHighAddress[
        config.chainId as keyof typeof checkBalanceHighAddress
      ],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link checkBalanceHighABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5d7103853f12109A7d27F118e54BbC654ad847E9)
 */
export function readCheckBalanceHigh<
  TAbi extends readonly unknown[] = typeof checkBalanceHighABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof checkBalanceHighAddress
  }
) {
  return readContract({
    abi: checkBalanceHighABI,
    address:
      checkBalanceHighAddress[
        config.chainId as keyof typeof checkBalanceHighAddress
      ],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link checkBalanceHighABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x7eC64a8a591bFf829ff6C8be76074D540ACb813F)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5d7103853f12109A7d27F118e54BbC654ad847E9)
 */
export function watchCheckBalanceHighEvent<
  TAbi extends readonly unknown[] = typeof checkBalanceHighABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof checkBalanceHighAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: checkBalanceHighABI,
      address:
        checkBalanceHighAddress[
          config.chainId as keyof typeof checkBalanceHighAddress
        ],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link checkBalanceLowABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x7Ce13D154FAEE5C8B3E6b19d4Add16f21d884474)
 */
export function getCheckBalanceLow(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof checkBalanceLowAddress
  }
) {
  return getContract({
    abi: checkBalanceLowABI,
    address:
      checkBalanceLowAddress[
        config.chainId as keyof typeof checkBalanceLowAddress
      ],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link checkBalanceLowABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x7Ce13D154FAEE5C8B3E6b19d4Add16f21d884474)
 */
export function readCheckBalanceLow<
  TAbi extends readonly unknown[] = typeof checkBalanceLowABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof checkBalanceLowAddress
  }
) {
  return readContract({
    abi: checkBalanceLowABI,
    address:
      checkBalanceLowAddress[
        config.chainId as keyof typeof checkBalanceLowAddress
      ],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link checkBalanceLowABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x381a4eFC2A2C914eA1889722bB4B44Fa6BD5b640)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x7Ce13D154FAEE5C8B3E6b19d4Add16f21d884474)
 */
export function watchCheckBalanceLowEvent<
  TAbi extends readonly unknown[] = typeof checkBalanceLowABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof checkBalanceLowAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: checkBalanceLowABI,
      address:
        checkBalanceLowAddress[
          config.chainId as keyof typeof checkBalanceLowAddress
        ],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link checkGelatoLowABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xF9c8a4Cb4021f57F9f6d69799cA9BefF64524862)
 */
export function getCheckGelatoLow(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof checkGelatoLowAddress
  }
) {
  return getContract({
    abi: checkGelatoLowABI,
    address:
      checkGelatoLowAddress[
        config.chainId as keyof typeof checkGelatoLowAddress
      ],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link checkGelatoLowABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xF9c8a4Cb4021f57F9f6d69799cA9BefF64524862)
 */
export function readCheckGelatoLow<
  TAbi extends readonly unknown[] = typeof checkGelatoLowABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof checkGelatoLowAddress
  }
) {
  return readContract({
    abi: checkGelatoLowABI,
    address:
      checkGelatoLowAddress[
        config.chainId as keyof typeof checkGelatoLowAddress
      ],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link checkGelatoLowABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4f7CFc43f6D262a085F3b946cAC69E7a8E39BBAa)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0xF9c8a4Cb4021f57F9f6d69799cA9BefF64524862)
 */
export function watchCheckGelatoLowEvent<
  TAbi extends readonly unknown[] = typeof checkGelatoLowABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof checkGelatoLowAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: checkGelatoLowABI,
      address:
        checkGelatoLowAddress[
          config.chainId as keyof typeof checkGelatoLowAddress
        ],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link checkTrueABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5c741a38cb11424711231777D71689C458eE835D)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5c741a38cb11424711231777D71689C458eE835D)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x47443D0C184e022F19BD1578F5bca6B8a9F58E32)
 */
export function getCheckTrue(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof checkTrueAddress
  }
) {
  return getContract({
    abi: checkTrueABI,
    address: checkTrueAddress[config.chainId as keyof typeof checkTrueAddress],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link checkTrueABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5c741a38cb11424711231777D71689C458eE835D)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5c741a38cb11424711231777D71689C458eE835D)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x47443D0C184e022F19BD1578F5bca6B8a9F58E32)
 */
export function readCheckTrue<
  TAbi extends readonly unknown[] = typeof checkTrueABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof checkTrueAddress
  }
) {
  return readContract({
    abi: checkTrueABI,
    address: checkTrueAddress[config.chainId as keyof typeof checkTrueAddress],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link drippieABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export function getDrippie(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof drippieAddress
  }
) {
  return getContract({ abi: drippieABI, address: drippieAddress[1], ...config })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link drippieABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export function readDrippie<
  TAbi extends readonly unknown[] = typeof drippieABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof drippieAddress
  }
) {
  return readContract({
    abi: drippieABI,
    address: drippieAddress[1],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link drippieABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export function writeDrippie<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof drippieAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof drippieABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof drippieAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof drippieABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof drippieAddress
      })
) {
  return writeContract({
    abi: drippieABI,
    address: drippieAddress[1],
    ...config,
  } as unknown as WriteContractArgs<typeof drippieABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link drippieABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export function prepareWriteDrippie<
  TAbi extends readonly unknown[] = typeof drippieABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof drippieAddress }
) {
  return prepareWriteContract({
    abi: drippieABI,
    address: drippieAddress[1],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link drippieABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export function watchDrippieEvent<
  TAbi extends readonly unknown[] = typeof drippieABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof drippieAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: drippieABI,
      address: drippieAddress[1],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link drippieGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export function getDrippieGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof drippieGoerliAddress
  }
) {
  return getContract({
    abi: drippieGoerliABI,
    address: drippieGoerliAddress[5],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link drippieGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export function readDrippieGoerli<
  TAbi extends readonly unknown[] = typeof drippieGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof drippieGoerliAddress
  }
) {
  return readContract({
    abi: drippieGoerliABI,
    address: drippieGoerliAddress[5],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link drippieGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export function writeDrippieGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof drippieGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof drippieGoerliABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof drippieGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof drippieGoerliABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof drippieGoerliAddress
      })
) {
  return writeContract({
    abi: drippieGoerliABI,
    address: drippieGoerliAddress[5],
    ...config,
  } as unknown as WriteContractArgs<typeof drippieGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link drippieGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export function prepareWriteDrippieGoerli<
  TAbi extends readonly unknown[] = typeof drippieGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof drippieGoerliAddress }
) {
  return prepareWriteContract({
    abi: drippieGoerliABI,
    address: drippieGoerliAddress[5],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link drippieGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x44b3A2a040057eBafC601A78647e805fd58B1f50)
 */
export function watchDrippieGoerliEvent<
  TAbi extends readonly unknown[] = typeof drippieGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof drippieGoerliAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: drippieGoerliABI,
      address: drippieGoerliAddress[5],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link drippieOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x8D8d533C16D23847EB04EEB0925be8900Dd3af86)
 */
export function getDrippieOptimismGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof drippieOptimismGoerliAddress
  }
) {
  return getContract({
    abi: drippieOptimismGoerliABI,
    address: drippieOptimismGoerliAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link drippieOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x8D8d533C16D23847EB04EEB0925be8900Dd3af86)
 */
export function readDrippieOptimismGoerli<
  TAbi extends readonly unknown[] = typeof drippieOptimismGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof drippieOptimismGoerliAddress
  }
) {
  return readContract({
    abi: drippieOptimismGoerliABI,
    address: drippieOptimismGoerliAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link drippieOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x8D8d533C16D23847EB04EEB0925be8900Dd3af86)
 */
export function writeDrippieOptimismGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof drippieOptimismGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof drippieOptimismGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof drippieOptimismGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof drippieOptimismGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof drippieOptimismGoerliAddress
      })
) {
  return writeContract({
    abi: drippieOptimismGoerliABI,
    address: drippieOptimismGoerliAddress[420],
    ...config,
  } as unknown as WriteContractArgs<typeof drippieOptimismGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link drippieOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x8D8d533C16D23847EB04EEB0925be8900Dd3af86)
 */
export function prepareWriteDrippieOptimismGoerli<
  TAbi extends readonly unknown[] = typeof drippieOptimismGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof drippieOptimismGoerliAddress }
) {
  return prepareWriteContract({
    abi: drippieOptimismGoerliABI,
    address: drippieOptimismGoerliAddress[420],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link drippieOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x8D8d533C16D23847EB04EEB0925be8900Dd3af86)
 */
export function watchDrippieOptimismGoerliEvent<
  TAbi extends readonly unknown[] = typeof drippieOptimismGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof drippieOptimismGoerliAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: drippieOptimismGoerliABI,
      address: drippieOptimismGoerliAddress[420],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link easABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4E0275Ea5a89e7a3c1B58411379D1a0eDdc5b088)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5A633F1cc84B03F7588486CF2F386c102061E6e1)
 */
export function getEas(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof easAddress
  }
) {
  return getContract({
    abi: easABI,
    address: easAddress[config.chainId as keyof typeof easAddress],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link easABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4E0275Ea5a89e7a3c1B58411379D1a0eDdc5b088)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5A633F1cc84B03F7588486CF2F386c102061E6e1)
 */
export function readEas<
  TAbi extends readonly unknown[] = typeof easABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof easAddress
  }
) {
  return readContract({
    abi: easABI,
    address: easAddress[config.chainId as keyof typeof easAddress],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link easABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4E0275Ea5a89e7a3c1B58411379D1a0eDdc5b088)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5A633F1cc84B03F7588486CF2F386c102061E6e1)
 */
export function writeEas<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof easAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof easABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared' ? TChainId : keyof typeof easAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof easABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared' ? TChainId : keyof typeof easAddress
      })
) {
  return writeContract({
    abi: easABI,
    address: easAddress[config.chainId as keyof typeof easAddress],
    ...config,
  } as unknown as WriteContractArgs<typeof easABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link easABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4E0275Ea5a89e7a3c1B58411379D1a0eDdc5b088)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5A633F1cc84B03F7588486CF2F386c102061E6e1)
 */
export function prepareWriteEas<
  TAbi extends readonly unknown[] = typeof easABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof easAddress }
) {
  return prepareWriteContract({
    abi: easABI,
    address: easAddress[config.chainId as keyof typeof easAddress],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link easABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4E0275Ea5a89e7a3c1B58411379D1a0eDdc5b088)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x5A633F1cc84B03F7588486CF2F386c102061E6e1)
 */
export function watchEasEvent<
  TAbi extends readonly unknown[] = typeof easABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof easAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: easABI,
      address: easAddress[config.chainId as keyof typeof easAddress],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link gasPriceOracleABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000000F)
 */
export function getGasPriceOracle(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof gasPriceOracleAddress
  }
) {
  return getContract({
    abi: gasPriceOracleABI,
    address: gasPriceOracleAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link gasPriceOracleABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000000F)
 */
export function readGasPriceOracle<
  TAbi extends readonly unknown[] = typeof gasPriceOracleABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof gasPriceOracleAddress
  }
) {
  return readContract({
    abi: gasPriceOracleABI,
    address: gasPriceOracleAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l1BlockABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000015)
 */
export function getL1Block(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l1BlockAddress
  }
) {
  return getContract({
    abi: l1BlockABI,
    address: l1BlockAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l1BlockABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000015)
 */
export function readL1Block<
  TAbi extends readonly unknown[] = typeof l1BlockABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l1BlockAddress
  }
) {
  return readContract({
    abi: l1BlockABI,
    address: l1BlockAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l1BlockABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000015)
 */
export function writeL1Block<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l1BlockAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof l1BlockABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1BlockAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof l1BlockABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1BlockAddress
      })
) {
  return writeContract({
    abi: l1BlockABI,
    address: l1BlockAddress[420],
    ...config,
  } as unknown as WriteContractArgs<typeof l1BlockABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l1BlockABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000015)
 */
export function prepareWriteL1Block<
  TAbi extends readonly unknown[] = typeof l1BlockABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1BlockAddress }
) {
  return prepareWriteContract({
    abi: l1BlockABI,
    address: l1BlockAddress[420],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l1CrossDomainMessengerABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1)
 */
export function getL1CrossDomainMessenger(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l1CrossDomainMessengerAddress
  }
) {
  return getContract({
    abi: l1CrossDomainMessengerABI,
    address: l1CrossDomainMessengerAddress[1],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l1CrossDomainMessengerABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1)
 */
export function readL1CrossDomainMessenger<
  TAbi extends readonly unknown[] = typeof l1CrossDomainMessengerABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l1CrossDomainMessengerAddress
  }
) {
  return readContract({
    abi: l1CrossDomainMessengerABI,
    address: l1CrossDomainMessengerAddress[1],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l1CrossDomainMessengerABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1)
 */
export function writeL1CrossDomainMessenger<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l1CrossDomainMessengerAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof l1CrossDomainMessengerABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1CrossDomainMessengerAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof l1CrossDomainMessengerABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1CrossDomainMessengerAddress
      })
) {
  return writeContract({
    abi: l1CrossDomainMessengerABI,
    address: l1CrossDomainMessengerAddress[1],
    ...config,
  } as unknown as WriteContractArgs<typeof l1CrossDomainMessengerABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l1CrossDomainMessengerABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1)
 */
export function prepareWriteL1CrossDomainMessenger<
  TAbi extends readonly unknown[] = typeof l1CrossDomainMessengerABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1CrossDomainMessengerAddress }
) {
  return prepareWriteContract({
    abi: l1CrossDomainMessengerABI,
    address: l1CrossDomainMessengerAddress[1],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l1CrossDomainMessengerABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1)
 */
export function watchL1CrossDomainMessengerEvent<
  TAbi extends readonly unknown[] = typeof l1CrossDomainMessengerABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1CrossDomainMessengerAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l1CrossDomainMessengerABI,
      address: l1CrossDomainMessengerAddress[1],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l1CrossDomainMessengerGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5086d1eEF304eb5284A0f6720f79403b4e9bE294)
 */
export function getL1CrossDomainMessengerGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l1CrossDomainMessengerGoerliAddress
  }
) {
  return getContract({
    abi: l1CrossDomainMessengerGoerliABI,
    address: l1CrossDomainMessengerGoerliAddress[5],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l1CrossDomainMessengerGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5086d1eEF304eb5284A0f6720f79403b4e9bE294)
 */
export function readL1CrossDomainMessengerGoerli<
  TAbi extends readonly unknown[] = typeof l1CrossDomainMessengerGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l1CrossDomainMessengerGoerliAddress
  }
) {
  return readContract({
    abi: l1CrossDomainMessengerGoerliABI,
    address: l1CrossDomainMessengerGoerliAddress[5],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l1CrossDomainMessengerGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5086d1eEF304eb5284A0f6720f79403b4e9bE294)
 */
export function writeL1CrossDomainMessengerGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l1CrossDomainMessengerGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof l1CrossDomainMessengerGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1CrossDomainMessengerGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof l1CrossDomainMessengerGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1CrossDomainMessengerGoerliAddress
      })
) {
  return writeContract({
    abi: l1CrossDomainMessengerGoerliABI,
    address: l1CrossDomainMessengerGoerliAddress[5],
    ...config,
  } as unknown as WriteContractArgs<typeof l1CrossDomainMessengerGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l1CrossDomainMessengerGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5086d1eEF304eb5284A0f6720f79403b4e9bE294)
 */
export function prepareWriteL1CrossDomainMessengerGoerli<
  TAbi extends readonly unknown[] = typeof l1CrossDomainMessengerGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1CrossDomainMessengerGoerliAddress }
) {
  return prepareWriteContract({
    abi: l1CrossDomainMessengerGoerliABI,
    address: l1CrossDomainMessengerGoerliAddress[5],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l1CrossDomainMessengerGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5086d1eEF304eb5284A0f6720f79403b4e9bE294)
 */
export function watchL1CrossDomainMessengerGoerliEvent<
  TAbi extends readonly unknown[] = typeof l1CrossDomainMessengerGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1CrossDomainMessengerGoerliAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l1CrossDomainMessengerGoerliABI,
      address: l1CrossDomainMessengerGoerliAddress[5],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l1Erc721BridgeABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5a7749f83b81B301cAb5f48EB8516B986DAef23D)
 */
export function getL1Erc721Bridge(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l1Erc721BridgeAddress
  }
) {
  return getContract({
    abi: l1Erc721BridgeABI,
    address: l1Erc721BridgeAddress[1],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l1Erc721BridgeABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5a7749f83b81B301cAb5f48EB8516B986DAef23D)
 */
export function readL1Erc721Bridge<
  TAbi extends readonly unknown[] = typeof l1Erc721BridgeABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l1Erc721BridgeAddress
  }
) {
  return readContract({
    abi: l1Erc721BridgeABI,
    address: l1Erc721BridgeAddress[1],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l1Erc721BridgeABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5a7749f83b81B301cAb5f48EB8516B986DAef23D)
 */
export function writeL1Erc721Bridge<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l1Erc721BridgeAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof l1Erc721BridgeABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1Erc721BridgeAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof l1Erc721BridgeABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1Erc721BridgeAddress
      })
) {
  return writeContract({
    abi: l1Erc721BridgeABI,
    address: l1Erc721BridgeAddress[1],
    ...config,
  } as unknown as WriteContractArgs<typeof l1Erc721BridgeABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l1Erc721BridgeABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5a7749f83b81B301cAb5f48EB8516B986DAef23D)
 */
export function prepareWriteL1Erc721Bridge<
  TAbi extends readonly unknown[] = typeof l1Erc721BridgeABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1Erc721BridgeAddress }
) {
  return prepareWriteContract({
    abi: l1Erc721BridgeABI,
    address: l1Erc721BridgeAddress[1],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l1Erc721BridgeABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x5a7749f83b81B301cAb5f48EB8516B986DAef23D)
 */
export function watchL1Erc721BridgeEvent<
  TAbi extends readonly unknown[] = typeof l1Erc721BridgeABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1Erc721BridgeAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l1Erc721BridgeABI,
      address: l1Erc721BridgeAddress[1],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l1Erc721BridgeGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x8DD330DdE8D9898d43b4dc840Da27A07dF91b3c9)
 */
export function getL1Erc721BridgeGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l1Erc721BridgeGoerliAddress
  }
) {
  return getContract({
    abi: l1Erc721BridgeGoerliABI,
    address: l1Erc721BridgeGoerliAddress[5],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l1Erc721BridgeGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x8DD330DdE8D9898d43b4dc840Da27A07dF91b3c9)
 */
export function readL1Erc721BridgeGoerli<
  TAbi extends readonly unknown[] = typeof l1Erc721BridgeGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l1Erc721BridgeGoerliAddress
  }
) {
  return readContract({
    abi: l1Erc721BridgeGoerliABI,
    address: l1Erc721BridgeGoerliAddress[5],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l1Erc721BridgeGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x8DD330DdE8D9898d43b4dc840Da27A07dF91b3c9)
 */
export function writeL1Erc721BridgeGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l1Erc721BridgeGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof l1Erc721BridgeGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1Erc721BridgeGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof l1Erc721BridgeGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1Erc721BridgeGoerliAddress
      })
) {
  return writeContract({
    abi: l1Erc721BridgeGoerliABI,
    address: l1Erc721BridgeGoerliAddress[5],
    ...config,
  } as unknown as WriteContractArgs<typeof l1Erc721BridgeGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l1Erc721BridgeGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x8DD330DdE8D9898d43b4dc840Da27A07dF91b3c9)
 */
export function prepareWriteL1Erc721BridgeGoerli<
  TAbi extends readonly unknown[] = typeof l1Erc721BridgeGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1Erc721BridgeGoerliAddress }
) {
  return prepareWriteContract({
    abi: l1Erc721BridgeGoerliABI,
    address: l1Erc721BridgeGoerliAddress[5],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l1Erc721BridgeGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x8DD330DdE8D9898d43b4dc840Da27A07dF91b3c9)
 */
export function watchL1Erc721BridgeGoerliEvent<
  TAbi extends readonly unknown[] = typeof l1Erc721BridgeGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1Erc721BridgeGoerliAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l1Erc721BridgeGoerliABI,
      address: l1Erc721BridgeGoerliAddress[5],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l1FeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000001a)
 */
export function getL1FeeVault(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l1FeeVaultAddress
  }
) {
  return getContract({
    abi: l1FeeVaultABI,
    address: l1FeeVaultAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l1FeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000001a)
 */
export function readL1FeeVault<
  TAbi extends readonly unknown[] = typeof l1FeeVaultABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l1FeeVaultAddress
  }
) {
  return readContract({
    abi: l1FeeVaultABI,
    address: l1FeeVaultAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l1FeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000001a)
 */
export function writeL1FeeVault<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l1FeeVaultAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof l1FeeVaultABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1FeeVaultAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof l1FeeVaultABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1FeeVaultAddress
      })
) {
  return writeContract({
    abi: l1FeeVaultABI,
    address: l1FeeVaultAddress[420],
    ...config,
  } as unknown as WriteContractArgs<typeof l1FeeVaultABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l1FeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000001a)
 */
export function prepareWriteL1FeeVault<
  TAbi extends readonly unknown[] = typeof l1FeeVaultABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1FeeVaultAddress }
) {
  return prepareWriteContract({
    abi: l1FeeVaultABI,
    address: l1FeeVaultAddress[420],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l1FeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x420000000000000000000000000000000000001a)
 */
export function watchL1FeeVaultEvent<
  TAbi extends readonly unknown[] = typeof l1FeeVaultABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1FeeVaultAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l1FeeVaultABI,
      address: l1FeeVaultAddress[420],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l1StandardBridgeABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1)
 */
export function getL1StandardBridge(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l1StandardBridgeAddress
  }
) {
  return getContract({
    abi: l1StandardBridgeABI,
    address: l1StandardBridgeAddress[1],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l1StandardBridgeABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1)
 */
export function readL1StandardBridge<
  TAbi extends readonly unknown[] = typeof l1StandardBridgeABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l1StandardBridgeAddress
  }
) {
  return readContract({
    abi: l1StandardBridgeABI,
    address: l1StandardBridgeAddress[1],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l1StandardBridgeABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1)
 */
export function writeL1StandardBridge<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l1StandardBridgeAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof l1StandardBridgeABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1StandardBridgeAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof l1StandardBridgeABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1StandardBridgeAddress
      })
) {
  return writeContract({
    abi: l1StandardBridgeABI,
    address: l1StandardBridgeAddress[1],
    ...config,
  } as unknown as WriteContractArgs<typeof l1StandardBridgeABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l1StandardBridgeABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1)
 */
export function prepareWriteL1StandardBridge<
  TAbi extends readonly unknown[] = typeof l1StandardBridgeABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1StandardBridgeAddress }
) {
  return prepareWriteContract({
    abi: l1StandardBridgeABI,
    address: l1StandardBridgeAddress[1],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l1StandardBridgeABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1)
 */
export function watchL1StandardBridgeEvent<
  TAbi extends readonly unknown[] = typeof l1StandardBridgeABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1StandardBridgeAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l1StandardBridgeABI,
      address: l1StandardBridgeAddress[1],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l1StandardBridgeGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x636Af16bf2f682dD3109e60102b8E1A089FedAa8)
 */
export function getL1StandardBridgeGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l1StandardBridgeGoerliAddress
  }
) {
  return getContract({
    abi: l1StandardBridgeGoerliABI,
    address: l1StandardBridgeGoerliAddress[5],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l1StandardBridgeGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x636Af16bf2f682dD3109e60102b8E1A089FedAa8)
 */
export function readL1StandardBridgeGoerli<
  TAbi extends readonly unknown[] = typeof l1StandardBridgeGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l1StandardBridgeGoerliAddress
  }
) {
  return readContract({
    abi: l1StandardBridgeGoerliABI,
    address: l1StandardBridgeGoerliAddress[5],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l1StandardBridgeGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x636Af16bf2f682dD3109e60102b8E1A089FedAa8)
 */
export function writeL1StandardBridgeGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l1StandardBridgeGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof l1StandardBridgeGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1StandardBridgeGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof l1StandardBridgeGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l1StandardBridgeGoerliAddress
      })
) {
  return writeContract({
    abi: l1StandardBridgeGoerliABI,
    address: l1StandardBridgeGoerliAddress[5],
    ...config,
  } as unknown as WriteContractArgs<typeof l1StandardBridgeGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l1StandardBridgeGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x636Af16bf2f682dD3109e60102b8E1A089FedAa8)
 */
export function prepareWriteL1StandardBridgeGoerli<
  TAbi extends readonly unknown[] = typeof l1StandardBridgeGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1StandardBridgeGoerliAddress }
) {
  return prepareWriteContract({
    abi: l1StandardBridgeGoerliABI,
    address: l1StandardBridgeGoerliAddress[5],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l1StandardBridgeGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x636Af16bf2f682dD3109e60102b8E1A089FedAa8)
 */
export function watchL1StandardBridgeGoerliEvent<
  TAbi extends readonly unknown[] = typeof l1StandardBridgeGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l1StandardBridgeGoerliAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l1StandardBridgeGoerliABI,
      address: l1StandardBridgeGoerliAddress[5],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l2CrossDomainMessengerABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000007)
 */
export function getL2CrossDomainMessenger(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l2CrossDomainMessengerAddress
  }
) {
  return getContract({
    abi: l2CrossDomainMessengerABI,
    address: l2CrossDomainMessengerAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l2CrossDomainMessengerABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000007)
 */
export function readL2CrossDomainMessenger<
  TAbi extends readonly unknown[] = typeof l2CrossDomainMessengerABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l2CrossDomainMessengerAddress
  }
) {
  return readContract({
    abi: l2CrossDomainMessengerABI,
    address: l2CrossDomainMessengerAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l2CrossDomainMessengerABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000007)
 */
export function writeL2CrossDomainMessenger<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l2CrossDomainMessengerAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof l2CrossDomainMessengerABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2CrossDomainMessengerAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof l2CrossDomainMessengerABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2CrossDomainMessengerAddress
      })
) {
  return writeContract({
    abi: l2CrossDomainMessengerABI,
    address: l2CrossDomainMessengerAddress[420],
    ...config,
  } as unknown as WriteContractArgs<typeof l2CrossDomainMessengerABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l2CrossDomainMessengerABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000007)
 */
export function prepareWriteL2CrossDomainMessenger<
  TAbi extends readonly unknown[] = typeof l2CrossDomainMessengerABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2CrossDomainMessengerAddress }
) {
  return prepareWriteContract({
    abi: l2CrossDomainMessengerABI,
    address: l2CrossDomainMessengerAddress[420],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l2CrossDomainMessengerABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000007)
 */
export function watchL2CrossDomainMessengerEvent<
  TAbi extends readonly unknown[] = typeof l2CrossDomainMessengerABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2CrossDomainMessengerAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l2CrossDomainMessengerABI,
      address: l2CrossDomainMessengerAddress[420],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l2Erc721BridgeABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000014)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000014)
 */
export function getL2Erc721Bridge(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l2Erc721BridgeAddress
  }
) {
  return getContract({
    abi: l2Erc721BridgeABI,
    address:
      l2Erc721BridgeAddress[
        config.chainId as keyof typeof l2Erc721BridgeAddress
      ],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l2Erc721BridgeABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000014)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000014)
 */
export function readL2Erc721Bridge<
  TAbi extends readonly unknown[] = typeof l2Erc721BridgeABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l2Erc721BridgeAddress
  }
) {
  return readContract({
    abi: l2Erc721BridgeABI,
    address:
      l2Erc721BridgeAddress[
        config.chainId as keyof typeof l2Erc721BridgeAddress
      ],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l2Erc721BridgeABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000014)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000014)
 */
export function writeL2Erc721Bridge<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l2Erc721BridgeAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof l2Erc721BridgeABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2Erc721BridgeAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof l2Erc721BridgeABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2Erc721BridgeAddress
      })
) {
  return writeContract({
    abi: l2Erc721BridgeABI,
    address:
      l2Erc721BridgeAddress[
        config.chainId as keyof typeof l2Erc721BridgeAddress
      ],
    ...config,
  } as unknown as WriteContractArgs<typeof l2Erc721BridgeABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l2Erc721BridgeABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000014)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000014)
 */
export function prepareWriteL2Erc721Bridge<
  TAbi extends readonly unknown[] = typeof l2Erc721BridgeABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2Erc721BridgeAddress }
) {
  return prepareWriteContract({
    abi: l2Erc721BridgeABI,
    address:
      l2Erc721BridgeAddress[
        config.chainId as keyof typeof l2Erc721BridgeAddress
      ],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l2Erc721BridgeABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000014)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000014)
 */
export function watchL2Erc721BridgeEvent<
  TAbi extends readonly unknown[] = typeof l2Erc721BridgeABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2Erc721BridgeAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l2Erc721BridgeABI,
      address:
        l2Erc721BridgeAddress[
          config.chainId as keyof typeof l2Erc721BridgeAddress
        ],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l2OutputOracleABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdfe97868233d1aa22e815a266982f2cf17685a27)
 */
export function getL2OutputOracle(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l2OutputOracleAddress
  }
) {
  return getContract({
    abi: l2OutputOracleABI,
    address: l2OutputOracleAddress[1],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l2OutputOracleABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdfe97868233d1aa22e815a266982f2cf17685a27)
 */
export function readL2OutputOracle<
  TAbi extends readonly unknown[] = typeof l2OutputOracleABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l2OutputOracleAddress
  }
) {
  return readContract({
    abi: l2OutputOracleABI,
    address: l2OutputOracleAddress[1],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l2OutputOracleABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdfe97868233d1aa22e815a266982f2cf17685a27)
 */
export function writeL2OutputOracle<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l2OutputOracleAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof l2OutputOracleABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2OutputOracleAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof l2OutputOracleABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2OutputOracleAddress
      })
) {
  return writeContract({
    abi: l2OutputOracleABI,
    address: l2OutputOracleAddress[1],
    ...config,
  } as unknown as WriteContractArgs<typeof l2OutputOracleABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l2OutputOracleABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdfe97868233d1aa22e815a266982f2cf17685a27)
 */
export function prepareWriteL2OutputOracle<
  TAbi extends readonly unknown[] = typeof l2OutputOracleABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2OutputOracleAddress }
) {
  return prepareWriteContract({
    abi: l2OutputOracleABI,
    address: l2OutputOracleAddress[1],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l2OutputOracleABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xdfe97868233d1aa22e815a266982f2cf17685a27)
 */
export function watchL2OutputOracleEvent<
  TAbi extends readonly unknown[] = typeof l2OutputOracleABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2OutputOracleAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l2OutputOracleABI,
      address: l2OutputOracleAddress[1],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l2OutputOracleGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0)
 */
export function getL2OutputOracleGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l2OutputOracleGoerliAddress
  }
) {
  return getContract({
    abi: l2OutputOracleGoerliABI,
    address: l2OutputOracleGoerliAddress[5],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l2OutputOracleGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0)
 */
export function readL2OutputOracleGoerli<
  TAbi extends readonly unknown[] = typeof l2OutputOracleGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l2OutputOracleGoerliAddress
  }
) {
  return readContract({
    abi: l2OutputOracleGoerliABI,
    address: l2OutputOracleGoerliAddress[5],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l2OutputOracleGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0)
 */
export function writeL2OutputOracleGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l2OutputOracleGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof l2OutputOracleGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2OutputOracleGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof l2OutputOracleGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2OutputOracleGoerliAddress
      })
) {
  return writeContract({
    abi: l2OutputOracleGoerliABI,
    address: l2OutputOracleGoerliAddress[5],
    ...config,
  } as unknown as WriteContractArgs<typeof l2OutputOracleGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l2OutputOracleGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0)
 */
export function prepareWriteL2OutputOracleGoerli<
  TAbi extends readonly unknown[] = typeof l2OutputOracleGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2OutputOracleGoerliAddress }
) {
  return prepareWriteContract({
    abi: l2OutputOracleGoerliABI,
    address: l2OutputOracleGoerliAddress[5],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l2OutputOracleGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0)
 */
export function watchL2OutputOracleGoerliEvent<
  TAbi extends readonly unknown[] = typeof l2OutputOracleGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2OutputOracleGoerliAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l2OutputOracleGoerliABI,
      address: l2OutputOracleGoerliAddress[5],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l2StandardBridgeABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000010)
 */
export function getL2StandardBridge(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l2StandardBridgeAddress
  }
) {
  return getContract({
    abi: l2StandardBridgeABI,
    address: l2StandardBridgeAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l2StandardBridgeABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000010)
 */
export function readL2StandardBridge<
  TAbi extends readonly unknown[] = typeof l2StandardBridgeABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l2StandardBridgeAddress
  }
) {
  return readContract({
    abi: l2StandardBridgeABI,
    address: l2StandardBridgeAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l2StandardBridgeABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000010)
 */
export function writeL2StandardBridge<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l2StandardBridgeAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof l2StandardBridgeABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2StandardBridgeAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof l2StandardBridgeABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2StandardBridgeAddress
      })
) {
  return writeContract({
    abi: l2StandardBridgeABI,
    address: l2StandardBridgeAddress[420],
    ...config,
  } as unknown as WriteContractArgs<typeof l2StandardBridgeABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l2StandardBridgeABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000010)
 */
export function prepareWriteL2StandardBridge<
  TAbi extends readonly unknown[] = typeof l2StandardBridgeABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2StandardBridgeAddress }
) {
  return prepareWriteContract({
    abi: l2StandardBridgeABI,
    address: l2StandardBridgeAddress[420],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l2StandardBridgeABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000010)
 */
export function watchL2StandardBridgeEvent<
  TAbi extends readonly unknown[] = typeof l2StandardBridgeABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2StandardBridgeAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l2StandardBridgeABI,
      address: l2StandardBridgeAddress[420],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link l2ToL1MessagePasserABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000016)
 */
export function getL2ToL1MessagePasser(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof l2ToL1MessagePasserAddress
  }
) {
  return getContract({
    abi: l2ToL1MessagePasserABI,
    address: l2ToL1MessagePasserAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link l2ToL1MessagePasserABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000016)
 */
export function readL2ToL1MessagePasser<
  TAbi extends readonly unknown[] = typeof l2ToL1MessagePasserABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof l2ToL1MessagePasserAddress
  }
) {
  return readContract({
    abi: l2ToL1MessagePasserABI,
    address: l2ToL1MessagePasserAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link l2ToL1MessagePasserABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000016)
 */
export function writeL2ToL1MessagePasser<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof l2ToL1MessagePasserAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof l2ToL1MessagePasserABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2ToL1MessagePasserAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof l2ToL1MessagePasserABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof l2ToL1MessagePasserAddress
      })
) {
  return writeContract({
    abi: l2ToL1MessagePasserABI,
    address: l2ToL1MessagePasserAddress[420],
    ...config,
  } as unknown as WriteContractArgs<typeof l2ToL1MessagePasserABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link l2ToL1MessagePasserABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000016)
 */
export function prepareWriteL2ToL1MessagePasser<
  TAbi extends readonly unknown[] = typeof l2ToL1MessagePasserABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2ToL1MessagePasserAddress }
) {
  return prepareWriteContract({
    abi: l2ToL1MessagePasserABI,
    address: l2ToL1MessagePasserAddress[420],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link l2ToL1MessagePasserABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000016)
 */
export function watchL2ToL1MessagePasserEvent<
  TAbi extends readonly unknown[] = typeof l2ToL1MessagePasserABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof l2ToL1MessagePasserAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: l2ToL1MessagePasserABI,
      address: l2ToL1MessagePasserAddress[420],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link mintManagerABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x5C4e7Ba1E219E47948e6e3F55019A647bA501005)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x038a8825A3C3B0c08d52Cc76E5E361953Cf6Dc76)
 */
export function getMintManager(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof mintManagerAddress
  }
) {
  return getContract({
    abi: mintManagerABI,
    address:
      mintManagerAddress[config.chainId as keyof typeof mintManagerAddress],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link mintManagerABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x5C4e7Ba1E219E47948e6e3F55019A647bA501005)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x038a8825A3C3B0c08d52Cc76E5E361953Cf6Dc76)
 */
export function readMintManager<
  TAbi extends readonly unknown[] = typeof mintManagerABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof mintManagerAddress
  }
) {
  return readContract({
    abi: mintManagerABI,
    address:
      mintManagerAddress[config.chainId as keyof typeof mintManagerAddress],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link mintManagerABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x5C4e7Ba1E219E47948e6e3F55019A647bA501005)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x038a8825A3C3B0c08d52Cc76E5E361953Cf6Dc76)
 */
export function writeMintManager<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof mintManagerAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof mintManagerABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof mintManagerAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof mintManagerABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof mintManagerAddress
      })
) {
  return writeContract({
    abi: mintManagerABI,
    address:
      mintManagerAddress[config.chainId as keyof typeof mintManagerAddress],
    ...config,
  } as unknown as WriteContractArgs<typeof mintManagerABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link mintManagerABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x5C4e7Ba1E219E47948e6e3F55019A647bA501005)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x038a8825A3C3B0c08d52Cc76E5E361953Cf6Dc76)
 */
export function prepareWriteMintManager<
  TAbi extends readonly unknown[] = typeof mintManagerABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof mintManagerAddress }
) {
  return prepareWriteContract({
    abi: mintManagerABI,
    address:
      mintManagerAddress[config.chainId as keyof typeof mintManagerAddress],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link mintManagerABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x5C4e7Ba1E219E47948e6e3F55019A647bA501005)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x038a8825A3C3B0c08d52Cc76E5E361953Cf6Dc76)
 */
export function watchMintManagerEvent<
  TAbi extends readonly unknown[] = typeof mintManagerABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof mintManagerAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: mintManagerABI,
      address:
        mintManagerAddress[config.chainId as keyof typeof mintManagerAddress],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link optimismMintableErc20FactoryABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function getOptimismMintableErc20Factory(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismMintableErc20FactoryAddress
  }
) {
  return getContract({
    abi: optimismMintableErc20FactoryABI,
    address: optimismMintableErc20FactoryAddress[1],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link optimismMintableErc20FactoryABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function readOptimismMintableErc20Factory<
  TAbi extends readonly unknown[] = typeof optimismMintableErc20FactoryABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismMintableErc20FactoryAddress
  }
) {
  return readContract({
    abi: optimismMintableErc20FactoryABI,
    address: optimismMintableErc20FactoryAddress[1],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link optimismMintableErc20FactoryABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function writeOptimismMintableErc20Factory<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof optimismMintableErc20FactoryAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof optimismMintableErc20FactoryABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismMintableErc20FactoryAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof optimismMintableErc20FactoryABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismMintableErc20FactoryAddress
      })
) {
  return writeContract({
    abi: optimismMintableErc20FactoryABI,
    address: optimismMintableErc20FactoryAddress[1],
    ...config,
  } as unknown as WriteContractArgs<typeof optimismMintableErc20FactoryABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link optimismMintableErc20FactoryABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function prepareWriteOptimismMintableErc20Factory<
  TAbi extends readonly unknown[] = typeof optimismMintableErc20FactoryABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimismMintableErc20FactoryAddress }
) {
  return prepareWriteContract({
    abi: optimismMintableErc20FactoryABI,
    address: optimismMintableErc20FactoryAddress[1],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link optimismMintableErc20FactoryABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function watchOptimismMintableErc20FactoryEvent<
  TAbi extends readonly unknown[] = typeof optimismMintableErc20FactoryABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimismMintableErc20FactoryAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: optimismMintableErc20FactoryABI,
      address: optimismMintableErc20FactoryAddress[1],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link optimismMintableErc20FactoryGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function getOptimismMintableErc20FactoryGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismMintableErc20FactoryGoerliAddress
  }
) {
  return getContract({
    abi: optimismMintableErc20FactoryGoerliABI,
    address: optimismMintableErc20FactoryGoerliAddress[5],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link optimismMintableErc20FactoryGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function readOptimismMintableErc20FactoryGoerli<
  TAbi extends readonly unknown[] = typeof optimismMintableErc20FactoryGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismMintableErc20FactoryGoerliAddress
  }
) {
  return readContract({
    abi: optimismMintableErc20FactoryGoerliABI,
    address: optimismMintableErc20FactoryGoerliAddress[5],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link optimismMintableErc20FactoryGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function writeOptimismMintableErc20FactoryGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof optimismMintableErc20FactoryGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof optimismMintableErc20FactoryGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismMintableErc20FactoryGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof optimismMintableErc20FactoryGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismMintableErc20FactoryGoerliAddress
      })
) {
  return writeContract({
    abi: optimismMintableErc20FactoryGoerliABI,
    address: optimismMintableErc20FactoryGoerliAddress[5],
    ...config,
  } as unknown as WriteContractArgs<typeof optimismMintableErc20FactoryGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link optimismMintableErc20FactoryGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function prepareWriteOptimismMintableErc20FactoryGoerli<
  TAbi extends readonly unknown[] = typeof optimismMintableErc20FactoryGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimismMintableErc20FactoryGoerliAddress }
) {
  return prepareWriteContract({
    abi: optimismMintableErc20FactoryGoerliABI,
    address: optimismMintableErc20FactoryGoerliAddress[5],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link optimismMintableErc20FactoryGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function watchOptimismMintableErc20FactoryGoerliEvent<
  TAbi extends readonly unknown[] = typeof optimismMintableErc20FactoryGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimismMintableErc20FactoryGoerliAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: optimismMintableErc20FactoryGoerliABI,
      address: optimismMintableErc20FactoryGoerliAddress[5],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link optimismMintableErc20FactoryOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function getOptimismMintableErc20FactoryOptimismGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismMintableErc20FactoryOptimismGoerliAddress
  }
) {
  return getContract({
    abi: optimismMintableErc20FactoryOptimismGoerliABI,
    address: optimismMintableErc20FactoryOptimismGoerliAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link optimismMintableErc20FactoryOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function readOptimismMintableErc20FactoryOptimismGoerli<
  TAbi extends readonly unknown[] = typeof optimismMintableErc20FactoryOptimismGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismMintableErc20FactoryOptimismGoerliAddress
  }
) {
  return readContract({
    abi: optimismMintableErc20FactoryOptimismGoerliABI,
    address: optimismMintableErc20FactoryOptimismGoerliAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link optimismMintableErc20FactoryOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function writeOptimismMintableErc20FactoryOptimismGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof optimismMintableErc20FactoryOptimismGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof optimismMintableErc20FactoryOptimismGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismMintableErc20FactoryOptimismGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof optimismMintableErc20FactoryOptimismGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismMintableErc20FactoryOptimismGoerliAddress
      })
) {
  return writeContract({
    abi: optimismMintableErc20FactoryOptimismGoerliABI,
    address: optimismMintableErc20FactoryOptimismGoerliAddress[420],
    ...config,
  } as unknown as WriteContractArgs<typeof optimismMintableErc20FactoryOptimismGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link optimismMintableErc20FactoryOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function prepareWriteOptimismMintableErc20FactoryOptimismGoerli<
  TAbi extends readonly unknown[] = typeof optimismMintableErc20FactoryOptimismGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & {
    chainId?: keyof typeof optimismMintableErc20FactoryOptimismGoerliAddress
  }
) {
  return prepareWriteContract({
    abi: optimismMintableErc20FactoryOptimismGoerliABI,
    address: optimismMintableErc20FactoryOptimismGoerliAddress[420],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link optimismMintableErc20FactoryOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000012)
 */
export function watchOptimismMintableErc20FactoryOptimismGoerliEvent<
  TAbi extends readonly unknown[] = typeof optimismMintableErc20FactoryOptimismGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & {
    chainId?: keyof typeof optimismMintableErc20FactoryOptimismGoerliAddress
  },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: optimismMintableErc20FactoryOptimismGoerliABI,
      address: optimismMintableErc20FactoryOptimismGoerliAddress[420],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link optimismMintableErc721FactoryABI}__.
 *
 * [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000017)
 */
export function getOptimismMintableErc721Factory(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismMintableErc721FactoryAddress
  }
) {
  return getContract({
    abi: optimismMintableErc721FactoryABI,
    address: optimismMintableErc721FactoryAddress[10],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link optimismMintableErc721FactoryABI}__.
 *
 * [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000017)
 */
export function readOptimismMintableErc721Factory<
  TAbi extends readonly unknown[] = typeof optimismMintableErc721FactoryABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismMintableErc721FactoryAddress
  }
) {
  return readContract({
    abi: optimismMintableErc721FactoryABI,
    address: optimismMintableErc721FactoryAddress[10],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link optimismMintableErc721FactoryABI}__.
 *
 * [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000017)
 */
export function writeOptimismMintableErc721Factory<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof optimismMintableErc721FactoryAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof optimismMintableErc721FactoryABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismMintableErc721FactoryAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof optimismMintableErc721FactoryABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismMintableErc721FactoryAddress
      })
) {
  return writeContract({
    abi: optimismMintableErc721FactoryABI,
    address: optimismMintableErc721FactoryAddress[10],
    ...config,
  } as unknown as WriteContractArgs<typeof optimismMintableErc721FactoryABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link optimismMintableErc721FactoryABI}__.
 *
 * [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000017)
 */
export function prepareWriteOptimismMintableErc721Factory<
  TAbi extends readonly unknown[] = typeof optimismMintableErc721FactoryABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimismMintableErc721FactoryAddress }
) {
  return prepareWriteContract({
    abi: optimismMintableErc721FactoryABI,
    address: optimismMintableErc721FactoryAddress[10],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link optimismMintableErc721FactoryABI}__.
 *
 * [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x4200000000000000000000000000000000000017)
 */
export function watchOptimismMintableErc721FactoryEvent<
  TAbi extends readonly unknown[] = typeof optimismMintableErc721FactoryABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimismMintableErc721FactoryAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: optimismMintableErc721FactoryABI,
      address: optimismMintableErc721FactoryAddress[10],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link optimismMintableErc721FactoryOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000017)
 */
export function getOptimismMintableErc721FactoryOptimismGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismMintableErc721FactoryOptimismGoerliAddress
  }
) {
  return getContract({
    abi: optimismMintableErc721FactoryOptimismGoerliABI,
    address: optimismMintableErc721FactoryOptimismGoerliAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link optimismMintableErc721FactoryOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000017)
 */
export function readOptimismMintableErc721FactoryOptimismGoerli<
  TAbi extends readonly unknown[] = typeof optimismMintableErc721FactoryOptimismGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismMintableErc721FactoryOptimismGoerliAddress
  }
) {
  return readContract({
    abi: optimismMintableErc721FactoryOptimismGoerliABI,
    address: optimismMintableErc721FactoryOptimismGoerliAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link optimismMintableErc721FactoryOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000017)
 */
export function writeOptimismMintableErc721FactoryOptimismGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof optimismMintableErc721FactoryOptimismGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof optimismMintableErc721FactoryOptimismGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismMintableErc721FactoryOptimismGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof optimismMintableErc721FactoryOptimismGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismMintableErc721FactoryOptimismGoerliAddress
      })
) {
  return writeContract({
    abi: optimismMintableErc721FactoryOptimismGoerliABI,
    address: optimismMintableErc721FactoryOptimismGoerliAddress[420],
    ...config,
  } as unknown as WriteContractArgs<typeof optimismMintableErc721FactoryOptimismGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link optimismMintableErc721FactoryOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000017)
 */
export function prepareWriteOptimismMintableErc721FactoryOptimismGoerli<
  TAbi extends readonly unknown[] = typeof optimismMintableErc721FactoryOptimismGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & {
    chainId?: keyof typeof optimismMintableErc721FactoryOptimismGoerliAddress
  }
) {
  return prepareWriteContract({
    abi: optimismMintableErc721FactoryOptimismGoerliABI,
    address: optimismMintableErc721FactoryOptimismGoerliAddress[420],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link optimismMintableErc721FactoryOptimismGoerliABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000017)
 */
export function watchOptimismMintableErc721FactoryOptimismGoerliEvent<
  TAbi extends readonly unknown[] = typeof optimismMintableErc721FactoryOptimismGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & {
    chainId?: keyof typeof optimismMintableErc721FactoryOptimismGoerliAddress
  },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: optimismMintableErc721FactoryOptimismGoerliABI,
      address: optimismMintableErc721FactoryOptimismGoerliAddress[420],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link optimismPortalABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xbEb5Fc579115071764c7423A4f12eDde41f106Ed)
 */
export function getOptimismPortal(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismPortalAddress
  }
) {
  return getContract({
    abi: optimismPortalABI,
    address: optimismPortalAddress[1],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link optimismPortalABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xbEb5Fc579115071764c7423A4f12eDde41f106Ed)
 */
export function readOptimismPortal<
  TAbi extends readonly unknown[] = typeof optimismPortalABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismPortalAddress
  }
) {
  return readContract({
    abi: optimismPortalABI,
    address: optimismPortalAddress[1],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link optimismPortalABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xbEb5Fc579115071764c7423A4f12eDde41f106Ed)
 */
export function writeOptimismPortal<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof optimismPortalAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof optimismPortalABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismPortalAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof optimismPortalABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismPortalAddress
      })
) {
  return writeContract({
    abi: optimismPortalABI,
    address: optimismPortalAddress[1],
    ...config,
  } as unknown as WriteContractArgs<typeof optimismPortalABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link optimismPortalABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xbEb5Fc579115071764c7423A4f12eDde41f106Ed)
 */
export function prepareWriteOptimismPortal<
  TAbi extends readonly unknown[] = typeof optimismPortalABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimismPortalAddress }
) {
  return prepareWriteContract({
    abi: optimismPortalABI,
    address: optimismPortalAddress[1],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link optimismPortalABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xbEb5Fc579115071764c7423A4f12eDde41f106Ed)
 */
export function watchOptimismPortalEvent<
  TAbi extends readonly unknown[] = typeof optimismPortalABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimismPortalAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: optimismPortalABI,
      address: optimismPortalAddress[1],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link optimismPortalGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383)
 */
export function getOptimismPortalGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismPortalGoerliAddress
  }
) {
  return getContract({
    abi: optimismPortalGoerliABI,
    address: optimismPortalGoerliAddress[5],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link optimismPortalGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383)
 */
export function readOptimismPortalGoerli<
  TAbi extends readonly unknown[] = typeof optimismPortalGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof optimismPortalGoerliAddress
  }
) {
  return readContract({
    abi: optimismPortalGoerliABI,
    address: optimismPortalGoerliAddress[5],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link optimismPortalGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383)
 */
export function writeOptimismPortalGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof optimismPortalGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof optimismPortalGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismPortalGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof optimismPortalGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimismPortalGoerliAddress
      })
) {
  return writeContract({
    abi: optimismPortalGoerliABI,
    address: optimismPortalGoerliAddress[5],
    ...config,
  } as unknown as WriteContractArgs<typeof optimismPortalGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link optimismPortalGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383)
 */
export function prepareWriteOptimismPortalGoerli<
  TAbi extends readonly unknown[] = typeof optimismPortalGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimismPortalGoerliAddress }
) {
  return prepareWriteContract({
    abi: optimismPortalGoerliABI,
    address: optimismPortalGoerliAddress[5],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link optimismPortalGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383)
 */
export function watchOptimismPortalGoerliEvent<
  TAbi extends readonly unknown[] = typeof optimismPortalGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimismPortalGoerliAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: optimismPortalGoerliABI,
      address: optimismPortalGoerliAddress[5],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link optimistABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 */
export function getOptimist(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof optimistAddress
  }
) {
  return getContract({
    abi: optimistABI,
    address: optimistAddress[config.chainId as keyof typeof optimistAddress],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link optimistABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 */
export function readOptimist<
  TAbi extends readonly unknown[] = typeof optimistABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof optimistAddress
  }
) {
  return readContract({
    abi: optimistABI,
    address: optimistAddress[config.chainId as keyof typeof optimistAddress],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link optimistABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 */
export function writeOptimist<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof optimistAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof optimistABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimistAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof optimistABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimistAddress
      })
) {
  return writeContract({
    abi: optimistABI,
    address: optimistAddress[config.chainId as keyof typeof optimistAddress],
    ...config,
  } as unknown as WriteContractArgs<typeof optimistABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link optimistABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 */
export function prepareWriteOptimist<
  TAbi extends readonly unknown[] = typeof optimistABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimistAddress }
) {
  return prepareWriteContract({
    abi: optimistABI,
    address: optimistAddress[config.chainId as keyof typeof optimistAddress],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link optimistABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5)
 */
export function watchOptimistEvent<
  TAbi extends readonly unknown[] = typeof optimistABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimistAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: optimistABI,
      address: optimistAddress[config.chainId as keyof typeof optimistAddress],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link optimistAllowlistABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180)
 */
export function getOptimistAllowlist(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof optimistAllowlistAddress
  }
) {
  return getContract({
    abi: optimistAllowlistABI,
    address:
      optimistAllowlistAddress[
        config.chainId as keyof typeof optimistAllowlistAddress
      ],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link optimistAllowlistABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x482b1945D58f2E9Db0CEbe13c7fcFc6876b41180)
 */
export function readOptimistAllowlist<
  TAbi extends readonly unknown[] = typeof optimistAllowlistABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof optimistAllowlistAddress
  }
) {
  return readContract({
    abi: optimistAllowlistABI,
    address:
      optimistAllowlistAddress[
        config.chainId as keyof typeof optimistAllowlistAddress
      ],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link optimistInviterABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 */
export function getOptimistInviter(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof optimistInviterAddress
  }
) {
  return getContract({
    abi: optimistInviterABI,
    address:
      optimistInviterAddress[
        config.chainId as keyof typeof optimistInviterAddress
      ],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link optimistInviterABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 */
export function readOptimistInviter<
  TAbi extends readonly unknown[] = typeof optimistInviterABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof optimistInviterAddress
  }
) {
  return readContract({
    abi: optimistInviterABI,
    address:
      optimistInviterAddress[
        config.chainId as keyof typeof optimistInviterAddress
      ],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link optimistInviterABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 */
export function writeOptimistInviter<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof optimistInviterAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof optimistInviterABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimistInviterAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof optimistInviterABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof optimistInviterAddress
      })
) {
  return writeContract({
    abi: optimistInviterABI,
    address:
      optimistInviterAddress[
        config.chainId as keyof typeof optimistInviterAddress
      ],
    ...config,
  } as unknown as WriteContractArgs<typeof optimistInviterABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link optimistInviterABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 */
export function prepareWriteOptimistInviter<
  TAbi extends readonly unknown[] = typeof optimistInviterABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimistInviterAddress }
) {
  return prepareWriteContract({
    abi: optimistInviterABI,
    address:
      optimistInviterAddress[
        config.chainId as keyof typeof optimistInviterAddress
      ],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link optimistInviterABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x073031A1E1b8F5458Ed41Ce56331F5fd7e1de929)
 */
export function watchOptimistInviterEvent<
  TAbi extends readonly unknown[] = typeof optimistInviterABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof optimistInviterAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: optimistInviterABI,
      address:
        optimistInviterAddress[
          config.chainId as keyof typeof optimistInviterAddress
        ],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link portalSenderABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x0A893d9576b9cFD9EF78595963dc973238E78210)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xe7FACd39531ee3C313330E93B4d7a8B8A3c84Aa4)
 */
export function getPortalSender(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof portalSenderAddress
  }
) {
  return getContract({
    abi: portalSenderABI,
    address:
      portalSenderAddress[config.chainId as keyof typeof portalSenderAddress],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link portalSenderABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x0A893d9576b9cFD9EF78595963dc973238E78210)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xe7FACd39531ee3C313330E93B4d7a8B8A3c84Aa4)
 */
export function readPortalSender<
  TAbi extends readonly unknown[] = typeof portalSenderABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof portalSenderAddress
  }
) {
  return readContract({
    abi: portalSenderABI,
    address:
      portalSenderAddress[config.chainId as keyof typeof portalSenderAddress],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link portalSenderABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x0A893d9576b9cFD9EF78595963dc973238E78210)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xe7FACd39531ee3C313330E93B4d7a8B8A3c84Aa4)
 */
export function writePortalSender<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof portalSenderAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof portalSenderABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof portalSenderAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof portalSenderABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof portalSenderAddress
      })
) {
  return writeContract({
    abi: portalSenderABI,
    address:
      portalSenderAddress[config.chainId as keyof typeof portalSenderAddress],
    ...config,
  } as unknown as WriteContractArgs<typeof portalSenderABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link portalSenderABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x0A893d9576b9cFD9EF78595963dc973238E78210)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xe7FACd39531ee3C313330E93B4d7a8B8A3c84Aa4)
 */
export function prepareWritePortalSender<
  TAbi extends readonly unknown[] = typeof portalSenderABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof portalSenderAddress }
) {
  return prepareWriteContract({
    abi: portalSenderABI,
    address:
      portalSenderAddress[config.chainId as keyof typeof portalSenderAddress],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link protocolVersionsABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x0C24F5098774aA366827D667494e9F889f7cFc08)
 */
export function getProtocolVersions(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof protocolVersionsAddress
  }
) {
  return getContract({
    abi: protocolVersionsABI,
    address: protocolVersionsAddress[5],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link protocolVersionsABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x0C24F5098774aA366827D667494e9F889f7cFc08)
 */
export function readProtocolVersions<
  TAbi extends readonly unknown[] = typeof protocolVersionsABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof protocolVersionsAddress
  }
) {
  return readContract({
    abi: protocolVersionsABI,
    address: protocolVersionsAddress[5],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link protocolVersionsABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x0C24F5098774aA366827D667494e9F889f7cFc08)
 */
export function writeProtocolVersions<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof protocolVersionsAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof protocolVersionsABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof protocolVersionsAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof protocolVersionsABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof protocolVersionsAddress
      })
) {
  return writeContract({
    abi: protocolVersionsABI,
    address: protocolVersionsAddress[5],
    ...config,
  } as unknown as WriteContractArgs<typeof protocolVersionsABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link protocolVersionsABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x0C24F5098774aA366827D667494e9F889f7cFc08)
 */
export function prepareWriteProtocolVersions<
  TAbi extends readonly unknown[] = typeof protocolVersionsABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof protocolVersionsAddress }
) {
  return prepareWriteContract({
    abi: protocolVersionsABI,
    address: protocolVersionsAddress[5],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link protocolVersionsABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x0C24F5098774aA366827D667494e9F889f7cFc08)
 */
export function watchProtocolVersionsEvent<
  TAbi extends readonly unknown[] = typeof protocolVersionsABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof protocolVersionsAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: protocolVersionsABI,
      address: protocolVersionsAddress[5],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link proxyAdminABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000018)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000018)
 */
export function getProxyAdmin(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof proxyAdminAddress
  }
) {
  return getContract({
    abi: proxyAdminABI,
    address:
      proxyAdminAddress[config.chainId as keyof typeof proxyAdminAddress],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link proxyAdminABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000018)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000018)
 */
export function readProxyAdmin<
  TAbi extends readonly unknown[] = typeof proxyAdminABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof proxyAdminAddress
  }
) {
  return readContract({
    abi: proxyAdminABI,
    address:
      proxyAdminAddress[config.chainId as keyof typeof proxyAdminAddress],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link proxyAdminABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000018)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000018)
 */
export function writeProxyAdmin<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof proxyAdminAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof proxyAdminABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof proxyAdminAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof proxyAdminABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof proxyAdminAddress
      })
) {
  return writeContract({
    abi: proxyAdminABI,
    address:
      proxyAdminAddress[config.chainId as keyof typeof proxyAdminAddress],
    ...config,
  } as unknown as WriteContractArgs<typeof proxyAdminABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link proxyAdminABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000018)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000018)
 */
export function prepareWriteProxyAdmin<
  TAbi extends readonly unknown[] = typeof proxyAdminABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof proxyAdminAddress }
) {
  return prepareWriteContract({
    abi: proxyAdminABI,
    address:
      proxyAdminAddress[config.chainId as keyof typeof proxyAdminAddress],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link proxyAdminABI}__.
 *
 * - [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x4200000000000000000000000000000000000018)
 * - [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x4200000000000000000000000000000000000018)
 */
export function watchProxyAdminEvent<
  TAbi extends readonly unknown[] = typeof proxyAdminABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof proxyAdminAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: proxyAdminABI,
      address:
        proxyAdminAddress[config.chainId as keyof typeof proxyAdminAddress],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link schemaRegistryABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x6232208d66bAc2305b46b4Cb6BCB3857B298DF13)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2545fa928d5d278cA75Fd47306e4a89096ff6403)
 */
export function getSchemaRegistry(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof schemaRegistryAddress
  }
) {
  return getContract({
    abi: schemaRegistryABI,
    address:
      schemaRegistryAddress[
        config.chainId as keyof typeof schemaRegistryAddress
      ],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link schemaRegistryABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x6232208d66bAc2305b46b4Cb6BCB3857B298DF13)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2545fa928d5d278cA75Fd47306e4a89096ff6403)
 */
export function readSchemaRegistry<
  TAbi extends readonly unknown[] = typeof schemaRegistryABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof schemaRegistryAddress
  }
) {
  return readContract({
    abi: schemaRegistryABI,
    address:
      schemaRegistryAddress[
        config.chainId as keyof typeof schemaRegistryAddress
      ],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link schemaRegistryABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x6232208d66bAc2305b46b4Cb6BCB3857B298DF13)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2545fa928d5d278cA75Fd47306e4a89096ff6403)
 */
export function writeSchemaRegistry<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof schemaRegistryAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof schemaRegistryABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof schemaRegistryAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof schemaRegistryABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof schemaRegistryAddress
      })
) {
  return writeContract({
    abi: schemaRegistryABI,
    address:
      schemaRegistryAddress[
        config.chainId as keyof typeof schemaRegistryAddress
      ],
    ...config,
  } as unknown as WriteContractArgs<typeof schemaRegistryABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link schemaRegistryABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x6232208d66bAc2305b46b4Cb6BCB3857B298DF13)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2545fa928d5d278cA75Fd47306e4a89096ff6403)
 */
export function prepareWriteSchemaRegistry<
  TAbi extends readonly unknown[] = typeof schemaRegistryABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof schemaRegistryAddress }
) {
  return prepareWriteContract({
    abi: schemaRegistryABI,
    address:
      schemaRegistryAddress[
        config.chainId as keyof typeof schemaRegistryAddress
      ],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link schemaRegistryABI}__.
 *
 * - [__View Contract on Op Mainnet Optimism Explorer__](https://explorer.optimism.io/address/0x6232208d66bAc2305b46b4Cb6BCB3857B298DF13)
 * - [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x2545fa928d5d278cA75Fd47306e4a89096ff6403)
 */
export function watchSchemaRegistryEvent<
  TAbi extends readonly unknown[] = typeof schemaRegistryABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof schemaRegistryAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: schemaRegistryABI,
      address:
        schemaRegistryAddress[
          config.chainId as keyof typeof schemaRegistryAddress
        ],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link sequencerFeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000011)
 */
export function getSequencerFeeVault(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof sequencerFeeVaultAddress
  }
) {
  return getContract({
    abi: sequencerFeeVaultABI,
    address: sequencerFeeVaultAddress[420],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link sequencerFeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000011)
 */
export function readSequencerFeeVault<
  TAbi extends readonly unknown[] = typeof sequencerFeeVaultABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof sequencerFeeVaultAddress
  }
) {
  return readContract({
    abi: sequencerFeeVaultABI,
    address: sequencerFeeVaultAddress[420],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link sequencerFeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000011)
 */
export function writeSequencerFeeVault<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof sequencerFeeVaultAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof sequencerFeeVaultABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof sequencerFeeVaultAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof sequencerFeeVaultABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof sequencerFeeVaultAddress
      })
) {
  return writeContract({
    abi: sequencerFeeVaultABI,
    address: sequencerFeeVaultAddress[420],
    ...config,
  } as unknown as WriteContractArgs<typeof sequencerFeeVaultABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link sequencerFeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000011)
 */
export function prepareWriteSequencerFeeVault<
  TAbi extends readonly unknown[] = typeof sequencerFeeVaultABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof sequencerFeeVaultAddress }
) {
  return prepareWriteContract({
    abi: sequencerFeeVaultABI,
    address: sequencerFeeVaultAddress[420],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link sequencerFeeVaultABI}__.
 *
 * [__View Contract on Optimism Goerli Etherscan__](https://goerli-optimism.etherscan.io/address/0x4200000000000000000000000000000000000011)
 */
export function watchSequencerFeeVaultEvent<
  TAbi extends readonly unknown[] = typeof sequencerFeeVaultABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof sequencerFeeVaultAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: sequencerFeeVaultABI,
      address: sequencerFeeVaultAddress[420],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link systemConfigABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x229047fed2591dbec1eF1118d64F7aF3dB9EB290)
 */
export function getSystemConfig(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof systemConfigAddress
  }
) {
  return getContract({
    abi: systemConfigABI,
    address: systemConfigAddress[1],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link systemConfigABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x229047fed2591dbec1eF1118d64F7aF3dB9EB290)
 */
export function readSystemConfig<
  TAbi extends readonly unknown[] = typeof systemConfigABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof systemConfigAddress
  }
) {
  return readContract({
    abi: systemConfigABI,
    address: systemConfigAddress[1],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link systemConfigABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x229047fed2591dbec1eF1118d64F7aF3dB9EB290)
 */
export function writeSystemConfig<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof systemConfigAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof systemConfigABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof systemConfigAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof systemConfigABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof systemConfigAddress
      })
) {
  return writeContract({
    abi: systemConfigABI,
    address: systemConfigAddress[1],
    ...config,
  } as unknown as WriteContractArgs<typeof systemConfigABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link systemConfigABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x229047fed2591dbec1eF1118d64F7aF3dB9EB290)
 */
export function prepareWriteSystemConfig<
  TAbi extends readonly unknown[] = typeof systemConfigABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof systemConfigAddress }
) {
  return prepareWriteContract({
    abi: systemConfigABI,
    address: systemConfigAddress[1],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link systemConfigABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x229047fed2591dbec1eF1118d64F7aF3dB9EB290)
 */
export function watchSystemConfigEvent<
  TAbi extends readonly unknown[] = typeof systemConfigABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof systemConfigAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: systemConfigABI,
      address: systemConfigAddress[1],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link systemConfigGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60)
 */
export function getSystemConfigGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof systemConfigGoerliAddress
  }
) {
  return getContract({
    abi: systemConfigGoerliABI,
    address: systemConfigGoerliAddress[5],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link systemConfigGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60)
 */
export function readSystemConfigGoerli<
  TAbi extends readonly unknown[] = typeof systemConfigGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof systemConfigGoerliAddress
  }
) {
  return readContract({
    abi: systemConfigGoerliABI,
    address: systemConfigGoerliAddress[5],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link systemConfigGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60)
 */
export function writeSystemConfigGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof systemConfigGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof systemConfigGoerliABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof systemConfigGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof systemConfigGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof systemConfigGoerliAddress
      })
) {
  return writeContract({
    abi: systemConfigGoerliABI,
    address: systemConfigGoerliAddress[5],
    ...config,
  } as unknown as WriteContractArgs<typeof systemConfigGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link systemConfigGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60)
 */
export function prepareWriteSystemConfigGoerli<
  TAbi extends readonly unknown[] = typeof systemConfigGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof systemConfigGoerliAddress }
) {
  return prepareWriteContract({
    abi: systemConfigGoerliABI,
    address: systemConfigGoerliAddress[5],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link systemConfigGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60)
 */
export function watchSystemConfigGoerliEvent<
  TAbi extends readonly unknown[] = typeof systemConfigGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof systemConfigGoerliAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: systemConfigGoerliABI,
      address: systemConfigGoerliAddress[5],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link systemDictatorABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xB4453CEb33d2e67FA244A24acf2E50CEF31F53cB)
 */
export function getSystemDictator(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof systemDictatorAddress
  }
) {
  return getContract({
    abi: systemDictatorABI,
    address: systemDictatorAddress[1],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link systemDictatorABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xB4453CEb33d2e67FA244A24acf2E50CEF31F53cB)
 */
export function readSystemDictator<
  TAbi extends readonly unknown[] = typeof systemDictatorABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof systemDictatorAddress
  }
) {
  return readContract({
    abi: systemDictatorABI,
    address: systemDictatorAddress[1],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link systemDictatorABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xB4453CEb33d2e67FA244A24acf2E50CEF31F53cB)
 */
export function writeSystemDictator<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof systemDictatorAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof systemDictatorABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof systemDictatorAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<typeof systemDictatorABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof systemDictatorAddress
      })
) {
  return writeContract({
    abi: systemDictatorABI,
    address: systemDictatorAddress[1],
    ...config,
  } as unknown as WriteContractArgs<typeof systemDictatorABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link systemDictatorABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xB4453CEb33d2e67FA244A24acf2E50CEF31F53cB)
 */
export function prepareWriteSystemDictator<
  TAbi extends readonly unknown[] = typeof systemDictatorABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof systemDictatorAddress }
) {
  return prepareWriteContract({
    abi: systemDictatorABI,
    address: systemDictatorAddress[1],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link systemDictatorABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0xB4453CEb33d2e67FA244A24acf2E50CEF31F53cB)
 */
export function watchSystemDictatorEvent<
  TAbi extends readonly unknown[] = typeof systemDictatorABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof systemDictatorAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: systemDictatorABI,
      address: systemDictatorAddress[1],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link systemDictatorGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x1f0613A44c9a8ECE7B3A2e0CdBdF0F5B47A50971)
 */
export function getSystemDictatorGoerli(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof systemDictatorGoerliAddress
  }
) {
  return getContract({
    abi: systemDictatorGoerliABI,
    address: systemDictatorGoerliAddress[5],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link systemDictatorGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x1f0613A44c9a8ECE7B3A2e0CdBdF0F5B47A50971)
 */
export function readSystemDictatorGoerli<
  TAbi extends readonly unknown[] = typeof systemDictatorGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof systemDictatorGoerliAddress
  }
) {
  return readContract({
    abi: systemDictatorGoerliABI,
    address: systemDictatorGoerliAddress[5],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link systemDictatorGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x1f0613A44c9a8ECE7B3A2e0CdBdF0F5B47A50971)
 */
export function writeSystemDictatorGoerli<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof systemDictatorGoerliAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<
          typeof systemDictatorGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof systemDictatorGoerliAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof systemDictatorGoerliABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof systemDictatorGoerliAddress
      })
) {
  return writeContract({
    abi: systemDictatorGoerliABI,
    address: systemDictatorGoerliAddress[5],
    ...config,
  } as unknown as WriteContractArgs<typeof systemDictatorGoerliABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link systemDictatorGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x1f0613A44c9a8ECE7B3A2e0CdBdF0F5B47A50971)
 */
export function prepareWriteSystemDictatorGoerli<
  TAbi extends readonly unknown[] = typeof systemDictatorGoerliABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof systemDictatorGoerliAddress }
) {
  return prepareWriteContract({
    abi: systemDictatorGoerliABI,
    address: systemDictatorGoerliAddress[5],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link systemDictatorGoerliABI}__.
 *
 * [__View Contract on Goerli Etherscan__](https://goerli.etherscan.io/address/0x1f0613A44c9a8ECE7B3A2e0CdBdF0F5B47A50971)
 */
export function watchSystemDictatorGoerliEvent<
  TAbi extends readonly unknown[] = typeof systemDictatorGoerliABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof systemDictatorGoerliAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: systemDictatorGoerliABI,
      address: systemDictatorGoerliAddress[5],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}

/**
 * Wraps __{@link getContract}__ with `abi` set to __{@link teleportrWithdrawerABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x78A25524D90E3D0596558fb43789bD800a5c3007)
 */
export function getTeleportrWithdrawer(
  config: Omit<GetContractArgs, 'abi' | 'address'> & {
    chainId?: keyof typeof teleportrWithdrawerAddress
  }
) {
  return getContract({
    abi: teleportrWithdrawerABI,
    address: teleportrWithdrawerAddress[1],
    ...config,
  })
}

/**
 * Wraps __{@link readContract}__ with `abi` set to __{@link teleportrWithdrawerABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x78A25524D90E3D0596558fb43789bD800a5c3007)
 */
export function readTeleportrWithdrawer<
  TAbi extends readonly unknown[] = typeof teleportrWithdrawerABI,
  TFunctionName extends string = string
>(
  config: Omit<ReadContractConfig<TAbi, TFunctionName>, 'abi' | 'address'> & {
    chainId?: keyof typeof teleportrWithdrawerAddress
  }
) {
  return readContract({
    abi: teleportrWithdrawerABI,
    address: teleportrWithdrawerAddress[1],
    ...config,
  } as unknown as ReadContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link writeContract}__ with `abi` set to __{@link teleportrWithdrawerABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x78A25524D90E3D0596558fb43789bD800a5c3007)
 */
export function writeTeleportrWithdrawer<
  TFunctionName extends string,
  TMode extends WriteContractMode,
  TChainId extends number = keyof typeof teleportrWithdrawerAddress
>(
  config:
    | (Omit<
        WriteContractPreparedArgs<typeof teleportrWithdrawerABI, TFunctionName>,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof teleportrWithdrawerAddress
      })
    | (Omit<
        WriteContractUnpreparedArgs<
          typeof teleportrWithdrawerABI,
          TFunctionName
        >,
        'abi' | 'address'
      > & {
        mode: TMode
        chainId?: TMode extends 'prepared'
          ? TChainId
          : keyof typeof teleportrWithdrawerAddress
      })
) {
  return writeContract({
    abi: teleportrWithdrawerABI,
    address: teleportrWithdrawerAddress[1],
    ...config,
  } as unknown as WriteContractArgs<typeof teleportrWithdrawerABI, TFunctionName>)
}

/**
 * Wraps __{@link prepareWriteContract}__ with `abi` set to __{@link teleportrWithdrawerABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x78A25524D90E3D0596558fb43789bD800a5c3007)
 */
export function prepareWriteTeleportrWithdrawer<
  TAbi extends readonly unknown[] = typeof teleportrWithdrawerABI,
  TFunctionName extends string = string
>(
  config: Omit<
    PrepareWriteContractConfig<TAbi, TFunctionName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof teleportrWithdrawerAddress }
) {
  return prepareWriteContract({
    abi: teleportrWithdrawerABI,
    address: teleportrWithdrawerAddress[1],
    ...config,
  } as unknown as PrepareWriteContractConfig<TAbi, TFunctionName>)
}

/**
 * Wraps __{@link watchContractEvent}__ with `abi` set to __{@link teleportrWithdrawerABI}__.
 *
 * [__View Contract on Ethereum Etherscan__](https://etherscan.io/address/0x78A25524D90E3D0596558fb43789bD800a5c3007)
 */
export function watchTeleportrWithdrawerEvent<
  TAbi extends readonly unknown[] = typeof teleportrWithdrawerABI,
  TEventName extends string = string
>(
  config: Omit<
    WatchContractEventConfig<TAbi, TEventName>,
    'abi' | 'address'
  > & { chainId?: keyof typeof teleportrWithdrawerAddress },
  callback: WatchContractEventCallback<TAbi, TEventName>
) {
  return watchContractEvent(
    {
      abi: teleportrWithdrawerABI,
      address: teleportrWithdrawerAddress[1],
      ...config,
    } as WatchContractEventConfig<TAbi, TEventName>,
    callback
  )
}
