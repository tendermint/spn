package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/monitoringc/types"
)

func (k Keeper) VerifiedClientIds(goCtx context.Context, req *types.QueryGetVerifiedClientIdsRequest) (*types.QueryGetVerifiedClientIdsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	verifiedClientID, found := k.GetVerifiedClientID(ctx, req.LaunchID)
	if !found {
		return nil, status.Errorf(codes.Internal, "launch id not found %d", req.LaunchID)
	}

	return &types.QueryGetVerifiedClientIdsResponse{
		ClientIds: verifiedClientID.ClientIDs,
	}, nil
}
