package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/claim/keeper"
)

func createTestAirdropSupply(keeper *keeper.Keeper, ctx sdk.Context) sdk.Coin {
	item := sample.Coin(r)
	keeper.SetAirdropSupply(ctx, item)
	return item
}

func TestAirdropSupplyGet(t *testing.T) {
	k, ctx := keepertest.ClaimKeeper(t)
	item := createTestAirdropSupply(k, ctx)
	rst, found := k.GetAirdropSupply(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
