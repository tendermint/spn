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

// DescribeChannel returns the channel data from its ID
func (k Keeper) DescribeChannel(c context.Context, req *types.QueryDescribeChannelRequest) (*types.QueryDescribeChannelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	channel, found := k.GetChannel(ctx, req.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChannel, fmt.Sprintf("channel not found"))
	}

	return &types.QueryDescribeChannelResponse{Channel: &channel}, nil
}
