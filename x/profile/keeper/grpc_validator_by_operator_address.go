package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/profile/types"
)

func (k Keeper) ValidatorByOperatorAddress(c context.Context, req *types.QueryGetValidatorByOperatorAddressRequest) (
	*types.QueryGetValidatorByOperatorAddressResponse,
	error,
) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetValidatorByOperatorAddress(ctx, req.OperatorAddress)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetValidatorByOperatorAddressResponse{ValidatorByOperatorAddress: val}, nil
}
