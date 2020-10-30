package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/chat/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListMessages(c context.Context, req *types.QueryListMessagesRequest) (*types.QueryListMessagesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	messages, found := k.GetAllMessagesFromChannel(ctx, req.ChannelId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("channel %v doesn't exist", req.ChannelId))
	}

	return &types.QueryListMessagesResponse{Messages: messages}, nil
}

func (k Keeper) SearchMessages(c context.Context, req *types.QuerySearchMessagesRequest) (*types.QuerySearchMessagesResponse, error) {

	return &types.QuerySearchMessagesResponse{Messages: messages}, nil
}
