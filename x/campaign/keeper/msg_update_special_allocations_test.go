package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/errors"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func Test_msgServer_UpdateSpecialAllocations(t *testing.T) {
	var (
		coordAddr           = sample.Address(r)
		coordAddrNoCampaign = sample.Address(r)
		sdkCtx, tk, ts      = testkeeper.NewTestSetup(t)
		ctx                 = sdk.WrapSDKContext(sdkCtx)
	)

	totalShares := tk.CampaignKeeper.GetTotalShares(sdkCtx)

	// Create two coordinators
	res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)
	coordID := res.CoordinatorID
	res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddrNoCampaign,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)

	// utility to initialize a sample campaign with the data we need
	newCampaign := func(
		campaignID uint64,
		as types.Shares,
		sa types.SpecialAllocations,
	) *types.Campaign {
		c := sample.Campaign(r, campaignID)
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
			chain.LaunchTimestamp = sample.Duration(r).Milliseconds()
		}
		return &chain
	}

	// inputState represents input state for TDT cases/
	// non-defined values are not initialized
	// if a mainnet is defined, it is the mainnet of the campaign
	type inputState struct {
		campaign *types.Campaign
		mainnet  *launchtypes.Chain
	}

	campaignNoExistentMainnet := newCampaign(100, types.EmptyShares(), types.EmptySpecialAllocations())
	campaignNoExistentMainnet.MainnetInitialized = true
	campaignNoExistentMainnet.MainnetID = 100

	tests := []struct {
		name                    string
		msg                     types.MsgUpdateSpecialAllocations
		state                   inputState
		expectedAllocatedShares types.Shares
		err                     error
	}{
		{
			name: "empty should not update empty allocated shares",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				1,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				campaign: newCampaign(1, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet:  nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
		},
		{
			name: "new special allocations should increase allocated shares",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				2,
				types.NewSpecialAllocations(tc.Shares(t, "50foo"), tc.Shares(t, "30foo")),
			),
			state: inputState{
				campaign: newCampaign(2, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet:  nil,
			},
			expectedAllocatedShares: tc.Shares(t, "80foo"),
		},
		{
			name: "removing special allocations should decrease allocated shares",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				3,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				campaign: newCampaign(3, tc.Shares(t, "200foo"),
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
				campaign: newCampaign(4, tc.Shares(t, "1000foo,1000bar,1000baz"),
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
				campaign: newCampaign(5, tc.Shares(t, "1000foo,1000bar,1000baz"),
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
				campaign: newCampaign(6, tc.Shares(t, "1000foo,200bar"),
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
				campaign: newCampaign(7, tc.Shares(t, "200foo,200bar"),
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
				campaign: newCampaign(8, tc.Shares(t, "1000foo,1000bar,1000baz"),
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
				campaign: newCampaign(9, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet:  nil,
			},
			expectedAllocatedShares: tc.Shares(t, fmt.Sprintf("%dfoo,%dbar", totalShares, totalShares)),
		},
		{
			name: "should fail if campaign doesn't exist",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				10000,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				campaign: nil,
				mainnet:  nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     types.ErrCampaignNotFound,
		},
		{
			name: "should fail if the coordinator doesn't exist",
			msg: *types.NewMsgUpdateSpecialAllocations(
				sample.Address(r),
				50,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				campaign: newCampaign(50, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet:  nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should fail if the signer is not the coordinator of the campaign",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddrNoCampaign,
				51,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				campaign: newCampaign(51, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet:  nil,
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
				campaign: newCampaign(52, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet:  newChain(50, true),
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
				campaign: newCampaign(53, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet:  nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     types.ErrTotalSharesLimit,
		},
		{
			name: "should fails with critical error if current special allocations are bigger than allocated shares",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				54,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				campaign: newCampaign(54, tc.Shares(t, "1000foo"),
					types.NewSpecialAllocations(tc.Shares(t, "600foo"), tc.Shares(t, "600foo")),
				),
				mainnet: nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     ignterrors.ErrCritical,
		},
		{
			name: "updating a campaign with a non-existent initialize mainnet should trigger a critical error",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				100,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				campaign: campaignNoExistentMainnet,
				mainnet:  nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     ignterrors.ErrCritical,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the campaign if defined
			if tt.state.campaign != nil {

				// link mainnet to campaign if defined
				if tt.state.mainnet != nil {
					tt.state.mainnet.IsMainnet = true
					tt.state.mainnet.CampaignID = tt.state.campaign.CampaignID
					tt.state.campaign.MainnetInitialized = true
					tt.state.campaign.MainnetID = tt.state.mainnet.LaunchID

					tk.LaunchKeeper.SetChain(sdkCtx, *tt.state.mainnet)
				}

				tk.CampaignKeeper.SetCampaign(sdkCtx, *tt.state.campaign)
			}

			_, err := ts.CampaignSrv.UpdateSpecialAllocations(ctx, &tt.msg)

			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}

			// fetch campaign
			camp, found := tk.CampaignKeeper.GetCampaign(sdkCtx, tt.msg.CampaignID)
			require.True(t, found)

			// check genesis distribution
			gdExpected := tt.msg.SpecialAllocations.GenesisDistribution
			gdGot := camp.SpecialAllocations.GenesisDistribution
			require.True(t, types.IsEqualShares(gdExpected, gdGot),
				"invalid genesis distribution, expected: %s, got: %s", gdExpected.String(), gdGot.String(),
			)

			// check claimable airdrop
			caExpected := tt.msg.SpecialAllocations.ClaimableAirdrop
			caGot := camp.SpecialAllocations.ClaimableAirdrop
			require.True(t, types.IsEqualShares(caExpected, caGot),
				"invalid claimable airdrop, expected: %s, got: %s", caExpected.String(), caGot.String(),
			)

			// check allocated shares
			asExpected := tt.expectedAllocatedShares
			asGot := camp.AllocatedShares
			require.True(t, types.IsEqualShares(asExpected, asGot),
				"invalid allocated shares, expected: %s, got: %s", asExpected.String(), asGot.String(),
			)

			// no other values should be edited
			camp.SpecialAllocations = types.EmptySpecialAllocations()
			tt.state.campaign.SpecialAllocations = types.EmptySpecialAllocations()
			camp.AllocatedShares = types.EmptyShares()
			tt.state.campaign.AllocatedShares = types.EmptyShares()
			require.EqualValues(t, *tt.state.campaign, camp)
		})
	}
}
