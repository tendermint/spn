package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

var (
	// ShareDenoms are the denom used for the shares in the simulation
	ShareDenoms = []string{"s/foo", "s/bar", "s/toto"}
)

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

// GetCoordSimAccountWithCampaignID finds an account associated with a coordinator profile from simulation accounts
// and a campaign created by this coordinator. The boolean flag `requireNoMainnetLaunchTriggered` is ignored if
// the flag `requireNoMainnetInitialized` is set to `true`
func GetCoordSimAccountWithCampaignID(
	r *rand.Rand,
	ctx sdk.Context,
	pk types.ProfileKeeper,
	k keeper.Keeper,
	accs []simtypes.Account,
	requireNoMainnetInitialized bool,
	requireNoMainnetLaunchTriggered bool,
) (simtypes.Account, uint64, bool) {
	campaigns := k.GetAllCampaign(ctx)
	campNb := len(campaigns)
	if campNb == 0 {
		return simtypes.Account{}, 0, false
	}

	r.Shuffle(campNb, func(i, j int) {
		campaigns[i], campaigns[j] = campaigns[j], campaigns[i]
	})

	// select first campaign after shuffle
	camp := campaigns[0]
	// If a criteria is required for the campaign, we simply fetch the first one that satisfies the criteria
	if requireNoMainnetInitialized {
		var campFound bool
		for _, campaign := range campaigns {
			if !campaign.MainnetInitialized {
				camp = campaign
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
		for _, campaign := range campaigns {
			launched, _ := k.IsCampaignMainnetLaunchTriggered(ctx, campaign.CampaignID)
			if !launched {
				camp = campaign
				campFound = true
				break
			}
		}
		if !campFound {
			return simtypes.Account{}, 0, false
		}
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

// GetSharesFromCampaign returns a small portion of shares that can be minted as vouchers or added to an account
func GetSharesFromCampaign(r *rand.Rand, ctx sdk.Context, k keeper.Keeper, campID uint64) (types.Shares, bool) {
	camp, found := k.GetCampaign(ctx, campID)
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
	k keeper.Keeper,
	accs []simtypes.Account,
	requireNoMainnetLaunchTriggered bool,
) (campID uint64, account simtypes.Account, coins sdk.Coins, found bool) {
	var err error
	var accountAddr sdk.AccAddress

	// Parse all account balances and find one with vouchers
	bk.IterateAllBalances(ctx, func(addr sdk.AccAddress, coin sdk.Coin) bool {
		campID, err = types.VoucherCampaign(coin.Denom)
		if err != nil {
			return false
		}

		if requireNoMainnetLaunchTriggered {
			campaign, found := k.GetCampaign(ctx, campID)
			if !found {
				return false
			}
			launched, err := k.IsCampaignMainnetLaunchTriggered(ctx, campaign.CampaignID)
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

	// Fetch from the vouchers of the campaign owned by the account
	bk.IterateAccountBalances(ctx, accountAddr, func(coin sdk.Coin) bool {
		coinCampID, err := types.VoucherCampaign(coin.Denom)
		if err == nil && coinCampID == campID {
			// retain a random portion of the balance in the range [0, coin.Amount)
			retainAmt := sdk.NewInt(rand.Int63n(coin.Amount.Int64()))
			coin.Amount = coin.Amount.Sub(retainAmt)
			coins = append(coins, coin)
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
			campaign, found := k.GetCampaign(ctx, mAcc.CampaignID)
			if !found {
				continue
			}
			launched, _ := k.IsCampaignMainnetLaunchTriggered(ctx, campaign.CampaignID)
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
			return mainnetAccount.CampaignID, acc, mainnetAccount.Shares, true
		}
	}
	return 0, simtypes.Account{}, nil, false
}
