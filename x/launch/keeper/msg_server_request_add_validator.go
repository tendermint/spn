package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestAddValidator(
	goCtx context.Context,
	msg *types.MsgRequestAddValidator,
) (*types.MsgRequestAddValidatorResponse, error) {
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

	content := types.NewGenesisValidator(
		msg.LaunchID,
		msg.ValAddress,
		msg.GenTx,
		msg.ConsPubKey,
		msg.SelfDelegation,
		msg.Peer,
	)
	request := types.Request{
		LaunchID:  msg.LaunchID,
		Creator:   msg.Creator,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	}

	var requestID uint64
	var err error
	approved := false
	if msg.Creator == coordAddress {
		err := ApplyRequest(ctx, k.Keeper, msg.LaunchID, request)
		if err != nil {
			return nil, err
		}
		approved = true
		err = ctx.EventManager().EmitTypedEvent(&types.EventValidatorAdded{
			GenesisValidator: *content.GetGenesisValidator(),
		})
	} else {
		requestID = k.AppendRequest(ctx, request)
		err = ctx.EventManager().EmitTypedEvent(&types.EventRequestCreated{
			Request: request,
		})
	}

	return &types.MsgRequestAddValidatorResponse{
		RequestID:    requestID,
		AutoApproved: approved,
	}, err
}
