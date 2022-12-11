package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	projecttypes "github.com/tendermint/spn/x/project/types"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) EditChain(goCtx context.Context, msg *types.MsgEditChain) (*types.MsgEditChainResponse, error) {
	var (
		err error
		ctx = sdk.UnwrapSDKContext(goCtx)
	)

	// check if the metadata length is valid
	maxMetadataLength := k.MaxMetadataLength(ctx)
	if uint64(len(msg.Metadata)) > maxMetadataLength {
		return nil, sdkerrors.Wrapf(types.ErrInvalidMetadataLength,
			"metadata length %d is greater than maximum %d",
			len(msg.Metadata),
			maxMetadataLength,
		)
	}

	chain, found := k.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainNotFound, "%d", msg.LaunchID)
	}

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	if chain.CoordinatorID != coordID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the chain is %d",
			chain.CoordinatorID,
		))
	}

	if len(msg.Metadata) > 0 {
		chain.Metadata = msg.Metadata
	}

	if msg.SetProjectID {
		// check if chain already has id associated
		if chain.HasProject {
			return nil, sdkerrors.Wrapf(types.ErrChainHasProject,
				"project with id %d already associated with chain %d",
				chain.ProjectID,
				chain.LaunchID,
			)
		}

		// check if chain coordinator is project coordinator
		project, found := k.projectKeeper.GetProject(ctx, msg.ProjectID)
		if !found {
			return nil, sdkerrors.Wrapf(projecttypes.ErrProjectNotFound, "project with id %d not found", msg.ProjectID)
		}

		if project.CoordinatorID != chain.CoordinatorID {
			return nil, sdkerrors.Wrapf(profiletypes.ErrCoordInvalid,
				"coordinator of the project is %d, chain coordinator is %d",
				project.CoordinatorID,
				chain.CoordinatorID,
			)
		}

		chain.ProjectID = msg.ProjectID
		chain.HasProject = true

		err = k.projectKeeper.AddChainToProject(ctx, chain.ProjectID, chain.LaunchID)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrAddChainToProject, err.Error())
		}
	}

	k.SetChain(ctx, chain)
	return &types.MsgEditChainResponse{}, err
}
