package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgAddMainnetAccount(t *testing.T) {
	var (
		addr                                               = sample.AccAddress()
		coordAddr1                                         = sample.AccAddress()
		coordAddr2                                         = sample.AccAddress()
		coordAddr3                                         = sample.AccAddress()
		campaignKeeper, _, campaignSrv, profileSrv, sdkCtx = setupMsgServer(t)
		ctx                                                = sdk.WrapSDKContext(sdkCtx)
		campaign1                                          = sample.Campaign(0)
		campaign2                                          = sample.Campaign(1)
		campaign3                                          = sample.Campaign(2)
	)

	// create shares
	allocatedShares, err := types.NewShares("91token")
	require.NoError(t, err)
	totalShares, err := types.NewShares("100token")
	require.NoError(t, err)
	share1, err := types.NewShares("10token")
	require.NoError(t, err)
	share2, err := types.NewShares("8token")
	require.NoError(t, err)

	// Create a campaign with coordinator
	res1, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	campaign1.CoordinatorID = res1.CoordinatorId
	campaign1.Id = campaignKeeper.AppendCampaign(sdkCtx, campaign1)
	campaign1.AllocatedShares = allocatedShares
	campaign1.TotalShares = totalShares

	res2, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	campaign2.CoordinatorID = res2.CoordinatorId
	campaign2.MainnetInitialized = true
	campaign2.AllocatedShares = allocatedShares
	campaign2.TotalShares = totalShares
	campaign2.Id = campaignKeeper.AppendCampaign(sdkCtx, campaign2)

	res3, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr3,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	campaign3.CoordinatorID = res3.CoordinatorId
	campaign3.AllocatedShares = allocatedShares
	campaign3.TotalShares = totalShares
	campaign3.Id = campaignKeeper.AppendCampaign(sdkCtx, campaign3)

	for _, tc := range []struct {
		name       string
		msg        types.MsgAddMainnetAccount
		expectedID uint64
		err        error
	}{
		{
			name: "invalid campaign id",
			msg: types.MsgAddMainnetAccount{
				Coordinator: coordAddr1,
				CampaignID:  100,
				Address:     addr,
				Shares:      sample.Shares(),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "invalid coordinator address",
			msg: types.MsgAddMainnetAccount{
				Coordinator: addr,
				CampaignID:  campaign1.Id,
				Address:     addr,
				Shares:      sample.Shares(),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator id",
			msg: types.MsgAddMainnetAccount{
				Coordinator: coordAddr2,
				CampaignID:  campaign1.Id,
				Address:     addr,
				Shares:      sample.Shares(),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "campaign already in mainnet",
			msg: types.MsgAddMainnetAccount{
				Coordinator: coordAddr2,
				CampaignID:  campaign2.Id,
				Address:     addr,
				Shares:      sample.Shares(),
			},
			err: types.ErrMainnetInitialized,
		},
		{
			name: "allocated shares greater them total shares",
			msg: types.MsgAddMainnetAccount{
				Coordinator: coordAddr3,
				CampaignID:  campaign3.Id,
				Address:     addr,
				Shares:      share1,
			},
			err: types.ErrTotalShareLimit,
		},
		{
			name: "valid message",
			msg: types.MsgAddMainnetAccount{
				Coordinator: coordAddr1,
				CampaignID:  campaign1.Id,
				Address:     addr,
				Shares:      share2,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var tmpCampaign types.Campaign
			if tc.err == nil {
				var found bool
				tmpCampaign, found = campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)
			}
			_, err := campaignSrv.AddMainnetAccount(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			account, found := campaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Address)
			require.True(t, found)

			campaign, found := campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)

			tmpShare, err := types.DecreaseShares(campaign.AllocatedShares, account.Shares)
			require.NoError(t, err)

			equal := types.IsEqualShares(tmpCampaign.AllocatedShares, tmpShare)
			require.True(t, equal)
		})
	}
}
