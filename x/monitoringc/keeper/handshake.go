package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	connectiontypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// DebugModeLaunchID is the launch ID automatically used when debug mode is set
const DebugModeLaunchID = 1

// VerifyClientIDFromChannelID verifies if the client ID associated with the provided channel ID
// is a verified client ID and if no connection is yet established with the provider chain
// this operation should be performed at OnChanOpenInit handshake phase
func (k Keeper) VerifyClientIDFromChannelID(ctx sdk.Context, channelID string) error {
	clientID, err := k.getClientIDFromChannelID(ctx, channelID)
	if err != nil {
		return err
	}

	// no verification if debug mode is set
	if !k.DebugMode(ctx) {
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
	}

	return nil
}

// RegisterProviderClientIDFromChannelID registers the verified client ID for the provider
// this operation should be performed at OnChanOpenAck
func (k Keeper) RegisterProviderClientIDFromChannelID(ctx sdk.Context, channelID string) error {
	clientID, err := k.getClientIDFromChannelID(ctx, channelID)
	if err != nil {
		return spnerrors.Criticalf(
			"client ID %s should be retrieved during registration, got error: %s",
			clientID,
			err.Error(),
		)
	}

	// if debug mode is set, the launch ID 1 is automatically registered for the client
	if k.DebugMode(ctx) {
		k.SetProviderClientID(ctx, types.ProviderClientID{
			ClientID: clientID,
			LaunchID: DebugModeLaunchID,
		})

		// TODO: add launch ID from channel ID

	} else {
		// get the launch ID from the client ID
		lidFromCid, found := k.GetLaunchIDFromVerifiedClientID(ctx, clientID)
		if !found {
			// client should be verified at this phase, so a critical error is returned
			return spnerrors.Criticalf("client ID %s should be verified during registration", clientID)
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

		// register the client for the provider
		k.SetProviderClientID(ctx, types.ProviderClientID{
			ClientID: clientID,
			LaunchID: lidFromCid.LaunchID,
		})
	}

	return nil
}

// getClientIDFromChannelID retrieves the client ID associated with the provided channel ID
func (k Keeper) getClientIDFromChannelID(ctx sdk.Context, channelID string) (string, error) {
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
			"must have direct connection to baby chain",
		)
	}
	connectionID := channel.ConnectionHops[0]
	conn, ok := k.connectionKeeper.GetConnection(ctx, connectionID)
	if !ok {
		return "", sdkerrors.Wrapf(
			connectiontypes.ErrConnectionNotFound,
			"connection not found for connection ID: %s",
			connectionID,
		)
	}
	return conn.GetClientID(), nil
}
