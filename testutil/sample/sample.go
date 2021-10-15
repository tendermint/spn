package sample

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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

// AccAddress returns a sample account address
func AccAddress() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr)
}

// Address returns a sample string account address
func Address() string {
	return AccAddress().String()
}

// Coin returns a sample coin structure
func Coin() sdk.Coin {
	return sdk.NewCoin(AlphaString(5), sdk.NewInt(int64(rand.Intn(10000)+1)))
}

// Coins returns a sample coins structure
func Coins() sdk.Coins {
	return sdk.NewCoins(Coin(), Coin(), Coin())
}