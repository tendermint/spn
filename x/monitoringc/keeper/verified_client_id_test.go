package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func createNVerifiedClientID(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VerifiedClientID {
	items := make([]types.VerifiedClientID, n)
	for i := range items {
		items[i].LaunchID = uint64(i)
		items[i].ClientIDs = []string{strconv.Itoa(i)}
		keeper.SetVerifiedClientID(ctx, items[i])
	}
	return items
}

func TestVerifiedClientIDGet(t *testing.T) {
	k, ctx := keepertest.Monitoringc(t)
	items := createNVerifiedClientID(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetVerifiedClientID(ctx, item.LaunchID)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestVerifiedClientIDGetAll(t *testing.T) {
	k, ctx := keepertest.Monitoringc(t)
	items := createNVerifiedClientID(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllVerifiedClientID(ctx)),
	)
}

func TestAddVerifiedClientID(t *testing.T) {
	k, ctx := keepertest.Monitoringc(t)
	t.Run("update a verified client id", func(t *testing.T) {
		var (
			launchID         = uint64(1)
			newClientID      = "2"
			verifiedClientID = types.VerifiedClientID{
				LaunchID:  launchID,
				ClientIDs: []string{"1"},
			}
		)
		k.SetVerifiedClientID(ctx, verifiedClientID)
		k.AddVerifiedClientID(ctx, launchID, newClientID)
		got, found := k.GetVerifiedClientID(ctx, launchID)
		require.True(t, found)
		verifiedClientID.ClientIDs = append(verifiedClientID.ClientIDs, newClientID)
		require.Equal(t, verifiedClientID, got)
	})

	t.Run("update a duplicated verified client id", func(t *testing.T) {
		var (
			launchID         = uint64(2)
			newClientID      = "2"
			verifiedClientID = types.VerifiedClientID{
				LaunchID:  launchID,
				ClientIDs: []string{"1", newClientID},
			}
		)
		k.SetVerifiedClientID(ctx, verifiedClientID)
		k.AddVerifiedClientID(ctx, launchID, newClientID)
		got, found := k.GetVerifiedClientID(ctx, launchID)
		require.True(t, found)
		require.Equal(t, verifiedClientID, got)
	})

	t.Run("update a non exiting verified client id", func(t *testing.T) {
		verifiedClientID := types.VerifiedClientID{
			LaunchID:  3,
			ClientIDs: []string{"1"},
		}
		k.AddVerifiedClientID(ctx, verifiedClientID.LaunchID, verifiedClientID.ClientIDs[0])
		got, found := k.GetVerifiedClientID(ctx, verifiedClientID.LaunchID)
		require.True(t, found)
		require.Equal(t, verifiedClientID, got)
	})
}
