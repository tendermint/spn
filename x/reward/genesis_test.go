package reward_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/reward"
	"github.com/tendermint/spn/x/reward/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RewardPoolList: []types.RewardPool{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Reward(t)
	reward.InitGenesis(ctx, *k, genesisState)
	got := reward.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RewardPoolList, got.RewardPoolList)
	// this line is used by starport scaffolding # genesis/test/assert
}
