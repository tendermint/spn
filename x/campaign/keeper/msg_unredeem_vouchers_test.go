package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgUnredeemVouchers(t *testing.T) {
	var (
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)

		accountAddr              = sample.Address()
		account                  = sample.MainnetAccount(0, accountAddr)
		accountShare, _          = types.NewShares("30foo,30bar")
		accountFewSharesAddr     = sample.Address()
		accountFewShares         = sample.MainnetAccount(0, accountFewSharesAddr)
		accountFewSharesShare, _ = types.NewShares("30foo,15bar")

		shares, _ = types.NewShares("10foo,10bar")
	)
	account.Shares = accountShare
	accountFewShares.Shares = accountFewSharesShare

	// Create campaign
	tk.CampaignKeeper.AppendCampaign(sdkCtx, sample.Campaign(0))

	// Create accounts
	tk.CampaignKeeper.SetMainnetAccount(sdkCtx, account)
	tk.CampaignKeeper.SetMainnetAccount(sdkCtx, accountFewShares)

	for _, tc := range []struct {
		name string
		msg  types.MsgUnredeemVouchers
		err  error
	}{
		{
			name: "unredeem vouchers",
			msg: types.MsgUnredeemVouchers{
				Sender:     accountAddr,
				CampaignID: 0,
				Shares:     shares,
			},
		},
		{
			name: "unredeem vouchers a second time",
			msg: types.MsgUnredeemVouchers{
				Sender:     accountAddr,
				CampaignID: 0,
				Shares:     shares,
			},
		},
		{
			name: "unredeem vouchers to zero",
			msg: types.MsgUnredeemVouchers{
				Sender:     accountAddr,
				CampaignID: 0,
				Shares:     shares,
			},
		},
		{
			name: "unredeem vouchers from another account",
			msg: types.MsgUnredeemVouchers{
				Sender:     accountFewSharesAddr,
				CampaignID: 0,
				Shares:     shares,
			},
		},
		{
			name: "not enough shares in balance",
			msg: types.MsgUnredeemVouchers{
				Sender:     accountFewSharesAddr,
				CampaignID: 0,
				Shares:     shares,
			},
			err: types.ErrSharesDecrease,
		},
		{
			name: "non existent campaign",
			msg: types.MsgUnredeemVouchers{
				Sender:     accountAddr,
				CampaignID: 1000,
				Shares:     shares,
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "non existent account",
			msg: types.MsgUnredeemVouchers{
				Sender:     sample.Address(),
				CampaignID: 0,
				Shares:     shares,
			},
			err: types.ErrAccountNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousAccount types.MainnetAccount
			var previousBalance sdk.Coins
			var found bool

			accountAddr, err := sdk.AccAddressFromBech32(tc.msg.Sender)
			require.NoError(t, err)

			// Get values before message execution
			if tc.err == nil {
				previousAccount, found = tk.CampaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Sender)
				require.True(t, found)

				previousBalance = tk.BankKeeper.GetAllBalances(sdkCtx, accountAddr)
			}

			// Execute message
			_, err = ts.CampaignSrv.UnredeemVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			if types.IsEqualShares(tc.msg.Shares, previousAccount.Shares) {
				// All unredeemed
				_, found := tk.CampaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Sender)
				require.False(t, found)

			} else {
				account, found := tk.CampaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Sender)
				require.True(t, found)

				expectedShares, err := types.DecreaseShares(previousAccount.Shares, tc.msg.Shares)
				require.NoError(t, err)
				require.True(t, types.IsEqualShares(expectedShares, account.Shares))
			}

			// Compare balance
			unredeemed, err := types.SharesToVouchers(tc.msg.Shares, tc.msg.CampaignID)
			require.NoError(t, err)
			balance := tk.BankKeeper.GetAllBalances(sdkCtx, accountAddr)
			require.True(t, balance.IsEqual(previousBalance.Add(unredeemed...)))
		})
	}
}
