package sample

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/rand"
)

const accountAddressPrefix = "spn"

var configSet = false

func setAddressPrefixes() {
	if !configSet {
		// Set prefixes
		accountPubKeyPrefix := accountAddressPrefix + "pub"
		validatorAddressPrefix := accountAddressPrefix + "valoper"
		validatorPubKeyPrefix := accountAddressPrefix + "valoperpub"
		consNodeAddressPrefix := accountAddressPrefix + "valcons"
		consNodePubKeyPrefix := accountAddressPrefix + "valconspub"

		// Set and seal config
		config := sdk.GetConfig()
		config.SetBech32PrefixForAccount(accountAddressPrefix, accountPubKeyPrefix)
		config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
		config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
		config.Seal()

		configSet = true
	}
}

// SampleString returns a random string of length n
func SampleString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[rand.Intn(len(letter))]
	}
	return string(randomString)
}

// SampleAlphaString returns a random string with lowercase alpha char of length n
func SampleAlphaString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[rand.Intn(len(letter))]
	}
	return string(randomString)
}

// SampleAccAddress returns a sample account address
func SampleAccAddress() string {
	setAddressPrefixes()

	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// SampleCoin returns a sample coin structure
func SampleCoin() sdk.Coin {
	return sdk.NewCoin(SampleAlphaString(5), sdk.NewInt(int64(rand.Intn(10000)+1)))
}

// SampleCoin returns a sample coins structure
func SampleCoins() sdk.Coins {
	return sdk.NewCoins(SampleCoin(), SampleCoin(), SampleCoin())
}