package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestNewDefaultInitialGenesis(t *testing.T) {
	t.Run("should return a default initial genesis", func(t *testing.T) {
		initialGenesis := types.NewDefaultInitialGenesis()

		initialDefaultGenesis := initialGenesis.GetDefaultInitialGenesis()
		require.NotNil(t, initialDefaultGenesis)

		genesisURL := initialGenesis.GetGenesisURL()
		require.Nil(t, genesisURL)
	})
}

func TestNewGenesisURL(t *testing.T) {
	t.Run("should return a genesis URL", func(t *testing.T) {
		url := sample.String(r, 30)
		hash := sample.GenesisHash(r)
		initialGenesis := types.NewGenesisURL(url, hash)

		genesisURL := initialGenesis.GetGenesisURL()
		require.NotNil(t, genesisURL)
		require.EqualValues(t, url, genesisURL.Url)
		require.EqualValues(t, hash, genesisURL.Hash)

		initialDefaultGenesis := initialGenesis.GetDefaultInitialGenesis()
		require.Nil(t, initialDefaultGenesis)
	})
}

func TestDefaultInitialGenesis_Validate(t *testing.T) {
	t.Run("should validate any default initial genesis", func(t *testing.T) {
		require.NoError(t, types.NewDefaultInitialGenesis().Validate(),
			" DefaultInitialGenesis should always be valid",
		)
	})
}

func TestGenesisURL_Validate(t *testing.T) {
	t.Run("should validate genesis URL", func(t *testing.T) {
		require.NoError(t, types.NewGenesisURL(sample.String(r, 30), sample.GenesisHash(r)).Validate(),
			" GenesisURL should be valid",
		)
	})

	t.Run("should prevent validate invalid genesis URL", func(t *testing.T) {
		require.Error(t, types.NewGenesisURL("", sample.GenesisHash(r)).Validate(),
			" GenesisURL must contain a url",
		)

		require.Error(t, types.NewGenesisURL(sample.String(r, 30), sample.String(r, types.HashLength-1)).Validate(),
			" GenesisURL must contain a valid sha256 hash",
		)
	})
}

func TestGenesisURLHash(t *testing.T) {
	t.Run("should return valid SHA256 hash", func(t *testing.T) {
		require.EqualValues(t,
			"2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae",
			types.GenesisURLHash("foo"),
		)
	})
}
