package project

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/sample"
	projectsim "github.com/tendermint/spn/x/project/simulation"
	"github.com/tendermint/spn/x/project/types"
)

const (
	defaultWeightMsgCreateProject            = 25
	defaultWeightMsgEditProject              = 20
	defaultWeightMsgUpdateTotalSupply        = 20
	defaultWeightMsgInitializeMainnet        = 15
	defaultWeightMsgUpdateSpecialAllocations = 20
	defaultWeightMsgMintVouchers             = 20
	defaultWeightMsgBurnVouchers             = 20
	defaultWeightMsgRedeemVouchers           = 20
	defaultWeightMsgUnredeemVouchers         = 20

	opWeightMsgCreateProject            = "op_weight_msg_create_project"
	opWeightMsgEditProject              = "op_weight_msg_edit_project"
	opWeightMsgUpdateTotalSupply        = "op_weight_msg_update_total_supply"
	opWeightMsgInitializeMainnet        = "op_weight_msg_initialize_mainnet"
	opWeightMsgUpdateSpecialAllocations = "op_weight_msg_update_special_allocations"
	opWeightMsgMintVouchers             = "op_weight_msg_mint_vouchers"
	opWeightMsgBurnVouchers             = "op_weight_msg_burn_vouchers"
	opWeightMsgRedeemVouchers           = "op_weight_msg_redeem_vouchers"
	opWeightMsgUnredeemVouchers         = "op_weight_msg_unredeem_vouchers"

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	projectGenesis := sample.ProjectGenesisState(simState.Rand)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&projectGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.LegacyParamChange {
	projectParams := types.DefaultParams()
	return []simtypes.LegacyParamChange{
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyTotalSupplyRange), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(projectParams.TotalSupplyRange))
		}),
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyProjectCreationFee), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(projectParams.ProjectCreationFee))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	var (
		weightMsgCreateProject            int
		weightMsgEditProject              int
		weightMsgUpdateTotalSupply        int
		weightMsgInitializeMainnet        int
		weightMsgUpdateSpecialAllocations int
		weightMsgMintVouchers             int
		weightMsgBurnVouchers             int
		weightMsgRedeemVouchers           int
		weightMsgUnredeemVouchers         int
	)

	appParams := simState.AppParams
	cdc := simState.Cdc
	appParams.GetOrGenerate(cdc, opWeightMsgCreateProject, &weightMsgCreateProject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProject = defaultWeightMsgCreateProject
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgEditProject, &weightMsgEditProject, nil,
		func(_ *rand.Rand) {
			weightMsgEditProject = defaultWeightMsgEditProject
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgUpdateTotalSupply, &weightMsgUpdateTotalSupply, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTotalSupply = defaultWeightMsgUpdateTotalSupply
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgInitializeMainnet, &weightMsgInitializeMainnet, nil,
		func(_ *rand.Rand) {
			weightMsgInitializeMainnet = defaultWeightMsgInitializeMainnet
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgUpdateSpecialAllocations, &weightMsgUpdateSpecialAllocations, nil,
		func(_ *rand.Rand) {
			weightMsgInitializeMainnet = defaultWeightMsgUpdateSpecialAllocations
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgMintVouchers, &weightMsgMintVouchers, nil,
		func(_ *rand.Rand) {
			weightMsgMintVouchers = defaultWeightMsgMintVouchers
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgBurnVouchers, &weightMsgBurnVouchers, nil,
		func(_ *rand.Rand) {
			weightMsgBurnVouchers = defaultWeightMsgBurnVouchers
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRedeemVouchers, &weightMsgRedeemVouchers, nil,
		func(_ *rand.Rand) {
			weightMsgRedeemVouchers = defaultWeightMsgRedeemVouchers
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgUnredeemVouchers, &weightMsgUnredeemVouchers, nil,
		func(_ *rand.Rand) {
			weightMsgUnredeemVouchers = defaultWeightMsgUnredeemVouchers
		},
	)

	return []simtypes.WeightedOperation{
		simulation.NewWeightedOperation(
			weightMsgCreateProject,
			projectsim.SimulateMsgCreateProject(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgEditProject,
			projectsim.SimulateMsgEditProject(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateTotalSupply,
			projectsim.SimulateMsgUpdateTotalSupply(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgInitializeMainnet,
			projectsim.SimulateMsgInitializeMainnet(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateSpecialAllocations,
			projectsim.SimulateMsgUpdateSpecialAllocations(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgMintVouchers,
			projectsim.SimulateMsgMintVouchers(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgBurnVouchers,
			projectsim.SimulateMsgBurnVouchers(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRedeemVouchers,
			projectsim.SimulateMsgRedeemVouchers(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUnredeemVouchers,
			projectsim.SimulateMsgUnredeemVouchers(am.accountKeeper, am.bankKeeper, am.keeper),
		),
	}
}
