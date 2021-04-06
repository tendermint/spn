package keeper_test

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sort"

	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/launch"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestListChains(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := launch.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	// Create several chains through messages
	var chainIDs []string
	for i := 0; i < 10; i++ {
		chainIDs = append(chainIDs, "AAA"+spnmocks.MockRandomAlphaString(5))
	}
	for i := 0; i < 10; i++ {
		chainIDs = append(chainIDs, "BBB"+spnmocks.MockRandomAlphaString(5))
	}
	for i := range chainIDs {
		msg := types.NewMsgChainCreate(
			chainIDs[i],
			spnmocks.MockAccAddress(),
			spnmocks.MockRandomAlphaString(5),
			spnmocks.MockRandomAlphaString(5),
			"",
			"",
		)
		_, err := h(ctx, msg)
		require.NoError(t, err)
	}

	// ListChains lists the chains and can filter with a prefix
	var listQuery types.QueryListChainsRequest
	listChainsRes, err := q.ListChains(context.Background(), &listQuery)
	require.NoError(t, err)

	listQuery.Prefix = "AAA"
	listChainsResAAA, err := q.ListChains(context.Background(), &listQuery)
	require.NoError(t, err)

	listQuery.Prefix = "BBB"
	listChainsResBBB, err := q.ListChains(context.Background(), &listQuery)
	require.NoError(t, err)

	listQuery.Prefix = "CCC"
	listChainsResCCC, err := q.ListChains(context.Background(), &listQuery)
	require.NoError(t, err)

	require.Equal(t, 20, len(listChainsRes.Chains))
	require.Equal(t, 10, len(listChainsResAAA.Chains))
	require.Equal(t, 10, len(listChainsResBBB.Chains))
	require.Equal(t, 0, len(listChainsResCCC.Chains))

	// Retrieve AAA prefixed chains
	for _, chainID := range chainIDs[0:10] {
		chain, found := k.GetChain(ctx, chainID)
		require.True(t, found)
		require.Contains(t, listChainsRes.Chains, &chain)
		require.Contains(t, listChainsResAAA.Chains, &chain)
		require.NotContains(t, listChainsResBBB.Chains, &chain)
	}
	// Retrieve BBB prefixed chains
	for _, chainID := range chainIDs[10:20] {
		chain, found := k.GetChain(ctx, chainID)
		require.True(t, found)
		require.Contains(t, listChainsRes.Chains, &chain)
		require.NotContains(t, listChainsResAAA.Chains, &chain)
		require.Contains(t, listChainsResBBB.Chains, &chain)
	}
}

func TestShowChain(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := launch.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	// Create a chain
	chainID := spnmocks.MockRandomAlphaString(5)
	msg := types.NewMsgChainCreate(
		chainID,
		spnmocks.MockAccAddress(),
		spnmocks.MockRandomAlphaString(5),
		spnmocks.MockRandomAlphaString(5),
		"",
		"",
	)
	res, err := h(ctx, msg)
	require.NoError(t, err)
	chain := types.UnmarshalChain(k.GetCodec(), res.Data)

	// Query a specific chain
	showQuery := types.QueryShowChainRequest{
		ChainID: chainID,
	}
	showChainRes, err := q.ShowChain(context.Background(), &showQuery)
	require.NoError(t, err)
	require.Equal(t, chain, *showChainRes.Chain)

	// Query on a non existing chain should fail
	showQuery = types.QueryShowChainRequest{
		ChainID: spnmocks.MockRandomAlphaString(10),
	}
	_, err = q.ShowChain(context.Background(), &showQuery)
	require.Error(t, err)

}

func TestProposalCount(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := launch.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	// Fails if the chain doesn't exist
	countQuery := types.QueryProposalCountRequest{
		ChainID: spnmocks.MockRandomAlphaString(5),
	}
	_, err := q.ProposalCount(context.Background(), &countQuery)
	require.Error(t, err)

	// Create a chain
	chainID := spnmocks.MockRandomAlphaString(5)
	msg := types.NewMsgChainCreate(
		chainID,
		spnmocks.MockAccAddress(),
		spnmocks.MockRandomAlphaString(5),
		spnmocks.MockRandomAlphaString(5),
		"",
		"",
	)
	_, err = h(ctx, msg)
	require.NoError(t, err)

	// 0 proposal
	countQuery = types.QueryProposalCountRequest{
		ChainID: chainID,
	}
	count, err := q.ProposalCount(context.Background(), &countQuery)
	require.NoError(t, err)
	require.Equal(t, int32(0), count.Count)

	// 3 proposals
	msgAddAccount := types.NewMsgProposalAddAccount(
		chainID,
		spnmocks.MockAccAddress(),
		spnmocks.MockProposalAddAccountPayload(),
	)
	_, err = h(ctx, msgAddAccount)
	require.NoError(t, err)
	msgAddAccount = types.NewMsgProposalAddAccount(
		chainID,
		spnmocks.MockAccAddress(),
		spnmocks.MockProposalAddAccountPayload(),
	)
	_, err = h(ctx, msgAddAccount)
	require.NoError(t, err)
	msgAddAccount = types.NewMsgProposalAddAccount(
		chainID,
		spnmocks.MockAccAddress(),
		spnmocks.MockProposalAddAccountPayload(),
	)
	_, err = h(ctx, msgAddAccount)
	require.NoError(t, err)
	countQuery = types.QueryProposalCountRequest{
		ChainID: chainID,
	}
	count, err = q.ProposalCount(context.Background(), &countQuery)
	require.NoError(t, err)
	require.Equal(t, int32(3), count.Count)
}

func TestShowProposal(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := launch.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	// Create a new chain
	chainID := spnmocks.MockRandomAlphaString(5)
	chain := spnmocks.MockChain()
	chain.ChainID = chainID
	k.SetChain(ctx, *chain)

	// Create an add account proposal
	payload1 := spnmocks.MockProposalAddAccountPayload()
	msg := types.NewMsgProposalAddAccount(
		chainID,
		spnmocks.MockAccAddress(),
		payload1,
	)
	res, err := h(ctx, msg)
	require.NoError(t, err)
	proposal1 := types.UnmarshalProposal(k.GetCodec(), res.Data)

	// Create an add validator proposal
	payload2 := spnmocks.MockProposalAddValidatorPayload()
	msg2 := types.NewMsgProposalAddValidator(
		chainID,
		spnmocks.MockAccAddress(),
		payload2,
	)
	res, err = h(ctx, msg2)
	require.NoError(t, err)
	proposal2 := types.UnmarshalProposal(k.GetCodec(), res.Data)

	// Can query a specific proposal
	showQuery := types.QueryShowProposalRequest{
		ChainID:    chainID,
		ProposalID: 0,
	}
	showRes, err := q.ShowProposal(context.Background(), &showQuery)
	require.NoError(t, err)
	_, ok := showRes.Proposal.Payload.(*types.Proposal_AddAccountPayload)
	require.True(t, ok)
	require.Equal(t, &proposal1, showRes.Proposal)

	// Test with the add validator query
	showQuery = types.QueryShowProposalRequest{
		ChainID:    chainID,
		ProposalID: 1,
	}
	showRes, err = q.ShowProposal(context.Background(), &showQuery)
	require.NoError(t, err)
	_, ok = showRes.Proposal.Payload.(*types.Proposal_AddValidatorPayload)
	require.True(t, ok)
	require.Equal(t, &proposal2, showRes.Proposal)

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

func TestListProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := launch.NewHandler(*k)
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
	for i := 0; i < 5; i++ {
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
	for i := 0; i < 5; i++ {
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
	for i := 0; i < 5; i++ {
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
	for i := 0; i < 5; i++ {
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
	for i := 0; i < 5; i++ {
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
	for i := 0; i < 5; i++ {
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
	h := launch.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	chainID := spnmocks.MockRandomAlphaString(5)
	coordinator := spnmocks.MockAccAddress()

	// Create a new chain
	msgChainCreate := types.NewMsgChainCreate(
		chainID,
		coordinator,
		spnmocks.MockRandomAlphaString(10),
		spnmocks.MockRandomAlphaString(10),
		"",
		"",
	)
	h(ctx, msgChainCreate)

	var accounts []*types.ProposalAddAccountPayload
	var gentxs [][]byte
	var peers []string

	// Test with 20 accounts and 10 validators
	for i := 0; i < 10; i++ {
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

		accounts = append(accounts, addAccountPayload)

		// Send add validator proposal
		msgAddValidator := types.NewMsgProposalAddValidator(
			chainID,
			accAddress,
			addValidatorpayload,
		)
		_, err = h(ctx, msgAddValidator)
		require.NoError(t, err)

		gentxs = append(gentxs, addValidatorpayload.GenTx)
		peers = append(peers, addValidatorpayload.Peer)
	}
	for i := 0; i < 10; i++ {
		addAccountPayload := spnmocks.MockProposalAddAccountPayload()

		// Send add account proposal
		msgAddAccount := types.NewMsgProposalAddAccount(
			chainID,
			addAccountPayload.Address,
			addAccountPayload,
		)
		_, err := h(ctx, msgAddAccount)
		require.NoError(t, err)

		accounts = append(accounts, addAccountPayload)
	}

	// Approve all proposals
	for i := 0; i < 30; i++ {
		msg := types.NewMsgApprove(
			chainID,
			int32(i),
			coordinator,
		)
		_, err := h(ctx, msg)
		require.NoError(t, err)
	}

	// Can retrieve the current launch with all the approved proposals
	var req types.QueryLaunchInformationRequest
	req.ChainID = chainID
	res, err := q.LaunchInformation(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, 20, len(res.LaunchInformation.Accounts))
	require.Equal(t, 10, len(res.LaunchInformation.GenTxs))
	require.Equal(t, 10, len(res.LaunchInformation.Peers))
	require.Equal(t, accounts, res.LaunchInformation.Accounts)
	require.Equal(t, gentxs, res.LaunchInformation.GenTxs)
	require.Equal(t, peers, res.LaunchInformation.Peers)
}

func TestSimulatedLaunchInformation(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := launch.NewHandler(*k)
	q := spnmocks.MockGenesisQueryClient(ctx, k)

	chainID := spnmocks.MockRandomAlphaString(5)
	coordinator := spnmocks.MockAccAddress()

	// Create a new chain
	msgChainCreate := types.NewMsgChainCreate(
		chainID,
		coordinator,
		spnmocks.MockRandomAlphaString(10),
		spnmocks.MockRandomAlphaString(10),
		"",
		"",
	)
	h(ctx, msgChainCreate)

	var accounts []*types.ProposalAddAccountPayload

	// Send 6 add account proposals
	for i := 0; i < 6; i++ {
		addAccountPayload := spnmocks.MockProposalAddAccountPayload()
		msgAddAccount := types.NewMsgProposalAddAccount(
			chainID,
			addAccountPayload.Address,
			addAccountPayload,
		)
		_, err := h(ctx, msgAddAccount)
		require.NoError(t, err)
		accounts = append(accounts, addAccountPayload)
	}
	// Approve 3 of them
	for i := 0; i < 3; i++ {
		msg := types.NewMsgApprove(
			chainID,
			int32(i),
			coordinator,
		)
		_, err := h(ctx, msg)
		require.NoError(t, err)
	}
	// Send an add validator proposal
	addValidatorPayload := spnmocks.MockProposalAddValidatorPayload()
	msgAddValidator := types.NewMsgProposalAddValidator(
		chainID,
		spnmocks.MockAccAddress(),
		addValidatorPayload,
	)
	_, err := h(ctx, msgAddValidator)
	require.NoError(t, err)

	// SimulatedLaunchInformation should contains the proposal to test
	var req types.QuerySimulatedLaunchInformationRequest
	req.ChainID = chainID
	req.ProposalIDs = []int32{int32(3)}
	res, err := q.SimulatedLaunchInformation(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, 4, len(res.LaunchInformation.Accounts))
	require.Equal(t, accounts[0:4], res.LaunchInformation.Accounts)

	// SimulatedLaunchInformation can test a add validator proposal
	req.ChainID = chainID
	req.ProposalIDs = []int32{int32(6)}
	res, err = q.SimulatedLaunchInformation(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, addValidatorPayload.GenTx, res.LaunchInformation.GenTxs[0])
	require.Equal(t, addValidatorPayload.Peer, res.LaunchInformation.Peers[0])

	// Fails if the proposal to test doesn't exist
	req.ChainID = chainID
	req.ProposalIDs = []int32{int32(10)}
	_, err = q.SimulatedLaunchInformation(context.Background(), &req)
	require.Error(t, err)

	// Fails if the proposal to test is not pending
	req.ChainID = chainID
	req.ProposalIDs = []int32{int32(0)}
	_, err = q.SimulatedLaunchInformation(context.Background(), &req)
	require.Error(t, err)

	// Fails if the chain doesn't exist
	req.ChainID = spnmocks.MockRandomAlphaString(5)
	req.ProposalIDs = []int32{int32(3)}
	_, err = q.SimulatedLaunchInformation(context.Background(), &req)
	require.Error(t, err)

	// Allows to test several proposals
	req.ChainID = chainID
	req.ProposalIDs = []int32{int32(3), int32(4), int32(5)}
	res, err = q.SimulatedLaunchInformation(context.Background(), &req)
	require.NoError(t, err)
	require.Equal(t, 6, len(res.LaunchInformation.Accounts))
	require.Equal(t, accounts, res.LaunchInformation.Accounts)

	// Fails if a proposal to test appears twice
	req.ChainID = chainID
	req.ProposalIDs = []int32{int32(3), int32(3)}
	_, err = q.SimulatedLaunchInformation(context.Background(), &req)
	require.Error(t, err)
}

func TestPendingProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := launch.NewHandler(*k)

	coordinator := spnmocks.MockAccAddress()
	coordinatorIdentity, _ := k.IdentityKeeper.GetIdentifier(ctx, coordinator)

	// Create a new chain
	chainID := spnmocks.MockRandomAlphaString(5)
	chain := spnmocks.MockChain()
	chain.Creator = coordinatorIdentity
	chain.ChainID = chainID
	k.SetChain(ctx, *chain)

	// Create an add account proposal
	payload1 := spnmocks.MockProposalAddAccountPayload()
	msg := types.NewMsgProposalAddAccount(
		chainID,
		spnmocks.MockAccAddress(),
		payload1,
	)
	_, err := h(ctx, msg)
	require.NoError(t, err)

	// Create an add validator proposal
	payload2 := spnmocks.MockProposalAddValidatorPayload()
	msg2 := types.NewMsgProposalAddValidator(
		chainID,
		spnmocks.MockAccAddress(),
		payload2,
	)
	_, err = h(ctx, msg2)
	require.NoError(t, err)

	// Create other proposal to test pending proposal command
	for i := 0; i < 8; i++ {
		msg = types.NewMsgProposalAddAccount(
			chainID,
			spnmocks.MockAccAddress(),
			spnmocks.MockProposalAddAccountPayload(),
		)
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

	// Let's approve 2 add accounts proposals inside the pending proposals
	msgApprove := types.NewMsgApprove(
		chainID,
		int32(2),
		coordinator,
	)
	_, err = h(ctx, msgApprove)
	require.NoError(t, err)
	msgApprove = types.NewMsgApprove(
		chainID,
		int32(3),
		coordinator,
	)
	_, err = h(ctx, msgApprove)
	require.NoError(t, err)

	// The result of pending proposals should be sorted
	proposals, err = k.PendingProposals(ctx, chainID)
	require.NoError(t, err)
	require.True(t, sort.SliceIsSorted(proposals, func(i, j int) bool {
		return proposals[i].ProposalInformation.ProposalID < proposals[j].ProposalInformation.ProposalID
	}))
}

func TestApprovedProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := launch.NewHandler(*k)

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
	_, err = k.ApprovedProposals(ctx, spnmocks.MockRandomAlphaString(6))
	require.Error(t, err)
}

func TestRejectedProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := launch.NewHandler(*k)

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
