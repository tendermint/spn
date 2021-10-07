package profile

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

const (
	weightMsgUpdateValidatorDescription   = 10
	weightMsgDeleteValidator              = 10
	weightMsgCreateCoordinator            = 10
	weightMsgUpdateCoordinatorDescription = 10
	weightMsgUpdateCoordinatorAddress     = 10
	weightMsgDeleteCoordinator            = 10
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	profileGenesis := sample.ProfileGenesisState()
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&profileGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return []simtypes.WeightedOperation{
		simulation.NewWeightedOperation(
			weightMsgCreateCoordinator,
			SimulateMsgCreateCoordinator(),
		),
	}
}

// SimulateMsgCreateCoordinator returns a MsgCreateCoordinator operation for simulation
func SimulateMsgCreateCoordinator() simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// If no op
		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateCoordinator, "skip coordinator creation"), nil, nil

		//return simtypes.NewOperationMsg(nil, true, "", nil), nil, nil
	}
}
