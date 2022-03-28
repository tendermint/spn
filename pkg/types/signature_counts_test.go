package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/pkg/types"
	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
)

var (
	opAddrFoo    = sample.OperatorAddress(r)
	opAddrBar    = sample.OperatorAddress(r)
	opAddrBaz    = sample.OperatorAddress(r)
	opAddrFoobar = sample.OperatorAddress(r)
)

func TestNewSignatureCounts(t *testing.T) {
	sc := types.NewSignatureCounts()
	require.Zero(t, sc)
}

func TestSignatureCounts_AddSignature(t *testing.T) {
	tests := []struct {
		name             string
		sc               types.SignatureCounts
		opAddres         string
		validatorSetSize int64
		expected         types.SignatureCounts
	}{
		{
			name: "a new signature in a empty object should contain only the signature",
			sc: types.SignatureCounts{
				BlockCount: 1,
				Counts:     []types.SignatureCount{},
			},
			opAddres:         opAddrFoo,
			validatorSetSize: 1,
			expected: types.SignatureCounts{
				BlockCount: 1,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "1"),
				},
			},
		},
		{
			name: "validator set size should influence the relative signatures",
			sc: types.SignatureCounts{
				BlockCount: 100,
				Counts:     []types.SignatureCount{},
			},
			opAddres:         opAddrFoo,
			validatorSetSize: 10000,
			expected: types.SignatureCounts{
				BlockCount: 100,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "0.0001"),
				},
			},
		},
		{
			name: "a new address should add a new entry in the object",
			sc: types.SignatureCounts{
				BlockCount: 100,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "1"),
					tc.SignatureCount(t, opAddrBar, "0.5"),
					tc.SignatureCount(t, opAddrBaz, "5.5"),
				},
			},
			opAddres:         opAddrFoobar,
			validatorSetSize: 10,
			expected: types.SignatureCounts{
				BlockCount: 100,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "1"),
					tc.SignatureCount(t, opAddrBar, "0.5"),
					tc.SignatureCount(t, opAddrBaz, "5.5"),
					tc.SignatureCount(t, opAddrFoobar, "0.1"),
				},
			},
		},
		{
			name: "an existing address should update then existing entry in the object",
			sc: types.SignatureCounts{
				BlockCount: 100,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "1"),
					tc.SignatureCount(t, opAddrBar, "0.5"),
					tc.SignatureCount(t, opAddrBaz, "5.5"),
				},
			},
			opAddres:         opAddrBar,
			validatorSetSize: 10,
			expected: types.SignatureCounts{
				BlockCount: 100,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "1"),
					tc.SignatureCount(t, opAddrBar, "0.6"),
					tc.SignatureCount(t, opAddrBaz, "5.5"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sc.AddSignature(tt.opAddres, tt.validatorSetSize)
			require.Equal(t, tt.expected, tt.sc)
		})
	}
}

func TestSignatureCounts_Validate(t *testing.T) {
	tests := []struct {
		name    string
		sc      types.SignatureCounts
		wantErr bool
	}{
		{
			name: "empty is valid",
			sc:   types.NewSignatureCounts(),
		},
		{
			name: "valid signature counts",
			sc: types.SignatureCounts{
				BlockCount: 2,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "1"),
					tc.SignatureCount(t, opAddrBar, "0.1"),
					tc.SignatureCount(t, opAddrFoobar, "0.5"),
				},
			},
		},
		{
			name: "sum of relative signatures equals block count is valid",
			sc: types.SignatureCounts{
				BlockCount: 2,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "0.5"),
					tc.SignatureCount(t, opAddrBar, "0.5"),
					tc.SignatureCount(t, opAddrBaz, "0.5"),
					tc.SignatureCount(t, opAddrFoobar, "0.5"),
				},
			},
		},
		{
			name: "duplicated consensus address",
			sc: types.SignatureCounts{
				BlockCount: 2,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "1"),
					tc.SignatureCount(t, opAddrBar, "0.1"),
					tc.SignatureCount(t, opAddrBar, "0.5"),
				},
			},
			wantErr: true,
		},
		{
			name: "sum of relative signatures equals is greater than block count",
			sc: types.SignatureCounts{
				BlockCount: 2,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "1"),
					tc.SignatureCount(t, opAddrBar, "1"),
					tc.SignatureCount(t, opAddrBaz, "0.5"),
				},
			},
			wantErr: true,
		},
		{
			name: "invalid operator address",
			sc: types.SignatureCounts{
				BlockCount: 2,
				Counts: []types.SignatureCount{
					tc.SignatureCount(t, opAddrFoo, "1"),
					tc.SignatureCount(t, "invalid_bech_32", "0.1"),
					tc.SignatureCount(t, opAddrFoobar, "0.5"),
				},
			},
			wantErr: true,
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

func TestSignatureCounts_GetOperatorAddress(t *testing.T) {
	converted, err := tc.SignatureCount(t, "cosmosvaloper17qqqsmyx43efr5ywp33h35l4dlmacfxvtkyed5", "1").GetOperatorAddress(types.AccountAddressPrefix)
	require.NoError(t, err)
	require.EqualValues(t, "spn17qqqsmyx43efr5ywp33h35l4dlmacfxvj7g20a", converted)

	_, err = tc.SignatureCount(t, "invalid_bech_32", "1").GetOperatorAddress(types.AccountAddressPrefix)
	require.Error(t, err)
}
