package keeper_test

import (
	"fmt"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func TestCalculateRewards(t *testing.T) {
	type args struct {
		blockRatio sdk.Dec
		sigRatio   sdk.Dec
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
				blockRatio: tc.Dec(t, "1.000001"),
				sigRatio:   sdk.ZeroDec(),
				coins:      sample.Coins(r),
			},
			wantErr: true,
		},
		{
			name: "prevent using signature ratio greater than 1",
			args: args{
				blockRatio: sdk.ZeroDec(),
				sigRatio:   tc.Dec(t, "1.000001"),
				coins:      sample.Coins(r),
			},
			wantErr: true,
		},
		{
			name: "zero ratios and zero coins should give zero rewards",
			args: args{
				blockRatio: sdk.ZeroDec(),
				sigRatio:   sdk.ZeroDec(),
				coins:      sdk.NewCoins(),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "nil coins should give zero rewards",
			args: args{
				blockRatio: sdk.OneDec(),
				sigRatio:   sdk.OneDec(),
				coins:      nil,
			},
			want: sdk.NewCoins(),
		},
		{
			name: "0 block ratio should give 0 rewards",
			args: args{
				blockRatio: sdk.ZeroDec(),
				sigRatio:   sdk.OneDec(),
				coins:      tc.Coins(t, "10aaa,10bbb,10ccc"),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "0 signature ratio should give 0 rewards",
			args: args{
				blockRatio: sdk.OneDec(),
				sigRatio:   sdk.ZeroDec(),
				coins:      tc.Coins(t, "10aaa,10bbb,10ccc"),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "full block and signature ratios should give all rewards",
			args: args{
				blockRatio: sdk.OneDec(),
				sigRatio:   sdk.OneDec(),
				coins:      tc.Coins(t, "10aaa,10bbb,10ccc"),
			},
			want: tc.Coins(t, "10aaa,10bbb,10ccc"),
		},
		{
			name: "0.5 block ratio should give half rewards",
			args: args{
				blockRatio: tc.Dec(t, "0.5"),
				sigRatio:   sdk.OneDec(),
				coins:      tc.Coins(t, "10aaa,100bbb,1000ccc"),
			},
			want: tc.Coins(t, "5aaa,50bbb,500ccc"),
		},
		{
			name: "0.5 signature ratio should give half rewards",
			args: args{
				blockRatio: sdk.OneDec(),
				sigRatio:   tc.Dec(t, "0.5"),
				coins:      tc.Coins(t, "10aaa,100bbb,1000ccc"),
			},
			want: tc.Coins(t, "5aaa,50bbb,500ccc"),
		},
		{
			name: "0.5 block ratio and 0.4 signature ratio should give 0.2 rewards",
			args: args{
				blockRatio: tc.Dec(t, "0.5"),
				sigRatio:   tc.Dec(t, "0.4"),
				coins:      tc.Coins(t, "10aaa,100bbb,1000ccc"),
			},
			want: tc.Coins(t, "2aaa,20bbb,200ccc"),
		},
		{
			name: "decimal rewards should be truncated",
			args: args{
				blockRatio: tc.Dec(t, "0.5"),
				sigRatio:   sdk.OneDec(),
				coins:      tc.Coins(t, "1aaa,11bbb,101ccc"),
			},
			want: tc.Coins(t, "5bbb,50ccc"),
		},
		{
			name: "0.1 block ratio and 0.1 signature ratio should give 0.01 rewards",
			args: args{
				blockRatio: tc.Dec(t, "0.1"),
				sigRatio:   tc.Dec(t, "0.1"),
				coins:      tc.Coins(t, "10aaa,100bbb,1000ccc"),
			},
			want: tc.Coins(t, "1bbb,10ccc"),
		},
		{
			name: "rewards should be empty coins if all rewards are fully truncated",
			args: args{
				blockRatio: tc.Dec(t, "0.0001"),
				sigRatio:   sdk.OneDec(),
				coins:      tc.Coins(t, "10aaa,100bbb,1000ccc"),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "empty coins should return empty coins",
			args: args{
				blockRatio: sdk.OneDec(),
				sigRatio:   sdk.OneDec(),
				coins:      sdk.NewCoins(),
			},
			want: sdk.NewCoins(),
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
		ctx, tk, _      = testkeeper.NewTestSetup(t)
		valFoo          = sample.Address(r)
		valBar          = sample.Address(r)
		valOpAddrFoo    = sample.Address(r)
		valOpAddrBar    = sample.Address(r)
		noProfileVal    = sample.Address(r)
		notFoundValAddr = sample.Address(r)
		provider        = sample.Address(r)
	)

	// set validator profiles
	tk.ProfileKeeper.SetValidator(ctx, profiletypes.Validator{
		Address:           valFoo,
		OperatorAddresses: []string{valOpAddrFoo},
	})
	tk.ProfileKeeper.SetValidatorByOperatorAddress(ctx, profiletypes.ValidatorByOperatorAddress{

		ValidatorAddress: valFoo,
		OperatorAddress:  valOpAddrFoo,
	})
	tk.ProfileKeeper.SetValidator(ctx, profiletypes.Validator{
		Address:           valBar,
		OperatorAddresses: []string{valOpAddrBar},
	})
	tk.ProfileKeeper.SetValidatorByOperatorAddress(ctx, profiletypes.ValidatorByOperatorAddress{
		ValidatorAddress: valBar,
		OperatorAddress:  valOpAddrBar,
	})
	tk.ProfileKeeper.SetValidatorByOperatorAddress(ctx, profiletypes.ValidatorByOperatorAddress{
		ValidatorAddress: sample.Address(r),
		OperatorAddress:  notFoundValAddr,
	})

	type args struct {
		launchID        uint64
		signatureCounts spntypes.SignatureCounts
		lastBlockHeight int64
		closeRewardPool bool
	}
	tests := []struct {
		name         string
		rewardPool   types.RewardPool
		args         args
		wantBalances map[string]sdk.Coins
		err          error
	}{
		{
			name: "should allow distributing rewards",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 10,
				closeRewardPool: true,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "50aaa,50bbb"),
				valBar:   tc.Coins(t, "50aaa,50bbb"),
			},
		},
		{
			name: "should allow distributing reward with different signature ratios",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,1000bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,1000bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.2"),
					tc.SignatureCount(t, valOpAddrBar, "0.8"),
				),
				lastBlockHeight: 10,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "20aaa,200bbb"),
				valBar:   tc.Coins(t, "80aaa,800bbb"),
			},
		},
		{
			name: "current reward height should influence the block ration for reward distribution",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 10,
				LastRewardHeight:    20,
				Closed:              false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 15,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "25aaa,25bbb"),
				valBar:   tc.Coins(t, "25aaa,25bbb"),
			},
		},
		{
			name: "closing the reward pool should distribute all rewards",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 5,
				closeRewardPool: true,
			},
			wantBalances: map[string]sdk.Coins{
				provider: tc.Coins(t, "50aaa,50bbb"),
				valFoo:   tc.Coins(t, "25aaa,25bbb"),
				valBar:   tc.Coins(t, "25aaa,25bbb"),
			},
		},
		{
			name: "last reward height not reached and reward pool not closed should distribute part of the reward",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 5,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "25aaa,25bbb"),
				valBar:   tc.Coins(t, "25aaa,25bbb"),
			},
		},
		{
			name: "last block height greater than reward pool last reward height should distribute all rewards",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 10,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "50aaa,50bbb"),
				valBar:   tc.Coins(t, "50aaa,50bbb"),
			},
		},
		{
			name: "clamp block height to 1 if ratio GT 1",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 9,
				LastRewardHeight:    10,
				Closed:              false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 11,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "50aaa,50bbb"),
				valBar:   tc.Coins(t, "50aaa,50bbb"),
			},
		},
		{
			name: "rewards for validator with no profile should be distributed to the operator address",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.3"),
					tc.SignatureCount(t, valOpAddrBar, "0.3"),
					tc.SignatureCount(t, noProfileVal, "0.3"),
				),
				lastBlockHeight: 10,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider:     tc.Coins(t, "10aaa,10bbb"),
				valFoo:       tc.Coins(t, "30aaa,30bbb"),
				valBar:       tc.Coins(t, "30aaa,30bbb"),
				noProfileVal: tc.Coins(t, "30aaa,30bbb"),
			},
		},
		{
			name: "rewards should all be refunded if the reward pool is closed and no signature counts are reported",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID:        1,
				signatureCounts: tc.SignatureCounts(1),
				lastBlockHeight: 5,
				closeRewardPool: true,
			},
			wantBalances: map[string]sdk.Coins{
				provider: tc.Coins(t, "100aaa,100bbb"),
			},
		},
		{
			name: "rewards should all be refunded if the last reward height is reached and no signature counts are reported",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID:        1,
				signatureCounts: tc.SignatureCounts(1),
				lastBlockHeight: 10,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: tc.Coins(t, "100aaa,100bbb"),
			},
		},
		{
			name: "reward should be refunded to the provider relative to the block ratio if the reward pool is not closed and no signature counts are reported",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID:        1,
				signatureCounts: tc.SignatureCounts(1),
				lastBlockHeight: 5,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: tc.Coins(t, "50aaa,50bbb"),
			},
		},
		{
			name: "invalid signature counts yields critical error for negative reward pool",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 10,
				LastRewardHeight:    20,
				Closed:              false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.6"),
				),
				lastBlockHeight: 20,
				closeRewardPool: false,
			},
			err: spnerrors.ErrCritical,
		},
		{
			name: "critical error for signatureRatio GT 1",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 10,
				LastRewardHeight:    20,
				Closed:              false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "1.0001"),
				),
				lastBlockHeight: 20,
				closeRewardPool: false,
			},
			err: spnerrors.ErrCritical,
		},
		{
			name: "should prevent distributing rewards with a non-existent reward pool",
			args: args{
				launchID: 99999,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
				),
				lastBlockHeight: 1,
				closeRewardPool: false,
			},
			err: types.ErrRewardPoolNotFound,
		},
		{
			name: "should prevent distributing rewards from a closed reward pool",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           true,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
				),
				lastBlockHeight: 1,
				closeRewardPool: false,
			},
			err: types.ErrRewardPoolClosed,
		},
		{
			name: "prevent distributing rewards if signature counts are invalid",
			rewardPool: types.RewardPool{
				LaunchID:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, "invalid-bech32-address", "0.5"),
				),
				lastBlockHeight: 1,
				closeRewardPool: false,
			},
			err: types.ErrInvalidSignatureCounts,
		},
		{
			name: "prevent providing a last block height lower than the current reward height",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 5,
				LastRewardHeight:    10,
				Closed:              false,
			},
			args: args{
				launchID:        1,
				signatureCounts: tc.SignatureCounts(1),
				lastBlockHeight: 1,
				closeRewardPool: false,
			},
			err: types.ErrInvalidLastBlockHeight,
		},
		{
			name: "prevent providing a last block height equals to the current reward height",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 5,
				LastRewardHeight:    10,
				Closed:              false,
			},
			args: args{
				launchID:        1,
				signatureCounts: tc.SignatureCounts(1),
				lastBlockHeight: 5,
				closeRewardPool: false,
			},
			err: types.ErrInvalidLastBlockHeight,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set test reward pool if contains coins
			if tt.rewardPool.RemainingCoins != nil {
				tk.RewardKeeper.SetRewardPool(ctx, tt.rewardPool)
				err := tk.BankKeeper.MintCoins(ctx, types.ModuleName, tt.rewardPool.RemainingCoins)
				require.NoError(t, err)
			}

			err := tk.RewardKeeper.DistributeRewards(ctx,
				tt.args.launchID,
				tt.args.signatureCounts,
				tt.args.lastBlockHeight,
				tt.args.closeRewardPool,
			)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			rewardPool, found := tk.RewardKeeper.GetRewardPool(ctx, tt.args.launchID)
			require.True(t, found)
			require.Equal(t, tt.rewardPool.InitialCoins, rewardPool.InitialCoins)
			require.Equal(t, tt.rewardPool.Provider, rewardPool.Provider)

			// check if reward pool should be closed
			if tt.args.closeRewardPool || tt.args.lastBlockHeight >= rewardPool.LastRewardHeight {
				require.Equal(t, true, rewardPool.Closed)
			} else {
				require.Equal(t, tt.args.lastBlockHeight, rewardPool.CurrentRewardHeight)
			}

			totalDistributedBalances := sdk.NewCoins()
			for wantAddr, wantBalance := range tt.wantBalances {
				t.Run(fmt.Sprintf("check balance %s", wantAddr), func(t *testing.T) {
					wantAcc, err := sdk.AccAddressFromBech32(wantAddr)
					require.NoError(t, err)

					balance := tk.BankKeeper.GetAllBalances(ctx, wantAcc)
					require.True(t, balance.IsEqual(wantBalance),
						fmt.Sprintf("address: %s,  want: %s, got: %s",
							wantAddr, wantBalance.String(), balance.String(),
						),
					)
					totalDistributedBalances = totalDistributedBalances.Add(balance...)

					// remove the test balance
					err = tk.BankKeeper.SendCoinsFromAccountToModule(ctx, wantAcc, types.ModuleName, balance)
					require.NoError(t, err)
					err = tk.BankKeeper.BurnCoins(ctx, types.ModuleName, balance)
					require.NoError(t, err)
				})
			}

			// assert currentRemainingCoins = previousRemainingCoins - distributedRewards
			expectedRemainingCoins, neg := tt.rewardPool.RemainingCoins.SafeSub(totalDistributedBalances)
			require.False(t, neg, "more coins have been distributed than coins in remaining coins %s > %s",
				totalDistributedBalances.String(),
				tt.rewardPool.RemainingCoins.String(),
			)
			require.True(t, rewardPool.RemainingCoins.IsEqual(expectedRemainingCoins), "expected remaining coins %s, got %s",
				expectedRemainingCoins.String(),
				rewardPool.RemainingCoins.String(),
			)

			// remove the reward pool used for the test
			tk.RewardKeeper.RemoveRewardPool(ctx, tt.rewardPool.LaunchID)
		})
	}
}
