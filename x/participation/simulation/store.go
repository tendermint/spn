package simulation

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	fundraisingkeeper "github.com/tendermint/fundraising/x/fundraising/keeper"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	"github.com/tendermint/spn/x/participation/keeper"
	"github.com/tendermint/spn/x/participation/types"
)

// RandomAccWithBalance returns random account with the desired available balance
func RandomAccWithBalance(ctx sdk.Context, r *rand.Rand,
	bk bankkeeper.Keeper,
	accs []simtypes.Account,
	desired sdk.Coins,
) (simtypes.Account, sdk.Coins, bool) {
	// Randomize the set
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})

	for _, acc := range accs {
		balances := bk.GetAllBalances(ctx, acc.Address)
		if len(balances) == 0 {
			continue
		}

		if balances.IsAllGTE(desired) {
			return acc, balances, true
		}
	}

	return simtypes.Account{}, sdk.NewCoins(), false
}

// RandomAuctionStandby returns random auction that is in standby
func RandomAuctionStandby(ctx sdk.Context, r *rand.Rand, fk fundraisingkeeper.Keeper) (auction fundraisingtypes.AuctionI, found bool) {
	auctions := fk.GetAuctions(ctx)
	if len(auctions) == 0 {
		return auction, false
	}

	r.Shuffle(len(auctions), func(i, j int) {
		auctions[i], auctions[j] = auctions[j], auctions[i]
	})

	for _, a := range auctions {
		// auction must be in standby status
		if a.GetStatus() == fundraisingtypes.AuctionStatusStandBy {
			return a, true
		}
	}

	return auction, false
}

// RandomAuctionParticipationEnabled returns random auction where participation is enabled
func RandomAuctionParticipationEnabled(
	ctx sdk.Context,
	r *rand.Rand,
	fk fundraisingkeeper.Keeper,
	k keeper.Keeper,
) (auction fundraisingtypes.AuctionI, found bool) {
	auctions := fk.GetAuctions(ctx)
	if len(auctions) == 0 {
		return auction, false
	}

	r.Shuffle(len(auctions), func(i, j int) {
		auctions[i], auctions[j] = auctions[j], auctions[i]
	})

	for _, a := range auctions {
		if a.GetStatus() != fundraisingtypes.AuctionStatusStandBy {
			continue
		}
		if !k.IsRegistrationEnabled(ctx, a.GetStartTime()) {
			continue
		}
		auction = a
		found = true
	}

	return
}

// RandomAuctionWithdrawEnabled returns random auction where used allocations can be withdrawn at blockTime
func RandomAuctionWithdrawEnabled(
	ctx sdk.Context,
	r *rand.Rand,
	fk fundraisingkeeper.Keeper,
	k keeper.Keeper,
) (auction fundraisingtypes.AuctionI, found bool) {
	auctions := fk.GetAuctions(ctx)
	withdrawalDelay := k.WithdrawalDelay(ctx)
	if len(auctions) == 0 {
		return auction, false
	}

	r.Shuffle(len(auctions), func(i, j int) {
		auctions[i], auctions[j] = auctions[j], auctions[i]
	})

	for _, a := range auctions {
		// if auction cancelled, withdraw is always enabled
		if a.GetStatus() == fundraisingtypes.AuctionStatusCancelled {
			return a, true
		}

		// check if withdrawal delay has passed and hence withdraw is enabled
		if ctx.BlockTime().After(a.GetStartTime().Add(withdrawalDelay)) {
			return a, true
		}
	}

	return auction, false
}

// RandomAccWithAvailableAllocations returns random account that has at least the desired amount of available allocations
// and can still participate in the specified auction
func RandomAccWithAvailableAllocations(ctx sdk.Context, r *rand.Rand,
	k keeper.Keeper,
	accs []simtypes.Account,
	desired sdkmath.Int,
	auctionID uint64,
) (simtypes.Account, sdkmath.Int, bool) {
	// Randomize the set
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})

	// account must have allocations but not already have participated
	for _, acc := range accs {
		amt, err := k.GetAvailableAllocations(ctx, acc.Address.String())
		if err != nil {
			continue
		}

		if amt.GTE(desired) {
			_, found := k.GetAuctionUsedAllocations(ctx, acc.Address.String(), auctionID)
			if found {
				continue
			}

			return acc, amt, true
		}
	}

	return simtypes.Account{}, sdkmath.ZeroInt(), false
}

// RandomAccWithAuctionUsedAllocationsNotWithdrawn returns random account that has used allocations for the given
// auction that have not yet been withdrawn
func RandomAccWithAuctionUsedAllocationsNotWithdrawn(
	ctx sdk.Context,
	r *rand.Rand,
	k keeper.Keeper,
	accs []simtypes.Account,
	auctionID uint64,
) (simtypes.Account, bool) {
	// Randomize the set
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})

	// account must have used allocations for this auction that have not yet been withdrawn
	for _, acc := range accs {
		usedAllocations, found := k.GetAuctionUsedAllocations(ctx, acc.Address.String(), auctionID)
		if !found || usedAllocations.Withdrawn {
			continue
		}

		return acc, true
	}

	return simtypes.Account{}, false
}

func RandomTierFromList(r *rand.Rand, tierList []types.Tier) (types.Tier, bool) {
	if len(tierList) == 0 {
		return types.Tier{}, false
	}

	index := r.Intn(len(tierList))
	return tierList[index], true
}

func FindLargestMaxBid(tierList []types.Tier) (types.Tier, bool) {
	largestMaxBid := sdkmath.ZeroInt()
	if len(tierList) == 0 {
		return types.Tier{}, false
	}

	index := 0
	for i, tier := range tierList {
		if tier.Benefits.MaxBidAmount.GT(largestMaxBid) {
			largestMaxBid = tier.Benefits.MaxBidAmount
			index = i
		}
	}

	return tierList[index], true
}
