package types_test

import (
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

	// Prevent creating a chain with a empty genesis
	_, err = types.NewChain(
		spnmocks.MockRandomString(5)+"_"+spnmocks.MockRandomString(5),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		time.Now(),
		[]byte{},
	)
	require.Error(t, err)
}
