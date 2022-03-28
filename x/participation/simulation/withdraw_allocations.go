package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	fundraisingkeeper "github.com/tendermint/fundraising/x/fundraising/keeper"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	"github.com/tendermint/spn/x/participation/keeper"
	"github.com/tendermint/spn/x/participation/types"
)

// RandomAuctionWithdrawEnabled returns random auction where used allocations can be withdrawn at blockTime
func RandomAuctionWithdrawEnabled(
	ctx sdk.Context,
	r *rand.Rand,
	fk fundraisingkeeper.Keeper,
	k keeper.Keeper,
) (auction fundraisingtypes.AuctionI, found bool) {
	auctions := fk.GetAuctions(ctx)
	withdrawalDelay := k.WithdrawalDelay(ctx)
	if len(auctions) == 0 {
		return auction, false
	}

	r.Shuffle(len(auctions), func(i, j int) {
		auctions[i], auctions[j] = auctions[j], auctions[i]
	})

	for _, a := range auctions {
		if ctx.BlockTime().After(auction.GetStartTime().Add(withdrawalDelay)) {
			return a, true
		}
	}

	return auction, false
}

// RandomAccWithAuctionUsedAllocationsNotWithdrawn returns random account that has used allocations for the given
// auction that have not yet been withdrawn
func RandomAccWithAuctionUsedAllocationsNotWithdrawn(
	ctx sdk.Context,
	r *rand.Rand,
	k keeper.Keeper,
	accs []simtypes.Account,
	auctionID uint64,
) (simtypes.Account, bool) {
	// Randomize the set
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})

	// account must have used allocations for this auction that have not yet been withdrawn
	for _, acc := range accs {
		usedAllocations, found := k.GetAuctionUsedAllocations(ctx, acc.Address.String(), auctionID)
		if !found || usedAllocations.Withdrawn {
			continue
		}

		return acc, true
	}

	return simtypes.Account{}, false
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
