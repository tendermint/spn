package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	"testing"
)

func TestMsgUnredeemVouchers(t *testing.T) {
	var (
		coordAddr1                                            = sample.AccAddress()
		coordAddr2                                            = sample.AccAddress()
		campaignKeeper, _, _, bankKeeper, campaignSrv, profileSrv, sdkCtx = setupMsgServer(t)
		ctx                                                   = sdk.WrapSDKContext(sdkCtx)
	)

	for _, tc := range []struct {
		name string
		msg  types.MsgUnredeemVouchers
		err  error
	} {
		{
			name: "unredeem vouchers",
			msg: types.MsgUnredeemVouchers{
			},
		},
	} {
		var previousAccount types.MainnetAccount
		var previousBalance sdk.Coins
		var found bool

		accountAddr, err := sdk.AccAddressFromBech32(tc.msg.Sender)
		require.NoError(t, err)

		// Get values before message execution
		if tc.err == nil {
			previousAccount, found = campaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Sender)
			require.True(t, found)

			previousBalance = bankKeeper.GetAllBalances(sdkCtx, accountAddr)
		}

		// Execute message
		_, err = campaignSrv.UnredeemVouchers(ctx, &tc.msg)
		if tc.err != nil {
			require.ErrorIs(t, err, tc.err)
			return
		}
		require.NoError(t, err)

		if types.IsEqualShares(tc.msg.Shares, previousAccount.Shares) {
			// All unredeemed
			_, found := campaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Sender)
			require.False(t, found)

		} else {
			account, found := campaignKeeper.GetMainnetAccount(sdkCtx, tc.msg.CampaignID, tc.msg.Sender)
			require.True(t, found)

			expectedShares, err := types.DecreaseShares(previousAccount.Shares, tc.msg.Shares)
			require.NoError(t, err)
			require.True(t, types.IsEqualShares(expectedShares, account.Shares))
		}

		// Compare balance
		unredeemed, err := types.SharesToVouchers(tc.msg.Shares, tc.msg.CampaignID)
		balance := bankKeeper.GetAllBalances(sdkCtx, accountAddr)
		require.True(t, balance.IsEqual(previousBalance.Add(unredeemed...)))
	}
}