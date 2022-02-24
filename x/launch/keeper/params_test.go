package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
)

func Test_GetParams(t *testing.T) {
	k, ctx := testkeeper.Launch(t)
	params := sample.LaunchParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.LaunchTimeRange.MaxLaunchTime, k.LaunchTimeRange(ctx).MaxLaunchTime)
	require.EqualValues(t, params.LaunchTimeRange.MinLaunchTime, k.LaunchTimeRange(ctx).MinLaunchTime)
	require.EqualValues(t, params.RevertDelay, k.RevertDelay(ctx))
}
