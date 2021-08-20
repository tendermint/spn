package sample

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	launch "github.com/tendermint/spn/x/launch/types"
	profile "github.com/tendermint/spn/x/profile/types"
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

// Codec returns a codec with preregistered interfaces
func Codec() codec.Marshaler {
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	launch.RegisterInterfaces(interfaceRegistry)
	profile.RegisterInterfaces(interfaceRegistry)

	return codec.NewProtoCodec(interfaceRegistry)
}

// Bytes returns a random array of bytes
func Bytes(n int) []byte {
	return []byte(String(n))
}

// Uint64 returns a random uint64
func Uint64() uint64 {
	return uint64(rand.Intn(10000))
}

// String returns a random string of length n
func String(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[rand.Intn(len(letter))]
	}
	return string(randomString)
}

// AlphaString returns a random string with lowercase alpha char of length n
func AlphaString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[rand.Intn(len(letter))]
	}
	return string(randomString)
}

// AccAddress returns a sample account address
func AccAddress() string {
	setAddressPrefixes()

	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// Coin returns a sample coin structure
func Coin() sdk.Coin {
	return sdk.NewCoin(AlphaString(5), sdk.NewInt(int64(rand.Intn(10000)+1)))
}

// Coins returns a sample coins structure
func Coins() sdk.Coins {
	return sdk.NewCoins(Coin(), Coin(), Coin())
}
