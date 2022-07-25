package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/claim/types"
)

func (k Keeper) ClaimRecordAll(c context.Context, req *types.QueryAllClaimRecordRequest) (*types.QueryAllClaimRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var claimRecords []types.ClaimRecord
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	claimRecordStore := prefix.NewStore(store, types.KeyPrefix(types.ClaimRecordKeyPrefix))

	pageRes, err := query.Paginate(claimRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var claimRecord types.ClaimRecord
		if err := k.cdc.Unmarshal(value, &claimRecord); err != nil {
			return err
		}

		claimRecords = append(claimRecords, claimRecord)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClaimRecordResponse{ClaimRecord: claimRecords, Pagination: pageRes}, nil
}

func (k Keeper) ClaimRecord(c context.Context, req *types.QueryGetClaimRecordRequest) (*types.QueryGetClaimRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetClaimRecord(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetClaimRecordResponse{ClaimRecord: val}, nil
}
