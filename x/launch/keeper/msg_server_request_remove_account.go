package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestRemoveAccount(
	goCtx context.Context,
	msg *types.MsgRequestRemoveAccount,
) (*types.MsgRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainNotFound, "%d", msg.LaunchID)
	}

	if chain.IsMainnet {
		return nil, sdkerrors.Wrapf(
			types.ErrRemoveMainnetAccount,
			"the chain %d is a mainnet",
			msg.LaunchID,
		)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	coordAddress, found := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainInactive,
			"the chain %d coordinator has been deleted", chain.LaunchID)
	}
	if msg.Creator != msg.Address && msg.Creator != coordAddress {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Creator)
	}

	var requestID uint64
	approved := false

	content := types.NewAccountRemoval(msg.Address)
	request := types.Request{
		LaunchID:  msg.LaunchID,
		Creator:   msg.Address,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	}

	if msg.Creator == coordAddress {
		err := ApplyRequest(ctx, k.Keeper, msg.LaunchID, request)
		if err != nil {
			return nil, err
		}
		approved = true
	} else {
		requestID = k.AppendRequest(ctx, request)
	}

	return &types.MsgRequestResponse{
		RequestID:    requestID,
		AutoApproved: approved,
	}, nil
}
