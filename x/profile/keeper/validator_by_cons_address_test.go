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

func createNValidatorByConsAddress(keeper *Keeper, ctx sdk.Context, n int) []types.ValidatorByConsAddress {
	items := make([]types.ValidatorByConsAddress, n)
	for i := range items {
		items[i].ConsAddress = strconv.Itoa(i)

		keeper.SetValidatorByConsAddress(ctx, items[i])
	}
	return items
}

func TestValidatorByConsAddressGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidatorByConsAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidatorByConsAddress(ctx,
			item.ConsAddress,
		)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestValidatorByConsAddressRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidatorByConsAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidatorByConsAddress(ctx,
			item.ConsAddress,
		)
		_, found := keeper.GetValidatorByConsAddress(ctx,
			item.ConsAddress,
		)
		assert.False(t, found)
	}
}

func TestValidatorByConsAddressGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidatorByConsAddress(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllValidatorByConsAddress(ctx))
}
