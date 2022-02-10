package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func createNProviderClientID(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ProviderClientID {
	items := make([]types.ProviderClientID, n)
	for i := range items {
		items[i].LaunchID = uint64(i)
		keeper.SetProviderClientID(ctx, items[i])
	}
	return items
}

func TestProviderClientIDGet(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	items := createNProviderClientID(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProviderClientID(ctx,
			item.LaunchID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestProviderClientIDRemove(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	items := createNProviderClientID(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProviderClientID(ctx,
			item.LaunchID,
		)
		_, found := keeper.GetProviderClientID(ctx,
			item.LaunchID,
		)
		require.False(t, found)
	}
}

func TestProviderClientIDGetAll(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	items := createNProviderClientID(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProviderClientID(ctx)),
	)
}
