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
	defaultWeightMsgCreateChain                 int = 50
	defaultWeightMsgEditChain                   int = 10
	defaultWeightMsgRequestAddGenesisAccount    int = 50
	defaultWeightMsgRequestRemoveGenesisAccount int = 25
	defaultWeightMsgRequestAddVestingAccount    int = 50
	defaultWeightMsgRequestRemoveVestingAccount int = 25
	defaultWeightMsgRequestAddValidator         int = 50
	defaultWeightMsgRequestRemoveValidator      int = 25
	defaultWeightMsgTriggerLaunch               int = 15
	defaultWeightMsgRevertLaunch                int = 15
	defaultWeightMsgSettleRequest               int = 50

	opWeightMsgCreateChain                 = "op_weight_msg_create_chain"
	opWeightMsgEditChain                   = "op_weight_msg_edit_chain"
	opWeightMsgRequestAddGenesisAccount    = "op_weight_msg_request_add_genesis_account"
	opWeightMsgRequestRemoveGenesisAccount = "op_weight_msg_request_remove_genesis_account"
	opWeightMsgRequestAddVestingAccount    = "op_weight_msg_request_add_vesting_account"
	opWeightMsgRequestRemoveVestingAccount = "op_weight_msg_request_remove_vesting_account"
	opWeightMsgRequestAddValidator         = "op_weight_msg_request_add_validator"
	opWeightMsgRequestRemoveValidator      = "op_weight_msg_request_remove_validator"
	opWeightMsgTriggerLaunch               = "op_weight_msg_trigger_launch"
	opWeightMsgRevertLaunch                = "op_weight_msg_revert_launch"
	opWeightMsgSettleRequest               = "op_weight_msg_settle_request"
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	launchGenesis := sample.LaunchGenesisState(accs...)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&launchGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
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
	var (
		weightMsgCreateChain                 int
		weightMsgEditChain                   int
		weightMsgRequestAddGenesisAccount    int
		weightMsgRequestRemoveGenesisAccount int
		weightMsgRequestAddVestingAccount    int
		weightMsgRequestRemoveVestingAccount int
		weightMsgRequestAddValidator         int
		weightMsgRequestRemoveValidator      int
		weightMsgTriggerLaunch               int
		weightMsgRevertLaunch                int
		weightMsgSettleRequest               int
	)

	appParams := simState.AppParams
	cdc := simState.Cdc
	appParams.GetOrGenerate(cdc, opWeightMsgCreateChain, &weightMsgCreateChain, nil,
		func(_ *rand.Rand) {
			weightMsgCreateChain = defaultWeightMsgCreateChain
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgEditChain, &weightMsgEditChain, nil,
		func(_ *rand.Rand) {
			weightMsgEditChain = defaultWeightMsgEditChain
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestAddGenesisAccount, &weightMsgRequestAddGenesisAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddGenesisAccount = defaultWeightMsgRequestAddGenesisAccount
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestAddGenesisAccount, &weightMsgRequestAddGenesisAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddGenesisAccount = defaultWeightMsgRequestAddGenesisAccount
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestRemoveGenesisAccount, &weightMsgRequestRemoveGenesisAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRemoveGenesisAccount = defaultWeightMsgRequestRemoveGenesisAccount
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestAddVestingAccount, &weightMsgRequestAddVestingAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddVestingAccount = defaultWeightMsgRequestAddVestingAccount
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestRemoveVestingAccount, &weightMsgRequestRemoveVestingAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRemoveVestingAccount = defaultWeightMsgRequestRemoveVestingAccount
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestAddValidator, &weightMsgRequestAddValidator, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddValidator = defaultWeightMsgRequestAddValidator
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRequestRemoveValidator, &weightMsgRequestRemoveValidator, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRemoveValidator = defaultWeightMsgRequestRemoveValidator
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgTriggerLaunch, &weightMsgTriggerLaunch, nil,
		func(_ *rand.Rand) {
			weightMsgTriggerLaunch = defaultWeightMsgTriggerLaunch
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgRevertLaunch, &weightMsgRevertLaunch, nil,
		func(_ *rand.Rand) {
			weightMsgRevertLaunch = defaultWeightMsgRevertLaunch
		},
	)
	appParams.GetOrGenerate(cdc, opWeightMsgSettleRequest, &weightMsgSettleRequest, nil,
		func(_ *rand.Rand) {
			weightMsgSettleRequest = defaultWeightMsgSettleRequest
		},
	)

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

// findChain find a chain
func findChain(ctx sdk.Context, k keeper.Keeper, launchTriggered bool) (types.Chain, bool) {
	found := false
	chains := k.GetAllChain(ctx)
	var chain types.Chain
	for _, c := range chains {
		if c.LaunchTriggered == launchTriggered {
			chain = c
			found = true
			break
		}
	}
	return chain, found
}

// findChainCoordinatorAccount find coordinator account by chain id
func findChainCoordinatorAccount(ctx sdk.Context, k keeper.Keeper, accs []simtypes.Account, chainID uint64) (simtypes.Account, error) {
	chain, found := k.GetChain(ctx, chainID)
	if !found {
		// No message if no coordinator address
		return simtypes.Account{}, fmt.Errorf("chain %d not found", chainID)
	}
	address, found := k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return simtypes.Account{}, fmt.Errorf("coordinator %d not found", chain.CoordinatorID)
	}
	return findAccount(accs, address)
}

// findAccount find account by string hex address
func findAccount(accs []simtypes.Account, address string) (simtypes.Account, error) {
	coordAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return simtypes.Account{}, err
	}
	simAccount, found := simtypes.FindAccount(accs, coordAddr)
	if !found {
		return simAccount, fmt.Errorf("address %s not found in the sim accounts", address)
	}
	return simAccount, nil
}

// SimulateMsgCreateChain simulates a MsgCreateChain message
func SimulateMsgCreateChain(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Check if the coordinator address is already in the store
		coordinators := k.GetProfileKeeper().GetAllCoordinator(ctx)
		if len(coordinators) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateChain, "coordinator not found"), nil, nil
		}
		coordinatorsNb := r.Intn(len(coordinators))
		coordinator := coordinators[coordinatorsNb]

		simAccount, err := findAccount(accs, coordinator.Address)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateChain, err.Error()), nil, err
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
		// Select a random chain
		chains := k.GetAllChain(ctx)
		if len(chains) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgEditChain, "chain not found"), nil, nil
		}
		chainNb := r.Intn(len(chains))
		chain := chains[chainNb]

		simAccount, err := findChainCoordinatorAccount(ctx, k, accs, chain.Id)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgEditChain, err.Error()), nil, err
		}
		msg := sample.MsgEditChain(
			simAccount.Address.String(),
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
		// Select a chain without launch triggered
		chain, found := findChain(ctx, k, false)
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
		genAccs := k.GetAllGenesisAccount(ctx)
		if len(genAccs) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveGenesisAccount, "genesis account not found"), nil, nil
		}
		genAcc := genAccs[len(genAccs)-1]

		// Find coordinator account
		simAccount, err := findAccount(accs, genAcc.Address)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveGenesisAccount, err.Error()), nil, err
		}
		msg := sample.MsgRequestRemoveAccount(
			genAcc.Address,
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
		chain, found := findChain(ctx, k, false)
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
		valAccs := k.GetAllGenesisValidator(ctx)
		if len(valAccs) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveValidator, "validator not found"), nil, nil
		}
		valAcc := valAccs[len(valAccs)-1]

		// Find coordinator account
		simAccount, err := findAccount(accs, valAcc.Address)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveValidator, err.Error()), nil, err
		}
		msg := sample.MsgRequestRemoveValidator(
			valAcc.Address,
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
		chain, found := findChain(ctx, k, false)
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
		vestAccs := k.GetAllVestingAccount(ctx)
		if len(vestAccs) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveVestingAccount, "vesting account not found"), nil, nil
		}
		vestAcc := vestAccs[len(vestAccs)-1]

		// Find coordinator account
		simAccount, err := findAccount(accs, vestAcc.Address)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRequestRemoveVestingAccount, err.Error()), nil, err
		}
		msg := sample.MsgRequestRemoveAccount(
			vestAcc.Address,
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
		chain, found := findChain(ctx, k, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgTriggerLaunch, "non-triggered chain not found"), nil, nil
		}

		// Find coordinator account
		simAccount, err := findChainCoordinatorAccount(ctx, k, accs, chain.Id)
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
		// Select a request without launch triggered
		chainFound := false
		requests := k.GetAllRequest(ctx)
		var request types.Request
		for _, r := range requests {
			chain, found := k.GetChain(ctx, r.ChainID)
			if found && !chain.LaunchTriggered {
				request = r
				chainFound = true
				break
			}
		}
		if !chainFound {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSettleRequest, "request for non-triggered chain not found"), nil, nil
		}

		// Find coordinator account
		simAccount, err := findChainCoordinatorAccount(ctx, k, accs, request.ChainID)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgSettleRequest, err.Error()), nil, err
		}
		approve := r.Intn(2) == 1
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
		chain, found := findChain(ctx, k, true)
		if !found {
			// No message if no triggered chain
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRevertLaunch, "triggered chain not found"), nil, nil
		}

		// Find coordinator account
		simAccount, err := findChainCoordinatorAccount(ctx, k, accs, chain.Id)
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
