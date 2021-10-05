package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgBurnVouchers(t *testing.T) {
	var (
		campaignKeeper, _, _, bankKeeper, campaignSrv, _, sdkCtx = setupMsgServer(t)

		ctx            = sdk.WrapSDKContext(sdkCtx)
		campaign       = sample.Campaign(0)
		addr           = sample.Address()
		vouchersTooBig = sdk.NewCoins(
			sdk.NewCoin("v/0/foo", sdk.NewInt(types.DefaultTotalShareNumber+1)),
		)
	)

	// Create shares
	shares, err := types.NewShares("1000foo,500bar,300foobar")
	require.NoError(t, err)

	// Set campaign
	campaign.AllocatedShares = shares
	campaign.Id = campaignKeeper.AppendCampaign(sdkCtx, campaign)

	// Create vouchers
	vouchers, err := types.SharesToVouchers(shares, campaign.Id)
	require.NoError(t, err)

	// Send coins to account
	err = bankKeeper.MintCoins(sdkCtx, types.ModuleName, vouchers)
	require.NoError(t, err)
	err = bankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, addr, vouchers)
	require.NoError(t, err)

	for _, tc := range []struct {
		name string
		msg  types.MsgBurnVouchers
		err  error
	}{
		{
			name: "non existing campaign",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: 1000,
				Vouchers:   sample.Coins(),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "invalid sender address",
			msg: types.MsgBurnVouchers{
				Sender:     "invalid_address",
				CampaignID: campaign.Id,
				Vouchers:   sample.Coins(),
			},
			err: spnerrors.ErrCritical,
		},
		{
			name: "should not burn more than allocated shares",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.Id,
				Vouchers:   vouchersTooBig,
			},
			err: types.ErrInsufficientVouchersBalance,
		},
		{
			name: "burn voucher one",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.Id,
				Vouchers:   sdk.NewCoins(vouchers[0]),
			},
		},
		{
			name: "insufficient funds for voucher one",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.Id,
				Vouchers:   sdk.NewCoins(vouchers[0]),
			},
			err: types.ErrInsufficientVouchersBalance,
		},
		{
			name: "burn voucher two",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.Id,
				Vouchers:   sdk.NewCoins(vouchers[1]),
			},
		},
		{
			name: "insufficient funds for voucher two",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.Id,
				Vouchers:   sdk.NewCoins(vouchers[1]),
			},
			err: types.ErrInsufficientVouchersBalance,
		},
		{
			name: "burn voucher three",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.Id,
				Vouchers:   sdk.NewCoins(vouchers[2]),
			},
		},
		{
			name: "insufficient funds for voucher three",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.Id,
				Vouchers:   sdk.NewCoins(vouchers[2]),
			},
			err: types.ErrInsufficientVouchersBalance,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousCampaign types.Campaign
			var previousBalance sdk.Coins
			var creatorAddr sdk.AccAddress

			// Get values before message execution
			if tc.err == nil {
				var found bool
				previousCampaign, found = campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)

				creatorAddr, err = sdk.AccAddressFromBech32(tc.msg.Sender)
				require.NoError(t, err)
				previousBalance = bankKeeper.GetAllBalances(sdkCtx, creatorAddr)
			}

			// Execute message
			_, err = campaignSrv.BurnVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			campaign, found := campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)

			// Allocated shares of the campaign must be decreased
			burned, err := types.VouchersToShares(tc.msg.Vouchers, tc.msg.CampaignID)
			require.NoError(t, err)

			expectedShares, err := types.DecreaseShares(previousCampaign.AllocatedShares, burned)
			require.NoError(t, err)
			require.True(t, types.IsEqualShares(expectedShares, campaign.AllocatedShares))

			// Check coordinator balance
			balance := bankKeeper.GetAllBalances(sdkCtx, creatorAddr)
			expectedBalance := previousBalance.Sub(tc.msg.Vouchers)
			require.True(t, balance.IsEqual(expectedBalance))
		})
	}
}
