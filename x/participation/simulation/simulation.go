package simulation

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"
	fundraisingkeeper "github.com/tendermint/fundraising/x/fundraising/keeper"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	"github.com/tendermint/spn/app/simutil"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/testutil/simulation"
	"github.com/tendermint/spn/x/participation/keeper"
	"github.com/tendermint/spn/x/participation/types"
)

func SimulateMsgParticipate(
	ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper,
	fk fundraisingkeeper.Keeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgParticipate{}
		auction, found := RandomAuctionParticipationEnabled(ctx, r, fk, k)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no valid auction found"), nil, nil
		}

		tierList := k.ParticipationTierList(ctx)
		tier, found := RandomTierFromList(r, tierList)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no valid tiers"), nil, nil
		}

		simAccount, _, found := RandomAccWithAvailableAllocations(ctx, r, k, accs, tier.RequiredAllocations, auction.GetId())
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no account with allocations"), nil, nil
		}

		msg = types.NewMsgParticipate(
			simAccount.Address.String(),
			auction.GetId(),
			tier.TierID,
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

func SimulateCreateAuction(
	ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper,
	fk fundraisingkeeper.Keeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &fundraisingtypes.MsgCreateFixedPriceAuction{}

		// fundraising simulation params must be set
		// since they are not initially set
		params := fundraisingtypes.DefaultParams()
		params.AuctionCreationFee = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(100_000)))
		fk.SetParams(ctx, params)
		fee := params.AuctionCreationFee

		largestTier, found := FindLargestMaxBid(k.ParticipationTierList(ctx))
		if !found {
			return simtypes.OperationMsg{},
				nil,
				fmt.Errorf("no tier in list")
		}
		// create a selling coin that at least covers all tiers in the simulation param
		requireAmt := sdk.NewCoin(simutil.AuctionCoinDenom, largestTier.Benefits.MaxBidAmount)
		sellCoin := sample.Coin(r)
		sellCoin.Denom = simutil.AuctionCoinDenom
		sellCoin = sellCoin.Add(requireAmt)

		desiredCoins := fee.Add(sellCoin)
		simAccount, _, found := RandomAccWithBalance(ctx, r, bk, accs, desiredCoins)
		if !found {
			return simtypes.NoOpMsg(
					types.ModuleName,
					msg.Type(),
					"no account with balance found"),
				nil,
				nil
		}

		startTime := ctx.BlockTime().Add(time.Hour * 24)
		endTime := startTime.Add(time.Hour * 24 * 7)
		msg = sample.MsgCreateFixedAuction(r, simAccount.Address.String(), sellCoin, startTime, endTime)

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
			ModuleName:      fundraisingtypes.ModuleName,
			CoinsSpentInMsg: desiredCoins,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
	}
}

func SimulateMsgWithdrawAllocations(
	ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper,
	fk fundraisingkeeper.Keeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgWithdrawAllocations{}
		auction, found := RandomAuctionWithdrawEnabled(ctx, r, fk, k)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no valid auction found"), nil, nil
		}

		simAccount, found := RandomAccWithAuctionUsedAllocationsNotWithdrawn(ctx, r, k, accs, auction.GetId())
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no account with used allocations found"), nil, nil
		}

		msg = types.NewMsgWithdrawAllocations(
			simAccount.Address.String(),
			auction.GetId(),
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

func SimulateMsgCancelAuction(
	ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper,
	fk fundraisingkeeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var simAccount simtypes.Account
		msg := &fundraisingtypes.MsgCancelAuction{}
		auction, found := RandomAuctionStandby(ctx, r, fk)
		if !found {
			return simtypes.NoOpMsg(fundraisingtypes.ModuleName, msg.Type(), "no valid auction found"), nil, nil
		}

		// find account of auctioneer
		found = false
		for _, acc := range accs {
			if acc.Address.Equals(auction.GetAuctioneer()) {
				simAccount = acc
				found = true
				break
			}
		}
		if !found {
			// return error, this should never happen
			return simtypes.OperationMsg{}, nil, errors.New("auctioneer not found within provided accounts")
		}

		msg = fundraisingtypes.NewMsgCancelAuction(simAccount.Address.String(), auction.GetId())

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
			ModuleName:      fundraisingtypes.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
	}
}
