package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func TestMsgSetRewards(t *testing.T) {
	var (
		k, lk, _, bk, srv, psrv, _, sdkCtx = setupMsgServer(t)

		ctx            = sdk.WrapSDKContext(sdkCtx)
		provider       = sample.AccAddress()
		invalidCoord   = sample.Address()
		noBalanceCoord = sample.Address()

		emptyCoinsBalance      = sample.Coins()
		zeroRewarHeightBalance = sample.Coins()
		newBalance             = sample.Coins()
		moduleBalance          = sample.Coins().
					Add(emptyCoinsBalance...).
					Add(zeroRewarHeightBalance...).
					Add(newBalance...)
	)
	coordMsg := sample.MsgCreateCoordinator(invalidCoord)
	res, err := psrv.CreateCoordinator(ctx, &coordMsg)
	require.NoError(t, err)

	coordMsg = sample.MsgCreateCoordinator(noBalanceCoord)
	res, err = psrv.CreateCoordinator(ctx, &coordMsg)
	require.NoError(t, err)
	noBalancelaunchID := lk.AppendChain(sdkCtx, sample.Chain(1, res.CoordinatorID))

	coordMsg = sample.MsgCreateCoordinator(provider.String())
	res, err = psrv.CreateCoordinator(ctx, &coordMsg)
	require.NoError(t, err)
	launchID := lk.AppendChain(sdkCtx, sample.Chain(1, res.CoordinatorID))

	launchTriggeredChain := sample.Chain(1, res.CoordinatorID)
	launchTriggeredChain.LaunchTriggered = true
	launchTriggeredChainID := lk.AppendChain(sdkCtx, launchTriggeredChain)

	emptyBalanceLaunchID := lk.AppendChain(sdkCtx, sample.Chain(4, res.CoordinatorID))
	k.SetRewardPool(sdkCtx, types.RewardPool{
		LaunchID:            emptyBalanceLaunchID,
		Coins:               emptyCoinsBalance,
		LastRewardHeight:    100,
		CurrentRewardHeight: 30,
	})
	zeroRewardHeightLaunchID := lk.AppendChain(sdkCtx, sample.Chain(5, res.CoordinatorID))
	k.SetRewardPool(sdkCtx, types.RewardPool{
		LaunchID:            zeroRewardHeightLaunchID,
		Coins:               zeroRewarHeightBalance,
		LastRewardHeight:    100,
		CurrentRewardHeight: 30,
	})

	err = bk.MintCoins(sdkCtx, types.ModuleName, moduleBalance)
	require.NoError(t, err)
	err = bk.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, newBalance)
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  types.MsgSetRewards
		err  error
	}{
		{
			name: "invalid chain",
			msg: types.MsgSetRewards{
				Provider:         provider.String(),
				LaunchID:         9999,
				Coins:            newBalance,
				LastRewardHeight: 1000,
			},
			err: launchtypes.ErrChainNotFound,
		},
		{
			name: "coordinator address not found",
			msg: types.MsgSetRewards{
				Provider:         sample.Address(),
				LaunchID:         launchID,
				Coins:            newBalance,
				LastRewardHeight: 1000,
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator id",
			msg: types.MsgSetRewards{
				Provider:         invalidCoord,
				LaunchID:         launchID,
				Coins:            newBalance,
				LastRewardHeight: 1000,
			},
			err: types.ErrInvalidCoordinatorID,
		},
		{
			name: "launch triggered chain",
			msg: types.MsgSetRewards{
				Provider:         provider.String(),
				LaunchID:         launchTriggeredChainID,
				Coins:            newBalance,
				LastRewardHeight: 1000,
			},
			err: launchtypes.ErrTriggeredLaunch,
		},
		{
			name: "coordinator with insufficient funds",
			msg: types.MsgSetRewards{
				Provider:         noBalanceCoord,
				LaunchID:         noBalancelaunchID,
				Coins:            newBalance,
				LastRewardHeight: 1000,
			},
			err: sdkerrors.ErrInsufficientFunds,
		},
		{
			name: "coordinator with insufficient funds",
			msg: types.MsgSetRewards{
				Provider:         noBalanceCoord,
				LaunchID:         noBalancelaunchID,
				Coins:            newBalance,
				LastRewardHeight: 1000,
			},
			err: sdkerrors.ErrInsufficientFunds,
		},
		{
			name: "empty coins",
			msg: types.MsgSetRewards{
				Provider:         provider.String(),
				LaunchID:         emptyBalanceLaunchID,
				Coins:            emptyCoinsBalance,
				LastRewardHeight: 1000,
			},
		},
		{
			name: "zero reward height",
			msg: types.MsgSetRewards{
				Provider:         provider.String(),
				LaunchID:         zeroRewardHeightLaunchID,
				Coins:            zeroRewarHeightBalance,
				LastRewardHeight: 0,
			},
		},
		{
			name: "valid message",
			msg: types.MsgSetRewards{
				Provider:         provider.String(),
				LaunchID:         launchID,
				Coins:            newBalance,
				LastRewardHeight: 1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			previusRewardPool, _ := k.GetRewardPool(sdkCtx, tt.msg.LaunchID)
			got, err := srv.SetRewards(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			require.Equal(t, previusRewardPool.Coins, got.PreviousCoins)
			require.Equal(t, previusRewardPool.LastRewardHeight, got.PreviousLastRewardHeight)

			rewardPool, found := k.GetRewardPool(sdkCtx, tt.msg.LaunchID)
			if tt.msg.Coins.Empty() || tt.msg.LastRewardHeight == 0 {
				require.False(t, found)
				require.Equal(t, uint64(0), got.NewLastRewardHeight)
				require.Equal(t, sdk.NewCoins(), got.NewCoins)
				return
			}
			require.True(t, found)
			require.Equal(t, tt.msg.Coins, rewardPool.Coins)
			require.Equal(t, tt.msg.Provider, rewardPool.Provider)
			require.Equal(t, tt.msg.LastRewardHeight, rewardPool.LastRewardHeight)

			require.Equal(t, tt.msg.Coins, got.NewCoins)
			require.Equal(t, tt.msg.LastRewardHeight, got.NewLastRewardHeight)
		})
	}
}

func TestSetBalance(t *testing.T) {
	var (
		_, _, _, bk, _, _, _, sdkCtx = setupMsgServer(t)

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
			name: "use the same module balance",
			args: args{
				provider:  provider,
				coins:     sample.Coins(),
				poolCoins: sample.Coins(),
			},
		},
		{
			name: "set new balance",
			args: args{
				provider:  provider,
				coins:     sample.Coins(),
				poolCoins: sample.Coins(),
			},
		},
		{
			name: "set the old balance",
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
			err := bk.MintCoins(sdkCtx, types.ModuleName, tt.args.coins.Add(tt.args.poolCoins...))
			require.NoError(t, err)
			err = bk.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, tt.args.coins)
			require.NoError(t, err)

			err = keeper.SetBalance(sdkCtx, bk, tt.args.provider, tt.args.coins, tt.args.poolCoins)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
