package types_test

import (
	"fmt"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	tc2 "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
)

func TestGenesisState_Validate(t *testing.T) {
	var (
		project1 = sample.Project(r, 0)
		project2 = sample.Project(r, 1)
		shares1  = sample.Shares(r)
		shares2  = sample.Shares(r)
		shares3  = sample.Shares(r)
		shares4  = sample.Shares(r)
	)
	sharesProject1 := types.IncreaseShares(shares1, shares2)
	project1.AllocatedShares = sharesProject1
	project1.CoordinatorID = 0

	sharesProject2 := types.IncreaseShares(shares3, shares4)
	project2.AllocatedShares = sharesProject2
	project2.CoordinatorID = 1

	for _, tc := range []struct {
		name         string
		genState     *types.GenesisState
		errorMessage string
	}{
		{
			name:     "should allow validation of valid default genesis",
			genState: types.DefaultGenesis(),
		},
		{
			name: "should allow validation of valid genesis",
			genState: &types.GenesisState{
				// this line is used by starport scaffolding # types/genesis/validField
				ProjectChains: []types.ProjectChains{
					{
						ProjectID: project1.ProjectID,
					},
					{
						ProjectID: project2.ProjectID,
					},
				},
				Projects: []types.Project{
					project1,
					project2,
				},
				ProjectCounter: 2,
				MainnetAccounts: []types.MainnetAccount{
					{
						ProjectID: project1.ProjectID,
						Address:   sample.Address(r),
						Shares:    shares1,
					},
					{
						ProjectID: project2.ProjectID,
						Address:   sample.Address(r),
						Shares:    shares3,
					},
				},
				TotalShares: spntypes.TotalShareNumber,
				Params:      types.DefaultParams(),
			},
		},
		{
			name: "should prevent validation of genesis with non existing project for mainnet account",
			genState: &types.GenesisState{
				ProjectChains: []types.ProjectChains{
					{
						ProjectID: 0,
					},
					{
						ProjectID: 1,
					},
				},
				Projects: []types.Project{
					sample.Project(r, 0),
					sample.Project(r, 1),
				},
				ProjectCounter: 2,
				MainnetAccounts: []types.MainnetAccount{
					sample.MainnetAccount(r, 330, "330"),
				},
				TotalShares: spntypes.TotalShareNumber,
			},
			errorMessage: "project id 330 doesn't exist for mainnet account 330",
		},
		{
			name: "should prevent validation of genesis with non existing project for chains",
			genState: &types.GenesisState{
				ProjectChains: []types.ProjectChains{
					{
						ProjectID: 2,
					},
					{
						ProjectID: 4,
					},
				},
				Projects: []types.Project{
					sample.Project(r, 99),
					sample.Project(r, 88),
				},
				ProjectCounter: 100,
				TotalShares:    spntypes.TotalShareNumber,
			},
			errorMessage: "project id 2 doesn't exist for chains",
		},
		{
			name: "should prevent validation of genesis with duplicated projectChains",
			genState: &types.GenesisState{
				Projects: []types.Project{
					sample.Project(r, 0),
				},
				ProjectCounter: 1,
				ProjectChains: []types.ProjectChains{
					{
						ProjectID: 0,
					},
					{
						ProjectID: 0,
					},
				},
				TotalShares: spntypes.TotalShareNumber,
			},
			errorMessage: "duplicated index for projectChains",
		},
		{
			name: "should prevent validation of genesis with duplicated project",
			genState: &types.GenesisState{
				Projects: []types.Project{
					sample.Project(r, 0),
					sample.Project(r, 0),
				},
				ProjectCounter: 2,
				TotalShares:    spntypes.TotalShareNumber,
			},
			errorMessage: "duplicated id for project",
		},
		{
			name: "should prevent validation of genesis with invalid project count",
			genState: &types.GenesisState{
				Projects: []types.Project{
					sample.Project(r, 1),
				},
				ProjectCounter: 0,
				TotalShares:    spntypes.TotalShareNumber,
			},
			errorMessage: "project id should be lower or equal than the last id",
		},
		{
			name: "should prevent validation of genesis with invalid project",
			genState: &types.GenesisState{
				Projects: []types.Project{
					types.NewProject(
						0,
						invalidProjectName,
						sample.Uint64(r),
						sample.TotalSupply(r),
						sample.Metadata(r, 20),
						sample.Duration(r).Milliseconds(),
					),
				},
				ProjectCounter: 1,
				TotalShares:    spntypes.TotalShareNumber,
			},
			errorMessage: "invalid project 0: project name can only contain alphanumerical characters or hyphen",
		},
		{
			name: "should prevent validation of genesis with duplicated mainnetAccount",
			genState: &types.GenesisState{
				Projects: []types.Project{
					sample.Project(r, 0),
				},
				ProjectCounter: 1,
				MainnetAccounts: []types.MainnetAccount{
					{
						ProjectID: 0,
						Address:   "0",
					},
					{
						ProjectID: 0,
						Address:   "0",
					},
				},
				TotalShares: spntypes.TotalShareNumber,
			},
			errorMessage: "duplicated index for mainnetAccount",
		},
		{
			name: "should prevent validation of genesis with invalid allocations",
			genState: &types.GenesisState{
				Projects: []types.Project{
					{
						ProjectID:          0,
						ProjectName:        "test",
						CoordinatorID:      0,
						MainnetID:          0,
						MainnetInitialized: false,
						TotalSupply:        nil,
						AllocatedShares:    types.NewSharesFromCoins(tc2.Coins(t, fmt.Sprintf("%dstake", spntypes.TotalShareNumber+1))),
						Metadata:           nil,
					},
				},
				ProjectCounter: 1,
				MainnetAccounts: []types.MainnetAccount{
					{
						ProjectID: 0,
						Address:   "0",
					},
				},
				TotalShares: spntypes.TotalShareNumber,
			},
			errorMessage: "invalid project 0: more allocated shares than total shares",
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.errorMessage != "" {
				require.Error(t, err)
				require.Equal(t, tc.errorMessage, err.Error())
				return
			}
			require.NoError(t, err)

			projectIDMap := make(map[uint64]types.Shares)
			for _, elem := range tc.genState.Projects {
				projectIDMap[elem.ProjectID] = elem.AllocatedShares
			}
			shares := make(map[uint64]types.Shares)

			for _, acc := range tc.genState.MainnetAccounts {
				// check if the project exists for mainnet accounts
				_, ok := projectIDMap[acc.ProjectID]
				require.True(t, ok)

				// sum mainnet account shares
				if _, ok := shares[acc.ProjectID]; !ok {
					shares[acc.ProjectID] = types.EmptyShares()
				}
				shares[acc.ProjectID] = types.IncreaseShares(
					shares[acc.ProjectID],
					acc.Shares,
				)
			}

			for projectID, share := range projectIDMap {
				// check if the project shares is equal all accounts shares
				accShares, ok := shares[projectID]
				require.True(t, ok)
				isLowerEqual := accShares.IsAllLTE(share)
				require.True(t, isLowerEqual)
			}
		})
	}
}

func TestGenesisState_ValidateParams(t *testing.T) {
	for _, tc := range []struct {
		name     string
		genState types.GenesisState
		valid    bool
	}{
		{
			name: "should prevent validation of genesis with max total supply below min total supply",
			genState: types.GenesisState{
				Params: types.NewParams(
					types.DefaultMinTotalSupply,
					types.DefaultMinTotalSupply.Sub(sdkmath.OneInt()),
					types.DefaultProjectCreationFee,
					types.DefaultMaxMetadataLength,
				),
			},
			valid: false,
		},
		{
			name: "should prevent validation of genesis with valid parameters",
			genState: types.GenesisState{
				Params: types.NewParams(
					types.DefaultMinTotalSupply,
					types.DefaultMinTotalSupply.Add(sdkmath.OneInt()),
					types.DefaultProjectCreationFee,
					types.DefaultMaxMetadataLength,
				),
			},
			valid: true,
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
