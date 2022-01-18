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

	k.SetChainCounter(ctx, genState.ChainCounter)

	// Set all the genesisAccount
	for _, elem := range genState.GenesisAccountList {
		k.SetGenesisAccount(ctx, elem)
	}

	// Set all the vestingAccount
	for _, elem := range genState.VestingAccountList {
		k.SetVestingAccount(ctx, elem)
	}

	// Set all the genesisValidator
	for _, elem := range genState.GenesisValidatorList {
		k.SetGenesisValidator(ctx, elem)
	}

	// Set all the request
	for _, elem := range genState.RequestList {
		k.SetRequest(ctx, elem)
	}

	// Set all request counter
	for _, elem := range genState.RequestCounterList {
		k.SetRequestCounter(ctx, elem.LaunchID, elem.Counter)
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.ChainList = k.GetAllChain(ctx)
	genesis.ChainCounter = k.GetChainCounter(ctx)
	genesis.GenesisAccountList = k.GetAllGenesisAccount(ctx)
	genesis.VestingAccountList = k.GetAllVestingAccount(ctx)
	genesis.GenesisValidatorList = k.GetAllGenesisValidator(ctx)
	genesis.RequestList = k.GetAllRequest(ctx)
	genesis.Params = k.GetParams(ctx)

	// Get request counts
	for _, elem := range genesis.ChainList {
		// Get request count
		counter := k.GetRequestCounter(ctx, elem.LaunchID)
		genesis.RequestCounterList = append(genesis.RequestCounterList, types.RequestCounter{
			LaunchID: elem.LaunchID,
			Counter:  counter,
		})
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
