package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/monitoringp/types"
)

func (k Keeper) MonitoringInfo(c context.Context, req *types.QueryGetMonitoringInfoRequest) (*types.QueryGetMonitoringInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMonitoringInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMonitoringInfoResponse{MonitoringInfo: val}, nil
}
