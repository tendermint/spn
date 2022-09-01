package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	connectiontypes "github.com/cosmos/ibc-go/v5/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	ignterrors "github.com/ignite/modules/errors"

	"github.com/tendermint/spn/x/monitoringc/types"
)

// VerifyClientIDFromConnID verifies if the client ID associated with the provided connection ID
// is a verified client ID and if no connection is yet established with the provider chain
// this operation should be performed at OnChanOpenInit handshake phase
func (k Keeper) VerifyClientIDFromConnID(ctx sdk.Context, connID string) error {
	clientID, err := k.getClientIDFromConnID(ctx, connID)
	if err != nil {
		return err
	}

	// check if the client ID is verified
	lidFromCid, found := k.GetLaunchIDFromVerifiedClientID(ctx, clientID)
	if !found {
		return sdkerrors.Wrapf(types.ErrClientNotVerified, clientID)
	}

	// check if the connection with the provider for this launch ID is already established
	pCid, found := k.GetProviderClientID(ctx, lidFromCid.LaunchID)
	if found {
		return sdkerrors.Wrapf(
			types.ErrConnectionAlreadyEstablished,
			"provider client ID for launch ID %d is: %s",
			pCid.LaunchID, pCid.ClientID,
		)
	}

	return nil
}

// RegisterProviderClientIDFromChannelID registers the verified client ID for the provider
// this operation should be performed at OnChanOpenAck
func (k Keeper) RegisterProviderClientIDFromChannelID(ctx sdk.Context, channelID string) error {
	connID, err := k.getConnIDFromChannelID(ctx, channelID)
	if err != nil {
		return err
	}

	clientID, err := k.getClientIDFromConnID(ctx, connID)
	if err != nil {
		return err
	}

	// get the launch ID from the client ID
	lidFromCid, found := k.GetLaunchIDFromVerifiedClientID(ctx, clientID)
	if !found {
		// client should be verified at this phase, so a critical error is returned
		return ignterrors.Criticalf("client ID %s should be verified during registration", clientID)
	}

	// another connection could have been established between OnChanOpenInit and OnChanOpenAck
	// so we check if provider client ID exists
	pCid, found := k.GetProviderClientID(ctx, lidFromCid.LaunchID)
	if found {
		return sdkerrors.Wrapf(
			types.ErrConnectionAlreadyEstablished,
			"provider connection for launch ID %d has been established: %s",
			pCid.LaunchID, pCid.ClientID,
		)
	}
	launchID := lidFromCid.LaunchID

	// update the chain since it is not MonitoringConnected
	if err = k.launchKeeper.EnableMonitoringConnection(ctx, launchID); err != nil {
		return err
	}

	// register the client for the provider
	k.SetProviderClientID(ctx, types.ProviderClientID{
		ClientID: clientID,
		LaunchID: launchID,
	})

	// associate the channel ID for the provider connection with the correct launch ID
	k.SetLaunchIDFromChannelID(ctx, types.LaunchIDFromChannelID{
		LaunchID:  launchID,
		ChannelID: channelID,
	})

	return nil
}

// getClientIDFromConnID retrieves the client ID associated with a connection ID
func (k Keeper) getClientIDFromConnID(ctx sdk.Context, connID string) (string, error) {
	conn, ok := k.connectionKeeper.GetConnection(ctx, connID)
	if !ok {
		return "", sdkerrors.Wrapf(
			connectiontypes.ErrConnectionNotFound,
			"connection not found for connection ID: %s",
			connID,
		)
	}
	return conn.GetClientID(), nil
}

// getConnIDFromChannelID retrieves the connection ID associated with the provided channel ID
func (k Keeper) getConnIDFromChannelID(ctx sdk.Context, channelID string) (string, error) {
	// retrieve the client ID from the channel
	channel, ok := k.channelKeeper.GetChannel(ctx, types.PortID, channelID)
	if !ok {
		return "", sdkerrors.Wrapf(
			channeltypes.ErrChannelNotFound,
			"channel not found for channel ID: %s",
			channelID,
		)
	}
	if len(channel.ConnectionHops) != 1 {
		return "", sdkerrors.Wrap(
			channeltypes.ErrTooManyConnectionHops,
			"must have direct connection to consumer chain",
		)
	}
	return channel.ConnectionHops[0], nil
}
