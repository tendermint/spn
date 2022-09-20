package launch

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

// InitGenesis initializes the launch module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the chain
	for _, elem := range genState.Chains {
		k.SetChain(ctx, elem)
	}

	k.SetChainCounter(ctx, genState.ChainCounter)

	// Set all the genesisAccount
	for _, elem := range genState.GenesisAccounts {
		k.SetGenesisAccount(ctx, elem)
	}

	// Set all the vestingAccount
	for _, elem := range genState.VestingAccounts {
		k.SetVestingAccount(ctx, elem)
	}

	// Set all the genesisValidator
	for _, elem := range genState.GenesisValidators {
		k.SetGenesisValidator(ctx, elem)
	}

	// Set all the request
	for _, elem := range genState.Requests {
		k.SetRequest(ctx, elem)
	}

	// Set all request counter
	for _, elem := range genState.RequestCounters {
		k.SetRequestCounter(ctx, elem.LaunchID, elem.Counter)
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the launch module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Chains = k.GetAllChain(ctx)
	genesis.ChainCounter = k.GetChainCounter(ctx)
	genesis.GenesisAccounts = k.GetAllGenesisAccount(ctx)
	genesis.VestingAccounts = k.GetAllVestingAccount(ctx)
	genesis.GenesisValidators = k.GetAllGenesisValidator(ctx)
	genesis.Requests = k.GetAllRequest(ctx)
	genesis.Params = k.GetParams(ctx)

	// Get request counts
	for _, elem := range genesis.Chains {
		// Get request count
		counter := k.GetRequestCounter(ctx, elem.LaunchID)
		genesis.RequestCounters = append(genesis.RequestCounters, types.RequestCounter{
			LaunchID: elem.LaunchID,
			Counter:  counter,
		})
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
