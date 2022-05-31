package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/testutil/simulation"
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
			_, err := k.GetProfileKeeper().CoordinatorIDFromAddress(ctx, acc.Address.String())
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

		// skip if account cannot cover creation fee
		creationFee := k.ChainCreationFee(ctx)

		msg := sample.MsgCreateChain(r,
			simAccount.Address.String(),
			"",
			false,
			0,
		)
		txCtx := sdksimulation.OperationInput{
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
			CoinsSpentInMsg: creationFee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
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
		setCampaignID := r.Intn(100) < 50
		// do not set campaignID if already set
		if chain.HasCampaign {
			setCampaignID = false
		}
		// ensure there is always a value to edit
		if !modify && !setCampaignID {
			modify = true
		}

		campaignID := uint64(0)
		ok := false
		if setCampaignID {
			campaignID, ok = FindCoordinatorCampaign(r, ctx, k.GetCampaignKeeper(), chain.CoordinatorID, chain.LaunchID)
			if !ok {
				modify = true
				setCampaignID = false
			}
		}

		msg := sample.MsgEditChain(r,
			simAccount.Address.String(),
			chain.LaunchID,
			setCampaignID,
			campaignID,
			modify,
		)

		txCtx := sdksimulation.OperationInput{
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
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

		// Select a random account no in genesis
		r.Shuffle(len(accs), func(i, j int) {
			accs[i], accs[j] = accs[j], accs[i]
		})
		var simAccount simtypes.Account
		var availableAccount bool
		for _, acc := range accs {
			_, found := k.GetGenesisAccount(ctx, chain.LaunchID, acc.Address.String())
			if found {
				continue
			}
			_, found = k.GetVestingAccount(ctx, chain.LaunchID, acc.Address.String())
			if found {
				continue
			}
			simAccount = acc
			availableAccount = true
			break
		}
		if !availableAccount {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddAccount, "no available account"), nil, nil
		}

		msg := sample.MsgRequestAddAccount(r,
			simAccount.Address.String(),
			simAccount.Address.String(),
			chain.LaunchID,
		)
		txCtx := sdksimulation.OperationInput{
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
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

		// Select a random account no in genesis
		r.Shuffle(len(accs), func(i, j int) {
			accs[i], accs[j] = accs[j], accs[i]
		})
		var simAccount simtypes.Account
		var availableAccount bool
		for _, acc := range accs {
			_, found := k.GetGenesisAccount(ctx, chain.LaunchID, acc.Address.String())
			if found {
				continue
			}
			_, found = k.GetVestingAccount(ctx, chain.LaunchID, acc.Address.String())
			if found {
				continue
			}
			simAccount = acc
			availableAccount = true
			break
		}
		if !availableAccount {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddAccount, "no available account"), nil, nil
		}

		msg := sample.MsgRequestAddVestingAccount(r,
			simAccount.Address.String(),
			simAccount.Address.String(),
			chain.LaunchID,
		)
		txCtx := sdksimulation.OperationInput{
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
	}
}

// SimulateMsgRequestRemoveAccount simulates a MsgRequestRemoveAccount message
func SimulateMsgRequestRemoveAccount(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		type accChain struct {
			address  string
			launchID uint64
		}

		// build list of genesis and vesting accounts
		accChainList := make([]accChain, 0)
		genAccs := k.GetAllGenesisAccount(ctx)
		for _, acc := range genAccs {
			accChainList = append(accChainList, accChain{
				address:  acc.Address,
				launchID: acc.LaunchID,
			})
		}
		vestAccs := k.GetAllVestingAccount(ctx)
		for _, acc := range vestAccs {
			accChainList = append(accChainList, accChain{
				address:  acc.Address,
				launchID: acc.LaunchID,
			})
		}

		// add entropy
		r.Shuffle(len(accChainList), func(i, j int) {
			accChainList[i], accChainList[j] = accChainList[j], accChainList[i]
		})

		var (
			simAccount simtypes.Account
			accAddr    string
			accChainID uint64
		)
		found := false
		for _, accChain := range accChainList {
			if IsLaunchTriggeredChain(ctx, k, accChain.launchID) {
				continue
			}
			// get coordinator account
			var err error
			simAccount, err = FindChainCoordinatorAccount(ctx, k, accs, accChain.launchID)
			if err != nil {
				continue
			}
			accAddr = accChain.address
			accChainID = accChain.launchID
			found = true
			break
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveAccount, "genesis account not found"), nil, nil
		}

		msg := types.NewMsgRequestRemoveAccount(
			simAccount.Address.String(),
			accChainID,
			accAddr,
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
		msg := sample.MsgRequestAddValidator(r,
			simAccount.Address.String(),
			simAccount.Address.String(),
			chain.LaunchID,
		)
		txCtx := sdksimulation.OperationInput{
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
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
		txCtx := sdksimulation.OperationInput{
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
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
		msg := sample.MsgTriggerLaunch(r, simAccount.Address.String(), chain.LaunchID)
		txCtx := sdksimulation.OperationInput{
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
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

		// if we cannot check the request, reject
		if err := keeper.CheckRequest(ctx, k, request.LaunchID, request); err != nil {
			msg.Approve = false
		}

		txCtx := sdksimulation.OperationInput{
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
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
		if ctx.BlockTime().Unix() < chain.LaunchTimestamp+k.RevertDelay(ctx) {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, "invalid chain launch timestamp"), nil, nil
		}

		// Find coordinator account
		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, chain.LaunchID)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, err.Error()), nil, nil
		}
		msg := sample.MsgRevertLaunch(simAccount.Address.String(), chain.LaunchID)
		txCtx := sdksimulation.OperationInput{
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
	}
}

// SimulateMsgUpdateLaunchInformation simulates a MsgUpdateLaunchInformation message
func SimulateMsgUpdateLaunchInformation(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a chain with a valid coordinator account
		chain, found := FindRandomChain(r, ctx, k, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgEditChain, "chain not found"), nil, nil
		}

		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, chain.LaunchID)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgEditChain, "coordinator account not found"), nil, nil
		}

		modify := r.Intn(100) < 50
		msg := sample.MsgUpdateLaunchInformation(r,
			simAccount.Address.String(),
			chain.LaunchID,
			modify,
			!modify,
			modify,
			!modify && r.Intn(100) < 50,
		)
		txCtx := sdksimulation.OperationInput{
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
	}
}
