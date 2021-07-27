package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) DeleteValidator(goCtx context.Context, msg *types.MsgDeleteValidator) (*types.MsgDeleteValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the validator address is already in the store
	_, found := k.GetValidatorByAddress(ctx, msg.Address)
	if !found {
		return &types.MsgDeleteValidatorResponse{},
			sdkerrors.Wrap(types.ErrValidatorNotFound,
				fmt.Sprintf("validator: %s", msg.Address))
	}
	k.RemoveValidatorByAddress(ctx, msg.Address)

	return &types.MsgDeleteValidatorResponse{}, nil
}
