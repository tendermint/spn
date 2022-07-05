package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/claim/keeper"
	"github.com/tendermint/spn/x/claim/types"
)

func createTestInitialClaim(keeper *keeper.Keeper, ctx sdk.Context) types.InitialClaim {
	item := types.InitialClaim{}
	keeper.SetInitialClaim(ctx, item)
	return item
}

func TestInitialClaimGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	item := createTestInitialClaim(tk.ClaimKeeper, ctx)
	rst, found := tk.ClaimKeeper.GetInitialClaim(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestInitialClaimRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	createTestInitialClaim(tk.ClaimKeeper, ctx)
	tk.ClaimKeeper.RemoveInitialClaim(ctx)
	_, found := tk.ClaimKeeper.GetInitialClaim(ctx)
	require.False(t, found)
}
