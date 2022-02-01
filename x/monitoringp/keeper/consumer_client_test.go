package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestKeeper_InitializeConsumerClient(t *testing.T) {

	t.Run("initialize consumer client", func(t *testing.T) {
		k, ibcKeeper, ctx := testkeeper.MonitoringpKeeper(t)

		// set params with valid values
		k.SetParams(ctx, types.NewParams(
			1000,
			types.DefautConsumerChainID,
			sample.ConsensusState(0),
			false,
		))
		clientID, err := k.InitializeConsumerClient(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, clientID)

		consumerClientID, found := k.GetConsumerClientID(ctx)
		require.True(t, found, "consumer client ID should be registered in the store")
		require.EqualValues(t, clientID, consumerClientID.ClientID)

		// IBC client is created
		_, found = ibcKeeper.ClientKeeper.GetClientState(ctx, clientID)
		require.True(t, found, "IBC consumer client state should be created")
	})

	t.Run("invalid consumer consensus state", func(t *testing.T) {
		k, _, ctx := testkeeper.MonitoringpKeeper(t)

		// default params contain an empty consensus state, therefore invalid
		_, err := k.InitializeConsumerClient(ctx)
		require.ErrorIs(t, err, types.ErrInvalidConsensusState)
	})
}
