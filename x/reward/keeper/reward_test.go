package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
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
		k, _, pk, bk, _, _, _, ctx = setupMsgServer(t)
		launchIDs                  = []uint64{1, 2, 3, 4, 5}
		validatorFoo               = sample.Address()
		validatorBar               = sample.Address()
		validatorConsAddrFoo       = sample.Address()
		validatorConsAddrBar       = sample.Address()
		validatorConsAddrBaz       = sample.Address()
		notFoundValidatorAddr      = sample.Address()
		provider                   = sample.Address()
		coins                      = sample.Coins()
		signatureCounts            = spntypes.SignatureCounts{
			BlockCount: 2,
			Counts: []spntypes.SignatureCount{
				{ConsAddress: validatorConsAddrFoo, RelativeSignatures: sdk.NewDec(1)},
				{ConsAddress: validatorConsAddrBar, RelativeSignatures: sdk.NewDecWithPrec(5, 1)},
				{ConsAddress: validatorConsAddrBaz, RelativeSignatures: sdk.NewDecWithPrec(55, 1)},
			},
		}
		signatureCountsValNotFound = signatureCounts
	)
	signatureCountsValNotFound.Counts = append(signatureCountsValNotFound.Counts, spntypes.SignatureCount{
		ConsAddress: notFoundValidatorAddr, RelativeSignatures: sdk.NewDec(1),
	})
	moduleSupply := coins
	for _, launchID := range launchIDs {
		k.SetRewardPool(ctx, types.RewardPool{
			LaunchID:            launchID,
			Provider:            provider,
			Coins:               coins,
			LastRewardHeight:    1,
			CurrentRewardHeight: 4,
		})
		moduleSupply = moduleSupply.Add(moduleSupply...)
	}
	pk.SetValidator(ctx, profiletypes.Validator{
		Address:          validatorFoo,
		ConsensusAddress: validatorConsAddrFoo,
	})
	pk.SetValidatorByConsAddress(ctx, profiletypes.ValidatorByConsAddress{
		ValidatorAddress: validatorFoo,
		ConsensusAddress: validatorConsAddrFoo,
	})
	pk.SetValidator(ctx, profiletypes.Validator{
		Address:          validatorBar,
		ConsensusAddress: validatorConsAddrBar,
	})
	pk.SetValidatorByConsAddress(ctx, profiletypes.ValidatorByConsAddress{
		ValidatorAddress: validatorBar,
		ConsensusAddress: validatorConsAddrBar,
	})
	pk.SetValidatorByConsAddress(ctx, profiletypes.ValidatorByConsAddress{
		ValidatorAddress: sample.Address(),
		ConsensusAddress: notFoundValidatorAddr,
	})
	err := bk.MintCoins(ctx, types.ModuleName, moduleSupply)
	require.NoError(t, err)

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
		{
			name: "invalid reward pool",
			args: args{
				launchID:        99999,
				signatureCounts: signatureCounts,
				lastBlockHeight: 1,
				closeRewardPool: false,
			},
			err: types.ErrRewardPoolNotFound,
		},
		{
			name: "not found validator",
			args: args{
				launchID:        launchIDs[0],
				signatureCounts: signatureCountsValNotFound,
				lastBlockHeight: 1,
				closeRewardPool: false,
			},
			err: spnerrors.ErrCritical,
		},
		{
			name: "valid close reward pool",
			args: args{
				launchID:        launchIDs[1],
				signatureCounts: signatureCounts,
				lastBlockHeight: 2,
				closeRewardPool: true,
			},
		},
		{
			name: "valid close reward pool with lower last block height",
			args: args{
				launchID:        launchIDs[2],
				signatureCounts: signatureCounts,
				lastBlockHeight: 1,
				closeRewardPool: true,
			},
		},
		{
			name: "valid distribute rewards without close",
			args: args{
				launchID:        launchIDs[3],
				signatureCounts: signatureCounts,
				lastBlockHeight: 2,
				closeRewardPool: false,
			},
		},
		{
			name: "valid distribute rewards with high last block height",
			args: args{
				launchID:        launchIDs[4],
				signatureCounts: signatureCounts,
				lastBlockHeight: 3,
				closeRewardPool: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := k.DistributeRewards(ctx, tt.args.launchID, tt.args.signatureCounts, tt.args.lastBlockHeight, tt.args.closeRewardPool)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			rewardPool, found := k.GetRewardPool(ctx, tt.args.launchID)
			if tt.args.closeRewardPool || tt.args.lastBlockHeight >= rewardPool.LastRewardHeight {
				require.False(t, found)
				return
			}
			require.True(t, found)
			require.Equal(t, tt.args.lastBlockHeight, rewardPool.CurrentRewardHeight)
		})
	}
}
