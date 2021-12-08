<!--![cannon](https://upload.wikimedia.org/wikipedia/commons/8/80/Cannon%2C_Château_du_Haut-Koenigsbourg%2C_France.jpg)-->
<!--![cannon](https://cdn1.epicgames.com/ue/product/Featured/SCIFIWEAPONBUNDLE_featured-894x488-83fbc936b6d86edcbbe892b1a6780224.png)-->
<!--![cannon](https://static.wikia.nocookie.net/ageofempires/images/8/80/Bombard_cannon_aoe2DE.png/revision/latest/top-crop/width/360/height/360?cb=20200331021834)-->
![cannon](https://paradacreativa.es/wp-content/uploads/2021/05/Canon-orbital-GTA-01.jpg)

The cannon (cannon cannon cannon) is an on chain interactive fraud prover

It's half geth, half of what I think truebit was supposed to be. It can prove L1 blocks aren't fraud.

* It's Go code
* ...that runs an EVM
* ...emulating a MIPS machine
* ...running compiled Go code
* ...that runs an EVM

## Directory Layout

```
minigeth -- A standalone "geth" capable of computing a block transition
mipigo -- minigeth compiled for MIPS. Outputs a MIPS binary that's run and mapped at 0x0
mipsevm -- A MIPS runtime in the EVM (see also contracts/)
```

## Usage
```
# build
(cd minigeth/ && go build)
mkdir -p /tmp/cannon

# compute the transition from 13284469 -> 13284470
minigeth/go-ethereum 13284469

# generate MIPS checkpoints for 13284469 -> 13284470
mipsevm/mipsevm 13284469
```

## Full Challenge / Response

```
npx hardhat node --fork https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161

# testing on hardhat (forked mainnet)
# challenger is pretending the block 10 transition is the transition for 1171895
# this will conflict at the first step
rm -rf /tmp/cannon/*
mipsevm/mipsevm
npx hardhat run scripts/deploy.js

minigeth/go-ethereum 1171895 && mipsevm/mipsevm 1171895
minigeth/go-ethereum 10 && mipsevm/mipsevm 10
BLOCK=1171895 npx hardhat run scripts/challenge.js

# do binary search
for i in {1..23}
do
ID=0 BLOCK=10 CHALLENGER=1 npx hardhat run scripts/respond.js
ID=0 BLOCK=1171895 npx hardhat run scripts/respond.js
done

# assert as challenger (fails)
ID=0 BLOCK=10 CHALLENGER=1 npx hardhat run scripts/assert.js

# assert as defender (passes)
ID=0 BLOCK=1171895 npx hardhat run scripts/assert.js
```

## State Oracle API

On chain / in MIPS, we have two oracles

* InputHash() -> hash        # this is a hash of the initial custom state of the system
* Preimage(hash) -> value    # hash(value) == hash

We generate the Preimages in x86 using geth RPC

* PrefetchAccount
* PrefetchStorage
* PrefetchCode
* PrefetchBlock

These are NOP in the VM

## License

All my code is MIT license, minigeth is LGPL3. Being developed under contract for @optimismPBC
