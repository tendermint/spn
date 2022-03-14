package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/participation/types"
)

func (k Keeper) TotalAllocation(goCtx context.Context, req *types.QueryGetTotalAllocationRequest) (*types.QueryGetTotalAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	acc, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	numAlloc, err := k.GetTotalAllocation(ctx, acc)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryGetTotalAllocationResponse{TotalAllocation: numAlloc}, nil
}
