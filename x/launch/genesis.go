package launch

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the chain
	for _, elem := range genState.ChainList {
		k.SetChain(ctx, elem)
	}

	// Set all the chainNameCount
	for _, elem := range genState.ChainNameCountList {
		k.SetChainNameCount(ctx, elem)
	}

	// Set all the genesisAccount
	for _, elem := range genState.GenesisAccountList {
		k.SetGenesisAccount(ctx, elem)
	}

	// Set all the vestedAccount
	for _, elem := range genState.VestedAccountList {
		k.SetVestedAccount(ctx, elem)
	}

	// Set all the genesisValidator
	for _, elem := range genState.GenesisValidatorList {
		k.SetGenesisValidator(ctx, elem)
	}

	// Set all the request
	for _, elem := range genState.RequestList {
		k.SetRequest(ctx, elem)
	}

	// Set all request count
	for _, elem := range genState.RequestCountList {
		k.SetRequestCount(ctx, elem.ChainID, elem.Count)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.ChainList = k.GetAllChain(ctx)
	genesis.ChainNameCountList = k.GetAllChainNameCount(ctx)
	genesis.GenesisAccountList = k.GetAllGenesisAccount(ctx)
	genesis.VestedAccountList = k.GetAllVestedAccount(ctx)
	genesis.GenesisValidatorList = k.GetAllGenesisValidator(ctx)
	genesis.RequestList = k.GetAllRequest(ctx)
	genesis.Params = k.GetParams(ctx)

	// Get request counts
	for _, elem := range genesis.ChainList {
		// Get request count
		count := k.GetRequestCount(ctx, elem.ChainID)
		genesis.RequestCountList = append(genesis.RequestCountList, types.RequestCount{
			ChainID: elem.ChainID,
			Count:   count,
		})
	}

	// this line is used by starport scaffolding # genesis/module/export
	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
