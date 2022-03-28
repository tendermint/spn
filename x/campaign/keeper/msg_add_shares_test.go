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

func TestMsgAddShares(t *testing.T) {
	var (
		addr1                          = sample.Address(r)
		addr2                          = sample.Address(r)
		coordAddr1                     = sample.Address(r)
		coordAddr2                     = sample.Address(r)
		coordAddrMainnetInitialized    = sample.Address(r)
		campaign                       = sample.Campaign(r, 0)
		campaignInvalidAllocatedShares = sample.Campaign(r, 2)
		campaignMainnetInitialized     = sample.Campaign(r, 1)

		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
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

	tk.CampaignKeeper.SetMainnetAccount(sdkCtx, sample.MainnetAccount(r, campaign.CampaignID, addr2))
	res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddrMainnetInitialized,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)
	campaignMainnetInitialized.CoordinatorID = res.CoordinatorID
	campaignMainnetInitialized.MainnetInitialized = true
	campaignMainnetInitialized.AllocatedShares = allocatedShares
	campaignMainnetInitialized.TotalShares = totalShares
	campaignMainnetInitialized.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaignMainnetInitialized)

	res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)
	campaign.CoordinatorID = res.CoordinatorID
	campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaign)
	campaign.AllocatedShares = allocatedShares
	campaign.TotalShares = totalShares

	res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)
	campaignInvalidAllocatedShares.CoordinatorID = res.CoordinatorID
	campaignInvalidAllocatedShares.AllocatedShares = allocatedShares
	campaignInvalidAllocatedShares.TotalShares = totalShares
	campaignInvalidAllocatedShares.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaignInvalidAllocatedShares)

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
				Shares:      sample.Shares(r),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "coordinator address not found",
			msg: types.MsgAddShares{
				Coordinator: addr1,
				CampaignID:  campaign.CampaignID,
				Address:     addr1,
				Shares:      sample.Shares(r),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator id",
			msg: types.MsgAddShares{
				Coordinator: coordAddrMainnetInitialized,
				CampaignID:  campaign.CampaignID,
				Address:     addr1,
				Shares:      sample.Shares(r),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "campaign with initialized mainnet",
			msg: types.MsgAddShares{
				Coordinator: coordAddrMainnetInitialized,
				CampaignID:  campaignMainnetInitialized.CampaignID,
				Address:     addr1,
				Shares:      sample.Shares(r),
			},
		},
		{
			name: "allocated shares greater than total shares",
			msg: types.MsgAddShares{
				Coordinator: coordAddr2,
				CampaignID:  campaignInvalidAllocatedShares.CampaignID,
				Address:     addr1,
				Shares:      highShare,
			},
			err: types.ErrTotalSharesLimit,
		},
		{
			name: "create new account with shares",
			msg: types.MsgAddShares{
				Coordinator: coordAddr1,
				CampaignID:  campaign.CampaignID,
				Address:     addr1,
				Shares:      lowShare,
			},
		},
		{
			name: "update existing account shares",
			msg: types.MsgAddShares{
				Coordinator: coordAddr1,
				CampaignID:  campaign.CampaignID,
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
				previousCampaign, found = tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)

				previousAccount, accountExists = tk.CampaignKeeper.GetMainnetAccount(
					sdkCtx,
					tc.msg.CampaignID,
					tc.msg.Address,
				)
			}
			_, err := ts.CampaignSrv.AddShares(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			account, found := tk.CampaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Address)
			require.True(t, found)

			campaign, found := tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
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
