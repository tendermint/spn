package genesis_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	// A chain can be create
	msg := types.NewMsgChainCreate(
		chainID,
		creator,
		sourceURL,
		sourceHash,
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
	)
	_, err = h(ctx, msg)
	require.Error(t, err)
}

func TestHandleMsgReject(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	chainID := spnmocks.MockRandomAlphaString(5)
	coordinator := spnmocks.MockAccAddress()

	// Prevent rejecting in a non existing chain
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
		)
	_, err = h(ctx, msgChainCreate)
	require.NoError(t, err)

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
	_, err = h(ctx, msgProposal)
	require.NoError(t, err)

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
	_, err = h(ctx, msgProposal)
	require.NoError(t, err)
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

func TestHandleMsgApprove(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	chainID := spnmocks.MockRandomAlphaString(5)
	coordinator := spnmocks.MockAccAddress()

	// Prevent approving in a non existing chain
	msg := types.NewMsgApprove(
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
	)
	_, err = h(ctx, msgChainCreate)
	require.NoError(t, err)

	// Prevent approving a non existing proposal
	msg = types.NewMsgApprove(
		chainID,
		0,
		coordinator,
	)
	_, err = h(ctx, msg)
	require.Error(t, err)

	// Create a proposal
	proposalCreator := spnmocks.MockAccAddress()
	addAccountPayload := spnmocks.MockProposalAddAccountPayload()
	msgProposal := types.NewMsgProposalAddAccount(
		chainID,
		proposalCreator,
		addAccountPayload,
	)
	_, err = h(ctx, msgProposal)
	require.NoError(t, err)

	// Prevent an address other than the coordinator to approve the proposal
	msg = types.NewMsgApprove(
		chainID,
		0,
		spnmocks.MockAccAddress(),
	)
	_, err = h(ctx, msg)
	require.Error(t, err)
	msg = types.NewMsgApprove(
		chainID,
		0,
		proposalCreator,
	)
	_, err = h(ctx, msg)
	require.Error(t, err)

	// The coordinator creator can approve the proposal
	msg = types.NewMsgApprove(
		chainID,
		0,
		coordinator,
	)
	_, err = h(ctx, msg)
	require.NoError(t, err)

	// The proposal is approved
	proposal, _ := k.GetProposal(ctx, chainID, 0)
	require.Equal(t, types.ProposalState_APPROVED, proposal.ProposalState.Status)

	// The proposal is not in pending pool
	pending := k.GetPendingProposals(ctx, chainID)
	require.NotContains(t, pending.ProposalIDs, int32(0))

	// The proposal is in approved pool
	approved := k.GetApprovedProposals(ctx, chainID)
	require.Contains(t, approved.ProposalIDs, int32(0))

	// The account address is set in the store
	accountAddressSet := k.IsAccountSet(ctx, chainID, addAccountPayload.Address)
	require.True(t, accountAddressSet)

	// Prevent approving an already approved proposal
	msg = types.NewMsgApprove(
		chainID,
		0,
		coordinator,
	)
	_, err = h(ctx, msg)
	require.Error(t, err)

	// Prevent approving a proposal with an account already in the genesis
	_, err = h(ctx, msgProposal)
	require.NoError(t, err)
	msg = types.NewMsgApprove(
		chainID,
		1,
		coordinator,
	)
	_, err = h(ctx, msg)
	require.Error(t, err)

	// Create a add validator proposal
	addValidatorPayload := spnmocks.MockProposalAddValidatorPayload()
	addAccountPayload = spnmocks.MockProposalAddAccountPayload()
	valAddress := addValidatorPayload.ValidatorAddress
	addAccountPayload.Address = sdk.AccAddress(valAddress)
	k.SetAccount(ctx, chainID, addAccountPayload.Address, addAccountPayload)	// Simulate account address already being provided
	msgProposalValidator := types.NewMsgProposalAddValidator(
		chainID,
		proposalCreator,
		addValidatorPayload,
	)
	_, err = h(ctx, msgProposalValidator)
	require.NoError(t, err)

	// The coordinator creator can approve the proposal
	msg = types.NewMsgApprove(
		chainID,
		2,
		coordinator,
	)
	_, err = h(ctx, msg)
	require.NoError(t, err)

	// The proposal is approved
	proposal, _ = k.GetProposal(ctx, chainID, 2)
	require.Equal(t, types.ProposalState_APPROVED, proposal.ProposalState.Status)

	// The validator address is set in the store
	validatorAddressSet := k.IsValidatorSet(ctx, chainID, valAddress)
	require.True(t, validatorAddressSet)

	// Prevent approving a proposal with an validator already in the genesis
	_, err = h(ctx, msgProposalValidator)
	require.NoError(t, err)
	msg = types.NewMsgApprove(
		chainID,
		3,
		coordinator,
	)
	_, err = h(ctx, msg)
	require.Error(t, err)
}

func TestHandleMsgProposalAddAccount(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	chainID := spnmocks.MockRandomAlphaString(5)

	coordinator := spnmocks.MockAccAddress()
	coordinatorIdentity, _ := k.IdentityKeeper.GetIdentifier(ctx, coordinator)

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
	chain.Creator = coordinatorIdentity
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
	require.Equal(t, types.ProposalState_PENDING, proposal.ProposalState.Status)
	require.Equal(t, creatorIdentity, proposal.ProposalInformation.Creator)
	_, ok := proposal.Payload.(*types.Proposal_AddAccountPayload)
	require.True(t, ok)

	// The proposal is added to the pending proposals
	pending := k.GetPendingProposals(ctx, chainID)
	require.Contains(t, pending.ProposalIDs, int32(0))

	// The proposal count is incremented
	count := k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(1), count)

	// A proposal created by the coordinator is pre-approved
	account := spnmocks.MockAccAddress()
	msg = types.NewMsgProposalAddAccount(
		chainID,
		coordinator,
		spnmocks.MockProposalAddAccountPayload(),
	)
	msg.Payload.Address = account
	_, err = h(ctx, msg)
	require.NoError(t, err)

	// Can retrieve the proposal
	proposal, found = k.GetProposal(ctx, chainID, 1)
	require.True(t, found)
	require.Equal(t, types.ProposalState_APPROVED, proposal.ProposalState.Status)
	require.Equal(t, coordinatorIdentity, proposal.ProposalInformation.Creator)
	_, ok = proposal.Payload.(*types.Proposal_AddAccountPayload)
	require.True(t, ok)

	// The proposal is added to the pending proposals
	approved := k.GetApprovedProposals(ctx, chainID)
	require.Contains(t, approved.ProposalIDs, int32(1))

	// The proposal count is incremented
	count = k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(2), count)

	// A proposal created by the coordinator is not appended if it is invalid
	// (invalid here because the account already exists)
	msg = types.NewMsgProposalAddAccount(
		chainID,
		coordinator,
		spnmocks.MockProposalAddAccountPayload(),
	)
	msg.Payload.Address = account
	_, err = h(ctx, msg)
	require.Error(t, err)
}

func TestHandleMsgProposalAddValidator(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
	chainID := spnmocks.MockRandomAlphaString(5)

	coordinator := spnmocks.MockAccAddress()
	coordinatorIdentity, _ := k.IdentityKeeper.GetIdentifier(ctx, coordinator)

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
	chain.Creator = coordinatorIdentity
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
	require.NoError(t, err)
	_, err = h(ctx, msg)
	require.NoError(t, err)
	_, err = h(ctx, msg)
	require.NoError(t, err)
	msgAccount := types.NewMsgProposalAddAccount(
		chainID,
		creator,
		spnmocks.MockProposalAddAccountPayload(),
	)
	_, err = h(ctx, msgAccount)
	require.NoError(t, err)
	_, err = h(ctx, msgAccount)
	require.NoError(t, err)

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
	require.Equal(t, types.ProposalState_PENDING, proposal.ProposalState.Status)
	require.Equal(t, creatorIdentity, proposal.ProposalInformation.Creator)
	_, ok := proposal.Payload.(*types.Proposal_AddValidatorPayload)
	require.True(t, ok)

	// The proposal is added to the pending proposals
	pending := k.GetPendingProposals(ctx, chainID)
	require.Contains(t, pending.ProposalIDs, int32(5))

	// The proposal count is incremented
	count := k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(6), count)

	// A proposal created by the coordinator is pre-approved
	// We need to first have an existing account for the validator proposal
	account := spnmocks.MockAccAddress()
	msgAccount = types.NewMsgProposalAddAccount(
		chainID,
		coordinator,
		spnmocks.MockProposalAddAccountPayload(),
	)
	msgAccount.Payload.Address = account
	_, err = h(ctx, msgAccount)
	require.NoError(t, err)

	msg = types.NewMsgProposalAddValidator(
		chainID,
		coordinator,
		spnmocks.MockProposalAddValidatorPayload(),
	)
	msg.Payload.ValidatorAddress = sdk.ValAddress(account)
	_, err = h(ctx, msg)
	require.NoError(t, err)

	// Can retrieve the proposal
	proposal, found = k.GetProposal(ctx, chainID, 7)
	require.True(t, found)
	require.Equal(t, types.ProposalState_APPROVED, proposal.ProposalState.Status)
	require.Equal(t, coordinatorIdentity, proposal.ProposalInformation.Creator)
	_, ok = proposal.Payload.(*types.Proposal_AddValidatorPayload)
	require.True(t, ok)

	// The proposal is added to the pending proposals
	approved := k.GetApprovedProposals(ctx, chainID)
	require.Contains(t, approved.ProposalIDs, int32(7))

	// The proposal count is incremented
	count = k.GetProposalCount(ctx, chainID)
	require.Equal(t, int32(8), count)

	// A proposal created by the coordinator is not appended if it is invalid
	msg = types.NewMsgProposalAddValidator(
		chainID,
		coordinator,
		spnmocks.MockProposalAddValidatorPayload(),
	)
	_, err = h(ctx, msg)
	require.Error(t, err)
}