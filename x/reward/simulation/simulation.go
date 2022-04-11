package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

const (
	ProbabilityCreateRewardPool = 70
	ProbabilityCloseRewardPool  = 70

	ActionCreate = iota
	ActionEdit
	ActionClose
)

func SimulateMsgSetRewards(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgSetRewards{}
		createRewardPool := r.Int63n(100) < ProbabilityCreateRewardPool
		closeRewardPool := r.Int63n(100) < ProbabilityCloseRewardPool
		checkBalance := true

		// choose action to be taken
		action := ActionCreate
		if !createRewardPool && closeRewardPool {
			action = ActionClose
			checkBalance = false
		} else if !createRewardPool && !closeRewardPool {
			action = ActionEdit
		}

		wantCoin := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(r.Int63n(1_000_000))))
		chain, found := FindRandomChainWithCoordBalance(r, ctx, k, bk, createRewardPool, checkBalance, wantCoin)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no viable chain to be found"), nil, nil
		}

		coordinator, found := k.GetProfileKeeper().GetCoordinator(ctx, chain.CoordinatorID)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "chain does not have a coordinator"), nil, nil
		}

		coordAccAddr, err := sdk.AccAddressFromBech32(coordinator.Address)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, err
		}

		simAccount, found := simtypes.FindAccount(accs, coordAccAddr)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "simulation account not found for chain coordinator"), nil, nil
		}

		// initialize basic message
		msg = &types.MsgSetRewards{
			Provider: simAccount.Address.String(),
			LaunchID: chain.LaunchID,
		}

		// set message based on action to be taken
		switch action {
		case ActionCreate:
			msg.LastRewardHeight = ctx.BlockHeight() + r.Int63n(1000)
			msg.Coins = wantCoin
		case ActionEdit:
			pool, _ := k.GetRewardPool(ctx, chain.LaunchID)
			msg.LastRewardHeight = pool.LastRewardHeight + r.Int63n(1000)
			msg.Coins = wantCoin
		case ActionClose:
			msg.LastRewardHeight = 0
			msg.Coins = sdk.NewCoins()
		}

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
			CoinsSpentInMsg: wantCoin,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
