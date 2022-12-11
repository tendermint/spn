package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/project/types"
)

func TestMsgUpdateTotalSupply(t *testing.T) {
	var (
		coordID        uint64
		coordAddr1     = sample.Address(r)
		coordAddr2     = sample.Address(r)
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
	)

	t.Run("should allow creating coordinators", func(t *testing.T) {
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddr1,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
		coordID = res.CoordinatorID
		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddr2,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
	})

	// Set a regular project and a project with an already initialized mainnet
	project := sample.Project(r, 0)
	project.CoordinatorID = coordID
	tk.ProjectKeeper.SetProject(sdkCtx, project)

	project = sample.Project(r, 1)
	project.CoordinatorID = coordID
	project.MainnetInitialized = true
	tk.ProjectKeeper.SetProject(sdkCtx, project)

	for _, tc := range []struct {
		name string
		msg  types.MsgUpdateTotalSupply
		err  error
	}{
		{
			name: "should update total supply",
			msg: types.MsgUpdateTotalSupply{
				ProjectID:         0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
		},
		{
			name: "should allow update total supply again",
			msg: types.MsgUpdateTotalSupply{
				ProjectID:         0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
		},
		{
			name: "should fail if project not found",
			msg: types.MsgUpdateTotalSupply{
				ProjectID:         100,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail with non existing coordinator",
			msg: types.MsgUpdateTotalSupply{
				ProjectID:         0,
				Coordinator:       sample.Address(r),
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should fail if coordinator is not associated with project",
			msg: types.MsgUpdateTotalSupply{
				ProjectID:         0,
				Coordinator:       coordAddr2,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "cannot update total supply when mainnet is initialized",
			msg: types.MsgUpdateTotalSupply{
				ProjectID:         1,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: types.ErrMainnetInitialized,
		},
		{
			name: "should fail if total supply outside of valid range",
			msg: types.MsgUpdateTotalSupply{
				ProjectID:         0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.CoinsWithRange(r, 10, 20),
			},
			err: types.ErrInvalidTotalSupply,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousTotalSupply sdk.Coins
			if tc.err == nil {
				project, found := tk.ProjectKeeper.GetProject(sdkCtx, tc.msg.ProjectID)
				require.True(t, found)
				previousTotalSupply = project.TotalSupply
			}

			_, err := ts.ProjectSrv.UpdateTotalSupply(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			project, found := tk.ProjectKeeper.GetProject(sdkCtx, tc.msg.ProjectID)
			require.True(t, found)
			require.True(t, project.TotalSupply.IsEqual(
				types.UpdateTotalSupply(previousTotalSupply, tc.msg.TotalSupplyUpdate),
			))
		})
	}
}
