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
	"github.com/tendermint/spn/x/project/keeper"
	"github.com/tendermint/spn/x/project/types"
)

// TypedMsg extends sdk.Msg with Type method
type TypedMsg interface {
	sdk.Msg
	Type() string
}

// deliverSimTx delivers the tx for simulation from the provided message
func deliverSimTx(
	r *rand.Rand,
	app *baseapp.BaseApp,
	ctx sdk.Context,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	simAccount simtypes.Account,
	msg TypedMsg,
	coinsSpent sdk.Coins,
) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
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
		CoinsSpentInMsg: coinsSpent,
	}
	return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
}

// SimulateMsgCreateProject simulates a MsgCreateProject message
func SimulateMsgCreateProject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _, found := GetCoordSimAccount(r, ctx, pk, accs)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateProject, "skip project creation"), nil, nil
		}

		creationFee := k.ProjectCreationFee(ctx)

		msg := types.NewMsgCreateProject(
			simAccount.Address.String(),
			sample.ProjectName(r),
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
		)

		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, creationFee)
	}
}

// SimulateMsgEditProject simulates a MsgEditProject message
func SimulateMsgEditProject(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgEditProject{}

		simAccount, campID, found := GetCoordSimAccountWithProjectID(r, ctx, pk, k, accs, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip edit project"), nil, nil
		}

		var newName string
		var newMetadata []byte

		name := r.Intn(100) < 50
		metadata := r.Intn(100) < 50
		// ensure there is always a value to edit
		if !name && !metadata {
			metadata = true
		}

		if name {
			newName = sample.ProjectName(r)
		}
		if metadata {
			newMetadata = sample.Metadata(r, 20)
		}

		msg = types.NewMsgEditProject(
			simAccount.Address.String(),
			campID,
			newName,
			newMetadata,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgUpdateTotalSupply simulates a MsgUpdateTotalSupply message
func SimulateMsgUpdateTotalSupply(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithProjectID(r, ctx, pk, k, accs, true, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateTotalSupply, "skip update total supply"), nil, nil
		}

		msg := types.NewMsgUpdateTotalSupply(
			simAccount.Address.String(),
			campID,
			sample.TotalSupply(r),
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgInitializeMainnet simulates a MsgInitializeMainnet message
func SimulateMsgInitializeMainnet(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithProjectID(r, ctx, pk, k, accs, true, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgInitializeMainnet, "skip initliaze mainnet"), nil, nil
		}

		msg := types.NewMsgInitializeMainnet(
			simAccount.Address.String(),
			campID,
			sample.String(r, 50),
			sample.String(r, 32),
			sample.GenesisChainID(r),
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgUpdateSpecialAllocations simulates a MsgUpdateSpecialAllocations message
func SimulateMsgUpdateSpecialAllocations(ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithProjectID(r, ctx, pk, k, accs, false, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateSpecialAllocations, "skip update special allocations"), nil, nil
		}

		// get shares for both genesis distribution and claimable airdrop
		genesisDistribution, getShares := GetSharesFromProject(r, ctx, k, campID)
		if !getShares {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateSpecialAllocations, "skip update special allocations"), nil, nil
		}
		claimableAirdrop, getShares := GetSharesFromProject(r, ctx, k, campID)
		if !getShares {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateSpecialAllocations, "skip update special allocations"), nil, nil
		}

		// GetSharesFromProject returns a number of available shares for a project
		// potentially genesisDistribution + claimableAirdrop can overflow the available shares
		// we divide by two all amounts to avoid overflowing available shares
		for i, s := range genesisDistribution {
			genesisDistribution[i].Amount = s.Amount.QuoRaw(2)
		}
		for i, s := range claimableAirdrop {
			claimableAirdrop[i].Amount = s.Amount.QuoRaw(2)
		}

		msg := types.NewMsgUpdateSpecialAllocations(
			simAccount.Address.String(),
			campID,
			types.NewSpecialAllocations(genesisDistribution, claimableAirdrop),
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgMintVouchers simulates a MsgMintVouchers message
func SimulateMsgMintVouchers(ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithProjectID(r, ctx, pk, k, accs, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgMintVouchers, "skip mint vouchers"), nil, nil
		}

		shares, getShares := GetSharesFromProject(r, ctx, k, campID)
		if !getShares {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgMintVouchers, "skip mint vouchers"), nil, nil
		}

		msg := types.NewMsgMintVouchers(
			simAccount.Address.String(),
			campID,
			shares,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgBurnVouchers simulates a MsgBurnVouchers message
func SimulateMsgBurnVouchers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		campID, simAccount, vouchers, found := GetAccountWithVouchers(r, ctx, bk, k, accs, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBurnVouchers, "skip burn vouchers"), nil, nil
		}

		msg := types.NewMsgBurnVouchers(
			simAccount.Address.String(),
			campID,
			vouchers,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, vouchers)
	}
}

// SimulateMsgRedeemVouchers simulates a MsgRedeemVouchers message
func SimulateMsgRedeemVouchers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		campID, simAccount, vouchers, found := GetAccountWithVouchers(r, ctx, bk, k, accs, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRedeemVouchers, "skip redeem vouchers"), nil, nil
		}

		// Select a random account to redeem vouchers into
		accountNb := r.Intn(len(accs))

		msg := types.NewMsgRedeemVouchers(
			simAccount.Address.String(),
			accs[accountNb].Address.String(),
			campID,
			vouchers,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, vouchers)
	}
}

// SimulateMsgUnredeemVouchers simulates a MsgUnredeemVouchers message
func SimulateMsgUnredeemVouchers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Find a random account from a random project that contains shares
		campID, simAccount, shares, found := GetAccountWithShares(r, ctx, k, accs, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnredeemVouchers, "skip unredeem vouchers"), nil, nil
		}

		msg := types.NewMsgUnredeemVouchers(
			simAccount.Address.String(),
			campID,
			shares,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}
