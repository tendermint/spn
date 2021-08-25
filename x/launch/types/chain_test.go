package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestNewGenesisChainID(t *testing.T) {
	require.Equal(t, "foo-1", types.NewGenesisChainID("foo", 1))
}

func TestParseGenesisChainID(t *testing.T) {
	for _, tc := range []struct {
		desc           string
		chainID        string
		expectedName   string
		expectedNumber uint64
		valid          bool
	}{
		{
			desc:           "Valid chainID",
			chainID:        "foo-42",
			expectedName:   "foo",
			expectedNumber: 42,
			valid:          true,
		},
		{
			desc:    "Empty",
			chainID: "",
			valid:   false,
		},
		{
			desc:    "No separator",
			chainID: "foo42",
			valid:   false,
		},
		{
			desc:    "Too many separators",
			chainID: "foo-42-32",
			valid:   false,
		},
		{
			desc:    "No number",
			chainID: "foo-",
			valid:   false,
		},
		{
			desc:    "Invalid number",
			chainID: "foo-fortytwo",
			valid:   false,
		},
		{
			desc:    "No chain name",
			chainID: "-42",
			valid:   false,
		},
		{
			desc:    "Invalid chain name",
			chainID: "foo/bar-42",
			valid:   false,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			if tc.valid {
				name, number, err := types.ParseGenesisChainID(tc.chainID)
				require.NoError(t, err)
				require.EqualValues(t, tc.expectedName, name)
				require.EqualValues(t, tc.expectedNumber, number)
			} else {
				_, _, err := types.ParseGenesisChainID(tc.chainID)
				require.Error(t, err)
			}
		})
	}
}

func TestCheckChainName(t *testing.T) {
	for _, tc := range []struct {
		desc  string
		name  string
		valid bool
	}{
		{
			desc:  "Valid name",
			name:  "foobar",
			valid: true,
		},
		{
			desc:  "No uppercase",
			name:  "FooBar",
			valid: false,
		},
		{
			desc:  "No uppercase",
			name:  "FooBar",
			valid: false,
		},
		{
			desc:  "No empty",
			name:  "",
			valid: false,
		},
		{
			desc:  "Too big",
			name:  sample.AlphaString(types.ChainNameMaxLength + 1),
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
