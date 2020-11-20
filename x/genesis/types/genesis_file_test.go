package types_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/genesis/types"
	"testing"
)

func TestValidateAndComplete(t *testing.T) {
	cdc := spnmocks.MockCodec()

	// Should return no error for valid genesis
	genesis := spnmocks.MockGenesis()
	err := genesis.ValidateAndComplete()
	require.NoError(t, err)

	// Return error if accounts present in the genesis
	addAccountPayload := spnmocks.MockProposalAddAccountPayload()	// Use proposals to add an account for testing purpose
	proposal, _ := types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		addAccountPayload,
		)
	genesis.ApplyProposals(cdc, []types.Proposal{*proposal})
	err = genesis.ValidateAndComplete()
	require.Error(t, err)

	// Return error if gentxs present in the genesis
	genesis = spnmocks.MockGenesis()
	addValidatorPayload := spnmocks.MockProposalAddValidatorPayload()	// Use proposals to add a gentx for testing purpose
	proposalGentx, _ := types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		addValidatorPayload,
	)
	genesis.ApplyProposals(cdc, []types.Proposal{*proposalGentx})
	err = genesis.ValidateAndComplete()
	require.Error(t, err)
}