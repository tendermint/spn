package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
)

func Test_GetParams(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	params := sample.ProjectParams(r)

	t.Run("should set and get params", func(t *testing.T) {
		tk.ProjectKeeper.SetParams(ctx, params)
		require.EqualValues(t, params, tk.ProjectKeeper.GetParams(ctx))
		require.EqualValues(t, params.TotalSupplyRange.MinTotalSupply, tk.ProjectKeeper.TotalSupplyRange(ctx).MinTotalSupply)
		require.EqualValues(t, params.TotalSupplyRange.MaxTotalSupply, tk.ProjectKeeper.TotalSupplyRange(ctx).MaxTotalSupply)
		require.EqualValues(t, params.ProjectCreationFee, tk.ProjectKeeper.ProjectCreationFee(ctx))
	})
}
