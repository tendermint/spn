package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	ignterrors "github.com/ignite/modules/pkg/errors"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/project/types"
)

func Test_msgServer_UpdateSpecialAllocations(t *testing.T) {
	var (
		coordID            uint64
		coordAddr          = sample.Address(r)
		coordAddrNoProject = sample.Address(r)
		sdkCtx, tk, ts     = testkeeper.NewTestSetup(t)
		ctx                = sdk.WrapSDKContext(sdkCtx)
	)

	totalShares := tk.ProjectKeeper.GetTotalShares(sdkCtx)

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

	// utility to initialize a sample project with the data we need
	newProject := func(
		projectID uint64,
		as types.Shares,
		sa types.SpecialAllocations,
	) *types.Project {
		c := sample.Project(r, projectID)
		c.CoordinatorID = coordID
		c.AllocatedShares = as
		c.SpecialAllocations = sa
		return &c
	}

	// utility to initialize a sample chain with the data we need
	newChain := func(launchID uint64, launchTriggered bool) *launchtypes.Chain {
		chain := sample.Chain(r, launchID, coordID)
		chain.LaunchTriggered = launchTriggered
		if launchTriggered {
			chain.LaunchTime = sample.Time(r)
		}
		return &chain
	}

	// inputState represents input state for TDT cases/
	// non-defined values are not initialized
	// if a mainnet is defined, it is the mainnet of the project
	type inputState struct {
		project *types.Project
		mainnet *launchtypes.Chain
	}

	projectNoExistentMainnet := newProject(100, types.EmptyShares(), types.EmptySpecialAllocations())
	projectNoExistentMainnet.MainnetInitialized = true
	projectNoExistentMainnet.MainnetID = 100

	tests := []struct {
		name                    string
		msg                     types.MsgUpdateSpecialAllocations
		state                   inputState
		expectedAllocatedShares types.Shares
		err                     error
	}{
		{
			name: "should not update empty allocated shares",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				1,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				project: newProject(1, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet: nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
		},
		{
			name: "should increase allocated shares",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				2,
				types.NewSpecialAllocations(tc.Shares(t, "50foo"), tc.Shares(t, "30foo")),
			),
			state: inputState{
				project: newProject(2, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet: nil,
			},
			expectedAllocatedShares: tc.Shares(t, "80foo"),
		},
		{
			name: "should decrease allocated shares",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				3,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				project: newProject(3, tc.Shares(t, "200foo"),
					types.NewSpecialAllocations(tc.Shares(t, "50foo"), tc.Shares(t, "30foo")),
				),
				mainnet: nil,
			},
			expectedAllocatedShares: tc.Shares(t, "120foo"),
		},
		{
			name: "should update allocated shares relative to the new special allocations 1",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				4,
				types.NewSpecialAllocations(tc.Shares(t, "200foo"), tc.Shares(t, "200bar")),
			),
			state: inputState{
				project: newProject(4, tc.Shares(t, "1000foo,1000bar,1000baz"),
					types.NewSpecialAllocations(tc.Shares(t, "100foo,100bar"), tc.Shares(t, "100foo,100bar")),
				),
				mainnet: nil,
			},
			expectedAllocatedShares: tc.Shares(t, "1000foo,1000bar,1000baz"),
		},
		{
			name: "should update allocated shares relative to the new special allocations 2",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				5,
				types.NewSpecialAllocations(tc.Shares(t, "100foo"), tc.Shares(t, "300bar, 500baz")),
			),
			state: inputState{
				project: newProject(5, tc.Shares(t, "1000foo,1000bar,1000baz"),
					types.NewSpecialAllocations(tc.Shares(t, "100foo,100bar"), tc.Shares(t, "100foo,100bar")),
				),
				mainnet: nil,
			},
			expectedAllocatedShares: tc.Shares(t, "900foo,1100bar,1500baz"),
		},
		{
			name: "should update allocated shares relative to the new special allocations 3",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				6,
				types.NewSpecialAllocations(tc.Shares(t, "200foo"), tc.Shares(t, "500baz")),
			),
			state: inputState{
				project: newProject(6, tc.Shares(t, "1000foo,200bar"),
					types.NewSpecialAllocations(tc.Shares(t, "100foo,100bar"), tc.Shares(t, "100foo,100bar")),
				),
				mainnet: nil,
			},
			expectedAllocatedShares: tc.Shares(t, "1000foo,500baz"),
		},
		{
			name: "should set allocated shares to empty if allocated shares are special allocations and special allocations are removed",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				7,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				project: newProject(7, tc.Shares(t, "200foo,200bar"),
					types.NewSpecialAllocations(tc.Shares(t, "100foo,100bar"), tc.Shares(t, "100foo,100bar")),
				),
				mainnet: nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
		},
		{
			name: "should allow updating special allocations if mainnet is not initialized",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				8,
				types.NewSpecialAllocations(tc.Shares(t, "100foo"), tc.Shares(t, "300bar, 500baz")),
			),
			state: inputState{
				project: newProject(8, tc.Shares(t, "1000foo,1000bar,1000baz"),
					types.NewSpecialAllocations(tc.Shares(t, "100foo,100bar"), tc.Shares(t, "100foo,100bar")),
				),
				mainnet: newChain(1, false),
			},
			expectedAllocatedShares: tc.Shares(t, "900foo,1100bar,1500baz"),
		},
		{
			name: "should allow to allocated all shares to special allocations",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				9,
				types.NewSpecialAllocations(tc.Shares(t, fmt.Sprintf("%dfoo", totalShares)), tc.Shares(t, fmt.Sprintf("%dbar", totalShares))),
			),
			state: inputState{
				project: newProject(9, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet: nil,
			},
			expectedAllocatedShares: tc.Shares(t, fmt.Sprintf("%dfoo,%dbar", totalShares, totalShares)),
		},
		{
			name: "should fail if project does not exist",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				10000,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				project: nil,
				mainnet: nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     types.ErrProjectNotFound,
		},
		{
			name: "should fail if the coordinator does not exist",
			msg: *types.NewMsgUpdateSpecialAllocations(
				sample.Address(r),
				50,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				project: newProject(50, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet: nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should fail if the signer is not the coordinator of the project",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddrNoProject,
				51,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				project: newProject(51, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet: nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     profiletypes.ErrCoordInvalid,
		},
		{
			name: "should fail if mainnet launch is triggered",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				52,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				project: newProject(52, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet: newChain(50, true),
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     types.ErrMainnetLaunchTriggered,
		},
		{
			name: "should fail if total shares is reached",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				53,
				types.NewSpecialAllocations(tc.Shares(t, fmt.Sprintf("%dfoo", totalShares)), tc.Shares(t, "1foo")),
			),
			state: inputState{
				project: newProject(53, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet: nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     types.ErrTotalSharesLimit,
		},
		{
			name: "should fail with critical error if current special allocations are bigger than allocated shares",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				54,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				project: newProject(54, tc.Shares(t, "1000foo"),
					types.NewSpecialAllocations(tc.Shares(t, "600foo"), tc.Shares(t, "600foo")),
				),
				mainnet: nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     ignterrors.ErrCritical,
		},
		{
			name: "should trigger a critical error when updating a project with a non-existent initialize mainnet ",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				100,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				project: projectNoExistentMainnet,
				mainnet: nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     ignterrors.ErrCritical,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the project if defined
			if tt.state.project != nil {

				// link mainnet to project if defined
				if tt.state.mainnet != nil {
					tt.state.mainnet.IsMainnet = true
					tt.state.mainnet.ProjectID = tt.state.project.ProjectID
					tt.state.project.MainnetInitialized = true
					tt.state.project.MainnetID = tt.state.mainnet.LaunchID

					tk.LaunchKeeper.SetChain(sdkCtx, *tt.state.mainnet)
				}

				tk.ProjectKeeper.SetProject(sdkCtx, *tt.state.project)
			}

			_, err := ts.ProjectSrv.UpdateSpecialAllocations(ctx, &tt.msg)

			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}

			// fetch project
			prjt, found := tk.ProjectKeeper.GetProject(sdkCtx, tt.msg.ProjectID)
			require.True(t, found)

			// check genesis distribution
			gdExpected := tt.msg.SpecialAllocations.GenesisDistribution
			gdGot := prjt.SpecialAllocations.GenesisDistribution
			require.True(t, types.IsEqualShares(gdExpected, gdGot),
				"invalid genesis distribution, expected: %s, got: %s", gdExpected.String(), gdGot.String(),
			)

			// check claimable airdrop
			caExpected := tt.msg.SpecialAllocations.ClaimableAirdrop
			caGot := prjt.SpecialAllocations.ClaimableAirdrop
			require.True(t, types.IsEqualShares(caExpected, caGot),
				"invalid claimable airdrop, expected: %s, got: %s", caExpected.String(), caGot.String(),
			)

			// check allocated shares
			asExpected := tt.expectedAllocatedShares
			asGot := prjt.AllocatedShares
			require.True(t, types.IsEqualShares(asExpected, asGot),
				"invalid allocated shares, expected: %s, got: %s", asExpected.String(), asGot.String(),
			)

			// no other values should be edited
			prjt.SpecialAllocations = types.EmptySpecialAllocations()
			tt.state.project.SpecialAllocations = types.EmptySpecialAllocations()
			prjt.AllocatedShares = types.EmptyShares()
			tt.state.project.AllocatedShares = types.EmptyShares()
			require.EqualValues(t, *tt.state.project, prjt)
		})
	}
}
