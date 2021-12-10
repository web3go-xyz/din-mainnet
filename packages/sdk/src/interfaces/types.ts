import {
  Provider,
  TransactionReceipt,
  TransactionResponse,
} from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { Contract, BigNumber, Overrides } from 'ethers'

/**
 * Represents Optimistic Ethereum contracts, assumed to be connected to their appropriate
 * providers and addresses.
 */
export interface OEContracts {
  /**
   * L1 contract references.
   */
  l1: {
    AddressManager: Contract
    L1CrossDomainMessenger: Contract
    L1StandardBridge: Contract
    StateCommitmentChain: Contract
    CanonicalTransactionChain: Contract
    BondManager: Contract
  }

  /**
   * L2 contract references.
   */
  l2: {
    L2CrossDomainMessenger: Contract
    L2StandardBridge: Contract
    OVM_L1BlockNumber: Contract
    OVM_L2ToL1MessagePasser: Contract
    OVM_DeployerWhitelist: Contract
    OVM_ETH: Contract
    OVM_GasPriceOracle: Contract
    OVM_SequencerFeeVault: Contract
    WETH: Contract
  }
}

/**
 * Enum describing the status of a message.
 */
export enum MessageStatus {
  /**
   * Message is an L1 to L2 message and has not been processed by the L2.
   */
  UNCONFIRMED_L1_TO_L2_MESSAGE,

  /**
   * Message is an L1 to L2 message and the transaction to execute the message failed.
   * When this status is returned, you will need to resend the L1 to L2 message, probably with a
   * higher gas limit.
   */
  FAILED_L1_TO_L2_MESSAGE,

  /**
   * Message is an L2 to L1 message and no state root has been published yet.
   */
  STATE_ROOT_NOT_PUBLISHED,

  /**
   * Message is an L2 to L1 message and awaiting the challenge period.
   */
  IN_CHALLENGE_PERIOD,

  /**
   * Message is ready to be relayed.
   */
  READY_FOR_RELAY,

  /**
   * Message has been relayed.
   */
  RELAYED,
}

/**
 * Enum describing the direction of a message.
 */
export enum MessageDirection {
  L1_TO_L2,
  L2_TO_L1,
}

/**
 * Partial message that needs to be signed and executed by a specific signer.
 */
export interface CrossChainMessageRequest {
  direction: MessageDirection
  target: string
  message: string
  l2GasLimit: NumberLike
}

/**
 * Describes a message that is sent between L1 and L2. Direction determines where the message was
 * sent from and where it's being sent to.
 */
export interface CrossChainMessage {
  direction: MessageDirection
  sender: string
  target: string
  message: string
  messageNonce: number
}

/**
 * Convenience type for when you don't care which direction the message is going in.
 */
export type DirectionlessCrossChainMessage = Omit<
  CrossChainMessage,
  'direction'
>

/**
 * Describes a token withdrawal or deposit, along with the underlying raw cross chain message
 * behind the deposit or withdrawal.
 */
export interface TokenBridgeMessage {
  direction: MessageDirection
  from: string
  to: string
  l1Token: string
  l2Token: string
  amount: BigNumber
  raw: CrossChainMessage
}

/**
 * Enum describing the status of a CrossDomainMessage message receipt.
 */
export enum MessageReceiptStatus {
  RELAYED_SUCCEEDED,
  RELAYED_FAILED,
}

/**
 * CrossDomainMessage receipt.
 */
export interface MessageReceipt {
  messageHash: string
  receiptStatus: MessageReceiptStatus
  transactionReceipt: TransactionReceipt
}

/**
 * Header for a state root batch.
 */
export interface StateRootBatchHeader {
  batchIndex: BigNumber
  batchRoot: string
  batchSize: BigNumber
  prevTotalElements: BigNumber
  extraData: string
}

/**
 * State root batch, including header and actual state roots.
 */
export interface StateRootBatch {
  header: StateRootBatchHeader
  stateRoots: string[]
}

/**
 * Utility type for deep partials.
 */
export type DeepPartial<T> = {
  [P in keyof T]?: DeepPartial<T[P]>
}

/**
 * Extended Ethers overrides object with an l2GasLimit field.
 * Only meant to be used for L1 to L2 messages, since L2 to L1 messages don't have a specified gas
 * limit field (gas used depends on the amount of gas provided).
 */
export type L1ToL2Overrides = Overrides & {
  l2GasLimit: NumberLike
}

/**
 * Stuff that can be coerced into a transaction.
 */
export type TransactionLike = string | TransactionReceipt | TransactionResponse

/**
 * Stuff that can be coerced into a message.
 */
export type MessageLike =
  | CrossChainMessage
  | TransactionLike
  | TokenBridgeMessage

/**
 * Stuff that can be coerced into a provider.
 */
export type ProviderLike = string | Provider | any

/**
 * Stuff that can be coerced into a signer.
 */
export type SignerLike = string | Signer

/**
 * Stuff that can be coerced into a signer or provider.
 */
export type SignerOrProviderLike = SignerLike | ProviderLike

/**
 * Stuff that can be coerced into an address.
 */
export type AddressLike = string | Contract

/**
 * Stuff that can be coerced into a number.
 */
export type NumberLike = string | number | BigNumber
