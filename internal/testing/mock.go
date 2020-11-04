package testing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/tendermint/tendermint/crypto/ed25519"
	"math/rand"
)

// MockRandomString returns a random string of length n
func MockRandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[rand.Intn(len(letter))]
	}
	return string(randomString)
}

// MockAccAddress mocks an account address for test purpose
func MockAccAddress() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr)
}

// MockValAddress mocks an operator address for test purpose
func MockValAddress() sdk.ValAddress {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.ValAddress(addr)
}

// MockCoin mocks a coin allocation structure
func MockCoin() sdk.Coin {
	return sdk.NewCoin(MockRandomString(5), sdk.NewInt(int64(rand.Intn(10000)+1)))
}

// MockCoins mocks coins allocation structure
func MockCoins() sdk.Coins {
	var coins sdk.Coins
	coins = append(coins, MockCoin())
	coins = append(coins, MockCoin())
	coins = append(coins, MockCoin())
	return sdk.NewCoins()
}

// MockDescription mocks a validator description structure
func MockDescription() staking.Description {
	return staking.NewDescription(
		MockRandomString(10),
		MockRandomString(10),
		MockRandomString(10),
		MockRandomString(10),
		MockRandomString(10),
	)
}

// MockCommissionRates mocks a commissionRates structure
func MockCommissionRates() staking.CommissionRates {
	return staking.NewCommissionRates(
		sdk.NewDec(int64(rand.Intn(10000)+1)),
		sdk.NewDec(int64(rand.Intn(10000)+1)),
		sdk.NewDec(int64(rand.Intn(10000)+1)),
	)
}
