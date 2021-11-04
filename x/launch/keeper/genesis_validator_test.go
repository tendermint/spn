package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func createNGenesisValidator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GenesisValidator {
	items := make([]types.GenesisValidator, n)
	for i := range items {
		items[i] = sample.GenesisValidator(uint64(i), strconv.Itoa(i))
		keeper.SetGenesisValidator(ctx, items[i])
	}
	return items
}

func TestGenesisValidatorGet(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNGenesisValidator(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGenesisValidator(ctx,
			item.LaunchID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestGenesisValidatorRemove(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNGenesisValidator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGenesisValidator(ctx,
			item.LaunchID,
			item.Address,
		)
		_, found := keeper.GetGenesisValidator(ctx,
			item.LaunchID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestGenesisValidatorGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNGenesisValidator(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllGenesisValidator(ctx))
}
