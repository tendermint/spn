package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNValidatorByConsAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ValidatorByConsAddress {
	items := make([]types.ValidatorByConsAddress, n)
	for i := range items {
		items[i].ConsensusAddress = strconv.Itoa(i)

		keeper.SetValidatorByConsAddress(ctx, items[i])
	}
	return items
}

func TestValidatorByConsAddressGet(t *testing.T) {
	keeper, ctx := keepertest.Profile(t)
	items := createNValidatorByConsAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidatorByConsAddress(ctx,
			item.ConsensusAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestValidatorByConsAddressRemove(t *testing.T) {
	keeper, ctx := keepertest.Profile(t)
	items := createNValidatorByConsAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidatorByConsAddress(ctx,
			item.ConsensusAddress,
		)
		_, found := keeper.GetValidatorByConsAddress(ctx,
			item.ConsensusAddress,
		)
		require.False(t, found)
	}
}

func TestValidatorByConsAddressGetAll(t *testing.T) {
	keeper, ctx := keepertest.Profile(t)
	items := createNValidatorByConsAddress(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllValidatorByConsAddress(ctx)),
	)
}
