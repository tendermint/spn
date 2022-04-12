package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
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
		coordAddrMainnetLaunched       = sample.Address(r)
		campaign                       = sample.Campaign(r, 0)
		campaignInvalidAllocatedShares = sample.Campaign(r, 2)
		campaignMainnetInitialized     = sample.Campaign(r, 1)
		campaignMainnetLaunched        = sample.Campaign(r, 3)

		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
	)

	// create shares
	allocatedShares := tc.Shares(t, "91token")
	invalidAllocatedShares := tc.Shares(t, fmt.Sprintf("%dtoken", spntypes.TotalShareNumber))
	highShare := tc.Shares(t, "1000token")
	lowShare := tc.Shares(t, "8token")

	// Create a campaigns
	res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddrMainnetInitialized,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)
	campaignMainnetInitialized.CoordinatorID = res.CoordinatorID
	campaignMainnetInitialized.MainnetInitialized = true
	campaignMainnetInitialized.AllocatedShares = allocatedShares
	chain := sample.Chain(r, 0, res.CoordinatorID)
	chain.IsMainnet = true
	chain.LaunchTriggered = false
	campaignMainnetInitialized.MainnetID = tk.LaunchKeeper.AppendChain(sdkCtx, chain)
	campaignMainnetInitialized.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaignMainnetInitialized)

	res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddrMainnetLaunched,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)
	campaignMainnetLaunched.CoordinatorID = res.CoordinatorID
	campaignMainnetLaunched.MainnetInitialized = true
	campaignMainnetLaunched.AllocatedShares = allocatedShares
	chainLaunched := sample.Chain(r, 1, res.CoordinatorID)
	chainLaunched.IsMainnet = true
	chainLaunched.LaunchTriggered = true
	campaignMainnetLaunched.MainnetID = tk.LaunchKeeper.AppendChain(sdkCtx, chainLaunched)
	campaignMainnetLaunched.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaignMainnetLaunched)

	res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)
	campaign.CoordinatorID = res.CoordinatorID
	campaign.AllocatedShares = allocatedShares
	campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaign)
	accShare := sample.MainnetVestingAccountWithShares(r, campaign.CampaignID, addr2, lowShare)
	tk.CampaignKeeper.SetMainnetVestingAccount(sdkCtx, accShare)

	res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(t, err)
	campaignInvalidAllocatedShares.CoordinatorID = res.CoordinatorID
	campaignInvalidAllocatedShares.AllocatedShares = invalidAllocatedShares
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
			name: "campaign with launched mainnet",
			msg: types.MsgAddVestingOptions{
				Coordinator:    coordAddrMainnetLaunched,
				CampaignID:     campaignMainnetLaunched.CampaignID,
				Address:        addr1,
				VestingOptions: sample.ShareVestingOptions(r),
			},
			err: types.ErrMainnetLaunchTriggered,
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
