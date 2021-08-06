package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestAddValidator(goCtx context.Context, msg *types.MsgRequestAddValidator) (*types.MsgRequestAddValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRequestAddValidatorResponse{}, nil
}
