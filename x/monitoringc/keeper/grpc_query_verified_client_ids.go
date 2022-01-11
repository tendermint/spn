package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/monitoringc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerifiedClientIds(goCtx context.Context, req *types.QueryVerifiedClientIdsRequest) (*types.QueryVerifiedClientIdsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var verifiedClientIDs []string

	store := ctx.KVStore(k.storeKey)
	clientIDsStore := prefix.NewStore(store, types.VerifiedClientIDsPrefix(req.LaunchID))
	pageRes, err := query.Paginate(clientIDsStore, req.Pagination, func(key []byte, value []byte) error {
		var client types.VerifiedClientID
		if err := k.cdc.Unmarshal(value, &client); err != nil {
			return err
		}

		verifiedClientIDs = append(verifiedClientIDs, client.ClientID)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryVerifiedClientIdsResponse{
		ClientIds:  verifiedClientIDs,
		Pagination: pageRes,
	}, nil
}
