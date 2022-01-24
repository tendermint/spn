package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/pkg/types"
	"testing"
)

func signatureCount(t *testing.T, consAddr, sig string) types.SignatureCount {
	sigDec, err := sdk.NewDecFromStr(sig)
	require.NoError(t, err)
	return types.SignatureCount{
		ConsAddress: consAddr,
		RelativeSignatures: sigDec,
	}
}

func TestNewSignatureCounts(t *testing.T) {
	sc := types.NewSignatureCounts()
	require.Zero(t, sc)
}

func TestSignatureCounts_AddSignature(t *testing.T) {
	tests := []struct {
		name   string
		sc types.SignatureCounts
		consAddres string
		validatorSetSize int64
		expected   types.SignatureCounts
	}{
		{
			name: "new signature",
			sc: types.SignatureCounts{
				BlockCount: 1,
				Counts: []types.SignatureCount{},
			},
			consAddres: "foo",
			validatorSetSize: 1,
			expected: types.SignatureCounts{
				BlockCount: 1,
				Counts: []types.SignatureCount{
					signatureCount(t, "foo", "1"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sc.AddSignature(tt.consAddres, tt.validatorSetSize)
			require.Equal(t, tt.expected, tt.sc)
		})
	}
}

func TestSignatureCounts_Validate(t *testing.T) {
	tests := []struct {
		name string
		sc types.SignatureCounts
		wantErr bool
	}{
		{
			name: "valid signature counts",
			sc: types.SignatureCounts{
				BlockCount: 2,
				Counts: []types.SignatureCount{
					signatureCount(t, "foo", "1"),
					signatureCount(t, "bar", "0.1"),
					signatureCount(t, "foobar", "0.5"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.sc.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
