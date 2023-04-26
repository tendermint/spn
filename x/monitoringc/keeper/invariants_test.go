package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestMissingVerifiedClientIDInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should allow valid case", func(t *testing.T) {
		n := sample.Uint64(r)
		launchID := sample.Uint64(r)
		for i := uint64(0); i < n; i++ {
			clientID := sample.AlphaString(r, 10)
			tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, launchID, clientID)
			tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
				ClientID: clientID,
				LaunchID: launchID,
			})
		}
		msg, broken := keeper.MissingVerifiedClientIDInvariant(*tk.MonitoringConsumerKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should prevent invalid case", func(t *testing.T) {
		n := sample.Uint64(r)
		launchID := sample.Uint64(r)
		for i := uint64(0); i < n; i++ {
			clientID := sample.AlphaString(r, 10)
			tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, launchID, clientID)
		}
		msg, broken := keeper.MissingVerifiedClientIDInvariant(*tk.MonitoringConsumerKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
