package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/monitoringc/types"
)

func (k Keeper) MonitoringHistory(c context.Context, req *types.QueryGetMonitoringHistoryRequest) (*types.QueryGetMonitoringHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMonitoringHistory(
		ctx,
		req.LaunchID,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMonitoringHistoryResponse{MonitoringHistory: val}, nil
}
