package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/spn/x/monitoringc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MonitoringHistoryAll(c context.Context, req *types.QueryAllMonitoringHistoryRequest) (*types.QueryAllMonitoringHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var monitoringHistorys []types.MonitoringHistory
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	monitoringHistoryStore := prefix.NewStore(store, types.KeyPrefix(types.MonitoringHistoryKeyPrefix))

	pageRes, err := query.Paginate(monitoringHistoryStore, req.Pagination, func(key []byte, value []byte) error {
		var monitoringHistory types.MonitoringHistory
		if err := k.cdc.Unmarshal(value, &monitoringHistory); err != nil {
			return err
		}

		monitoringHistorys = append(monitoringHistorys, monitoringHistory)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMonitoringHistoryResponse{MonitoringHistory: monitoringHistorys, Pagination: pageRes}, nil
}

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
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetMonitoringHistoryResponse{MonitoringHistory: val}, nil
}
