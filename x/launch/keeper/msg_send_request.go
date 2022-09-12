package keeper

import (
	"context"
	ignterrors "github.com/ignite/modules/errors"
	profiletypes "github.com/tendermint/spn/x/profile/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) SendRequest(
	goCtx context.Context,
	msg *types.MsgSendRequest,
) (*types.MsgSendRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainNotFound, "%d", msg.LaunchID)
	}

	// check if request is valid for mainnet
	err := msg.Content.IsValidForMainnet()
	if err != nil && chain.IsMainnet {
		return nil, sdkerrors.Wrap(types.ErrInvalidRequestForMainnet, err.Error())
	}

	// no request can be sent if the launch of the chain is triggered
	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	coord, found := k.profileKeeper.GetCoordinator(ctx, chain.CoordinatorID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainInactive,
			"the chain %d coordinator not found", chain.LaunchID)
	}

	// only chain with active coordinator can receive a request
	if !coord.Active {
		return nil, sdkerrors.Wrapf(profiletypes.ErrCoordInactive,
			"the chain %d coordinator is inactive", chain.LaunchID)
	}

	// create the request from the content
	request := types.Request{
		LaunchID:  msg.LaunchID,
		Creator:   msg.Creator,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   msg.Content,
		Status:    types.Request_PENDING,
	}

	var requestID uint64
	approved := false
	if msg.Creator == coord.Address {
		err := ApplyRequest(ctx, k.Keeper, chain, request, coord)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrRequestApplicationFailure, err.Error())
		}
		approved = true
		request.Status = types.Request_APPROVED
	}

	// deduct request fee if set
	requestFee := k.RequestFee(ctx)
	if !requestFee.Empty() {
		sender, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return nil, ignterrors.Criticalf("invalid coordinator bech32 address %s", err.Error())
		}
		if err = k.distrKeeper.FundCommunityPool(ctx, requestFee, sender); err != nil {
			return nil, err
		}
	}

	requestID = k.AppendRequest(ctx, request)
	err = ctx.EventManager().EmitTypedEvent(&types.EventRequestCreated{
		Creator: msg.Creator,
		Request: request,
	})

	return &types.MsgSendRequestResponse{
		RequestID:    requestID,
		AutoApproved: approved,
	}, err
}
