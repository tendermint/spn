package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/project/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func initCreationFeeAndFundCoordAccounts(
	t *testing.T,
	keeper *keeper.Keeper,
	bk bankkeeper.Keeper,
	sdkCtx sdk.Context,
	fee sdk.Coins,
	numCreations int64,
	addrs ...string,
) {
	// set fee param to `coins`
	params := keeper.GetParams(sdkCtx)
	params.ProjectCreationFee = fee
	keeper.SetParams(sdkCtx, params)

	coins := sdk.NewCoins()
	for _, coin := range fee {
		coin.Amount = coin.Amount.MulRaw(numCreations)
		coins = coins.Add(coin)
	}

	t.Run("should add coins to balance of each coordinator address", func(t *testing.T) {
		for _, addr := range addrs {
			accAddr, err := sdk.AccAddressFromBech32(addr)
			require.NoError(t, err)
			err = bk.MintCoins(sdkCtx, types.ModuleName, coins)
			require.NoError(t, err)
			err = bk.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, accAddr, coins)
			require.NoError(t, err)
		}
	})
}

func TestMsgCreateProject(t *testing.T) {
	var (
		coordAddrs         = make([]string, 3)
		coordMap           = make(map[string]uint64)
		sdkCtx, tk, ts     = testkeeper.NewTestSetup(t)
		ctx                = sdk.WrapSDKContext(sdkCtx)
		projectCreationFee = sample.Coins(r)
		maxMetadataLength  = tk.ProjectKeeper.MaxMetadataLength(sdkCtx)
	)

	t.Run("should allow creation of coordinators", func(t *testing.T) {
		for i := range coordAddrs {
			addr := sample.Address(r)
			coordAddrs[i] = addr
			coordMap[addr], _ = ts.CreateCoordinatorWithAddr(ctx, r, addr)
		}
	})

	// assign random sdk.Coins to `projectCreationFee` param and provide balance to coordinators
	// coordAddrs[2] is not funded
	initCreationFeeAndFundCoordAccounts(t, tk.ProjectKeeper, tk.BankKeeper, sdkCtx, projectCreationFee, 1, coordAddrs[:2]...)

	for _, tc := range []struct {
		name       string
		msg        types.MsgCreateProject
		expectedID uint64
		err        error
	}{
		{
			name: "should allow create a project 1",
			msg: types.MsgCreateProject{
				ProjectName: sample.ProjectName(r),
				Coordinator: coordAddrs[0],
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, 20),
			},
			expectedID: uint64(0),
		},
		{
			name: "should allow create a project from a different coordinator",
			msg: types.MsgCreateProject{
				ProjectName: sample.ProjectName(r),
				Coordinator: coordAddrs[1],
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, 20),
			},
			expectedID: uint64(1),
		},
		{
			name: "should allow create a project from a non existing coordinator",
			msg: types.MsgCreateProject{
				ProjectName: sample.ProjectName(r),
				Coordinator: sample.Address(r),
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should allow create a project with an invalid token supply",
			msg: types.MsgCreateProject{
				ProjectName: sample.ProjectName(r),
				Coordinator: coordAddrs[0],
				TotalSupply: sample.CoinsWithRange(r, 10, 20),
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "should fail for insufficient balance to cover creation fee",
			msg: types.MsgCreateProject{
				ProjectName: sample.ProjectName(r),
				Coordinator: coordAddrs[2],
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrFundCommunityPool,
		},
		{
			name: "should fail when the change had too long metadata",
			msg: types.MsgCreateProject{
				Coordinator: sample.Address(r),
				ProjectName: sample.ProjectName(r),
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, maxMetadataLength+1),
			},
			err: types.ErrInvalidMetadataLength,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// get account initial balance
			accAddr, err := sdk.AccAddressFromBech32(tc.msg.Coordinator)
			require.NoError(t, err)
			preBalance := tk.BankKeeper.SpendableCoins(sdkCtx, accAddr)

			got, err := ts.ProjectSrv.CreateProject(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectedID, got.ProjectID)
			project, found := tk.ProjectKeeper.GetProject(sdkCtx, got.ProjectID)
			require.True(t, found)
			require.EqualValues(t, got.ProjectID, project.ProjectID)
			require.EqualValues(t, tc.msg.ProjectName, project.ProjectName)
			require.EqualValues(t, coordMap[tc.msg.Coordinator], project.CoordinatorID)
			require.False(t, project.MainnetInitialized)
			require.True(t, tc.msg.TotalSupply.IsEqual(project.TotalSupply))
			require.EqualValues(t, types.Shares(nil), project.AllocatedShares)

			// Empty list of project chains
			projectChains, found := tk.ProjectKeeper.GetProjectChains(sdkCtx, got.ProjectID)
			require.True(t, found)
			require.EqualValues(t, got.ProjectID, projectChains.ProjectID)
			require.Empty(t, projectChains.Chains)

			// check fee deduction
			postBalance := tk.BankKeeper.SpendableCoins(sdkCtx, accAddr)
			require.True(t, preBalance.Sub(projectCreationFee...).IsEqual(postBalance))
		})
	}
}
