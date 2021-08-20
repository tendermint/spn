package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNValidator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Validator {
	items := make([]types.Validator, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

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
		assert.True(t, found)
		assert.Equal(t, item, rst)
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
		assert.False(t, found)
	}
}

func TestValidatorGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNValidator(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllValidator(ctx))
}
