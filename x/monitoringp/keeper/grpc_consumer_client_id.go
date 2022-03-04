package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/monitoringp/types"
)

func (k Keeper) ConsumerClientID(c context.Context, req *types.QueryGetConsumerClientIDRequest) (*types.QueryGetConsumerClientIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetConsumerClientID(ctx)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetConsumerClientIDResponse{ConsumerClientID: val}, nil
}
