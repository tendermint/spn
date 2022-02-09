package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func TestCoordinatorAddrNotFoundInvariant(t *testing.T) {
	ctx, k, _ := setupMsgServer(t)
	t.Run("valid case", func(t *testing.T) {
		coord := sample.Coordinator(sample.Address())
		coord.CoordinatorID = k.AppendCoordinator(ctx, coord)
		k.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
			Address:       sample.Address(),
			CoordinatorID: coord.CoordinatorID,
		})
		_, isValid := keeper.CoordinatorAddrNotFoundInvariant(*k)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case 1", func(t *testing.T) {
		k.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
			Address:       sample.Address(),
			CoordinatorID: 10,
		})
		_, isValid := keeper.CoordinatorAddrNotFoundInvariant(*k)(ctx)
		require.Equal(t, true, isValid)
	})
}
