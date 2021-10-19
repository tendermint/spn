package launch

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
			_, found = k.GetProfileKeeper().CoordinatorIDFromAddress(ctx, acc.Address.String())
			if found {
				simAccount = acc
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
		var (
			err        error
			found      bool
			chain      uint64
			simAccount simtypes.Account
		)
		chains := k.GetAllChain(ctx)
		for _, c := range chains {
			simAccount, err = FindChainCoordinatorAccount(ctx, k, accs, c.Id)
			if err != nil {
				continue
			}
			chain = c.Id
			found = true
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgEditChain, "chain not found"), nil, nil
		}

		msg := sample.MsgEditChain(
			simAccount.Address.String(),
			chain,
			true,
			true,
			true,
			true,
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
		chain, found := FindChain(ctx, k, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddAccount, "non-triggered chain not found"), nil, nil
		}

		// Select a random account
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := sample.MsgRequestAddAccount(
			simAccount.Address.String(),
			chain.Id,
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

// SimulateMsgRequestRemoveGenesisAccount simulates a MsgRequestRemoveAccount message for genesis account
func SimulateMsgRequestRemoveGenesisAccount(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a genesis account
		var (
			simAccount simtypes.Account
			genAcc     types.GenesisAccount
			err        error
		)
		found := false
		genAccs := k.GetAllGenesisAccount(ctx)
		for _, acc := range genAccs {
			if IsLaunchTriggeredChain(ctx, k, acc.ChainID) {
				continue
			}
			// get coordinator account for removal
			simAccount, err = FindChainCoordinatorAccount(ctx, k, accs, acc.ChainID)
			if err != nil {
				continue
			}
			genAcc = acc
			found = true
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveGenesisAccount, "genesis account not found"), nil, nil
		}

		msg := sample.MsgRequestRemoveAccount(
			simAccount.Address.String(),
			genAcc.Address,
			genAcc.ChainID,
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
		chain, found := FindChain(ctx, k, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddValidator, "non-triggered chain not found"), nil, nil
		}
		// Select a random account
		simAccount, _ := simtypes.RandomAcc(r, accs)
		// Select between new address or coordinator address randomly
		msg := sample.MsgRequestAddValidator(
			simAccount.Address.String(),
			chain.Id,
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
		var (
			valAcc     types.GenesisValidator
			simAccount simtypes.Account
			err        error
		)
		found := false
		valAccs := k.GetAllGenesisValidator(ctx)
		for _, acc := range valAccs {
			if IsLaunchTriggeredChain(ctx, k, acc.ChainID) {
				continue
			}
			// get coordinator account for removal
			simAccount, err = FindChainCoordinatorAccount(ctx, k, accs, acc.ChainID)
			if err != nil {
				continue
			}
			valAcc = acc
			found = true
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveValidator, "genesis account not found"), nil, nil
		}

		msg := sample.MsgRequestRemoveValidator(
			simAccount.Address.String(),
			valAcc.Address,
			valAcc.ChainID,
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
		chain, found := FindChain(ctx, k, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgTriggerLaunch, "non-triggered chain not found"), nil, nil
		}

		// Select a random account
		simAccount, _ := simtypes.RandomAcc(r, accs)
		creator := simAccount.Address.String()
		msg := sample.MsgRequestAddVestingAccount(
			creator,
			chain.Id,
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

// SimulateMsgRequestRemoveVestingAccount simulates a MsgRequestRemoveAccount message for vesting account
func SimulateMsgRequestRemoveVestingAccount(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a vesting account
		var (
			vestAcc    types.VestingAccount
			simAccount simtypes.Account
			err        error
		)
		found := false
		vestAccs := k.GetAllVestingAccount(ctx)
		for _, acc := range vestAccs {
			if IsLaunchTriggeredChain(ctx, k, acc.ChainID) {
				continue
			}
			// get coordinator account for removal
			simAccount, err = FindChainCoordinatorAccount(ctx, k, accs, acc.ChainID)
			if err != nil {
				continue
			}
			vestAcc = acc
			found = true
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveVestingAccount, "genesis account not found"), nil, nil
		}

		msg := sample.MsgRequestRemoveAccount(
			simAccount.Address.String(),
			vestAcc.Address,
			vestAcc.ChainID,
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
		chain, found := FindChain(ctx, k, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgTriggerLaunch, "non-triggered chain not found"), nil, nil
		}

		// Find coordinator account
		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, chain.Id)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgTriggerLaunch, err.Error()), nil, err
		}
		msg := sample.MsgTriggerLaunch(simAccount.Address.String(), chain.Id)
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
		requests := k.GetAllRequest(ctx)
		var request types.Request
		chainNotFound := true
		for chainNotFound {
			requestNb := r.Intn(len(requests))
			request = requests[requestNb]
			chain, found := k.GetChain(ctx, request.ChainID)
			if !found || chain.LaunchTriggered {
				continue
			}
			// check if the coordinator is still in the store
			_, found = k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
			if !found {
				continue
			}
			switch content := request.Content.Content.(type) {
			case *types.RequestContent_ValidatorRemoval:
				// if is validator removal, check if the validator exist
				if _, found := k.GetGenesisValidator(
					ctx,
					chain.Id,
					content.ValidatorRemoval.ValAddress,
				); !found {
					continue
				}
			case *types.RequestContent_AccountRemoval:
				// if is account removal, check if account exist
				found, err := keeper.CheckAccount(ctx, k, chain.Id, content.AccountRemoval.Address)
				if err != nil || !found {
					continue
				}
			}
			chainNotFound = false
		}
		if chainNotFound {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSettleRequest, "request for non-triggered chain not found"), nil, nil
		}

		// Find coordinator account
		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, request.ChainID)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSettleRequest, err.Error()), nil, err
		}

		approve := r.Intn(100) < 50
		msg := sample.MsgSettleRequest(
			simAccount.Address.String(),
			request.ChainID,
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
		chain, found := FindChain(ctx, k, true)
		if !found {
			// No message if no triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, "triggered chain not found"), nil, nil
		}

		// Wait for a specific delay once the chain is launched
		if ctx.BlockTime().Unix() < chain.LaunchTimestamp+types.RevertDelay {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, "invalid chain launch timestamp"), nil, nil
		}

		// Find coordinator account
		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, chain.Id)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, err.Error()), nil, err
		}
		msg := sample.MsgRevertLaunch(simAccount.Address.String(), chain.Id)
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
