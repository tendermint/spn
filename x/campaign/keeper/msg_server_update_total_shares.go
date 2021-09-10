package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) UpdateTotalShares(goCtx context.Context, msg *types.MsgUpdateTotalShares) (*types.MsgUpdateTotalSharesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateTotalSharesResponse{}, nil
}
