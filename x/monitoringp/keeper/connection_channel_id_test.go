package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringp/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func createTestConnectionChannelID(keeper *keeper.Keeper, ctx sdk.Context) types.ConnectionChannelID {
	item := types.ConnectionChannelID{}
	keeper.SetConnectionChannelID(ctx, item)
	return item
}

func TestConnectionChannelIDGet(t *testing.T) {
	keeper, _, _, ctx := keepertest.MonitoringpKeeper(t)
	item := createTestConnectionChannelID(keeper, ctx)
	rst, found := keeper.GetConnectionChannelID(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}