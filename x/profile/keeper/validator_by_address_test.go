package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/tendermint/spn/x/profile/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNValidatorByAddress(keeper *Keeper, ctx sdk.Context, n int) []types.ValidatorByAddress {
	items := make([]types.ValidatorByAddress, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetValidatorByAddress(ctx, items[i])
	}
	return items
}

func TestValidatorByAddressGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidatorByAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidatorByAddress(ctx,
			item.Address,
		)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestValidatorByAddressRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidatorByAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidatorByAddress(ctx,
			item.Address,
		)
		_, found := keeper.GetValidatorByAddress(ctx,
			item.Address,
		)
		assert.False(t, found)
	}
}

func TestValidatorByAddressGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidatorByAddress(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllValidatorByAddress(ctx))
}
