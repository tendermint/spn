package simulation

import (
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	fundraisingkeeper "github.com/tendermint/fundraising/x/fundraising/keeper"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"
	"github.com/tendermint/spn/testutil/sample"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/tendermint/spn/x/participation/keeper"
	"github.com/tendermint/spn/x/participation/types"
)

func RandomAccWithBalance(ctx sdk.Context, r *rand.Rand,
	bk bankkeeper.Keeper,
	accs []simtypes.Account,
	desired sdk.Coins,
) (account simtypes.Account, coins sdk.Coins, found bool) {
	// Randomize the set
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})

	for _, acc := range accs {
		balances := bk.GetAllBalances(ctx, acc.Address)
		if balances.IsAllGT(desired) {
			return acc, balances, true
		}
	}

	return simtypes.Account{}, sdk.NewCoins(), false
}

func RandomAuction(ctx sdk.Context, r *rand.Rand, fk fundraisingkeeper.Keeper) (auction fundraisingtypes.AuctionI, found bool) {
	auctions := fk.GetAuctions(ctx)
	if len(auctions) == 0 {
		return auction, false
	}

	index := r.Intn(len(auctions))
	return auctions[index], true
}

func SimulateMsgParticipate(
	ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper,
	fk fundraisingkeeper.Keeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgParticipate{}
		auction, found := RandomAuction(ctx, r, fk)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no valid auction found"), nil, nil
		}

		numTiers := len(k.GetParams(ctx).ParticipationTierList)
		if numTiers == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no valid tiers"), nil, nil
		}

		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg = &types.MsgParticipate{
			Participant: simAccount.Address.String(),
			AuctionID:   auction.GetId(),
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
			ModuleName:      fundraisingtypes.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateCreateAuction(
	ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper,
	fk fundraisingkeeper.Keeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// fundraising simulation params must be set
		params := fundraisingtypes.DefaultParams()
		fk.SetParams(ctx, params)
		fee := params.AuctionCreationFee
		sellCoin := sample.Coin()

		simAccount, _, found := RandomAccWithBalance(ctx, r, bk, accs, fee)
		if !found {
			return simtypes.NoOpMsg(
					types.ModuleName,
					fundraisingtypes.MsgCreateFixedPriceAuction{}.Type(),
					"no account with balance found"),
				nil,
				nil
		}

		startTime := ctx.BlockTime().Add(time.Hour * 24)
		endTime := startTime.Add(time.Hour * 24 * 7)
		msg := sample.MsgCreateFixedAuction(simAccount.Address.String(), sellCoin, startTime, endTime)

		mintAmt := sdk.NewCoins(msg.SellingCoin)
		// must mint and send new coins to auctioneer
		err := bk.MintCoins(ctx, minttypes.ModuleName, mintAmt)
		if err != nil {
			return simtypes.NoOpMsg(
					types.ModuleName,
					fundraisingtypes.MsgCreateFixedPriceAuction{}.Type(),
					"error setting up balance"),
				nil,
				nil
		}
		err = bk.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, simAccount.Address, mintAmt)
		if err != nil {
			return simtypes.NoOpMsg(
					types.ModuleName,
					fundraisingtypes.MsgCreateFixedPriceAuction{}.Type(),
					"error setting up balance"),
				nil,
				nil
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
			ModuleName:      fundraisingtypes.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		// choose custom fee that only uses the default bond denom
		// otherwise the custom sellingCoin denom could be chosen
		customFee, err := simtypes.RandomFees(txCtx.R, txCtx.Context, fee)
		if err != nil {
			return simtypes.NoOpMsg(
					types.ModuleName,
					fundraisingtypes.MsgCreateFixedPriceAuction{}.Type(),
					"error setting up custom fee"),
				nil,
				nil
		}

		return simulation.GenAndDeliverTx(txCtx, customFee)
	}
}

func SimulateMsgWithdrawAllocations(
	_ authkeeper.AccountKeeper,
	_ bankkeeper.Keeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgWithdrawAllocations{
			Participant: simAccount.Address.String(),
		}

		// TODO: Handling the WithdrawAllocations simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "WithdrawAllocations simulation not implemented"), nil, nil
	}
}
