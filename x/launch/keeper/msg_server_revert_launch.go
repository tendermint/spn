package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) RevertLaunch(goCtx context.Context, msg *types.MsgRevertLaunch) (*types.MsgRevertLaunchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainNotFound, msg.ChainID)
	}

	// Check sender is the coordinator of the chain
	coordinatorID, found := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}
	if chain.CoordinatorID != coordinatorID {
		return nil, sdkerrors.Wrapf(
			profiletypes.ErrCoordInvalid, 
			"coordinator of the chain is %v", 
			chain.CoordinatorID,
		)
	}

	if !chain.LaunchTriggered {
		return nil, sdkerrors.Wrap(types.ErrLaunchNotTriggered, msg.ChainID)
	}

	// The LaunchTimestamp must always be a non-zero value if LaunchTriggered is set
	if chain.LaunchTimestamp == 0 {
		return nil, spnerrors.Critical("LaunchTimestamp is not set while LaunchTriggered is set")
	}

	// We must wait for a specific delay once the chain is launched before being able to revert it
	if ctx.BlockTime().Unix() < chain.LaunchTimestamp+types.RevertDelay {
		return nil, sdkerrors.Wrap(types.ErrRevertDelayNotReached, msg.ChainID)
	}

	chain.LaunchTriggered = false
	chain.LaunchTimestamp = 0
	k.SetChain(ctx, chain)

	return &types.MsgRevertLaunchResponse{}, nil
}
