package types_test

import (
	"testing"
	"time"

	spnmocks "github.com/tendermint/spn/internal/testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/launch/types"
)

func TestNewChain(t *testing.T) {
	chainID := spnmocks.MockRandomString(5)+"-"+spnmocks.MockRandomString(5)
	creator := spnmocks.MockRandomString(20)
	sourceURL := spnmocks.MockRandomString(20)
		sourceHash := spnmocks.MockRandomString(20)
		creationTime := time.Now()

	// Can create a chain
	chain, err := types.NewChain(
		chainID,
		creator,
		sourceURL,
		sourceHash,
		creationTime,
		"",
		"",
	)
	require.NoError(t, err)
	require.Equal(t, chainID, chain.ChainID)
	require.Equal(t, creator, chain.Creator)
	require.Equal(t, sourceURL, chain.SourceURL)
	require.Equal(t, sourceHash, chain.SourceHash)
	require.Equal(t, creationTime.Unix(), chain.CreatedAt)
	require.Equal(t, 0, len(chain.Peers))
	initialGenesisType, err := chain.InitialGenesis.GetType()
	require.NoError(t, err)
	require.Equal(t, types.InitialGenesisType_DEFAULT, initialGenesisType)

	// Can append peers to the chain
	peer1 := spnmocks.MockRandomString(20)
	peer2 := spnmocks.MockRandomString(20)
	chain.AppendPeer(peer1)
	chain.AppendPeer(peer2)
	require.Equal(t, 2, len(chain.Peers))
	require.Equal(t, []string{peer1, peer2}, chain.Peers)

	// Prevent creating a chain with a invalid name
	_, err = types.NewChain(
		spnmocks.MockRandomString(5)+"_"+spnmocks.MockRandomString(5),
		creator,
		sourceURL,
		sourceHash,
		creationTime,
		"",
		"",
	)
	require.Error(t, err)

	// Can create a chain with a custom genesis
	genesisURL, err := types.NewGenesisURL(
		spnmocks.MockRandomString(100),
		spnmocks.MockRandomString(32),
	)
	require.NoError(t, err)
	chain, err = types.NewChain(
		chainID,
		creator,
		sourceURL,
		sourceHash,
		creationTime,
		genesisURL.Url,
		genesisURL.Hash,
	)
	require.NoError(t, err)
	initialGenesisType, err = chain.InitialGenesis.GetType()
	require.NoError(t, err)
	require.Equal(t, types.InitialGenesisType_URL, initialGenesisType)
	chainGenesisURL, err := chain.InitialGenesis.GenesisURL()
	require.NoError(t, err)
	require.Equal(t, genesisURL, chainGenesisURL)
}
