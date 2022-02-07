package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func Test_msgServer_SetRewards(t *testing.T) {
	var (
		k, lk, _, bk, srv, psrv, _, sdkCtx = setupMsgServer(t)

		ctx            = sdk.WrapSDKContext(sdkCtx)
		moduleBalance  = sample.Coins()
		newBalance     = sample.Coins()
		provider       = sample.AccAddress()
		invalidCoord   = sample.Address()
		noBalanceCoord = sample.Address()
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

	err = bk.MintCoins(sdkCtx, types.ModuleName, append(moduleBalance, newBalance...).Sort())
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.SetRewards(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			rewardPool, found := k.GetRewardPool(sdkCtx, tt.msg.LaunchID)
			require.True(t, found)
			require.Equal(t, tt.msg.Coins, rewardPool.Coins)
			require.Equal(t, tt.msg.Provider, rewardPool.Provider)
			require.Equal(t, tt.msg.LastRewardHeight, rewardPool.LastRewardHeight)
		})
	}
}

func TestSetBalance(t *testing.T) {
	var (
		_, _, _, bk, _, _, _, sdkCtx = setupMsgServer(t)

		poolCoins     = sample.Coins()
		newBalance    = sample.Coins()
		provider      = sample.AccAddress()
		noBalanceAddr = sample.AccAddress()
	)
	totalBalance := append(newBalance.Add(newBalance...).Add(newBalance...))
	err := bk.MintCoins(sdkCtx, types.ModuleName, append(poolCoins, totalBalance...).Sort())
	require.NoError(t, err)
	err = bk.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, totalBalance)
	require.NoError(t, err)

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
				coins:     poolCoins,
				poolCoins: poolCoins,
			},
		},
		{
			name: "set new balance",
			args: args{
				provider:  provider,
				coins:     newBalance,
				poolCoins: poolCoins,
			},
		},
		{
			name: "empty reward pool",
			args: args{
				provider:  provider,
				coins:     newBalance,
				poolCoins: sdk.NewCoins(),
			},
		},
		{
			name: "nil reward pool",
			args: args{
				provider:  provider,
				coins:     newBalance,
				poolCoins: nil,
			},
		},
		{
			name: "no balance provider",
			args: args{
				provider:  provider,
				coins:     append(newBalance, sample.Coins()...).Sort(),
				poolCoins: poolCoins,
			},
			wantErr: true,
		},
		{
			name: "no balance address",
			args: args{
				provider:  noBalanceAddr,
				coins:     sample.Coins(),
				poolCoins: poolCoins,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := keeper.SetBalance(sdkCtx, bk, tt.args.provider, tt.args.coins, tt.args.poolCoins)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}