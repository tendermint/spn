package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgAddShares(t *testing.T) {
	var (
		addr1                       = sample.AccAddress()
		addr2                       = sample.AccAddress()
		coordAddr1                  = sample.AccAddress()
		coordAddr2                  = sample.AccAddress()
		coordAddrMainnetInitialized = sample.AccAddress()
		campaign1                   = sample.Campaign(0)
		campaign2                   = sample.Campaign(2)
		campaignMainnetInitialized  = sample.Campaign(1)

		campaignKeeper, _, campaignSrv, profileSrv, sdkCtx = setupMsgServer(t)
		ctx                                                = sdk.WrapSDKContext(sdkCtx)
	)

	// create shares
	allocatedShares, err := types.NewShares("91token")
	require.NoError(t, err)
	totalShares, err := types.NewShares("100token")
	require.NoError(t, err)
	highShare, err := types.NewShares("1000token")
	require.NoError(t, err)
	lowShare, err := types.NewShares("8token")
	require.NoError(t, err)

	campaignKeeper.SetMainnetAccount(sdkCtx, sample.MainnetAccount(campaign1.Id, addr2))
	res, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddrMainnetInitialized,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	campaignMainnetInitialized.CoordinatorID = res.CoordinatorId
	campaignMainnetInitialized.MainnetInitialized = true
	campaignMainnetInitialized.AllocatedShares = allocatedShares
	campaignMainnetInitialized.TotalShares = totalShares
	campaignMainnetInitialized.Id = campaignKeeper.AppendCampaign(sdkCtx, campaignMainnetInitialized)

	// Create a campaign with coordinator
	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	campaign1.CoordinatorID = res.CoordinatorId
	campaign1.Id = campaignKeeper.AppendCampaign(sdkCtx, campaign1)
	campaign1.AllocatedShares = allocatedShares
	campaign1.TotalShares = totalShares

	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	campaign2.CoordinatorID = res.CoordinatorId
	campaign2.AllocatedShares = allocatedShares
	campaign2.TotalShares = totalShares
	campaign2.Id = campaignKeeper.AppendCampaign(sdkCtx, campaign2)

	for _, tc := range []struct {
		name       string
		msg        types.MsgAddShares
		expectedID uint64
		err        error
	}{
		{
			name: "invalid campaign id",
			msg: types.MsgAddShares{
				Coordinator: coordAddr1,
				CampaignID:  100,
				Address:     addr1,
				Shares:      sample.Shares(),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "coordinator address not found",
			msg: types.MsgAddShares{
				Coordinator: addr1,
				CampaignID:  campaign1.Id,
				Address:     addr1,
				Shares:      sample.Shares(),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator id",
			msg: types.MsgAddShares{
				Coordinator: coordAddrMainnetInitialized,
				CampaignID:  campaign1.Id,
				Address:     addr1,
				Shares:      sample.Shares(),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "campaign already in mainnet",
			msg: types.MsgAddShares{
				Coordinator: coordAddrMainnetInitialized,
				CampaignID:  campaignMainnetInitialized.Id,
				Address:     addr1,
				Shares:      sample.Shares(),
			},
			err: types.ErrMainnetInitialized,
		},
		{
			name: "allocated shares greater than total shares",
			msg: types.MsgAddShares{
				Coordinator: coordAddr2,
				CampaignID:  campaign2.Id,
				Address:     addr1,
				Shares:      highShare,
			},
			err: types.ErrTotalSharesLimit,
		},
		{
			name: "create new account with shares",
			msg: types.MsgAddShares{
				Coordinator: coordAddr1,
				CampaignID:  campaign1.Id,
				Address:     addr1,
				Shares:      lowShare,
			},
		},
		{
			name: "update existing account shares",
			msg: types.MsgAddShares{
				Coordinator: coordAddr1,
				CampaignID:  campaign1.Id,
				Address:     addr2,
				Shares:      highShare,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var (
				accountExists    bool
				previousAccount  types.MainnetAccount
				previousCampaign types.Campaign
			)
			if tc.err == nil {
				var found bool
				previousCampaign, found = campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)

				previousAccount, accountExists = campaignKeeper.GetMainnetAccount(
					sdkCtx,
					tc.msg.CampaignID,
					tc.msg.Address,
				)
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

			if accountExists {
				shares, err := types.DecreaseShares(account.Shares, previousAccount.Shares)
				require.NoError(t, err)
				equal := types.IsEqualShares(shares, tc.msg.Shares)
				require.True(t, equal)
			} else {
				tmpShare, err := types.DecreaseShares(campaign.AllocatedShares, account.Shares)
				require.NoError(t, err)
				equal := types.IsEqualShares(previousCampaign.AllocatedShares, tmpShare)
				require.True(t, equal)
			}
		})
	}
}
