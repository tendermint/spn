package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/account/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) DeleteCoordinator(
	goCtx context.Context,
	msg *types.MsgDeleteCoordinator,
) (*types.MsgDeleteCoordinatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the coordinator address
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return &types.MsgDeleteCoordinatorResponse{},
			status.Error(codes.InvalidArgument,
				fmt.Sprintf("invalid coordinator address (%s): %s", msg.Address, err.Error()))
	}

	// Check if the coordinator address is already in the store
	coordByAddress, found := k.GetCoordinatorByAddress(ctx, msg.Address)
	if !found {
		return &types.MsgDeleteCoordinatorResponse{},
			status.Error(codes.NotFound,
				fmt.Sprintf("coordinator address not found: %s", msg.Address))
	}

	coord := k.GetCoordinator(ctx, coordByAddress.CoordinatorId)
	k.RemoveCoordinatorByAddress(ctx, msg.Address)
	k.RemoveCoordinator(ctx, coord.CoordinatorId)
	return &types.MsgDeleteCoordinatorResponse{}, nil
}
