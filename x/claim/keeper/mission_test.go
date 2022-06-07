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

func createNMission(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mission {
	items := make([]types.Mission, n)
	for i := range items {
		items[i].MissionID = uint64(i)
		items[i].Weight = sdk.NewDec(r.Int63())
		keeper.SetMission(ctx, items[i])
	}
	return items
}

func TestMissionGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	items := createNMission(tk.ClaimKeeper, ctx, 10)
	for _, item := range items {
		got, found := tk.ClaimKeeper.GetMission(ctx, item.MissionID)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMissionGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	items := createNMission(tk.ClaimKeeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(tk.ClaimKeeper.GetAllMission(ctx)),
	)
}

func TestMissionRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	items := createNMission(tk.ClaimKeeper, ctx, 10)
	for _, item := range items {
		tk.ClaimKeeper.RemoveMission(ctx, item.MissionID)
		_, found := tk.ClaimKeeper.GetMission(ctx, item.MissionID)
		require.False(t, found)
	}
}

func TestKeeper_CompleteMission(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	tests := []struct {
		name            string
		missionID       uint64
		address         string
		expectedBalance sdk.Coins
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
