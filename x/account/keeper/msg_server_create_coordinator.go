package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/account/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			status.Error(codes.AlreadyExists,
				fmt.Sprintf("coordinator address already exist: %d", coord.CoordinatorId))
	}

	// Validate the coordinator address
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return &types.MsgCreateCoordinatorResponse{},
			status.Error(codes.InvalidArgument,
				fmt.Sprintf("invalid coordinator address (%s): %s", msg.Address, err.Error()))
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
