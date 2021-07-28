package keeper

import (
	"context"
	spnerrors "github.com/tendermint/spn/pkg/errors"

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

	if !k.HasCoordinator(ctx, coordByAddress.CoordinatorId) {
		return &types.MsgDeleteCoordinatorResponse{},
			spnerrors.Critical("a coordinator address is associated to a non-existent coordinator ID")
	}
	coord := k.GetCoordinator(ctx, coordByAddress.CoordinatorId)

	k.RemoveCoordinatorByAddress(ctx, msg.Address)
	k.RemoveCoordinator(ctx, coord.CoordinatorId)
	return &types.MsgDeleteCoordinatorResponse{
		CoordinatorId: coord.CoordinatorId,
	}, nil
}
