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

func TestMsgAddVestingOptions(t *testing.T) {
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
	allocatedShares, err := types.NewShares("999token")
	require.NoError(t, err)
	totalShares, err := types.NewShares("9999token")
	require.NoError(t, err)
	highShare, err := types.NewShares("9999token")
	require.NoError(t, err)
	lowShare, err := types.NewShares("8token")
	require.NoError(t, err)

	// Create a campaigns
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
	campaign.AllocatedShares = allocatedShares
	campaign.TotalShares = totalShares
	campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaign)
	accShare := sample.MainnetVestingAccountWithShares(r, campaign.CampaignID, addr2, lowShare)
	tk.CampaignKeeper.SetMainnetVestingAccount(sdkCtx, accShare)

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
		msg        types.MsgAddVestingOptions
		expectedID uint64
		err        error
	}{
		{
			name: "invalid campaign id",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddr1,
				CampaignID:     100,
				Address:        addr1,
				VestingOptions: sample.ShareVestingOptions(r),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "invalid coordinator address",
			msg: types.MsgAddVestingOptions{
				Coordinator:    addr1,
				CampaignID:     campaign.CampaignID,
				Address:        addr1,
				VestingOptions: sample.ShareVestingOptions(r),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator id",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddrMainnetInitialized,
				CampaignID:     campaign.CampaignID,
				Address:        addr1,
				VestingOptions: sample.ShareVestingOptions(r),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "campaign with initialized mainnet",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddrMainnetInitialized,
				CampaignID:     campaignMainnetInitialized.CampaignID,
				Address:        addr1,
				VestingOptions: sample.ShareVestingOptions(r),
			},
		},
		{
			name: "allocated shares greater them total shares",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddr2,
				CampaignID:     campaignInvalidAllocatedShares.CampaignID,
				Address:        addr1,
				VestingOptions: sample.CustomShareVestingOptions(r, highShare),
			},
			err: types.ErrTotalSharesLimit,
		},
		{
			name: "create new account with shares",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddr1,
				CampaignID:     campaign.CampaignID,
				Address:        addr1,
				VestingOptions: sample.CustomShareVestingOptions(r, lowShare),
			},
		},
		{
			name: "update existing account shares",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddr1,
				CampaignID:     campaign.CampaignID,
				Address:        addr2,
				VestingOptions: sample.CustomShareVestingOptions(r, lowShare),
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var (
				accountExists    bool
				previousAccount  types.MainnetVestingAccount
				previousCampaign types.Campaign
			)
			if tc.err == nil {
				var found bool
				previousCampaign, found = tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)

				previousAccount, accountExists = tk.CampaignKeeper.GetMainnetVestingAccount(
					sdkCtx,
					tc.msg.CampaignID,
					tc.msg.Address,
				)
			}
			_, err := ts.CampaignSrv.AddVestingOptions(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			account, found := tk.CampaignKeeper.GetMainnetVestingAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Address)
			require.True(t, found)

			campaign, found := tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)

			totalShares, err := account.GetTotalShares()
			require.NoError(t, err)
			tmpShare := types.IncreaseShares(previousCampaign.AllocatedShares, totalShares)

			if accountExists {
				tmpAccShares, err := previousAccount.GetTotalShares()
				require.NoError(t, err)
				tmpShare, err = types.DecreaseShares(tmpShare, tmpAccShares)
				require.NoError(t, err)
			}

			require.Equal(t, tc.msg.VestingOptions, account.VestingOptions)
			equal := types.IsEqualShares(campaign.AllocatedShares, tmpShare)
			require.True(t, equal)
		})
	}
}
