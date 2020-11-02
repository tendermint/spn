package keeper

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/chat/types"
)

var _ types.QueryServer = Keeper{}

// ShowChannel returns the channel data from its ID
func (k Keeper) ShowChannel(c context.Context, req *types.QueryShowChannelRequest) (*types.QueryShowChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	channel, found := k.GetChannel(ctx, req.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChannel, fmt.Sprintf("channel not found"))
	}

	return &types.QueryShowChannelResponse{Channel: &channel}, nil
}

// ListChannels list all the channels
func (k Keeper) ListChannels(c context.Context, req *types.QueryListChannelsRequest) (*types.QueryListChannelsResponse, error) {
	// TODO: Implement
	return nil, nil
}

// ListMessages lists all the messages in a channel
func (k Keeper) ListMessages(c context.Context, req *types.QueryListMessagesRequest) (*types.QueryListMessagesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	messages, found := k.GetAllMessagesFromChannel(ctx, req.ChannelId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("channel %v doesn't exist", req.ChannelId))
	}

	// Get pointers on messages
	messagePtrs := make([]*types.Message, len(messages))

	for i := range messages {
		messagePtrs[i] = &messages[i]
	}

	return &types.QueryListMessagesResponse{Messages: messagePtrs}, nil
}

// SearchMessages lists all the message in a channel containing a specific tag
func (k Keeper) SearchMessages(c context.Context, req *types.QuerySearchMessagesRequest) (*types.QuerySearchMessagesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var messagePtrs []*types.Message

	// Get the tag references
	tagReferences := k.GetTagReferencesFromChannel(ctx, req.Tag, req.ChannelId)
	if len(tagReferences) == 0 {
		return &types.QuerySearchMessagesResponse{Messages: messagePtrs}, nil
	}

	// Get the messages from the tag references
	messages := k.GetMessagesByIDs(ctx, tagReferences)

	// Get pointers on messages
	messagePtrs = make([]*types.Message, len(messages))

	for i := range messages {
		messagePtrs[i] = &messages[i]
	}

	return &types.QuerySearchMessagesResponse{Messages: messagePtrs}, nil
}
