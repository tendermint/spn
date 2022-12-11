package simulation

import (
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

// IsLaunchTriggeredChain check if chain is launch triggered
func IsLaunchTriggeredChain(ctx sdk.Context, k keeper.Keeper, chainID uint64) bool {
	chain, found := k.GetChain(ctx, chainID)
	if !found {
		return false
	}
	return chain.LaunchTriggered
}

// FindAccount find account by string hex address
func FindAccount(accs []simtypes.Account, address string) (simtypes.Account, error) {
	coordAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return simtypes.Account{}, err
	}
	simAccount, found := simtypes.FindAccount(accs, coordAddr)
	if !found {
		return simAccount, fmt.Errorf("address %s not found in the sim accounts", address)
	}
	return simAccount, nil
}

// RandomAccWithBalance returns random account with the desired available balance
func RandomAccWithBalance(ctx sdk.Context, r *rand.Rand,
	bk types.BankKeeper,
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

// FindCoordinatorProject finds a project associated with a coordinator
// and returns if it is associated with a chain
func FindCoordinatorProject(
	r *rand.Rand,
	ctx sdk.Context,
	ck types.ProjectKeeper,
	coordID uint64,
	chainID uint64,
) (uint64, bool) {
	projects := ck.GetAllProject(ctx)

	campNb := len(projects)
	if campNb == 0 {
		return 0, false
	}

	// Randomize the set
	r.Shuffle(len(projects), func(i, j int) {
		projects[i], projects[j] = projects[j], projects[i]
	})

	// check if project is already associated with chain
	for _, project := range projects {
		if project.CoordinatorID == coordID {
			// get chain ids
			projectChains, hasChains := ck.GetProjectChains(ctx, project.ProjectID)
			if !hasChains {
				return project.ProjectID, true
			}

			for _, projectChain := range projectChains.Chains {
				if projectChain == chainID {
					return 0, false
				}
			}

			return project.ProjectID, true
		}
	}

	return 0, false
}

// FindChainCoordinatorAccount find coordinator account by chain id
func FindChainCoordinatorAccount(
	ctx sdk.Context,
	k keeper.Keeper,
	accs []simtypes.Account,
	chainID uint64,
) (simtypes.Account, error) {
	chain, found := k.GetChain(ctx, chainID)
	if !found {
		// No message if no coordinator address
		return simtypes.Account{}, fmt.Errorf("chain %d not found", chainID)
	}
	coord, found := k.GetProfileKeeper().GetCoordinator(ctx, chain.CoordinatorID)
	if !found {
		return simtypes.Account{}, fmt.Errorf("coordinator %d not found", chain.CoordinatorID)
	}

	if !coord.Active {
		return simtypes.Account{}, fmt.Errorf("coordinator %d inactive", chain.CoordinatorID)
	}

	return FindAccount(accs, coord.Address)
}

// FindRandomChain find a random chain from store
func FindRandomChain(
	r *rand.Rand,
	ctx sdk.Context,
	k keeper.Keeper,
	launchTriggered,
	noMainnet bool,
) (chain types.Chain, found bool) {
	chains := k.GetAllChain(ctx)
	r.Shuffle(len(chains), func(i, j int) {
		chains[i], chains[j] = chains[j], chains[i]
	})
	for _, c := range chains {
		if c.LaunchTriggered != launchTriggered {
			continue
		}
		if noMainnet && c.IsMainnet {
			continue
		}
		// check if the coordinator is still in the store and active
		coord, coordFound := k.GetProfileKeeper().GetCoordinator(ctx, c.CoordinatorID)
		if !coordFound || !coord.Active {
			continue
		}
		chain = c
		found = true
		break
	}
	return chain, found
}

// FindRandomRequest find a valid random request from store
func FindRandomRequest(
	r *rand.Rand,
	ctx sdk.Context,
	k keeper.Keeper,
) (request types.Request, found bool) {
	// Select a random request without launch triggered
	requests := k.GetAllRequest(ctx)
	r.Shuffle(len(requests), func(i, j int) {
		requests[i], requests[j] = requests[j], requests[i]
	})
	for _, req := range requests {
		if req.Status != types.Request_PENDING {
			continue
		}

		chain, chainFound := k.GetChain(ctx, req.LaunchID)
		if !chainFound || chain.LaunchTriggered {
			continue
		}
		// check if the coordinator is still in the store and active
		coord, coordFound := k.GetProfileKeeper().GetCoordinator(ctx, chain.CoordinatorID)
		if !coordFound || !coord.Active {
			continue
		}

		switch content := req.Content.Content.(type) {
		case *types.RequestContent_ValidatorRemoval:
			// if is validator removal, check if the validator exist
			if _, found := k.GetGenesisValidator(
				ctx,
				chain.LaunchID,
				content.ValidatorRemoval.ValAddress,
			); !found {
				continue
			}
		case *types.RequestContent_AccountRemoval:
			// if is account removal, check if account exist
			found, err := keeper.CheckAccount(ctx, k, chain.LaunchID, content.AccountRemoval.Address)
			if err != nil || !found {
				continue
			}
		}

		return req, true
	}

	return request, false
}

// FindRandomValidator find a valid validator from store
func FindRandomValidator(
	r *rand.Rand,
	ctx sdk.Context,
	k keeper.Keeper,
	accs []simtypes.Account,
) (simAccount simtypes.Account, valAcc types.GenesisValidator, found bool) {
	valAccs := k.GetAllGenesisValidator(ctx)
	r.Shuffle(len(valAccs), func(i, j int) {
		valAccs[i], valAccs[j] = valAccs[j], valAccs[i]
	})
	for _, acc := range valAccs {
		if IsLaunchTriggeredChain(ctx, k, acc.LaunchID) {
			continue
		}
		// get coordinator account for removal
		var err error
		simAccount, err = FindChainCoordinatorAccount(ctx, k, accs, acc.LaunchID)
		if err != nil {
			continue
		}
		valAcc = acc
		found = true
		break
	}
	return simAccount, valAcc, found
}
