package keeper_test

import (
	"testing"
	"time"

	ibctmtypes "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestKeeper_InitializeConsumerClient(t *testing.T) {
	t.Run("initialize consumer client", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)

		// set params with valid values
		tk.MonitoringProviderKeeper.SetParams(ctx, types.NewParams(
			1000,
			types.DefaultConsumerChainID,
			sample.ConsensusState(0),
			spntypes.DefaultUnbondingPeriod,
			spntypes.DefaultRevisionHeight,
		))
		clientID, err := tk.MonitoringProviderKeeper.InitializeConsumerClient(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, clientID)

		consumerClientID, found := tk.MonitoringProviderKeeper.GetConsumerClientID(ctx)
		require.True(t, found, "consumer client ID should be registered in the store")
		require.EqualValues(t, clientID, consumerClientID.ClientID)

		// IBC client should be created
		clientState, found := tk.IBCKeeper.ClientKeeper.GetClientState(ctx, clientID)
		require.True(t, found, "IBC consumer client state should be created")

		cs, ok := clientState.(*ibctmtypes.ClientState)
		require.True(t, ok)
		require.EqualValues(t, tk.MonitoringProviderKeeper.ConsumerRevisionHeight(ctx), cs.LatestHeight.RevisionHeight)
		require.EqualValues(t, time.Second*time.Duration(tk.MonitoringProviderKeeper.ConsumerUnbondingPeriod(ctx)), cs.UnbondingPeriod)
	})

	t.Run("invalid consumer consensus state", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)

		// default params contain an empty consensus state, therefore invalid
		_, err := tk.MonitoringProviderKeeper.InitializeConsumerClient(ctx)
		require.ErrorIs(t, err, types.ErrInvalidConsensusState)
	})
}
