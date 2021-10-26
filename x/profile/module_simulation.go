package profile

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/spn/testutil/sample"
	profilesimulation "github.com/tendermint/spn/x/profile/simulation"
	"github.com/tendermint/spn/x/profile/types"
)

const (
	defaultWeightMsgUpdateValidatorDescription   = 50
	defaultWeightMsgDeleteValidator              = 10
	defaultWeightMsgCreateCoordinator            = 50
	defaultWeightMsgUpdateCoordinatorDescription = 20
	defaultWeightMsgUpdateCoordinatorAddress     = 20
	defaultWeightMsgDeleteCoordinator            = 5

	opWeightMsgUpdateValidatorDescription   = "op_weight_msg_update_validator_description"
	opWeightMsgDeleteValidator              = "op_weight_msg_delete_validator"
	opWeightMsgCreateCoordinator            = "op_weight_msg_create_coordinator"
	opWeightMsgUpdateCoordinatorDescription = "op_weight_msg_update_coordinator_description"
	opWeightMsgUpdateCoordinatorAddress     = "op_weight_msg_update_coordinator_address"
	opWeightMsgDeleteCoordinator            = "op_weight_msg_delete_coordinator"
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	profileGenesis := sample.ProfileGenesisState(accs...)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&profileGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	var (
		weightMsgUpdateValidatorDescription   int
		weightMsgDeleteValidator              int
		weightMsgCreateCoordinator            int
		weightMsgUpdateCoordinatorDescription int
		weightMsgUpdateCoordinatorAddress     int
		weightMsgDeleteCoordinator            int
	)

	appParams := simState.AppParams
	cdc := simState.Cdc
	appParams.GetOrGenerate(cdc, opWeightMsgUpdateValidatorDescription, &weightMsgUpdateValidatorDescription, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateValidatorDescription = defaultWeightMsgUpdateValidatorDescription
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgDeleteValidator, &weightMsgDeleteValidator, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteValidator = defaultWeightMsgDeleteValidator
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgCreateCoordinator, &weightMsgCreateCoordinator, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCoordinator = defaultWeightMsgCreateCoordinator
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgUpdateCoordinatorDescription, &weightMsgUpdateCoordinatorDescription, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCoordinatorDescription = defaultWeightMsgUpdateCoordinatorDescription
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgUpdateCoordinatorAddress, &weightMsgUpdateCoordinatorAddress, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCoordinatorAddress = defaultWeightMsgUpdateCoordinatorAddress
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgDeleteCoordinator, &weightMsgDeleteCoordinator, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCoordinator = defaultWeightMsgDeleteCoordinator
		},
	)

	return []simtypes.WeightedOperation{
		simulation.NewWeightedOperation(
			weightMsgUpdateValidatorDescription,
			profilesimulation.SimulateMsgUpdateValidatorDescription(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgDeleteValidator,
			profilesimulation.SimulateMsgDeleteValidator(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateCoordinator,
			profilesimulation.SimulateMsgCreateCoordinator(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateCoordinatorDescription,
			profilesimulation.SimulateMsgUpdateCoordinatorDescription(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateCoordinatorAddress,
			profilesimulation.SimulateMsgUpdateCoordinatorAddress(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgDeleteCoordinator,
			profilesimulation.SimulateMsgDeleteCoordinator(am.accountKeeper, am.bankKeeper, am.keeper),
		),
	}
}
