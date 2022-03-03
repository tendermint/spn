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

	genesisState := sample.LaunchGenesisState()
	launch.InitGenesis(ctx, *tk.LaunchKeeper, genesisState)
	got := launch.ExportGenesis(ctx, *tk.LaunchKeeper)

	// Compare lists
	require.ElementsMatch(t, genesisState.ChainList, got.ChainList)
	require.Equal(t, genesisState.ChainCounter, got.ChainCounter)

	require.ElementsMatch(t, genesisState.GenesisAccountList, got.GenesisAccountList)
	require.ElementsMatch(t, genesisState.VestingAccountList, got.VestingAccountList)
	require.ElementsMatch(t, genesisState.GenesisValidatorList, got.GenesisValidatorList)
	require.ElementsMatch(t, genesisState.RequestList, got.RequestList)
	require.ElementsMatch(t, genesisState.RequestCounterList, got.RequestCounterList)

	require.Equal(t, genesisState.Params, got.Params)

	// this line is used by starport scaffolding # genesis/test/assert
}
