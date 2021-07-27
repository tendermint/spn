package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/tendermint/spn/x/profile/types"
)

func createNCoordinatorByAddress(keeper *Keeper, ctx sdk.Context, n int) []types.CoordinatorByAddress {
	items := make([]types.CoordinatorByAddress, n)
	for i := range items {
		items[i].Address = fmt.Sprintf("%d", i)
		keeper.SetCoordinatorByAddress(ctx, items[i])
	}
	return items
}

func TestCoordinatorByAddressGet(t *testing.T) {
	keeper, ctx := SetupTestKeeper(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCoordinatorByAddress(ctx, item.Address)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestCoordinatorByAddressRemove(t *testing.T) {
	keeper, ctx := SetupTestKeeper(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoordinatorByAddress(ctx, item.Address)
		_, found := keeper.GetCoordinatorByAddress(ctx, item.Address)
		assert.False(t, found)
	}
}

func TestCoordinatorByAddressGetAll(t *testing.T) {
	keeper, ctx := SetupTestKeeper(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllCoordinatorByAddress(ctx))
}
