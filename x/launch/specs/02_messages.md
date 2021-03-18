<!--
order: 2
-->

# Messages

## MsgChainCreate

```go
type MsgChainCreate struct {
	ChainID string
	Creator sdk.AccAddress
	SourceURL string
	SourceHash string
	Genesis GenesisFile
}
```

Fails if:

1. A chain with `ChainID` exists in the store.
2. The genesis is not in the correct format

This message creates a new `Chain`. `Peers` array of the chain is empty.

## MsgProposalAddAccount

```go
type MsgProposalAddAccount struct {
	ChainID string
	Creator sdk.AccAddress
	Payload ProposalAddAccountPayload
}
```

Fails if:

1. Error if a chain with `ChainID` **doesn't exists** in the store.
2. Error if a proposal with `ProposalID` **exists** in the chain store.

This message creates a new Proposal in the chain store, sets its status to pending, appends the proposal ID in the pending proposals pool, and increments the proposal count of the chain.

## MsgProposalAddValidator

```go
type MsgProposalAddValidator struct {
	ChainID string
	Creator sdk.AccAddress
	Payload ProposalAddValidatorPayload
}
```

Fails if:

1. A chain with `ChainID` **doesn't exists** in the store.

This message creates a new Proposal in the chain store, sets its status to pending, appends the proposal ID in the pending proposals pool, and increments the proposal count of the chain.

## MsgApprove

```go
type MsgApprove struct {
	ChainID string
	ProposalID int32
	Approver sdk.AccAddress
}
```

Fails if:

1. A chain with `ChainID` **doesn't exists** in the store.
2. A proposal with `ProposalID` doesn't exists in the chain store.
3. The approver is not the coordinator of the chain.
4. The proposal is not in `pending` state.
5. The proposal can't be approved with the current state of the genesis.

This message sets the status of the proposal to `approved`, appends the `proposalID` to the approved proposals pool, and updates the genesis internal state of the chain.

## MsgReject

```go
type MsgReject struct {
	ChainID string
	ProposalID int32
	Rejector sdk.AccAddress
}
```

Fails if:

1. A chain with `ChainID` **doesn't exists** in the store.
2. A proposal with `ProposalID` doesn't exists in the chain store.
3. The rejector is neither the coordinator of the chain nor the creator of the proposal
4. The proposal is not in `pending` state.

This message sets the status of the proposal to `rejected` and appends the `proposalID` to the rejected proposals pool.