package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgRedeemVouchers(t *testing.T) {
	var (
		campaignKeeper, _, _, bankKeeper, campaignSrv, _, sdkCtx = setupMsgServer(t)

		ctx            = sdk.WrapSDKContext(sdkCtx)
		addr           = sample.Address()
		existAddr      = sample.Address()
		campaign       = sample.Campaign(0)
		shareFoo       = sdk.NewCoin("foo", sdk.NewInt(1000))
		shareBar       = sdk.NewCoin("bar", sdk.NewInt(500))
		shareFooBar    = sdk.NewCoin("foobar", sdk.NewInt(300))
		shares         = types.NewSharesFromCoins(sdk.NewCoins(shareFoo, shareBar, shareFooBar))
		vouchersTooBig = sdk.NewCoins(
			sdk.NewCoin("v/0/foo", sdk.NewInt(types.DefaultTotalShareNumber+1)),
		)
	)

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

	campaignKeeper.SetMainnetAccount(sdkCtx, types.MainnetAccount{
		CampaignID: campaign.Id,
		Address:    existAddr.String(),
		Shares:     shares,
	})
	err = bankKeeper.MintCoins(sdkCtx, types.ModuleName, vouchers)
	require.NoError(t, err)
	err = bankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, existAddr, vouchers)
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
				CampaignID: campaign.Id,
				Vouchers:   sample.Coins(),
			},
			err: spnerrors.ErrCritical,
		},
		{
			name: "invalid sender address",
			msg: types.MsgRedeemVouchers{
				Sender:     "invalid_address",
				Account:    addr.String(),
				CampaignID: campaign.Id,
				Vouchers:   vouchers,
			},
			err: spnerrors.ErrCritical,
		},
		{
			name: "insufficient funds",
			msg: types.MsgRedeemVouchers{
				Sender:     addr.String(),
				Account:    addr.String(),
				CampaignID: campaign.Id,
				Vouchers:   vouchersTooBig,
			},
			err: types.ErrInsufficientFunds,
		},
		{
			name: "new account redeem all",
			msg: types.MsgRedeemVouchers{
				Sender:     addr.String(),
				Account:    sample.AccAddress(),
				CampaignID: campaign.Id,
				Vouchers:   vouchers,
			},
		},
		{
			name: "exist account redeem voucher one",
			msg: types.MsgRedeemVouchers{
				Sender:     existAddr.String(),
				Account:    existAddr.String(),
				CampaignID: campaign.Id,
				Vouchers:   sdk.NewCoins(vouchers[0]),
			},
		},
		{
			name: "exist account redeem voucher two",
			msg: types.MsgRedeemVouchers{
				Sender:     existAddr.String(),
				Account:    existAddr.String(),
				CampaignID: campaign.Id,
				Vouchers:   sdk.NewCoins(vouchers[1]),
			},
		},
		{
			name: "exist account redeem voucher three",
			msg: types.MsgRedeemVouchers{
				Sender:     existAddr.String(),
				Account:    existAddr.String(),
				CampaignID: campaign.Id,
				Vouchers:   sdk.NewCoins(vouchers[2]),
			},
		},
		{
			name: "account without funds for vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:     existAddr.String(),
				Account:    existAddr.String(),
				CampaignID: campaign.Id,
				Vouchers:   vouchers,
			},
			err: types.ErrInsufficientFunds,
		},
		{
			name: "account without funds for voucher one",
			msg: types.MsgRedeemVouchers{
				Sender:     existAddr.String(),
				Account:    existAddr.String(),
				CampaignID: campaign.Id,
				Vouchers:   sdk.NewCoins(vouchers[0]),
			},
			err: types.ErrInsufficientFunds,
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

				previousAccount, foundAccount = campaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Account)
				if foundAccount {
					previousBalance = bankKeeper.GetAllBalances(sdkCtx, accountAddr)
				}
			}

			// Execute message
			_, err = campaignSrv.RedeemVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			shares, err := types.VouchersToShares(tc.msg.Vouchers, tc.msg.CampaignID)
			require.NoError(t, err)

			account, found := campaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Account)
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
			balance := bankKeeper.GetAllBalances(sdkCtx, accountAddr)
			require.True(t, expectedVouchers.IsEqual(balance))
		})
	}
}
