package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringp/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func createTestConnectionChannelID(ctx sdk.Context, keeper *keeper.Keeper) types.ConnectionChannelID {
	item := types.ConnectionChannelID{}
	keeper.SetConnectionChannelID(ctx, item)
	return item
}

func TestConnectionChannelIDGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)

	t.Run("should allow get", func(t *testing.T) {
		item := createTestConnectionChannelID(ctx, tk.MonitoringProviderKeeper)
		rst, found := tk.MonitoringProviderKeeper.GetConnectionChannelID(ctx)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	})
}
