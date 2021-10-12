package launch

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

const (
	weightMsgCreateChain                 = 50
	weightMsgEditChain                   = 20
	weightMsgRequestAddGenesisAccount    = 50
	weightMsgRequestRemoveGenesisAccount = 50
	weightMsgRequestAddVestingAccount    = 50
	weightMsgRequestRemoveVestingAccount = 10
	weightMsgRequestAddValidator         = 50
	weightMsgRequestRemoveValidator      = 10
	weightMsgTriggerLaunch               = 10
	weightMsgRevertLaunch                = 10
	weightMsgSettleRequest               = 50
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	launchGenesis := sample.LaunchGenesisState()
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&launchGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	launchParams := sample.LaunchParams()

	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMinLaunchTime), func(r *rand.Rand) string {
			return fmt.Sprintf("\"%d\"", launchParams.MinLaunchTime)
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMaxLaunchTime), func(r *rand.Rand) string {
			return fmt.Sprintf("\"%d\"", launchParams.MaxLaunchTime)

		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return []simtypes.WeightedOperation{
		simulation.NewWeightedOperation(
			weightMsgCreateChain,
			SimulateMsgCreateChain(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgEditChain,
			SimulateMsgEditChain(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestAddGenesisAccount,
			SimulateMsgRequestAddGenesisAccount(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestRemoveGenesisAccount,
			SimulateMsgRequestRemoveGenesisAccount(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestAddVestingAccount,
			SimulateMsgRequestAddVestingAccount(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestRemoveVestingAccount,
			SimulateMsgRequestRemoveVestingAccount(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestAddValidator,
			SimulateMsgRequestAddValidator(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRequestRemoveValidator,
			SimulateMsgRequestRemoveValidator(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgTriggerLaunch,
			SimulateMsgTriggerLaunch(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRevertLaunch,
			SimulateMsgRevertLaunch(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgSettleRequest,
			SimulateMsgSettleRequest(am.accountKeeper, am.bankKeeper, am.keeper),
		),
	}
}

// SimulateMsgCreateChain simulates a MsgCreateChain message
func SimulateMsgCreateChain(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a random account
		accountNb := r.Intn(len(accs))
		msg := sample.MsgCreateChain(
			accs[accountNb].Address.String(),
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
			SimAccount:      accs[accountNb],
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
		// Select a random chain
		chains := k.GetAllChain(ctx)
		if len(chains) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgEditChain, "chain not found"), nil, nil
		}
		chainNb := r.Intn(len(chains))
		chain := chains[chainNb]

		coordAddr, found := k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
		if !found {
			// No message if no coordinator address
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgEditChain, "coordinator address not found"), nil, nil
		}

		var simAccount simtypes.Account
		for _, acc := range accs {
			if acc.Address.String() == coordAddr {
				simAccount = acc
				break
			}
		}
		if len(simAccount.Address.Bytes()) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgEditChain, "coordinator address not found in the sim accounts"), nil, nil
		}

		msg := sample.MsgEditChain(
			coordAddr,
			chain.Id,
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
		// Select a random account
		accountNb := r.Intn(len(accs))

		// Select a random chain
		chains := k.GetAllChain(ctx)
		if len(chains) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddAccount, "chain not found"), nil, nil
		}
		chainNb := r.Intn(len(chains))
		chain := chains[chainNb]

		// Select between new address or coordinator address randomly
		creator := sample.Address()
		if rand.Intn(2) == 1 {
			var found bool
			creator, found = k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
			if !found {
				// No message if no coordinator address
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddAccount, "coordinator address not found"), nil, nil
			}
		}

		msg := sample.MsgRequestAddAccount(
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
			SimAccount:      accs[accountNb],
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
		// Select a random account
		accountNb := r.Intn(len(accs))

		// Select a random chain
		chains := k.GetAllChain(ctx)
		if len(chains) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveGenesisAccount, "chain not found"), nil, nil
		}
		chainNb := r.Intn(len(chains))
		chain := chains[chainNb]

		// Select a random genesis account
		genAccs := k.GetAllGenesisAccount(ctx)
		if len(genAccs) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveGenesisAccount, "genesis account not found"), nil, nil
		}
		genAccNb := r.Intn(len(genAccs))
		genAcc := genAccs[genAccNb]

		// Select between new address or coordinator address randomly
		creator := sample.Address()
		if rand.Intn(2) == 1 {
			var found bool
			creator, found = k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
			if !found {
				// No message if no coordinator address
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveGenesisAccount, "coordinator address not found"), nil, nil
			}
		}

		msg := sample.MsgRequestRemoveAccount(
			creator,
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
			SimAccount:      accs[accountNb],
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
		// Select a random account
		accountNb := r.Intn(len(accs))

		// Select a random chain
		chains := k.GetAllChain(ctx)
		if len(chains) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddValidator, "chain not found"), nil, nil
		}
		chainNb := r.Intn(len(chains))
		chain := chains[chainNb]

		// Select between new address or coordinator address randomly
		creator := sample.Address()
		if rand.Intn(2) == 1 {
			var found bool
			creator, found = k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
			if !found {
				// No message if no coordinator address
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddValidator, "coordinator address not found"), nil, nil
			}
		}

		msg := sample.MsgRequestAddValidator(
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
			SimAccount:      accs[accountNb],
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
		// Select a random account
		accountNb := r.Intn(len(accs))

		// Select a random chain
		chains := k.GetAllChain(ctx)
		if len(chains) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveValidator, "chain not found"), nil, nil
		}
		chainNb := r.Intn(len(chains))
		chain := chains[chainNb]

		// Select a random validator
		valAccs := k.GetAllGenesisValidator(ctx)
		if len(valAccs) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveValidator, "validator not found"), nil, nil
		}
		valAccNb := r.Intn(len(valAccs))
		valAcc := valAccs[valAccNb]

		// Select between new address or coordinator address randomly
		creator := sample.Address()
		if rand.Intn(2) == 1 {
			var found bool
			creator, found = k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
			if !found {
				// No message if no coordinator address
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveValidator, "coordinator address not found"), nil, nil
			}
		}

		msg := sample.MsgRequestRemoveValidator(
			creator,
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
			SimAccount:      accs[accountNb],
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
		// Select a random account
		accountNb := r.Intn(len(accs))

		// Select a random chain
		chains := k.GetAllChain(ctx)
		if len(chains) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddVestingAccount, "chain not found"), nil, nil
		}
		chainNb := r.Intn(len(chains))
		chain := chains[chainNb]

		// Select between new address or coordinator address randomly
		creator := sample.Address()
		if rand.Intn(2) == 1 {
			var found bool
			creator, found = k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
			if !found {
				// No message if no coordinator address
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestAddVestingAccount, "coordinator address not found"), nil, nil
			}
		}

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
			SimAccount:      accs[accountNb],
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
		// Select a random account
		accountNb := r.Intn(len(accs))

		// Select a random chain
		chains := k.GetAllChain(ctx)
		if len(chains) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveVestingAccount, "chain not found"), nil, nil
		}
		chainNb := r.Intn(len(chains))
		chain := chains[chainNb]

		// Select a random vesting account
		vestAccs := k.GetAllVestingAccount(ctx)
		if len(vestAccs) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveVestingAccount, "vesting account not found"), nil, nil
		}
		vestAccNb := r.Intn(len(vestAccs))
		vestAcc := vestAccs[vestAccNb]

		// Select between new address or coordinator address randomly
		creator := sample.Address()
		if rand.Intn(2) == 1 {
			var found bool
			creator, found = k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
			if !found {
				// No message if no coordinator address
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveVestingAccount, "coordinator address not found"), nil, nil
			}
		}

		msg := sample.MsgRequestRemoveAccount(
			creator,
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
			SimAccount:      accs[accountNb],
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
		// Select a random account
		accountNb := r.Intn(len(accs))

		// Select a random chain
		var (
			found bool
			chain types.Chain
		)
		chains := k.GetAllChain(ctx)
		for _, c := range chains {
			if !c.LaunchTriggered {
				chain, found = c, true
				break
			}
		}
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgTriggerLaunch, "non-triggered chain not found"), nil, nil
		}

		coordinator, found := k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
		if !found {
			// No message if no coordinator address
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgTriggerLaunch, "coordinator address not found"), nil, nil
		}

		msg := sample.MsgTriggerLaunch(coordinator, chain.Id)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      accs[accountNb],
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
		// Select a random account
		accountNb := r.Intn(len(accs))

		// Select a random request
		requests := k.GetAllRequest(ctx)
		if len(requests) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSettleRequest, "request not found"), nil, nil
		}

		requestNb := r.Intn(len(requests))
		request := requests[requestNb]

		chain, found := k.GetChain(ctx, request.ChainID)
		if !found {
			// No message if the chain not found
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSettleRequest, "chain not found"), nil, nil
		}

		coordinator, found := k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
		if !found {
			// No message if no coordinator address
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgTriggerLaunch, "coordinator address not found"), nil, nil
		}

		approve := rand.Intn(2) == 1
		msg := sample.MsgSettleRequest(
			coordinator,
			chain.Id,
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
			SimAccount:      accs[accountNb],
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
		// Select a random account
		accountNb := r.Intn(len(accs))

		// Select a random chain
		var (
			found bool
			chain types.Chain
		)
		chains := k.GetAllChain(ctx)
		for _, c := range chains {
			if c.LaunchTriggered {
				chain, found = c, true
				break
			}
		}
		if !found {
			// No message if no triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, "triggered chain not found"), nil, nil
		}

		coordinator, found := k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
		if !found {
			// No message if no coordinator address
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, "coordinator address not found"), nil, nil
		}

		msg := sample.MsgRevertLaunch(coordinator, chain.Id)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      accs[accountNb],
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
