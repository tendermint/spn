package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	// signMessage is the size message constant
	signMessage = "StarportNetwork-MsgSetValidatorConsAddress"
	// signSeparator is the size message separator
	signSeparator = byte('/')
)

// ValidatorKey stores the validator private key
type ValidatorKey struct {
	Address crypto.Address
	PubKey  crypto.PubKey
	PrivKey crypto.PrivKey
}

// Sign signs the message with privateKey and returns a signature
func (v ValidatorKey) Sign(nonce uint64) ([]byte, error) {
	return v.PrivKey.Sign(CreateSignMessage(nonce))
}

// VerifySignature reports whether sig is a valid signature of mes
func (v ValidatorKey) VerifySignature(nonce uint64, sig []byte) bool {
	return v.PubKey.VerifySignature(CreateSignMessage(nonce), sig)
}

// GetConsAddress return the validator consensus address
func (v ValidatorKey) GetConsAddress() types.ConsAddress {
	return types.ConsAddress(v.PubKey.Address())
}

// LoadValidatorKey load the validator key file into the ValidatorKey struct
func LoadValidatorKey(keyJSONBytes []byte) (pvKey ValidatorKey, err error) {
	err = json.Unmarshal(keyJSONBytes, &pvKey)
	pvKey.PubKey = pvKey.PrivKey.PubKey()
	pvKey.Address = pvKey.PubKey.Address()
	return
}

// CreateSignMessage create the sign message with nonce and chain id
func CreateSignMessage(nonce uint64) []byte {
	msg := append([]byte(signMessage), signSeparator)
	nonceBytes := append(UintBytes(nonce), signSeparator)
	msg = append(msg, nonceBytes...)
	return msg
}
