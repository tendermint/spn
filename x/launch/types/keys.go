package types

import (
	"strconv"
)

const (
	// ModuleName defines the module name
	ModuleName = "launch"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// ChainKey is the key to store chain info
	ChainKey = "chain-"

	// ProposalKey is the key to store the proposals
	ProposalKey = "proposal-"

	// ProposalCountKey is the keyu to retrieve the count of proposal
	ProposalCountKey = "proposalcount-"

	// ApprovedProposalKey is the key to store the approved proposal ids
	ApprovedProposalKey = "approvedproposal-"

	// PendingProposalKey is the key to store the pending proposal ids
	PendingProposalKey = "pendingproposal-"

	// RejectedProposalKey is the key to store the rejected proposal ids
	RejectedProposalKey = "rejectedproposal-"

	// AccountKey is the key to store retrieve existing accounts in the genesis
	AccountKey = "account-"

	// ValidatorKey is the key to store retrieve existing validators in the genesis
	ValidatorKey = "validator-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetChainKey returns the key for the chain store
func GetChainKey(chainID string) []byte {
	return append(KeyPrefix(ChainKey), []byte(chainID)...)
}

// GetProposalKey returns the key for the proposal store
func GetProposalKey(chainID string, proposalID int32) []byte {
	key := append(KeyPrefix(ProposalKey), []byte(chainID)...)

	// We use "_" to separate chainID and proposalID to avoid prefix conflic since "-" is allowed in chainID
	key = append(key, []byte("_")...)
	return append(key, []byte(strconv.Itoa(int(proposalID)))...)
}

// GetProposalCountKey returns the key for the proposal count store
func GetProposalCountKey(chainID string) []byte {
	return append(KeyPrefix(ProposalCountKey), []byte(chainID)...)
}

// GetApprovedProposalKey returns the the key for the approved proposal id store
func GetApprovedProposalKey(chainID string) []byte {
	return append(KeyPrefix(ApprovedProposalKey), []byte(chainID)...)
}

// GetPendingProposalKey returns the the key for the pending proposal id store
func GetPendingProposalKey(chainID string) []byte {
	return append(KeyPrefix(PendingProposalKey), []byte(chainID)...)
}

// GetRejectedProposalKey returns the the key for the rejected proposal id store
func GetRejectedProposalKey(chainID string) []byte {
	return append(KeyPrefix(RejectedProposalKey), []byte(chainID)...)
}

// GetAccountKey returns the key for accounts store
func GetAccountKey(chainID string, accountAddress string) []byte {
	key := append(KeyPrefix(AccountKey), []byte(chainID)...)

	// We use "_" to separate chainID and proposalID to avoid prefix conflic since "-" is allowed in chainID
	key = append(key, []byte("_")...)
	return append(key, []byte(accountAddress)...)
}

// GetValidatorKey returns the key for validators store
func GetValidatorKey(chainID string, valAddress string) []byte {
	key := append(KeyPrefix(ValidatorKey), []byte(chainID)...)

	// We use "_" to separate chainID and proposalID to avoid prefix conflic since "-" is allowed in chainID
	key = append(key, []byte("_")...)
	return append(key, []byte(valAddress)...)
}
