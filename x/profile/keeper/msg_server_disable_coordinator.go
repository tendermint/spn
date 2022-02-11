package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) DisableCoordinator(
	goCtx context.Context,
	msg *types.MsgDisableCoordinator,
) (*types.MsgDisableCoordinatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the coordinator address is already in the store
	coordByAddress, found := k.GetCoordinatorByAddress(ctx, msg.Address)
	if !found {
		return &types.MsgDisableCoordinatorResponse{},
			sdkerrors.Wrap(types.ErrCoordAddressNotFound, msg.Address)
	}

	coord, found := k.GetCoordinator(ctx, coordByAddress.CoordinatorID)
	if !found {
		return &types.MsgDisableCoordinatorResponse{},
			spnerrors.Criticalf("a coordinator address is associated to a non-existent coordinator ID: %d",
				coordByAddress.CoordinatorID)
	}

	// disable by setting to inactive and remove CoordByAddress
	coord.Active = false
	k.SetCoordinator(ctx, coord)
	k.RemoveCoordinatorByAddress(ctx, msg.Address)

	return &types.MsgDisableCoordinatorResponse{
		CoordinatorID: coord.CoordinatorID,
	}, nil
}
