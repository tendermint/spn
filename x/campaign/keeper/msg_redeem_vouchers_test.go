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

func TestMsgRedeemVouchers(t *testing.T) {
	var (
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)

		ctx            = sdk.WrapSDKContext(sdkCtx)
		addr           = sample.AccAddress()
		existAddr      = sample.AccAddress()
		campaign       = sample.Campaign(0)
		vouchersTooBig = sdk.NewCoins(
			sdk.NewCoin("v/0/foo", sdk.NewInt(spntypes.TotalShareNumber+1)),
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

	tk.CampaignKeeper.SetMainnetAccount(sdkCtx, types.MainnetAccount{
		CampaignID: campaign.CampaignID,
		Address:    existAddr.String(),
		Shares:     shares,
	})
	err = tk.BankKeeper.MintCoins(sdkCtx, types.ModuleName, vouchers)
	require.NoError(t, err)
	err = tk.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, existAddr, vouchers)
	require.NoError(t, err)

	for _, tc := range []struct {
		name string
		msg  types.MsgRedeemVouchers
		err  error
	}{
		{
			name: "non existing campaign",
			msg: types.MsgRedeemVouchers{
				Sender:     addr.String(),
				Account:    addr.String(),
				CampaignID: 10000,
				Vouchers:   sample.Coins(),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "invalid vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:     addr.String(),
				Account:    addr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sample.Coins(),
			},
			err: spnerrors.ErrCritical,
		},
		{
			name: "invalid sender address",
			msg: types.MsgRedeemVouchers{
				Sender:     "invalid_address",
				Account:    addr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   vouchers,
			},
			err: spnerrors.ErrCritical,
		},
		{
			name: "insufficient funds",
			msg: types.MsgRedeemVouchers{
				Sender:     addr.String(),
				Account:    addr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   vouchersTooBig,
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "new account redeem all",
			msg: types.MsgRedeemVouchers{
				Sender:     addr.String(),
				Account:    sample.Address(),
				CampaignID: campaign.CampaignID,
				Vouchers:   vouchers,
			},
		},
		{
			name: "exist account redeem voucher one",
			msg: types.MsgRedeemVouchers{
				Sender:     existAddr.String(),
				Account:    existAddr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sdk.NewCoins(vouchers[0]),
			},
		},
		{
			name: "exist account redeem voucher two",
			msg: types.MsgRedeemVouchers{
				Sender:     existAddr.String(),
				Account:    existAddr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sdk.NewCoins(vouchers[1]),
			},
		},
		{
			name: "exist account redeem voucher three",
			msg: types.MsgRedeemVouchers{
				Sender:     existAddr.String(),
				Account:    existAddr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sdk.NewCoins(vouchers[2]),
			},
		},
		{
			name: "account without funds for vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:     existAddr.String(),
				Account:    existAddr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   vouchers,
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "account without funds for voucher one",
			msg: types.MsgRedeemVouchers{
				Sender:     existAddr.String(),
				Account:    existAddr.String(),
				CampaignID: campaign.CampaignID,
				Vouchers:   sdk.NewCoins(vouchers[0]),
			},
			err: types.ErrInsufficientVouchers,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousAccount types.MainnetAccount
			var previousBalance sdk.Coins
			var foundAccount bool
			var accountAddr sdk.AccAddress

			// Get values before message execution
			if tc.err == nil {
				accountAddr, err = sdk.AccAddressFromBech32(tc.msg.Account)
				require.NoError(t, err)

				previousAccount, foundAccount = tk.CampaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Account)
				if foundAccount {
					previousBalance = tk.BankKeeper.GetAllBalances(sdkCtx, accountAddr)
				}
			}

			// Execute message
			_, err = ts.CampaignSrv.RedeemVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			shares, err := types.VouchersToShares(tc.msg.Vouchers, tc.msg.CampaignID)
			require.NoError(t, err)

			account, found := tk.CampaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Account)
			require.True(t, found)

			// Check account shares
			expectedShares := shares
			if foundAccount {
				expectedShares = types.IncreaseShares(previousAccount.Shares, shares)
			}
			require.True(t, types.IsEqualShares(expectedShares, account.Shares))

			// Check account balance
			expectedVouchers := sdk.Coins{}
			if foundAccount {
				var negative bool
				expectedVouchers, negative = previousBalance.SafeSub(tc.msg.Vouchers)
				require.False(t, negative)
			}
			balance := tk.BankKeeper.GetAllBalances(sdkCtx, accountAddr)
			require.True(t, expectedVouchers.IsEqual(balance))
		})
	}
}
