package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	projectkeeper "github.com/tendermint/spn/x/project/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func createNProject(keeper *projectkeeper.Keeper, ctx sdk.Context, n int) []types.Project {
	items := make([]types.Project, n)
	for i := range items {
		items[i] = sample.Project(r, 0)
		items[i].ProjectID = keeper.AppendProject(ctx, items[i])
	}
	return items
}

func TestProjectGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should get projects", func(t *testing.T) {
		items := createNProject(tk.ProjectKeeper, ctx, 10)
		for _, item := range items {
			got, found := tk.ProjectKeeper.GetProject(ctx, item.ProjectID)
			require.True(t, found)
			require.Equal(t, item, got)
		}
	})
}

func TestProjectRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should remove projects", func(t *testing.T) {
		items := createNProject(tk.ProjectKeeper, ctx, 10)
		for _, item := range items {
			tk.ProjectKeeper.RemoveProject(ctx, item.ProjectID)
			_, found := tk.ProjectKeeper.GetProject(ctx, item.ProjectID)
			require.False(t, found)
		}
	})
}

func TestProjectGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNProject(tk.ProjectKeeper, ctx, 10)
	t.Run("should get all projects", func(t *testing.T) {
		require.ElementsMatch(t, items, tk.ProjectKeeper.GetAllProject(ctx))
	})
}

func TestProjectCount(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should get project count", func(t *testing.T) {
		items := createNProject(tk.ProjectKeeper, ctx, 10)
		counter := uint64(len(items))
		require.Equal(t, counter, tk.ProjectKeeper.GetProjectCounter(ctx))
	})
}
