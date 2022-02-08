package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) SettleRequest(
	goCtx context.Context,
	msg *types.MsgSettleRequest,
) (*types.MsgSettleRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainNotFound, "%d", msg.LaunchID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	coord, found := k.profileKeeper.GetCoordinator(ctx, chain.CoordinatorID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainInactive,
			"the chain %d coordinator not found", chain.LaunchID)
	}

	if !coord.Active {
		return nil, sdkerrors.Wrapf(profiletypes.ErrCoordInactive,
			"the chain %d coordinator inactive", chain.LaunchID)
	}

	if msg.Approve && msg.Signer != coord.Address {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Signer)
	}

	// first check if the request exists
	request, found := k.GetRequest(ctx, msg.LaunchID, msg.RequestID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrRequestNotFound,
			"request %d for chain %d not found",
			msg.RequestID,
			msg.LaunchID,
		)
	}

	if msg.Signer != request.Creator && msg.Signer != coord.Address {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Signer)
	}

	// perform request action
	k.RemoveRequest(ctx, msg.LaunchID, request.RequestID)
	if msg.Approve {
		err := ApplyRequest(ctx, k.Keeper, msg.LaunchID, request)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgSettleRequestResponse{}, nil
}
