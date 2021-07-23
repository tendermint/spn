package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) UpdateValidatorDescription(goCtx context.Context, msg *types.MsgUpdateValidatorDescription) (*types.MsgUpdateValidatorDescriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateValidatorDescriptionResponse{}, nil
}
