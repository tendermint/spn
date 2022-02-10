package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func createNConsensusKeyNonce(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ConsensusKeyNonce {
	items := make([]types.ConsensusKeyNonce, n)
	for i := range items {
		items[i].ConsensusAddress = []byte(strconv.Itoa(i))
		keeper.SetConsensusKeyNonce(ctx, items[i])
	}
	return items
}

func TestConsensusKeyNonceGet(t *testing.T) {
	keeper, ctx := keepertest.Profile(t)
	items := createNConsensusKeyNonce(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetConsensusKeyNonce(ctx,
			item.ConsensusAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestConsensusKeyNonceRemove(t *testing.T) {
	keeper, ctx := keepertest.Profile(t)
	items := createNConsensusKeyNonce(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveConsensusKeyNonce(ctx,
			item.ConsensusAddress,
		)
		_, found := keeper.GetConsensusKeyNonce(ctx,
			item.ConsensusAddress,
		)
		require.False(t, found)
	}
}
