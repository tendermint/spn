// Package sample provides methods to initialize sample object of various types for test purposes
package sample

import (
	"math/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cosmosed25519 "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctypes "github.com/cosmos/ibc-go/v2/modules/core/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"

	campaign "github.com/tendermint/spn/x/campaign/types"

	launch "github.com/tendermint/spn/x/launch/types"
	profile "github.com/tendermint/spn/x/profile/types"
)

// Codec returns a codec with preregistered interfaces
func Codec() codec.Codec {
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	cryptocodec.RegisterInterfaces(interfaceRegistry)
	authtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	launch.RegisterInterfaces(interfaceRegistry)
	profile.RegisterInterfaces(interfaceRegistry)
	ibctypes.RegisterInterfaces(interfaceRegistry)

	return codec.NewProtoCodec(interfaceRegistry)
}

// Bool returns randomly true or false
func Bool() bool {
	r := rand.Intn(100)
	return r < 50
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

// PubKey returns a sample account PubKey
func PubKey() crypto.PubKey {
	return ed25519.GenPrivKey().PubKey()
}

// ConsAddress returns a sample consensus address
func ConsAddress() sdk.ConsAddress {
	return sdk.ConsAddress(PubKey().Address())
}

// AccAddress returns a sample account address
func AccAddress() sdk.AccAddress {
	addr := PubKey().Address()
	return sdk.AccAddress(addr)
}

// Address returns a sample string account address
func Address() string {
	return AccAddress().String()
}

// ValAddress returns a sample validator operator address
func ValAddress() sdk.ValAddress {
	return sdk.ValAddress(PubKey().Address())
}

// OperatorAddress returns a sample string validator operator address
func OperatorAddress() string {
	return ValAddress().String()
}

// Validator returns a sample staking validator
func Validator(t testing.TB) stakingtypes.Validator {
	val, err := stakingtypes.NewValidator(
		ValAddress(),
		cosmosed25519.GenPrivKey().PubKey(),
		stakingtypes.Description{})
	require.NoError(t, err)
	return val
}

// Coin returns a sample coin structure
func Coin() sdk.Coin {
	return sdk.NewCoin(AlphaString(5), sdk.NewInt(rand.Int63n(10000)+1))
}

// CoinWithRange returns a sample coin structure where the amount is a random number between provided min and max values
// with a random denom
func CoinWithRange(min, max int64) sdk.Coin {
	return sdk.NewCoin(AlphaString(5), sdk.NewInt(rand.Int63n(max-min)+min))
}

// CoinWithRangeAmount returns a sample coin structure where the amount is a random number between provided min and max values
// with a given denom
func CoinWithRangeAmount(denom string, min, max int64) sdk.Coin {
	return sdk.NewCoin(denom, sdk.NewInt(rand.Int63n(max-min)+min))
}

// Coins returns a sample coins structure
func Coins() sdk.Coins {
	return sdk.NewCoins(Coin(), Coin(), Coin())
}

// CoinsWithRange returns a sample coins structure where the amount is a random number between provided min and max values
func CoinsWithRange(min, max int64) sdk.Coins {
	return sdk.NewCoins(CoinWithRange(min, max), CoinWithRange(min, max), CoinWithRange(min, max))
}

// CoinsWithRangeAmount returns a sample coins structure where the amount is a random number between provided min and max values
// with a set of given denoms
func CoinsWithRangeAmount(denom1, denom2, denom3 string, min, max int64) sdk.Coins {
	return sdk.NewCoins(CoinWithRangeAmount(denom1, min, max), CoinWithRangeAmount(denom2, min, max), CoinWithRangeAmount(denom3, min, max))
}

// TotalSupply returns a sample coins structure where each denom's total supply is within the default
// allowed supply range
func TotalSupply() sdk.Coins {
	return CoinsWithRange(campaign.DefaultMinTotalSupply.Int64(), campaign.DefaultMaxTotalSupply.Int64())
}
