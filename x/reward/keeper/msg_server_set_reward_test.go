package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func TestSetBalance(t *testing.T) {
	var (
		k, lk, pk, srv, _, _, sdkCtx = setupMsgServer(t)
		ctx                          = sdk.WrapSDKContext(sdkCtx)
	)
	type args struct {
		provider sdk.AccAddress
		coins    sdk.Coins
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := keeper.SetBalance(ctx, authkeeper, bankKeeper, tt.args.provider, tt.args.coins)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func Test_msgServer_SetReward(t *testing.T) {
	var (
		k, lk, pk, srv, _, _, sdkCtx = setupMsgServer(t)
		ctx                          = sdk.WrapSDKContext(sdkCtx)
	)
	tests := []struct {
		name string
		msg  *types.MsgSetReward
		want *types.MsgSetRewardResponse
		err  error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.SetReward(ctx, tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
