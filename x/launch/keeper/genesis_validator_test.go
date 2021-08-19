package keeper_test

import (
	"strconv"
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGenesisValidator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GenesisValidator {
	items := make([]types.GenesisValidator, n)
	for i := range items {
		items[i] = *sample.GenesisValidator(strconv.Itoa(i), strconv.Itoa(i))
		keeper.SetGenesisValidator(ctx, items[i])
	}
	return items
}

func TestGenesisValidatorGet(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNGenesisValidator(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGenesisValidator(ctx,
			item.ChainID,
			item.Address,
		)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestGenesisValidatorRemove(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNGenesisValidator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGenesisValidator(ctx,
			item.ChainID,
			item.Address,
		)
		_, found := keeper.GetGenesisValidator(ctx,
			item.ChainID,
			item.Address,
		)
		assert.False(t, found)
	}
}

func TestGenesisValidatorGetAll(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNGenesisValidator(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllGenesisValidator(ctx))
}
