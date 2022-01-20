package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	connectiontypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)


func (k Keeper) getClientIDFromChannel(ctx sdk.Context, channelID string) (string, error) {
	// retrieve the client ID from the channel
	channel, ok := k.ChannelKeeper.GetChannel(ctx, types.PortID, channelID)
	if !ok {
		return "", sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "channel not found for channel ID: %s", channelID)
	}
	if len(channel.ConnectionHops) != 1 {
		return "", sdkerrors.Wrap(channeltypes.ErrTooManyConnectionHops, "must have direct connection to baby chain")
	}
	connectionID := channel.ConnectionHops[0]
	conn, ok := k.connectionKeeper.GetConnection(ctx, connectionID)
	if !ok {
		return "", sdkerrors.Wrapf(connectiontypes.ErrConnectionNotFound, "connection not found for connection ID: %s", connectionID)
	}

	//// check if the client ID is verified
	//k.GetVerifiedClientID()
	//
	//// check if the connection with the provider for this launch ID is already established
	//k.GetProviderClientID()

	return conn.ClientId, nil
}
