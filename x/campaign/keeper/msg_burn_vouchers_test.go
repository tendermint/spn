package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgBurnVouchers(t *testing.T) {
	var (
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)

		ctx            = sdk.WrapSDKContext(sdkCtx)
		campaign       = sample.Campaign(r, 0)
		addr           = sample.AccAddress(r)
		vouchersTooBig = sdk.NewCoins(
			sdk.NewCoin("v/0/foo", sdkmath.NewInt(spntypes.TotalShareNumber+1)),
		)
	)

	// Create shares
	shares, err := types.NewShares("1000foo,500bar,300foobar")
	require.NoError(t, err)

	// Set campaign
	campaign.AllocatedShares = shares
	campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaign)

	// Create vouchers
	vouchers, err := types.SharesToVouchers(shares, campaign.CampaignID)
	require.NoError(t, err)

	// Send coins to account
	err = tk.BankKeeper.MintCoins(sdkCtx, types.ModuleName, vouchers)
	require.NoError(t, err)
	err = tk.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, addr, vouchers)
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
				Vouchers:   sample.Coins(r),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "invalid sender address",
			msg: types.MsgBurnVouchers{
				Sender:     "invalid_address",
				CampaignID: campaign.CampaignID,
				Vouchers:   sample.Coins(r),
			},
			err: spnerrors.ErrCritical,
		},
		{
			name: "should not burn more than allocated shares",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   vouchersTooBig,
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "burn voucher one",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sdk.NewCoins(vouchers[0]),
			},
		},
		{
			name: "insufficient funds for voucher one",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sdk.NewCoins(vouchers[0]),
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "burn voucher two",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sdk.NewCoins(vouchers[1]),
			},
		},
		{
			name: "insufficient funds for voucher two",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sdk.NewCoins(vouchers[1]),
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "burn voucher three",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sdk.NewCoins(vouchers[2]),
			},
		},
		{
			name: "insufficient funds for voucher three",
			msg: types.MsgBurnVouchers{
				Sender:     addr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sdk.NewCoins(vouchers[2]),
			},
			err: types.ErrInsufficientVouchers,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousCampaign types.Campaign
			var previousBalance sdk.Coins
			var creatorAddr sdk.AccAddress

			// Get values before message execution
			if tc.err == nil {
				var found bool
				previousCampaign, found = tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)

				creatorAddr, err = sdk.AccAddressFromBech32(tc.msg.Sender)
				require.NoError(t, err)
				previousBalance = tk.BankKeeper.GetAllBalances(sdkCtx, creatorAddr)
			}

			// Execute message
			_, err = ts.CampaignSrv.BurnVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			campaign, found := tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)

			// Allocated shares of the campaign must be decreased
			burned, err := types.VouchersToShares(tc.msg.Vouchers, tc.msg.CampaignID)
			require.NoError(t, err)

			expectedShares, err := types.DecreaseShares(previousCampaign.AllocatedShares, burned)
			require.NoError(t, err)
			require.True(t, types.IsEqualShares(expectedShares, campaign.AllocatedShares))

			// Check coordinator balance
			balance := tk.BankKeeper.GetAllBalances(sdkCtx, creatorAddr)
			expectedBalance := previousBalance.Sub(tc.msg.Vouchers...)
			require.True(t, balance.IsEqual(expectedBalance))
		})
	}
}
