package campaign

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/sample"
	campaignsim "github.com/tendermint/spn/x/campaign/simulation"
	"github.com/tendermint/spn/x/campaign/types"
)

const (
	defaultWeightMsgCreateCampaign           = 25
	defaultWeightMsgEditCampaign             = 20
	defaultWeightMsgUpdateTotalSupply        = 20
	defaultWeightMsgInitializeMainnet        = 15
	defaultWeightMsgUpdateSpecialAllocations = 20
	defaultWeightMsgAddShares                = 20
	defaultWeightMsgMintVouchers             = 20
	defaultWeightMsgBurnVouchers             = 20
	defaultWeightMsgRedeemVouchers           = 20
	defaultWeightMsgUnredeemVouchers         = 20

	opWeightMsgCreateCampaign           = "op_weight_msg_create_campaign"
	opWeightMsgEditCampaign             = "op_weight_msg_edit_campaign"
	opWeightMsgUpdateTotalSupply        = "op_weight_msg_update_total_supply"
	opWeightMsgInitializeMainnet        = "op_weight_msg_initialize_mainnet"
	opWeightMsgUpdateSpecialAllocations = "op_weight_msg_update_special_allocations"
	opWeightMsgAddShares                = "op_weight_msg_add_shares"
	opWeightMsgMintVouchers             = "op_weight_msg_mint_vouchers"
	opWeightMsgBurnVouchers             = "op_weight_msg_burn_vouchers"
	opWeightMsgRedeemVouchers           = "op_weight_msg_redeem_vouchers"
	opWeightMsgUnredeemVouchers         = "op_weight_msg_unredeem_vouchers"

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	campaignGenesis := sample.CampaignGenesisState(simState.Rand)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&campaignGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	campaignParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyTotalSupplyRange), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(campaignParams.TotalSupplyRange))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyCampaignCreationFee), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(campaignParams.CampaignCreationFee))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	var (
		weightMsgCreateCampaign           int
		weightMsgEditCampaign             int
		weightMsgUpdateTotalSupply        int
		weightMsgInitializeMainnet        int
		weightMsgUpdateSpecialAllocations int
		weightMsgAddShares                int
		weightMsgMintVouchers             int
		weightMsgBurnVouchers             int
		weightMsgRedeemVouchers           int
		weightMsgUnredeemVouchers         int
	)

	appParams := simState.AppParams
	cdc := simState.Cdc
	appParams.GetOrGenerate(cdc, opWeightMsgCreateCampaign, &weightMsgCreateCampaign, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCampaign = defaultWeightMsgCreateCampaign
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgEditCampaign, &weightMsgEditCampaign, nil,
		func(_ *rand.Rand) {
			weightMsgEditCampaign = defaultWeightMsgEditCampaign
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
	appParams.GetOrGenerate(cdc, opWeightMsgAddShares, &weightMsgAddShares, nil,
		func(_ *rand.Rand) {
			weightMsgAddShares = defaultWeightMsgAddShares
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
			weightMsgCreateCampaign,
			campaignsim.SimulateMsgCreateCampaign(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgEditCampaign,
			campaignsim.SimulateMsgEditCampaign(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateTotalSupply,
			campaignsim.SimulateMsgUpdateTotalSupply(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
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
			weightMsgAddShares,
			campaignsim.SimulateMsgUpdateSpecialAllocations(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgMintVouchers,
			campaignsim.SimulateMsgMintVouchers(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgBurnVouchers,
			campaignsim.SimulateMsgBurnVouchers(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRedeemVouchers,
			campaignsim.SimulateMsgRedeemVouchers(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUnredeemVouchers,
			campaignsim.SimulateMsgUnredeemVouchers(am.accountKeeper, am.bankKeeper, am.keeper),
		),
	}
}
