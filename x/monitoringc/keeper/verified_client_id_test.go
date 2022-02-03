package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func createNVerifiedClientID(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VerifiedClientID {
	items := make([]types.VerifiedClientID, n)
	for i := range items {
		items[i].LaunchID = uint64(i)
		items[i].ClientIDs = []string{strconv.Itoa(i)}
		keeper.SetVerifiedClientID(ctx, items[i])
	}
	return items
}

func TestVerifiedClientIDGet(t *testing.T) {
	k, ctx := keepertest.Monitoringc(t)
	items := createNVerifiedClientID(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetVerifiedClientID(ctx, item.LaunchID)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestVerifiedClientIDRemove(t *testing.T) {
	k, ctx := keepertest.Monitoringc(t)
	items := createNVerifiedClientID(k, ctx, 10)
	for _, item := range items {
		k.RemoveVerifiedClientID(ctx, item.LaunchID)
		_, found := k.GetVerifiedClientID(ctx, item.LaunchID)
		require.False(t, found)
	}
}

func TestVerifiedClientIDGetAll(t *testing.T) {
	k, ctx := keepertest.Monitoringc(t)
	items := createNVerifiedClientID(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllVerifiedClientID(ctx)),
	)
}
