package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func initRewardPool(
	t *testing.T,
	k *keeper.Keeper,
	bk bankkeeper.Keeper,
	lk *launchkeeper.Keeper,
	sdkCtx sdk.Context,
	psrv profiletypes.MsgServer,
) types.RewardPool {
	var (
		provider = sample.AccAddress()
		coins    = sample.Coins()
		ctx      = sdk.WrapSDKContext(sdkCtx)
		coordMsg = sample.MsgCreateCoordinator(provider.String())
	)

	res, err := psrv.CreateCoordinator(ctx, &coordMsg)
	require.NoError(t, err)
	launchID := lk.AppendChain(sdkCtx, sample.Chain(1, res.CoordinatorID))
	rewardPool := types.RewardPool{
		Provider:            provider.String(),
		LaunchID:            launchID,
		InitialCoins:        coins,
		RemainingCoins:      coins,
		LastRewardHeight:    100,
		CurrentRewardHeight: 30,
		Closed:              false,
	}
	k.SetRewardPool(sdkCtx, rewardPool)

	err = bk.MintCoins(sdkCtx, types.ModuleName, coins.Add(coins...))
	require.NoError(t, err)
	err = bk.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, coins)
	require.NoError(t, err)

	return rewardPool
}

func TestMsgSetRewards(t *testing.T) {
	var (
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)

		ctx          = sdk.WrapSDKContext(sdkCtx)
		invalidCoord = sample.Address()
	)
	invalidCoordMsg := sample.MsgCreateCoordinator(invalidCoord)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &invalidCoordMsg)
	require.NoError(t, err)

	var (
		rewardPool                 = initRewardPool(t, tk.RewardKeeper, tk.BankKeeper, tk.LaunchKeeper, sdkCtx, ts.ProfileSrv)
		noBalanceRewardPool        = initRewardPool(t, tk.RewardKeeper, tk.BankKeeper, tk.LaunchKeeper, sdkCtx, ts.ProfileSrv)
		emptyCoinsRewardPool       = initRewardPool(t, tk.RewardKeeper, tk.BankKeeper, tk.LaunchKeeper, sdkCtx, ts.ProfileSrv)
		zeroRewardHeightRewardPool = initRewardPool(t, tk.RewardKeeper, tk.BankKeeper, tk.LaunchKeeper, sdkCtx, ts.ProfileSrv)
		launchedRewardPool         = initRewardPool(t, tk.RewardKeeper, tk.BankKeeper, tk.LaunchKeeper, sdkCtx, ts.ProfileSrv)
	)
	launchTriggeredChain, found := tk.LaunchKeeper.GetChain(sdkCtx, launchedRewardPool.LaunchID)
	require.True(t, found)
	launchTriggeredChain.LaunchTriggered = true
	tk.LaunchKeeper.SetChain(sdkCtx, launchTriggeredChain)

	tests := []struct {
		name string
		msg  types.MsgSetRewards
		err  error
	}{
		{
			name: "invalid chain",
			msg: types.MsgSetRewards{
				Provider:         rewardPool.Provider,
				LaunchID:         9999,
				Coins:            rewardPool.RemainingCoins,
				LastRewardHeight: 1000,
			},
			err: launchtypes.ErrChainNotFound,
		},
		{
			name: "coordinator address not found",
			msg: types.MsgSetRewards{
				Provider:         sample.Address(),
				LaunchID:         rewardPool.LaunchID,
				Coins:            rewardPool.RemainingCoins,
				LastRewardHeight: 1000,
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator id",
			msg: types.MsgSetRewards{
				Provider:         invalidCoord,
				LaunchID:         rewardPool.LaunchID,
				Coins:            rewardPool.RemainingCoins,
				LastRewardHeight: 1000,
			},
			err: types.ErrInvalidCoordinatorID,
		},
		{
			name: "launch triggered chain",
			msg: types.MsgSetRewards{
				Provider:         launchedRewardPool.Provider,
				LaunchID:         launchedRewardPool.LaunchID,
				Coins:            launchedRewardPool.RemainingCoins,
				LastRewardHeight: 1000,
			},
			err: launchtypes.ErrTriggeredLaunch,
		},
		{
			name: "coordinator with insufficient funds",
			msg: types.MsgSetRewards{
				Provider:         noBalanceRewardPool.Provider,
				LaunchID:         noBalanceRewardPool.LaunchID,
				Coins:            sample.Coins(),
				LastRewardHeight: 1000,
			},
			err: sdkerrors.ErrInsufficientFunds,
		},
		{
			name: "empty coins",
			msg: types.MsgSetRewards{
				Provider:         emptyCoinsRewardPool.Provider,
				LaunchID:         emptyCoinsRewardPool.LaunchID,
				Coins:            sdk.NewCoins(),
				LastRewardHeight: 1000,
			},
		},
		{
			name: "zero reward height",
			msg: types.MsgSetRewards{
				Provider:         zeroRewardHeightRewardPool.Provider,
				LaunchID:         zeroRewardHeightRewardPool.LaunchID,
				Coins:            zeroRewardHeightRewardPool.RemainingCoins,
				LastRewardHeight: 0,
			},
		},
		{
			name: "valid message",
			msg: types.MsgSetRewards{
				Provider:         rewardPool.Provider,
				LaunchID:         rewardPool.LaunchID,
				Coins:            rewardPool.RemainingCoins,
				LastRewardHeight: 1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			previousRewardPool, _ := tk.RewardKeeper.GetRewardPool(sdkCtx, tt.msg.LaunchID)
			got, err := ts.RewardSrv.SetRewards(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			require.Equal(t, previousRewardPool.RemainingCoins, got.PreviousCoins)
			require.Equal(t, previousRewardPool.LastRewardHeight, got.PreviousLastRewardHeight)

			rewardPool, found := tk.RewardKeeper.GetRewardPool(sdkCtx, tt.msg.LaunchID)
			if tt.msg.Coins.Empty() || tt.msg.LastRewardHeight == 0 {
				require.False(t, found)
				require.Equal(t, int64(0), got.NewLastRewardHeight)
				require.Equal(t, sdk.NewCoins(), got.NewCoins)
				return
			}
			require.True(t, found)
			require.False(t, rewardPool.Closed)
			require.Equal(t, tt.msg.Coins, rewardPool.InitialCoins)
			require.Equal(t, tt.msg.Coins, rewardPool.RemainingCoins)
			require.Equal(t, tt.msg.Provider, rewardPool.Provider)
			require.Equal(t, tt.msg.LastRewardHeight, rewardPool.LastRewardHeight)

			require.Equal(t, tt.msg.Coins, got.NewCoins)
			require.Equal(t, tt.msg.LastRewardHeight, got.NewLastRewardHeight)
		})
	}
}

func TestSetBalance(t *testing.T) {
	var (
		sdkCtx, tk, _ = testkeeper.NewTestSetup(t)

		provider = sample.AccAddress()
	)
	type args struct {
		provider  sdk.AccAddress
		coins     sdk.Coins
		poolCoins sdk.Coins
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "set new balance",
			args: args{
				provider:  provider,
				coins:     sample.Coins(),
				poolCoins: sample.Coins(),
			},
		},
		{
			name: "empty reward pool",
			args: args{
				provider:  provider,
				coins:     sample.Coins(),
				poolCoins: sdk.NewCoins(),
			},
		},
		{
			name: "equal coins and pool coins",
			args: args{
				provider:  provider,
				coins:     tc.Coins(t, "101aaa,102bbb"),
				poolCoins: tc.Coins(t, "101aaa,102bbb"),
			},
		},
		{
			name: "extra coin",
			args: args{
				provider:  provider,
				coins:     tc.Coins(t, "101aaa,102bbb"),
				poolCoins: tc.Coins(t, "33aaa,22bbb"),
			},
		},
		{
			name: "extra pool coin",
			args: args{
				provider:  provider,
				coins:     tc.Coins(t, "101aaa,102bbb"),
				poolCoins: tc.Coins(t, "33aaa,22bbb,11ccc"),
			},
		},
		{
			name: "nil reward pool",
			args: args{
				provider:  provider,
				coins:     sample.Coins(),
				poolCoins: nil,
			},
		},
		{
			name: "no balance address",
			args: args{
				provider:  sample.AccAddress(),
				coins:     sample.Coins(),
				poolCoins: sample.Coins(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tk.BankKeeper.MintCoins(sdkCtx, types.ModuleName, tt.args.coins.Add(tt.args.poolCoins...))
			require.NoError(t, err)
			err = tk.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, tt.args.coins)
			require.NoError(t, err)

			err = keeper.SetBalance(sdkCtx, tk.BankKeeper, tt.args.provider, tt.args.coins, tt.args.poolCoins)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
