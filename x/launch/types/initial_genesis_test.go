package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestNewDefaultInitialGenesis(t *testing.T) {
	initialGenesis := types.NewDefaultInitialGenesis()

	initialDefaultGenesis := initialGenesis.GetDefaultInitialGenesis()
	require.NotNil(t, initialDefaultGenesis)

	genesisURL := initialGenesis.GetGenesisURL()
	require.Nil(t, genesisURL)
}

func TestNewGenesisURL(t *testing.T) {
	url := sample.String(30)
	hash := sample.GenesisHash()
	initialGenesis := types.NewGenesisURL(url, hash)

	genesisURL := initialGenesis.GetGenesisURL()
	require.NotNil(t, genesisURL)
	require.EqualValues(t, url, genesisURL.Url)
	require.EqualValues(t, hash, genesisURL.Hash)

	initialDefaultGenesis := initialGenesis.GetDefaultInitialGenesis()
	require.Nil(t, initialDefaultGenesis)
}

func TestDefaultInitialGenesis_Validate(t *testing.T) {
	require.NoError(t, types.NewDefaultInitialGenesis().Validate(),
		" DefaultInitialGenesis should always be valid",
	)
}

func TestGenesisURL_Validate(t *testing.T) {
	require.NoError(t, types.NewGenesisURL(sample.String(30), sample.GenesisHash()).Validate(),
		" GenesisURL should be valid",
	)

	require.Error(t, types.NewGenesisURL("", sample.GenesisHash()).Validate(),
		" GenesisURL must contain a url",
	)

	require.Error(t, types.NewGenesisURL(sample.String(30), sample.String(types.HashLength - 1)).Validate(),
		" GenesisURL must contain a valid sha256 hash",
	)
}

func TestGenesisURLHash(t *testing.T) {
	require.EqualValues(t,
		"2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae",
		types.GenesisURLHash("foo"),
	)
}
