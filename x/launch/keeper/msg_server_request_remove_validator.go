package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestRemoveValidator(
	goCtx context.Context,
	msg *types.MsgRequestRemoveValidator,
) (*types.MsgRequestRemoveValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainNotFound, "%d", msg.LaunchID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	coordAddress, found := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainInactive,
			"the chain %d coordinator has been deleted", chain.LaunchID)
	}
	if msg.Creator != msg.ValidatorAddress && msg.Creator != coordAddress {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Creator)
	}

	var requestID uint64
	var err error
	approved := false

	content := types.NewValidatorRemoval(msg.ValidatorAddress)
	request := types.Request{
		LaunchID:  msg.LaunchID,
		Creator:   msg.ValidatorAddress,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	}

	if msg.Creator == coordAddress {
		err := ApplyRequest(ctx, k.Keeper, msg.LaunchID, request)
		if err != nil {
			return nil, err
		}
		approved = true
		err = ctx.EventManager().EmitTypedEvent(&types.EventValidatorRemoved{
			GenesisValidatorAccount: msg.ValidatorAddress,
			LaunchID:                msg.LaunchID,
		})
	} else {
		requestID = k.AppendRequest(ctx, request)
		err = ctx.EventManager().EmitTypedEvent(&types.EventRequestCreated{
			Request: request,
		})
	}

	return &types.MsgRequestRemoveValidatorResponse{
		RequestID:    requestID,
		AutoApproved: approved,
	}, err
}
