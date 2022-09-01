package keeper

import (
	"context"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) RevertLaunch(goCtx context.Context, msg *types.MsgRevertLaunch) (*types.MsgRevertLaunchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
		return nil, sdkerrors.Wrapf(
			profiletypes.ErrCoordInvalid,
			"coordinator of the chain is %d",
			chain.CoordinatorID,
		)
	}

	if !chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrNotTriggeredLaunch, "%d", msg.LaunchID)
	}

	if chain.MonitoringConnected {
		return nil, sdkerrors.Wrapf(types.ErrChainMonitoringConnected, "%d", msg.LaunchID)
	}

	// We must wait for a specific delay once the chain is launched before being able to revert it
	if ctx.BlockTime().Before(chain.LaunchTime.Add(k.RevertDelay(ctx))) {
		return nil, sdkerrors.Wrapf(types.ErrRevertDelayNotReached, "%d", msg.LaunchID)
	}

	chain.LaunchTriggered = false
	chain.LaunchTime = time.Unix(0, 0).UTC()
	k.SetChain(ctx, chain)

	// clear associated client IDs from monitoring
	k.monitoringcKeeper.ClearVerifiedClientIDs(ctx, msg.LaunchID)
	err = ctx.EventManager().EmitTypedEvent(&types.EventLaunchReverted{
		LaunchID: msg.LaunchID,
	})

	return &types.MsgRevertLaunchResponse{}, err
}
