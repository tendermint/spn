package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	connectiontypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
)

// monitoringpKeeperWithFooClient returns a test monitoring keeper containing necessary IBC mocks for a client with ID foo
func monitoringpKeeperWithFooClient(t *testing.T) (sdk.Context, testkeeper.TestKeepers, testkeeper.TestMsgServers) {
	return testkeeper.NewTestSetupWithIBCMocksMonitoringp(
		t,
		[]testkeeper.Connection{
			{
				ConnID: "foo",
				Conn: connectiontypes.ConnectionEnd{
					ClientId: "foo",
				},
			},
		},
		[]testkeeper.Channel{
			{
				ChannelID: "foo",
				Channel: channeltypes.Channel{
					ConnectionHops: []string{"foo"},
				},
			},
		},
	)
}

func TestKeeper_VerifyClientIDFromChannelID(t *testing.T) {
	t.Run("should return no error if the client exists and no connection is established", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromConnID(ctx, "foo")
		require.NoError(t, err)
	})

	t.Run("should fail if connection client ID is not equal to consumerClientID", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "notequal",
		})
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromConnID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrInvalidClient)
	})

	t.Run("should fail if connection doesn't exist", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromConnID(ctx, "bar")
		require.ErrorIs(t, err, connectiontypes.ErrConnectionNotFound)
	})

	t.Run("should fail if the consumer client doesn't exist", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromConnID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrNoConsumerClient)
	})

	t.Run("should fail if connection has already been established", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		tk.MonitoringProviderKeeper.SetConnectionChannelID(ctx, types.ConnectionChannelID{
			ChannelID: "bar",
		})
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromConnID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConsumerConnectionEstablished)
	})
}

func TestKeeper_RegisterConnectionChannelID(t *testing.T) {
	t.Run("should register the channel id", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "foo")
		require.NoError(t, err)
		channelID, found := tk.MonitoringProviderKeeper.GetConnectionChannelID(ctx)
		require.True(t, found)
		require.EqualValues(t, types.ConnectionChannelID{
			ChannelID: "foo",
		}, channelID)
	})

	t.Run("should fail if connection has already been established", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		tk.MonitoringProviderKeeper.SetConnectionChannelID(ctx, types.ConnectionChannelID{
			ChannelID: "bar",
		})
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConsumerConnectionEstablished)
	})

	t.Run("should fail if the consumer client doesn't exist", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrNoConsumerClient)
	})

	t.Run("should fail if the channel doesn't exist", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "bar")
		require.ErrorIs(t, err, channeltypes.ErrChannelNotFound)
	})

	t.Run("should fail if the channel has more than 1 hop connection", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithIBCMocksMonitoringp(
			t,
			[]testkeeper.Connection{},
			[]testkeeper.Channel{
				{
					ChannelID: "foo",
					Channel: channeltypes.Channel{
						ConnectionHops: []string{"foo", "bar"},
					},
				},
			},
		)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "foo")
		require.ErrorIs(t, err, channeltypes.ErrTooManyConnectionHops)
	})

	t.Run("should fail if the connection doesn't exist", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithIBCMocksMonitoringp(
			t,
			[]testkeeper.Connection{},
			[]testkeeper.Channel{
				{
					ChannelID: "foo",
					Channel: channeltypes.Channel{
						ConnectionHops: []string{"foo"},
					},
				},
			},
		)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "foo")
		require.ErrorIs(t, err, connectiontypes.ErrConnectionNotFound)
	})
}
