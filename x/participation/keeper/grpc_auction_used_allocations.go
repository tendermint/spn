package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/participation/types"
)

func (k Keeper) AuctionUsedAllocationsAll(c context.Context, req *types.QueryAllAuctionUsedAllocationsRequest) (*types.QueryAllAuctionUsedAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var auctionUsedAllocationss []types.AuctionUsedAllocations
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	auctionUsedAllocationsStore := prefix.NewStore(store, types.KeyPrefix(types.AuctionUsedAllocationsKeyPrefix))
	addressAuctionUsedAllocationsStore := prefix.NewStore(auctionUsedAllocationsStore, types.KeyPrefix(req.Address))

	pageRes, err := query.Paginate(addressAuctionUsedAllocationsStore, req.Pagination, func(key []byte, value []byte) error {
		var auctionUsedAllocations types.AuctionUsedAllocations
		if err := k.cdc.Unmarshal(value, &auctionUsedAllocations); err != nil {
			return err
		}

		auctionUsedAllocationss = append(auctionUsedAllocationss, auctionUsedAllocations)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAuctionUsedAllocationsResponse{AuctionUsedAllocations: auctionUsedAllocationss, Pagination: pageRes}, nil
}

func (k Keeper) AuctionUsedAllocations(c context.Context, req *types.QueryGetAuctionUsedAllocationsRequest) (*types.QueryGetAuctionUsedAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAuctionUsedAllocations(
		ctx,
		req.Address,
		req.AuctionID,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAuctionUsedAllocationsResponse{AuctionUsedAllocations: val}, nil
}
