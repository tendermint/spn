package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Monitoringc(t)

	params := types.DefaultParams()
	k.SetParams(ctx, params)
	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, false, k.DebugMode(ctx))

	params = types.NewParams(true)
	k.SetParams(ctx, params)
	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, true, k.DebugMode(ctx))
}
