package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func createNConsensusKeyNonce(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ConsensusKeyNonce {
	items := make([]types.ConsensusKeyNonce, n)
	for i := range items {
		items[i].ConsensusAddress = sample.ConsAddress().Bytes()
		keeper.SetConsensusKeyNonce(ctx, items[i])
	}
	return items
}

func TestConsensusKeyNonceGet(t *testing.T) {
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNConsensusKeyNonce(tk.ProfileKeeper, ctx, 10)
	for _, item := range items {
		rst, found := tk.ProfileKeeper.GetConsensusKeyNonce(ctx,
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
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNConsensusKeyNonce(tk.ProfileKeeper, ctx, 10)
	for _, item := range items {
		tk.ProfileKeeper.RemoveConsensusKeyNonce(ctx,
			item.ConsensusAddress,
		)
		_, found := tk.ProfileKeeper.GetConsensusKeyNonce(ctx,
			item.ConsensusAddress,
		)
		require.False(t, found)
	}
}
