package keeper_test

import (
	"testing"
	
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
)

func Test_GetParams(t *testing.T) {
	k, _, ctx, _ := testkeeper.Launch(t)
	params := sample.LaunchParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.MaxLaunchTime, k.MaxLaunchTime(ctx))
	require.EqualValues(t, params.MinLaunchTime, k.MinLaunchTime(ctx))
}