package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"

	launchtypes "github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

// FindRandomChainWithCoordBalance find a random chain from store
func FindRandomChainWithCoordBalance(
	r *rand.Rand,
	ctx sdk.Context,
	k keeper.Keeper,
	bk types.BankKeeper,
	hasRewardPool,
	checkBalance bool,
	wantCoins sdk.Coins,
) (chain launchtypes.Chain, found bool) {
	chains := k.GetLaunchKeeper().GetAllChain(ctx)
	r.Shuffle(len(chains), func(i, j int) {
		chains[i], chains[j] = chains[j], chains[i]
	})
	for _, c := range chains {
		_, poolFound := k.GetRewardPool(ctx, c.LaunchID)
		if hasRewardPool != poolFound {
			continue
		}

		// chain cannot be launch triggered
		if c.LaunchTriggered || c.IsMainnet {
			continue
		}

		// check if the coordinator is still in the store and active
		coord, coordFound := k.GetProfileKeeper().GetCoordinator(ctx, c.CoordinatorID)
		if !coordFound || !coord.Active {
			continue
		}

		coordAccAddr, err := sdk.AccAddressFromBech32(coord.Address)
		if err != nil {
			continue
		}

		if checkBalance {
			balance := bk.SpendableCoins(ctx, coordAccAddr)
			if !balance.IsAllGTE(wantCoins) {
				continue
			}
		}

		chain = c
		found = true
		break
	}
	return chain, found
}
