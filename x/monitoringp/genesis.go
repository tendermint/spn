package monitoringp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/monitoringp/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.ConsumerClientID != nil {
		k.SetConsumerClientID(ctx, *genState.ConsumerClientID)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)

	// initialize and setup the consumer IBC client
	if genState.Params.ConsumerConsensusState.Timestamp != "" {
		_, err := k.InitializeConsumerClient(ctx)
		if err != nil {
			panic("couldn't initialize the consumer client ID" + err.Error())
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	// Get all consumerClientID
	consumerClientID, found := k.GetConsumerClientID(ctx)
	if found {
		genesis.ConsumerClientID = &consumerClientID
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
