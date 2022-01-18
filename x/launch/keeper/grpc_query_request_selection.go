package keeper

import (
	"context"

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

	result := make([]types.Request, 0)
	for _, id := range ids {
		request, found := k.GetRequest(ctx, req.LaunchID, id)
		if !found {
			return nil, status.Errorf(codes.NotFound, "request with id %d not found", id)
		}
		result = append(result, request)
	}

	return &types.QueryRequestSelectionResponse{Request: result}, nil
}
