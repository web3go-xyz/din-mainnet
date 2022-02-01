import { ethers, Overrides, Signer, BigNumber } from 'ethers'
import {
  TransactionRequest,
  TransactionResponse,
} from '@ethersproject/abstract-provider'
import { predeploys } from '@eth-optimism/contracts'

import {
  CrossChainMessageRequest,
  ICrossChainMessenger,
  ICrossChainProvider,
  MessageLike,
  NumberLike,
  AddressLike,
  MessageDirection,
} from './interfaces'

export class CrossChainMessenger implements ICrossChainMessenger {
  provider: ICrossChainProvider
  l1Signer: Signer
  l2Signer: Signer

  /**
   * Creates a new CrossChainMessenger instance.
   *
   * @param opts Options for the messenger.
   * @param opts.provider CrossChainProvider to use to send messages.
   * @param opts.l1Signer Signer to use to send messages on L1.
   * @param opts.l2Signer Signer to use to send messages on L2.
   */
  constructor(opts: {
    provider: ICrossChainProvider
    l1Signer: Signer
    l2Signer: Signer
  }) {
    this.provider = opts.provider
    this.l1Signer = opts.l1Signer
    this.l2Signer = opts.l2Signer
  }

  public async sendMessage(
    message: CrossChainMessageRequest,
    opts?: {
      l2GasLimit?: NumberLike
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    const tx = await this.populateTransaction.sendMessage(message, opts)
    if (message.direction === MessageDirection.L1_TO_L2) {
      return this.l1Signer.sendTransaction(tx)
    } else {
      return this.l2Signer.sendTransaction(tx)
    }
  }

  public async resendMessage(
    message: MessageLike,
    messageGasLimit: NumberLike,
    opts?: {
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return this.l1Signer.sendTransaction(
      await this.populateTransaction.resendMessage(
        message,
        messageGasLimit,
        opts
      )
    )
  }

  public async finalizeMessage(
    message: MessageLike,
    opts?: {
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return this.l1Signer.sendTransaction(
      await this.populateTransaction.finalizeMessage(message, opts)
    )
  }

  public async depositETH(
    amount: NumberLike,
    opts?: {
      l2GasLimit?: NumberLike
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return this.l1Signer.sendTransaction(
      await this.populateTransaction.depositETH(amount, opts)
    )
  }

  public async withdrawETH(
    amount: NumberLike,
    opts?: {
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return this.l2Signer.sendTransaction(
      await this.populateTransaction.withdrawETH(amount, opts)
    )
  }

  public async depositERC20(
    l1Token: AddressLike,
    l2Token: AddressLike,
    amount: NumberLike,
    opts?: {
      l2GasLimit?: NumberLike
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return this.l1Signer.sendTransaction(
      await this.populateTransaction.depositERC20(
        l1Token,
        l2Token,
        amount,
        opts
      )
    )
  }

  public async withdrawERC20(
    l1Token: AddressLike,
    l2Token: AddressLike,
    amount: NumberLike,
    opts?: {
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return this.l2Signer.sendTransaction(
      await this.populateTransaction.withdrawERC20(
        l1Token,
        l2Token,
        amount,
        opts
      )
    )
  }

  populateTransaction = {
    sendMessage: async (
      message: CrossChainMessageRequest,
      opts?: {
        l2GasLimit?: NumberLike
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      if (message.direction === MessageDirection.L1_TO_L2) {
        return this.provider.contracts.l1.L1CrossDomainMessenger.populateTransaction.sendMessage(
          message.target,
          message.message,
          opts?.l2GasLimit ||
            (await this.provider.estimateL2MessageGasLimit(message)),
          opts?.overrides || {}
        )
      } else {
        return this.provider.contracts.l2.L2CrossDomainMessenger.populateTransaction.sendMessage(
          message.target,
          message.message,
          0, // Gas limit goes unused when sending from L2 to L1
          opts?.overrides || {}
        )
      }
    },

    resendMessage: async (
      message: MessageLike,
      messageGasLimit: NumberLike,
      opts?: {
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      const resolved = await this.provider.toCrossChainMessage(message)
      if (resolved.direction === MessageDirection.L2_TO_L1) {
        throw new Error(`cannot resend L2 to L1 message`)
      }

      return this.provider.contracts.l1.L1CrossDomainMessenger.populateTransaction.replayMessage(
        resolved.target,
        resolved.sender,
        resolved.message,
        resolved.messageNonce,
        resolved.gasLimit,
        messageGasLimit,
        opts?.overrides || {}
      )
    },

    finalizeMessage: async (
      message: MessageLike,
      opts?: {
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      const resolved = await this.provider.toCrossChainMessage(message)
      if (resolved.direction === MessageDirection.L1_TO_L2) {
        throw new Error(`cannot finalize L1 to L2 message`)
      }

      const proof = await this.provider.getMessageProof(resolved)
      return this.provider.contracts.l1.L1CrossDomainMessenger.populateTransaction.relayMessage(
        resolved.target,
        resolved.sender,
        resolved.message,
        resolved.messageNonce,
        proof,
        opts?.overrides || {}
      )
    },

    depositETH: async (
      amount: NumberLike,
      opts?: {
        l2GasLimit?: NumberLike
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      return this.provider.bridges.ETH.populateTransaction.deposit(
        ethers.constants.AddressZero,
        predeploys.OVM_ETH,
        amount,
        opts
      )
    },

    withdrawETH: async (
      amount: NumberLike,
      opts?: {
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      return this.provider.bridges.ETH.populateTransaction.withdraw(
        ethers.constants.AddressZero,
        predeploys.OVM_ETH,
        amount,
        opts
      )
    },

    depositERC20: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        l2GasLimit?: NumberLike
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      const bridge = await this.provider.getBridgeForTokenPair(l1Token, l2Token)
      return bridge.populateTransaction.deposit(l1Token, l2Token, amount, opts)
    },

    withdrawERC20: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      const bridge = await this.provider.getBridgeForTokenPair(l1Token, l2Token)
      return bridge.populateTransaction.withdraw(l1Token, l2Token, amount, opts)
    },
  }

  estimateGas = {
    sendMessage: async (
      message: CrossChainMessageRequest,
      opts?: {
        l2GasLimit?: NumberLike
        overrides?: Overrides
      }
    ): Promise<BigNumber> => {
      const tx = await this.populateTransaction.sendMessage(message, opts)
      if (message.direction === MessageDirection.L1_TO_L2) {
        return this.provider.l1Provider.estimateGas(tx)
      } else {
        return this.provider.l2Provider.estimateGas(tx)
      }
    },

    resendMessage: async (
      message: MessageLike,
      messageGasLimit: NumberLike,
      opts?: {
        overrides?: Overrides
      }
    ): Promise<BigNumber> => {
      return this.provider.l1Provider.estimateGas(
        await this.populateTransaction.resendMessage(
          message,
          messageGasLimit,
          opts
        )
      )
    },

    finalizeMessage: async (
      message: MessageLike,
      opts?: {
        overrides?: Overrides
      }
    ): Promise<BigNumber> => {
      return this.provider.l1Provider.estimateGas(
        await this.populateTransaction.finalizeMessage(message, opts)
      )
    },

    depositETH: async (
      amount: NumberLike,
      opts?: {
        l2GasLimit?: NumberLike
        overrides?: Overrides
      }
    ): Promise<BigNumber> => {
      return this.provider.l1Provider.estimateGas(
        await this.populateTransaction.depositETH(amount, opts)
      )
    },

    withdrawETH: async (
      amount: NumberLike,
      opts?: {
        overrides?: Overrides
      }
    ): Promise<BigNumber> => {
      return this.provider.l2Provider.estimateGas(
        await this.populateTransaction.withdrawETH(amount, opts)
      )
    },

    depositERC20: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        l2GasLimit?: NumberLike
        overrides?: Overrides
      }
    ): Promise<BigNumber> => {
      return this.provider.l1Provider.estimateGas(
        await this.populateTransaction.depositERC20(
          l1Token,
          l2Token,
          amount,
          opts
        )
      )
    },

    withdrawERC20: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        overrides?: Overrides
      }
    ): Promise<BigNumber> => {
      return this.provider.l2Provider.estimateGas(
        await this.populateTransaction.withdrawERC20(
          l1Token,
          l2Token,
          amount,
          opts
        )
      )
    },
  }
}
