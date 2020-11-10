package genesis_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/genesis"
	"github.com/tendermint/spn/x/genesis/types"
	"testing"
)

func TestHandleMsgChainCreate(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	h := genesis.NewHandler(*k)

	chainID := spnmocks.MockRandomAlphaString(5)
	creator := spnmocks.MockAccAddress()
	sourceURL := spnmocks.MockRandomString(20)
	sourceHash := spnmocks.MockRandomString(20)
	genesis := spnmocks.MockGenesis()

	// A chain can be create
	msg := types.NewMsgChainCreate(
		chainID,
		creator,
		sourceURL,
		sourceHash,
		genesis,
		)
	_, err := h(ctx, msg)
	require.NoError(t, err, "NewMsgChainCreate with a correct chain should succeed")
	retrieved, found := k.GetChain(ctx, chainID)
	require.True(t, found, "NewMsgChainCreate should add the chain in the store")
	creatorIdentity, _ := k.IdentityKeeper.GetIdentifier(ctx, creator)
	require.Equal(t, creatorIdentity, retrieved.Creator, "NewMsgChainCreate should add the correct chain")
	require.Equal(t, sourceURL, retrieved.SourceURL, "NewMsgChainCreate should add the correct chain")
	require.Equal(t, sourceHash, retrieved.SourceHash, "NewMsgChainCreate should add the correct chain")

	// Prevent adding an existing chain id
	msg = types.NewMsgChainCreate(
		chainID,
		creator,
		sourceURL,
		sourceHash,
		genesis,
	)
	_, err = h(ctx, msg)
	require.Error(t, err, "NewMsgChainCreate should prevent adding an existing chain")
}