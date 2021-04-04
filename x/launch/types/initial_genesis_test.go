package types_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestNewInitialGenesisDefault(t *testing.T) {
	initialGenesis := types.NewInitialGenesisDefault()

	gt, err := initialGenesis.GetType()
	require.NoError(t, err)
	require.Equal(t, types.InitialGenesisType_DEFAULT, gt)

	_, err = initialGenesis.GenesisURL()
	require.Error(t, err)
}

func TestNewInitialGenesisURL(t *testing.T) {
	url := spnmocks.MockRandomString(100)
	hash := spnmocks.MockRandomString(types.HashLength)
	genesisURL, err := types.NewGenesisURL(url, hash)
	require.NoError(t, err)

	initialGenesis := types.NewInitialGenesisURL(genesisURL)
	gt, err := initialGenesis.GetType()
	require.NoError(t, err)
	require.Equal(t, types.InitialGenesisType_URL, gt)

	gu, err := initialGenesis.GenesisURL()
	require.NoError(t, err)
	require.Equal(t, genesisURL, gu)
}
