package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestChainIDFromChainName(t *testing.T) {
	require.Equal(t, "foo-1", types.ChainIDFromChainName("foo", 1))
}

func TestCheckChainName(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		name string
		valid bool
	}{
		{
			desc:  "Valid name",
			name:  "FooBar999",
			valid: true,
		},
		{
			desc:  "No empty",
			name:  "",
			valid: false,
		},
		{
			desc:  "No special character",
			name:  "foo-bar",
			valid: false,
		},
		{
			desc:  "No space",
			name:  "foo bar",
			valid: false,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			if tc.valid {
				require.NoError(t, types.CheckChainName(tc.name))
			} else {
				require.Error(t, types.CheckChainName(tc.name))
			}
		})
	}
}