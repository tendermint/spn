package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestGetParams(t *testing.T) {
	ctx, tk := testkeeper.NewTestKeepers(t)

	params := types.DefaultParams()
	tk.MonitoringConsumerKeeper.SetParams(ctx, params)
	require.EqualValues(t, params, tk.MonitoringConsumerKeeper.GetParams(ctx))
	require.EqualValues(t, false, tk.MonitoringConsumerKeeper.DebugMode(ctx))

	params = types.NewParams(true)
	tk.MonitoringConsumerKeeper.SetParams(ctx, params)
	require.EqualValues(t, params, tk.MonitoringConsumerKeeper.GetParams(ctx))
	require.EqualValues(t, true, tk.MonitoringConsumerKeeper.DebugMode(ctx))
}
