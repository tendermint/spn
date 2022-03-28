package campaign

import (
	"fmt"
	"github.com/tendermint/spn/testutil/sample"
	"math/rand"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	campaignsim "github.com/tendermint/spn/x/campaign/simulation"
	"github.com/tendermint/spn/x/campaign/types"
)

const (
	defaultWeightMsgCreateCampaign    = 25
	defaultWeightMsgUpdateTotalSupply = 20
	defaultWeightMsgUpdateTotalShares = 20
	defaultWeightMsgInitializeMainnet = 5
	defaultWeightMsgAddShares         = 20
	defaultWeightMsgAddVestingOptions = 20
	defaultWeightMsgMintVouchers      = 20
	defaultWeightMsgBurnVouchers      = 20
	defaultWeightMsgRedeemVouchers    = 20
	defaultWeightMsgUnredeemVouchers  = 20
	defaultWeightMsgSendVouchers      = 20

	opWeightMsgCreateCampaign    = "op_weight_msg_create_campaign"
	opWeightMsgUpdateTotalSupply = "op_weight_msg_update_total_supply"
	opWeightMsgUpdateTotalShares = "op_weight_msg_update_total_share"
	opWeightMsgInitializeMainnet = "op_weight_msg_initialize_mainnet"
	opWeightMsgAddShares         = "op_weight_msg_add_shares"
	opWeightMsgAddVestingOptions = "op_weight_msg_add_vesting_options"
	opWeightMsgMintVouchers      = "op_weight_msg_mint_vouchers"
	opWeightMsgBurnVouchers      = "op_weight_msg_burn_vouchers"
	opWeightMsgRedeemVouchers    = "op_weight_msg_redeem_vouchers"
	opWeightMsgUnredeemVouchers  = "op_weight_msg_unredeem_vouchers"
	opWeightMsgSendVouchers      = "op_weight_msg_send_vouchers"
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
	creationFee := make([]string, len(campaignParams.CampaignCreationFee))
	for i := range campaignParams.CampaignCreationFee {
		creationFee[i] = fmt.Sprintf(
			"{\"denom\":\"%v\",\"amount\":\"%v\"}",
			campaignParams.CampaignCreationFee[i].Denom,
			campaignParams.CampaignCreationFee[i].Amount.String(),
		)
	}
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyTotalSupplyRange), func(r *rand.Rand) string {
			return fmt.Sprintf(
				"{\"minTotalSupply\":\"%v\",\"maxTotalSupply\":\"%v\"}",
				campaignParams.TotalSupplyRange.MinTotalSupply,
				campaignParams.TotalSupplyRange.MaxTotalSupply)
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyCampaignCreationFee), func(r *rand.Rand) string {
			return fmt.Sprintf("[%v]", strings.Join(creationFee, ","))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	var (
		weightMsgCreateCampaign    int
		weightMsgUpdateTotalSupply int
		weightMsgUpdateTotalShares int
		weightMsgInitializeMainnet int
		weightMsgAddShares         int
		weightMsgAddVestingOptions int
		weightMsgMintVouchers      int
		weightMsgBurnVouchers      int
		weightMsgRedeemVouchers    int
		weightMsgUnredeemVouchers  int
		weightMsgSendVouchers      int
	)

	appParams := simState.AppParams
	cdc := simState.Cdc
	appParams.GetOrGenerate(cdc, opWeightMsgCreateCampaign, &weightMsgCreateCampaign, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCampaign = defaultWeightMsgCreateCampaign
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgUpdateTotalSupply, &weightMsgUpdateTotalSupply, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTotalSupply = defaultWeightMsgUpdateTotalSupply
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgUpdateTotalShares, &weightMsgUpdateTotalShares, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTotalShares = defaultWeightMsgUpdateTotalShares
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgInitializeMainnet, &weightMsgInitializeMainnet, nil,
		func(_ *rand.Rand) {
			weightMsgInitializeMainnet = defaultWeightMsgInitializeMainnet
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgAddShares, &weightMsgAddShares, nil,
		func(_ *rand.Rand) {
			weightMsgAddShares = defaultWeightMsgAddShares
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgAddVestingOptions, &weightMsgAddVestingOptions, nil,
		func(_ *rand.Rand) {
			weightMsgAddVestingOptions = defaultWeightMsgAddVestingOptions
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
	appParams.GetOrGenerate(cdc, opWeightMsgSendVouchers, &weightMsgSendVouchers, nil,
		func(_ *rand.Rand) {
			weightMsgSendVouchers = defaultWeightMsgSendVouchers
		},
	)

	return []simtypes.WeightedOperation{
		simulation.NewWeightedOperation(
			weightMsgCreateCampaign,
			campaignsim.SimulateMsgCreateCampaign(am.keeper, am.accountKeeper, am.bankKeeper, am.profileKeeper),
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
