## Localnet

Localnet is a simple local testnet for Starport Network that includes 3 validators.

Running localnet requires Python 3.

The three nodes use the following config:

```
rpc: 26657
p2p: 26656
api: 1317
```
*Node 1 (Joe)*
```
p2p: 26655
```
*Node 2 (Steve)*
```
p2p: 26654
```
*Node 3 (Olivia)*

### Starting the Localnet

Before running any scripts, `spnd` must be built for the current branch
```
starport chain build
```

The localnet can be started from the following script:
(must be run in `localnet` directory)
```
python3 start.py
```

The logs of the Node 1 are printed in the console.

The chain state is reset everytime the script is executed.

### Configuration

Configuration for the localnet is defined in `conf.yml`.
Values are automatically updated when restarting the localnet.

### Updating the Validator Set

The validator set can be updated dynamically when the testnet is running through the `delegate.py` and `undelegate.py` scripts.
Usage:
```
python3 delegate.py   [val1_stake] [val2_stake] [val3_stake]
python3 undelegate.py [val1_stake] [val2_stake] [val3_stake]
```
When `0` is provided for a validator, the delegation is not updated

### Example

Run the localnet in a terminal
```
python3 start.py
```

Show the validator set
```
spnd q tendermint-validator-set
block_height: "3"
total: "3"
validators:
- address: spnvalcons1ypme3zv7zgh9v4ec0ae5t22vfkpk8cjvara424
  proposer_priority: "0"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: FcHc+u/13lyZ3fm2bWE46fpjbjqjfyrb6/X5r66gjFA=
  voting_power: "10"
- address: spnvalcons1xha40hx6k605kj3c2p80xwxmzkqvlwq22e92mk
  proposer_priority: "0"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: FyTmyvZhwRjwqhY6eWykTfiE+0mwe+U0aSo3ti8DCW8=
  voting_power: "10"
- address: spnvalcons1k3mq649cjr8vhhaw65raqq9v2d9j6ktpce95lj
  proposer_priority: "0"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: SIbr/rY/55BiXE6NBay7PmzBw25ADIrVtfVRqsqQBZM=
  voting_power: "10"
```

Perform delegation
```
python3 delegate.py 10000000 0 0
```

Validator 1 has now more voting power
```
spnd q tendermint-validator-set
block_height: "21"
total: "3"
validators:
- address: spnvalcons1xha40hx6k605kj3c2p80xwxmzkqvlwq22e92mk
  proposer_priority: "20"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: FyTmyvZhwRjwqhY6eWykTfiE+0mwe+U0aSo3ti8DCW8=
  voting_power: "20"
- address: spnvalcons1ypme3zv7zgh9v4ec0ae5t22vfkpk8cjvara424
  proposer_priority: "-10"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: FcHc+u/13lyZ3fm2bWE46fpjbjqjfyrb6/X5r66gjFA=
  voting_power: "10"
- address: spnvalcons1k3mq649cjr8vhhaw65raqq9v2d9j6ktpce95lj
  proposer_priority: "-10"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: SIbr/rY/55BiXE6NBay7PmzBw25ADIrVtfVRqsqQBZM=
  voting_power: "10"
```

Perform undelegation
```
python3 undelegate.py 10000000 0 0
```

Validator 1 voting power is decreased back
```
spnd q tendermint-validator-set
block_height: "67"
total: "3"
validators:
- address: spnvalcons1ypme3zv7zgh9v4ec0ae5t22vfkpk8cjvara424
  proposer_priority: "-20"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: FcHc+u/13lyZ3fm2bWE46fpjbjqjfyrb6/X5r66gjFA=
  voting_power: "10"
- address: spnvalcons1xha40hx6k605kj3c2p80xwxmzkqvlwq22e92mk
  proposer_priority: "10"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: FyTmyvZhwRjwqhY6eWykTfiE+0mwe+U0aSo3ti8DCW8=
  voting_power: "10"
- address: spnvalcons1k3mq649cjr8vhhaw65raqq9v2d9j6ktpce95lj
  proposer_priority: "10"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: SIbr/rY/55BiXE6NBay7PmzBw25ADIrVtfVRqsqQBZM=
  voting_power: "10"
```