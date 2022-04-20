package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func TestCoordinatorAddrNotFoundInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		coord := sample.Coordinator(r, sample.Address(r))
		coord.CoordinatorID = tk.ProfileKeeper.AppendCoordinator(ctx, coord)
		tk.ProfileKeeper.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
			Address:       sample.Address(r),
			CoordinatorID: coord.CoordinatorID,
		})
		msg, broken := keeper.CoordinatorAddrNotFoundInvariant(*tk.ProfileKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("invalid case", func(t *testing.T) {
		tk.ProfileKeeper.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
			Address:       sample.Address(r),
			CoordinatorID: 10,
		})
		msg, broken := keeper.CoordinatorAddrNotFoundInvariant(*tk.ProfileKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
