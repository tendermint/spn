// Package constructor provides constructors to easily initialize objects for test purpose with automatic error handling
package constructor

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
)

// Coins returns a sdk.Coins from a string
func Coins(t testing.TB, str string) sdk.Coins {
	coins, err := sdk.ParseCoinsNormalized(str)
	require.NoError(t, err)
	return coins
}

// Dec returns a sdk.Dec from a string
func Dec(t testing.TB, str string) sdk.Dec {
	dec, err := sdk.NewDecFromStr(str)
	require.NoError(t, err)
	return dec
}

// SignatureCount returns a signature count object for test from a cons address and a decimal string for relative signatures
func SignatureCount(t *testing.T, consAddr []byte, relSig string) spntypes.SignatureCount {
	return spntypes.SignatureCount{
		ConsAddress:        consAddr,
		RelativeSignatures: Dec(t, relSig),
	}
}

// SignatureCounts returns a signature counts object for tests from a a block count and list of signature counts
func SignatureCounts(blockCount uint64, sc ...spntypes.SignatureCount) spntypes.SignatureCounts {
	return spntypes.SignatureCounts{
		BlockCount: blockCount,
		Counts:     sc,
	}
}
