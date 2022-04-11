package simulation

import (
	"errors"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

// FindCoordinatorAccount find a sim account for a coordinator that exists or not
func FindCoordinatorAccount(
	r *rand.Rand,
	ctx sdk.Context,
	k keeper.Keeper,
	accs []simtypes.Account,
	exist bool,
) (simtypes.Account, bool) {
	// Randomize the set for coordinator operation entropy
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})

	for _, acc := range accs {
		coordByAddress, err := k.GetCoordinatorByAddress(ctx, acc.Address.String())
		found := !errors.Is(err, types.ErrCoordAddressNotFound)
		if found == exist {
			coord, found := k.GetCoordinator(ctx, coordByAddress.CoordinatorID)
			if found && !coord.Active {
				continue
			}
			return acc, true
		}
	}
	return simtypes.Account{}, false
}
