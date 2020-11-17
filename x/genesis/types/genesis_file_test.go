package types_test

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

// TODO: This test could be considered as an integration test -> Move it to the integration test
func TestApplyProposals(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)
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
		h(ctx, msgAddAccount)

		// Send add validator proposal
		msgAddValidator := types.NewMsgProposalAddValidator(
			chainID,
			accAddress,
			addValidatorpayload,
		)
		h(ctx, msgAddValidator)
	}
	for i:=0; i<10; i++ {
		addAccountPayload := spnmocks.MockProposalAddAccountPayload()

		// Send add account proposal
		msgAddAccount := types.NewMsgProposalAddAccount(
			chainID,
			addAccountPayload.Address,
			addAccountPayload,
		)
		h(ctx, msgAddAccount)
	}

	// Approve all proposals
	for i:=0; i<20; i++ {
		msg := types.NewMsgApprove(
			chainID,
			int32(i),
			coordinator,
		)
		h(ctx, msg)
	}

	// Can retrieve the current genesis with all the approved proposals
	var req *types.QueryCurrentGenesisRequest
	req.ChainID = chainID
	res, err := k.CurrentGenesis(context.Background(), req)
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

	// Analyse velidators
	appGenesisState, err := genutiltypes.GenesisStateFromGenDoc(genesisDoc)
	require.NoError(t, err)
	genesisState := genutiltypes.GetGenesisStateFromAppState(cdc, appGenesisState)
	require.Equal(t, 10, len(genesisState.GenTxs))
}