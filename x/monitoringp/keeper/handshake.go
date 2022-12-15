package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	connectiontypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"

	"github.com/tendermint/spn/x/monitoringp/types"
)

// VerifyClientIDFromConnID verifies if the client ID associated with the provided connection ID
// is the consumer client ID and if no connection is yet established with the consumer chain
// this operation should be performed at OnChanOpenTry handshake phase
func (k Keeper) VerifyClientIDFromConnID(ctx sdk.Context, connID string) error {
	// get the current client ID and check if it's the consumer client ID
	clientID, err := k.getClientIDFromConnID(ctx, connID)
	if err != nil {
		return err
	}

	// check no connection is already established
	_, found := k.GetConnectionChannelID(ctx)
	if found {
		return types.ErrConsumerConnectionEstablished
	}

	// check if the consumer client ID exists
	consumerClient, found := k.GetConsumerClientID(ctx)
	if !found {
		return types.ErrNoConsumerClient
	}

	if consumerClient.ClientID != clientID {
		return sdkerrors.Wrapf(
			types.ErrInvalidClient,
			"the client is not the consumer client, got %s, expected %s",
			clientID,
			consumerClient.ClientID,
		)
	}

	return nil
}

// RegisterConnectionChannelID registers the channel ID for the connection used with the consumer chain
func (k Keeper) RegisterConnectionChannelID(ctx sdk.Context, channelID string) error {
	connID, err := k.getConnIDFromChannelID(ctx, channelID)
	if err != nil {
		return err
	}

	// verify the channel during registration
	// the channel could have been registered with another connection, if it's the case, an error is returned
	// any other error is a critical error
	if err = k.VerifyClientIDFromConnID(ctx, connID); err != nil {
		return err
	}

	// register the connection channel ID
	k.SetConnectionChannelID(ctx, types.ConnectionChannelID{
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
