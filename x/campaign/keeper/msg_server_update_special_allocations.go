package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) UpdateSpecialAllocations(goCtx context.Context, msg *types.MsgUpdateSpecialAllocations) (*types.MsgUpdateSpecialAllocationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateSpecialAllocationsResponse{}, nil
}
