package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
)

func TestMsgUnredeemVouchers(t *testing.T) {
	var (
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)

		accountAddr              = sample.Address(r)
		account                  = sample.MainnetAccount(r, 0, accountAddr)
		accountShare, _          = types.NewShares("30foo,30bar")
		accountFewSharesAddr     = sample.Address(r)
		accountFewShares         = sample.MainnetAccount(r, 0, accountFewSharesAddr)
		accountFewSharesShare, _ = types.NewShares("30foo,15bar")

		project                = sample.Project(r, 0)
		projectMainnetLaunched = sample.Project(r, 1)
		shares, _              = types.NewShares("10foo,10bar")
	)
	account.Shares = accountShare
	accountFewShares.Shares = accountFewSharesShare

	// Create projects
	tk.ProjectKeeper.AppendProject(sdkCtx, project)

	projectMainnetLaunched.MainnetInitialized = true
	chainLaunched := sample.Chain(r, 0, 0)
	chainLaunched.LaunchTriggered = true
	chainLaunched.IsMainnet = true
	projectMainnetLaunched.MainnetID = tk.LaunchKeeper.AppendChain(sdkCtx, chainLaunched)
	projectMainnetLaunched.ProjectID = tk.ProjectKeeper.AppendProject(sdkCtx, projectMainnetLaunched)

	// Create accounts
	tk.ProjectKeeper.SetMainnetAccount(sdkCtx, account)
	tk.ProjectKeeper.SetMainnetAccount(sdkCtx, accountFewShares)

	for _, tc := range []struct {
		name string
		msg  types.MsgUnredeemVouchers
		err  error
	}{
		{
			name: "should allow unredeem vouchers",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountAddr,
				ProjectID: 0,
				Shares:    shares,
			},
		},
		{
			name: "should allow unredeem vouchers a second time",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountAddr,
				ProjectID: 0,
				Shares:    shares,
			},
		},
		{
			name: "should allow unredeem vouchers to zero",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountAddr,
				ProjectID: 0,
				Shares:    shares,
			},
		},
		{
			name: "should allow unredeem vouchers from another account",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountFewSharesAddr,
				ProjectID: 0,
				Shares:    shares,
			},
		},
		{
			name: "should prevent if not enough shares in balance",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountFewSharesAddr,
				ProjectID: 0,
				Shares:    shares,
			},
			err: types.ErrSharesDecrease,
		},
		{
			name: "should prevent for non existent project",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountAddr,
				ProjectID: 1000,
				Shares:    shares,
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should prevent for non existent account",
			msg: types.MsgUnredeemVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Shares:    shares,
			},
			err: types.ErrAccountNotFound,
		},
		{
			name: "should prevent for project with launched mainnet",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountAddr,
				ProjectID: projectMainnetLaunched.ProjectID,
				Shares:    sample.Shares(r),
			},
			err: types.ErrMainnetLaunchTriggered,
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
				previousAccount, found = tk.ProjectKeeper.GetMainnetAccount(sdkCtx, tc.msg.ProjectID, tc.msg.Sender)
				require.True(t, found)

				previousBalance = tk.BankKeeper.GetAllBalances(sdkCtx, accountAddr)
			}

			// Execute message
			_, err = ts.ProjectSrv.UnredeemVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			if types.IsEqualShares(tc.msg.Shares, previousAccount.Shares) {
				// All unredeemed
				_, found := tk.ProjectKeeper.GetMainnetAccount(sdkCtx, tc.msg.ProjectID, tc.msg.Sender)
				require.False(t, found)

			} else {
				account, found := tk.ProjectKeeper.GetMainnetAccount(sdkCtx, tc.msg.ProjectID, tc.msg.Sender)
				require.True(t, found)

				expectedShares, err := types.DecreaseShares(previousAccount.Shares, tc.msg.Shares)
				require.NoError(t, err)
				require.True(t, types.IsEqualShares(expectedShares, account.Shares))
			}

			// Compare balance
			unredeemed, err := types.SharesToVouchers(tc.msg.Shares, tc.msg.ProjectID)
			require.NoError(t, err)
			balance := tk.BankKeeper.GetAllBalances(sdkCtx, accountAddr)
			require.True(t, balance.IsEqual(previousBalance.Add(unredeemed...)))
		})
	}
}
