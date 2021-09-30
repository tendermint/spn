package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) UnredeemVouchers(goCtx context.Context, msg *types.MsgUnredeemVouchers) (*types.MsgUnredeemVouchersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUnredeemVouchersResponse{}, nil
}
