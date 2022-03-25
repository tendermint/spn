package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/campaign/types"
)

func (k Keeper) MaximumShares(goCtx context.Context, req *types.QueryMaximumSharesRequest) (*types.QueryMaximumSharesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	shares := k.GetMaximumShares(ctx)

	return &types.QueryMaximumSharesResponse{MaximumShares: shares}, nil
}
