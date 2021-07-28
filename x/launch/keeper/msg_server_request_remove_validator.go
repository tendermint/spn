package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestRemoveValidator(goCtx context.Context, msg *types.MsgRequestRemoveValidator) (*types.MsgRequestRemoveValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainIdNotFound, msg.ChainID)
	}
	request := types.Request{
		ChainID:   msg.ChainID,
		Creator:   msg.ValidatorAddress,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   nil,
	}

	return &types.MsgRequestRemoveValidatorResponse{}, nil
}
