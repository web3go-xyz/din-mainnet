# atst cli docs

![preview](../assets/preview.gif)

## Installation

```bash
npm install @eth-optimism/atst --global
```

## Usage

```bash
npx atst <command> [options]
```

## Commands

read read an attestation
write write an attestation

For more info, run any command with the `--help` flag:

```bash
npx atst read --help
npx atst write --help
```

## Options:

-h, --help Display this message
-v, --version Display version number

## Usage:

```bash
npx atst <command> [options]
```

Commands:
read read an attestation
write write an attestation

For more info, run any command with the `--help` flag:

```bash
npx atst read --help
npx atst write --help
```

### Read

`--creator <string> Address of the creator of the attestation`
`--about <string> Address of the subject of the attestation`
`--key <string> Key of the attestation either as string or hex number`
`--data-type [string] Zod validator for the DataType type string | bytes | number | bool | address (default: string)`
`--rpc-url [url] Rpc url to use (default: https://mainnet.optimism.io)`
`--contract [address] Contract address to read from (default: 0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77)`
`-h, --help Display this message`

Example:

```bash
npx atst read --key "optimist.base-uri" --about 0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5 --creator 0x60c5C9c98bcBd0b0F2fD89B24c16e533BaA8CdA3
```

### write

`--private-key <string> Address of the creator of the attestation`
`--data-type [string] Zod validator for the DataType type string | bytes | number | bool | address (default: string)`
`--about <string> Address of the subject of the attestation`
`--key <string> Key of the attestation either as string or hex number`
`--value <string> undefined`
`--rpc-url [url] Rpc url to use (default: https://mainnet.optimism.io)`
`--contract [address] Contract address to read from (default: 0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77) -h, --help Display this message`

```bash
atst write --key "optimist.base-uri" --about 0x2335022c740d17c2837f9C884Bfe4fFdbf0A95D5 --value "my attestation" --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --rpc-url http://localhost:8545
atst/0.0.0
```
