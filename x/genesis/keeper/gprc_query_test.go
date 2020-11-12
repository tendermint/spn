package keeper_test

import (
	"context"

	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/genesis"
	"github.com/tendermint/spn/x/genesis/types"
	"testing"
)

func TestListChains(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	// Create several chains through messages
	var chainIDs []string
	var sourceURLs []string
	for i := 0; i < 10; i++ {
		chainIDs = append(chainIDs, spnmocks.MockRandomAlphaString(5))
		sourceURLs = append(sourceURLs, spnmocks.MockRandomAlphaString(5))
	}
	for i := range chainIDs {
		msg := types.NewMsgChainCreate(
			chainIDs[i],
			spnmocks.MockAccAddress(),
			sourceURLs[i],
			spnmocks.MockRandomAlphaString(5),
			spnmocks.MockGenesis(),
		)
		h(ctx, msg)
	}

	// Query the created chains
	var listQuery types.QueryListChainsRequest
	listChainsRes, err := q.ListChains(context.Background(), &listQuery)
	require.NoError(t, err)
	require.Equal(t, 10, len(listChainsRes.ChainIDs))
	for _, chainID := range chainIDs {
		require.Contains(t, listChainsRes.ChainIDs, chainID)
	}

	// Query a specific chain
	showQuery := types.QueryShowChainRequest{
		ChainID: chainIDs[5],
	}
	showChainRes, err := q.ShowChain(context.Background(), &showQuery)
	require.NoError(t, err)
	require.Equal(t, sourceURLs[5], showChainRes.Chain.SourceURL)

	// Query on a non existing chain should fail
	showQuery = types.QueryShowChainRequest{
		ChainID: spnmocks.MockRandomAlphaString(10),
	}
	_, err = q.ShowChain(context.Background(), &showQuery)
	require.Error(t, err)
}

func TestPendingProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	// Create a new chain
	chainID := spnmocks.MockRandomAlphaString(5)
	chain := spnmocks.MockChain()
	chain.ChainID = chainID
	k.SetChain(ctx, *chain)

	// Create an add account proposal
	proposal1 := spnmocks.MockProposalAddAccountPayload()
	msg := types.NewMsgProposalAddAccount(
		chainID,
		spnmocks.MockAccAddress(),
		proposal1,
	)
	h(ctx, msg)

	// Create an add validator porposal
	proposal2 := spnmocks.MockProposalAddValidatorPayload()
	msg2 := types.NewMsgProposalAddValidator(
		chainID,
		spnmocks.MockAccAddress(),
		proposal2,
	)
	h(ctx, msg2)

	// Create other proposal to test pending proposal command
	for i := 0; i < 8; i++ {
		h(ctx, msg)
	}

	// Can query pending proposals
	pendingQuery := types.QueryPendingProposalsRequest{
		ChainID: chainID,
	}
	pendingRes, err := q.PendingProposals(context.Background(), &pendingQuery)
	require.NoError(t, err)
	require.Equal(t, 10, len(pendingRes.Proposals))
	require.NotEqual(t, t, *pendingRes.Proposals[0], t, *pendingRes.Proposals[1]) // Simple check to ensure all elements are the same

	// PendingProposals fails if the chain doesn't exist
	pendingQuery = types.QueryPendingProposalsRequest{
		ChainID: spnmocks.MockRandomAlphaString(6),
	}
	pendingRes, err = q.PendingProposals(context.Background(), &pendingQuery)
	require.Error(t, err)

	// Can query a specific proposal
	showQuery := types.QueryShowProposalRequest{
		ChainID: chainID,
		ProposalID: 0,
	}
	showRes, err := q.ShowProposal(context.Background(), &showQuery)
	require.NoError(t, err)
	retrievedPayload1, ok := showRes.Proposal.Payload.(*types.Proposal_AddAccountPayload)
	require.True(t, ok)
	require.Equal(t, proposal1, retrievedPayload1.AddAccountPayload)

	// Test with the add validator query
	showQuery = types.QueryShowProposalRequest{
		ChainID: chainID,
		ProposalID: 1,
	}
	showRes, err = q.ShowProposal(context.Background(), &showQuery)
	require.NoError(t, err)
	retrievedPayload2, ok := showRes.Proposal.Payload.(*types.Proposal_AddValidatorPayload)
	require.True(t, ok)
	require.Equal(t, proposal2, retrievedPayload2.AddValidatorPayload)

	// ShowProposal fails if the proposal doesn't exist
	showQuery = types.QueryShowProposalRequest{
		ChainID: chainID,
		ProposalID: 1000,
	}
	_, err = q.ShowProposal(context.Background(), &showQuery)
	require.Error(t, err)

	// ShowProposal fails if the chain doesn't exist
	showQuery = types.QueryShowProposalRequest{
		ChainID: spnmocks.MockRandomAlphaString(7),
		ProposalID: 0,
	}
	_, err = q.ShowProposal(context.Background(), &showQuery)
	require.Error(t, err)
}
