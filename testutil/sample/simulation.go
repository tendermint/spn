package sample

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

const (
	simAccountsNb = 100
)

// SimAccounts returns a sample array of account for simulation
func SimAccounts() (accounts []simtypes.Account) {
	for i := 0; i < simAccountsNb; i++ {
		privKey := ed25519.GenPrivKey()
		pubKey := privKey.PubKey()
		acc := simtypes.Account{
			PrivKey: privKey,
			PubKey:  pubKey,
			Address: sdk.AccAddress(pubKey.Address()),
			ConsKey: privKey,
		}
		accounts = append(accounts, acc)
	}
	return
}

// Rand returns a sample Rand object for randomness
func Rand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().Unix()))
}

// Fees returns a random fee by selecting a random amount of bond denomination
// from the account's available balance. If the user doesn't have enough funds for
// paying fees, it returns empty coins.
func Fees(r *rand.Rand, spendableCoins sdk.Coins) (sdk.Coins, error) {
	if spendableCoins.Empty() {
		return nil, nil
	}

	bondDenomAmt := spendableCoins.AmountOf(sdk.DefaultBondDenom)
	if bondDenomAmt.IsZero() {
		return nil, nil
	}

	amt, err := simtypes.RandPositiveInt(r, bondDenomAmt)
	if err != nil {
		return nil, err
	}

	if amt.IsZero() {
		return nil, nil
	}

	fees := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, amt))
	return fees, nil
}
