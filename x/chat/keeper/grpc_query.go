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
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
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
