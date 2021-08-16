package profile

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the validator
	for _, elem := range genState.ValidatorList {
		k.SetValidator(ctx, elem)
	}

	// Set all the coordinatorByAddress
	for _, elem := range genState.CoordinatorByAddressList {
		k.SetCoordinatorByAddress(ctx, elem)
	}

	// Set all the coordinator
	for _, elem := range genState.CoordinatorList {
		k.SetCoordinator(ctx, elem)
	}

	// Set coordinator count
	k.SetCoordinatorCount(ctx, genState.CoordinatorCount)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.ValidatorList = k.GetAllValidator(ctx)
	genesis.CoordinatorByAddressList = k.GetAllCoordinatorByAddress(ctx)
	genesis.CoordinatorList = k.GetAllCoordinator(ctx)
	genesis.CoordinatorCount = k.GetCoordinatorCount(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
