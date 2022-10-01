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

func createNProviderClientID(ctx sdk.Context, keeper *keeper.Keeper, n int) []types.ProviderClientID {
	items := make([]types.ProviderClientID, n)
	for i := range items {
		items[i].LaunchID = uint64(i)
		keeper.SetProviderClientID(ctx, items[i])
	}
	return items
}

func TestProviderClientIDGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get", func(t *testing.T) {
		items := createNProviderClientID(ctx, tk.MonitoringConsumerKeeper, 10)
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
	})
}

func TestProviderClientIDGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get all", func(t *testing.T) {
		items := createNProviderClientID(ctx, tk.MonitoringConsumerKeeper, 10)
		require.ElementsMatch(t,
			nullify.Fill(items),
			nullify.Fill(tk.MonitoringConsumerKeeper.GetAllProviderClientID(ctx)),
		)
	})
}
