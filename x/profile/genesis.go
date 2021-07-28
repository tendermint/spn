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
		k.SetValidator(ctx, *elem)
	}

	// Set all the coordinatorByAddress
	for _, elem := range genState.CoordinatorByAddressList {
		k.SetCoordinatorByAddress(ctx, *elem)
	}

	// Set all the coordinator
	for _, elem := range genState.CoordinatorList {
		k.SetCoordinator(ctx, *elem)
	}

	// Set coordinator count
	k.SetCoordinatorCount(ctx, genState.CoordinatorCount)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all validator
	validatorList := k.GetAllValidator(ctx)
	for _, elem := range validatorList {
		elem := elem
		genesis.ValidatorList = append(genesis.ValidatorList, &elem)
	}

	// Get all coordinatorByAddress
	coordinatorByAddressList := k.GetAllCoordinatorByAddress(ctx)
	for _, elem := range coordinatorByAddressList {
		elem := elem
		genesis.CoordinatorByAddressList = append(genesis.CoordinatorByAddressList, &elem)
	}

	// Get all coordinator
	coordinatorList := k.GetAllCoordinator(ctx)
	for _, elem := range coordinatorList {
		elem := elem
		genesis.CoordinatorList = append(genesis.CoordinatorList, &elem)
	}

	// Set the current count
	genesis.CoordinatorCount = k.GetCoordinatorCount(ctx)

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
