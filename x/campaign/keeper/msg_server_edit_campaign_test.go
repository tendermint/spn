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
	campaign.CampaignID = campaignKeeper.AppendCampaign(sdkCtx, campaign)

	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddrNoCampaign,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)

	for _, tc := range []struct {
		name string
		msg  types.MsgEditCampaign
		err  error
	}{
		{
			name: "invalid campaign id",
			msg: types.MsgEditCampaign{
				Coordinator: coordAddr,
				CampaignID:  100,
				Name:        "new_name",
				Metadata:    sample.Metadata(20),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "invalid coordinator address",
			msg: types.MsgEditCampaign{
				Coordinator: sample.Address(),
				CampaignID:  campaign.CampaignID,
				Name:        "new_name",
				Metadata:    sample.Metadata(20),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "wrong coordinator id",
			msg: types.MsgEditCampaign{
				Coordinator: coordAddrNoCampaign,
				CampaignID:  campaign.CampaignID,
				Name:        "new_name",
				Metadata:    sample.Metadata(20),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "valid transaction",
			msg: types.MsgEditCampaign{
				Coordinator: coordAddr,
				CampaignID:  campaign.CampaignID,
				Name:        "new_name",
				Metadata:    sample.Metadata(20),
			},
		},
		{
			name: "valid transaction - unmodified metadata",
			msg: types.MsgEditCampaign{
				Coordinator: coordAddr,
				CampaignID:  campaign.CampaignID,
				Name:        "new_name",
				Metadata:    []byte{},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			previousCampaign, found := campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			_, err := campaignSrv.EditCampaign(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			campaign, found := campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)
			require.Equal(t, tc.msg.Name, campaign.CampaignName)
			if len(tc.msg.Metadata) > 0 {
				require.EqualValues(t, tc.msg.Metadata, campaign.Metadata)
			} else {
				require.EqualValues(t, previousCampaign.Metadata, campaign.Metadata)
			}
		})
	}
}
