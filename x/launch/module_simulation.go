package launch

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/sample"
	launchsim "github.com/tendermint/spn/x/launch/simulation"
	"github.com/tendermint/spn/x/launch/types"
)

const (
	defaultWeightMsgCreateChain              int = 50
	defaultWeightMsgEditChain                int = 20
	defaultWeightMsgRequestAddGenesisAccount int = 50
	defaultWeightMsgRequestAddVestingAccount int = 50
	defaultWeightMsgRequestRemoveAccount     int = 15
	defaultWeightMsgRequestAddValidator      int = 50
	defaultWeightMsgRequestRemoveValidator   int = 15
	defaultWeightMsgRequestParamChange       int = 15
	defaultWeightMsgSettleRequest            int = 50
	defaultWeightMsgTriggerLaunch            int = 15
	defaultWeightMsgRevertLaunch             int = 0
	defaultWeightMsgUpdateLaunchInformation  int = 20

	opWeightMsgCreateChain              = "op_weight_msg_create_chain"
	opWeightMsgEditChain                = "op_weight_msg_edit_chain"
	opWeightMsgRequestAddGenesisAccount = "op_weight_msg_request_add_genesis_account"
	opWeightMsgRequestAddVestingAccount = "op_weight_msg_request_add_vesting_account"
	opWeightMsgRequestRemoveAccount     = "op_weight_msg_request_remove_account"
	opWeightMsgRequestAddValidator      = "op_weight_msg_request_add_validator"
	opWeightMsgRequestRemoveValidator   = "op_weight_msg_request_remove_validator"
	opWeightMsgRequestParamChange       = "op_weight_msg_request_change_param"
	opWeightMsgTriggerLaunch            = "op_weight_msg_trigger_launch"
	opWeightMsgRevertLaunch             = "op_weight_msg_revert_launch"
	opWeightMsgSettleRequest            = "op_weight_msg_settle_request"
	opWeightMsgUpdateLaunchInformation  = "op_weight_msg_update_launch_information"

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	launchGenesis := sample.LaunchGenesisState(simState.Rand, accs...)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&launchGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	launchParams := sample.LaunchParams(r)
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyLaunchTimeRange), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(launchParams.LaunchTimeRange))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyChainCreationFee), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(launchParams.ChainCreationFee))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	// this line is used by starport scaffolding # simapp/module/operation
	var (
		weightMsgCreateChain              int
		weightMsgEditChain                int
		weightMsgRequestAddGenesisAccount int
		weightMsgRequestAddVestingAccount int
		weightMsgRequestRemoveAccount     int
		weightMsgRequestAddValidator      int
		weightMsgRequestRemoveValidator   int
		weightMsgRequestParamChange       int
		weightMsgTriggerLaunch            int
		weightMsgRevertLaunch             int
		weightMsgSettleRequest            int
		weightMsgUpdateLaunchInformation  int
	)

	appParams := simState.AppParams
	cdc := simState.Cdc
	appParams.GetOrGenerate(cdc, opWeightMsgCreateChain, &weightMsgCreateChain, nil,
		func(_ *rand.Rand) {
			weightMsgCreateChain = defaultWeightMsgCreateChain
		},
	)
	appParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateLaunchInformation, &weightMsgUpdateLaunchInformation, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLaunchInformation = defaultWeightMsgUpdateLaunchInformation
		},
	)

	appParams.GetOrGenerate(cdc, opWeightMsgEditChain, &weightMsgEditChain, nil,
		func(_ *rand.Rand) {
			weightMsgEditChain = defaultWeightMsgEditChain
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestAddGenesisAccount, &weightMsgRequestAddGenesisAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddGenesisAccount = defaultWeightMsgRequestAddGenesisAccount
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestAddGenesisAccount, &weightMsgRequestAddGenesisAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddGenesisAccount = defaultWeightMsgRequestAddGenesisAccount
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestAddVestingAccount, &weightMsgRequestAddVestingAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddVestingAccount = defaultWeightMsgRequestAddVestingAccount
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestRemoveAccount, &weightMsgRequestRemoveAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRemoveAccount = defaultWeightMsgRequestRemoveAccount
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestAddValidator, &weightMsgRequestAddValidator, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddValidator = defaultWeightMsgRequestAddValidator
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestRemoveValidator, &weightMsgRequestRemoveValidator, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRemoveValidator = defaultWeightMsgRequestRemoveValidator
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestParamChange, &weightMsgRequestParamChange, nil,
		func(_ *rand.Rand) {
			weightMsgRequestParamChange = defaultWeightMsgRequestParamChange
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgTriggerLaunch, &weightMsgTriggerLaunch, nil,
		func(_ *rand.Rand) {
			weightMsgTriggerLaunch = defaultWeightMsgTriggerLaunch
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRevertLaunch, &weightMsgRevertLaunch, nil,
		func(_ *rand.Rand) {
			weightMsgRevertLaunch = defaultWeightMsgRevertLaunch
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgSettleRequest, &weightMsgSettleRequest, nil,
		func(_ *rand.Rand) {
			weightMsgSettleRequest = defaultWeightMsgSettleRequest
		},
	)

	return []simtypes.WeightedOperation{
		simulation.NewWeightedOperation(
			weightMsgCreateChain,
			launchsim.SimulateMsgCreateChain(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgEditChain,
			launchsim.SimulateMsgEditChain(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestAddGenesisAccount,
			launchsim.SimulateMsgRequestAddGenesisAccount(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestAddVestingAccount,
			launchsim.SimulateMsgRequestAddVestingAccount(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestRemoveAccount,
			launchsim.SimulateMsgRequestRemoveAccount(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestAddValidator,
			launchsim.SimulateMsgRequestAddValidator(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestRemoveValidator,
			launchsim.SimulateMsgRequestRemoveValidator(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestParamChange,
			launchsim.SimulateMsgRequestParamChange(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgTriggerLaunch,
			launchsim.SimulateMsgTriggerLaunch(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgSettleRequest,
			launchsim.SimulateMsgSettleRequest(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRevertLaunch,
			launchsim.SimulateMsgRevertLaunch(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateLaunchInformation,
			launchsim.SimulateMsgUpdateLaunchInformation(am.accountKeeper, am.bankKeeper, am.keeper),
		),
	}
}
