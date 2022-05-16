package types_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	tc2 "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestGenesisState_Validate(t *testing.T) {
	var (
		campaign1 = sample.Campaign(r, 0)
		campaign2 = sample.Campaign(r, 1)
		shares1   = sample.Shares(r)
		shares2   = sample.Shares(r)
		shares3   = sample.Shares(r)
		shares4   = sample.Shares(r)
	)
	sharesCampaign1 := types.IncreaseShares(shares1, shares2)
	campaign1.AllocatedShares = sharesCampaign1
	campaign1.CoordinatorID = 0

	sharesCampaign2 := types.IncreaseShares(shares3, shares4)
	campaign2.AllocatedShares = sharesCampaign2
	campaign2.CoordinatorID = 1

	for _, tc := range []struct {
		desc         string
		genState     *types.GenesisState
		errorMessage string
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				// this line is used by starport scaffolding # types/genesis/validField
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: campaign1.CampaignID,
					},
					{
						CampaignID: campaign2.CampaignID,
					},
				},
				CampaignList: []types.Campaign{
					campaign1,
					campaign2,
				},
				CampaignCounter: 2,
				MainnetAccountList: []types.MainnetAccount{
					{
						CampaignID: campaign1.CampaignID,
						Address:    sample.Address(r),
						Shares:     shares1,
					},
					{
						CampaignID: campaign2.CampaignID,
						Address:    sample.Address(r),
						Shares:     shares3,
					},
				},
				MainnetVestingAccountList: []types.MainnetVestingAccount{
					{
						CampaignID:     campaign1.CampaignID,
						Address:        sample.Address(r),
						VestingOptions: *types.NewShareDelayedVesting(shares2, shares2, time.Now().Unix()),
					},
					{
						CampaignID:     campaign2.CampaignID,
						Address:        sample.Address(r),
						VestingOptions: *types.NewShareDelayedVesting(shares4, shares4, time.Now().Unix()),
					},
				},
				TotalShares: spntypes.TotalShareNumber,
				Params:      types.DefaultParams(),
			},
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
					sample.Campaign(r, 0),
					sample.Campaign(r, 1),
				},
				CampaignCounter: 2,
				MainnetAccountList: []types.MainnetAccount{
					sample.MainnetAccount(r, 0, sample.Address(r)),
					sample.MainnetAccount(r, 1, sample.Address(r)),
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
				TotalShares: spntypes.TotalShareNumber,
			},
			errorMessage: "campaign id 33333 doesn't exist for mainnet vesting account 33333",
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
					sample.Campaign(r, 0),
					sample.Campaign(r, 1),
				},
				CampaignCounter: 2,
				MainnetAccountList: []types.MainnetAccount{
					sample.MainnetAccount(r, 330, "330"),
				},
				TotalShares: spntypes.TotalShareNumber,
			},
			errorMessage: "campaign id 330 doesn't exist for mainnet account 330",
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
					sample.Campaign(r, 99),
					sample.Campaign(r, 88),
				},
				CampaignCounter: 100,
				TotalShares:     spntypes.TotalShareNumber,
			},
			errorMessage: "campaign id 2 doesn't exist for chains",
		},
		{
			desc: "duplicated campaignChains",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					sample.Campaign(r, 0),
				},
				CampaignCounter: 1,
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: 0,
					},
					{
						CampaignID: 0,
					},
				},
				TotalShares: spntypes.TotalShareNumber,
			},
			errorMessage: "duplicated index for campaignChains",
		},
		{
			desc: "duplicated mainnetVestingAccount",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					sample.Campaign(r, 0),
				},
				CampaignCounter: 1,
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
				TotalShares: spntypes.TotalShareNumber,
			},
			errorMessage: "duplicated index for mainnetVestingAccount",
		},
		{
			desc: "duplicated campaign",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					sample.Campaign(r, 0),
					sample.Campaign(r, 0),
				},
				CampaignCounter: 2,
				TotalShares:     spntypes.TotalShareNumber,
			},
			errorMessage: "duplicated id for campaign",
		},
		{
			desc: "invalid campaign count",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					sample.Campaign(r, 1),
				},
				CampaignCounter: 0,
				TotalShares:     spntypes.TotalShareNumber,
			},
			errorMessage: "campaign id should be lower or equal than the last id",
		},
		{
			desc: "invalid campaign",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					types.NewCampaign(
						0,
						invalidCampaignName,
						sample.Uint64(r),
						sample.TotalSupply(r),
						sample.Metadata(r, 20),
						sample.Duration(r).Milliseconds(),
					),
				},
				CampaignCounter: 1,
				TotalShares:     spntypes.TotalShareNumber,
			},
			errorMessage: "invalid campaign 0: campaign name can only contain alphanumerical characters or hyphen",
		},
		{
			desc: "duplicated mainnetAccount",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					sample.Campaign(r, 0),
				},
				CampaignCounter: 1,
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
				TotalShares: spntypes.TotalShareNumber,
			},
			errorMessage: "duplicated index for mainnetAccount",
		},
		{
			desc: "invalid allocations",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					{
						CampaignID:         0,
						CampaignName:       "test",
						CoordinatorID:      0,
						MainnetID:          0,
						MainnetInitialized: false,
						TotalSupply:        nil,
						AllocatedShares:    types.NewSharesFromCoins(tc2.Coins(t, fmt.Sprintf("%dstake", spntypes.TotalShareNumber+1))),
						Metadata:           nil,
					},
				},
				CampaignCounter: 1,
				MainnetAccountList: []types.MainnetAccount{
					{
						CampaignID: 0,
						Address:    "0",
					},
				},
				TotalShares: spntypes.TotalShareNumber,
			},
			errorMessage: "invalid campaign 0: more allocated shares than total shares",
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.errorMessage != "" {
				require.Error(t, err)
				require.Equal(t, tc.errorMessage, err.Error())
				return
			}
			require.NoError(t, err)

			campaignIDMap := make(map[uint64]types.Shares)
			for _, elem := range tc.genState.CampaignList {
				campaignIDMap[elem.CampaignID] = elem.AllocatedShares
			}
			shares := make(map[uint64]types.Shares)

			for _, acc := range tc.genState.MainnetAccountList {
				// check if the campaign exists for mainnet accounts
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

			for _, acc := range tc.genState.MainnetVestingAccountList {
				// check if the campaign exists for mainnet accounts
				_, ok := campaignIDMap[acc.CampaignID]
				require.True(t, ok)

				// sum mainnet account shares
				if _, ok := shares[acc.CampaignID]; !ok {
					shares[acc.CampaignID] = types.EmptyShares()
				}
				totalShares, err := acc.GetTotalShares()
				require.NoError(t, err)

				shares[acc.CampaignID] = types.IncreaseShares(
					shares[acc.CampaignID],
					totalShares,
				)
			}

			for campaignID, share := range campaignIDMap {
				// check if the campaign shares is equal all accounts shares
				accShares, ok := shares[campaignID]
				require.True(t, ok)
				isLowerEqual := accShares.IsAllLTE(share)
				require.True(t, isLowerEqual)
			}
		})
	}
}

func TestGenesisState_ValidateParams(t *testing.T) {
	for _, tc := range []struct {
		desc          string
		genState      types.GenesisState
		shouldBeValid bool
	}{
		{
			desc: "max total supply below min total supply",
			genState: types.GenesisState{
				Params: types.NewParams(types.DefaultMinTotalSupply, types.DefaultMinTotalSupply.Sub(sdk.OneInt()), types.DefaultCampaignCreationFee),
			},
			shouldBeValid: false,
		},
		{
			desc: "valid parameters",
			genState: types.GenesisState{
				Params: types.NewParams(types.DefaultMinTotalSupply, types.DefaultMinTotalSupply.Add(sdk.OneInt()), types.DefaultCampaignCreationFee),
			},
			shouldBeValid: true,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.shouldBeValid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
