package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
)

func Test_GetParams(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	params := sample.LaunchParams()

	tk.LaunchKeeper.SetParams(ctx, params)

	require.EqualValues(t, params, tk.LaunchKeeper.GetParams(ctx))
	require.EqualValues(t, params.LaunchTimeRange.MaxLaunchTime, tk.LaunchKeeper.LaunchTimeRange(ctx).MaxLaunchTime)
	require.EqualValues(t, params.LaunchTimeRange.MinLaunchTime, tk.LaunchKeeper.LaunchTimeRange(ctx).MinLaunchTime)
	require.EqualValues(t, params.RevertDelay, tk.LaunchKeeper.RevertDelay(ctx))
	require.EqualValues(t, params.ChainCreationFee, tk.LaunchKeeper.ChainCreationFee(ctx))
}
