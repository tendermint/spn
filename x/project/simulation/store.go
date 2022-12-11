package simulation

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/tendermint/spn/x/project/keeper"
	"github.com/tendermint/spn/x/project/types"
)

// ShareDenoms are the denom used for the shares in the simulation
var ShareDenoms = []string{"s/foo", "s/bar", "s/toto"}

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

// GetCoordSimAccountWithProjectID finds an account associated with a coordinator profile from simulation accounts
// and a project created by this coordinator. The boolean flag `requireNoMainnetLaunchTriggered` is ignored if
// the flag `requireNoMainnetInitialized` is set to `true`
func GetCoordSimAccountWithProjectID(
	r *rand.Rand,
	ctx sdk.Context,
	pk types.ProfileKeeper,
	k keeper.Keeper,
	accs []simtypes.Account,
	requireNoMainnetInitialized bool,
	requireNoMainnetLaunchTriggered bool,
) (simtypes.Account, uint64, bool) {
	projects := k.GetAllProject(ctx)
	campNb := len(projects)
	if campNb == 0 {
		return simtypes.Account{}, 0, false
	}

	r.Shuffle(campNb, func(i, j int) {
		projects[i], projects[j] = projects[j], projects[i]
	})

	// select first project after shuffle
	camp := projects[0]
	// If a criteria is required for the project, we simply fetch the first one that satisfies the criteria
	if requireNoMainnetInitialized {
		var campFound bool
		for _, project := range projects {
			if !project.MainnetInitialized {
				camp = project
				campFound = true
				break
			}
		}
		if !campFound {
			return simtypes.Account{}, 0, false
		}
	}
	if !requireNoMainnetInitialized && requireNoMainnetLaunchTriggered {
		var campFound bool
		for _, project := range projects {
			launched, _ := k.IsProjectMainnetLaunchTriggered(ctx, project.ProjectID)
			if !launched {
				camp = project
				campFound = true
				break
			}
		}
		if !campFound {
			return simtypes.Account{}, 0, false
		}
	}

	// Find the sim account of the project coordinator
	coord, found := pk.GetCoordinator(ctx, camp.CoordinatorID)
	if !found {
		return simtypes.Account{}, 0, false
	}
	for _, acc := range accs {
		if acc.Address.String() == coord.Address && coord.Active {
			return acc, camp.ProjectID, true
		}
	}

	return simtypes.Account{}, 0, false
}

// GetSharesFromProject returns a small portion of shares that can be minted as vouchers or added to an account
func GetSharesFromProject(r *rand.Rand, ctx sdk.Context, k keeper.Keeper, campID uint64) (types.Shares, bool) {
	camp, found := k.GetProject(ctx, campID)
	if !found {
		return types.EmptyShares(), false
	}

	var shares sdk.Coins
	for _, share := range ShareDenoms {
		remaining := int64(k.GetTotalShares(ctx)) - camp.AllocatedShares.AmountOf(share)
		if remaining == 0 {
			continue
		}

		shareNb := r.Int63n(5000) + 10
		if shareNb > remaining {
			shareNb = remaining
		}
		shares = append(shares, sdk.NewCoin(share, sdkmath.NewInt(shareNb)))
	}

	// No shares can be distributed
	if shares.Empty() {
		return types.EmptyShares(), false
	}
	shares = shares.Sort()
	return types.Shares(shares), true
}

// GetAccountWithVouchers returns an account that has vouchers for a project
func GetAccountWithVouchers(
	r *rand.Rand,
	ctx sdk.Context,
	bk types.BankKeeper,
	k keeper.Keeper,
	accs []simtypes.Account,
	requireNoMainnetLaunchTriggered bool,
) (campID uint64, account simtypes.Account, coins sdk.Coins, found bool) {
	var err error
	var accountAddr sdk.AccAddress

	// Parse all account balances and find one with vouchers
	bk.IterateAllBalances(ctx, func(addr sdk.AccAddress, coin sdk.Coin) bool {
		campID, err = types.VoucherProject(coin.Denom)
		if err != nil {
			return false
		}

		if requireNoMainnetLaunchTriggered {
			project, found := k.GetProject(ctx, campID)
			if !found {
				return false
			}
			launched, err := k.IsProjectMainnetLaunchTriggered(ctx, project.ProjectID)
			if err != nil || launched {
				return false
			}
		}

		found = true
		accountAddr = addr
		return true
	})

	// No account has vouchers
	if !found {
		return 0, account, coins, false
	}

	// Fetch from the vouchers of the project owned by the account
	bk.IterateAccountBalances(ctx, accountAddr, func(coin sdk.Coin) bool {
		coinCampID, err := types.VoucherProject(coin.Denom)
		if err == nil && coinCampID == campID {
			// fetch a part of each voucher hold by the account
			amt, err := simtypes.RandPositiveInt(r, coin.Amount)
			if err == nil {
				coins = append(coins, sdk.NewCoin(coin.Denom, amt))
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

// GetAccountWithShares returns an account that contains allocated shares with its associated project
func GetAccountWithShares(
	r *rand.Rand,
	ctx sdk.Context,
	k keeper.Keeper,
	accs []simtypes.Account,
	requireNoMainnetLaunchTriggered bool,
) (uint64, simtypes.Account, types.Shares, bool) {
	mainnetAccounts := k.GetAllMainnetAccount(ctx)
	nb := len(mainnetAccounts)

	// No account have been created yet
	if nb == 0 {
		return 0, simtypes.Account{}, nil, false
	}

	r.Shuffle(nb, func(i, j int) {
		mainnetAccounts[i], mainnetAccounts[j] = mainnetAccounts[j], mainnetAccounts[i]
	})

	// select a mainnet account
	var mainnetAccount types.MainnetAccount
	for _, mAcc := range mainnetAccounts {
		if requireNoMainnetLaunchTriggered {
			project, found := k.GetProject(ctx, mAcc.ProjectID)
			if !found {
				continue
			}
			launched, _ := k.IsProjectMainnetLaunchTriggered(ctx, project.ProjectID)
			if launched {
				continue
			}
		}
		mainnetAccount = mAcc
		break
	}

	// Find the associated sim account
	for _, acc := range accs {
		if acc.Address.String() == mainnetAccount.Address {
			return mainnetAccount.ProjectID, acc, mainnetAccount.Shares, true
		}
	}
	return 0, simtypes.Account{}, nil, false
}
