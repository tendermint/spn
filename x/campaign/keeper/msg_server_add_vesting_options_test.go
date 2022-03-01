package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgAddVestingOptions(t *testing.T) {
	var (
		addr1                          = sample.Address()
		addr2                          = sample.Address()
		coordAddr1                     = sample.Address()
		coordAddr2                     = sample.Address()
		coordAddrMainnetInitialized    = sample.Address()
		campaign                       = sample.Campaign(0)
		campaignInvalidAllocatedShares = sample.Campaign(2)
		campaignMainnetInitialized     = sample.Campaign(1)

		sdkCtx, tk, campaignSrv, profileSrv = setupMsgServer(t)
		ctx                                 = sdk.WrapSDKContext(sdkCtx)
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
	res, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddrMainnetInitialized,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	campaignMainnetInitialized.CoordinatorID = res.CoordinatorID
	campaignMainnetInitialized.MainnetInitialized = true
	campaignMainnetInitialized.AllocatedShares = allocatedShares
	campaignMainnetInitialized.TotalShares = totalShares
	campaignMainnetInitialized.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaignMainnetInitialized)

	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	campaign.CoordinatorID = res.CoordinatorID
	campaign.AllocatedShares = allocatedShares
	campaign.TotalShares = totalShares
	campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaign)
	accShare := sample.MainnetVestingAccountWithShares(campaign.CampaignID, addr2, lowShare)
	tk.CampaignKeeper.SetMainnetVestingAccount(sdkCtx, accShare)

	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(),
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
				VestingOptions: sample.ShareVestingOptions(),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "invalid coordinator address",
			msg: types.MsgAddVestingOptions{
				Coordinator:    addr1,
				CampaignID:     campaign.CampaignID,
				Address:        addr1,
				VestingOptions: sample.ShareVestingOptions(),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator id",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddrMainnetInitialized,
				CampaignID:     campaign.CampaignID,
				Address:        addr1,
				VestingOptions: sample.ShareVestingOptions(),
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "campaign with initialized mainnet",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddrMainnetInitialized,
				CampaignID:     campaignMainnetInitialized.CampaignID,
				Address:        addr1,
				VestingOptions: sample.ShareVestingOptions(),
			},
		},
		{
			name: "allocated shares greater them total shares",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddr2,
				CampaignID:     campaignInvalidAllocatedShares.CampaignID,
				Address:        addr1,
				VestingOptions: sample.CustomShareVestingOptions(highShare),
			},
			err: types.ErrTotalSharesLimit,
		},
		{
			name: "create new account with shares",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddr1,
				CampaignID:     campaign.CampaignID,
				Address:        addr1,
				VestingOptions: sample.CustomShareVestingOptions(lowShare),
			},
		},
		{
			name: "update existing account shares",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddr1,
				CampaignID:     campaign.CampaignID,
				Address:        addr2,
				VestingOptions: sample.CustomShareVestingOptions(lowShare),
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
			_, err := campaignSrv.AddVestingOptions(ctx, &tc.msg)
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
