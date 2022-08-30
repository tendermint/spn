// Package types defines types for monitored validator activities
package types

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewSignatureCounts returns a new SignatureCounts
func NewSignatureCounts() SignatureCounts {
	return SignatureCounts{}
}

// AddSignature adds a signature for the consensus address at a specific validator set size
func (m *SignatureCounts) AddSignature(opAddress string, validatorSetSize int64) {
	// relative signature is the signature relative to the validator set size
	relSignature := sdk.OneDec().QuoInt(sdkmath.NewInt(validatorSetSize))

	// search for the consensus address
	for i, c := range m.Counts {
		if c.OpAddress == opAddress {
			m.Counts[i].RelativeSignatures = c.RelativeSignatures.Add(relSignature)
			return
		}
	}

	// consensus address not found, a new one is added to the signature counts
	m.Counts = append(m.Counts, SignatureCount{
		OpAddress:          opAddress,
		RelativeSignatures: relSignature,
	})
}

// Validate checks if the signature counts object is valid
// the sum of all relative signatures should not exceed the number of block
func (m SignatureCounts) Validate() error {
	opAddr := make(map[string]struct{})
	sumSig := sdk.ZeroDec()

	// iterate all signature count
	for _, sc := range m.Counts {
		// check is the signer has a invalid bech32 address
		_, err := sc.GetOperatorAddress(AccountAddressPrefix)
		if err != nil {
			return errors.Wrapf(err, "invalid bech32 operator address: %s", sc.OpAddress)
		}

		// a consensus address must have a single entry
		if _, ok := opAddr[sc.OpAddress]; ok {
			return fmt.Errorf("duplicated operator address %s", sc.OpAddress)
		}
		opAddr[sc.OpAddress] = struct{}{}
		sumSig = sumSig.Add(sc.RelativeSignatures)
	}

	blockCountDec := sdk.NewDecFromInt(sdkmath.NewIntFromUint64(m.BlockCount))
	if sumSig.GT(blockCountDec) {
		return fmt.Errorf(
			"sum of relative signatures is higher than block number %s > %s",
			sumSig.String(),
			blockCountDec.String(),
		)
	}
	return nil
}

// GetOperatorAddress returns the operator address for the signer with the SPN prefix format
func (m SignatureCount) GetOperatorAddress(accountPrefix string) (string, error) {
	_, decoded, err := bech32.DecodeAndConvert(m.OpAddress)
	if err != nil {
		return "", err
	}
	return bech32.ConvertAndEncode(accountPrefix, decoded)
}
