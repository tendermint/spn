package profile

import (
	"github.com/tendermint/spn/x/profile/keeper"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
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
			SimulateMsgCreateCoordinator(am.accountKeeper, am.bankKeeper, am.keeper),
		),
	}
}

// SimulateMsgCreateCoordinator simulates a MsgCreateCoordinator message
func SimulateMsgCreateCoordinator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var simAccount simtypes.Account

		// Find an account with no coordinator
		allCreated := true
		for _, acc := range accs {
			_, found := k.GetCoordinatorByAddress(ctx, acc.Address.String())
			if !found {
				simAccount, allCreated = acc, false
				break
			}
		}
		if allCreated {
			// No message if all coordinator created
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateCoordinator, "skip coordinator creation"), nil, nil
		}

		msg := types.NewMsgCreateCoordinator(
			simAccount.Address.String(),
			sample.String(30),
			sample.String(30),
			sample.String(30),
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
