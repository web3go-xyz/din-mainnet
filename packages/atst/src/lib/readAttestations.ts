import { readContracts } from '@wagmi/core'
import { formatBytes32String } from 'ethers/lib/utils.js'

import { ATTESTATION_STATION_ADDRESS } from '../constants/attestationStationAddress'
import type { AttestationReadParams } from '../types/AttestationReadParams'
import type { WagmiBytes } from '../types/WagmiBytes'
import { abi } from './abi'
import { parseAttestationBytes } from './parseAttestationBytes'

/**
 * reads attestations from the attestation station contract
 *
 * @returns an array of attestation values
 * @throws Error if key is longer than 32 bytes
 * @example
 * const attestations = await readAttestations(
 *  {
 *    creator: creatorAddress,
 *    about: aboutAddress,
 *    key: 'my_key',
 *    allowFailure: false,
 *  },
 *  {
 *    creator: creatorAddress2,
 *    about: aboutAddress2,
 *    key: 'my_key',
 *    dataType: 'number',
 *    contractAddress: '0x1234',
 *    allowFailure: false,
 *   },
 * )
 */
export const readAttestations = async (
  ...attestationReads: Array<AttestationReadParams>
) => {
  const calls = attestationReads.map((attestation) => {
    const {
      creator,
      about,
      key,
      contractAddress = ATTESTATION_STATION_ADDRESS,
      allowFailure = false,
    } = attestation
    if (key.length > 32) {
      throw new Error(
        'Key is longer than the max length of 32 for attestation keys'
      )
    }
    return {
      address: contractAddress,
      abi,
      functionName: 'attestations',
      args: [creator, about, formatBytes32String(key) as WagmiBytes],
      allowFailure,
    } as const
  })

  const results = await readContracts({
    contracts: calls,
  })

  return results.map((dataBytes, i) => {
    const dataType = attestationReads[i].dataType ?? 'string'
    return parseAttestationBytes(dataBytes, dataType)
  })
}
