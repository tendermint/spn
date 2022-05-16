package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/claim/types"
)

func (k Keeper) AirdropSupply(c context.Context, req *types.QueryGetAirdropSupplyRequest) (*types.QueryGetAirdropSupplyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAirdropSupply(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAirdropSupplyResponse{AirdropSupply: val}, nil
}
