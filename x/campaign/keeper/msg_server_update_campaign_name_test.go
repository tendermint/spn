package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCampaignName(t *testing.T) {
	var (
		coordAddr           = sample.Address()
		coordAddrNoCampaign = sample.Address()
		campaign            = sample.Campaign(0)

		campaignKeeper, _, _, _, campaignSrv, profileSrv, sdkCtx = setupMsgServer(t)
		ctx                                                      = sdk.WrapSDKContext(sdkCtx)
	)
	res, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	campaign.CoordinatorID = res.CoordinatorID
	campaign.Id = campaignKeeper.AppendCampaign(sdkCtx, campaign)

	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddrNoCampaign,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)

	for _, tc := range []struct {
		name       string
		msg        types.MsgUpdateCampaignName
		expectedID uint64
		err        error
	}{
		{
			name: "invalid campaign id",
			msg: types.MsgUpdateCampaignName{
				Coordinator: coordAddr,
				CampaignID:  100,
				Name:        "new_name",
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "invalid coordinator address",
			msg: types.MsgUpdateCampaignName{
				Coordinator: sample.Address(),
				CampaignID:  campaign.Id,
				Name:        "new_name",
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "wrong coordinator id",
			msg: types.MsgUpdateCampaignName{
				Coordinator: coordAddrNoCampaign,
				CampaignID:  campaign.Id,
				Name:        "new_name",
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "valid message",
			msg: types.MsgUpdateCampaignName{
				Coordinator: coordAddr,
				CampaignID:  campaign.Id,
				Name:        "new_name",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := campaignSrv.UpdateCampaignName(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			campaign, found := campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)
			require.Equal(t, tc.msg.Name, campaign.CampaignName)
		})
	}
}
