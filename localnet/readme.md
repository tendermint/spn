## Localnet

Localnet is a simple local testnet for Starport Network that includes 3 validators.

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

### Starting the localnet

The localnet can be started from the following script:
(must be run in `localnet` directory)
```
python3 start.py
```

The chain state is reset everytime the script is executed.