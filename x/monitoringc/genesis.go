package monitoringc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// InitGenesis initializes the monitoringc module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the verifiedClientID
	for _, elem := range genState.VerifiedClientIDList {
		k.SetVerifiedClientID(ctx, elem)
	}
	// Set all the providerClientID
	for _, elem := range genState.ProviderClientIDList {
		k.SetProviderClientID(ctx, elem)
	}
	// Set all the launchIDFromVerifiedClientID
	for _, elem := range genState.LaunchIDFromVerifiedClientIDList {
		k.SetLaunchIDFromVerifiedClientID(ctx, elem)
	}
	// Set all the launchIDFromChannelID
	for _, elem := range genState.LaunchIDFromChannelIDList {
		k.SetLaunchIDFromChannelID(ctx, elem)
	}
	// Set all the monitoringHistory
	for _, elem := range genState.MonitoringHistoryList {
		k.SetMonitoringHistory(ctx, elem)
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
}

// ExportGenesis returns the monitoringc module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.VerifiedClientIDList = k.GetAllVerifiedClientID(ctx)
	genesis.ProviderClientIDList = k.GetAllProviderClientID(ctx)
	genesis.LaunchIDFromVerifiedClientIDList = k.GetAllLaunchIDFromVerifiedClientID(ctx)
	genesis.LaunchIDFromChannelIDList = k.GetAllLaunchIDFromChannelID(ctx)
	genesis.MonitoringHistoryList = k.GetAllMonitoringHistory(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
