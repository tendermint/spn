package types_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestNewGenesisURL(t *testing.T) {
	url := spnmocks.MockRandomString(100)
	hash := spnmocks.MockRandomString(types.HashLength)

	genesisURL, err := types.NewGenesisURL(url, hash)
	require.NoError(t, err)
	require.Equal(t, genesisURL.Url, url)
	require.Equal(t, genesisURL.Hash, hash)

	_, err = types.NewGenesisURL("", hash)
	require.Error(t, err)
	_, err = types.NewGenesisURL(url, spnmocks.MockRandomString(types.HashLength+1))
	require.Error(t, err)
}

func TestNewGenesisURLFromContent(t *testing.T) {
	url := spnmocks.MockRandomString(100)
	content := spnmocks.MockRandomString(100)

	genesisURL, err := types.NewGenesisURLFromContent(url, content)
	require.NoError(t, err)
	require.Equal(t, genesisURL.Url, url)
	require.Equal(t, genesisURL.Hash, types.GenesisURLHash(content))

	_, err = types.NewGenesisURLFromContent("", content)
	require.Error(t, err)
	_, err = types.NewGenesisURLFromContent(url, "")
	require.Error(t, err)
}

func TestGenesisURLHash(t *testing.T) {
	// Hash is deterministic
	content1 := spnmocks.MockRandomString(100)
	hash1 := types.GenesisURLHash(content1)
	hash2 := types.GenesisURLHash(content1)
	require.Len(t, hash1, types.HashLength)
	require.Equal(t, hash1, hash2)

	// Hash is unique
	content2 := spnmocks.MockRandomString(100)
	hash2 = types.GenesisURLHash(content2)
	require.Len(t, hash2, types.HashLength)
	require.NotEqual(t, hash1, hash2)
}
