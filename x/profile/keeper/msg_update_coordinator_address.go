package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) UpdateCoordinatorAddress(
	goCtx context.Context,
	msg *types.MsgUpdateCoordinatorAddress,
) (*types.MsgUpdateCoordinatorAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	coordByAddress, found := k.getCoordinatorByAddress(ctx, msg.Address)
	if !found {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			sdkerrors.Wrap(types.ErrCoordAddressNotFound, msg.Address)
	}

	// Check if the new coordinator address is already in the store
	newCoordAddr, found := k.getCoordinatorByAddress(ctx, msg.NewAddress)
	if found {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			sdkerrors.Wrap(types.ErrCoordAlreadyExist,
				fmt.Sprintf("new address already have a coordinator: %d", newCoordAddr.CoordinatorID))
	}

	coord, found := k.GetCoordinator(ctx, coordByAddress.CoordinatorID)
	if !found {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			ignterrors.Criticalf("a coordinator address is associated to a non-existent coordinator ID: %d",
				coordByAddress.CoordinatorID)
	}

	// Check if the coordinator is inactive
	if !coord.Active {
		return &types.MsgUpdateCoordinatorAddressResponse{},
			ignterrors.Criticalf("inactive coordinator address should not exist in store, ID: %d",
				coordByAddress.CoordinatorID)
	}

	coord.Address = msg.NewAddress

	// Remove the old coordinator by address and create a new one
	k.RemoveCoordinatorByAddress(ctx, msg.Address)
	k.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       msg.NewAddress,
		CoordinatorID: coord.CoordinatorID,
	})
	k.SetCoordinator(ctx, coord)

	return &types.MsgUpdateCoordinatorAddressResponse{},
		ctx.EventManager().EmitTypedEvent(
			&types.EventCoordinatorAddressUpdated{
				CoordinatorID: coord.CoordinatorID,
				NewAddress:    msg.NewAddress,
			})
}
