package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/reward/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			name:     "should allow valid default genesis",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			name: "should allow valid genesis state",
			genState: &types.GenesisState{
				RewardPools: []types.RewardPool{
					sample.RewardPool(r, 1),
					sample.RewardPool(r, 2),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
		{
			name: "should prevent duplicated rewardPool",
			genState: &types.GenesisState{
				RewardPools: []types.RewardPool{
					sample.RewardPool(r, 1),
					sample.RewardPool(r, 1),
				},
			},
			valid: false,
		},
		{
			name: "should prevent invalid rewardPool",
			genState: &types.GenesisState{
				RewardPools: []types.RewardPool{
					sample.RewardPool(r, 1),
					{}, // invalid reward pool
				},
			},
			valid: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
