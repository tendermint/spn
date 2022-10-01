package reward_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/reward"
	"github.com/tendermint/spn/x/reward/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RewardPools: []types.RewardPool{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	t.Run("should allow importing and exporting genesis", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		reward.InitGenesis(ctx, *tk.RewardKeeper, genesisState)
		got := reward.ExportGenesis(ctx, *tk.RewardKeeper)
		require.NotNil(t, got)

		nullify.Fill(&genesisState)
		nullify.Fill(got)

		require.ElementsMatch(t, genesisState.RewardPools, got.RewardPools)
	})
	// this line is used by starport scaffolding # genesis/test/assert
}
