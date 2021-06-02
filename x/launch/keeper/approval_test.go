package keeper_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestCheckProposalApproval(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chain := spnmocks.MockChain()
	chainID := chain.ChainID

	// Add the chain
	k.SetChain(ctx, *chain)

	// No error if account is not set
	payload := spnmocks.MockProposalAddAccountPayload()
	proposal, _ := types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		payload,
	)
	err := k.CheckProposalApproval(ctx, chainID, *proposal)
	require.NoError(t, err)

	// Prevent making checks on non existent chain
	err = k.CheckProposalApproval(ctx, spnmocks.MockRandomAlphaString(5), *proposal)
	require.Error(t, err)
}

func TestApplyProposalAddAccount(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chain := spnmocks.MockChain()
	chainID := chain.ChainID

	// Add the chain
	k.SetChain(ctx, *chain)

	// Can apply a new account
	payload := spnmocks.MockProposalAddAccountPayload()
	proposal, _ := types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		payload,
	)
	err := k.ApplyProposalApproval(ctx, chainID, *proposal)
	require.NoError(t, err)

	// Prevent perform on non existent chain
	err = k.ApplyProposalApproval(ctx, spnmocks.MockRandomAlphaString(5), *proposal)
	require.Error(t, err)
}

func TestApplyProposalAddValidator(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chain := spnmocks.MockChain()
	chainID := chain.ChainID

	// Add the chain
	k.SetChain(ctx, *chain)

	// Can apply a new validator
	payload := spnmocks.MockProposalAddValidatorPayload()
	proposal, _ := types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		payload,
	)
	err := k.ApplyProposalApproval(ctx, chainID, *proposal)
	require.NoError(t, err)

	// The peer ids of the chain is updated upon validator approval
	retrievedChain, _ := k.GetChain(ctx, chainID)
	require.Contains(t, retrievedChain.Peers, payload.Peer)

	// Prevent perform on non existent chain
	err = k.ApplyProposalApproval(ctx, spnmocks.MockRandomAlphaString(5), *proposal)
	require.Error(t, err)
}
