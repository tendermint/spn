package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/tendermint/spn/x/launch/types"
)

func createNGenesisAccount(keeper *Keeper, ctx sdk.Context, n int) []types.GenesisAccount {
	items := make([]types.GenesisAccount, n)
	for i := range items {
		items[i].ChainID = strconv.Itoa(i)
		items[i].Address = strconv.Itoa(i)

		keeper.SetGenesisAccount(ctx, items[i])
	}
	return items
}

func TestGenesisAccountGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNGenesisAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGenesisAccount(ctx,
			item.ChainID,
			item.Address,
		)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestGenesisAccountRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
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
		assert.False(t, found)
	}
}

func TestGenesisAccountGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNGenesisAccount(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllGenesisAccount(ctx))
}
