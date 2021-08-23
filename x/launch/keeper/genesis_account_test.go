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

func createNGenesisAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GenesisAccount {
	items := make([]types.GenesisAccount, n)
	for i := range items {
		items[i] = *sample.GenesisAccount(strconv.Itoa(i), strconv.Itoa(i))
		keeper.SetGenesisAccount(ctx, items[i])
	}
	return items
}

func TestGenesisAccountGet(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNGenesisAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGenesisAccount(ctx,
			item.ChainID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestGenesisAccountRemove(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNGenesisAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGenesisAccount(ctx,
			item.ChainID,
			item.Address,
		)
		_, found := keeper.GetGenesisAccount(ctx,
			item.ChainID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestGenesisAccountGetAll(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNGenesisAccount(keeper, ctx, 10)
	require.Equal(t, items, keeper.GetAllGenesisAccount(ctx))
}
