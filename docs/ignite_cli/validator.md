# Validator CLI Guide

Validators are Cosmos SDK blockchain validators using `spn`

---

## List all published chains

```shell
ignite n chain list
```

### Output

```shell
Launch Id 	Chain Id 	Source ...
6 		      planet-1 	https://github.com/lubtd/planet
5 		      spn-11 		https://github.com/tendermint/spn.git
4 		      spn-11 		https://github.com/tendermint/spn.git
3 		      mars-1 		https://github.com/lubtd/planet.git
2 		      spn-10 		https://github.com/tendermint/spn.git
1 		      spn-1 		https://github.com/tendermint/spn
```

---


## Request network participation

### Simple Flow

`ignite` can handle validator setup automatically.  Initialize the node and generate a gentx file with default values:

```shell
ignite n chain init 6
```

*NOTE: here "6" is specifying the `LaunchID`*

#### Output

```shell
✔ Source code fetched
✔ Blockchain set up
✔ Blockchain initialized
✔ Genesis initialized
? Staking amount 95000000stake
? Commission rate 0.10
? Commission max rate 0.20
? Commission max change rate 0.01
⋆ Gentx generated: /Users/lucas/spn/6/config/gentx/gentx.json
```

Now, create and broadcast a request to join a chain as a validator:

```shell
ignite n chain join 6 --amount 100000000stake
```

The join command accepts an optional --amount flag with a comma-separated list of tokens. If the flag is provided, the 
command will broadcast a request to add the validator’s address as an account to the genesis with the specific amount.

#### Output

```shell
? Peer's address 84.118.211.157:26656
✔ Source code fetched
✔ Blockchain set up
✔ Account added to the network by the coordinator!
✔ Validator added to the network by the coordinator!
```

---

### Advanced Flow

Using a more advanced setup (e.g. custom `gentx`), validators must provide an additional flag to their command
to point to the custom file:

```shell
ignite n chain join 6 --amount 100000000stake --gentx ~/chain/config/gentx/gentx.json
```

---


## Launch the network


### Simple Flow

Generate the final genesis and config of the node:

```shell
ignite n chain prepare 6
```

#### Output

```shell
✔ Source code fetched
✔ Blockchain set up
✔ Chain's binary built
✔ Genesis initialized
✔ Genesis built
✔ Chain is prepared for launch
```

Next, start the node:

```shell
planetd start --home ~/spn/6
```


---

### Advanced Flow

Fetch the final genesis for the chain:

```shell
ignite n chain show genesis 6
```

#### Output

```shell
✔ Source code fetched
✔ Blockchain set up
✔ Blockchain initialized
✔ Genesis initialized
✔ Genesis built
⋆ Genesis generated: ./genesis.json
```

Next, fetch the persistent peer list:

```shell
ignite n chain show peers 6
```

#### Output

```shell
⋆ Peer list generated: ./peers.txt
```

The fetched genesis file and peer list can be used for a manual node setup.


