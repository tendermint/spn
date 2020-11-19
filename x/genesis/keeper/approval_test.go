package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/genesis/types"
	"testing"
)

func CheckProposalAddAccount(t  *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chain := spnmocks.MockChain()
	chainID := chain.ChainID

	// Add the chain
	k.SetChain(ctx, *chain)

	// No error if account is not set
	payload := spnmocks.MockProposalAddAccountPayload()
	address := payload.Address
	proposal, _ := types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		payload,
	)
	err := k.CheckProposalApproval(ctx, chainID, *proposal)
	require.NoError(t, err)

	// Prevent making checks on non existent chain
	err = k.CheckProposalApproval(ctx, spnmocks.MockRandomAlphaString(5), *proposal)
	require.Error(t, err)

	// Error if account is already set
	k.SetAccount(ctx, chainID, address)
	payload = spnmocks.MockProposalAddAccountPayload()
	payload.Address = address
	proposal, _ = types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		payload,
	)
	err = k.CheckProposalApproval(ctx, chainID, *proposal)
	require.Error(t, err)
}

func CheckProposalAddValidator(t  *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chain := spnmocks.MockChain()
	chainID := chain.ChainID

	// Add the chain
	k.SetChain(ctx, *chain)

	// Error if corresponding account is not set
	payload := spnmocks.MockProposalAddValidatorPayload()
	proposal, _ := types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		payload,
	)
	err := k.CheckProposalApproval(ctx, chainID, *proposal)
	require.Error(t, err)

	// No error if corresponding account is set
	payload = spnmocks.MockProposalAddValidatorPayload()
	msg, _ := payload.GetCreateValidatorMessage()
	valAddress, _ := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	accAddress := sdk.AccAddress(valAddress)
	k.SetAccount(ctx, chainID, accAddress)
	proposal, _ = types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		payload,
	)
	err = k.CheckProposalApproval(ctx, chainID, *proposal)
	require.NoError(t, err)

	// Prevent making checks on non existent chain
	err = k.CheckProposalApproval(ctx, spnmocks.MockRandomAlphaString(5), *proposal)
	require.Error(t, err)

	// Error if validator is already set
	k.SetValidator(ctx, chainID, valAddress)
	err = k.CheckProposalApproval(ctx, chainID, *proposal)
	require.Error(t, err)
}

func TestApplyProposalAddAccount(t  *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chain := spnmocks.MockChain()
	chainID := chain.ChainID

	// Add the chain
	k.SetChain(ctx, *chain)

	// Can apply a new account
	payload := spnmocks.MockProposalAddAccountPayload()
	address := payload.Address
	proposal, _ := types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		payload,
	)
	err := k.ApplyProposalApproval(ctx, chainID, *proposal)
	require.NoError(t, err)
	isAccountSet := k.IsAccountSet(ctx, chainID, address)
	require.True(t, isAccountSet)

	// Prevent perform on non existent chain
	err = k.ApplyProposalApproval(ctx, spnmocks.MockRandomAlphaString(5), *proposal)
	require.Error(t, err)
}

func TestApplyProposalAddValidator(t  *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chain := spnmocks.MockChain()
	chainID := chain.ChainID

	// Add the chain
	k.SetChain(ctx, *chain)

	// Can apply a new validator
	payload := spnmocks.MockProposalAddValidatorPayload()
	msg, _ := payload.GetCreateValidatorMessage()
	valAddress, _ := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	proposal, _ := types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		payload,
	)
	err := k.ApplyProposalApproval(ctx, chainID, *proposal)
	require.NoError(t, err)
	isValidatorSet := k.IsValidatorSet(ctx, chainID, valAddress)
	require.True(t, isValidatorSet)

	// The peer ids of the chain is updated upon validator approval
	retrievedChain, _ := k.GetChain(ctx, chainID)
	require.Contains(t, retrievedChain.Peers, payload.Peer)

	// Prevent perform on non existent chain
	err = k.ApplyProposalApproval(ctx, spnmocks.MockRandomAlphaString(5), *proposal)
	require.Error(t, err)
}