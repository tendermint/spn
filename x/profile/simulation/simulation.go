package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/testutil/simulation"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

// generate a Tx with 2 signatures if validator account is not equal to operator address account
func genAndDeliverTxWithRandFeesAddOpAddr(txCtx sdksimulation.OperationInput, opSimAcc simtypes.Account) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	spendable := txCtx.Bankkeeper.SpendableCoins(txCtx.Context, account.GetAddress())
	opAccout := txCtx.AccountKeeper.GetAccount(txCtx.Context, opSimAcc.Address)

	accNumbers := []uint64{account.GetAccountNumber()}
	accSequences := []uint64{account.GetSequence()}
	privs := []cryptotypes.PrivKey{txCtx.SimAccount.PrivKey}

	if account != opAccout {
		accNumbers = append(accNumbers, opAccout.GetAccountNumber())
		accSequences = append(accSequences, opAccout.GetSequence())
		privs = append(privs, opSimAcc.PrivKey)
	}

	var fees sdk.Coins
	var err error

	coins, hasNeg := spendable.SafeSub(txCtx.CoinsSpentInMsg)
	if hasNeg {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "message doesn't leave room for fees"), nil, err
	}

	fees, err = sample.Fees(txCtx.R, coins)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to generate fees"), nil, err
	}

	tx, err := helpers.GenTx(
		txCtx.TxGen,
		[]sdk.Msg{txCtx.Msg},
		fees,
		helpers.DefaultGenTxGas,
		txCtx.Context.ChainID(),
		accNumbers,
		accSequences,
		privs...,
	)

	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to generate mock tx"), nil, err
	}

	_, _, err = txCtx.App.Deliver(txCtx.TxGen.TxEncoder(), tx)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to deliver tx"), nil, err
	}

	return simtypes.NewOperationMsg(txCtx.Msg, true, "", txCtx.Cdc), nil, nil
}

// SimulateMsgUpdateValidatorDescription simulates a MsgUpdateValidatorDescription message
func SimulateMsgUpdateValidatorDescription(ak types.AccountKeeper, bk types.BankKeeper, _ keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a random account
		simAccount, _ := simtypes.RandomAcc(r, accs)

		desc := sample.ValidatorDescription(sample.String(r, 50))
		msg := types.NewMsgUpdateValidatorDescription(
			simAccount.Address.String(),
			desc.Identity,
			desc.Moniker,
			desc.Website,
			desc.SecurityContact,
			desc.Details,
		)
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

// SimulateMsgAddValidatorOperatorAddress simulates a MsgAddValidatorOperatorAddress message
func SimulateMsgAddValidatorOperatorAddress(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// choose two addresses that are not equal
		simAccount, _ := simtypes.RandomAcc(r, accs)
		opAccount, _ := simtypes.RandomAcc(r, accs)
		for simAccount.Address.Equals(opAccount.Address) {
			opAccount, _ = simtypes.RandomAcc(r, accs)
		}

		msg := &types.MsgAddValidatorOperatorAddress{
			ValidatorAddress: simAccount.Address.String(),
			OperatorAddress:  opAccount.Address.String(),
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
		return genAndDeliverTxWithRandFeesAddOpAddr(txCtx, opAccount)
	}
}

// SimulateMsgCreateCoordinator simulates a MsgCreateCoordinator message
func SimulateMsgCreateCoordinator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Find an account with no coordinator
		simAccount, found := FindCoordinatorAccount(r, ctx, k, accs, false)
		if !found {
			// No message if all coordinator created
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateCoordinator, "skip coordinator creation"), nil, nil
		}

		msg := types.NewMsgCreateCoordinator(
			simAccount.Address.String(),
			sample.String(r, 30),
			sample.String(r, 30),
			sample.String(r, 30),
		)
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

// SimulateMsgUpdateCoordinatorDescription simulates a MsgUpdateCoordinatorDescription message
func SimulateMsgUpdateCoordinatorDescription(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Find an account with coordinator associated
		simAccount, found := FindCoordinatorAccount(r, ctx, k, accs, true)
		if !found {
			// No message if no coordinator
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateCoordinatorDescription, "skip update coordinator description"), nil, nil
		}

		desc := sample.CoordinatorDescription(r)
		msg := types.NewMsgUpdateCoordinatorDescription(
			simAccount.Address.String(),
			desc.Identity,
			desc.Website,
			desc.Details,
		)

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

// SimulateMsgUpdateCoordinatorAddress simulates a MsgUpdateCoordinatorAddress message
func SimulateMsgUpdateCoordinatorAddress(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a random account
		coord, found := FindCoordinatorAccount(r, ctx, k, accs, true)
		if !found {
			// No message if no coordinator
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateCoordinatorAddress, "skip update coordinator address"), nil, nil
		}
		simAccount, found := FindCoordinatorAccount(r, ctx, k, accs, false)
		if !found && coord.Address.String() != simAccount.Address.String() {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateCoordinatorAddress, "skip update coordinator address"), nil, nil
		}
		msg := types.NewMsgUpdateCoordinatorAddress(coord.Address.String(), simAccount.Address.String())
		txCtx := sdksimulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      coord,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
	}
}

// SimulateMsgDisableCoordinator simulates a MsgDisableCoordinator message
func SimulateMsgDisableCoordinator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Find an account with coordinator associated
		// avoid delete coordinator associated a chain (id 0,1,2)
		simAccount, found := FindCoordinatorAccount(r, ctx, k, accs[3:], true)
		if !found {
			// No message if no coordinator
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgDisableCoordinator, "skip update coordinator delete"), nil, nil
		}

		msg := types.NewMsgDisableCoordinator(simAccount.Address.String())
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
