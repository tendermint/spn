package types_test

import (
	"encoding/json"
	tmtypes "github.com/tendermint/tendermint/types"
	"testing"
	"time"

	spnmocks "github.com/tendermint/spn/internal/testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/genesis/types"
)

func TestNewChain(t *testing.T) {
	// Can create a chain
	chain, err := types.NewChain(
		spnmocks.MockRandomString(5)+"-"+spnmocks.MockRandomString(5),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		time.Now(),
		spnmocks.MockGenesis(),
	)
	require.NoError(t, err)
	require.Equal(t, 0, len(chain.Peers))

	// The chain ID should be added to the genesis of the chain
	chainID := chain.ChainID
	genesisDoc, err := chain.Genesis.GetGenesisDoc()
	require.NoError(t, err)
	require.Equal(t, chainID, genesisDoc.ChainID)

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
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		time.Now(),
		spnmocks.MockGenesis(),
	)
	require.Error(t, err)

	// Prevent creating a chain with a invalid genesis
	_, err = types.NewChain(
		spnmocks.MockRandomString(5)+"_"+spnmocks.MockRandomString(5),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		time.Now(),
		[]byte(spnmocks.MockRandomString(500)),
	)
	require.Error(t, err)

	var genesisObject tmtypes.GenesisDoc
	genesisObject.ConsensusParams = tmtypes.DefaultConsensusParams()
	genesisObject.ChainID = ""
	genesis, err := json.Marshal(genesisObject)
	if err != nil {
		panic("Cannot marshal genesis")
	}
	_, err = types.NewChain(
		spnmocks.MockRandomString(5)+"_"+spnmocks.MockRandomString(5),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		time.Now(),
		genesis,
	)
	require.Error(t, err)
}
