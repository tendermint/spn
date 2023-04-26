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

func createNValidatorByOperatorAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ValidatorByOperatorAddress {
	items := make([]types.ValidatorByOperatorAddress, n)
	for i := range items {
		items[i].OperatorAddress = sample.Address(r)
		keeper.SetValidatorByOperatorAddress(ctx, items[i])
	}
	return items
}

func TestValidatorByOperatorAddressGet(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNValidatorByOperatorAddress(tk.ProfileKeeper, sdkCtx, 10)

	t.Run("should allow getting validator by operator address", func(t *testing.T) {
		for _, item := range items {
			rst, found := tk.ProfileKeeper.GetValidatorByOperatorAddress(sdkCtx,
				item.OperatorAddress,
			)
			require.True(t, found)
			require.Equal(t,
				nullify.Fill(&item),
				nullify.Fill(&rst),
			)
		}
	})
}

func TestValidatorByOperatorAddressRemove(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNValidatorByOperatorAddress(tk.ProfileKeeper, sdkCtx, 10)

	t.Run("should allow removing validator by operator address", func(t *testing.T) {
		for _, item := range items {
			tk.ProfileKeeper.RemoveValidatorByOperatorAddress(sdkCtx,
				item.OperatorAddress,
			)
			_, found := tk.ProfileKeeper.GetValidatorByOperatorAddress(sdkCtx,
				item.OperatorAddress,
			)
			require.False(t, found)
		}
	})
}

func TestValidatorByOperatorAddressGetAll(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNValidatorByOperatorAddress(tk.ProfileKeeper, sdkCtx, 10)

	t.Run("should allow getting all validator by operator address", func(t *testing.T) {
		require.ElementsMatch(t,
			nullify.Fill(items),
			nullify.Fill(tk.ProfileKeeper.GetAllValidatorByOperatorAddress(sdkCtx)),
		)
	})
}
