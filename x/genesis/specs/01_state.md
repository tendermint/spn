<!--
order: 1
-->

# State

## Chain

Basic information of the chain is recorded on-chain. It contains the necessary information for the users to be able to retrieve download and build the chain and store the initial genesis.

key: `"chain-[chainID]" -> Chain`

```go
type Chain struct {
	ChainID string // Alphanumeric, with "-" allowed
	Creator string
	Peers []string
  SourceURL string // Source code URL
	SourceHash string
	CreatedAt uint64 // Unix time
	Genesis GenesisFile
	Final bool
}

// Represent the bytes of a tendermint/types.GenesisDoc structure
type GenesisFile []byte
```

## Proposal

Proposals represent propositions to update a genesis before launching the chain. It can be about adding new genesis accounts, genesis transactions (with `genutils` module), or modification on the params of a module. A proposal contains generic information, a state of its approval process, and a payload related to the type of change.

Currently, a proposal can be approved by the coordinator of the chain and rejected by the coordinator or the creator of the proposal.

`ProposalID` is an incrementing integer local to `Chain`.

key: `"proposal-[chainID]_[proposalID]" -> Proposal`

```go
type Proposal struct {
	proposalInformation ProposalInformation
	proposalState ProposalState
	payload [ProposalAddAccountPayload|ProposalAddValidatorPayload]
}

// Generic information about a proposal
type Proposal struct {
	chainID string
	proposalID int32
	creator string
	createdAt uint64 // Unix time
}

type ProposalState struct {
	status ProposalState_Status // [APPROVED|REJECTED|PENDING]
}

type ProposalAddAccountPayload struct {
	address sdk.AccAddress
	coins sdk.Coins
}

type ProposalAddValidatorPayload struct {
	gentTx tx.Tx // Tx type definition from cosmos-sdk/types/tx.Tx
	peer string // Peer definition of the node stored in persistent peers
}
```

## ProposalPools

They contain array of proposal ID to retrieve proposal based on their status

```go
"approvedProposal-[chainID]" -> []int32
"rejectedProposal-[chainID]" -> []int32
"pendingProposal-[chainID]" -> []int32
```

## GenesisInternalState

In order to approve a new proposal, we may need to perform some check on the current state of the genesis. For example, a validator is only valid if it contains an associated account with enough coins for self-delegation. To perform these checks efficiently without having the need to iterate through all approved proposals, we store on-chain an internal genesis state that is never interpreted by the users but used internally by `spn` to efficiently determine if a proposal can be approved.

### Accounts

This contains the payload of an AddAccountProposal directly indexed by the account address.

key: `"account-[chainID]_[accAddress]" -> ProposalAddAccountPayload`

### Validators

This contains a constant byte literal indexed by the validator address to determine if a validator exists for a specific chain's genesis.

key: `"validator-[chainID]_[valAddress]" -> [nil|1]`
