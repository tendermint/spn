package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgMintVouchers(t *testing.T) {
	var (
		sdkCtx, tk, ts  = testkeeper.NewTestSetup(t)
		ctx             = sdk.WrapSDKContext(sdkCtx)
		coordID         uint64
		coord           = sample.Address(r)
		coordNoProject = sample.Address(r)

		shares, _    = types.NewShares("1000foo,500bar,300foobar")
		sharesTooBig = types.NewSharesFromCoins(sdk.NewCoins(
			sdk.NewCoin("foo", sdkmath.NewInt(spntypes.TotalShareNumber+1)),
		))
	)

	t.Run("should allow creation of coordinators", func(t *testing.T) {
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coord,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
		coordID = res.CoordinatorID
		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordNoProject,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
	})

	// Set project
	project := sample.Project(r, 0)
	project.CoordinatorID = coordID
	project.ProjectID = tk.ProjectKeeper.AppendProject(sdkCtx, project)

	for _, tc := range []struct {
		name string
		msg  types.MsgMintVouchers
		err  error
	}{
		{
			name: "should allow minting  vouchers",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				ProjectID:  0,
				Shares:      shares,
			},
		},
		{
			name: "should allow minting same vouchers again",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				ProjectID:  0,
				Shares:      shares,
			},
		},
		{
			name: "should allow minting other vouchers",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				ProjectID:  0,
				Shares:      sample.Shares(r),
			},
		},
		{
			name: "should not mint more than total shares",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				ProjectID:  0,
				Shares:      sharesTooBig,
			},
			err: types.ErrTotalSharesLimit,
		},
		{
			name: "should fail with non existing project",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				ProjectID:  1000,
				Shares:      shares,
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail with non existing coordinator",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				ProjectID:  0,
				Shares:      shares,
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should fail with invalid coordinator",
			msg: types.MsgMintVouchers{
				Coordinator: coordNoProject,
				ProjectID:  0,
				Shares:      shares,
			},
			err: profiletypes.ErrCoordInvalid,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousProject types.Project
			var previousBalance sdk.Coins

			coordAddr, err := sdk.AccAddressFromBech32(tc.msg.Coordinator)
			require.NoError(t, err)

			// Get values before message execution
			if tc.err == nil {
				var found bool
				previousProject, found = tk.ProjectKeeper.GetProject(sdkCtx, tc.msg.ProjectID)
				require.True(t, found)

				previousBalance = tk.BankKeeper.GetAllBalances(sdkCtx, coordAddr)
			}

			// Execute message
			_, err = ts.ProjectSrv.MintVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			project, found := tk.ProjectKeeper.GetProject(sdkCtx, tc.msg.ProjectID)
			require.True(t, found)

			// Allocated shares of the project must be increased
			expectedShares := types.IncreaseShares(previousProject.AllocatedShares, tc.msg.Shares)
			require.True(t, types.IsEqualShares(expectedShares, project.AllocatedShares))

			// Check coordinator balance
			minted, err := types.SharesToVouchers(tc.msg.Shares, tc.msg.ProjectID)
			require.NoError(t, err)
			balance := tk.BankKeeper.GetAllBalances(sdkCtx, coordAddr)
			require.True(t, balance.IsEqual(previousBalance.Add(minted...)))
		})
	}
}
