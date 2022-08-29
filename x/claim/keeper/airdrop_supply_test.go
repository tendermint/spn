package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	claim "github.com/tendermint/spn/x/claim/types"
)

func TestAirdropSupplyGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	sampleSupply := sample.Coin(r)
	tk.ClaimKeeper.SetAirdropSupply(ctx, sampleSupply)

	rst, found := tk.ClaimKeeper.GetAirdropSupply(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&sampleSupply),
		nullify.Fill(&rst),
	)
}

func TestAirdropSupplyRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	tk.ClaimKeeper.SetAirdropSupply(ctx, sample.Coin(r))
	_, found := tk.ClaimKeeper.GetAirdropSupply(ctx)
	require.True(t, found)
	tk.ClaimKeeper.RemoveAirdropSupply(ctx)
	_, found = tk.ClaimKeeper.GetAirdropSupply(ctx)
	require.False(t, found)
}

func TestKeeper_InitializeAirdropSupply(t *testing.T) {
	// TODO: use mock for bank module to test critical errors
	// https://github.com/tendermint/spn/issues/838
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	tests := []struct {
		name          string
		airdropSupply sdk.Coin
	}{
		{
			name:          "should allows setting airdrop supply",
			airdropSupply: tc.Coin(t, "10000foo"),
		},
		{
			name:          "should allows specifying a new token for the supply",
			airdropSupply: tc.Coin(t, "125000bar"),
		},
		{
			name:          "should allows modifying a token for the supply",
			airdropSupply: tc.Coin(t, "525000bar"),
		},
		{
			name:          "should allows setting airdrop supply to zero",
			airdropSupply: sdk.NewCoin("foo", sdkmath.ZeroInt()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tk.ClaimKeeper.InitializeAirdropSupply(ctx, tt.airdropSupply)
			require.NoError(t, err)

			airdropSupply, found := tk.ClaimKeeper.GetAirdropSupply(ctx)
			require.True(t, found)
			require.True(t, airdropSupply.IsEqual(tt.airdropSupply))

			moduleBalance := tk.BankKeeper.GetBalance(
				ctx,
				tk.AccountKeeper.GetModuleAddress(claim.ModuleName),
				airdropSupply.Denom,
			)
			require.True(t, moduleBalance.IsEqual(tt.airdropSupply))
		})
	}
}
