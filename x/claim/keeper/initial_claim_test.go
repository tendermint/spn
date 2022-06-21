package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/claim/keeper"
	"github.com/tendermint/spn/x/claim/types"
)

func createTestInitialClaim(keeper *keeper.Keeper, ctx sdk.Context) types.InitialClaim {
	item := types.InitialClaim{}
	keeper.SetInitialClaim(ctx, item)
	return item
}

func TestInitialClaimGet(t *testing.T) {
	keeper, ctx := keepertest.ClaimKeeper(t)
	item := createTestInitialClaim(keeper, ctx)
	rst, found := keeper.GetInitialClaim(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestInitialClaimRemove(t *testing.T) {
	keeper, ctx := keepertest.ClaimKeeper(t)
	createTestInitialClaim(keeper, ctx)
	keeper.RemoveInitialClaim(ctx)
	_, found := keeper.GetInitialClaim(ctx)
	require.False(t, found)
}
