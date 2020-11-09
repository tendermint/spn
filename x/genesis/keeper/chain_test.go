package keeper_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"testing"
)

func TestGetChain(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chain := spnmocks.MockChain()

	// A non set chain should not exist
	_, found := k.GetChain(ctx, chain.ChainID)
	require.False(t, found, "GetChain should not find a non existent chain")

	// Set and get a chain
	k.SetChain(ctx, *chain)
	retrieved, found := k.GetChain(ctx, chain.ChainID)
	require.True(t, found, "GetChain should find a chain")
	require.Equal(t, *chain, retrieved, "GetChain should find a chain")

	// Can get all the chain
	chain2 := spnmocks.MockChain()
	chain3 := spnmocks.MockChain()
	chain4 := spnmocks.MockChain()
	chain5 := spnmocks.MockChain()
	chain6 := spnmocks.MockChain()
	k.SetChain(ctx, *chain2)
	k.SetChain(ctx, *chain3)
	k.SetChain(ctx, *chain4)
	k.SetChain(ctx, *chain5)
	k.SetChain(ctx, *chain6)
	allChains := k.GetAllChains(ctx)
	require.Equal(t, 6, len(allChains), "GetAllChains should retrieve all chains")
	require.Contains(t, allChains, *chain, "GetAllChains should retrieve all chains")
	require.Contains(t, allChains, *chain2, "GetAllChains should retrieve all chains")
	require.Contains(t, allChains, *chain3, "GetAllChains should retrieve all chains")
	require.Contains(t, allChains, *chain4, "GetAllChains should retrieve all chains")
	require.Contains(t, allChains, *chain5, "GetAllChains should retrieve all chains")
	require.Contains(t, allChains, *chain6, "GetAllChains should retrieve all chains")
}
