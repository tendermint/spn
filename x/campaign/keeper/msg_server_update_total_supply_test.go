package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRequestAddAccount(t *testing.T) {
	var (
		coordAddr1                                         = sample.AccAddress()
		coordAddr2                                         = sample.AccAddress()
		campaignKeeper, _, campaignSrv, profileSrv, sdkCtx = setupMsgServer(t)
		ctx                                                = sdk.WrapSDKContext(sdkCtx)
	)

	// Create coordinators and set a campaign in the store with the first coordinator
	res, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	coordID := res.CoordinatorId
	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)

	campaign := sample.Campaign(0)
	campaign.CoordinatorID = coordID
	campaignKeeper.SetCampaign(sdkCtx, campaign)

	for _, tc := range []struct {
		name       string
		msg        types.MsgUpdateTotalSupply
		err        error
	}{
		{
			name: "update total shares",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:  0,
				Coordinator:   coordAddr1,
				TotalSupply:   sample.Coins(),
			},
		},
		{
			name: "can update total shares again",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:  0,
				Coordinator:   coordAddr1,
				TotalSupply:   sample.Coins(),
			},
		},
		{
			name: "campaign not found",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:  100,
				Coordinator:   coordAddr1,
				TotalSupply:   sample.Coins(),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "non existing coordinator",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:  0,
				Coordinator:   sample.AccAddress(),
				TotalSupply:   sample.Coins(),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "not the coordinator of the campaign",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:  0,
				Coordinator:  coordAddr2,
				TotalSupply:   sample.Coins(),
			},
			err: profiletypes.ErrCoordInvalid,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := campaignSrv.UpdateTotalSupply(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			campaign, found := campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)
			require.True(t, tc.msg.TotalSupply.IsEqual(campaign.TotalSupply))
		})
	}
}