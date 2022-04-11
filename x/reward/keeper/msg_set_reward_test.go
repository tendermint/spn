package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func initRewardPool(
	t *testing.T,
	sdkCtx sdk.Context,
	tk testkeeper.TestKeepers,
	ts testkeeper.TestMsgServers,
) types.RewardPool {
	var (
		ctx               = sdk.WrapSDKContext(sdkCtx)
		coordID, provider = ts.CreateCoordinator(ctx, r)
		coins             = sample.Coins(r)
	)

	launchID := tk.LaunchKeeper.AppendChain(sdkCtx, sample.Chain(r, 1, coordID))
	rewardPool := types.RewardPool{
		Provider:            provider.String(),
		LaunchID:            launchID,
		InitialCoins:        coins,
		RemainingCoins:      coins,
		LastRewardHeight:    100,
		CurrentRewardHeight: 30,
		Closed:              false,
	}
	tk.RewardKeeper.SetRewardPool(sdkCtx, rewardPool)

	err := tk.BankKeeper.MintCoins(sdkCtx, types.ModuleName, coins.Add(coins...))
	require.NoError(t, err)
	err = tk.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, coins)
	require.NoError(t, err)

	return rewardPool
}

func TestMsgSetRewards(t *testing.T) {
	var (
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
		invalidCoord   = sample.Address(r)
	)
	invalidCoordMsg := sample.MsgCreateCoordinator(invalidCoord)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &invalidCoordMsg)
	require.NoError(t, err)

	var (
		rewardPool                 = initRewardPool(t, sdkCtx, tk, ts)
		noBalanceRewardPool        = initRewardPool(t, sdkCtx, tk, ts)
		emptyCoinsRewardPool       = initRewardPool(t, sdkCtx, tk, ts)
		zeroRewardHeightRewardPool = initRewardPool(t, sdkCtx, tk, ts)
		launchedRewardPool         = initRewardPool(t, sdkCtx, tk, ts)
	)
	launchTriggeredChain, found := tk.LaunchKeeper.GetChain(sdkCtx, launchedRewardPool.LaunchID)
	require.True(t, found)
	launchTriggeredChain.LaunchTriggered = true
	tk.LaunchKeeper.SetChain(sdkCtx, launchTriggeredChain)

	// setup a chain with no reward pool
	noPoolCoordID, noPoolCoordAddr := ts.CreateCoordinator(ctx, r)
	noPoolChainID := tk.LaunchKeeper.AppendChain(sdkCtx, sample.Chain(r, 0, noPoolCoordID))
	noPoolCoins := sample.Coins(r)
	tk.Mint(sdkCtx, noPoolCoordAddr.String(), noPoolCoins)

	noPoolInsufficientFundsChainID := tk.LaunchKeeper.AppendChain(sdkCtx, sample.Chain(r, 0, noPoolCoordID))

	tests := []struct {
		name string
		msg  types.MsgSetRewards
		err  error
	}{
		{
			name: "should prevent set rewards when chain not found",
			msg: types.MsgSetRewards{
				Provider:         rewardPool.Provider,
				LaunchID:         9999,
				Coins:            rewardPool.RemainingCoins,
				LastRewardHeight: 1000,
			},
			err: launchtypes.ErrChainNotFound,
		},
		{
			name: "should prevent set rewards when coordinator address not found",
			msg: types.MsgSetRewards{
				Provider:         sample.Address(r),
				LaunchID:         rewardPool.LaunchID,
				Coins:            rewardPool.RemainingCoins,
				LastRewardHeight: 1000,
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should prevent set rewards when invalid coordinator id",
			msg: types.MsgSetRewards{
				Provider:         invalidCoord,
				LaunchID:         rewardPool.LaunchID,
				Coins:            rewardPool.RemainingCoins,
				LastRewardHeight: 1000,
			},
			err: types.ErrInvalidCoordinatorID,
		},
		{
			name: "should prevent set rewards when launch is triggered for the chain",
			msg: types.MsgSetRewards{
				Provider:         launchedRewardPool.Provider,
				LaunchID:         launchedRewardPool.LaunchID,
				Coins:            launchedRewardPool.RemainingCoins,
				LastRewardHeight: 1000,
			},
			err: launchtypes.ErrTriggeredLaunch,
		},
		{
			name: "should prevent update rewards in existing pool when coordinator has insufficient funds",
			msg: types.MsgSetRewards{
				Provider:         noBalanceRewardPool.Provider,
				LaunchID:         noBalanceRewardPool.LaunchID,
				Coins:            sample.Coins(r),
				LastRewardHeight: 1000,
			},
			err: sdkerrors.ErrInsufficientFunds,
		},
		{
			name: "should prevent set rewards for new pool when coordinator has insufficient funds",
			msg: types.MsgSetRewards{
				Provider:         noPoolCoordAddr.String(),
				LaunchID:         noPoolInsufficientFundsChainID,
				Coins:            sample.Coins(r),
				LastRewardHeight: 1000,
			},
			err: sdkerrors.ErrInsufficientFunds,
		},
		{
			name: "empty coins should remove the reward pool",
			msg: types.MsgSetRewards{
				Provider:         emptyCoinsRewardPool.Provider,
				LaunchID:         emptyCoinsRewardPool.LaunchID,
				Coins:            sdk.NewCoins(),
				LastRewardHeight: 1000,
			},
		},
		{
			name: "zero reward height should remove the reward pool",
			msg: types.MsgSetRewards{
				Provider:         zeroRewardHeightRewardPool.Provider,
				LaunchID:         zeroRewardHeightRewardPool.LaunchID,
				Coins:            zeroRewardHeightRewardPool.RemainingCoins,
				LastRewardHeight: 0,
			},
		},
		{
			name: "should allows to update rewards in an existent pool",
			msg: types.MsgSetRewards{
				Provider:         rewardPool.Provider,
				LaunchID:         rewardPool.LaunchID,
				Coins:            rewardPool.RemainingCoins,
				LastRewardHeight: 1000,
			},
		},
		{
			name: "should allows to set rewards for a new pool",
			msg: types.MsgSetRewards{
				Provider:         noPoolCoordAddr.String(),
				LaunchID:         noPoolChainID,
				Coins:            noPoolCoins,
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
		provider      = sample.AccAddress(r)
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
				coins:     sample.Coins(r),
				poolCoins: sample.Coins(r),
			},
		},
		{
			name: "empty reward pool",
			args: args{
				provider:  provider,
				coins:     sample.Coins(r),
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
				coins:     sample.Coins(r),
				poolCoins: nil,
			},
		},
		{
			name: "no balance address",
			args: args{
				provider:  sample.AccAddress(r),
				coins:     sample.Coins(r),
				poolCoins: sample.Coins(r),
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
