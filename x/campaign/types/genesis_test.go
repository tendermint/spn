package types_test

import (
	"github.com/tendermint/spn/testutil/sample"
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
					sample.Campaign(0),
					sample.Campaign(1),
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
					sample.Campaign(0),
					sample.Campaign(0),
				},
				CampaignCount: 2,
			},
			valid: false,
		},
		{
			desc: "invalid campaign count",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					sample.Campaign(1),
				},
				CampaignCount: 0,
			},
			valid: false,
		},
		{
			desc: "invalid campaign",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					types.NewCampaign(0, invalidCampaignName, sample.Uint64(), sample.Coins(), false),
				},
				CampaignCount: 1,
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
