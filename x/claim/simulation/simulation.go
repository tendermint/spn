package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/simulation"
	"github.com/tendermint/spn/x/claim/keeper"
	"github.com/tendermint/spn/x/claim/types"
)

func SimulateMsgClaimInitial(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgClaimInitial{}

		// find an account
		simAccount, _ := simtypes.RandomAcc(r, accs)

		// check the account has a claim record and initial claim has not been completed
		cr, found := k.GetClaimRecord(ctx, simAccount.Address.String())
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "account has no claim record"), nil, nil
		}
		if cr.IsMissionCompleted(1) {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "account already completed initial claim"), nil, nil
		}

		// initialize basic message
		msg = &types.MsgClaimInitial{
			Claimer: simAccount.Address.String(),
		}

		txCtx := sdksimulation.OperationInput{
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

		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
	}
}
