package sample

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

const (
	//BondDenom defines the bond denom used in testing
	BondDenom = "stake"

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
