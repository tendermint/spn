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
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()

	t.Run("should allow import and export the genesis state", func(t *testing.T) {
		genesisState := sample.ProfileGenesisState(r)
		profile.InitGenesis(ctx, *tk.ProfileKeeper, genesisState)
		got := profile.ExportGenesis(ctx, *tk.ProfileKeeper)

		// Compare lists
		require.ElementsMatch(t, genesisState.Validators, got.Validators)
		require.ElementsMatch(t, genesisState.ValidatorByOperatorAddresses, got.ValidatorByOperatorAddresses)
		require.ElementsMatch(t, genesisState.Coordinators, got.Coordinators)
		require.ElementsMatch(t, genesisState.CoordinatorByAddresses, got.CoordinatorByAddresses)
		require.Equal(t, genesisState.CoordinatorCounter, got.CoordinatorCounter)
	})

	// this line is used by starport scaffolding # genesis/test/assert
}
