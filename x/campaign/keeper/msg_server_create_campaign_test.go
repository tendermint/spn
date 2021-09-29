package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgCreateCampaign(t *testing.T) {
	var (
		coordAddr1                                         = sample.AccAddress()
		coordAddr2                                         = sample.AccAddress()
		campaignKeeper, _, campaignSrv, profileSrv, sdkCtx = setupMsgServer(t)
		ctx                                                = sdk.WrapSDKContext(sdkCtx)
	)

	// Create coordinators
	coordMap := make(map[string]uint64)
	res, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	coordMap[coordAddr1] = res.CoordinatorId
	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	coordMap[coordAddr2] = res.CoordinatorId

	for _, tc := range []struct {
		name       string
		msg        types.MsgCreateCampaign
		expectedID uint64
		err        error
	}{
		{
			name: "create a campaign 1",
			msg: types.MsgCreateCampaign{
				CampaignName:  sample.CampaignName(),
				Coordinator:   coordAddr1,
				TotalSupply:   sample.Coins(),
				DynamicShares: false,
			},
			expectedID: uint64(1),
		},
		{
			name: "create a campaign 2 with dynamic shares",
			msg: types.MsgCreateCampaign{
				CampaignName:  sample.CampaignName(),
				Coordinator:   coordAddr1,
				TotalSupply:   sample.Coins(),
				DynamicShares: true,
			},
			expectedID: uint64(2),
		},
		{
			name: "create a campaign from a different coordinator",
			msg: types.MsgCreateCampaign{
				CampaignName:  sample.CampaignName(),
				Coordinator:   coordAddr2,
				TotalSupply:   sample.Coins(),
				DynamicShares: false,
			},
			expectedID: uint64(3),
		},
		{
			name: "create a campaign from a non existing coordinator",
			msg: types.MsgCreateCampaign{
				CampaignName:  sample.CampaignName(),
				Coordinator:   sample.AccAddress(),
				TotalSupply:   sample.Coins(),
				DynamicShares: false,
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := campaignSrv.CreateCampaign(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectedID, got.CampaignID)
			campaign, found := campaignKeeper.GetCampaign(sdkCtx, got.CampaignID)
			require.True(t, found)
			require.EqualValues(t, got.CampaignID, campaign.Id)
			require.EqualValues(t, tc.msg.CampaignName, campaign.CampaignName)
			require.EqualValues(t, coordMap[tc.msg.Coordinator], campaign.CoordinatorID)
			require.False(t, campaign.MainnetInitialized)
			require.True(t, tc.msg.TotalSupply.IsEqual(campaign.TotalSupply))
			require.EqualValues(t, tc.msg.DynamicShares, campaign.DynamicShares)
			require.EqualValues(t, types.EmptyShares(), campaign.AllocatedShares)
			require.EqualValues(t, types.EmptyShares(), campaign.TotalShares)
		})
	}
}
