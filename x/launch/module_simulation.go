package launch

import (
	"fmt"
	"math/rand"

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
	defaultWeightMsgRequestRemoveGenesisAccount int = 15
	defaultWeightMsgRequestAddVestingAccount    int = 50
	defaultWeightMsgRequestRemoveVestingAccount int = 15
	defaultWeightMsgRequestAddValidator         int = 50
	defaultWeightMsgRequestRemoveValidator      int = 15
	defaultWeightMsgSettleRequest               int = 50
	defaultWeightMsgTriggerLaunch               int = 15
	defaultWeightMsgRevertLaunch                int = 5

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
			weightMsgSettleRequest,
			SimulateMsgSettleRequest(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRevertLaunch,
			SimulateMsgRevertLaunch(am.accountKeeper, am.bankKeeper, am.keeper),
		),
	}
}

// IsLaunchTriggeredChain check if chain is launch triggered
func IsLaunchTriggeredChain(ctx sdk.Context, k keeper.Keeper, chainID uint64) bool {
	chain, found := k.GetChain(ctx, chainID)
	if !found {
		return false
	}
	return chain.LaunchTriggered
}

// FindChain find a chain
func FindChain(ctx sdk.Context, k keeper.Keeper, launchTriggered bool) (types.Chain, bool) {
	found := false
	chains := k.GetAllChain(ctx)
	var chain types.Chain
	for _, c := range chains {
		if c.LaunchTriggered != launchTriggered {
			continue
		}
		// check if the coordinator is still in the store
		_, found = k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, c.CoordinatorID)
		if !found {
			continue
		}
		chain = c
		break
	}
	return chain, found
}

// FindChainCoordinatorAccount find coordinator account by chain id
func FindChainCoordinatorAccount(ctx sdk.Context, k keeper.Keeper, accs []simtypes.Account, chainID uint64) (simtypes.Account, error) {
	chain, found := k.GetChain(ctx, chainID)
	if !found {
		// No message if no coordinator address
		return simtypes.Account{}, fmt.Errorf("chain %d not found", chainID)
	}
	address, found := k.GetProfileKeeper().GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return simtypes.Account{}, fmt.Errorf("coordinator %d not found", chain.CoordinatorID)
	}
	return FindAccount(accs, address)
}

// FindAccount find account by string hex address
func FindAccount(accs []simtypes.Account, address string) (simtypes.Account, error) {
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
