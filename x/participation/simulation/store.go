package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	fundraisingkeeper "github.com/tendermint/fundraising/x/fundraising/keeper"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	"github.com/tendermint/spn/x/participation/keeper"
	"github.com/tendermint/spn/x/participation/types"
)

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

func RandomAuction(ctx sdk.Context, r *rand.Rand, fk fundraisingkeeper.Keeper) (auction fundraisingtypes.AuctionI, found bool) {
	auctions := fk.GetAuctions(ctx)
	if len(auctions) == 0 {
		return auction, false
	}

	r.Shuffle(len(auctions), func(i, j int) {
		auctions[i], auctions[j] = auctions[j], auctions[i]
	})

	for _, a := range auctions {
		// auction must not be started and must not be cancelled
		if !a.IsAuctionStarted(ctx.BlockTime()) && a.GetStatus() != fundraisingtypes.AuctionStatusCancelled {
			return a, true
		}
	}

	return auction, false
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

func RandomAccWithAvailableAllocations(ctx sdk.Context, r *rand.Rand,
	k keeper.Keeper,
	accs []simtypes.Account,
	desired uint64,
	auctionID uint64,
) (simtypes.Account, uint64, bool) {
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

		if amt >= desired {
			_, found := k.GetAuctionUsedAllocations(ctx, acc.Address.String(), auctionID)
			if found {
				continue
			}

			return acc, amt, true
		}
	}

	return simtypes.Account{}, 0, false
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
