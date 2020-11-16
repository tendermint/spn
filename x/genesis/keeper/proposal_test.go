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
	require.Equal(t, int32(0), count)

	// Set new count
	k.SetProposalCount(ctx, chainID, 5)
	count = k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(5), count)
	count = k.GetProposalCount(ctx, "AnotherChain")
	require.Equal(t, int32(0), count)
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
	require.False(t, found)

	// Set and get a proposal
	k.SetProposal(ctx, *proposal)
	retrieved, found := k.GetProposal(
		ctx,
		proposal.GetProposalInformation().GetChainID(),
		proposal.GetProposalInformation().GetProposalID(),
	)
	require.True(t, found)
	require.Equal(t, *proposal, retrieved)

	otherProposal := spnmocks.MockProposal()
	_, found = k.GetProposal(
		ctx,
		otherProposal.GetProposalInformation().GetChainID(),
		otherProposal.GetProposalInformation().GetProposalID(),
	)
	require.False(t, found)
}

func TestGetProposalChange(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()

	// The retrieved proposal contains a valid payload
	proposal, _ := types.NewProposalChange(
		spnmocks.MockProposalInformation(),
		spnmocks.MockProposalChangePayload(),
	)
	k.SetProposal(ctx, *proposal)
	retrieved, _ := k.GetProposal(
		ctx,
		proposal.GetProposalInformation().GetChainID(),
		proposal.GetProposalInformation().GetProposalID(),
	)
	payload, ok := retrieved.Payload.(*types.Proposal_ChangePayload)
	require.True(t, ok)
	err := types.ValidateProposalPayloadChange(payload.ChangePayload)
	require.NoError(t, err)
}

func TestGetProposalAddAccount(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()

	// The retrieved proposal contains a valid payload
	proposal, _ := types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		spnmocks.MockProposalAddAccountPayload(),
		)
	k.SetProposal(ctx, *proposal)
	retrieved, _ := k.GetProposal(
		ctx,
		proposal.GetProposalInformation().GetChainID(),
		proposal.GetProposalInformation().GetProposalID(),
	)
	payload, ok := retrieved.Payload.(*types.Proposal_AddAccountPayload)
	require.True(t, ok)
	err := types.ValidateProposalPayloadAddAccount(payload.AddAccountPayload)
	require.NoError(t, err)
}

func TestGetProposalAddValidator(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()

	// The retrieved proposal contains a valid payload
	proposal, _ := types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		spnmocks.MockProposalAddValidatorPayload(),
	)
	k.SetProposal(ctx, *proposal)
	retrieved, _ := k.GetProposal(
		ctx,
		proposal.GetProposalInformation().GetChainID(),
		proposal.GetProposalInformation().GetProposalID(),
	)
	payload, ok := retrieved.Payload.(*types.Proposal_AddValidatorPayload)
	require.True(t, ok)
	err := types.ValidateProposalPayloadAddValidator(payload.AddValidatorPayload)
	require.NoError(t, err)
}


func TestGetApprovedProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chainID := spnmocks.MockRandomAlphaString(5)
	list := spnmocks.MockProposalList()
	emptyList := types.ProposalList{
		ProposalIDs: []int32{},
	}

	retrieved := k.GetApprovedProposals(ctx, chainID)
	require.Equal(t, emptyList, retrieved)

	k.SetApprovedProposals(ctx, chainID, *list)
	retrieved = k.GetApprovedProposals(ctx, chainID)
	require.Equal(t, *list, retrieved,)
	retrieved = k.GetApprovedProposals(ctx, "AnotherChain")
	require.Equal(t, emptyList, retrieved)
}

func TestGetPendingProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chainID := spnmocks.MockRandomAlphaString(5)
	list := spnmocks.MockProposalList()
	emptyList := types.ProposalList{
		ProposalIDs: []int32{},
	}

	retrieved := k.GetPendingProposals(ctx, chainID)
	require.Equal(t, emptyList, retrieved)

	k.SetPendingProposals(ctx, chainID, *list)
	retrieved = k.GetPendingProposals(ctx, chainID)
	require.Equal(t, *list, retrieved)
	retrieved = k.GetPendingProposals(ctx, "AnotherChain")
	require.Equal(t, emptyList, retrieved)
}

func TestGetRejectedProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chainID := spnmocks.MockRandomAlphaString(5)
	list := spnmocks.MockProposalList()
	emptyList := types.ProposalList{
		ProposalIDs: []int32{},
	}

	retrieved := k.GetRejectedProposals(ctx, chainID)
	require.Equal(t, emptyList, retrieved)

	k.SetRejectedProposals(ctx, chainID, *list)
	retrieved = k.GetRejectedProposals(ctx, chainID)
	require.Equal(t, *list, retrieved)
	retrieved = k.GetRejectedProposals(ctx, "AnotherChain")
	require.Equal(t, emptyList, retrieved)
}
