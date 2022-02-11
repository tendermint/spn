package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"
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
	coord, err := k.profileKeeper.GetActiveCoordinatorByAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	if chain.CoordinatorID != coord.CoordinatorID {
		return nil, sdkerrors.Wrapf(
			profiletypes.ErrCoordInvalid,
			"coordinator of the chain is %d",
			chain.CoordinatorID,
		)
	}

	if !chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrNotTriggeredLaunch, "%d", msg.LaunchID)
	}

	// The LaunchTimestamp must always be a non-zero value if LaunchTriggered is set
	if chain.LaunchTimestamp == 0 {
		return nil, spnerrors.Critical("LaunchTimestamp is not set while LaunchTriggered is set")
	}

	// We must wait for a specific delay once the chain is launched before being able to revert it
	if ctx.BlockTime().Unix() < chain.LaunchTimestamp+types.RevertDelay {
		return nil, sdkerrors.Wrapf(types.ErrRevertDelayNotReached, "%d", msg.LaunchID)
	}

	chain.LaunchTriggered = false
	chain.LaunchTimestamp = 0
	k.SetChain(ctx, chain)

	return &types.MsgRevertLaunchResponse{}, nil
}
