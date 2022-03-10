package participation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/x/participation/types"
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
	participationGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&participationGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	participationParams := sample.ParticipationParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyAllocationPrice), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(participationParams.AllocationPrice))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyParticipationTierList), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(participationParams.ParticipationTierList))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyWithdrawalAllocationDelay), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(participationParams.WithdrawalAllocationDelay))
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
