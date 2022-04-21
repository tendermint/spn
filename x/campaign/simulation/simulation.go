package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
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
		CoinsSpentInMsg: coinsSpent,
	}
	return simulation.GenAndDeliverTxWithRandFees(txCtx)
}

// SimulateMsgCreateCampaign simulates a MsgCreateCampaign message
func SimulateMsgCreateCampaign(
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
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateCampaign, "skip campaign creation"), nil, nil
		}

		creationFee := k.CampaignCreationFee(ctx)

		msg := types.NewMsgCreateCampaign(
			simAccount.Address.String(),
			sample.CampaignName(r),
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
		)

		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, creationFee)
	}
}

// SimulateMsgEditCampaign simulates a MsgEditCampaign message
func SimulateMsgEditCampaign(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgEditCampaign{}

		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip edit campaign"), nil, nil
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
			newName = sample.CampaignName(r)
		}
		if metadata {
			newMetadata = sample.Metadata(r, 20)
		}

		msg = types.NewMsgEditCampaign(
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
		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, true, true)
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
		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, true, true)
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

// SimulateMsgAddShares simulates a MsgAddShares message
func SimulateMsgAddShares(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, false, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddShares, "skip add shares"), nil, nil
		}

		shares, getShares := GetSharesFromCampaign(r, ctx, k, campID)
		if !getShares {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddShares, "skip add shares"), nil, nil
		}

		// Select a random account to give shares
		accountNb := r.Intn(len(accs))

		msg := types.NewMsgAddShares(
			campID,
			simAccount.Address.String(),
			accs[accountNb].Address.String(),
			shares,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgAddVestingOptions simulates a MsgAddVestingOptions message
func SimulateMsgAddVestingOptions(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, false, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddVestingOptions, "skip add vesting options"), nil, nil
		}

		shares, getShares := GetSharesFromCampaign(r, ctx, k, campID)
		if !getShares {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddVestingOptions, "skip add vesting options"), nil, nil
		}

		// Select a random account to give vesting options
		accountNb := r.Intn(len(accs))

		msg := types.NewMsgAddVestingOptions(
			campID,
			simAccount.Address.String(),
			accs[accountNb].Address.String(),
			*types.NewShareDelayedVesting(shares, shares, int64(sample.Duration(r))),
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
		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgMintVouchers, "skip mint vouchers"), nil, nil
		}

		shares, getShares := GetSharesFromCampaign(r, ctx, k, campID)
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
		// Find a random account from a random campaign that contains shares
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

// SimulateMsgSendVouchers simulates a Msg message from the bank module with vouchers
func SimulateMsgSendVouchers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, simAccount, vouchers, found := GetAccountWithVouchers(r, ctx, bk, k, accs, false)
		if !found {
			return simtypes.NoOpMsg(banktypes.ModuleName, banktypes.TypeMsgSend, "skip send vouchers"), nil, nil
		}

		// Select a random receiver account
		accountNb := r.Intn(len(accs))
		if accs[accountNb].Address.Equals(simAccount.Address) {
			if accountNb == 0 {
				accountNb++
			} else {
				accountNb--
			}
		}

		msg := banktypes.NewMsgSend(
			simAccount.Address,
			accs[accountNb].Address,
			vouchers,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, vouchers)
	}
}
