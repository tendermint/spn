// Package types defines types for monitored validator activities
package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewSignatureCounts returns a new SignatureCounts
func NewSignatureCounts() SignatureCounts {
	return SignatureCounts{}
}

// AddSignature adds a signature for the consensus address at a specific validator set size
func (m *SignatureCounts) AddSignature(consAddress string, validatorSetSize int64) {
	// relative signature is the signature relative to the validator set size
	relSignature := sdk.NewDec(1).QuoInt(sdk.NewInt(validatorSetSize))

	// search for the consensus address
	for i, c := range m.Counts {
		if c.ConsAddress == consAddress {
			m.Counts[i].RelativeSignatures = c.RelativeSignatures.Add(relSignature)
			return
		}
	}

	// consensus address not found, a new one is added to the signature counts
	m.Counts = append(m.Counts, SignatureCount{
		ConsAddress:        consAddress,
		RelativeSignatures: relSignature,
	})
	return
}

// Validate checks if the signature counts object is valid
// the sum of all relative signatures should not exceed the number of block
func (m SignatureCounts) Validate() error {
	consAddr := make(map[string]struct{})
	sumSig := sdk.NewDec(0)

	// iterate all signature count
	for _, sc := range m.Counts {
		// a consensus address must have a single entry
		if _, ok := consAddr[sc.ConsAddress]; ok {
			return fmt.Errorf("duplicated consensus address %s", sc.ConsAddress)
		}
		consAddr[sc.ConsAddress] = struct{}{}
		sumSig = sumSig.Add(sc.RelativeSignatures)
	}

	blockCountDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(m.BlockCount))
	if sumSig.GT(blockCountDec) {
		return fmt.Errorf(
			"sum of relative signatures is higher than block number %s > %s",
			sumSig.String(),
			blockCountDec.String(),
		)
	}
	return nil
}
