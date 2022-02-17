package simulation

import (
	"math/rand"
	"time"

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

var (
	// ShareDenoms are the denom used for the shares in the simulation
	ShareDenoms = []string{"s/foo", "s/bar", "s/toto"}
)

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

// GetCoordSimAccount finds an account associated with a coordinator profile from simulation accounts
func GetCoordSimAccount(
	r *rand.Rand,
	ctx sdk.Context,
	pk types.ProfileKeeper,
	accs []simtypes.Account,
) (simtypes.Account, uint64, bool) {
	// Choose a random coordinator
	coords := pk.GetAllCoordinator(ctx)
	coordNb := len(coords)
	if coordNb == 0 {
		return simtypes.Account{}, 0, false
	}
	coord := coords[r.Intn(coordNb)]

	// Find the account linked to this address
	for _, acc := range accs {
		if acc.Address.String() == coord.Address && coord.Active {
			return acc, coord.CoordinatorID, true
		}
	}
	return simtypes.Account{}, 0, false
}

// GetCoordSimAccountWithCampaignID finds an account associated with a coordinator profile from simulation accounts and a campaign created by this coordinator
func GetCoordSimAccountWithCampaignID(
	r *rand.Rand,
	ctx sdk.Context,
	pk types.ProfileKeeper,
	k keeper.Keeper,
	accs []simtypes.Account,
	requireDynamicShares,
	requireNoMainnetInitialized bool,
) (simtypes.Account, uint64, bool) {
	campaigns := k.GetAllCampaign(ctx)
	campNb := len(campaigns)
	if campNb == 0 {
		return simtypes.Account{}, 0, false
	}

	var camp types.Campaign
	if requireDynamicShares || requireNoMainnetInitialized {
		// If a criteria is required for the campaign, we simply fetch the first on that satisfies the criteria
		var campFound bool
		for _, campaign := range campaigns {
			if (!requireDynamicShares || campaign.DynamicShares) &&
				(!requireNoMainnetInitialized || !campaign.MainnetInitialized) {
				camp = campaign
				campFound = true
				break
			}
		}
		if !campFound {
			return simtypes.Account{}, 0, false
		}
	} else {
		// No criteria, choose a random campaign
		camp = campaigns[r.Intn(campNb)]
	}

	// Find the sim account of the campaign coordinator
	coord, found := pk.GetCoordinator(ctx, camp.CoordinatorID)
	if !found {
		return simtypes.Account{}, 0, false
	}
	for _, acc := range accs {
		if acc.Address.String() == coord.Address && coord.Active {
			return acc, camp.CampaignID, true
		}
	}

	return simtypes.Account{}, 0, false
}

// GetSharesFromCampaign returns a small portion of shares that can be minted as vouchers or added to an accounts
func GetSharesFromCampaign(r *rand.Rand, ctx sdk.Context, k keeper.Keeper, campID uint64) (types.Shares, bool) {
	camp, found := k.GetCampaign(ctx, campID)
	if !found {
		return types.EmptyShares(), false
	}

	var shares sdk.Coins
	for _, share := range ShareDenoms {
		total := camp.TotalShares.AmountOf(share)
		if total == 0 {
			total = types.DefaultTotalShareNumber
		}
		remaining := total - camp.AllocatedShares.AmountOf(share)
		if remaining == 0 {
			continue
		}

		shareNb := r.Int63n(5000) + 10
		if shareNb > remaining {
			shareNb = remaining
		}
		shares = append(shares, sdk.NewCoin(share, sdk.NewInt(shareNb)))
	}

	// No shares can be distributed
	if shares.Empty() {
		return types.EmptyShares(), false
	}
	shares = shares.Sort()
	return types.Shares(shares), true
}

// GetAccountWithVouchers returns an account that has vouchers for a campaign
func GetAccountWithVouchers(
	ctx sdk.Context,
	bk types.BankKeeper,
	accs []simtypes.Account,
) (campID uint64, account simtypes.Account, coins sdk.Coins, found bool) {
	var err error
	var accountAddr sdk.AccAddress

	// Parse all account balances and find one with vouchers
	bk.IterateAllBalances(ctx, func(addr sdk.AccAddress, coin sdk.Coin) bool {
		campID, err = types.VoucherCampaign(coin.Denom)
		if err != nil {
			return false
		}

		// Look for accounts with at least 10 vouchers
		if coin.Amount.Int64() < 10 {
			return false
		}

		found = true
		accountAddr = addr
		return true
	})

	// No account has vouchers
	if !found {
		return 0, account, coins, false
	}

	// Fetch all the vouchers of the campaign owned by the account
	bk.IterateAccountBalances(ctx, accountAddr, func(coin sdk.Coin) bool {
		coinCampID, err := types.VoucherCampaign(coin.Denom)
		if err == nil && coinCampID == campID {

			// Get a portion of the balance
			// If the balance is 1, we don't include it in the vouchers
			// There is a issue: insufficient fees that can occur when the whole balance for a voucher is sent
			// TODO: Investigate this issue
			if coin.Amount.Int64() > 1 {
				coin.Amount = coin.Amount.Quo(sdk.NewInt(2))
				coins = append(coins, coin)
			}
		}
		return false
	})
	if coins.Empty() {
		return 0, account, coins, false
	}

	coins = coins.Sort()

	// Find the sim account
	for _, acc := range accs {
		if found = acc.Address.Equals(accountAddr); found {
			return campID, acc, coins, true
		}
	}
	return 0, account, coins, false
}

// GetAccountWithShares returns an account that contains allocated shares with its associated campaign
func GetAccountWithShares(
	r *rand.Rand,
	ctx sdk.Context,
	k keeper.Keeper,
	accs []simtypes.Account,
) (uint64, simtypes.Account, types.Shares, bool) {
	mainnetAccounts := k.GetAllMainnetAccount(ctx)
	nb := len(mainnetAccounts)

	// No account have been created yet
	if nb == 0 {
		return 0, simtypes.Account{}, nil, false
	}

	mainnetAccount := mainnetAccounts[r.Intn(nb)]

	// Find the associated sim acocunt
	for _, acc := range accs {
		if acc.Address.String() == mainnetAccount.Address {
			return mainnetAccount.CampaignID, acc, mainnetAccount.Shares, true
		}
	}
	return 0, simtypes.Account{}, nil, false
}

// SimulateMsgCreateCampaign simulates a MsgCreateCampaign message
func SimulateMsgCreateCampaign(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _, found := GetCoordSimAccount(r, ctx, pk, accs)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateCampaign, "skip campaign creation"), nil, nil
		}

		msg := types.NewMsgCreateCampaign(
			simAccount.Address.String(),
			sample.CampaignName(),
			sample.Coins(),
			sample.Metadata(20),
		)

		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// TODO add SimulateMsgEditCampaign

// SimulateMsgUpdateTotalSupply simulates a MsgUpdateTotalSupply message
func SimulateMsgUpdateTotalSupply(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, false, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateTotalSupply, "skip update total supply"), nil, nil
		}

		msg := types.NewMsgUpdateTotalSupply(
			simAccount.Address.String(),
			campID,
			sample.Coins(),
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgUpdateTotalShares simulates a MsgUpdateTotalShares message
func SimulateMsgUpdateTotalShares(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, true, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateTotalShares, "skip update total shares"), nil, nil
		}

		camp, _ := k.GetCampaign(ctx, campID)

		// Defined the value to update
		var newTotalShares sdk.Coins
		for _, share := range ShareDenoms {
			currentTotal := camp.TotalShares.AmountOf(share)
			if currentTotal == 0 {
				currentTotal = types.DefaultTotalShareNumber
			}
			allocatedShare := camp.AllocatedShares.AmountOf(share)
			newTotal := r.Int63n(currentTotal+types.DefaultTotalShareNumber) + allocatedShare
			newTotalShares = append(newTotalShares, sdk.NewCoin(share, sdk.NewInt(newTotal)))
		}
		newTotalShares = newTotalShares.Sort()

		msg := types.NewMsgUpdateTotalShares(
			simAccount.Address.String(),
			campID,
			types.Shares(newTotalShares),
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgInitializeMainnet simulates a MsgInitializeMainnet message
func SimulateMsgInitializeMainnet(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, false, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgInitializeMainnet, "skip initliaze mainnet"), nil, nil
		}

		msg := types.NewMsgInitializeMainnet(
			simAccount.Address.String(),
			campID,
			sample.String(50),
			sample.String(32),
			sample.GenesisChainID(),
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgAddShares simulates a MsgAddShares message
func SimulateMsgAddShares(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, false, false)
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
func SimulateMsgAddVestingOptions(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := GetCoordSimAccountWithCampaignID(r, ctx, pk, k, accs, false, false)
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
			*types.NewShareDelayedVesting(shares, shares, time.Now().Unix()),
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgMintVouchers simulates a MsgMintVouchers message
func SimulateMsgMintVouchers(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
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
func SimulateMsgBurnVouchers(ak types.AccountKeeper, bk types.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		campID, simAccount, vouchers, found := GetAccountWithVouchers(ctx, bk, accs)
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
func SimulateMsgRedeemVouchers(ak types.AccountKeeper, bk types.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		campID, simAccount, vouchers, found := GetAccountWithVouchers(ctx, bk, accs)
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
func SimulateMsgUnredeemVouchers(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Find a random account from a random campaign that contains shares
		campID, simAccount, shares, found := GetAccountWithShares(r, ctx, k, accs)
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
// TODO: This message constantly fails in simulation log, investigate why
func SimulateMsgSendVouchers(ak types.AccountKeeper, bk types.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, simAccount, vouchers, found := GetAccountWithVouchers(ctx, bk, accs)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, banktypes.TypeMsgSend, "skip send vouchers"), nil, nil
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
