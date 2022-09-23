package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/launch/types"
)

func (k Keeper) ParamChangeAll(goCtx context.Context, req *types.QueryAllParamChangeRequest) (
	*types.QueryAllParamChangeResponse, error,
) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var paramChanges []types.ParamChange
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	paramChangeStore := prefix.NewStore(store, types.ParamChangeAllKey(req.LaunchID))

	pageRes, err := query.Paginate(paramChangeStore, req.Pagination, func(key []byte, value []byte) error {
		var paramChange types.ParamChange
		if err := k.cdc.Unmarshal(value, &paramChange); err != nil {
			return err
		}

		paramChanges = append(paramChanges, paramChange)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllParamChangeResponse{ParamChanges: paramChanges, Pagination: pageRes}, nil
}
