package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) UpdateTotalSupply(goCtx context.Context, msg *types.MsgUpdateTotalSupply) (*types.MsgUpdateTotalSupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateTotalSupplyResponse{}, nil
}
