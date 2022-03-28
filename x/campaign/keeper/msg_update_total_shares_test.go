package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateTotalShares(t *testing.T) {
	var (
		coordAddr1     = sample.Address(r)
		coordAddr2     = sample.Address(r)
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
	)

	// Create coordinators
	res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)
	coordID := res.CoordinatorID
	res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)

	// Set different campaigns
	campaign := sample.Campaign(r, 0)
	campaign.CoordinatorID = coordID
	campaign.DynamicShares = true
	campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaign)

	campaignMainnetInitialized := sample.Campaign(r, 1)
	campaignMainnetInitialized.CoordinatorID = coordID
	campaignMainnetInitialized.DynamicShares = true
	campaignMainnetInitialized.MainnetInitialized = true
	campaignMainnetInitialized.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaignMainnetInitialized)

	campaignNoDynamicShares := sample.Campaign(r, 2)
	campaignNoDynamicShares.CoordinatorID = coordID
	campaignNoDynamicShares.DynamicShares = false
	campaignNoDynamicShares.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaignNoDynamicShares)

	campaignWithAllocatedShares := sample.Campaign(r, 3)
	campaignWithAllocatedShares.CoordinatorID = coordID
	campaignWithAllocatedShares.DynamicShares = true
	campaignWithAllocatedShares.AllocatedShares, _ = types.NewShares("100foo")
	campaignWithAllocatedShares.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaignWithAllocatedShares)
	smallerTotalShares, _ := types.NewShares("50foo")

	for _, tc := range []struct {
		name string
		msg  types.MsgUpdateTotalShares
		err  error
	}{
		{
			name: "update total shares",
			msg: types.MsgUpdateTotalShares{
				CampaignID:  campaign.CampaignID,
				Coordinator: coordAddr1,
				TotalShares: sample.Shares(r),
			},
		},
		{
			name: "can update total shares again",
			msg: types.MsgUpdateTotalShares{
				CampaignID:  campaign.CampaignID,
				Coordinator: coordAddr1,
				TotalShares: sample.Shares(r),
			},
		},
		{
			name: "campaign not found",
			msg: types.MsgUpdateTotalShares{
				CampaignID:  100,
				Coordinator: coordAddr1,
				TotalShares: sample.Shares(r),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "non existing coordinator",
			msg: types.MsgUpdateTotalShares{
				CampaignID:  campaign.CampaignID,
				Coordinator: sample.Address(r),
				TotalShares: sample.Shares(r),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "not the coordinator of the campaign",
			msg: types.MsgUpdateTotalShares{
				CampaignID:  campaign.CampaignID,
				Coordinator: coordAddr2,
				TotalShares: sample.Shares(r),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "cannot update total shares when mainnet is initialized",
			msg: types.MsgUpdateTotalShares{
				CampaignID:  campaignMainnetInitialized.CampaignID,
				Coordinator: coordAddr1,
				TotalShares: sample.Shares(r),
			},
			err: types.ErrMainnetInitialized,
		},
		{
			name: "cannot update total shares when dynamic shares option not set",
			msg: types.MsgUpdateTotalShares{
				CampaignID:  campaignNoDynamicShares.CampaignID,
				Coordinator: coordAddr1,
				TotalShares: sample.Shares(r),
			},
			err: types.ErrNoDynamicShares,
		},
		{
			name: "cannot update total shares when below allocated shares",
			msg: types.MsgUpdateTotalShares{
				CampaignID:  campaignWithAllocatedShares.CampaignID,
				Coordinator: coordAddr1,
				TotalShares: smallerTotalShares,
			},
			err: types.ErrInvalidShares,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ts.CampaignSrv.UpdateTotalShares(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			campaign, found := tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)
			require.True(t, sdk.Coins(tc.msg.TotalShares).IsEqual(sdk.Coins(campaign.TotalShares)))
		})
	}
}
