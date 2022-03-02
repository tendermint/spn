package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func createNValidatorByConsAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ValidatorByConsAddress {
	items := make([]types.ValidatorByConsAddress, n)
	for i := range items {
		items[i].ConsensusAddress = sample.ConsAddress().Bytes()
		keeper.SetValidatorByConsAddress(ctx, items[i])
	}
	return items
}

func TestValidatorByConsAddressGet(t *testing.T) {
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNValidatorByConsAddress(tk.ProfileKeeper, ctx, 10)
	for _, item := range items {
		rst, found := tk.ProfileKeeper.GetValidatorByConsAddress(ctx,
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
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNValidatorByConsAddress(tk.ProfileKeeper, ctx, 10)
	for _, item := range items {
		tk.ProfileKeeper.RemoveValidatorByConsAddress(ctx,
			item.ConsensusAddress,
		)
		_, found := tk.ProfileKeeper.GetValidatorByConsAddress(ctx,
			item.ConsensusAddress,
		)
		require.False(t, found)
	}
}

func TestValidatorByConsAddressGetAll(t *testing.T) {
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNValidatorByConsAddress(tk.ProfileKeeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(tk.ProfileKeeper.GetAllValidatorByConsAddress(ctx)),
	)
}
