package monitoringp

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringp/types"
)

const (
// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	monitoringpGenesis := types.GenesisState{
		PortId: types.PortID,
		Params: sample.MonitoringpParams(simState.Rand),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&monitoringpGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.LegacyParamChange {
	monitoringpParams := sample.MonitoringpParams(r)
	return []simtypes.LegacyParamChange{
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyLastBlockHeight), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(monitoringpParams.LastBlockHeight))
		}),
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyConsumerConsensusState), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(monitoringpParams.ConsumerConsensusState))
		}),
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyConsumerChainID), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(monitoringpParams.ConsumerChainID))
		}),
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyConsumerUnbondingPeriod), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(monitoringpParams.ConsumerUnbondingPeriod))
		}),
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyConsumerRevisionHeight), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(monitoringpParams.ConsumerRevisionHeight))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
