package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/account/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) UpdateCoordinatorDescription(
	goCtx context.Context,
	msg *types.MsgUpdateCoordinatorDescription,
) (*types.MsgUpdateCoordinatorDescriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the coordinator address
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return &types.MsgUpdateCoordinatorDescriptionResponse{},
			status.Error(codes.InvalidArgument,
				fmt.Sprintf("invalid coordinator address (%s): %s", msg.Address, err.Error()))
	}

	// Check if the coordinator address is already in the store
	coordByAddress, found := k.GetCoordinatorByAddress(ctx, msg.Address)
	if !found {
		return &types.MsgUpdateCoordinatorDescriptionResponse{},
			status.Error(codes.NotFound,
				fmt.Sprintf("coordinator address not found: %s", msg.Address))
	}

	coord := k.GetCoordinator(ctx, coordByAddress.CoordinatorId)
	if len(msg.Description.Identity) > 0 {
		coord.Description.Identity = msg.Description.Identity
	}
	if len(msg.Description.Website) > 0 {
		coord.Description.Website = msg.Description.Website
	}
	if len(msg.Description.Details) > 0 {
		coord.Description.Details = msg.Description.Details
	}

	k.SetCoordinator(ctx, coord)
	return &types.MsgUpdateCoordinatorDescriptionResponse{}, nil
}
