package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/claim/types"
)

func (k msgServer) ClaimInitial(goCtx context.Context, msg *types.MsgClaimInitial) (*types.MsgClaimInitialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgClaimInitialResponse{}, nil
}
