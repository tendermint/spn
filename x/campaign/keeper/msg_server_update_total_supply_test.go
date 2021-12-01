package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateTotalSupply(t *testing.T) {
	var (
		coordAddr1                                               = sample.Address()
		coordAddr2                                               = sample.Address()
		campaignKeeper, _, _, _, campaignSrv, profileSrv, sdkCtx = setupMsgServer(t)
		ctx                                                      = sdk.WrapSDKContext(sdkCtx)
	)

	// Create coordinators
	res, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	coordID := res.CoordinatorID
	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)

	// Set a regular campaign and a campaign with an already initialized mainnet
	campaign := sample.Campaign(0)
	campaign.CoordinatorID = coordID
	campaignKeeper.SetCampaign(sdkCtx, campaign)

	campaign = sample.Campaign(1)
	campaign.CoordinatorID = coordID
	campaign.MainnetInitialized = true
	campaignKeeper.SetCampaign(sdkCtx, campaign)

	for _, tc := range []struct {
		name string
		msg  types.MsgUpdateTotalSupply
		err  error
	}{
		{
			name: "update total supply",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.Coins(),
			},
		},
		{
			name: "can update total supply again",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.Coins(),
			},
		},
		{
			name: "campaign not found",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        100,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.Coins(),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "non existing coordinator",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        0,
				Coordinator:       sample.Address(),
				TotalSupplyUpdate: sample.Coins(),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "not the coordinator of the campaign",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        0,
				Coordinator:       coordAddr2,
				TotalSupplyUpdate: sample.Coins(),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "cannot update total supply when mainnet is initialized",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        1,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.Coins(),
			},
			err: types.ErrMainnetInitialized,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousTotalSupply sdk.Coins
			if tc.err == nil {
				campaign, found := campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)
				previousTotalSupply = campaign.TotalSupply
			}

			_, err := campaignSrv.UpdateTotalSupply(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			campaign, found := campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)
			require.True(t, campaign.TotalSupply.IsEqual(
				types.UpdateTotalSupply(previousTotalSupply, tc.msg.TotalSupplyUpdate),
			))
		})
	}
}
