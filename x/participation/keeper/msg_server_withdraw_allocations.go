package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/participation/types"
)

func (k msgServer) WithdrawAllocations(goCtx context.Context, msg *types.MsgWithdrawAllocations) (*types.MsgWithdrawAllocationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWithdrawAllocationsResponse{}, nil
}
