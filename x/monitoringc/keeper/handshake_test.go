package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	connectiontypes "github.com/cosmos/ibc-go/v5/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	ignterrors "github.com/ignite/modules/pkg/errors"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
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
	t.Run("should return no error if the client is verified and provider has no connection yet", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromConnID(ctx, "foo")
		require.NoError(t, err)
	})

	t.Run("should fail if connection doesn't exist", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromConnID(ctx, "bar")
		require.ErrorIs(t, err, connectiontypes.ErrConnectionNotFound)
	})

	t.Run("should fail if the client is not verified", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromConnID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrClientNotVerified)
	})

	t.Run("should fail if the provider already has an established connection", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		tk.MonitoringConsumerKeeper.SetProviderClientID(ctx, types.ProviderClientID{
			LaunchID: 1,
			ClientID: "bar",
		})
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromConnID(ctx, "foo")
		require.ErrorIs(t, err, types.ErrConnectionAlreadyEstablished)
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
		launchIDFromChanID, found := tk.MonitoringConsumerKeeper.GetLaunchIDFromChannelID(ctx, "foo")
		require.True(t, found)
		require.EqualValues(t, 1, launchIDFromChanID.LaunchID)
		require.EqualValues(t, "foo", launchIDFromChanID.ChannelID)
	})

	t.Run("should fail if the channel doesn't exist", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "bar")
		require.ErrorIs(t, err, channeltypes.ErrChannelNotFound)
	})

	t.Run("should fail with critical error if the client is not verified", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, ignterrors.ErrCritical)
	})

	t.Run("should fail if the provider already has an established connection", func(t *testing.T) {
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

	t.Run("should fail if monitoring connection already enabled on chain", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: "foo",
		})
		chain := launchtypes.Chain{
			LaunchID:            1,
			MonitoringConnected: true,
		}
		tk.LaunchKeeper.SetChain(ctx, chain)
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(ctx, "foo")
		require.ErrorIs(t, err, launchtypes.ErrChainMonitoringConnected)
	})

	t.Run("should fail if the channel has more than 1 hop connection", func(t *testing.T) {
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
		require.ErrorIs(t, err, channeltypes.ErrTooManyConnectionHops)
	})

	t.Run("should fail if the connection doesn't exist", func(t *testing.T) {
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
		require.ErrorIs(t, err, connectiontypes.ErrConnectionNotFound)
	})
}
