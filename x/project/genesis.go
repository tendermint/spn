package project

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/project/keeper"
	"github.com/tendermint/spn/x/project/types"
)

// InitGenesis initializes the project module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the project
	for _, elem := range genState.Projects {
		k.SetProject(ctx, elem)
	}
	// Set project counter
	k.SetProjectCounter(ctx, genState.ProjectCounter)

	// Set all the projectChains
	for _, elem := range genState.ProjectChains {
		k.SetProjectChains(ctx, elem)
	}

	// Set all the mainnetAccount
	for _, elem := range genState.MainnetAccounts {
		k.SetMainnetAccount(ctx, elem)
	}

	k.SetParams(ctx, genState.Params)

	// set maximum shares constant value
	k.SetTotalShares(ctx, genState.TotalShares)

	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the project module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Projects = k.GetAllProject(ctx)
	genesis.ProjectCounter = k.GetProjectCounter(ctx)
	genesis.ProjectChains = k.GetAllProjectChains(ctx)
	genesis.MainnetAccounts = k.GetAllMainnetAccount(ctx)
	genesis.Params = k.GetParams(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
