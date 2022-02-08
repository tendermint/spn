package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/reward/keeper"
)

func TestCalculateReward(t *testing.T) {
	var (
		coinA = sdk.NewCoin("abc", sdk.NewInt(9999999))
		coinB = sdk.NewCoin("bcd", sdk.NewInt(3000))
		coinC = sdk.NewCoin("cde", sdk.NewInt(10))
		coins = sdk.NewCoins(coinA, coinB, coinC)
	)
	type args struct {
		blockRatio float64
		ratio      float64
		coins      sdk.Coins
	}
	tests := []struct {
		name    string
		args    args
		want    sdk.Coins
		wantErr bool
	}{
		{
			name: "test zero values",
			args: args{
				blockRatio: 0,
				ratio:      0,
				coins:      sdk.NewCoins(),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "test nil coins",
			args: args{
				blockRatio: 0,
				ratio:      0,
				coins:      nil,
			},
			want: sdk.NewCoins(),
		},
		{
			name: "negative coin amount",
			args: args{
				blockRatio: 1000,
				ratio:      100,
				coins:      sample.Coins(),
			},
			wantErr: true,
		},
		{
			name: "valid case",
			args: args{
				blockRatio: 0.6,
				ratio:      0.34,
				coins:      coins,
			},
			want: sdk.NewCoins(
				sdk.NewCoin(coinA.Denom, sdk.NewInt(7960000)),
				sdk.NewCoin(coinB.Denom, sdk.NewInt(2388)),
				sdk.NewCoin(coinC.Denom, sdk.NewInt(8)),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := keeper.CalculateReward(tt.args.blockRatio, tt.args.ratio, tt.args.coins)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestKeeper_DistributeRewards(t *testing.T) {
	var (
		k, _, _, _, _, _, _, sdkCtx = setupMsgServer(t)
		//ctx = sdk.WrapSDKContext(sdkCtx)
	)

	type args struct {
		launchID        uint64
		signatureCounts spntypes.SignatureCounts
		lastBlockHeight uint64
		closeRewardPool bool
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
			err := k.DistributeRewards(sdkCtx, tt.args.launchID, tt.args.signatureCounts, tt.args.lastBlockHeight, tt.args.closeRewardPool)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
