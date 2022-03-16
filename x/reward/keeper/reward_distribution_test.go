package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	spntypes "github.com/tendermint/spn/pkg/types"
	tc "github.com/tendermint/spn/testutil/constructor"
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
				blockRatio: tc.Dec(t, "1.1"),
				sigRatio:   sdk.ZeroDec(),
				coins:      sample.Coins(),
			},
			wantErr: true,
		},
		{
			name: "prevent using signature ratio greater than 1",
			args: args{
				blockRatio: sdk.ZeroDec(),
				sigRatio:   tc.Dec(t, "1.1"),
				coins:      sample.Coins(),
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
				blockRatio: sdk.ZeroDec(),
				sigRatio:   sdk.ZeroDec(),
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
			name: "0.5 block ratio should give half rewards",
			args: args{
				blockRatio: tc.Dec(t, "0.5"),
				sigRatio:   sdk.OneDec(),
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
		valFoo          = sample.Address()
		valBar          = sample.Address()
		valOpAddrFoo    = sample.Address()
		valOpAddrBar    = sample.Address()
		noProfileVal    = sample.Address()
		notFoundValAddr = sample.Address()
		provider        = sample.Address()
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
		ValidatorAddress: sample.Address(),
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
			name: "valid close reward pool",
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
			name: "valid close reward pool with lower last block height",
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
			name: "valid distribute rewards without close",
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
			name: "valid distribute rewards with last block height greater than reward pool last reward height",
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
				provider: tc.Coins(t, "10aaa,10bbb"),
				valFoo:   tc.Coins(t, "30aaa,30bbb"),
				valBar:   tc.Coins(t, "30aaa,30bbb"),
				noProfileVal:   tc.Coins(t, "30aaa,30bbb"),
			},
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
			coinTotal := rewardPool.RemainingCoins.Add(totalDistributedBalances...)
			require.True(t, tt.rewardPool.RemainingCoins.IsEqual(coinTotal))

			// remove the reward pool used for the test
			tk.RewardKeeper.RemoveRewardPool(ctx, tt.rewardPool.LaunchID)
		})
	}
}
