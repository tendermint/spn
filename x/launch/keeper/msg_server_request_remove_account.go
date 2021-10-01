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

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainNotFound, "%v", msg.ChainID)
	}

	if chain.IsMainnet {
		return nil, sdkerrors.Wrapf(types.ErrChainIsMainnet, "%v", msg.ChainID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%v", msg.ChainID)
	}

	coordAddress, found := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainInactive,
			"the chain %v coordinator has been deleted", chain.Id)
	}
	if msg.Creator != msg.Address && msg.Creator != coordAddress {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Creator)
	}

	var requestID uint64
	approved := false

	content := types.NewAccountRemoval(msg.Address)
	request := types.Request{
		ChainID:   msg.ChainID,
		Creator:   msg.Address,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	}

	if msg.Creator == coordAddress {
		err := ApplyRequest(ctx, k.Keeper, msg.ChainID, request)
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
