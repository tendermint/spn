package simulation

import (
	"fmt"

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

// FindChain find a chain
func FindChain(ctx sdk.Context, k keeper.Keeper, launchTriggered bool) (types.Chain, bool) {
	found := false
	chains := k.GetAllChain(ctx)
	var chain types.Chain
	for _, c := range chains {
		if c.LaunchTriggered != launchTriggered {
			continue
		}
		// check if the coordinator is still in the store
		_, found = k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, c.CoordinatorID)
		if !found {
			continue
		}
		chain = c
		break
	}
	return chain, found
}

// FindChainCoordinatorAccount find coordinator account by chain id
func FindChainCoordinatorAccount(ctx sdk.Context, k keeper.Keeper, accs []simtypes.Account, chainID uint64) (simtypes.Account, error) {
	chain, found := k.GetChain(ctx, chainID)
	if !found {
		// No message if no coordinator address
		return simtypes.Account{}, fmt.Errorf("chain %d not found", chainID)
	}
	address, found := k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return simtypes.Account{}, fmt.Errorf("coordinator %d not found", chain.CoordinatorID)
	}
	return FindAccount(accs, address)
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

// FindRequest find a valid request from store
func FindRequest(ctx sdk.Context, k keeper.Keeper) (request types.Request, found bool) {
	// Select a random request without launch triggered
	requests := k.GetAllRequest(ctx)
	for _, req := range requests {
		chain, chainFound := k.GetChain(ctx, req.ChainID)
		if !chainFound || chain.LaunchTriggered {
			continue
		}
		// check if the coordinator is still in the store
		_, coordFound := k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
		if !coordFound {
			continue
		}
		switch content := req.Content.Content.(type) {
		case *types.RequestContent_ValidatorRemoval:
			// if is validator removal, check if the validator exist
			if _, found := k.GetGenesisValidator(
				ctx,
				chain.Id,
				content.ValidatorRemoval.ValAddress,
			); !found {
				continue
			}
		case *types.RequestContent_AccountRemoval:
			// if is account removal, check if account exist
			found, err := keeper.CheckAccount(ctx, k, chain.Id, content.AccountRemoval.Address)
			if err != nil || !found {
				continue
			}
		}
		found = true
		request = req
		break
	}
	return request, found
}

// FindValidator find a valid validator from store
func FindValidator(
	ctx sdk.Context,
	k keeper.Keeper,
	accs []simtypes.Account,
) (simAccount simtypes.Account, valAcc types.GenesisValidator, found bool) {
	valAccs := k.GetAllGenesisValidator(ctx)
	for _, acc := range valAccs {
		if IsLaunchTriggeredChain(ctx, k, acc.ChainID) {
			continue
		}
		// get coordinator account for removal
		var err error
		simAccount, err = FindChainCoordinatorAccount(ctx, k, accs, acc.ChainID)
		if err != nil {
			continue
		}
		valAcc = acc
		found = true
		break
	}
	return simAccount, valAcc, found
}
