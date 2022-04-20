package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func createNVerifiedClientID(ctx sdk.Context, keeper *keeper.Keeper, n int) []types.VerifiedClientID {
	items := make([]types.VerifiedClientID, n)
	for i := range items {
		items[i].LaunchID = uint64(i)
		items[i].ClientIDs = []string{strconv.Itoa(i)}
		keeper.SetVerifiedClientID(ctx, items[i])
	}
	return items
}

func TestVerifiedClientIDGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNVerifiedClientID(ctx, tk.MonitoringConsumerKeeper, 10)
	for _, item := range items {
		rst, found := tk.MonitoringConsumerKeeper.GetVerifiedClientID(ctx, item.LaunchID)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestVerifiedClientIDClear(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("successfully clear entries", func(t *testing.T) {
		items := createNVerifiedClientID(ctx, tk.MonitoringConsumerKeeper, 1)
		launchID := items[0].LaunchID
		clientID := items[0].ClientIDs[0]

		tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			ClientID: clientID,
			LaunchID: launchID,
		})
		rst, found := tk.MonitoringConsumerKeeper.GetVerifiedClientID(ctx, launchID)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&items[0]),
			nullify.Fill(&rst),
		)

		tk.MonitoringConsumerKeeper.ClearVerifiedClientIDs(ctx, launchID)
		_, found = tk.MonitoringConsumerKeeper.GetVerifiedClientID(ctx, launchID)
		require.False(t, found)

		_, found = tk.MonitoringConsumerKeeper.GetLaunchIDFromVerifiedClientID(ctx, clientID)
		require.False(t, found)
	})
}

func TestVerifiedClientIDGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNVerifiedClientID(ctx, tk.MonitoringConsumerKeeper, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(tk.MonitoringConsumerKeeper.GetAllVerifiedClientID(ctx)),
	)
}

func TestAddVerifiedClientID(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("update a verified client id", func(t *testing.T) {
		var (
			launchID         = uint64(1)
			newClientID      = "2"
			verifiedClientID = types.VerifiedClientID{
				LaunchID:  launchID,
				ClientIDs: []string{"1"},
			}
		)
		tk.MonitoringConsumerKeeper.SetVerifiedClientID(ctx, verifiedClientID)
		tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, launchID, newClientID)
		got, found := tk.MonitoringConsumerKeeper.GetVerifiedClientID(ctx, launchID)
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
		tk.MonitoringConsumerKeeper.SetVerifiedClientID(ctx, verifiedClientID)
		tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, launchID, newClientID)
		got, found := tk.MonitoringConsumerKeeper.GetVerifiedClientID(ctx, launchID)
		require.True(t, found)
		require.Equal(t, verifiedClientID, got)
	})

	t.Run("update a non exiting verified client id", func(t *testing.T) {
		verifiedClientID := types.VerifiedClientID{
			LaunchID:  3,
			ClientIDs: []string{"1"},
		}
		tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, verifiedClientID.LaunchID, verifiedClientID.ClientIDs[0])
		got, found := tk.MonitoringConsumerKeeper.GetVerifiedClientID(ctx, verifiedClientID.LaunchID)
		require.True(t, found)
		require.Equal(t, verifiedClientID, got)
	})
}
