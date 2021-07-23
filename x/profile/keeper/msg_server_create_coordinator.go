package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) CreateCoordinator(
	goCtx context.Context,
	msg *types.MsgCreateCoordinator,
) (*types.MsgCreateCoordinatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the coordinator address is already in the store
	coord, found := k.GetCoordinatorByAddress(ctx, msg.Address)
	if found {
		return &types.MsgCreateCoordinatorResponse{},
			sdkerrors.Wrap(types.ErrCoordAlreadyExist,
				fmt.Sprintf("coordinatorId: %d", coord.CoordinatorId))
	}

	coordID := k.AppendCoordinator(ctx, types.Coordinator{
		Address:     msg.Address,
		Description: msg.Description,
	})
	k.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       msg.Address,
		CoordinatorId: coordID,
	})

	return &types.MsgCreateCoordinatorResponse{
		CoordinatorId: coordID,
	}, nil
}
