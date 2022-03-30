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
		// auction must not be started
		if a.GetStatus() != fundraisingtypes.AuctionStatusStarted && !a.IsAuctionStarted(ctx.BlockTime()) {
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

func RandomTierFromList(r *rand.Rand, tierList []types.Tier) uint64 {
	return uint64(1 + r.Intn(len(tierList)))
}
