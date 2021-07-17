package keeper

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/account/types"
)

func (k msgServer) CreateCoordinator(goCtx context.Context, msg *types.MsgCreateCoordinator) (*types.MsgCreateCoordinatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the coordinator address is already in the store
	coord, found := k.GetCoordinatorByAddress(ctx, msg.Address)
	if found {
		return &types.MsgCreateCoordinatorResponse{},
			status.Error(codes.AlreadyExists,
				fmt.Sprintf("coordinator address already exist: %d", coord.CoordinatorId))
	}
	
	coordID := k.AppendCoordinator(ctx, types.Coordinator{
		Address: msg.Address,
		Description: &types.CoordinatorDescription{
			Identity: msg.Identity,
			Website:  msg.Website,
			Details:  msg.Details,
		},
	})
	k.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       msg.Address,
		CoordinatorId: coordID,
	})

	return &types.MsgCreateCoordinatorResponse{
		CoordinatorId: coordID,
	}, nil
}
