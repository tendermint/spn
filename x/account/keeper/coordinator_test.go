package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/spn/x/account/types"
)

func createNCoordinator(keeper *Keeper, ctx sdk.Context, n int) []types.Coordinator {
	items := make([]types.Coordinator, n)
	for i := range items {
		items[i].CoordinatorId = keeper.AppendCoordinator(ctx, items[i])
	}
	return items
}

func TestCoordinatorGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoordinator(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetCoordinator(ctx, item.CoordinatorId))
	}
}

func TestCoordinatorExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoordinator(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasCoordinator(ctx, item.CoordinatorId))
	}
}

func TestCoordinatorRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoordinator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoordinator(ctx, item.CoordinatorId)
		assert.False(t, keeper.HasCoordinator(ctx, item.CoordinatorId))
	}
}

func TestCoordinatorGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoordinator(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllCoordinator(ctx))
}

func TestCoordinatorCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoordinator(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetCoordinatorCount(ctx))
}
