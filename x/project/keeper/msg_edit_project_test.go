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

func TestMsgUpdateProjectName(t *testing.T) {
	var (
		coordAddr          = sample.Address(r)
		coordAddrNoProject = sample.Address(r)
		project            = sample.Project(r, 0)

		sdkCtx, tk, ts    = testkeeper.NewTestSetup(t)
		ctx               = sdk.WrapSDKContext(sdkCtx)
		maxMetadataLength = tk.ProjectKeeper.MaxMetadataLength(sdkCtx)
	)

	t.Run("should allow creation of coordinators", func(t *testing.T) {
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddr,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
		project.CoordinatorID = res.CoordinatorID
		project.ProjectID = tk.ProjectKeeper.AppendProject(sdkCtx, project)

		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddrNoProject,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
	})

	for _, tc := range []struct {
		name string
		msg  types.MsgEditProject
		err  error
	}{
		{
			name: "should allow edit name and metadata",
			msg: types.MsgEditProject{
				Coordinator: coordAddr,
				ProjectID:   project.ProjectID,
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "should allow edit name",
			msg: types.MsgEditProject{
				Coordinator: coordAddr,
				ProjectID:   project.ProjectID,
				Name:        sample.ProjectName(r),
				Metadata:    []byte{},
			},
		},
		{
			name: "should allow edit metadata",
			msg: types.MsgEditProject{
				Coordinator: coordAddr,
				ProjectID:   project.ProjectID,
				Name:        "",
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "should fail if invalid project id",
			msg: types.MsgEditProject{
				Coordinator: coordAddr,
				ProjectID:   100,
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail with invalid coordinator address",
			msg: types.MsgEditProject{
				Coordinator: sample.Address(r),
				ProjectID:   project.ProjectID,
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should fail with wrong coordinator id",
			msg: types.MsgEditProject{
				Coordinator: coordAddrNoProject,
				ProjectID:   project.ProjectID,
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "should fail when the change had too long metadata",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, maxMetadataLength+1),
			},
			err: types.ErrInvalidMetadataLength,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			previousProject, found := tk.ProjectKeeper.GetProject(sdkCtx, tc.msg.ProjectID)
			_, err := ts.ProjectSrv.EditProject(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			project, found := tk.ProjectKeeper.GetProject(sdkCtx, tc.msg.ProjectID)
			require.True(t, found)

			if len(tc.msg.Name) > 0 {
				require.EqualValues(t, tc.msg.Name, project.ProjectName)
			} else {
				require.EqualValues(t, previousProject.ProjectName, project.ProjectName)
			}

			if len(tc.msg.Metadata) > 0 {
				require.EqualValues(t, tc.msg.Metadata, project.Metadata)
			} else {
				require.EqualValues(t, previousProject.Metadata, project.Metadata)
			}
		})
	}
}
