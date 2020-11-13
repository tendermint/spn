package keeper_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/genesis/types"
	"testing"
)

func TestGetProposalCount(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chainID := spnmocks.MockRandomAlphaString(5)

	// If count not set: return 0
	count := k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(0), count, "GetProposalCount should return 0 if no value")

	// Set new count
	k.SetProposalCount(ctx, chainID, 5)
	count = k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(5), count, "GetProposalCount should return the set value")
	count = k.GetProposalCount(ctx, "AnotherChain")
	require.Equal(t, int32(0), count, "GetProposalCount should return 0 if no value")
}

func TestGetProposal(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	proposal := spnmocks.MockProposal()

	// A non set proposal should not exist
	_, found := k.GetProposal(
		ctx,
		proposal.GetProposalInformation().GetChainID(),
		proposal.GetProposalInformation().GetProposalID(),
	)
	require.False(t, found, "GetProposal should not find a non existent proposal")

	// Set and get a proposal
	k.SetProposal(ctx, *proposal)
	retrieved, found := k.GetProposal(
		ctx,
		proposal.GetProposalInformation().GetChainID(),
		proposal.GetProposalInformation().GetProposalID(),
	)
	require.True(t, found, "GetProposal should find a proposal")
	require.Equal(t, *proposal, retrieved, "GetProposal should find a proposal")

	otherProposal := spnmocks.MockProposal()
	_, found = k.GetProposal(
		ctx,
		otherProposal.GetProposalInformation().GetChainID(),
		otherProposal.GetProposalInformation().GetProposalID(),
	)
	require.False(t, found, "GetProposal should not find a non existent proposal")
}

func TestGetApprovedProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chainID := spnmocks.MockRandomAlphaString(5)
	list := spnmocks.MockProposalList()
	emptyList := types.ProposalList{
		ProposalIDs: []int32{},
	}

	retrieved := k.GetApprovedProposals(ctx, chainID)
	require.Equal(t, emptyList, retrieved, "GetApprovedProposals should return empty list if not set")

	k.SetApprovedProposals(ctx, chainID, *list)
	retrieved = k.GetApprovedProposals(ctx, chainID)
	require.Equal(t, *list, retrieved, "GetApprovedProposals should return the set value")
	retrieved = k.GetApprovedProposals(ctx, "AnotherChain")
	require.Equal(t, emptyList, retrieved, "GetApprovedProposals should return empty list if not set")
}

func TestGetPendingProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chainID := spnmocks.MockRandomAlphaString(5)
	list := spnmocks.MockProposalList()
	emptyList := types.ProposalList{
		ProposalIDs: []int32{},
	}

	retrieved := k.GetPendingProposals(ctx, chainID)
	require.Equal(t, emptyList, retrieved, "GetPendingProposals should return empty list if not set")

	k.SetPendingProposals(ctx, chainID, *list)
	retrieved = k.GetPendingProposals(ctx, chainID)
	require.Equal(t, *list, retrieved, "GetPendingProposals should return the set value")
	retrieved = k.GetPendingProposals(ctx, "AnotherChain")
	require.Equal(t, emptyList, retrieved, "GetPendingProposals should return empty list if not set")
}

func TestGetRejectedProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chainID := spnmocks.MockRandomAlphaString(5)
	list := spnmocks.MockProposalList()
	emptyList := types.ProposalList{
		ProposalIDs: []int32{},
	}

	retrieved := k.GetRejectedProposals(ctx, chainID)
	require.Equal(t, emptyList, retrieved, "GetRejectedProposals should return empty list if not set")

	k.SetRejectedProposals(ctx, chainID, *list)
	retrieved = k.GetRejectedProposals(ctx, chainID)
	require.Equal(t, *list, retrieved, "GetRejectedProposals should return the set value")
	retrieved = k.GetRejectedProposals(ctx, "AnotherChain")
	require.Equal(t, emptyList, retrieved, "GetRejectedProposals should return empty list if not set")
}
