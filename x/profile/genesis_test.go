package profile_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile"
)

/*
// We use a genesis template from sample package, therefore this placeholder is not used
// this line is used by starport scaffolding # genesis/test/state
*/

func TestGenesis(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)

	genesisState := sample.ProfileGenesisState()
	profile.InitGenesis(ctx, *keeper, genesisState)
	got := profile.ExportGenesis(ctx, *keeper)

	// Compare lists
	require.Len(t, got.ValidatorList, len(genesisState.ValidatorList))
	require.Subset(t, genesisState.ValidatorList, got.ValidatorList)

	require.Len(t, got.CoordinatorList, len(genesisState.CoordinatorList))
	require.Subset(t, genesisState.CoordinatorList, got.CoordinatorList)

	require.Len(t, got.CoordinatorByAddressList, len(genesisState.CoordinatorByAddressList))
	require.Subset(t, genesisState.CoordinatorByAddressList, got.CoordinatorByAddressList)

	require.Equal(t, genesisState.CoordinatorCount, got.CoordinatorCount)

	// this line is used by starport scaffolding # genesis/test/assert
}
