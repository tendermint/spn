package keeper_test

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

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

	// Create an add validator proposal
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
	require.NotEqual(t, *pendingRes.Proposals[0], *pendingRes.Proposals[1]) // Simple check to ensure all elements are the same

	// PendingProposals fails if the chain doesn't exist
	pendingQuery = types.QueryPendingProposalsRequest{
		ChainID: spnmocks.MockRandomAlphaString(6),
	}
	pendingRes, err = q.PendingProposals(context.Background(), &pendingQuery)
	require.Error(t, err)

	// Can query a specific proposal
	showQuery := types.QueryShowProposalRequest{
		ChainID:    chainID,
		ProposalID: 0,
	}
	showRes, err := q.ShowProposal(context.Background(), &showQuery)
	require.NoError(t, err)
	retrievedPayload1, ok := showRes.Proposal.Payload.(*types.Proposal_AddAccountPayload)
	require.True(t, ok)
	require.True(t, proposal1.Address.Equals(retrievedPayload1.AddAccountPayload.Address))

	// Test with the add validator query
	showQuery = types.QueryShowProposalRequest{
		ChainID:    chainID,
		ProposalID: 1,
	}
	showRes, err = q.ShowProposal(context.Background(), &showQuery)
	require.NoError(t, err)
	retrievedPayload2, ok := showRes.Proposal.Payload.(*types.Proposal_AddValidatorPayload)
	require.True(t, ok)
	require.Equal(t, proposal2.Peer, retrievedPayload2.AddValidatorPayload.Peer)

	// ShowProposal fails if the proposal doesn't exist
	showQuery = types.QueryShowProposalRequest{
		ChainID:    chainID,
		ProposalID: 1000,
	}
	_, err = q.ShowProposal(context.Background(), &showQuery)
	require.Error(t, err)

	// ShowProposal fails if the chain doesn't exist
	showQuery = types.QueryShowProposalRequest{
		ChainID:    spnmocks.MockRandomAlphaString(7),
		ProposalID: 0,
	}
	_, err = q.ShowProposal(context.Background(), &showQuery)
	require.Error(t, err)
}

func TestApprovedProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	// Create a new chain
	coordinator := spnmocks.MockAccAddress()
	chainID := spnmocks.MockRandomAlphaString(5)
	chain := spnmocks.MockChain()
	chain.ChainID = chainID
	chain.Creator, _ = k.IdentityKeeper.GetIdentifier(ctx, coordinator)
	k.SetChain(ctx, *chain)

	// Create and send add account proposals
	for i := 0; i < 10; i++ {
		msg := types.NewMsgProposalAddAccount(
			chainID,
			spnmocks.MockAccAddress(),
			spnmocks.MockProposalAddAccountPayload(),
		)
		h(ctx, msg)
	}

	// Approve half the proposals
	for i := 0; i < 5; i++ {
		msgApprove := types.NewMsgApprove(
			chainID,
			int32(i),
			coordinator,
		)
		h(ctx, msgApprove)
	}

	// Can query approved proposals
	approvedQuery := types.QueryApprovedProposalsRequest{
		ChainID: chainID,
	}
	approvedRes, err := q.ApprovedProposals(context.Background(), &approvedQuery)
	require.NoError(t, err)
	require.Equal(t, 5, len(approvedRes.Proposals))
	require.NotEqual(t, *approvedRes.Proposals[0], *approvedRes.Proposals[1]) // Simple check to ensure all elements are the same
	approvedProposal, _ := k.GetProposal(ctx, chainID, 0)
	require.Equal(t, *approvedRes.Proposals[0], approvedProposal)

	// PendingProposals fails if the chain doesn't exist
	approvedQuery = types.QueryApprovedProposalsRequest{
		ChainID: spnmocks.MockRandomAlphaString(6),
	}
	approvedRes, err = q.ApprovedProposals(context.Background(), &approvedQuery)
	require.Error(t, err)
}

func TestRejectedProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	// Create a new chain
	coordinator := spnmocks.MockAccAddress()
	chainID := spnmocks.MockRandomAlphaString(5)
	chain := spnmocks.MockChain()
	chain.ChainID = chainID
	chain.Creator, _ = k.IdentityKeeper.GetIdentifier(ctx, coordinator)
	chain.Creator, _ = k.IdentityKeeper.GetIdentifier(ctx, coordinator)
	k.SetChain(ctx, *chain)

	// Create and send add account proposals
	for i := 0; i < 10; i++ {
		msg := types.NewMsgProposalAddAccount(
			chainID,
			spnmocks.MockAccAddress(),
			spnmocks.MockProposalAddAccountPayload(),
		)
		h(ctx, msg)
	}

	// Reject half the proposals
	for i := 0; i < 5; i++ {
		msgReject := types.NewMsgReject(
			chainID,
			int32(i),
			coordinator,
		)
		h(ctx, msgReject)
	}

	// Can query approved proposals
	rejectQuery := types.QueryRejectedProposalsRequest{
		ChainID: chainID,
	}
	rejectedRes, err := q.RejectedProposals(context.Background(), &rejectQuery)
	require.NoError(t, err)
	require.Equal(t, 5, len(rejectedRes.Proposals))
	require.NotEqual(t, *rejectedRes.Proposals[0], *rejectedRes.Proposals[1]) // Simple check to ensure all elements are the same
	rejectedProposal, _ := k.GetProposal(ctx, chainID, 0)
	require.Equal(t, *rejectedRes.Proposals[0], rejectedProposal)

	// PendingProposals fails if the chain doesn't exist
	rejectQuery = types.QueryRejectedProposalsRequest{
		ChainID: spnmocks.MockRandomAlphaString(6),
	}
	_, err = q.RejectedProposals(context.Background(), &rejectQuery)
	require.Error(t, err)
}

// TODO: This test could be considered as an integration test -> Move it to the integration test
func TestCurrentGenesis(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)
	cdc := spnmocks.MockCodec()

	chainID := spnmocks.MockRandomAlphaString(5)
	coordinator := spnmocks.MockAccAddress()

	// Create a new chain
	msgChainCreate := types.NewMsgChainCreate(
		chainID,
		coordinator,
		spnmocks.MockRandomAlphaString(10),
		spnmocks.MockRandomAlphaString(10),
		spnmocks.MockGenesis(),
	)
	h(ctx, msgChainCreate)

	// Test with 20 accounts and 10 validators
	for i:=0; i<10; i++ {
		// Add validator payload
		addValidatorpayload := spnmocks.MockProposalAddValidatorPayload()
		msg, _ := addValidatorpayload.GetCreateValidatorMessage()
		valAddress, _ := sdk.ValAddressFromBech32(msg.ValidatorAddress)
		accAddress := sdk.AccAddress(valAddress)

		// Add account payload (for each validator we need an account)
		addAccountPayload := spnmocks.MockProposalAddAccountPayload()
		addAccountPayload.Address = accAddress

		// Send add account proposal
		msgAddAccount := types.NewMsgProposalAddAccount(
			chainID,
			accAddress,
			addAccountPayload,
		)
		_, err := h(ctx, msgAddAccount)
		require.NoError(t, err)

		// Send add validator proposal
		msgAddValidator := types.NewMsgProposalAddValidator(
			chainID,
			accAddress,
			addValidatorpayload,
		)
		_, err = h(ctx, msgAddValidator)
		require.NoError(t, err)
	}
	for i:=0; i<10; i++ {
		addAccountPayload := spnmocks.MockProposalAddAccountPayload()

		// Send add account proposal
		msgAddAccount := types.NewMsgProposalAddAccount(
			chainID,
			addAccountPayload.Address,
			addAccountPayload,
		)
		_, err := h(ctx, msgAddAccount)
		require.NoError(t, err)
	}

	// Approve all proposals
	for i:=0; i<30; i++ {
		msg := types.NewMsgApprove(
			chainID,
			int32(i),
			coordinator,
		)
		_, err := h(ctx, msg)
		require.NoError(t, err)
	}

	// Can retrieve the current genesis with all the approved proposals
	var req types.QueryCurrentGenesisRequest
	req.ChainID = chainID
	res, err := q.CurrentGenesis(context.Background(), &req)
	require.NoError(t, err)

	// Parse the retrieved genesis
	var genesis types.GenesisFile
	genesis = res.Genesis
	genesisDoc, err := genesis.GetGenesisDoc()
	require.NoError(t, err)
	appState, err := genutiltypes.GenesisStateFromGenDoc(genesisDoc)
	require.NoError(t, err)

	// Analyse accounts
	authGenState := authtypes.GetGenesisStateFromAppState(cdc, appState)
	accs, err := authtypes.UnpackAccounts(authGenState.Accounts)
	require.Equal(t, 20, len(accs))

	// Analyse validators
	appGenesisState, err := genutiltypes.GenesisStateFromGenDoc(genesisDoc)
	require.NoError(t, err)
	genesisState := genutiltypes.GetGenesisStateFromAppState(cdc, appGenesisState)
	require.Equal(t, 10, len(genesisState.GenTxs))
}