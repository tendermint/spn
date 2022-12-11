package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	"github.com/tendermint/spn/x/project/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) UpdateSpecialAllocations(goCtx context.Context, msg *types.MsgUpdateSpecialAllocations) (*types.MsgUpdateSpecialAllocationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	project, found := k.GetProject(ctx, msg.ProjectID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrProjectNotFound, "%d", msg.ProjectID)
	}

	// get the coordinator ID associated to the sender address
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

	// verify mainnet launch is not triggered
	mainnetLaunched, err := k.IsProjectMainnetLaunchTriggered(ctx, project.ProjectID)
	if err != nil {
		return nil, ignterrors.Critical(err.Error())
	}
	if mainnetLaunched {
		return nil, sdkerrors.Wrap(types.ErrMainnetLaunchTriggered, fmt.Sprintf(
			"mainnet %d launch is already triggered",
			project.MainnetID,
		))
	}

	// decrease allocated shares from current special allocations
	project.AllocatedShares, err = types.DecreaseShares(project.AllocatedShares, project.SpecialAllocations.TotalShares())
	if err != nil {
		return nil, ignterrors.Critical("project allocated shares should be bigger than current special allocations" + err.Error())
	}

	// increase with new special allocations
	project.AllocatedShares = types.IncreaseShares(project.AllocatedShares, msg.SpecialAllocations.TotalShares())

	// increase the project shares
	reached, err := types.IsTotalSharesReached(project.AllocatedShares, k.GetTotalShares(ctx))
	if err != nil {
		return nil, ignterrors.Criticalf("verified shares are invalid %s", err.Error())
	}
	if reached {
		return nil, sdkerrors.Wrapf(types.ErrTotalSharesLimit, "%d", msg.ProjectID)
	}

	project.SpecialAllocations = msg.SpecialAllocations
	k.SetProject(ctx, project)
	err = ctx.EventManager().EmitTypedEvents(
		&types.EventProjectSharesUpdated{
			ProjectID:         project.ProjectID,
			CoordinatorAddress: msg.Coordinator,
			AllocatedShares:    project.AllocatedShares,
		})
	return &types.MsgUpdateSpecialAllocationsResponse{}, err
}
