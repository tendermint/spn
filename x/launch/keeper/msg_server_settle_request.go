package keeper

import (
	"context"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) SettleRequest(
	goCtx context.Context,
	msg *types.MsgSettleRequest,
) (*types.MsgSettleRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainNotFound, msg.ChainID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrap(types.ErrTriggeredLaunch, msg.ChainID)
	}

	coordAddress := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if msg.Coordinator != coordAddress {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Coordinator)
	}

	request, found := k.GetRequest(ctx, msg.ChainID, msg.RequestID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrRequestNotFound,
			"request %s for chain %s not found",
			msg.RequestID,
			msg.ChainID,
		)
	}

	k.RemoveRequest(ctx, msg.ChainID, msg.RequestID)

	if msg.Approve {
		cdc := codectypes.NewInterfaceRegistry()

		var content types.RequestContent
		if err := cdc.UnpackAny(request.Content, &content); err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidRequestContent, err.Error())
		}
		switch content.(type) {
		case *types.AccountRemoval:
			// TODO: handle account removal
		default:
			return nil, sdkerrors.Wrap(types.ErrInvalidRequestContent,
				"unknown request content type")
		}
	}

	return &types.MsgSettleRequestResponse{}, nil
}
