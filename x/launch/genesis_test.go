package launch_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch"
)

/*
// We use a genesis template from sample package, therefore this placeholder is not used
// this line is used by starport scaffolding # genesis/test/state
*/

func TestGenesis(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()

	t.Run("should allow import and export the genesis state", func(t *testing.T) {
		genesisState := sample.LaunchGenesisState(r)
		launch.InitGenesis(ctx, *tk.LaunchKeeper, genesisState)
		got := launch.ExportGenesis(ctx, *tk.LaunchKeeper)

		// Compare lists
		require.ElementsMatch(t, genesisState.Chains, got.Chains)
		require.Equal(t, genesisState.ChainCounter, got.ChainCounter)

		require.ElementsMatch(t, genesisState.GenesisAccounts, got.GenesisAccounts)
		require.ElementsMatch(t, genesisState.VestingAccounts, got.VestingAccounts)
		require.ElementsMatch(t, genesisState.GenesisValidators, got.GenesisValidators)
		require.ElementsMatch(t, genesisState.Requests, got.Requests)
		require.ElementsMatch(t, genesisState.RequestCounters, got.RequestCounters)

		require.Equal(t, genesisState.Params, got.Params)
	})
	// this line is used by starport scaffolding # genesis/test/assert
}
