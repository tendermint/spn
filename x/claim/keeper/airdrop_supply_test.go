package keeper_test

import (
	tc "github.com/tendermint/spn/testutil/constructor"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
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

func TestKeeper_InitializeAirdropSupply(t *testing.T) {
	// TODO: use mock for bank module to test critical errors
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	tests := []struct {
		name          string
		airdropSupply sdk.Coin
	}{
		{
			name:          "should allows to set airdrop supply",
			airdropSupply: tc.Coin(t, "10000foo"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tk.ClaimKeeper.InitializeAirdropSupply(ctx, tt.airdropSupply)
			require.NoError(t, err)
		})
	}
}
