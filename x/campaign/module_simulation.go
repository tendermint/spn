package campaign

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

const (
	weightMsgCreateCampaign    = 50
	weightMsgUpdateTotalSupply = 20
	weightMsgUpdateTotalShares = 20
	weightMsgInitializeMainnet = 20
	weightMsgAddShares         = 20
	weightMsgAddVestingOptions = 20
	weightMsgMintVouchers      = 20
	weightMsgBurnVouchers      = 20
	weightMsgRedeemVouchers    = 20
	weightMsgUnredeemVouchers  = 20
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	campaignGenesis := sample.CampaignGenesisState()
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&campaignGenesis)
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
			SimulateMsgCreateCampaign(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateTotalSupply,
			SimulateMsgUpdateTotalSupply(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateTotalShares,
			SimulateMsgUpdateTotalShares(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgInitializeMainnet,
			SimulateMsgInitializeMainnet(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgAddShares,
			SimulateMsgAddShares(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgAddVestingOptions,
			SimulateMsgAddVestingOptions(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgMintVouchers,
			SimulateMsgMintVouchers(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgBurnVouchers,
			SimulateMsgBurnVouchers(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRedeemVouchers,
			SimulateMsgRedeemVouchers(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUnredeemVouchers,
			SimulateMsgUnredeemVouchers(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
	}
}

// getCoordinatorSimAccount finds an account associated with a coordinator profile from simulation accounts
func getCoordinatorSimAccount(ctx sdk.Context, pk types.ProfileKeeper, accs []simtypes.Account) (simtypes.Account, bool) {
	for _, acc := range accs {
		_, found := pk.CoordinatorIDFromAddress(ctx, acc.Address.String())
		if found {
			return acc, true
		}
	}
	return simtypes.Account{}, false
}

// SimulateMsgCreateCampaign simulates a MsgCreateCampaign message
func SimulateMsgCreateCampaign(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, found := getCoordinatorSimAccount(ctx, pk, accs)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateCampaign, "skip campaign creation"), nil, nil
		}

		dynamicShares := r.Intn(100) > 80

		msg := types.NewMsgCreateCampaign(
			simAccount.Address.String(),
			sample.CampaignName(),
			sample.Coins(),
			dynamicShares,
		)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateTotalSupply simulates a MsgUpdateTotalSupply message
func SimulateMsgUpdateTotalSupply(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateTotalSupply, "skip update total supply"), nil, nil
	}
}

// SimulateMsgUpdateTotalShares simulates a MsgUpdateTotalShares message
func SimulateMsgUpdateTotalShares(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateTotalShares, "skip update total shares"), nil, nil
	}
}

// SimulateMsgInitializeMainnet simulates a MsgInitializeMainnet message
func SimulateMsgInitializeMainnet(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgInitializeMainnet, "skip initliaze mainnet"), nil, nil
	}
}

// SimulateMsgAddShares simulates a MsgAddShares message
func SimulateMsgAddShares(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddShares, "skip add shares"), nil, nil
	}
}

// SimulateMsgAddVestingOptions simulates a MsgAddVestingOptions message
func SimulateMsgAddVestingOptions(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddVestingOptions, "skip add vesting options"), nil, nil
	}
}

// SimulateMsgMintVouchers simulates a MsgMintVouchers message
func SimulateMsgMintVouchers(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgMintVouchers, "skip mint vouchers"), nil, nil
	}
}

// SimulateMsgBurnVouchers simulates a MsgBurnVouchers message
func SimulateMsgBurnVouchers(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBurnVouchers, "skip burn vouchers"), nil, nil
	}
}

// SimulateMsgRedeemVouchers simulates a MsgRedeemVouchers message
func SimulateMsgRedeemVouchers(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRedeemVouchers, "skip redeem vouchers"), nil, nil
	}
}

// SimulateMsgUnredeemVouchers simulates a MsgUnredeemVouchers message
func SimulateMsgUnredeemVouchers(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnredeemVouchers, "skip unredeem vouchers"), nil, nil
	}
}
