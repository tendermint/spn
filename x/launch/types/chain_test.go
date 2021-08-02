package types_test

import (
	"testing"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestChain_GetDefaultInitialGenesis(t *testing.T) {
	cdc := sample.Codec()
	chain := sample.Chain("foo", 0)

	// Should return the default genesis
	chain.InitialGenesis, _ = codec.NewAnyWithValue(&types.DefaultInitialGenesis{})
	defaultGen, err := chain.GetDefaultInitialGenesis(cdc)
	require.NoError(t, err)
	require.Equal(t, &types.DefaultInitialGenesis{}, defaultGen)

	// Should return an error
	url, hash := sample.String(10), sample.String(10)
	chain.InitialGenesis, _ = codec.NewAnyWithValue(&types.GenesisURL{url, hash})
	_, err = chain.GetDefaultInitialGenesis(cdc)
	require.Error(t, err)
}

func TestChain_GetInitialGenesis(t *testing.T) {
	cdc := sample.Codec()
	chain := sample.Chain("foo", 0)

	url, hash := sample.String(10), sample.String(10)
	chain.InitialGenesis, _ = codec.NewAnyWithValue(&types.GenesisURL{url, hash})
	genUrl, err := chain.GetGenesisURL(cdc)
	require.NoError(t, err)
	require.Equal(t, &types.GenesisURL{
		Url:  url,
		Hash: hash,
	}, genUrl)

	chain.InitialGenesis, _ = codec.NewAnyWithValue(&types.DefaultInitialGenesis{})
	_, err = chain.GetGenesisURL(cdc)
	require.Error(t, err)
}

func TestChainIDFromChainName(t *testing.T) {
	require.Equal(t, "foo-1", types.ChainIDFromChainName("foo", 1))
}

func TestParseChainID(t *testing.T) {
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
				name, number, err := types.ParseChainID(tc.chainID)
				require.NoError(t, err)
				require.EqualValues(t, tc.expectedName, name)
				require.EqualValues(t, tc.expectedNumber, number)
			} else {
				_, _, err := types.ParseChainID(tc.chainID)
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
