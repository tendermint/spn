package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
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
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNProviderClientID(tk.MonitoringConsumerKeeper, ctx, 10)
	for _, item := range items {
		rst, found := tk.MonitoringConsumerKeeper.GetProviderClientID(ctx,
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
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNProviderClientID(tk.MonitoringConsumerKeeper, ctx, 10)
	for _, item := range items {
		tk.MonitoringConsumerKeeper.RemoveProviderClientID(ctx,
			item.LaunchID,
		)
		_, found := tk.MonitoringConsumerKeeper.GetProviderClientID(ctx,
			item.LaunchID,
		)
		require.False(t, found)
	}
}

func TestProviderClientIDGetAll(t *testing.T) {
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNProviderClientID(tk.MonitoringConsumerKeeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(tk.MonitoringConsumerKeeper.GetAllProviderClientID(ctx)),
	)
}
