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

func createTestConsumerClientID(ctx sdk.Context, keeper *keeper.Keeper) types.ConsumerClientID {
	item := types.ConsumerClientID{}
	keeper.SetConsumerClientID(ctx, item)
	return item
}

func TestConsumerClientIDGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)
	item := createTestConsumerClientID(ctx, tk.MonitoringProviderKeeper)
	rst, found := tk.MonitoringProviderKeeper.GetConsumerClientID(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
