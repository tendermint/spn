package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) DeleteValidator(goCtx context.Context, msg *types.MsgDeleteValidator) (*types.MsgDeleteValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the validator address is already in the store
	_, found := k.GetValidator(ctx, msg.Address)
	if !found {
		return &types.MsgDeleteValidatorResponse{},
			sdkerrors.Wrap(types.ErrValidatorNotFound, msg.Address)
	}
	k.RemoveValidator(ctx, msg.Address)

	return &types.MsgDeleteValidatorResponse{}, nil
}
