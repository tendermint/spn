package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func TestGetParams(t *testing.T) {
	ctx, tk := testkeeper.NewTestKeepers(t)
	params := types.DefaultParams()

	tk.RewardKeeper.SetParams(ctx, params)

	require.EqualValues(t, params, tk.RewardKeeper.GetParams(ctx))
}
