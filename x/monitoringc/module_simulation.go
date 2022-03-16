package monitoringc

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/sample"
	monitoringcsimulation "github.com/tendermint/spn/x/monitoringc/simulation"
	"github.com/tendermint/spn/x/monitoringc/types"
)

const (
	opWeightMsgCreateClient          = "op_weight_msg_create_chain"
	defaultWeightMsgCreateClient int = 50

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	monitoringcGenesis := types.GenesisState{
		PortId: types.PortID,
		Params: sample.MonitoringcParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&monitoringcGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	monitoringcParams := sample.MonitoringcParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyDebugMode), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(monitoringcParams.DebugMode))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateClient int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateClient, &weightMsgCreateClient, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClient = defaultWeightMsgCreateClient
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateClient,
		monitoringcsimulation.SimulateMsgCreateClient(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
