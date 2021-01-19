# Chain Name Service Draft

The goal of this module is to be a simple registry with the ability to transfer ownership. Auctions and other mechanisms are out of scope and can be implemented alongside or on top of this system (perhaps, in a custodial way).

# State

## Chain name

Chain name contains information about the chain, most importantly, chain ID used to uniquly identify the chain and a list of peer nodes.

Key: `chain-[chainID] -> Chain Name` 

```go
type ChainName struct {
	ChainID string
	Owner sdk.AccAddress
	Peers []peer
	SourceURL string
	SourceCommitHash string
}
```

## Offer

Offers are used to transfer ownership of a chain name.

```go
type Offer struct {
	ID int32
	From sdk.AccAddress
	Amount sdk.Coins
	Status status // pending | rejected | approved
}
```

# Messages

## MsgEditChainName

Chain owner can make edits to the chain name.

```go
type MsgEditChainName struct {
	
}
```

## MsgTransferChainName

To transfer a name from one owner from another without creating an offer with a zero amount.

```go
type MsgTransferChainName struct {
	From string
	To string
	ChainID string
}
```

## MsgCreateOffer

```go
type MsgCreateOffer struct {
	ChainID string
	Amount sdk.Coins
}
```

## MsgRejectOffer

```go
type MsgRejectOffer struct {
	ChainID string
}
```

## MsgApproveOffer

```go
type MsgApproveOffer struct {
	ChainID string
}
```

## MsgRemoveOffer

```go
type MsgRejectOffer struct {
	ChainID string
}
```

## MsgRegisterName

To accuire a name that doesn't have an owner. By default each name has a price. When registering a name, amount equivalent to the price goes to the community pool.

```go
type MsgRejectOffer struct {
	ChainID string
}
```

# Ideas

- Reserve names for existing chains
- Names with a chain IDs that have `tesnet-` prefix are registered for free
- Off-chain tool to check that nodes in peer list are part of a network with a given chain ID. Can loop through all the chain names and create a table.
- If you're launching a chain and don't want a single entity to be responsible for managing a chain name, use a multisig
- This should probably be an IBC module
- Open question: how does relate to IBC, relayers. How can this module be used with IBC.