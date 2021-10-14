package campaign

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
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

const (
	weightMsgCreateCampaign    = 25
	weightMsgUpdateTotalSupply = 20
	weightMsgUpdateTotalShares = 20
	weightMsgInitializeMainnet = 5
	weightMsgAddShares         = 20
	weightMsgAddVestingOptions = 20
	weightMsgMintVouchers      = 20
	weightMsgBurnVouchers      = 20
	weightMsgRedeemVouchers    = 20
	weightMsgUnredeemVouchers  = 20
	weightMsgSendVouchers      = 20
)

var (
	// shareDenoms are the denom used for the shares in the simulation
	shareDenoms = []string{"s/foo", "s/bar", "s/toto"}
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	campaignGenesis := sample.CampaignGenesisState()
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&campaignGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return []simtypes.WeightedOperation{
		simulation.NewWeightedOperation(
			weightMsgCreateCampaign,
			SimulateMsgCreateCampaign(am.accountKeeper, am.bankKeeper, am.profileKeeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateTotalSupply,
			SimulateMsgUpdateTotalSupply(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateTotalShares,
			SimulateMsgUpdateTotalShares(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgInitializeMainnet,
			SimulateMsgInitializeMainnet(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgAddShares,
			SimulateMsgAddShares(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgAddVestingOptions,
			SimulateMsgAddVestingOptions(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgMintVouchers,
			SimulateMsgMintVouchers(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgBurnVouchers,
			SimulateMsgBurnVouchers(am.accountKeeper, am.bankKeeper),
		),
		simulation.NewWeightedOperation(
			weightMsgRedeemVouchers,
			SimulateMsgRedeemVouchers(am.accountKeeper, am.bankKeeper),
		),
		simulation.NewWeightedOperation(
			weightMsgUnredeemVouchers,
			SimulateMsgUnredeemVouchers(am.accountKeeper, am.bankKeeper, am.keeper),
		),
		simulation.NewWeightedOperation(
			weightMsgSendVouchers,
			SimulateMsgSendVouchers(am.accountKeeper, am.bankKeeper),
		),
	}
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

// getCoordSimAccount finds an account associated with a coordinator profile from simulation accounts
func getCoordSimAccount(ctx sdk.Context, pk types.ProfileKeeper, accs []simtypes.Account) (simtypes.Account, uint64, bool) {
	for _, acc := range accs {
		coordID, found := pk.CoordinatorIDFromAddress(ctx, acc.Address.String())
		if found {
			return acc, coordID, true
		}
	}
	return simtypes.Account{}, 0, false
}

// getCoordSimAccountWithCampaignID finds an account associated with a coordinator profile from simulation accounts and a campaign created by this coordinator
func getCoordSimAccountWithCampaignID(
	ctx sdk.Context,
	pk types.ProfileKeeper,
	k keeper.Keeper,
	accs []simtypes.Account,
	requireDynamicShares,
	requireNoMainnetInitialized bool,
) (simtypes.Account, uint64, bool) {
	coord, coordID, found := getCoordSimAccount(ctx, pk, accs)
	if !found {
		return coord, 0, false
	}

	// Find a campaign associated to this account
	for _, camp := range k.GetAllCampaign(ctx) {
		if camp.CoordinatorID == coordID &&
			(!requireDynamicShares || camp.DynamicShares) &&
			(!requireNoMainnetInitialized || !camp.MainnetInitialized) {
			return coord, camp.Id, true
		}
	}

	return coord, 0, false
}

// getSharesFromCampaign returns a small portion of shares that can be minted as vouchers or added to an accounts
func getSharesFromCampaign(r *rand.Rand, ctx sdk.Context, k keeper.Keeper, campID uint64) (types.Shares, bool) {
	camp, found := k.GetCampaign(ctx, campID)
	if !found {
		return types.EmptyShares(), false
	}

	// store current values in map
	totalShares := make(map[string]int64)
	allocatedShares := make(map[string]int64)
	for _, total := range camp.TotalShares {
		totalShares[total.Denom] = total.Amount.Int64()
	}
	for _, allocated := range camp.AllocatedShares {
		allocatedShares[allocated.Denom] = allocated.Amount.Int64()
	}

	var shares sdk.Coins
	for _, share := range shareDenoms {
		total := totalShares[share]
		if total == 0 {
			total = types.DefaultTotalShareNumber
		}
		remaining := total - allocatedShares[share]
		if remaining == 0 {
			continue
		}

		shareNb := r.Int63n(5000) + 1
		if shareNb > remaining {
			shareNb = remaining
		}
		shares = append(shares, sdk.NewCoin(share, sdk.NewInt(shareNb)))
	}

	// No shares can be distributed
	if len(shares) == 0 {
		return types.EmptyShares(), false
	}
	shares = shares.Sort()
	return types.Shares(shares), true
}

// getAccountWithVouchers returns an account that has vouchers for a campaign
func getAccountWithVouchers(
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
		found = true
		accountAddr = addr
		return true
	})

	// No account has vouchers
	if !found {
		return
	}

	// Fetch all the vouchers of the campaign owned by the account
	bk.IterateAccountBalances(ctx, accountAddr, func(coin sdk.Coin) bool {
		coinCampID, err := types.VoucherCampaign(coin.Denom)
		if err == nil && coinCampID == campID {
			coins = append(coins, coin)
		}
		return false
	})
	coins = coins.Sort()

	// Find the sim account
	for _, acc := range accs {
		if found = acc.Address.Equals(accountAddr); found {
			account = acc
			return
		}
	}
	return
}

// getAccountWithShares returns an account that contains allocated shares with its associated campaign
func getAccountWithShares(
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
		simAccount, _, found := getCoordSimAccount(ctx, pk, accs)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateCampaign, "skip campaign creation"), nil, nil
		}

		dynamicShares := r.Intn(100) > 80

		msg := types.NewMsgCreateCampaign(
			simAccount.Address.String(),
			sample.CampaignName(),
			sample.Coins(),
			dynamicShares,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgUpdateTotalSupply simulates a MsgUpdateTotalSupply message
func SimulateMsgUpdateTotalSupply(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := getCoordSimAccountWithCampaignID(ctx, pk, k, accs, false, true)
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
		simAccount, campID, found := getCoordSimAccountWithCampaignID(ctx, pk, k, accs, true, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateTotalShares, "skip update total shares"), nil, nil
		}

		camp, _ := k.GetCampaign(ctx, campID)

		// Update all total shares to a number between allocated and current+100000
		totalShares := make(map[string]int64)
		allocatedShares := make(map[string]int64)
		for _, total := range camp.TotalShares {
			totalShares[total.Denom] = total.Amount.Int64()
		}
		for _, allocated := range camp.AllocatedShares {
			allocatedShares[allocated.Denom] = allocated.Amount.Int64()
		}

		// Defined the value to update
		var newTotalShares sdk.Coins
		for _, share := range shareDenoms {
			currentTotal := totalShares[share]
			if currentTotal == 0 {
				currentTotal = types.DefaultTotalShareNumber
			}
			newTotal := r.Int63n(currentTotal+types.DefaultTotalShareNumber) + allocatedShares[share]
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
		simAccount, campID, found := getCoordSimAccountWithCampaignID(ctx, pk, k, accs, false, true)
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
		simAccount, campID, found := getCoordSimAccountWithCampaignID(ctx, pk, k, accs, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddShares, "skip add shares"), nil, nil
		}

		shares, getShares := getSharesFromCampaign(r, ctx, k, campID)
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
		simAccount, campID, found := getCoordSimAccountWithCampaignID(ctx, pk, k, accs, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddVestingOptions, "skip add vesting options"), nil, nil
		}

		shares, getShares := getSharesFromCampaign(r, ctx, k, campID)
		if !getShares {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddVestingOptions, "skip add vesting options"), nil, nil
		}

		// Select a random account to give vesting options
		accountNb := r.Intn(len(accs))

		msg := types.NewMsgAddVestingOptions(
			campID,
			simAccount.Address.String(),
			accs[accountNb].Address.String(),
			types.EmptyShares(),
			*types.NewShareDelayedVesting(shares, time.Now().Unix()),
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins())
	}
}

// SimulateMsgMintVouchers simulates a MsgMintVouchers message
func SimulateMsgMintVouchers(ak types.AccountKeeper, bk types.BankKeeper, pk types.ProfileKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, campID, found := getCoordSimAccountWithCampaignID(ctx, pk, k, accs, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgMintVouchers, "skip mint vouchers"), nil, nil
		}

		shares, getShares := getSharesFromCampaign(r, ctx, k, campID)
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
		campID, simAccount, vouchers, found := getAccountWithVouchers(ctx, bk, accs)
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
		campID, simAccount, vouchers, found := getAccountWithVouchers(ctx, bk, accs)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRedeemVouchers, "skip redeem vouchers"), nil, nil
		}

		// Select a random account to redeem vouchers into
		accountNb := r.Intn(len(accs))

		msg := types.NewMsgRedeemVouchers(
			simAccount.Address.String(),
			campID,
			accs[accountNb].Address.String(),
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
		campID, simAccount, shares, found := getAccountWithShares(r, ctx, k, accs)
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
func SimulateMsgSendVouchers(ak types.AccountKeeper, bk types.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, simAccount, vouchers, found := getAccountWithVouchers(ctx, bk, accs)
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
