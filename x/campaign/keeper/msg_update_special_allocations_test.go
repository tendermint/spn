package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"testing"
)

func Test_msgServer_UpdateSpecialAllocations(t *testing.T) {
	var (
		coordAddr           = sample.Address(r)
		coordAddrNoCampaign = sample.Address(r)
		sdkCtx, tk, ts      = testkeeper.NewTestSetup(t)
		ctx                 = sdk.WrapSDKContext(sdkCtx)
	)

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

	tests := []struct {
		name                    string
		msg                     types.MsgUpdateSpecialAllocations
		state                   inputState
		expectedAllocatedShares types.Shares
		err                     error
	}{
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
				1,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				campaign: newCampaign(1, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet:  nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should fail if the signer is not the coordinator of the campaign",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddrNoCampaign,
				2,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				campaign: newCampaign(2, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet:  nil,
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     profiletypes.ErrCoordInvalid,
		},
		{
			name: "should fail if mainnet launch is triggered",
			msg: *types.NewMsgUpdateSpecialAllocations(
				coordAddr,
				3,
				types.EmptySpecialAllocations(),
			),
			state: inputState{
				campaign: newCampaign(3, types.EmptyShares(), types.EmptySpecialAllocations()),
				mainnet:  newChain(1, true),
			},
			expectedAllocatedShares: types.EmptyShares(),
			err:                     types.ErrMainnetLaunchTriggered,
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
			require.EqualValues(t, tt.state.campaign, camp)
		})
	}
}
