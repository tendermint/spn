package launch

import (
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	launchGenesis := sample.LaunchGenesisState()
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&launchGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	launchParams := sample.LaunchParams()

	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMinLaunchTime), func(r *rand.Rand) string {
			return fmt.Sprintf("\"%d\"", launchParams.MinLaunchTime)
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMaxLaunchTime), func(r *rand.Rand) string {
			return fmt.Sprintf("\"%d\"", launchParams.MaxLaunchTime)

		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return []simtypes.WeightedOperation{}
}
