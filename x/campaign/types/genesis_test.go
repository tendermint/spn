package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestGenesisState_Validate(t *testing.T) {
	var (
		campaign1 = sample.Campaign(0)
		campaign2 = sample.Campaign(1)
		shares1   = sample.Shares()
		shares2   = sample.Shares()
		shares3   = sample.Shares()
		shares4   = sample.Shares()
	)
	campaign1.AllocatedShares = types.IncreaseShares(shares1, shares2)
	campaign1.TotalShares = campaign1.AllocatedShares
	campaign2.AllocatedShares = types.IncreaseShares(shares3, shares4)
	campaign2.TotalShares = campaign2.AllocatedShares

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
						CampaignID: campaign1.Id,
					},
					{
						CampaignID: campaign2.Id,
					},
				},
				CampaignList: []types.Campaign{
					campaign1,
					campaign2,
				},
				CampaignCount: 2,
				MainnetAccountList: []types.MainnetAccount{
					{
						CampaignID: campaign1.Id,
						Address:    sample.AccAddress(),
						Shares:     shares1,
					},
					{
						CampaignID: campaign2.Id,
						Address:    sample.AccAddress(),
						Shares:     shares3,
					},
				},
				MainnetVestingAccountList: []types.MainnetVestingAccount{
					{
						CampaignID: campaign1.Id,
						Address:    sample.AccAddress(),
						Shares:     shares2,
					},
					{
						CampaignID: campaign2.Id,
						Address:    sample.AccAddress(),
						Shares:     shares4,
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
					sample.MainnetAccount(0, sample.AccAddress()),
					sample.MainnetAccount(1, sample.AccAddress()),
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
					sample.MainnetAccount(330, sample.AccAddress()),
					sample.MainnetAccount(434, sample.AccAddress()),
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
			if !tc.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			campaignIDMap := make(map[uint64]types.Shares)
			for _, elem := range tc.genState.CampaignList {
				campaignIDMap[elem.Id] = elem.AllocatedShares
			}
			shares := make(map[uint64]types.Shares)

			for _, acc := range tc.genState.MainnetAccountList {
				// check if the campaign exist for mainnet accounts
				_, ok := campaignIDMap[acc.CampaignID]
				require.True(t, ok)

				// sum mainnet account shares
				if _, ok := shares[acc.CampaignID]; !ok {
					shares[acc.CampaignID] = types.EmptyShares()
				}
				shares[acc.CampaignID] = types.IncreaseShares(
					shares[acc.CampaignID],
					acc.Shares,
				)
			}
			for campaignID, share := range campaignIDMap {
				// check if the campaign shares is equal all accounts shares
				accShares, ok := shares[campaignID]
				require.True(t, ok)
				isEqualShares := types.IsEqualShares(accShares, share)
				require.True(t, isEqualShares)
			}
		})
	}
}
