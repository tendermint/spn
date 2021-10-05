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
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
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

// Address returns a sample account address
func Address() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr)
}

// AccAddress returns a sample account address
func AccAddress() string {
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

// Voucher returns a sample voucher structure
func Voucher(campaignID uint64) sdk.Coin {
	denom := campaigntypes.VoucherDenom(campaignID, AlphaString(5))
	return sdk.NewCoin(denom, sdk.NewInt(int64(rand.Intn(10000)+1)))
}

// Vouchers returns a sample vouchers structure
func Vouchers(campaignID uint64) sdk.Coins {
	return sdk.NewCoins(Voucher(campaignID), Voucher(campaignID), Voucher(campaignID))
}
