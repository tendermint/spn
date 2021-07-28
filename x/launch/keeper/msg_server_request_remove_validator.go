package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestRemoveValidator(goCtx context.Context, msg *types.MsgRequestRemoveValidator) (*types.MsgRequestRemoveValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainIdNotFound, msg.ChainID)
	}

	content, err := types.AnyFromRequest()
	if err != nil {
		// TODO better error handler
		return nil, sdkerrors.Wrap(err, msg.ChainID)
	}

	request := types.Request{
		ChainID:   msg.ChainID,
		Creator:   msg.ValidatorAddress,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	}
	k.AppendRequest(ctx, request)

	return &types.MsgRequestRemoveValidatorResponse{}, nil
}
