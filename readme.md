# Starport Network

Starport Network (SPN) is a network for launching [Cosmos](https://cosmos.network) blockchains by matching coordinators with validators. This repository contains an SPN node built with Cosmos SDK and generated with [Starport](https://github.com/tendermint/starport).

## Using SPN

To use SPN you will need [Starport](https://github.com/tendermint/starport) installed. SPN is being actively developed, please, build Starport from source and use `develop` branch.

### Initiating a blockchain launch

To initiate a blockchain launch run the following command:

```
starport network chain create --chain [chainID] --source [sourceURL]
```

`chainID` is a string that uniquely identifies your blockchain on SPN. `sourceURL` is a URL that can be used to clone the repository containing a Cosmos SDK blockchain node (for example, `https://github.com/tendermint/spn`). By running the `create` command you act as a "coordinator" and initiate the launch of a blockchain.

By default a coordinator does not propose themselves as a validator. To do so, run `join` command and your proposal will be automatically approved.

### Joining as a validator

Run the following command from a server to propose yourself as a validator:

```
starport network chain join [chainID]
```

Follow the prompts to provide information about the validator. Starport will download the source code of the blockchain node, build, initialize and create and send two proposals to SPN: to add an account and to add a validator with self-delegation. By running a `join` command you act as a "validator".

### Listing pending proposals

```
starport network proposal list
```

This command lists all pending proposals. To see accepted and rejected proposals, use `--accepted` and `--rejected` flags. Each proposal has a `proposalID` (integer, unique to the chain), this ID is used to approve and reject a proposal.

### Accepting and rejecting proposals

As a coordinator run the following command to approve proposals:

```
starport network proposal approve [chainID] 1,4,5,6
```

Replace comma-separated values with a list of `proposalID` being accepted. Replace `approve` with `reject` to reject proposals instead.

### Starting a blockchain node

Once validator proposals have been accepted, run the following command to start a blockchain node:

```
starport network chain start [chainID]
```

This command will use SPN to create a correct genesis file, configure and launch your blockchain node. Once the node is started and the required number of validators are online, you will see output with incrementing block height number, which means that the blockchain has been successfully started.

## Learn more

* [Starport](https://github.com/tendermint/starport)
* [Cosmos Network](https://cosmos.network)
* [Cosmos Community Discord](https://discord.com/invite/W8trcGV) (check out the #starport channel)
