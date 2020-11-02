package chat

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/chat/keeper"
	"github.com/tendermint/spn/x/chat/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateChannel:
			return handleMsgCreateChannel(ctx, k, msg)
		case *types.MsgSendMessage:
			return handleMsgSendMessage(ctx, k, msg)
		case *types.MsgVotePoll:
			return handleMsgVotePoll(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateChannel(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateChannel) (*sdk.Result, error) {
	// Get the identity of the user
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Create the channel, the channel id will be modified by CreateChannel
	channel, err := types.NewChannel(0, identity, msg.Title, msg.Description, msg.Payload)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Append the channel
	err = k.CreateChannel(ctx, channel)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgSendMessage(ctx sdk.Context, k keeper.Keeper, msg *types.MsgSendMessage) (*sdk.Result, error) {
	// Get the identity of the user
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Create the message with the options provided and the block time
	message, err := types.NewMessage(
		msg.ChannelID,
		0, // Modified by the keeper
		identity,
		msg.Content,
		msg.Tags,
		ctx.BlockTime(),
		msg.PollOptions,
		msg.Payload,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Append the message
	found, err := k.AppendMessageToChannel(ctx, message)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Channel not found")
	}
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgVotePoll(ctx sdk.Context, k keeper.Keeper, msg *types.MsgVotePoll) (*sdk.Result, error) {
	// Get the identity of the user
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Create a new vote
	vote, err := types.NewVote(identity, msg.Value, msg.Payload)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Retrieve the message
	messageID := types.GetMessageIDFromChannelIDandIndex(msg.ChannelID, msg.MessageIndex)

	// Vote to the poll
	found, err := k.AppendVoteToPoll(ctx, messageID, &vote)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Message not found")
	}
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
