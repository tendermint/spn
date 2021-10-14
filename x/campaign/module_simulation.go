package campaign

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	campaignsim "github.com/tendermint/spn/x/campaign/simulation"
	"github.com/tendermint/spn/x/campaign/types"
)

const (
	weightMsgCreateCampaign    = 25
	weightMsgUpdateTotalSupply = 20
	weightMsgUpdateTotalShares = 20
	weightMsgInitializeMainnet = 5
	weightMsgAddShares         = 20
	weightMsgAddVestingOptions = 20
	weightMsgMintVouchers      = 20
	weightMsgBurnVouchers      = 20
	weightMsgRedeemVouchers    = 20
	weightMsgUnredeemVouchers  = 20
	weightMsgSendVouchers      = 20
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(types.DefaultGenesis())
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
			weightMsgCreateCampaign,
			campaignsim.SimulateMsgCreateCampaign(am.accountKeeper, am.bankKeeper, am.profileKeeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateTotalSupply,
			campaignsim.SimulateMsgUpdateTotalSupply(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateTotalShares,
			campaignsim.SimulateMsgUpdateTotalShares(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgInitializeMainnet,
			campaignsim.SimulateMsgInitializeMainnet(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgAddShares,
			campaignsim.SimulateMsgAddShares(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgAddVestingOptions,
			campaignsim.SimulateMsgAddVestingOptions(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgMintVouchers,
			campaignsim.SimulateMsgMintVouchers(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgBurnVouchers,
			campaignsim.SimulateMsgBurnVouchers(am.accountKeeper, am.bankKeeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRedeemVouchers,
			campaignsim.SimulateMsgRedeemVouchers(am.accountKeeper, am.bankKeeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUnredeemVouchers,
			campaignsim.SimulateMsgUnredeemVouchers(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgSendVouchers,
			campaignsim.SimulateMsgSendVouchers(am.accountKeeper, am.bankKeeper),
		),
	}
}
