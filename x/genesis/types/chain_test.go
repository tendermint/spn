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
		[]byte(spnmocks.MockRandomString(2000)),
	)
	require.NoError(t, err, "NewChain should create a new chain")
	require.Equal(t, 0, len(chain.Peers), "chain should have no peer when create")

	// Can append peers to the chain
	peer1 := spnmocks.MockRandomString(20)
	peer2 := spnmocks.MockRandomString(20)
	chain.AppendPeer(peer1)
	chain.AppendPeer(peer2)
	require.Equal(t, 2, len(chain.Peers), "AppendPeer should append new peers to the chain")
	require.Equal(t, []string{peer1, peer2}, chain.Peers, "AppendPeer should append new peers to the chain")

	// Prevent creating a chain with a invalid name
	_, err = types.NewChain(
		spnmocks.MockRandomString(5)+"_"+spnmocks.MockRandomString(5),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		spnmocks.MockRandomString(20),
		time.Now(),
		[]byte(spnmocks.MockRandomString(2000)),
	)
	require.Error(t, err, "NewChain should prevent creating chains with an invalid name")
}
