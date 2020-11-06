package testing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/tendermint/tendermint/crypto"
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

// MockRandomAlphaString returns a random string with lowercase alpha char of length n
func MockRandomAlphaString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz")

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

// MockValAddress mocks an operator address with associated private key for test purpose
func MockValAddress() (crypto.PrivKey, sdk.ValAddress) {
	privKey := secp256k1.GenPrivKey()
	pk := privKey.PubKey()
	addr := pk.Address()
	return privKey, sdk.ValAddress(addr)
}

// MockPubKey mocks a public key
func MockPubKey() crypto.PubKey {
	return ed25519.GenPrivKey().PubKey()
}

// MockCoin mocks a coin allocation structure
func MockCoin() sdk.Coin {
	return sdk.NewCoin(MockRandomAlphaString(5), sdk.NewInt(int64(rand.Intn(10000)+1)))
}

// MockCoins mocks coins allocation structure
func MockCoins() sdk.Coins {
	var coins sdk.Coins
	// Coin denomination must be sorted
	coin := sdk.NewCoin("a"+MockRandomAlphaString(5), sdk.NewInt(int64(rand.Intn(10000)+1)))
	coins = append(coins, coin)
	coin = sdk.NewCoin("b"+MockRandomAlphaString(5), sdk.NewInt(int64(rand.Intn(10000)+1)))
	coins = append(coins, coin)
	coin = sdk.NewCoin("c"+MockRandomAlphaString(5), sdk.NewInt(int64(rand.Intn(10000)+1)))
	coins = append(coins, coin)
	return coins
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
	return staking.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())
}
