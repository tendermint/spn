package keeper_test

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

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
		)
		h(ctx, msg)
	}

	// Query the created chains
	var listQuery types.QueryListChainsRequest
	listChainsRes, err := q.ListChains(context.Background(), &listQuery)
	require.NoError(t, err)
	require.Equal(t, 10, len(listChainsRes.Chains))
	for _, chainID := range chainIDs {
		chain, found := k.GetChain(ctx, chainID)
		require.True(t, found)
		require.Contains(t, listChainsRes.Chains, &chain)
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

func TestListProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	coordinator := spnmocks.MockAccAddress()
	coordinatorIdentity, _ := k.IdentityKeeper.GetIdentifier(ctx, coordinator)

	// Create a new chain
	chainID := spnmocks.MockRandomAlphaString(5)
	chain := spnmocks.MockChain()
	chain.Creator = coordinatorIdentity
	chain.ChainID = chainID
	k.SetChain(ctx, *chain)

	// Array of pending add account
	var pendingAddAccount []*types.Proposal
	for i:=0; i<5; i++ {
		payloadAddAccount := spnmocks.MockProposalAddAccountPayload()
		msgAddAccount := types.NewMsgProposalAddAccount(
			chainID,
			spnmocks.MockAccAddress(),
			payloadAddAccount,
		)
		res, err := h(ctx, msgAddAccount)
		require.NoError(t, err)
		proposal := types.UnmarshalProposal(k.GetCodec(), res.Data)
		pendingAddAccount = append(pendingAddAccount, &proposal)
	}

	// Array of pending add validator
	var pendingAddValidator []*types.Proposal
	for i:=0; i<5; i++ {
		payloadAddValidator := spnmocks.MockProposalAddValidatorPayload()
		msgAddValidator := types.NewMsgProposalAddValidator(
			chainID,
			spnmocks.MockAccAddress(),
			payloadAddValidator,
		)
		res, err := h(ctx, msgAddValidator)
		require.NoError(t, err)
		proposal := types.UnmarshalProposal(k.GetCodec(), res.Data)
		pendingAddValidator = append(pendingAddValidator, &proposal)
	}

	// Array of rejected add account
	var rejectedAddAccount []*types.Proposal
	for i:=0; i<5; i++ {
		payloadAddAccount := spnmocks.MockProposalAddAccountPayload()
		msgAddAccount := types.NewMsgProposalAddAccount(
			chainID,
			spnmocks.MockAccAddress(),
			payloadAddAccount,
		)
		res, err := h(ctx, msgAddAccount)
		require.NoError(t, err)
		proposal := types.UnmarshalProposal(k.GetCodec(), res.Data)

		msgReject := types.NewMsgReject(
			chainID,
			proposal.ProposalInformation.ProposalID,
			coordinator,
		)
		res, err = h(ctx, msgReject)
		require.NoError(t, err)
		proposal = types.UnmarshalProposal(k.GetCodec(), res.Data)

		rejectedAddAccount = append(rejectedAddAccount, &proposal)
	}

	// Array of rejected add validator
	var rejectedAddValidator []*types.Proposal
	for i:=0; i<5; i++ {
		payloadAddValidator := spnmocks.MockProposalAddValidatorPayload()
		msgAddValidator := types.NewMsgProposalAddValidator(
			chainID,
			spnmocks.MockAccAddress(),
			payloadAddValidator,
		)
		res, err := h(ctx, msgAddValidator)
		require.NoError(t, err)
		proposal := types.UnmarshalProposal(k.GetCodec(), res.Data)

		msgReject := types.NewMsgReject(
			chainID,
			proposal.ProposalInformation.ProposalID,
			coordinator,
		)
		res, err = h(ctx, msgReject)
		require.NoError(t, err)
		proposal = types.UnmarshalProposal(k.GetCodec(), res.Data)

		rejectedAddValidator = append(rejectedAddValidator, &proposal)
	}

	// Array of approved add account
	var approvedAddAccount []*types.Proposal
	var approvedAddress []sdk.AccAddress
	for i:=0; i<5; i++ {
		accountAddress := spnmocks.MockAccAddress()
		approvedAddress = append(approvedAddress, accountAddress)
		payloadAddAccount := spnmocks.MockProposalAddAccountPayload()
		payloadAddAccount.Address = accountAddress
		msgAddAccount := types.NewMsgProposalAddAccount(
			chainID,
			spnmocks.MockAccAddress(),
			payloadAddAccount,
		)
		res, err := h(ctx, msgAddAccount)
		require.NoError(t, err)
		proposal := types.UnmarshalProposal(k.GetCodec(), res.Data)

		msgApprove := types.NewMsgApprove(
			chainID,
			proposal.ProposalInformation.ProposalID,
			coordinator,
		)
		res, err = h(ctx, msgApprove)
		require.NoError(t, err)
		proposal = types.UnmarshalProposal(k.GetCodec(), res.Data)

		approvedAddAccount = append(approvedAddAccount, &proposal)
	}

	// Array of approved add validator
	var approvedAddValidator []*types.Proposal
	for i:=0; i<5; i++ {
		payloadAddValidator := spnmocks.MockProposalAddValidatorPayload()
		payloadAddValidator.ValidatorAddress = sdk.ValAddress(approvedAddress[i]) // Need an existing account to be approved
		msgAddValidator := types.NewMsgProposalAddValidator(
			chainID,
			spnmocks.MockAccAddress(),
			payloadAddValidator,
		)
		res, err := h(ctx, msgAddValidator)
		require.NoError(t, err)
		proposal := types.UnmarshalProposal(k.GetCodec(), res.Data)

		msgApprove := types.NewMsgApprove(
			chainID,
			proposal.ProposalInformation.ProposalID,
			coordinator,
		)
		res, err = h(ctx, msgApprove)
		require.NoError(t, err)
		proposal = types.UnmarshalProposal(k.GetCodec(), res.Data)

		approvedAddValidator = append(approvedAddValidator, &proposal)
	}

	// Can fetch all proposal
	var req types.QueryListProposalsRequest
	req.ChainID = chainID
	req.Type = types.ProposalType_ANY_TYPE
	req.Status = types.ProposalStatus_ANY_STATUS
	fetched, err := q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.Subset(t, fetched.Proposals, pendingAddAccount)
	require.Subset(t, fetched.Proposals, pendingAddValidator)
	require.Subset(t, fetched.Proposals, rejectedAddAccount)
	require.Subset(t, fetched.Proposals, rejectedAddValidator)
	require.Subset(t, fetched.Proposals, approvedAddAccount)
	require.Subset(t, fetched.Proposals, approvedAddValidator)

	// Can fetch pending proposals
	req.ChainID = chainID
	req.Type = types.ProposalType_ANY_TYPE
	req.Status = types.ProposalStatus_PENDING
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.Subset(t, fetched.Proposals, pendingAddAccount)
	require.Subset(t, fetched.Proposals, pendingAddValidator)
	require.NotSubset(t, fetched.Proposals, rejectedAddAccount)
	require.NotSubset(t, fetched.Proposals, rejectedAddValidator)
	require.NotSubset(t, fetched.Proposals, approvedAddAccount)
	require.NotSubset(t, fetched.Proposals, approvedAddValidator)

	// Can fetch rejected proposals
	req.ChainID = chainID
	req.Type = types.ProposalType_ANY_TYPE
	req.Status = types.ProposalStatus_REJECTED
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.NotSubset(t, fetched.Proposals, pendingAddAccount)
	require.NotSubset(t, fetched.Proposals, pendingAddValidator)
	require.Subset(t, fetched.Proposals, rejectedAddAccount)
	require.Subset(t, fetched.Proposals, rejectedAddValidator)
	require.NotSubset(t, fetched.Proposals, approvedAddAccount)
	require.NotSubset(t, fetched.Proposals, approvedAddValidator)

	// Can fetch approved proposals
	req.ChainID = chainID
	req.Type = types.ProposalType_ANY_TYPE
	req.Status = types.ProposalStatus_APPROVED
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.NotSubset(t, fetched.Proposals, pendingAddAccount)
	require.NotSubset(t, fetched.Proposals, pendingAddValidator)
	require.NotSubset(t, fetched.Proposals, rejectedAddAccount)
	require.NotSubset(t, fetched.Proposals, rejectedAddValidator)
	require.Subset(t, fetched.Proposals, approvedAddAccount)
	require.Subset(t, fetched.Proposals, approvedAddValidator)

	// Can fetch add account proposals
	req.ChainID = chainID
	req.Type = types.ProposalType_ADD_ACCOUNT
	req.Status = types.ProposalStatus_ANY_STATUS
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.Subset(t, fetched.Proposals, pendingAddAccount)
	require.NotSubset(t, fetched.Proposals, pendingAddValidator)
	require.Subset(t, fetched.Proposals, rejectedAddAccount)
	require.NotSubset(t, fetched.Proposals, rejectedAddValidator)
	require.Subset(t, fetched.Proposals, approvedAddAccount)
	require.NotSubset(t, fetched.Proposals, approvedAddValidator)

	// Can fetch add validator proposals
	req.ChainID = chainID
	req.Type = types.ProposalType_ADD_VALIDATOR
	req.Status = types.ProposalStatus_ANY_STATUS
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.NotSubset(t, fetched.Proposals, pendingAddAccount)
	require.Subset(t, fetched.Proposals, pendingAddValidator)
	require.NotSubset(t, fetched.Proposals, rejectedAddAccount)
	require.Subset(t, fetched.Proposals, rejectedAddValidator)
	require.NotSubset(t, fetched.Proposals, approvedAddAccount)
	require.Subset(t, fetched.Proposals, approvedAddValidator)

	// Can fetch pending add account
	req.ChainID = chainID
	req.Type = types.ProposalType_ADD_ACCOUNT
	req.Status = types.ProposalStatus_PENDING
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, pendingAddAccount, fetched.Proposals)

	// Can fetch rejected add account
	req.ChainID = chainID
	req.Type = types.ProposalType_ADD_ACCOUNT
	req.Status = types.ProposalStatus_REJECTED
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, rejectedAddAccount, fetched.Proposals)

	// Can fetch approved add account
	req.ChainID = chainID
	req.Type = types.ProposalType_ADD_ACCOUNT
	req.Status = types.ProposalStatus_APPROVED
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, approvedAddAccount, fetched.Proposals)

	// Can fetch pending add validator
	req.ChainID = chainID
	req.Type = types.ProposalType_ADD_VALIDATOR
	req.Status = types.ProposalStatus_PENDING
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, pendingAddValidator, fetched.Proposals)

	// Can fetch rejected add validator
	req.ChainID = chainID
	req.Type = types.ProposalType_ADD_VALIDATOR
	req.Status = types.ProposalStatus_REJECTED
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, rejectedAddValidator, fetched.Proposals)

	// Can fetch approved add validator
	req.ChainID = chainID
	req.Type = types.ProposalType_ADD_VALIDATOR
	req.Status = types.ProposalStatus_APPROVED
	fetched, err = q.ListProposals(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, approvedAddValidator, fetched.Proposals)
}

func TestLaunchInformation(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)
	_ = spnmocks.MockCodec()

	chainID := spnmocks.MockRandomAlphaString(5)
	coordinator := spnmocks.MockAccAddress()

	// Create a new chain
	msgChainCreate := types.NewMsgChainCreate(
		chainID,
		coordinator,
		spnmocks.MockRandomAlphaString(10),
		spnmocks.MockRandomAlphaString(10),
	)
	h(ctx, msgChainCreate)

	// Test with 20 accounts and 10 validators
	for i:=0; i<10; i++ {
		// Add validator payload
		addValidatorpayload := spnmocks.MockProposalAddValidatorPayload()
		valAddress := addValidatorpayload.ValidatorAddress
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
	var req types.QueryLaunchInformationRequest
	req.ChainID = chainID
	launchInformation, err := q.LaunchInformation(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, 20, len(launchInformation.Accounts))
	require.Equal(t, 10, len(launchInformation.GenTxs))
	require.Equal(t, 10, len(launchInformation.Peers))
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
	_, err := h(ctx, msg)
	require.NoError(t, err)

	// Create an add validator proposal
	proposal2 := spnmocks.MockProposalAddValidatorPayload()
	msg2 := types.NewMsgProposalAddValidator(
		chainID,
		spnmocks.MockAccAddress(),
		proposal2,
	)
	_, err = h(ctx, msg2)
	require.NoError(t, err)

	// Create other proposal to test pending proposal command
	for i := 0; i < 8; i++ {
		_, err := h(ctx, msg)
		require.NoError(t, err)
	}

	proposals, err := k.PendingProposals(ctx, chainID)
	require.NoError(t, err)
	require.Equal(t, 10, len(proposals))
	require.NotEqual(t, proposals[0].Payload, proposals[1].Payload)

	// PendingProposals fails if the chain doesn't exist
	_, err = k.PendingProposals(ctx, spnmocks.MockRandomAlphaString(6))
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
		_, err := h(ctx, msg)
		require.NoError(t, err)
	}

	// Approve half the proposals
	for i := 0; i < 5; i++ {
		msgApprove := types.NewMsgApprove(
			chainID,
			int32(i),
			coordinator,
		)
		_, err := h(ctx, msgApprove)
		require.NoError(t, err)
	}

	// Can query approved proposals
	proposals, err := k.ApprovedProposals(ctx, chainID)
	require.NoError(t, err)
	require.Equal(t, 5, len(proposals))
	require.NotEqual(t, proposals[0], proposals[1])
	approvedProposal, _ := k.GetProposal(ctx, chainID, 0)
	require.Equal(t, approvedProposal, proposals[0])

	// ApprovedProposals fails if the chain doesn't exist
	_, err = k.ApprovedProposals(ctx,  spnmocks.MockRandomAlphaString(6))
	require.Error(t, err)
}

func TestRejectedProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)

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
		_, err := h(ctx, msg)
		require.NoError(t, err)
	}

	// Reject half the proposals
	for i := 0; i < 5; i++ {
		msgReject := types.NewMsgReject(
			chainID,
			int32(i),
			coordinator,
		)
		_, err := h(ctx, msgReject)
		require.NoError(t, err)
	}

	// Can query approved proposals
	proposals, err := k.RejectedProposals(ctx, chainID)
	require.NoError(t, err)
	require.Equal(t, 5, len(proposals))
	require.NotEqual(t, proposals[0], proposals[1])
	rejectedProposal, _ := k.GetProposal(ctx, chainID, 0)
	require.Equal(t, rejectedProposal, proposals[0])

	// PendingProposals fails if the chain doesn't exist
	_, err = k.RejectedProposals(ctx, spnmocks.MockRandomAlphaString(6))
	require.Error(t, err)
}
