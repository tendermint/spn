package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) UpdateCoordinatorDescription(
	goCtx context.Context,
	msg *types.MsgUpdateCoordinatorDescription,
) (*types.MsgUpdateCoordinatorDescriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the coordinator address is already in the store
	coordByAddress, found := k.GetCoordinatorByAddress(ctx, msg.Address)
	if !found {
		return &types.MsgUpdateCoordinatorDescriptionResponse{},
			sdkerrors.Wrap(types.ErrCoordAddressNotFound, msg.Address)
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
	return &types.MsgUpdateCoordinatorDescriptionResponse{
		CoordinatorId: coord.CoordinatorId,
	}, nil
}
