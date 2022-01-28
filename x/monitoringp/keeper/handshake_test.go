package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	connectiontypes "github.com/cosmos/ibc-go/v2/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	monitoringpmodulekeeper "github.com/tendermint/spn/x/monitoringp/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
)

// monitoringpKeeperWithFooClient returns a test monitoring keeper containing necessary IBC mocks for a client with ID foo
func monitoringpKeeperWithFooClient(t *testing.T) (*monitoringpmodulekeeper.Keeper, sdk.Context) {
	k, _, ctx := testkeeper.MonitoringpKeeperWithIBCMock(
		t,
		[]testkeeper.Connection{
			{
				"foo",
				connectiontypes.ConnectionEnd{
					ClientId: "foo",
				},
			},
		},
		[]testkeeper.Channel{
			{
				"foo",
				channeltypes.Channel{
					ConnectionHops: []string{"foo"},
				},
			},
		},
	)
	return k, ctx
}

func TestKeeper_VerifyClientIDFromChannelID(t *testing.T) {
	t.Run("should returns no error if the client exists and no connection is established", func(t *testing.T) {
		k, ctx := monitoringpKeeperWithFooClient(t)
		k.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := k.VerifyClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err)
	})

	t.Run("should fails if channel doesn't exist", func(t *testing.T) {
		k, ctx := monitoringpKeeperWithFooClient(t)
		k.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := k.VerifyClientIDFromChannelID(ctx, "bar")
		require.ErrorIs(t, err, channeltypes.ErrChannelNotFound)
	})

	t.Run("should fails if the channel has more than 1 hop connection", func(t *testing.T) {
		k, _, ctx := testkeeper.MonitoringpKeeperWithIBCMock(
			t,
			[]testkeeper.Connection{},
			[]testkeeper.Channel{
				{
					"foo",
					channeltypes.Channel{
						ConnectionHops: []string{"foo", "bar"},
					},
				},
			},
		)
		k.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := k.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, channeltypes.ErrTooManyConnectionHops)
	})

	t.Run("should fails if the connection doesn't exist", func(t *testing.T) {
		k, _, ctx := testkeeper.MonitoringpKeeperWithIBCMock(
			t,
			[]testkeeper.Connection{},
			[]testkeeper.Channel{
				{
					"foo",
					channeltypes.Channel{
						ConnectionHops: []string{"foo"},
					},
				},
			},
		)
		k.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := k.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, connectiontypes.ErrConnectionNotFound)
	})

	t.Run("should fails if the client doesn't exist", func(t *testing.T) {
		k, ctx := monitoringpKeeperWithFooClient(t)
		err := k.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrNoConsumerClient)
	})

	t.Run("should fails if connection has already been established", func(t *testing.T) {
		k, ctx := monitoringpKeeperWithFooClient(t)
		k.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		k.SetConnectionChannelID(ctx, types.ConnectionChannelID{
			ChannelID: "bar",
		})
		err := k.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConsumerConnectionEstablished)
	})
}

func TestKeeper_RegisterConnectionChannelID(t *testing.T) {
	t.Run("should register the channel id", func(t *testing.T) {
		k, ctx := monitoringpKeeperWithFooClient(t)
		k.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		err := k.RegisterConnectionChannelID(ctx, "foo")
		require.NoError(t, err)
		channelID, found := k.GetConnectionChannelID(ctx)
		require.True(t, found)
		require.EqualValues(t, types.ConnectionChannelID{
			ChannelID: "foo",
		}, channelID)
	})

	t.Run("should fails with no critical if connection has already been established", func(t *testing.T) {
		k, ctx := monitoringpKeeperWithFooClient(t)
		k.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: "foo",
		})
		k.SetConnectionChannelID(ctx, types.ConnectionChannelID{
			ChannelID: "bar",
		})
		err := k.RegisterConnectionChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConsumerConnectionEstablished)
		require.NotErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("should fails with critical if verify channel id fails with other error than connection established", func(t *testing.T) {
		k, ctx := monitoringpKeeperWithFooClient(t)
		err := k.RegisterConnectionChannelID(ctx, "foo")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})
}
