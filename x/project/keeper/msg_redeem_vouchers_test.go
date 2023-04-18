package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	ignterrors "github.com/ignite/modules/pkg/errors"

	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
)

func TestMsgRedeemVouchers(t *testing.T) {
	var (
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)

		ctx                    = sdk.WrapSDKContext(sdkCtx)
		addr                   = sample.AccAddress(r)
		existAddr              = sample.AccAddress(r)
		project                = sample.Project(r, 0)
		projectMainnetLaunched = sample.Project(r, 1)
		shares                 types.Shares
		vouchers               sdk.Coins
		err                    error
		vouchersTooBig         = sdk.NewCoins(
			sdk.NewCoin("v/0/foo", sdkmath.NewInt(spntypes.TotalShareNumber+1)),
		)
	)

	t.Run("should allow creation of valid shares", func(t *testing.T) {
		shares, err = types.NewShares("1000foo,500bar,300foobar")
		require.NoError(t, err)
	})

	// Set projects
	project.AllocatedShares = shares
	project.ProjectID = tk.ProjectKeeper.AppendProject(sdkCtx, project)

	projectMainnetLaunched.MainnetInitialized = true
	projectMainnetLaunched.AllocatedShares = shares
	chainLaunched := sample.Chain(r, 0, 0)
	chainLaunched.LaunchTriggered = true
	chainLaunched.IsMainnet = true
	projectMainnetLaunched.MainnetID = tk.LaunchKeeper.AppendChain(sdkCtx, chainLaunched)
	projectMainnetLaunched.ProjectID = tk.ProjectKeeper.AppendProject(sdkCtx, projectMainnetLaunched)

	t.Run("should allow creation of valid vouchers", func(t *testing.T) {
		vouchers, err = types.SharesToVouchers(shares, project.ProjectID)
		require.NoError(t, err)
	})

	t.Run("should allow setting test balances", func(t *testing.T) {
		err = tk.BankKeeper.MintCoins(sdkCtx, types.ModuleName, vouchers)
		require.NoError(t, err)
		err = tk.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, addr, vouchers)
		require.NoError(t, err)

		tk.ProjectKeeper.SetMainnetAccount(sdkCtx, types.MainnetAccount{
			ProjectID: project.ProjectID,
			Address:   existAddr.String(),
			Shares:    shares,
		})
		err = tk.BankKeeper.MintCoins(sdkCtx, types.ModuleName, vouchers)
		require.NoError(t, err)
		err = tk.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, existAddr, vouchers)
		require.NoError(t, err)
	})

	for _, tc := range []struct {
		name string
		msg  types.MsgRedeemVouchers
		err  error
	}{
		{
			name: "should allow redeem voucher one",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[0]),
			},
		},
		{
			name: "should allow redeem voucher two",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[1]),
			},
		},
		{
			name: "should allow redeem voucher three",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[2]),
			},
		},
		{
			name: "should allow redeem all",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   sample.Address(r),
				ProjectID: project.ProjectID,
				Vouchers:  vouchers,
			},
		},
		{
			name: "should fail with non existing project",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectID: 10000,
				Vouchers:  sample.Coins(r),
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail with invalid vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sample.Coins(r),
			},
			err: ignterrors.ErrCritical,
		},
		{
			name: "should fail with invalid sender address",
			msg: types.MsgRedeemVouchers{
				Sender:    "invalid_address",
				Account:   addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  vouchers,
			},
			err: ignterrors.ErrCritical,
		},
		{
			name: "should fail with insufficient funds",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  vouchersTooBig,
			},
			err: types.ErrInsufficientVouchers,
		},

		{
			name: "should fail with account without funds for vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  vouchers,
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "should fail with account without funds for voucher one",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[0]),
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "should fail with project with launched mainnet",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectID: projectMainnetLaunched.ProjectID,
				Vouchers:  sample.Coins(r),
			},
			err: types.ErrMainnetLaunchTriggered,
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

				previousAccount, foundAccount = tk.ProjectKeeper.GetMainnetAccount(sdkCtx, tc.msg.ProjectID, tc.msg.Account)
				if foundAccount {
					previousBalance = tk.BankKeeper.GetAllBalances(sdkCtx, accountAddr)
				}
			}

			// Execute message
			_, err = ts.ProjectSrv.RedeemVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			shares, err := types.VouchersToShares(tc.msg.Vouchers, tc.msg.ProjectID)
			require.NoError(t, err)

			account, found := tk.ProjectKeeper.GetMainnetAccount(sdkCtx, tc.msg.ProjectID, tc.msg.Account)
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
				expectedVouchers, negative = previousBalance.SafeSub(tc.msg.Vouchers...)
				require.False(t, negative)
			}
			balance := tk.BankKeeper.GetAllBalances(sdkCtx, accountAddr)
			require.True(t, expectedVouchers.IsEqual(balance))
		})
	}
}
