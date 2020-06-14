# store

The store package is an internal dependency of [texa](https://github.com/TexaProject/texa), used to store the CID hashes of the result data from the interrogation session into the ethereum blockchain.

## Usage

1. Update the smart contract `TexaResultStorage.sol`, as needed. Now, install solcjs using the command:
```sh
$ sudo npm -g install solc
```

2. Generate the ABI and bytecode

```sh
$ solcjs --abi TexaResultStorage.sol
$ solcjs --bin TexaResultStorage.sol
```
This will generate 2 files in the PWD.

3. Install abigen and generate the `store.go` file using the abigen command

First, let us install the abigen:

```sh
$ go get -u github.com/ethereum/go-ethereum
$ cd $GOPATH/src/github.com/ethereum/go-ethereum/
$ make
$ make devtools
```

Now, let us use the abigen command to generate the `store.go` file:

```sh
$ abigen --bin=TexaResultStorage_sol_TexaResultStorage.bin --abi=TexaResultStorage_sol_TexaResultStorage.abi --pkg=store --out=store.go
```
