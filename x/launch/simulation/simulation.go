package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

// SimulateMsgCreateChain simulates a MsgCreateChain message
func SimulateMsgCreateChain(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Check if the coordinator address is already in the store
		var found bool
		var simAccount simtypes.Account
		for _, acc := range accs {
			_, err := k.GetProfileKeeper().GetCoordinatorByAddress(ctx, acc.Address.String())
			if err == nil {
				simAccount = acc
				found = true
				break
			}
		}
		if !found {
			// No message if no coordinator
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateChain, "skip create a new chain"), nil, nil
		}
		msg := sample.MsgCreateChain(
			simAccount.Address.String(),
			"",
			false,
			0,
		)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
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

// SimulateMsgEditChain simulates a MsgEditChain message
func SimulateMsgEditChain(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a chain with a valid coordinator account
		chain, found := FindRandomChain(r, ctx, k, r.Intn(100) < 50, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgEditChain, "chain not found"), nil, nil
		}

		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, chain.LaunchID)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgEditChain, "coordinator account not found"), nil, nil
		}

		modify := r.Intn(100) < 50
		msg := sample.MsgEditChain(
			simAccount.Address.String(),
			chain.LaunchID,
			modify,
			!modify,
			modify,
			!modify && r.Intn(100) < 50,
		)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
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

// SimulateMsgRequestAddGenesisAccount simulates a MsgRequestAddGenesisAccount message
func SimulateMsgRequestAddGenesisAccount(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, true)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddAccount, "non-triggered chain not found"), nil, nil
		}

		// Select a random account
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := sample.MsgRequestAddAccount(
			simAccount.Address.String(),
			simAccount.Address.String(),
			chain.LaunchID,
		)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
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

// SimulateMsgRequestAddVestingAccount simulates a MsgRequestAddVestingAccount message
func SimulateMsgRequestAddVestingAccount(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, true)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgTriggerLaunch, "non-triggered chain not found"), nil, nil
		}

		// Select a random account
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := sample.MsgRequestAddVestingAccount(
			simAccount.Address.String(),
			simAccount.Address.String(),
			chain.LaunchID,
		)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
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

// SimulateMsgRequestRemoveAccount simulates a MsgRequestRemoveAccount message
func SimulateMsgRequestRemoveAccount(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		accsChain := make(map[string]uint64)
		genAccs := k.GetAllGenesisAccount(ctx)
		for _, acc := range genAccs {
			accsChain[acc.Address] = acc.LaunchID
		}
		vestAccs := k.GetAllVestingAccount(ctx)
		for _, acc := range vestAccs {
			accsChain[acc.Address] = acc.LaunchID
		}

		var (
			simAccount simtypes.Account
			accAddr    string
			accChainID uint64
		)
		found := false
		for acc, chainID := range accsChain {
			if IsLaunchTriggeredChain(ctx, k, chainID) {
				continue
			}
			// get coordinator account for removal
			var err error
			simAccount, err = FindChainCoordinatorAccount(ctx, k, accs, chainID)
			if err != nil {
				continue
			}
			accAddr = acc
			accChainID = chainID
			found = true
			break
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveAccount, "genesis account not found"), nil, nil
		}

		msg := sample.MsgRequestRemoveAccount(
			simAccount.Address.String(),
			accAddr,
			accChainID,
		)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
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

// SimulateMsgRequestAddValidator simulates a MsgRequestAddValidator message
func SimulateMsgRequestAddValidator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddValidator, "non-triggered chain not found"), nil, nil
		}
		// Select a random account
		simAccount, _ := simtypes.RandomAcc(r, accs)
		// Select between new address or coordinator address randomly
		msg := sample.MsgRequestAddValidator(
			simAccount.Address.String(),
			simAccount.Address.String(),
			chain.LaunchID,
		)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
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

// SimulateMsgRequestRemoveValidator simulates a MsgRequestRemoveValidator message
func SimulateMsgRequestRemoveValidator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a validator
		simAccount, valAcc, found := FindRandomValidator(r, ctx, k, accs)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveValidator, "validator not found"), nil, nil
		}

		msg := sample.MsgRequestRemoveValidator(
			simAccount.Address.String(),
			valAcc.Address,
			valAcc.LaunchID,
		)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
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

// SimulateMsgTriggerLaunch simulates a MsgTriggerLaunch message
func SimulateMsgTriggerLaunch(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgTriggerLaunch, "non-triggered chain not found"), nil, nil
		}

		// Find coordinator account
		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, chain.LaunchID)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgTriggerLaunch, err.Error()), nil, nil
		}
		msg := sample.MsgTriggerLaunch(simAccount.Address.String(), chain.LaunchID)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
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

// SimulateMsgSettleRequest simulates a MsgSettleRequest message
func SimulateMsgSettleRequest(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a random request without launch triggered
		request, found := FindRandomRequest(r, ctx, k)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSettleRequest, "request for non-triggered chain not found"), nil, nil
		}

		// Find coordinator account
		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, request.LaunchID)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSettleRequest, err.Error()), nil, nil
		}

		approve := r.Intn(100) < 50
		msg := sample.MsgSettleRequest(
			simAccount.Address.String(),
			request.LaunchID,
			request.RequestID,
			approve,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
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

// SimulateMsgRevertLaunch simulates a MsgRevertLaunch message
func SimulateMsgRevertLaunch(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a chain with launch triggered
		chain, found := FindRandomChain(r, ctx, k, true, false)
		if !found {
			// No message if no triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, "triggered chain not found"), nil, nil
		}

		// Wait for a specific delay once the chain is launched
		if ctx.BlockTime().Unix() < chain.LaunchTimestamp+types.RevertDelay {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, "invalid chain launch timestamp"), nil, nil
		}

		// Find coordinator account
		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, chain.LaunchID)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, err.Error()), nil, nil
		}
		msg := sample.MsgRevertLaunch(simAccount.Address.String(), chain.LaunchID)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
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
