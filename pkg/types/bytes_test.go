package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/pkg/types"
)

func TestUintBytes(t *testing.T) {
	tests := []struct {
		name string
		v    uint64
		want []byte
	}{
		{name: "zero value", v: 0, want: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{name: "one value", v: 1, want: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}},
		{name: "teen value", v: 100, want: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := types.UintBytes(tt.v)
			require.Equal(t, tt.want, got)
		})
	}
}
