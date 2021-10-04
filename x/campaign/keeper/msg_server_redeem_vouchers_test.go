package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgRedeemVouchers(t *testing.T) {
	var (
		campaignKeeper, _, _, bankKeeper, campaignSrv, _, sdkCtx = setupMsgServer(t)

		addr     = sample.AccAddress()
		campaign = sample.Campaign(0)
		ctx      = sdk.WrapSDKContext(sdkCtx)
		//vouchersTooBig = sdk.NewCoins(
		//	sdk.NewCoin("v/0/foo", sdk.NewInt(types.DefaultTotalShareNumber+1)),
		//)
	)

	// Create shares
	shares, err := types.NewShares("1000foo,500bar,300foobar")
	require.NoError(t, err)

	// Create vouchers
	vouchers, err := types.SharesToVouchers(shares, campaign.Id)
	require.NoError(t, err)

	for _, tc := range []struct {
		name string
		msg  types.MsgRedeemVouchers
		err  error
	}{
		{
			name: "redeem vouchers",
			msg: types.MsgRedeemVouchers{
				Creator:    addr,
				CampaignID: 0,
				Vouchers:   vouchers,
			},
		},
		{
			name: "redeem vouchers a second time",
			msg: types.MsgRedeemVouchers{
				Creator:    addr,
				CampaignID: 0,
				Vouchers:   vouchers,
			},
		},
		{
			name: "redeem vouchers to zero",
			msg: types.MsgRedeemVouchers{
				Creator:    addr,
				CampaignID: 0,
				Vouchers:   vouchers,
			},
		},
		{
			name: "redeem vouchers from another account",
			msg: types.MsgRedeemVouchers{
				Creator:    addr,
				CampaignID: 0,
				Vouchers:   vouchers,
			},
		},
		{
			name: "not enough shares in balance",
			msg: types.MsgRedeemVouchers{
				Creator:    addr,
				CampaignID: 0,
				Vouchers:   vouchers,
			},
			err: types.ErrSharesDecrease,
		},
		{
			name: "non existent campaign",
			msg: types.MsgRedeemVouchers{
				Creator:    addr,
				CampaignID: 1000,
				Vouchers:   vouchers,
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "non existent account",
			msg: types.MsgRedeemVouchers{
				Creator:    sample.AccAddress(),
				CampaignID: 0,
				Vouchers:   vouchers,
			},
			err: types.ErrAccountNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousAccount types.MainnetAccount
			var previousBalance sdk.Coins
			var found bool

			accountAddr, err := sdk.AccAddressFromBech32(tc.msg.Creator)
			require.NoError(t, err)

			// Get values before message execution
			if tc.err == nil {
				previousAccount, found = campaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Creator)
				require.True(t, found)

				previousBalance = bankKeeper.GetAllBalances(sdkCtx, accountAddr)
			}

			// Execute message
			_, err = campaignSrv.RedeemVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			if types.IsEqualShares(tc.msg.Vouchers, previousAccount.Shares) {
				// All redeemed
				_, found := campaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Creator)
				require.False(t, found)

			} else {
				account, found := campaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Creator)
				require.True(t, found)

				expectedvouchers, err := types.DecreaseShares(previousAccount.Shares, tc.msg.Vouchers)
				require.NoError(t, err)
				require.True(t, types.IsEqualShares(expectedvouchers, account.Shares))
			}

			// Compare balance
			redeemed, err := types.SharesToVouchers(tc.msg.Vouchers, tc.msg.CampaignID)
			require.NoError(t, err)
			balance := bankKeeper.GetAllBalances(sdkCtx, accountAddr)
			require.True(t, balance.IsEqual(previousBalance.Add(redeemed...)))
		})
	}
}
