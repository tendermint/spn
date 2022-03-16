package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/participation/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AvailableAllocations(goCtx context.Context,
	req *types.QueryGetAvailableAllocationsRequest,
) (*types.QueryGetAvailableAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	availableAlloc, err := k.GetAvailableAllocations(ctx, req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryGetAvailableAllocationsResponse{AvailableAllocations: availableAlloc}, nil
}
