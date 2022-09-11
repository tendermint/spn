package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) TriggerLaunch(goCtx context.Context, msg *types.MsgTriggerLaunch) (*types.MsgTriggerLaunchResponse, error) {
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
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the chain is %d",
			chain.CoordinatorID,
		))
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	if msg.LaunchTime.Before(ctx.BlockTime().Add(k.LaunchTimeRange(ctx).MinLaunchTime)) {
		return nil, sdkerrors.Wrapf(types.ErrLaunchTimeTooLow, "%s", msg.LaunchTime.String())
	}
	if msg.LaunchTime.After(ctx.BlockTime().Add(k.LaunchTimeRange(ctx).MaxLaunchTime)) {
		return nil, sdkerrors.Wrapf(types.ErrLaunchTimeTooHigh, "%s", msg.LaunchTime.String())
	}

	// set launch timestamp
	chain.LaunchTriggered = true
	chain.LaunchTime = msg.LaunchTime

	// set revision height for monitoring IBC client
	chain.ConsumerRevisionHeight = ctx.BlockHeight()

	k.SetChain(ctx, chain)

	err = ctx.EventManager().EmitTypedEvent(&types.EventLaunchTriggered{
		LaunchID:        msg.LaunchID,
		LaunchTimeStamp: chain.LaunchTime,
	})

	return &types.MsgTriggerLaunchResponse{}, err
}
