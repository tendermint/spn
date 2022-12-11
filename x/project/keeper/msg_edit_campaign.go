package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/project/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) EditProject(goCtx context.Context, msg *types.MsgEditProject) (*types.MsgEditProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if the metadata length is valid
	maxMetadataLength := k.MaxMetadataLength(ctx)
	if uint64(len(msg.Metadata)) > maxMetadataLength {
		return nil, sdkerrors.Wrapf(types.ErrInvalidMetadataLength,
			"metadata length %d is greater than maximum %d",
			len(msg.Metadata),
			maxMetadataLength,
		)
	}

	project, found := k.GetProject(ctx, msg.ProjectID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrProjectNotFound, "%d", msg.ProjectID)
	}

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	if project.CoordinatorID != coordID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the project is %d",
			project.CoordinatorID,
		))
	}

	if len(msg.Name) > 0 {
		project.ProjectName = msg.Name
	}

	if len(msg.Metadata) > 0 {
		project.Metadata = msg.Metadata
	}

	k.SetProject(ctx, project)

	err = ctx.EventManager().EmitTypedEvent(&types.EventProjectInfoUpdated{
		ProjectID:         project.ProjectID,
		CoordinatorAddress: msg.Coordinator,
		ProjectName:       project.ProjectName,
		Metadata:           project.Metadata,
	})

	return &types.MsgEditProjectResponse{}, err
}
