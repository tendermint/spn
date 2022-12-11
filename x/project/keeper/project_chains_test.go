package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func TestKeeper_AddChainToProject(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should fail if project does not exist", func(t *testing.T) {
		err := tk.ProjectKeeper.AddChainToProject(ctx, 0, 0)
		require.Error(t, err)
	})

	// Chains can be added
	t.Run("should allow adding chains to project", func(t *testing.T) {
		tk.ProjectKeeper.SetProject(ctx, sample.Project(r, 0))
		err := tk.ProjectKeeper.AddChainToProject(ctx, 0, 0)
		require.NoError(t, err)
		err = tk.ProjectKeeper.AddChainToProject(ctx, 0, 1)
		require.NoError(t, err)
		err = tk.ProjectKeeper.AddChainToProject(ctx, 0, 2)
		require.NoError(t, err)

		projectChains, found := tk.ProjectKeeper.GetProjectChains(ctx, 0)
		require.True(t, found)
		require.EqualValues(t, projectChains.ProjectID, uint64(0))
		require.Len(t, projectChains.Chains, 3)
		require.EqualValues(t, []uint64{0, 1, 2}, projectChains.Chains)
	})

	// Can't add an existing chain
	t.Run("should prevent adding existing chain to project", func(t *testing.T) {
		err := tk.ProjectKeeper.AddChainToProject(ctx, 0, 0)
		require.Error(t, err)
	})
}

func createNProjectChains(k *keeper.Keeper, ctx sdk.Context, n int) []types.ProjectChains {
	items := make([]types.ProjectChains, n)
	for i := range items {
		items[i].ProjectID = uint64(i)
		k.SetProjectChains(ctx, items[i])
	}
	return items
}

func TestProjectChainsGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should get all projects", func(t *testing.T) {
		items := createNProjectChains(tk.ProjectKeeper, ctx, 10)
		for _, item := range items {
			rst, found := tk.ProjectKeeper.GetProjectChains(ctx,
				item.ProjectID,
			)
			require.True(t, found)
			require.Equal(t, item, rst)
		}
	})
}

func TestProjectChainsGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should get all projects", func(t *testing.T) {
		items := createNProjectChains(tk.ProjectKeeper, ctx, 10)
		require.ElementsMatch(t, items, tk.ProjectKeeper.GetAllProjectChains(ctx))
	})
}
