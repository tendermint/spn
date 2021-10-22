package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) UpdateCoordinatorAddress(
	goCtx context.Context,
	msg *types.MsgUpdateCoordinatorAddress,
) (*types.MsgUpdateCoordinatorAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	coordByAddress, found := k.GetCoordinatorByAddress(ctx, msg.Address)
	if !found {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			sdkerrors.Wrap(types.ErrCoordAddressNotFound, msg.Address)
	}

	// Check if the new coordinator address is already in the store
	newCoordAddr, found := k.GetCoordinatorByAddress(ctx, msg.NewAddress)
	if found {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			sdkerrors.Wrap(types.ErrCoordAlreadyExist,
				fmt.Sprintf("new address already have a coordinator: %d", newCoordAddr.CoordinatorId))
	}

	coord, found := k.GetCoordinator(ctx, coordByAddress.CoordinatorId)
	if !found {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			spnerrors.Criticalf("a coordinator address is associated to a non-existent coordinator ID: %d",
				coordByAddress.CoordinatorId)
	}
	coord.Address = msg.NewAddress

	// Remove the old coordinator by address and create a new one
	k.RemoveCoordinatorByAddress(ctx, msg.Address)
	k.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       msg.NewAddress,
		CoordinatorId: coord.CoordinatorId,
	})
	k.SetCoordinator(ctx, coord)

	return &types.MsgUpdateCoordinatorAddressResponse{}, nil
}
