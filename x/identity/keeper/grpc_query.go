package keeper

import (
	"context"

	"github.com/tendermint/spn/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Username(c context.Context, req *types.QueryGetUsernameRequest) (*types.QueryGetUsernameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "address cannot be empty")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)
	username, err := k.GetUsername(ctx, addr)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetUsernameResponse{Username: username}, nil
}
