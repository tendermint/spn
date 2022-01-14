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
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if len(ids) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no request id provided")
	}

	store := ctx.KVStore(k.storeKey)
	keyPrefix := append(types.KeyPrefix(types.RequestKeyPrefix), types.RequestPoolKey(req.LaunchID)...)
	requestStore := prefix.NewStore(store, keyPrefix)

	result := make([]types.Request, 0)
	for _, id := range ids {
		var request types.Request

		b := requestStore.Get(types.RequestIDBytes(id))

		if b == nil {
			// TODO: decide on error content
			return nil, status.Error(codes.NotFound, "Not found")
		}

		err := k.cdc.Unmarshal(b, &request)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		result = append(result, request)
	}

	return &types.QueryRequestSelectionResponse{Request: result}, nil
}
