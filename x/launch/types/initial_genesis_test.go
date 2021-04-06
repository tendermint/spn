package types_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestNewInitialGenesisDefault(t *testing.T) {
	initialGenesis := types.NewInitialGenesisDefault()

	g := initialGenesis.GetDefaultGenesis()
	require.NotNil(t, g)

	gu := initialGenesis.GetGenesisURL()
	require.Nil(t, gu)
}

func TestNewInitialGenesisURL(t *testing.T) {
	url := spnmocks.MockRandomString(100)
	hash := spnmocks.MockRandomString(types.HashLength)
	genesisURL, err := types.NewGenesisURL(url, hash)
	require.NoError(t, err)

	initialGenesis := types.NewInitialGenesisURL(genesisURL)
	gu := initialGenesis.GetGenesisURL()
	require.NotNil(t, gu)
	require.Equal(t, genesisURL, *gu)
}
