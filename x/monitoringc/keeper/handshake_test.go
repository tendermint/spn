package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	connectiontypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	monitoringcmodulekeeper "github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// monitoringcKeeperWithFooClient returns a test monitoring keeper containing necessary IBC mocks for a client with ID foo
func monitoringcKeeperWithFooClient(t *testing.T) (*monitoringcmodulekeeper.Keeper, sdk.Context) {
	return testkeeper.MonitoringcWithIBCMocks(
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
}

func TestKeeper_VerifyClientIDFromChannelID(t *testing.T) {
	t.Run("should returns no error if the client is verified and provider has no connection yet", func(t *testing.T) {
		k, ctx := monitoringcKeeperWithFooClient(t)
		k.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		err := k.VerifyClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err)
	})

	t.Run("should fails if channel doesn't exist", func(t *testing.T) {
		k, ctx := monitoringcKeeperWithFooClient(t)
		err := k.VerifyClientIDFromChannelID(ctx, "bar")
		require.ErrorIs(t, err, channeltypes.ErrChannelNotFound)
	})

	t.Run("should fails if the channel has more than 1 hop connection", func(t *testing.T) {
		k, ctx := testkeeper.MonitoringcWithIBCMocks(
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
		err := k.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, channeltypes.ErrTooManyConnectionHops)
	})

	t.Run("should fails if the connection doesn't exist", func(t *testing.T) {
		k, ctx := testkeeper.MonitoringcWithIBCMocks(
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
		err := k.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, connectiontypes.ErrConnectionNotFound)
	})

	t.Run("should fails if the client is not verified", func(t *testing.T) {
		k, ctx := monitoringcKeeperWithFooClient(t)
		err := k.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrClientNotVerified)
	})

	t.Run("should fails if the provider already has an established connection", func(t *testing.T) {
		k, ctx := monitoringcKeeperWithFooClient(t)
		k.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		k.SetProviderClientID(ctx, types.ProviderClientID{
			LaunchID: 1,
			ClientID: "bar",
		})
		err := k.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConnectionAlreadyEstablished)
	})
}

func TestKeeper_RegisterProviderClientIDFromChannelID(t *testing.T) {
	t.Run("should register the client id", func(t *testing.T) {
		k, ctx := monitoringcKeeperWithFooClient(t)
		k.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		err := k.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err)

		// the provider client ID should be created
		pCid, found := k.GetProviderClientID(ctx, 1)
		require.True(t, found)
		require.EqualValues(t, 1, pCid.LaunchID)
		require.EqualValues(t, "foo", pCid.ClientID)
	})

	t.Run("should fails with critical if channel doesn't exist", func(t *testing.T) {
		k, ctx := monitoringcKeeperWithFooClient(t)
		err := k.RegisterProviderClientIDFromChannelID(ctx, "bar")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("should fails with critical if the channel has more than 1 hop connection", func(t *testing.T) {
		k, ctx := testkeeper.MonitoringcWithIBCMocks(
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
		err := k.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("should fails with critical if the connection doesn't exist", func(t *testing.T) {
		k, ctx := testkeeper.MonitoringcWithIBCMocks(
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
		err := k.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("should fails with critical if the client is not verified", func(t *testing.T) {
		k, ctx := monitoringcKeeperWithFooClient(t)
		err := k.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("should fails if the provider already has an established connection", func(t *testing.T) {
		k, ctx := monitoringcKeeperWithFooClient(t)
		k.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		k.SetProviderClientID(ctx, types.ProviderClientID{
			LaunchID: 1,
			ClientID: "bar",
		})
		err := k.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConnectionAlreadyEstablished)
	})
}
