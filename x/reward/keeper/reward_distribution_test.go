package keeper_test

import (
	"fmt"
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
	type args struct {
		blockRatio sdk.Dec
		sigRatio      sdk.Dec
		coins      sdk.Coins
	}
	tests := []struct {
		name    string
		args    args
		want    sdk.Coins
		wantErr bool
	}{
		{
			name: "prevent using block ratio greater than 1",
			args: args{
				blockRatio: decFromString(t, "1.1"),
				sigRatio:      sdk.NewDec(0),
				coins:      sample.Coins(),
			},
			wantErr: true,
		},
		{
			name: "prevent using signature ratio greater than 1",
			args: args{
				blockRatio: sdk.NewDec(0),
				sigRatio:   decFromString(t, "1.1"),
				coins:      sample.Coins(),
			},
			wantErr: true,
		},
		{
			name: "zero ratios and zero coins should give zero rewards",
			args: args{
				blockRatio: sdk.NewDec(0),
				sigRatio:      sdk.NewDec(0),
				coins:      sdk.NewCoins(),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "nil coins should give zero rewards",
			args: args{
				blockRatio: sdk.NewDec(0),
				sigRatio:      sdk.NewDec(0),
				coins:      nil,
			},
			want: sdk.NewCoins(),
		},
		{
			name: "0 block ratio should give 0 rewards",
			args: args{
				blockRatio: sdk.NewDec(0),
				sigRatio:   sdk.NewDec(1),
				coins:     coinsFromString(t, "10aaa,10bbb,10ccc"),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "0 signature ratio should give 0 rewards",
			args: args{
				blockRatio: sdk.NewDec(1),
				sigRatio:   sdk.NewDec(0),
				coins:     coinsFromString(t, "10aaa,10bbb,10ccc"),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "0.5 block ratio should give half rewards",
			args: args{
				blockRatio: decFromString(t, "0.5"),
				sigRatio:   sdk.NewDec(1),
				coins:      coinsFromString(t, "10aaa,100bbb,1000ccc"),
			},
			want: coinsFromString(t, "5aaa,50bbb,500ccc"),
		},
		{
			name: "0.5 block ratio and 0.4 signature ratio should give 0.2 rewards",
			args: args{
				blockRatio: decFromString(t, "0.5"),
				sigRatio:   decFromString(t, "0.4"),
				coins:      coinsFromString(t, "10aaa,100bbb,1000ccc"),
			},
			want: coinsFromString(t, "2aaa,20bbb,200ccc"),
		},
		{
			name: "decimal rewards should be truncated",
			args: args{
				blockRatio: decFromString(t, "0.5"),
				sigRatio:   sdk.NewDec(1),
				coins:      coinsFromString(t, "1aaa,11bbb,101ccc"),
			},
			want: coinsFromString(t, "5bbb,50ccc"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := keeper.CalculateRewards(tt.args.blockRatio, tt.args.sigRatio, tt.args.coins)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.True(t, got.IsEqual(tt.want),
				fmt.Sprintf("want: %s, got: %s", tt.want.String(), got.String()),
			)
		})
	}
}

func TestKeeper_DistributeRewards(t *testing.T) {
	var (
		k, _, pk, bk, _, _, _, ctx = setupMsgServer(t)
		launchIDs                  = []uint64{1, 2, 3, 4, 5}
		validatorFoo               = sample.Address()
		validatorBar               = sample.Address()
		validatorConsAddrFoo       = sample.ConsAddress()
		validatorConsAddrBar       = sample.ConsAddress()
		validatorConsAddrBaz       = sample.ConsAddress()
		notFoundValidatorAddr      = sample.ConsAddress()
		provider                   = sample.Address()
		coins = coinsFromString(t, "11bar,222baz,3333foo,4444foobar")

		signatureCounts = spntypes.SignatureCounts{
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
	for _, launchID := range launchIDs {
		k.SetRewardPool(ctx, types.RewardPool{
			LaunchID:            launchID,
			Provider:            provider,
			Coins:               coins,
			LastRewardHeight:    5,
			CurrentRewardHeight: 10,
		})
	}
	pk.SetValidator(ctx, profiletypes.Validator{
		Address:            validatorFoo,
		ConsensusAddresses: [][]byte{validatorConsAddrFoo},
	})
	pk.SetValidatorByConsAddress(ctx, profiletypes.ValidatorByConsAddress{
		ValidatorAddress: validatorFoo,
		ConsensusAddress: validatorConsAddrFoo,
	})
	pk.SetValidator(ctx, profiletypes.Validator{
		Address:            validatorBar,
		ConsensusAddresses: [][]byte{validatorConsAddrBar},
	})
	pk.SetValidatorByConsAddress(ctx, profiletypes.ValidatorByConsAddress{
		ValidatorAddress: validatorBar,
		ConsensusAddress: validatorConsAddrBar,
	})
	pk.SetValidatorByConsAddress(ctx, profiletypes.ValidatorByConsAddress{
		ValidatorAddress: sample.Address(),
		ConsensusAddress: notFoundValidatorAddr,
	})

	type args struct {
		launchID        uint64
		signatureCounts spntypes.SignatureCounts
		lastBlockHeight uint64
		closeRewardPool bool
	}
	tests := []struct {
		name        string
		args        args
		wantBalance map[string]sdk.Coins
		err         error
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
			wantBalance: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				validatorBar: coinsFromString(t,"2bar,60baz,936foo,1248foobar"),
				validatorFoo: coinsFromString(t,"2bar,60baz,936foo,1248foobar"),
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
			wantBalance: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				validatorBar: coinsFromString(t,"2bar,60baz,936foo,1248foobar"),
				validatorFoo: coinsFromString(t,"2bar,60baz,936foo,1248foobar"),
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
			wantBalance: map[string]sdk.Coins{
				provider: coinsFromString(t,"2bar,60baz,936foo,1248foobar"),
				validatorBar: coinsFromString(t,"16bar,336baz,5000foo,6668foobar"),
				validatorFoo: coinsFromString(t,"24bar,444baz,6668foo,8888foobar"),
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
			wantBalance: map[string]sdk.Coins{
				provider: coinsFromString(t,"6baz,104foo,138foobar"),
				validatorBar: coinsFromString(t,"4bar,84baz,1250foo,1667foobar"),
				validatorFoo: coinsFromString(t,"6bar,111baz,1667foo,2222foobar"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if rewardPool, found := k.GetRewardPool(ctx, tt.args.launchID); found {
				err := bk.MintCoins(ctx, types.ModuleName, rewardPool.Coins)
				require.NoError(t, err)
			}

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

			for wantAddr, wantBalance := range tt.wantBalance {
				t.Run(fmt.Sprintf("check balance %s", wantAddr), func(t *testing.T) {
					wantAcc, err := sdk.AccAddressFromBech32(wantAddr)
					require.NoError(t, err)

					balance := bk.GetAllBalances(ctx, wantAcc)
					require.Equal(t, wantBalance, balance)

					// remove the test balance
					err = bk.SendCoinsFromAccountToModule(ctx, wantAcc, types.ModuleName, balance)
					require.NoError(t, err)
					err = bk.BurnCoins(ctx, types.ModuleName, balance)
					require.NoError(t, err)
				})
			}
		})
	}
}
