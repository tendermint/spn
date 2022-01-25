package keeper

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	connectiontypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/monitoringp/types"
)

// VerifyClientIDFromChannelID verifies if the client ID associated with the provided channel ID
// is the consumer client ID and if no connection is yet established with the consumer chain
// this operation should be performed at OnChanOpenTry handshake phase
func (k Keeper) VerifyClientIDFromChannelID(ctx sdk.Context, channelID string) error {
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

	// get the current client ID and check if it's the consumer client ID
	clientID, err := k.getClientIDFromChannelID(ctx, channelID)
	if err != nil {
		return err
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
	// verify the channel during registration
	// the channel could have been registered with another connection, if it's the case, an error is returned
	// any other error is a critical error
	err := k.VerifyClientIDFromChannelID(ctx, channelID)
	if errors.Is(err, types.ErrConsumerConnectionEstablished) {
		return err
	} else if err != nil {
		return spnerrors.Criticalf(
			"failed to verify channelID during end of handshake: %s",
			err.Error(),
		)
	}

	// register the connection channel ID
	k.SetConnectionChannelID(ctx, types.ConnectionChannelID{
		ChannelID: channelID,
	})
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
