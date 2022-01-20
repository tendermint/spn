package types

import (
	"encoding/binary"

	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

const (
	signMessage   = "StarportNetwork-MsgSetValidatorConsAddress"
	signSeparator = byte('/')
)

// SignValidatorMessage sign the validator consensus message with
// chain id and nonce using the mnemonic
func SignValidatorMessage(mnemonic, chainID string, nonce uint64) ([]byte, error) {
	priv := ed25519.GenPrivKeyFromSecret([]byte(mnemonic))
	msg := createSignMessage(chainID, nonce)
	return priv.Sign(msg)
}

// CheckValidatorSignature check the validator consensus signed message with
// chain id, nonce and public key
func CheckValidatorSignature(signature []byte, consPubKey, chainID string, nonce uint64) error {
	pub := ed25519.PubKey(consPubKey)
	msg := createSignMessage(chainID, nonce)
	valid := pub.VerifySignature(msg, signature)
	if !valid {
		return errors.New("invalid signature")
	}
	return nil
}

// createSignMessage create the sign message with nonce and chain id
func createSignMessage(chainID string, nonce uint64) []byte {
	msg := append([]byte(signMessage), signSeparator)
	nonceBytes := append(uintBytes(nonce), signSeparator)
	msg = append(msg, nonceBytes...)
	chainIDBytes := append([]byte(chainID), signSeparator)
	msg = append(msg, chainIDBytes...)
	return msg
}

// uintBytes convert uint64 to byte slice
func uintBytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
