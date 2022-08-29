package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) CreateCoordinator(
	goCtx context.Context,
	msg *types.MsgCreateCoordinator,
) (*types.MsgCreateCoordinatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the coordinator address is already in the store
	coordByAddr, found := k.getCoordinatorByAddress(ctx, msg.Address)
	if found {
		return &types.MsgCreateCoordinatorResponse{},
			sdkerrors.Wrap(types.ErrCoordAlreadyExist,
				fmt.Sprintf("coordinatorId: %d", coordByAddr.CoordinatorID))
	}

	coordID := k.AppendCoordinator(ctx, types.Coordinator{
		Address:     msg.Address,
		Description: msg.Description,
		Active:      true,
	})
	k.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       msg.Address,
		CoordinatorID: coordID,
	})

	return &types.MsgCreateCoordinatorResponse{
			CoordinatorID: coordID,
		}, ctx.EventManager().EmitTypedEvent(
			&types.EventCoordinatorCreated{
				CoordinatorID: coordID,
				Address:       msg.Address,
			})
}
