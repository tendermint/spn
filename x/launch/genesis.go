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
		k.SetChain(ctx, *elem)
	}

	// Set all the genesisAccount
	for _, elem := range genState.GenesisAccountList {
		k.SetGenesisAccount(ctx, *elem)
	}

	// Set all the vestedAccount
	for _, elem := range genState.VestedAccountList {
		k.SetVestedAccount(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export

	// Get all chain
	chainList := k.GetAllChain(ctx)
	for _, elem := range k.GetAllChain(ctx) {
		elem := elem
		genesis.ChainList = append(genesis.ChainList, &elem)
	}

	// Get all genesisAccount
	genesisAccountList := k.GetAllGenesisAccount(ctx)
	for _, elem := range genesisAccountList {
		elem := elem
		genesis.GenesisAccountList = append(genesis.GenesisAccountList, &elem)
	}

	// Get all vestedAccount
	vestedAccountList := k.GetAllVestedAccount(ctx)
	for _, elem := range vestedAccountList {
		elem := elem
		genesis.VestedAccountList = append(genesis.VestedAccountList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
