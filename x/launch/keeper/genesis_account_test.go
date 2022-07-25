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

func createNGenesisAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GenesisAccount {
	items := make([]types.GenesisAccount, n)
	for i := range items {
		keeper.SetChain(ctx, sample.Chain(r, uint64(i), sample.Uint64(r)))
		items[i] = sample.GenesisAccount(r, uint64(i), sample.Address(r))
		keeper.SetGenesisAccount(ctx, items[i])
	}
	return items
}

func TestGenesisAccountGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNGenesisAccount(tk.LaunchKeeper, ctx, 10)

	t.Run("should get a genesis account", func(t *testing.T) {
		for _, item := range items {
			rst, found := tk.LaunchKeeper.GetGenesisAccount(ctx,
				item.LaunchID,
				item.Address,
			)
			require.True(t, found)
			require.Equal(t, item, rst)
		}
	})
}

func TestGenesisAccountRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNGenesisAccount(tk.LaunchKeeper, ctx, 10)

	t.Run("should remove a genesis account", func(t *testing.T) {
		for _, item := range items {
			tk.LaunchKeeper.RemoveGenesisAccount(ctx,
				item.LaunchID,
				item.Address,
			)
			_, found := tk.LaunchKeeper.GetGenesisAccount(ctx,
				item.LaunchID,
				item.Address,
			)
			require.False(t, found)
		}
	})
}

func TestGenesisAccountGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNGenesisAccount(tk.LaunchKeeper, ctx, 10)

	t.Run("should get all genesis account", func(t *testing.T) {
		require.ElementsMatch(t, items, tk.LaunchKeeper.GetAllGenesisAccount(ctx))
	})
}
