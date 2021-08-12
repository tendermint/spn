package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
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

	coordAddress, found := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return nil, spnerrors.Critical(
			fmt.Sprintf("Coordinator id not found: %d", chain.CoordinatorID))
	}
	if msg.Coordinator != coordAddress {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Coordinator)
	}

	// first check if the request exists
	request, found := k.GetRequest(ctx, msg.ChainID, msg.RequestID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrRequestNotFound,
			"request %d for chain %s not found",
			msg.RequestID,
			msg.ChainID,
		)
	}

	// perform request action
	k.RemoveRequest(ctx, msg.ChainID, request.RequestID)
	if msg.Approve {
		err := applyRequest(ctx, k.Keeper, msg.ChainID, request)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgSettleRequestResponse{}, nil
}
