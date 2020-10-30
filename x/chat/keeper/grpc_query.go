package keeper

import (
	"fmt"
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

func (k Keeper) ListMessages(c context.Context, req *types.QueryListMessagesRequest) (*types.QueryListMessagesResponse, error) {\
	ctx := sdk.UnwrapSDKContext(c)
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	messages, found := k.GetAllMessagesFromChannel(ctx, req.ChannelId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("channel %v doesn't exist", req.ChannelId))
	}

	// Get pointers on messages
	var messagePtrs []*types.Message
	for _, mes := range messages {
		messagePtrs = append(messagePtrs, &mes)
	}

	return &types.QueryListMessagesResponse{Messages: messagePtrs}, nil
}

func (k Keeper) SearchMessages(c context.Context, req *types.QuerySearchMessagesRequest) (*types.QuerySearchMessagesResponse, error) {

	return nil, nil
}


