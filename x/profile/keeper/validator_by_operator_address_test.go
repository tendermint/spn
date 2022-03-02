package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func createNValidatorByOperatorAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ValidatorByOperatorAddress {
	items := make([]types.ValidatorByOperatorAddress, n)
	for i := range items {
		items[i].OperatorAddress = sample.Address()
		keeper.SetValidatorByOperatorAddress(ctx, items[i])
	}
	return items
}

func TestValidatorByOperatorAddressGet(t *testing.T) {
	k, ctx := keepertest.Profile(t)
	items := createNValidatorByOperatorAddress(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetValidatorByOperatorAddress(ctx,
			item.OperatorAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestValidatorByOperatorAddressRemove(t *testing.T) {
	k, ctx := keepertest.Profile(t)
	items := createNValidatorByOperatorAddress(k, ctx, 10)
	for _, item := range items {
		k.RemoveValidatorByOperatorAddress(ctx,
			item.OperatorAddress,
		)
		_, found := k.GetValidatorByOperatorAddress(ctx,
			item.OperatorAddress,
		)
		require.False(t, found)
	}
}

func TestValidatorByOperatorAddressGetAll(t *testing.T) {
	k, ctx := keepertest.Profile(t)
	items := createNValidatorByOperatorAddress(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllValidatorByOperatorAddress(ctx)),
	)
}
