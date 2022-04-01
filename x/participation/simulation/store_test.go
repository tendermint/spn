package simulation_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/require"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	participationsim "github.com/tendermint/spn/x/participation/simulation"
	"github.com/tendermint/spn/x/participation/types"
)

func TestRandomAccWithBalance(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		r          = sample.Rand()
		accs       = simulation.RandomAccounts(r, 5)
	)

	// give one account balance
	newCoins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)))
	err := tk.BankKeeper.MintCoins(ctx, minttypes.ModuleName, newCoins)
	require.NoError(t, err)
	err = tk.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, accs[0].Address, newCoins)
	require.NoError(t, err)

	tests := []struct {
		name         string
		accounts     []simulation.Account
		desiredCoins sdk.Coins
		wantAccount  simulation.Account
		found        bool
	}{
		{
			name:     "no accounts with balance",
			accounts: accs[1:],
			found:    false,
		},
		{
			name:         "one account has balance",
			accounts:     accs,
			desiredCoins: newCoins,
			wantAccount:  accs[0],
			found:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, coins, found := participationsim.RandomAccWithBalance(ctx, r, tk.BankKeeper, tt.accounts, tt.desiredCoins)
			require.Equal(t, tt.found, found)
			if !tt.found {
				return
			}

			require.Equal(t, tt.wantAccount, got)
			require.Equal(t, tt.desiredCoins, coins)
		})
	}
}

func TestRandomAuction(t *testing.T) {
	var (
		r            = sample.Rand()
		accs         = simulation.RandomAccounts(r, 5)
		sellingCoin1 = sample.Coin(r)
	)

	t.Run("no auction to be found that satisfy requirements", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		startTime := ctx.BlockTime().Add(-time.Hour)
		endTime := ctx.BlockTime().Add(time.Hour * 24 * 7)

		// set custom auction with status AuctionStarted
		var allowedBidders []fundraisingtypes.AllowedBidder
		endTimes := []time.Time{endTime}
		ba := &fundraisingtypes.BaseAuction{
			Id:             0,
			Type:           fundraisingtypes.AuctionTypeFixedPrice,
			AllowedBidders: allowedBidders,
			Auctioneer:     accs[0].Address.String(),
			StartTime:      startTime,
			EndTimes:       endTimes,
			Status:         fundraisingtypes.AuctionStatusStarted,
		}
		auction := fundraisingtypes.NewFixedPriceAuction(ba)
		tk.FundraisingKeeper.SetAuction(ctx, auction)

		// set custom auction with status AuctionCancelled
		ba = &fundraisingtypes.BaseAuction{
			Id:             1,
			Type:           fundraisingtypes.AuctionTypeFixedPrice,
			AllowedBidders: allowedBidders,
			Auctioneer:     accs[0].Address.String(),
			StartTime:      startTime,
			EndTimes:       endTimes,
			Status:         fundraisingtypes.AuctionStatusCancelled,
		}
		auction = fundraisingtypes.NewFixedPriceAuction(ba)
		tk.FundraisingKeeper.SetAuction(ctx, auction)

		_, found := participationsim.RandomAuction(ctx, r, *tk.FundraisingKeeper)
		require.False(t, found)
	})

	t.Run("one auction to be found", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		startTime := ctx.BlockTime().Add(time.Hour * 10)
		endTime := ctx.BlockTime().Add(time.Hour * 24 * 7)

		// initialize auction
		tk.Mint(ctx, accs[0].Address.String(), sdk.NewCoins(sellingCoin1))
		auctionID1 := tk.CreateFixedPriceAuction(ctx, r, accs[0].Address.String(), sellingCoin1, startTime, endTime)
		auction1, found := tk.FundraisingKeeper.GetAuction(ctx, auctionID1)
		require.True(t, found)

		got, found := participationsim.RandomAuction(ctx, r, *tk.FundraisingKeeper)
		require.True(t, found)
		require.Equal(t, auction1, got)
	})

	t.Run("no auctions", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		_, found := participationsim.RandomAuction(ctx, r, *tk.FundraisingKeeper)
		require.False(t, found)
	})
}

func TestRandomAuctionWithdrawEnabled(t *testing.T) {
	var (
		ctx, tk, _       = testkeeper.NewTestSetup(t)
		r                = sample.Rand()
		accs             = simulation.RandomAccounts(r, 5)
		withdrawalDelay  = time.Hour * 5
		invalidStartTime = ctx.BlockTime()
		validStartTime   = ctx.BlockTime().Add(-withdrawalDelay).Add(-time.Hour)
		endTime          = ctx.BlockTime().Add(time.Hour * 24 * 7)
		sellingCoin      = sample.Coin(r)
	)

	params := types.DefaultParams()
	params.WithdrawalDelay = withdrawalDelay
	tk.ParticipationKeeper.SetParams(ctx, params)

	t.Run("no auctions", func(t *testing.T) {
		_, found := participationsim.RandomAuctionWithdrawEnabled(ctx, r, *tk.FundraisingKeeper, *tk.ParticipationKeeper)
		require.False(t, found)
	})

	// populate keeper with some invalid auctions
	var allowedBidders []fundraisingtypes.AllowedBidder
	endTimes := []time.Time{endTime}
	ba := &fundraisingtypes.BaseAuction{
		Id:             0,
		Type:           fundraisingtypes.AuctionTypeFixedPrice,
		AllowedBidders: allowedBidders,
		Auctioneer:     accs[0].Address.String(),
		StartTime:      invalidStartTime, // started, but withdrawal delay not reached
		EndTimes:       endTimes,
		Status:         fundraisingtypes.AuctionStatusStarted,
	}
	invalidAuction := fundraisingtypes.NewFixedPriceAuction(ba)
	tk.FundraisingKeeper.SetAuction(ctx, invalidAuction)
	ba = &fundraisingtypes.BaseAuction{
		Id:             1,
		Type:           fundraisingtypes.AuctionTypeFixedPrice,
		AllowedBidders: allowedBidders,
		Auctioneer:     accs[0].Address.String(),
		StartTime:      invalidStartTime.Add(time.Hour), // not yet started, but withdrawal delay not reached
		EndTimes:       endTimes,
		Status:         fundraisingtypes.AuctionStatusStandBy,
	}
	invalidAuction = fundraisingtypes.NewFixedPriceAuction(ba)
	tk.FundraisingKeeper.SetAuction(ctx, invalidAuction)

	t.Run("no auction to be found that satisfy requirements", func(t *testing.T) {
		_, found := participationsim.RandomAuctionWithdrawEnabled(ctx, r, *tk.FundraisingKeeper, *tk.ParticipationKeeper)
		require.False(t, found)
	})

	// add valid auction
	tk.Mint(ctx, accs[0].Address.String(), sdk.NewCoins(sellingCoin))
	validAuctionID := tk.CreateFixedPriceAuction(ctx, r, accs[0].Address.String(), sellingCoin, validStartTime, endTime)
	validAuction, found := tk.FundraisingKeeper.GetAuction(ctx, validAuctionID)
	require.True(t, found)

	t.Run("find auction where withdrawal delay has passed", func(t *testing.T) {
		foundAuction, found := participationsim.RandomAuctionWithdrawEnabled(ctx, r, *tk.FundraisingKeeper, *tk.ParticipationKeeper)
		require.True(t, found)
		require.Equal(t, validAuction, foundAuction)
		require.True(t, ctx.BlockTime().After(foundAuction.GetStartTime().Add(withdrawalDelay)))
	})

	// forcefully set status for auction created above to cancelled
	err := validAuction.SetStatus(fundraisingtypes.AuctionStatusCancelled)
	require.NoError(t, err)
	tk.FundraisingKeeper.SetAuction(ctx, validAuction)

	t.Run("find cancelled auction", func(t *testing.T) {
		foundAuction, found := participationsim.RandomAuctionWithdrawEnabled(ctx, r, *tk.FundraisingKeeper, *tk.ParticipationKeeper)
		require.True(t, found)
		require.Equal(t, validAuction, foundAuction)
		require.Equal(t, fundraisingtypes.AuctionStatusCancelled, foundAuction.GetStatus())
	})
}

func TestRandomAccWithAvailableAllocations(t *testing.T) {
	var (
		ctx, tk, _        = testkeeper.NewTestSetup(t)
		r                 = sample.Rand()
		accs              = simulation.RandomAccounts(r, 5)
		auctionID  uint64 = 0
	)

	allocationPrice := types.AllocationPrice{Bonded: sdk.NewInt(100)}
	params := types.DefaultParams()
	params.AllocationPrice = allocationPrice
	tk.ParticipationKeeper.SetParams(ctx, params)

	tk.DelegateN(ctx, r, accs[0].Address.String(), 100, 10)
	tk.DelegateN(ctx, r, accs[1].Address.String(), 100, 10)

	tk.ParticipationKeeper.SetUsedAllocations(ctx, types.UsedAllocations{
		Address:        accs[1].Address.String(),
		NumAllocations: 2,
	})
	tk.ParticipationKeeper.SetAuctionUsedAllocations(ctx, types.AuctionUsedAllocations{
		Address:        accs[1].Address.String(),
		AuctionID:      auctionID,
		NumAllocations: 2,
		Withdrawn:      false,
	})

	tests := []struct {
		name               string
		accounts           []simulation.Account
		desiredAllocations uint64
		wantAccount        simulation.Account
		found              bool
	}{
		{
			name:               "no accounts with allocations",
			accounts:           accs[2:],
			desiredAllocations: 10,
			found:              false,
		},
		{
			name:               "one account with insufficient allocations",
			accounts:           accs[1:],
			desiredAllocations: 10,
			found:              false,
		},
		{
			name:               "one account has sufficient allocations",
			accounts:           accs,
			desiredAllocations: 10,
			wantAccount:        accs[0],
			found:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, allocations, found := participationsim.RandomAccWithAvailableAllocations(
				ctx,
				r,
				*tk.ParticipationKeeper,
				tt.accounts,
				tt.desiredAllocations,
				auctionID,
			)
			require.Equal(t, tt.found, found)
			if !tt.found {
				return
			}

			require.Equal(t, tt.wantAccount, got)
			require.Equal(t, tt.desiredAllocations, allocations)
		})
	}
}

func TestRandomAccWithAuctionUsedAllocationsNotWithdrawn(t *testing.T) {
	var (
		ctx, tk, _        = testkeeper.NewTestSetup(t)
		r                 = sample.Rand()
		accs              = simulation.RandomAccounts(r, 5)
		auctionID  uint64 = 0
	)

	allocationPrice := types.AllocationPrice{Bonded: sdk.NewInt(100)}
	params := types.DefaultParams()
	params.AllocationPrice = allocationPrice
	tk.ParticipationKeeper.SetParams(ctx, params)

	// add some delegations
	tk.DelegateN(ctx, r, accs[0].Address.String(), 100, 10)
	tk.DelegateN(ctx, r, accs[1].Address.String(), 100, 10)
	tk.DelegateN(ctx, r, accs[2].Address.String(), 100, 10)

	// add withdrawn allocations for accs[1]
	tk.ParticipationKeeper.SetUsedAllocations(ctx, types.UsedAllocations{
		Address:        accs[1].Address.String(),
		NumAllocations: 2,
	})
	tk.ParticipationKeeper.SetAuctionUsedAllocations(ctx, types.AuctionUsedAllocations{
		Address:        accs[1].Address.String(),
		AuctionID:      auctionID,
		NumAllocations: 2,
		Withdrawn:      true,
	})

	// add used allocations not yet withdrawn for accs[2]
	tk.ParticipationKeeper.SetUsedAllocations(ctx, types.UsedAllocations{
		Address:        accs[2].Address.String(),
		NumAllocations: 2,
	})
	tk.ParticipationKeeper.SetAuctionUsedAllocations(ctx, types.AuctionUsedAllocations{
		Address:        accs[2].Address.String(),
		AuctionID:      auctionID,
		NumAllocations: 2,
		Withdrawn:      false,
	})

	tests := []struct {
		name        string
		accounts    []simulation.Account
		wantAccount simulation.Account
		found       bool
	}{
		{
			name:     "no accounts with allocations",
			accounts: accs[3:],
			found:    false,
		},
		{
			name:     "no account with used allocations that can be withdrawn",
			accounts: accs[:2],
			found:    false,
		},
		{
			name:        "one account has allocations that can be withdrawn",
			accounts:    accs,
			wantAccount: accs[2],
			found:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := participationsim.RandomAccWithAuctionUsedAllocationsNotWithdrawn(
				ctx,
				r,
				*tk.ParticipationKeeper,
				tt.accounts,
				auctionID,
			)
			require.Equal(t, tt.found, found)
			if !tt.found {
				return
			}

			require.Equal(t, tt.wantAccount, got)
		})
	}
}

func TestRandomTierFromList(t *testing.T) {
	r := sample.Rand()

	// find the existing 1 tier
	tierList := []types.Tier{
		{
			TierID:              1,
			RequiredAllocations: 10,
			Benefits:            types.TierBenefits{},
		},
	}

	tier, found := participationsim.RandomTierFromList(r, tierList)
	require.True(t, found)
	require.Equal(t, tier, tierList[0])

	// no tier found with empty list
	_, found = participationsim.RandomTierFromList(r, []types.Tier{})
	require.False(t, found)
}
