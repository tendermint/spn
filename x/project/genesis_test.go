package project_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project"
)

/*
// We use a genesis template from sample package, therefore this placeholder is not used
// this line is used by starport scaffolding # genesis/test/state
*/

func TestGenesis(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()

	t.Run("should allow importing and exporting genesis", func(t *testing.T) {
		genesisState := sample.ProjectGenesisStateWithAccounts(r)

		project.InitGenesis(ctx, *tk.ProjectKeeper, genesisState)
		got := *project.ExportGenesis(ctx, *tk.ProjectKeeper)

		require.ElementsMatch(t, genesisState.ProjectChains, got.ProjectChains)
		require.ElementsMatch(t, genesisState.Projects, got.Projects)
		require.Equal(t, genesisState.ProjectCounter, got.ProjectCounter)
		require.ElementsMatch(t, genesisState.MainnetAccounts, got.MainnetAccounts)
		require.Equal(t, genesisState.Params, got.Params)
		maxShares := tk.ProjectKeeper.GetTotalShares(ctx)
		require.Equal(t, uint64(spntypes.TotalShareNumber), maxShares)
	})

	// this line is used by starport scaffolding # genesis/test/assert
}
