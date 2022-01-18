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
		coordAddr1                                               = sample.Address()
		coordAddr2                                               = sample.Address()
		campaignKeeper, _, _, _, campaignSrv, profileSrv, sdkCtx = setupMsgServer(t)
		ctx                                                      = sdk.WrapSDKContext(sdkCtx)
	)

	// Create coordinators
	coordMap := make(map[string]uint64)
	res, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	coordMap[coordAddr1] = res.CoordinatorID
	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	coordMap[coordAddr2] = res.CoordinatorID

	for _, tc := range []struct {
		name       string
		msg        types.MsgCreateCampaign
		expectedID uint64
		err        error
	}{
		{
			name: "create a campaign 1",
			msg: types.MsgCreateCampaign{
				CampaignName: sample.CampaignName(),
				Coordinator:  coordAddr1,
				TotalSupply:  sample.Coins(),
			},
			expectedID: uint64(0),
		},
		{
			name: "create a campaign from a different coordinator",
			msg: types.MsgCreateCampaign{
				CampaignName: sample.CampaignName(),
				Coordinator:  coordAddr2,
				TotalSupply:  sample.Coins(),
			},
			expectedID: uint64(1),
		},
		{
			name: "create a campaign from a non existing coordinator",
			msg: types.MsgCreateCampaign{
				CampaignName: sample.CampaignName(),
				Coordinator:  sample.Address(),
				TotalSupply:  sample.Coins(),
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
			require.EqualValues(t, got.CampaignID, campaign.CampaignID)
			require.EqualValues(t, tc.msg.CampaignName, campaign.CampaignName)
			require.EqualValues(t, coordMap[tc.msg.Coordinator], campaign.CoordinatorID)
			require.False(t, campaign.MainnetInitialized)
			require.True(t, tc.msg.TotalSupply.IsEqual(campaign.TotalSupply))
			require.EqualValues(t, types.EmptyShares(), campaign.AllocatedShares)
			require.EqualValues(t, types.EmptyShares(), campaign.TotalShares)

			// dynamic share is always disabled
			require.EqualValues(t, false, campaign.DynamicShares)

			// Empty list of campaign chains
			campaignChains, found := campaignKeeper.GetCampaignChains(sdkCtx, got.CampaignID)
			require.True(t, found)
			require.EqualValues(t, got.CampaignID, campaignChains.CampaignID)
			require.Empty(t, campaignChains.Chains)
		})
	}
}
