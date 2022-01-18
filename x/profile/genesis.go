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

	// Set all the validatorByConsAddress
	for _, elem := range genState.ValidatorByConsAddressList {
		k.SetValidatorByConsAddress(ctx, elem)
	}

	// Set all the consensusKeyNonce
	for _, elem := range genState.ConsensusKeyNonceList {
		k.SetConsensusKeyNonce(ctx, elem)
	}

	// Set all the coordinator
	for _, elem := range genState.CoordinatorList {
		k.SetCoordinator(ctx, elem)
	}

	// Set coordinator counter
	k.SetCoordinatorCounter(ctx, genState.CoordinatorCounter)

	// Set all the coordinatorByAddress
	for _, elem := range genState.CoordinatorByAddressList {
		k.SetCoordinatorByAddress(ctx, elem)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.CoordinatorList = k.GetAllCoordinator(ctx)
	genesis.CoordinatorCounter = k.GetCoordinatorCounter(ctx)
	genesis.CoordinatorByAddressList = k.GetAllCoordinatorByAddress(ctx)
	genesis.ValidatorList = k.GetAllValidator(ctx)
	genesis.ValidatorByConsAddressList = k.GetAllValidatorByConsAddress(ctx)
	genesis.ConsensusKeyNonceList = k.GetAllConsensusKeyNonce(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
