package keeper_test

import (
	"testing"

	launchtypes "github.com/tendermint/spn/x/launch/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	connectiontypes "github.com/cosmos/ibc-go/v2/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	monitoringcmodulekeeper "github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// testSetupWithFooClient returns a test setup with monitoring keeper containing necessary IBC mocks for a client with ID foo
func testSetupWithFooClient(t *testing.T) (sdk.Context, testkeeper.TestKeepers, testkeeper.TestMsgServers) {
	return testkeeper.NewTestSetupWithIBCMocks(
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
	t.Run("should returns no error if the client is verified and provider has no connection yet", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err)
	})

	t.Run("should fails if channel doesn't exist", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromChannelID(ctx, "bar")
		require.ErrorIs(t, err, channeltypes.ErrChannelNotFound)
	})

	t.Run("should fails if the channel has more than 1 hop connection", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithIBCMocks(
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
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, channeltypes.ErrTooManyConnectionHops)
	})

	t.Run("should fails if the connection doesn't exist", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithIBCMocks(
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
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, connectiontypes.ErrConnectionNotFound)
	})

	t.Run("should fails if the client is not verified", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrClientNotVerified)
	})

	t.Run("should fails if the provider already has an established connection", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		tk.MonitoringConsumerKeeper.SetProviderClientID(ctx, types.ProviderClientID{
			LaunchID: 1,
			ClientID: "bar",
		})
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConnectionAlreadyEstablished)
	})

	t.Run("debug mode should fail if client ID can't be retrieve from channel ID", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetParams(ctx, types.NewParams(true))
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromChannelID(ctx, "bar")
		require.ErrorIs(t, err, channeltypes.ErrChannelNotFound)
	})

	t.Run("should return no error when debug mode is set and client is not verified", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetParams(ctx, types.NewParams(true))
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err)
	})

	t.Run("should return no error when debug mode is set and connection is already established", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetParams(ctx, types.NewParams(true))
		tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		tk.MonitoringConsumerKeeper.SetProviderClientID(ctx, types.ProviderClientID{
			LaunchID: 1,
			ClientID: "bar",
		})

		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err)
	})
}

func TestKeeper_RegisterProviderClientIDFromChannelID(t *testing.T) {
	t.Run("should register the client id", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})

		chain := launchtypes.Chain{
			LaunchID: 1,
		}
		tk.LaunchKeeper.SetChain(ctx, chain)

		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err)

		// check that the chain is properly set to have MonitoringConnected be true
		chain, found := tk.LaunchKeeper.GetChain(ctx, 1)
		require.True(t, found)
		require.True(t, chain.MonitoringConnected)

		// the provider client ID should be created
		pCid, found := tk.MonitoringConsumerKeeper.GetProviderClientID(ctx, 1)
		require.True(t, found)
		require.EqualValues(t, 1, pCid.LaunchID)
		require.EqualValues(t, "foo", pCid.ClientID)

		// the channel ID is associated with the correct launch ID
		launcIDFromChanID, found := tk.MonitoringConsumerKeeper.GetLaunchIDFromChannelID(ctx, "foo")
		require.True(t, found)
		require.EqualValues(t, 1, launcIDFromChanID.LaunchID)
		require.EqualValues(t, "foo", launcIDFromChanID.ChannelID)
	})

	t.Run("should fails with critical if channel doesn't exist", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "bar")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("should fails with critical if the channel has more than 1 hop connection", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithIBCMocks(
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
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("should fails with critical if the connection doesn't exist", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithIBCMocks(
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
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("should fails with critical if the client is not verified", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("should fails if the provider already has an established connection", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		tk.MonitoringConsumerKeeper.SetProviderClientID(ctx, types.ProviderClientID{
			LaunchID: 1,
			ClientID: "bar",
		})
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConnectionAlreadyEstablished)
	})

	t.Run("debug mode should fail with critical if client ID can't be retrieve from channel ID", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetParams(ctx, types.NewParams(true))
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "bar")
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	})

	t.Run("debug mode allows to automatically register the client for a predefined launch ID", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetParams(ctx, types.NewParams(true))
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err, spnerrors.ErrCritical)

		pCid, found := tk.MonitoringConsumerKeeper.GetProviderClientID(ctx, monitoringcmodulekeeper.DebugModeLaunchID)
		require.True(t, found)
		require.EqualValues(t, monitoringcmodulekeeper.DebugModeLaunchID, pCid.LaunchID)
		require.EqualValues(t, "foo", pCid.ClientID)

		launcIDFromChanID, found := tk.MonitoringConsumerKeeper.GetLaunchIDFromChannelID(ctx, "foo")
		require.True(t, found)
		require.EqualValues(t, monitoringcmodulekeeper.DebugModeLaunchID, launcIDFromChanID.LaunchID)
		require.EqualValues(t, "foo", launcIDFromChanID.ChannelID)
	})

	t.Run("debug mode allows to register a new client and replace previous one", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetParams(ctx, types.NewParams(true))
		tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		tk.MonitoringConsumerKeeper.SetProviderClientID(ctx, types.ProviderClientID{
			LaunchID: 1,
			ClientID: "bar",
		})
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.NoError(t, err, spnerrors.ErrCritical)

		pCid, found := tk.MonitoringConsumerKeeper.GetProviderClientID(ctx, monitoringcmodulekeeper.DebugModeLaunchID)
		require.True(t, found)
		require.EqualValues(t, monitoringcmodulekeeper.DebugModeLaunchID, pCid.LaunchID)
		require.EqualValues(t, "foo", pCid.ClientID)

		launcIDFromChanID, found := tk.MonitoringConsumerKeeper.GetLaunchIDFromChannelID(ctx, "foo")
		require.True(t, found)
		require.EqualValues(t, monitoringcmodulekeeper.DebugModeLaunchID, launcIDFromChanID.LaunchID)
		require.EqualValues(t, "foo", launcIDFromChanID.ChannelID)
	})
}
