package types_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestNewGenesisURL(t *testing.T) {
	url := spnmocks.MockRandomString(100)
	content := spnmocks.MockRandomString(100)

	genesisURL, err := types.NewGenesisURL(url, content)
	require.NoError(t, err)
	require.Equal(t, genesisURL.Url, url)
	require.Equal(t, genesisURL.Hash, types.GenesisURLHash(content))

	_, err = types.NewGenesisURL("", content)
	require.Error(t, err)
	_, err = types.NewGenesisURL(url, "")
	require.Error(t, err)
}

func TestGenesisURLHash(t *testing.T) {
	// Hash is deterministic
	content1 := spnmocks.MockRandomString(100)
	hash1 := types.GenesisURLHash(content1)
	hash2 := types.GenesisURLHash(content1)
	require.Len(t, hash1, 32)
	require.Equal(t, hash1, hash2)

	// Hash is unique
	content2 := spnmocks.MockRandomString(100)
	hash2 = types.GenesisURLHash(content2)
	require.Len(t, hash2, 32)
	require.NotEqual(t, hash1, hash2)
}