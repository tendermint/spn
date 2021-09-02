package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CoordinatorByAddress(c context.Context, req *types.QueryGetCoordinatorByAddressRequest) (*types.QueryGetCoordinatorByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCoordinatorByAddress(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetCoordinatorByAddressResponse{CoordinatorByAddress: val}, nil
}
