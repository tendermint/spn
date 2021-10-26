package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func createNValidator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Validator {
	items := make([]types.Validator, n)
	for i := range items {
		items[i].Address = sample.Address()
		keeper.SetValidator(ctx, items[i])
	}
	return items
}

func TestValidatorGet(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNValidator(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidator(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestValidatorRemove(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNValidator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidator(ctx,
			item.Address,
		)
		_, found := keeper.GetValidator(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestValidatorGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNValidator(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllValidator(ctx))
}
