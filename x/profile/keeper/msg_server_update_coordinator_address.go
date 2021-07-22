package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) UpdateCoordinatorAddress(
	goCtx context.Context,
	msg *types.MsgUpdateCoordinatorAddress,
) (*types.MsgUpdateCoordinatorAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the coordinator address
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			status.Error(codes.InvalidArgument,
				fmt.Sprintf("invalid coordinator address (%s): %s", msg.Address, err.Error()))
	}

	// Validate the new coordinator address
	_, err = sdk.AccAddressFromBech32(msg.NewAddress)
	if err != nil {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			status.Error(codes.InvalidArgument,
				fmt.Sprintf("invalid new coordinator address (%s): %s", msg.NewAddress, err.Error()))
	}

	// Check if the addresses are the same
	if msg.Address == msg.NewAddress {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			status.Error(codes.InvalidArgument,
				fmt.Sprintf("address are equal of new address (%s)", msg.Address))
	}

	// Check if the new coordinator address is already in the store
	newCoordAddr, found := k.GetCoordinatorByAddress(ctx, msg.NewAddress)
	if found {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			status.Error(codes.AlreadyExists,
				fmt.Sprintf("coordinator address already exist: %d", newCoordAddr.CoordinatorId))
	}

	coordByAddress, found := k.GetCoordinatorByAddress(ctx, msg.Address)
	if !found {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			status.Error(codes.NotFound,
				fmt.Sprintf("coordinator address not found: %s", msg.Address))
	}

	coord := k.GetCoordinator(ctx, coordByAddress.CoordinatorId)
	coord.Address = msg.NewAddress

	// Remove the old coordinator by address and create a new one
	k.RemoveCoordinatorByAddress(ctx, msg.Address)
	k.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       msg.NewAddress,
		CoordinatorId: coord.CoordinatorId,
	})
	k.SetCoordinator(ctx, coord)

	return &types.MsgUpdateCoordinatorAddressResponse{CoordinatorId: coord.CoordinatorId}, nil
}
