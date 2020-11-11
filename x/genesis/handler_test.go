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
	require.NoError(t, err, "NewMsgChainCreate with a correct chain should succeed")
	retrieved, found := k.GetChain(ctx, chainID)
	require.True(t, found, "NewMsgChainCreate should add the chain in the store")
	creatorIdentity, _ := k.IdentityKeeper.GetIdentifier(ctx, creator)
	require.Equal(t, creatorIdentity, retrieved.Creator, "NewMsgChainCreate should add the correct chain")
	require.Equal(t, sourceURL, retrieved.SourceURL, "NewMsgChainCreate should add the correct chain")
	require.Equal(t, sourceHash, retrieved.SourceHash, "NewMsgChainCreate should add the correct chain")

	// Prevent adding an existing chain id
	msg = types.NewMsgChainCreate(
		chainID,
		creator,
		sourceURL,
		sourceHash,
		genesis,
	)
	_, err = h(ctx, msg)
	require.Error(t, err, "NewMsgChainCreate should prevent adding an existing chain")
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
	require.Error(t, err, "MsgProposalAddAccount should fail for non existing chains")

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
	require.Error(t, err, "MsgProposalAddAccount should append a new proposal")

	// Can retrieve the proposal
	proposal, found := k.GetProposal(ctx, chainID, 0)
	require.True(t, found, "MsgProposalAddAccount should append a new proposal")
	require.Equal(t, creatorIdentity, proposal.ProposalInformation.Creator, "MsgProposalAddAccount should append a new proposal")
	_, ok := proposal.Payload.(*types.Proposal_AddAccountPayload)
	require.True(t, ok, "MsgProposalAddAccount should append a new proposal to add account")

	// The proposal is added to the pending proposals
	pending := k.GetPendingProposals(ctx, chainID)
	require.Contains(t, pending, 0, "MsgProposalAddAccount should append the proposal to the pending pool")

	// The proposal count is incremented
	count := k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(1), count, "MsgProposalAddAccount should increment proposal count")
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
	require.Error(t, err, "MsgProposalAddValidator should fail for non existing chains")

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
	require.Error(t, err, "MsgProposalAddValidator should append a new proposal")

	// Can retrieve the proposal
	proposal, found := k.GetProposal(ctx, chainID, 5)
	require.True(t, found, "MsgProposalAddValidator should append a new proposal")
	require.Equal(t, creatorIdentity, proposal.ProposalInformation.Creator, "MsgProposalAddValidator should append a new proposal")
	_, ok := proposal.Payload.(*types.Proposal_AddValidatorPayload)
	require.True(t, ok, "MsgProposalAddValidator should append a new proposal to add validator")

	// The proposal is added to the pending proposals
	pending := k.GetPendingProposals(ctx, chainID)
	require.Contains(t, pending, 5, "MsgProposalAddValidator should append the proposal to the pending pool")

	// The proposal count is incremented
	count := k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(6), count, "MsgProposalAddValidator should increment proposal count")
}