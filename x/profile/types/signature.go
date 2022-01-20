package types

import (
	"encoding/binary"
	"strconv"

	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

const (
	// PublicKeySize is the size, in bytes
	PublicKeySize = 32
	// signMessage is the size message constant
	signMessage = "StarportNetwork-MsgSetValidatorConsAddress"
	// signSeparator is the size message separator
	signSeparator = byte('/')
)

// SignValidatorMessage sign the validator consensus message with
// chain id and nonce using the mnemonic
func SignValidatorMessage(mnemonic, chainID string, nonce uint64) ([]byte, error) {
	priv := ed25519.GenPrivKeyFromSecret([]byte(mnemonic))
	msg := CreateSignMessage(chainID, nonce)
	return priv.Sign(msg)
}

// CheckValidatorSignature check the validator consensus signed message with
// chain id, nonce and public key
func CheckValidatorSignature(signature, consPubKey []byte, chainID string, nonce uint64) error {
	if l := len(consPubKey); l != PublicKeySize {
		return errors.New("ed25519: bad public key length: " + strconv.Itoa(l))
	}
	pub := ed25519.PubKey(consPubKey)
	msg := CreateSignMessage(chainID, nonce)
	valid := pub.VerifySignature(msg, signature)
	if !valid {
		return errors.New("invalid signature")
	}
	return nil
}

// CreateSignMessage create the sign message with nonce and chain id
func CreateSignMessage(chainID string, nonce uint64) []byte {
	msg := append([]byte(signMessage), signSeparator)
	nonceBytes := append(UintBytes(nonce), signSeparator)
	msg = append(msg, nonceBytes...)
	chainIDBytes := append([]byte(chainID), signSeparator)
	msg = append(msg, chainIDBytes...)
	return msg
}

// UintBytes convert uint64 to byte slice
func UintBytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
