package profile

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

// InitGenesis initializes the profile module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init

	// Set all the validator
	for _, elem := range genState.Validators {
		k.SetValidator(ctx, elem)
	}

	// Set all the validatorByOperatorAddress
	for _, elem := range genState.ValidatorsByOperatorAddress {
		k.SetValidatorByOperatorAddress(ctx, elem)
	}

	// Set all the coordinator
	for _, elem := range genState.Coordinators {
		k.SetCoordinator(ctx, elem)
	}

	// Set coordinator counter
	k.SetCoordinatorCounter(ctx, genState.CoordinatorCounter)

	// Set all the coordinatorByAddress
	for _, elem := range genState.CoordinatorsByAddress {
		k.SetCoordinatorByAddress(ctx, elem)
	}
}

// ExportGenesis returns the profile module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Coordinators = k.GetAllCoordinator(ctx)
	genesis.CoordinatorCounter = k.GetCoordinatorCounter(ctx)
	genesis.CoordinatorsByAddress = k.GetAllCoordinatorByAddress(ctx)
	genesis.Validators = k.GetAllValidator(ctx)
	genesis.ValidatorsByOperatorAddress = k.GetAllValidatorByOperatorAddress(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
