package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) DeleteCoordinator(
	goCtx context.Context,
	msg *types.MsgDeleteCoordinator,
) (*types.MsgDeleteCoordinatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the coordinator address is already in the store
	coordByAddress, found := k.GetCoordinatorByAddress(ctx, msg.Address)
	if !found {
		return &types.MsgDeleteCoordinatorResponse{},
			sdkerrors.Wrap(types.ErrCoordAddressNotFound, msg.Address)
	}

	coord := k.GetCoordinator(ctx, coordByAddress.CoordinatorId)
	if (coord == types.Coordinator{}) {
		panic("Inconsistency error: coordinator id not exist into the keeper")
	}
	k.RemoveCoordinatorByAddress(ctx, msg.Address)
	k.RemoveCoordinator(ctx, coord.CoordinatorId)
	return &types.MsgDeleteCoordinatorResponse{
		CoordinatorId: coord.CoordinatorId,
	}, nil
}
