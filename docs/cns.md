# Chain Name Service Draft

The goal of this module is to be a simple registry of Cosmos Blockchains with the ability to transfer ownership.

The development of CNS is planned to be divided into 3 separate phases in order to ensure there's enough time for existing chains and prevent front running

**Phase 0** - Here, we expect owners of existing chains to claim their chain names which are approved by a previously agreed upon multi-sig. Estimated: 6 months.

Ideas:

1) Have a fixed price for reserved names, which would be transfered to community pool.

**Phase 1** - After an agreed-upon time is over, we would open CNS to general users to register chain names of their choice.

From this phase, the name will have an expiry time.

**Phase 2** - Finally, we add a way to transfer ownership of chain names. It's still to be determined if auction should be implemented as a module, or make CNS NFT compatible and let NFT handle the transfer or a CosmWasm smart contract which handles the auction utilizing the transfer ownership feature of CNS.

# State

Goal 1 (end-users): let users determine whether the tokens they got through IBC came from a chain they claim to come from. This will be possible through the mapping of chain names to IBC clients.

Goal 2 (validators): let nodes (validators, full-nodes, etc.) join networks.

```go
type ChainInfo struct {
	ChainName     string
	Owner         OwnerProps
	Expiration    int64
	Metadata      [][2]string

	// Goal 1
	CanonicalIBCClientId string

	// Goal 2
	Seed          []seed
	SourceCodeURL string //how likely is the possibility of code url changing
	Version       VersionInfo
}
```

```go
type VersionInfo struct {
	 Version          int64
	 SourceCommitHash string
	 GenesisHash      string
}
```

```go
type OwnerProps interface {
	String()  string
}
```

```go
type Claim struct {
  ID        int64
  ChainName string
  Owner     OwnerProps
  Proof     string
}
```

# Messages

### Phase - 0

Ideally, we would want the chains to approve via a gov proposal the owner (likely address) of the chain name. This makes it easy for approvers and also prevent collision for the same namespace.

### MsgClaimChainName

Claim is about mapping chain name to owner. Name, owner, IBC client, meta.

```go
type MsgClaimChainName struct {
	ChainName     string
	CanonicalIBCClientId string
	Proof         string
	Owner         OwnerProps
}
```

### MsgApproveClaim

```go
type MsgApproveClaim struct {
	ClaimID  int32
	Approver sdk.AccAddress
}
```

### MsgRejectClaim

```go
type MsgRejectClaim struct{
	ClaimID int32
	Approver sdk.AccAddress
}
```

### MsgUpdateInfo

```go
type MsgUpdateInfo struct{
	ChainName     string
	Owner         OwnerProps
	Seed          []seed
	CanonicalIBCClientId string
	Version       VersionInfo
	Metadata      [][2]string
}
```