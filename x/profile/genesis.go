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
	// Set all the consensusKeyNonce
	for _, elem := range genState.ConsensusKeyNonceList {
		k.SetConsensusKeyNonce(ctx, *elem)
	}

	// Set all the validatorByConsAddress
	for _, elem := range genState.ValidatorByConsAddressList {
		k.SetValidatorByConsAddress(ctx, *elem)
	}

	// Set all the validatorByAddress
	for _, elem := range genState.ValidatorByAddressList {
		k.SetValidatorByAddress(ctx, *elem)
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
	// Get all consensusKeyNonce
	consensusKeyNonceList := k.GetAllConsensusKeyNonce(ctx)
	for _, elem := range consensusKeyNonceList {
		elem := elem
		genesis.ConsensusKeyNonceList = append(genesis.ConsensusKeyNonceList, &elem)
	}

	// Get all validatorByConsAddress
	validatorByConsAddressList := k.GetAllValidatorByConsAddress(ctx)
	for _, elem := range validatorByConsAddressList {
		elem := elem
		genesis.ValidatorByConsAddressList = append(genesis.ValidatorByConsAddressList, &elem)
	}

	// Get all validatorByAddress
	validatorByAddressList := k.GetAllValidatorByAddress(ctx)
	for _, elem := range validatorByAddressList {
		elem := elem
		genesis.ValidatorByAddressList = append(genesis.ValidatorByAddressList, &elem)
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
