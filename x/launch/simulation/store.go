package simulation

import (
	"fmt"
	"math/rand"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"

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

func checkRequest(
	ctx sdk.Context,
	k keeper.Keeper,
	request types.Request,
) error {
	if err := request.Content.Validate(); err != nil {
		return err
	}

	switch requestContent := request.Content.Content.(type) {
	case *types.RequestContent_GenesisAccount:
		ga := requestContent.GenesisAccount
		found, err := keeper.CheckAccount(ctx, k, request.LaunchID, ga.Address)
		if err != nil {
			return err
		}
		if found {
			return sdkerrors.Wrapf(types.ErrAccountAlreadyExist,
				"account %s for chain %d already exist",
				ga.Address, request.LaunchID,
			)
		}

	case *types.RequestContent_VestingAccount:
		va := requestContent.VestingAccount
		found, err := keeper.CheckAccount(ctx, k, request.LaunchID, va.Address)
		if err != nil {
			return err
		}
		if found {
			return sdkerrors.Wrapf(types.ErrAccountAlreadyExist,
				"account %s for chain %d already exist",
				va.Address, request.LaunchID,
			)
		}

	case *types.RequestContent_AccountRemoval:
		ar := requestContent.AccountRemoval
		found, err := keeper.CheckAccount(ctx, k, request.LaunchID, ar.Address)
		if err != nil {
			return err
		}
		if !found {
			return sdkerrors.Wrapf(types.ErrAccountNotFound,
				"account %s for chain %d not found",
				ar.Address, request.LaunchID,
			)
		}

	case *types.RequestContent_GenesisValidator:
		ga := requestContent.GenesisValidator
		if _, found := k.GetGenesisValidator(ctx, request.LaunchID, ga.Address); found {
			return sdkerrors.Wrapf(types.ErrValidatorAlreadyExist,
				"genesis validator %s for chain %d already exist",
				ga.Address, request.LaunchID,
			)
		}

	case *types.RequestContent_ValidatorRemoval:
		vr := requestContent.ValidatorRemoval
		if _, found := k.GetGenesisValidator(ctx, request.LaunchID, vr.ValAddress); !found {
			return sdkerrors.Wrapf(types.ErrValidatorNotFound,
				"genesis validator %s for chain %d not found",
				vr.ValAddress, request.LaunchID,
			)
		}

	default:
		return spnerrors.Critical("unknown request content type")
	}
	return nil
}
