package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				// this line is used by starport scaffolding # types/genesis/validField
				CampaignList: []types.Campaign{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				CampaignCount: 2,
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
		{
			desc: "duplicated campaign",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid campaign count",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					{
						Id: 1,
					},
				},
				CampaignCount: 0,
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
