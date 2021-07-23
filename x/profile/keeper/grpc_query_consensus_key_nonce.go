package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ConsensusKeyNonce(c context.Context, req *types.QueryGetConsensusKeyNonceRequest) (*types.QueryGetConsensusKeyNonceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetConsensusKeyNonce(
		ctx,
		req.ConsAddress,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetConsensusKeyNonceResponse{ConsensusKeyNonce: &val}, nil
}
