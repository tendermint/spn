package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func createNParamChange(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ParamChange {
	items := make([]types.ParamChange, n)
	for i := range items {
		keeper.SetChain(ctx, sample.Chain(r, uint64(i), sample.Uint64(r)))
		items[i] = sample.ParamChange(r, uint64(i))
		keeper.SetParamChange(ctx, items[i])
	}
	return items
}

func TestParamChangeGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNParamChange(tk.LaunchKeeper, ctx, 10)

	t.Run("should get a change param", func(t *testing.T) {
		for i, item := range items {
			rst, found := tk.LaunchKeeper.GetParamChange(ctx,
				uint64(i),
				item.Module,
				item.Param,
			)
			require.True(t, found)
			require.Equal(t, item, rst)
		}
	})
}

func TestParamChangeGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNParamChange(tk.LaunchKeeper, ctx, 10)

	t.Run("should get all change param", func(t *testing.T) {
		require.ElementsMatch(t, items, tk.LaunchKeeper.GetAllParamChange(ctx))
	})
}
