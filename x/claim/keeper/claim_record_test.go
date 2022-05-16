package keeper_test

import (
	"github.com/tendermint/spn/testutil/sample"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/claim/keeper"
	"github.com/tendermint/spn/x/claim/types"
)

func createNClaimRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ClaimRecord {
	items := make([]types.ClaimRecord, n)
	for i := range items {
		items[i].Address = sample.Address(r)

		keeper.SetClaimRecord(ctx, items[i])
	}
	return items
}

func TestClaimRecordGet(t *testing.T) {
	k, ctx := keepertest.ClaimKeeper(t)
	items := createNClaimRecord(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetClaimRecord(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestClaimRecordRemove(t *testing.T) {
	k, ctx := keepertest.ClaimKeeper(t)
	items := createNClaimRecord(k, ctx, 10)
	for _, item := range items {
		k.RemoveClaimRecord(ctx,
			item.Address,
		)
		_, found := k.GetClaimRecord(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestClaimRecordGetAll(t *testing.T) {
	k, ctx := keepertest.ClaimKeeper(t)
	items := createNClaimRecord(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllClaimRecord(ctx)),
	)
}
