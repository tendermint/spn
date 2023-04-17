package simulation

import (
	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/testutil/simulation"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
	"math/rand"
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
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
		setProjectID := r.Intn(100) < 50
		// do not set projectID if already set
		if chain.HasProject {
			setProjectID = false
		}
		// ensure there is always a value to edit
		if !modify && !setProjectID {
			modify = true
		}

		projectID := uint64(0)
		ok := false
		if setProjectID {
			projectID, ok = FindCoordinatorProject(r, ctx, k.GetProjectKeeper(), chain.CoordinatorID, chain.LaunchID)
			if !ok {
				modify = true
				setProjectID = false
			}
		}

		msg := sample.MsgEditChain(r,
			simAccount.Address.String(),
			chain.LaunchID,
			setProjectID,
			projectID,
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
	}
}

// SimulateMsgRequestAddGenesisAccount simulates a MsgRequestAddGenesisAccount message
func SimulateMsgRequestAddGenesisAccount(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		fee := k.RequestFee(ctx)

		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, true)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSendRequest, "non-triggered chain not found"), nil, nil
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
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSendRequest, "no available account"), nil, nil
		}

		msg := sample.MsgSendRequestWithAddAccount(r,
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
	}
}

// SimulateMsgRequestAddVestingAccount simulates a MsgRequestAddVestingAccount message
func SimulateMsgRequestAddVestingAccount(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		fee := k.RequestFee(ctx)

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
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSendRequest, "no available account"), nil, nil
		}

		msg := sample.MsgSendRequestWithAddVestingAccount(r,
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
	}
}

// SimulateMsgRequestRemoveAccount simulates a MsgRequestRemoveAccount message
func SimulateMsgRequestRemoveAccount(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		fee := k.RequestFee(ctx)

		type accChain struct {
			address  string
			launchID uint64
		}

		// build list of genesis and vesting accounts
		accChains := make([]accChain, 0)
		genAccs := k.GetAllGenesisAccount(ctx)
		for _, acc := range genAccs {
			accChains = append(accChains, accChain{
				address:  acc.Address,
				launchID: acc.LaunchID,
			})
		}
		vestAccs := k.GetAllVestingAccount(ctx)
		for _, acc := range vestAccs {
			accChains = append(accChains, accChain{
				address:  acc.Address,
				launchID: acc.LaunchID,
			})
		}

		// add entropy
		r.Shuffle(len(accChains), func(i, j int) {
			accChains[i], accChains[j] = accChains[j], accChains[i]
		})

		var (
			simAccount simtypes.Account
			accAddr    string
			accChainID uint64
		)
		found := false
		for _, accChain := range accChains {
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
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSendRequest, "genesis account not found"), nil, nil
		}

		msg := types.NewMsgSendRequest(
			simAccount.Address.String(),
			accChainID,
			types.NewAccountRemoval(accAddr),
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
	}
}

// SimulateMsgRequestAddValidator simulates a MsgRequestAddValidator message
func SimulateMsgRequestAddValidator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		fee := k.RequestFee(ctx)

		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSendRequest, "non-triggered chain not found"), nil, nil
		}
		// Select a random account
		simAccount, _ := simtypes.RandomAcc(r, accs)
		// Select between new address or coordinator address randomly
		msg := sample.MsgSendRequestWithAddValidator(r,
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
	}
}

// SimulateMsgRequestRemoveValidator simulates a MsgRequestRemoveValidator message
func SimulateMsgRequestRemoveValidator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		fee := k.RequestFee(ctx)

		// Select a validator
		simAccount, valAcc, found := FindRandomValidator(r, ctx, k, accs)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSendRequest, "validator not found"), nil, nil
		}

		msg := sample.MsgSendRequestWithRemoveValidator(
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
	}
}

// SimulateMsgRequestParamChange simulates a MsgSendRequest message with ParamChange content
func SimulateMsgRequestParamChange(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		fee := k.RequestFee(ctx)

		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSendRequest, "non-triggered chain not found"), nil, nil
		}
		simAccount, _ := simtypes.RandomAcc(r, accs)

		msg := sample.MsgSendRequestWithParamChange(
			r,
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
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
		msg := sample.MsgTriggerLaunch(r, simAccount.Address.String(), chain.LaunchID, ctx.BlockTime())
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
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
		if ctx.BlockTime().Before(chain.LaunchTime.Add(k.RevertDelay(ctx))) {
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
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
		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
	}
}
