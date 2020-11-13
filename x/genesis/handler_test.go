package genesis_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/genesis"
	"github.com/tendermint/spn/x/genesis/types"
	"testing"
)

func TestHandleMsgChainCreate(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)

	chainID := spnmocks.MockRandomAlphaString(5)
	creator := spnmocks.MockAccAddress()
	sourceURL := spnmocks.MockRandomString(20)
	sourceHash := spnmocks.MockRandomString(20)
	genesis := spnmocks.MockGenesis()

	// A chain can be create
	msg := types.NewMsgChainCreate(
		chainID,
		creator,
		sourceURL,
		sourceHash,
		genesis,
	)
	_, err := h(ctx, msg)
	require.NoError(t, err)
	retrieved, found := k.GetChain(ctx, chainID)
	require.True(t, found)
	creatorIdentity, _ := k.IdentityKeeper.GetIdentifier(ctx, creator)
	require.Equal(t, creatorIdentity, retrieved.Creator)
	require.Equal(t, sourceURL, retrieved.SourceURL)
	require.Equal(t, sourceHash, retrieved.SourceHash)

	// Prevent adding an existing chain id
	msg = types.NewMsgChainCreate(
		chainID,
		creator,
		sourceURL,
		sourceHash,
		genesis,
	)
	_, err = h(ctx, msg)
	require.Error(t, err)
}

func TestHandleMsgReject(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	chainID := spnmocks.MockRandomAlphaString(5)
	coordinator := spnmocks.MockAccAddress()

	// Prevent rejecting in a new existing chain
	msg := types.NewMsgReject(
		chainID,
		0,
		coordinator,
		)
	_, err := h(ctx, msg)
	require.Error(t, err)

	// Create a new chain
	msgChainCreate := types.NewMsgChainCreate(
		chainID,
		coordinator,
		spnmocks.MockRandomAlphaString(10),
		spnmocks.MockRandomAlphaString(10),
		spnmocks.MockGenesis(),
		)
	h(ctx, msgChainCreate)

	// Prevent rejecting a non existing proposal
	msg = types.NewMsgReject(
		chainID,
		0,
		coordinator,
	)
	_, err = h(ctx, msg)
	require.Error(t, err)

	// Create the proposal
	proposalCreator := spnmocks.MockAccAddress()
	msgProposal := types.NewMsgProposalAddAccount(
		chainID,
		proposalCreator,
		spnmocks.MockProposalAddAccountPayload(),
	)
	h(ctx, msgProposal)

	// Prevent an address other than the coordinator or the proposal creator to reject the proposal
	msg = types.NewMsgReject(
		chainID,
		0,
		spnmocks.MockAccAddress(),
	)
	_, err = h(ctx, msg)
	require.Error(t, err)

	// The proposal creator can reject the proposal
	msg = types.NewMsgReject(
		chainID,
		0,
		proposalCreator,
	)
	_, err = h(ctx, msg)
	require.NoError(t, err)

	// The proposal is rejected
	proposal, _ := k.GetProposal(ctx, chainID, 0)
	require.Equal(t, types.ProposalState_REJECTED, proposal.ProposalState.Status)

	// The proposal is not in pending pool
	pending := k.GetPendingProposals(ctx, chainID)
	require.NotContains(t, pending.ProposalIDs, int32(0))

	// The proposal is in rejected pool
	rejected := k.GetRejectedProposals(ctx, chainID)
	require.Contains(t, rejected.ProposalIDs, int32(0))

	// Prevent rejecting an already rejected proposal
	msg = types.NewMsgReject(
		chainID,
		0,
		proposalCreator,
	)
	_, err = h(ctx, msg)
	require.Error(t, err)

	// The coordinator can reject a proposal
	msgProposal = types.NewMsgProposalAddAccount(
		chainID,
		spnmocks.MockAccAddress(),
		spnmocks.MockProposalAddAccountPayload(),
	)
	h(ctx, msgProposal)
	msg = types.NewMsgReject(
		chainID,
		1,
		coordinator,
	)
	_, err = h(ctx, msg)
	require.NoError(t, err)
	proposal, _ = k.GetProposal(ctx, chainID, 1)
	require.Equal(t, types.ProposalState_REJECTED, proposal.ProposalState.Status)
	pending = k.GetPendingProposals(ctx, chainID)
	require.NotContains(t, pending.ProposalIDs, int32(1))
	rejected = k.GetRejectedProposals(ctx, chainID)
	require.Contains(t, rejected.ProposalIDs, int32(1))
}

func TestHandleMsgProposalAddAccount(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	chainID := spnmocks.MockRandomAlphaString(5)

	// Prevent creating a proposal for a non existing chain
	msg := types.NewMsgProposalAddAccount(
		chainID,
		spnmocks.MockAccAddress(),
		spnmocks.MockProposalAddAccountPayload(),
	)
	_, err := h(ctx, msg)
	require.Error(t, err)

	// Create a new chain
	chain := spnmocks.MockChain()
	chain.ChainID = chainID
	k.SetChain(ctx, *chain)

	// Can add the proposal
	creator := spnmocks.MockAccAddress()
	creatorIdentity, _ := k.IdentityKeeper.GetIdentifier(ctx, creator)
	msg = types.NewMsgProposalAddAccount(
		chainID,
		creator,
		spnmocks.MockProposalAddAccountPayload(),
	)
	_, err = h(ctx, msg)
	require.NoError(t, err)

	// Can retrieve the proposal
	proposal, found := k.GetProposal(ctx, chainID, 0)
	require.True(t, found)
	require.Equal(t, creatorIdentity, proposal.ProposalInformation.Creator)
	_, ok := proposal.Payload.(*types.Proposal_AddAccountPayload)
	require.True(t, ok)

	// The proposal is added to the pending proposals
	pending := k.GetPendingProposals(ctx, chainID)
	require.Contains(t, pending.ProposalIDs, int32(0))

	// The proposal count is incremented
	count := k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(1), count)
}

func TestHandleMsgProposalAddValidator(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	chainID := spnmocks.MockRandomAlphaString(5)

	// Prevent creating a proposal for a non existing chain
	msg := types.NewMsgProposalAddValidator(
		chainID,
		spnmocks.MockAccAddress(),
		spnmocks.MockProposalAddValidatorPayload(),
	)
	_, err := h(ctx, msg)
	require.Error(t, err)

	// Create a new chain
	chain := spnmocks.MockChain()
	chain.ChainID = chainID
	k.SetChain(ctx, *chain)

	// We send n-1 proposals to check proposal n behavior
	creator := spnmocks.MockAccAddress()
	creatorIdentity, _ := k.IdentityKeeper.GetIdentifier(ctx, creator)
	msg = types.NewMsgProposalAddValidator(
		chainID,
		creator,
		spnmocks.MockProposalAddValidatorPayload(),
	)
	_, err = h(ctx, msg)
	_, err = h(ctx, msg)
	_, err = h(ctx, msg)
	msgAccount := types.NewMsgProposalAddAccount(
		chainID,
		creator,
		spnmocks.MockProposalAddAccountPayload(),
	)
	_, err = h(ctx, msgAccount)
	_, err = h(ctx, msgAccount)

	// Can add the new proposal n
	msg = types.NewMsgProposalAddValidator(
		chainID,
		creator,
		spnmocks.MockProposalAddValidatorPayload(),
	)
	_, err = h(ctx, msg)
	require.NoError(t, err)

	// Can retrieve the proposal
	proposal, found := k.GetProposal(ctx, chainID, 5)
	require.True(t, found)
	require.Equal(t, creatorIdentity, proposal.ProposalInformation.Creator)
	_, ok := proposal.Payload.(*types.Proposal_AddValidatorPayload)
	require.True(t, ok)

	// The proposal is added to the pending proposals
	pending := k.GetPendingProposals(ctx, chainID)
	require.Contains(t, pending.ProposalIDs, int32(5))

	// The proposal count is incremented
	count := k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(6), count)
}