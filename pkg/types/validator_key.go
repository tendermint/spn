package types

import (
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	tmjson "github.com/tendermint/tendermint/libs/json"
)

const (
	// signMessage is the size message constant
	signMessage = "StarportNetwork-MsgSetValidatorConsAddress"
	// signSeparator is the size message separator
	signSeparator = byte('/')
)

// ValidatorKey stores the validator private key
type ValidatorKey struct {
	Address crypto.Address `json:"address"`
	PubKey  crypto.PubKey  `json:"pub_key"`
	PrivKey crypto.PrivKey `json:"priv_key"`
}

// Sign signs the message with privateKey and returns a signature
func (v ValidatorKey) Sign(nonce uint64, chainID string) (string, error) {
	sign, err := v.PrivKey.Sign(CreateSignMessage(nonce, chainID))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

// VerifySignature reports whether sig is a valid signature of mes
func (v ValidatorKey) VerifySignature(nonce uint64, chainID, sig string) bool {
	sigBytes, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return false
	}
	return v.PubKey.VerifySignature(CreateSignMessage(nonce, chainID), sigBytes)
}

// GetConsAddress return the validator consensus address
func (v ValidatorKey) GetConsAddress() types.ConsAddress {
	return types.ConsAddress(v.PubKey.Address())
}

// LoadValidatorKey load the validator key file into the ValidatorKey struct
func LoadValidatorKey(keyJSONBytes []byte) (pvKey ValidatorKey, err error) {
	err = tmjson.Unmarshal(keyJSONBytes, &pvKey)
	if err != nil {
		return pvKey, fmt.Errorf("error reading PrivValidator key: %s", err)
	}

	// overwrite pubkey and address for convenience
	pvKey.PubKey = pvKey.PrivKey.PubKey()
	pvKey.Address = pvKey.PubKey.Address()
	return
}

// CreateSignMessage create the sign message with nonce and chain id
func CreateSignMessage(nonce uint64, chainID string) []byte {
	msg := append([]byte(signMessage), signSeparator)
	nonceBytes := append(UintBytes(nonce), signSeparator)
	msg = append(msg, nonceBytes...)
	chainIDBytes := append([]byte(chainID), signSeparator)
	msg = append(msg, chainIDBytes...)
	return msg
}
