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
	keeper, ctx := testkeeper.Launch(t)

	genesisState := sample.LaunchGenesisState()
	launch.InitGenesis(ctx, *keeper, genesisState)
	got := launch.ExportGenesis(ctx, *keeper)

	// Compare lists
	require.Len(t, got.ChainList, len(genesisState.ChainList))
	require.Subset(t, genesisState.ChainList, got.ChainList)

	require.Equal(t, genesisState.ChainCount, got.ChainCount)

	require.Len(t, got.GenesisAccountList, len(genesisState.GenesisAccountList))
	require.Subset(t, genesisState.GenesisAccountList, got.GenesisAccountList)

	require.Len(t, got.VestingAccountList, len(genesisState.VestingAccountList))
	require.Subset(t, genesisState.VestingAccountList, got.VestingAccountList)

	require.Len(t, got.GenesisValidatorList, len(genesisState.GenesisValidatorList))
	require.Subset(t, genesisState.GenesisValidatorList, got.GenesisValidatorList)

	require.Len(t, got.RequestList, len(genesisState.RequestList))
	require.Subset(t, genesisState.RequestList, got.RequestList)

	require.Len(t, got.RequestCountList, len(genesisState.RequestCountList))
	require.Subset(t, genesisState.RequestCountList, got.RequestCountList)

	require.Equal(t, genesisState.Params, got.Params)

	// this line is used by starport scaffolding # genesis/test/assert
}
