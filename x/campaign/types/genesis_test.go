package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
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
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: 0,
					},
					{
						CampaignID: 1,
					},
				},
				CampaignList: []types.Campaign{
					sample.Campaign(0),
					sample.Campaign(1),
				},
				CampaignCount: 2,
				MainnetAccountList: []types.MainnetAccount{
					sample.MainnetAccount(0, sample.Address()),
					sample.MainnetAccount(1, sample.Address()),
				},
				MainnetVestingAccountList: []types.MainnetVestingAccount{
					{
						CampaignID: 0,
						Address:    "0",
					},
					{
						CampaignID: 1,
						Address:    "1",
					},
				},
			},
			valid: true,
		},
		{
			desc: "non existing campaign for mainnet vesting account",
			genState: &types.GenesisState{
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: 0,
					},
					{
						CampaignID: 1,
					},
				},
				CampaignList: []types.Campaign{
					sample.Campaign(0),
					sample.Campaign(1),
				},
				CampaignCount: 2,
				MainnetAccountList: []types.MainnetAccount{
					sample.MainnetAccount(0, sample.Address()),
					sample.MainnetAccount(1, sample.Address()),
				},
				MainnetVestingAccountList: []types.MainnetVestingAccount{
					{
						CampaignID: 33333,
						Address:    "33333",
					},
					{
						CampaignID: 9999,
						Address:    "9999",
					},
				},
			},
			valid: false,
		},
		{
			desc: "non existing campaign for mainnet account",
			genState: &types.GenesisState{
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: 0,
					},
					{
						CampaignID: 1,
					},
				},
				CampaignList: []types.Campaign{
					sample.Campaign(0),
					sample.Campaign(1),
				},
				CampaignCount: 2,
				MainnetAccountList: []types.MainnetAccount{
					sample.MainnetAccount(330, sample.Address()),
					sample.MainnetAccount(434, sample.Address()),
				},
			},
			valid: false,
		},
		{
			desc: "non existing campaign for chains",
			genState: &types.GenesisState{
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: 2,
					},
					{
						CampaignID: 4,
					},
				},
				CampaignList: []types.Campaign{
					sample.Campaign(99),
					sample.Campaign(88),
				},
				CampaignCount: 100,
			},
			valid: false,
		},
		{
			desc: "duplicated campaignChains",
			genState: &types.GenesisState{
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: 0,
					},
					{
						CampaignID: 0,
					},
				},
			},
		},
		{
			desc: "duplicated mainnetVestingAccount",
			genState: &types.GenesisState{
				MainnetVestingAccountList: []types.MainnetVestingAccount{
					{
						CampaignID: 0,
						Address:    "0",
					},
					{
						CampaignID: 0,
						Address:    "0",
					},
				},
			},
			valid: false,
		},
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
		{
			desc: "duplicated mainnetAccount",
			genState: &types.GenesisState{
				MainnetAccountList: []types.MainnetAccount{
					{
						CampaignID: 0,
						Address:    "0",
					},
					{
						CampaignID: 0,
						Address:    "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
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
