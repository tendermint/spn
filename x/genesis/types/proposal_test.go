package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/genesis/types"
	"testing"
)

func TestNewProposalState(t *testing.T) {
	// Can create a proposal state
	ps := types.NewProposalState()
	require.Equal(t, types.ProposalState_PENDING, ps.Status, "NewProposalState should create a pending proposal")
	require.Equal(t, 0, len(ps.Votes), "NewProposalState should create a proposal with no vote")

	// Can create a new vote

	// Can append a new vote in the proposal

	// Prevent creating an invalid vote

	// Prevent voting twice for the same identity

}

func TestNewProposalChange(t *testing.T) {
	// Can create a proposal for a genesis change

	// Prevent invalid change path
}

func TestNewProposalAddAccount(t *testing.T) {
	// Can create a new proposal to add an account

	// Prevent add account with invalid address

	// Prevent invalid coins allocation
}

func TestNewProposalAddValidator(t *testing.T) {
	// Can create a proposal to add a validator

}
