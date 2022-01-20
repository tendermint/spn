package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	connectiontypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// VerifyClientIDFromChannelID verifies if the client ID associated with the provided channel ID
// is a verified client ID and if no connection is yet established with the provider chain
func (k Keeper) VerifyClientIDFromChannelID(ctx sdk.Context, channelID string) error {
	// retrieve the client ID from the channel
	channel, ok := k.ChannelKeeper.GetChannel(ctx, types.PortID, channelID)
	if !ok {
		return sdkerrors.Wrapf(
			channeltypes.ErrChannelNotFound,
			"channel not found for channel ID: %s",
			channelID,
		)
	}
	if len(channel.ConnectionHops) != 1 {
		return sdkerrors.Wrap(
			channeltypes.ErrTooManyConnectionHops,
			"must have direct connection to baby chain",
		)
	}
	connectionID := channel.ConnectionHops[0]
	conn, ok := k.connectionKeeper.GetConnection(ctx, connectionID)
	if !ok {
		return sdkerrors.Wrapf(
			connectiontypes.ErrConnectionNotFound,
			"connection not found for connection ID: %s",
			connectionID,
		)
	}
	clientID := conn.GetClientID()

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
