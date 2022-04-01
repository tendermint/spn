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
	opWeightMsgParticipate            = "op_weight_msg_participate"
	defaultWeightMsgParticipate   int = 50
	opWeightMsgCreateAuction          = "op_weight_create_auction"
	defaultWeightMsgCreateAuction int = 20

	opWeightMsgWithdrawAllocations          = "op_weight_withdraw_allocations"
	defaultWeightMsgWithdrawAllocations int = 50

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	participationGenesis := types.GenesisState{
		Params: sample.ParticipationParams(simState.Rand),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&participationGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	participationParams := sample.ParticipationParams(r)
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

	var (
		weightMsgParticipate         int
		weightMsgCreateAuction       int
		weightMsgWithdrawAllocations int
	)

	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgParticipate, &weightMsgParticipate, nil,
		func(_ *rand.Rand) {
			weightMsgParticipate = defaultWeightMsgParticipate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgParticipate,
		participationsim.SimulateMsgParticipate(am.accountKeeper, am.bankKeeper, am.fundraisingKeeper, am.keeper),
	))

	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAuction, &weightMsgCreateAuction, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAuction = defaultWeightMsgCreateAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAuction,
		participationsim.SimulateCreateAuction(am.accountKeeper, am.bankKeeper, am.fundraisingKeeper, am.keeper),
	))

	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawAllocations, &weightMsgWithdrawAllocations, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawAllocations = defaultWeightMsgWithdrawAllocations
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawAllocations,
		participationsim.SimulateMsgWithdrawAllocations(am.accountKeeper, am.bankKeeper, am.FundraisingKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
