package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/starport/starport/pkg/numbers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RequestSelection(goCtx context.Context, req *types.QueryRequestSelectionRequest) (*types.QueryRequestSelectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ids, err := numbers.ParseList(req.RequestIDs)
	if err != nil || len(ids) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	store := ctx.KVStore(k.storeKey)
	keyPrefix := append(types.KeyPrefix(types.RequestKeyPrefix), types.RequestPoolKey(req.LaunchID)...)
	requestStore := prefix.NewStore(store, keyPrefix)

	result := make([]types.Request, 0)
	iterator := requestStore.Iterator(
		types.RequestIDBytes(ids[0]),
		types.RequestIDBytes(ids[len(ids)-1]+1),
	)
	defer iterator.Close()
	selectionMap := prepareSelectionMap(ids)

	for ; iterator.Valid(); iterator.Next() {
		var request types.Request
		err := k.cdc.Unmarshal(iterator.Value(), &request)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if _, ok := selectionMap[request.RequestID]; ok {
			result = append(result, request)
			selectionMap[request.RequestID] = struct{}{}
		}
	}

	return &types.QueryRequestSelectionResponse{Request: result}, nil
}

func prepareSelectionMap(ids []uint64) map[uint64]struct{} {
	selectionMap := make(map[uint64]struct{})
	for _, id := range ids {
		selectionMap[id] = struct{}{}
	}
	return selectionMap
}
