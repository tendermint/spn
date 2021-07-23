package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/tendermint/spn/x/profile/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNConsensusKeyNonce(keeper *Keeper, ctx sdk.Context, n int) []types.ConsensusKeyNonce {
	items := make([]types.ConsensusKeyNonce, n)
	for i := range items {
		items[i].ConsAddress = strconv.Itoa(i)

		keeper.SetConsensusKeyNonce(ctx, items[i])
	}
	return items
}

func TestConsensusKeyNonceGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNConsensusKeyNonce(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetConsensusKeyNonce(ctx,
			item.ConsAddress,
		)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestConsensusKeyNonceRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNConsensusKeyNonce(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveConsensusKeyNonce(ctx,
			item.ConsAddress,
		)
		_, found := keeper.GetConsensusKeyNonce(ctx,
			item.ConsAddress,
		)
		assert.False(t, found)
	}
}

func TestConsensusKeyNonceGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNConsensusKeyNonce(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllConsensusKeyNonce(ctx))
}
