package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/project/types"
)

func TestMsgInitializeMainnet(t *testing.T) {
	var (
		coordID                     uint64
		projectID                   uint64 = 0
		projectMainnetInitializedID uint64 = 1
		projectIncorrectCoordID     uint64 = 2
		projectEmptySupplyID        uint64 = 3
		coordAddr                          = sample.Address(r)
		coordAddrNoProject                 = sample.Address(r)

		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
	)

	t.Run("should allow creation of coordinators", func(t *testing.T) {
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddr,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
		coordID = res.CoordinatorID
		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddrNoProject,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
	})

	project := sample.Project(r, projectID)
	project.CoordinatorID = coordID
	tk.ProjectKeeper.SetProject(sdkCtx, project)

	projectMainnetInitialized := sample.Project(r, projectMainnetInitializedID)
	projectMainnetInitialized.CoordinatorID = coordID
	projectMainnetInitialized.MainnetInitialized = true
	tk.ProjectKeeper.SetProject(sdkCtx, projectMainnetInitialized)

	projectEmptySupply := sample.Project(r, projectEmptySupplyID)
	projectEmptySupply.CoordinatorID = coordID
	projectEmptySupply.TotalSupply = sdk.NewCoins()
	tk.ProjectKeeper.SetProject(sdkCtx, projectEmptySupply)

	projectIncorrectCoord := sample.Project(r, projectIncorrectCoordID)
	projectIncorrectCoord.CoordinatorID = coordID
	tk.ProjectKeeper.SetProject(sdkCtx, projectIncorrectCoord)

	for _, tc := range []struct {
		name string
		msg  types.MsgInitializeMainnet
		err  error
	}{
		{
			name: "should allow initialize mainnet",
			msg: types.MsgInitializeMainnet{
				ProjectID:      projectID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
		},
		{
			name: "should fail if project not found",
			msg: types.MsgInitializeMainnet{
				ProjectID:      1000,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail if mainnet already initialized",
			msg: types.MsgInitializeMainnet{
				ProjectID:      projectMainnetInitializedID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrMainnetInitialized,
		},
		{
			name: "should fail if project has empty supply",
			msg: types.MsgInitializeMainnet{
				ProjectID:      projectEmptySupplyID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "should fail with non-existent coordinator",
			msg: types.MsgInitializeMainnet{
				ProjectID:      projectIncorrectCoordID,
				Coordinator:    sample.Address(r),
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should fail with invalid coordinator",
			msg: types.MsgInitializeMainnet{
				ProjectID:      projectIncorrectCoordID,
				Coordinator:    coordAddrNoProject,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: profiletypes.ErrCoordInvalid,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ts.ProjectSrv.InitializeMainnet(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			project, found := tk.ProjectKeeper.GetProject(sdkCtx, tc.msg.ProjectID)
			require.True(t, found)
			require.True(t, project.MainnetInitialized)
			require.EqualValues(t, res.MainnetID, project.MainnetID)

			// Chain is in launch module
			chain, found := tk.LaunchKeeper.GetChain(sdkCtx, project.MainnetID)
			require.True(t, found)
			require.True(t, chain.HasProject)
			require.True(t, chain.IsMainnet)
			require.EqualValues(t, tc.msg.ProjectID, chain.ProjectID)

			// Mainnet ID is listed in project chains
			projectChains, found := tk.ProjectKeeper.GetProjectChains(sdkCtx, tc.msg.ProjectID)
			require.True(t, found)
			require.Contains(t, projectChains.Chains, project.MainnetID)
		})
	}
}
