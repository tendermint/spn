package keeper_test

import (
	"testing"

	"github.com/tendermint/spn/testutil/sample"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestMissingVerifiedClientIDInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		n := sample.Uint64(r)
		launchID := sample.Uint64(r)
		for i := uint64(0); i < n; i++ {
			chanID := sample.AlphaString(r, 10)
			tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, launchID, chanID)
			tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
				ClientID: chanID,
				LaunchID: launchID,
			})
		}
		mes, broken := keeper.MissingVerifiedClientIDInvariant(*tk.MonitoringConsumerKeeper)(ctx)
		require.False(t, broken, mes)
	})

	t.Run("invalid case", func(t *testing.T) {
		n := sample.Uint64(r)
		launchID := sample.Uint64(r)
		for i := uint64(0); i < n; i++ {
			chanID := sample.AlphaString(r, 10)
			tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, launchID, chanID)
		}
		mes, broken := keeper.MissingVerifiedClientIDInvariant(*tk.MonitoringConsumerKeeper)(ctx)
		require.True(t, broken, mes)
	})
}
