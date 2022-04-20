package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/participation/types"
)

func (k Keeper) TotalAllocations(goCtx context.Context, req *types.QueryGetTotalAllocationsRequest) (*types.QueryGetTotalAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	numAlloc, err := k.GetTotalAllocations(ctx, req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryGetTotalAllocationsResponse{TotalAllocations: numAlloc}, nil
}
