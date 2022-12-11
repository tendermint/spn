package keeper

import (
	"context"
	"errors"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/project/types"
)

func (k msgServer) UpdateTotalSupply(goCtx context.Context, msg *types.MsgUpdateTotalSupply) (*types.MsgUpdateTotalSupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	if project.MainnetInitialized {
		return nil, sdkerrors.Wrapf(types.ErrMainnetInitialized, "%d", msg.ProjectID)
	}

	// Validate provided totalSupply
	totalSupplyRange := k.TotalSupplyRange(ctx)
	if err := types.ValidateTotalSupply(msg.TotalSupplyUpdate, totalSupplyRange); err != nil {
		if errors.Is(err, types.ErrInvalidSupplyRange) {
			return nil, ignterrors.Critical(err.Error())
		}
		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, err.Error())
	}

	project.TotalSupply = types.UpdateTotalSupply(project.TotalSupply, msg.TotalSupplyUpdate)
	k.SetProject(ctx, project)

	err = ctx.EventManager().EmitTypedEvent(&types.EventProjectTotalSupplyUpdated{
		ProjectID:          project.ProjectID,
		CoordinatorAddress: msg.Coordinator,
		TotalSupply:        project.TotalSupply,
	})

	return &types.MsgUpdateTotalSupplyResponse{}, err
}
