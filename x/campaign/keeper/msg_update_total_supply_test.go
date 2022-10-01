package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateTotalSupply(t *testing.T) {
	var (
		coordID        uint64
		coordAddr1     = sample.Address(r)
		coordAddr2     = sample.Address(r)
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
	)

	t.Run("should allow creating coordinators", func(t *testing.T) {
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddr1,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
		coordID = res.CoordinatorID
		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddr2,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
	})

	// Set a regular campaign and a campaign with an already initialized mainnet
	campaign := sample.Campaign(r, 0)
	campaign.CoordinatorID = coordID
	tk.CampaignKeeper.SetCampaign(sdkCtx, campaign)

	campaign = sample.Campaign(r, 1)
	campaign.CoordinatorID = coordID
	campaign.MainnetInitialized = true
	tk.CampaignKeeper.SetCampaign(sdkCtx, campaign)

	for _, tc := range []struct {
		name string
		msg  types.MsgUpdateTotalSupply
		err  error
	}{
		{
			name: "should update total supply",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
		},
		{
			name: "should allow update total supply again",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
		},
		{
			name: "should fail if campaign not found",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        100,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "should fail with non existing coordinator",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        0,
				Coordinator:       sample.Address(r),
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should fail if coordinator is not associated with campaign",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        0,
				Coordinator:       coordAddr2,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "cannot update total supply when mainnet is initialized",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        1,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: types.ErrMainnetInitialized,
		},
		{
			name: "should fail if total supply outside of valid range",
			msg: types.MsgUpdateTotalSupply{
				CampaignID:        0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.CoinsWithRange(r, 10, 20),
			},
			err: types.ErrInvalidTotalSupply,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousTotalSupply sdk.Coins
			if tc.err == nil {
				campaign, found := tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)
				previousTotalSupply = campaign.TotalSupply
			}

			_, err := ts.CampaignSrv.UpdateTotalSupply(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			campaign, found := tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)
			require.True(t, campaign.TotalSupply.IsEqual(
				types.UpdateTotalSupply(previousTotalSupply, tc.msg.TotalSupplyUpdate),
			))
		})
	}
}
