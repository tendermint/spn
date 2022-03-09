package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	connectiontypes "github.com/cosmos/ibc-go/v2/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	spntypes "github.com/tendermint/spn/pkg/types"
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
	t.Run("should returns no error if the client exists and no connection is established", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err)
	})

	t.Run("should fail if channel doesn't exist", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromChannelID(ctx, "bar")
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
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromChannelID(ctx, "foo")
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
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, connectiontypes.ErrConnectionNotFound)
	})

	t.Run("should fail if the client doesn't exist", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrNoConsumerClient)
	})

	t.Run("should fails if connection has already been established", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		tk.MonitoringProviderKeeper.SetConnectionChannelID(ctx, types.ConnectionChannelID{
			ChannelID: "bar",
		})
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConsumerConnectionEstablished)
	})

	t.Run("debug mode should fail if client ID can't be retrieve from channel ID", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetParams(ctx, types.NewParams(
			1,
			"foo-1",
			spntypes.ConsensusState{},
			spntypes.DefaultUnbondingPeriod,
			1,
			true,
		))
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromChannelID(ctx, "bar")
		require.ErrorIs(t, err, channeltypes.ErrChannelNotFound)
	})

	t.Run("should return no error when debug mode is set and client doesn't exist", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetParams(ctx, types.NewParams(
			1,
			"foo-1",
			spntypes.ConsensusState{},
			spntypes.DefaultUnbondingPeriod,
			1,
			true,
		))
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err)
	})

	t.Run("should return no error when debug mode is set and connection has already been established", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetParams(ctx, types.NewParams(
			1,
			"foo-1",
			spntypes.ConsensusState{},
			spntypes.DefaultUnbondingPeriod,
			1,
			true,
		))
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		tk.MonitoringProviderKeeper.SetConnectionChannelID(ctx, types.ConnectionChannelID{
			ChannelID: "bar",
		})
		err := tk.MonitoringProviderKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err)
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

	t.Run("should fails with no critical if connection has already been established", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		tk.MonitoringProviderKeeper.SetConnectionChannelID(ctx, types.ConnectionChannelID{
			ChannelID: "bar",
		})
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConsumerConnectionEstablished)
		require.NotErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("should fails with critical if verify channel id fails with other error than connection established", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "foo")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("debug mode should fail with critical if client ID can't be retrieve from channel ID", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetParams(ctx, types.NewParams(
			1,
			"foo-1",
			spntypes.ConsensusState{},
			spntypes.DefaultUnbondingPeriod,
			1,
			true,
		))
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "bar")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("debug mode allow to register a channel ID when consumer client ID doesn't exist", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetParams(ctx, types.NewParams(
			1,
			"foo-1",
			spntypes.ConsensusState{},
			spntypes.DefaultUnbondingPeriod,
			1,
			true,
		))
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "foo")
		require.NoError(t, err)
		channelID, found := tk.MonitoringProviderKeeper.GetConnectionChannelID(ctx)
		require.True(t, found)
		require.EqualValues(t, types.ConnectionChannelID{
			ChannelID: "foo",
		}, channelID)
	})

	t.Run("debug mode allow to register a new channel ID and replace previous one", func(t *testing.T) {
		ctx, tk, _ := monitoringpKeeperWithFooClient(t)
		tk.MonitoringProviderKeeper.SetParams(ctx, types.NewParams(
			1,
			"foo-1",
			spntypes.ConsensusState{},
			spntypes.DefaultUnbondingPeriod,
			1,
			true,
		))
		tk.MonitoringProviderKeeper.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		tk.MonitoringProviderKeeper.SetConnectionChannelID(ctx, types.ConnectionChannelID{
			ChannelID: "bar",
		})
		err := tk.MonitoringProviderKeeper.RegisterConnectionChannelID(ctx, "foo")
		require.NoError(t, err)
		channelID, found := tk.MonitoringProviderKeeper.GetConnectionChannelID(ctx)
		require.True(t, found)
		require.EqualValues(t, types.ConnectionChannelID{
			ChannelID: "foo",
		}, channelID)
	})
}
