package participation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/sample"
	participationsim "github.com/tendermint/spn/x/participation/simulation"
	"github.com/tendermint/spn/x/participation/types"
)

const (
	opWeightMsgParticipate = "op_weight_msg_participate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgParticipate int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	participationGenesis := types.GenesisState{
		Params: sample.ParticipationParams(),
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
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyRegistrationPeriod), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(participationParams.RegistrationPeriod))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyWithdrawalDelay), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(participationParams.WithdrawalDelay))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgParticipate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgParticipate, &weightMsgParticipate, nil,
		func(_ *rand.Rand) {
			weightMsgParticipate = defaultWeightMsgParticipate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgParticipate,
		participationsim.SimulateMsgParticipate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
