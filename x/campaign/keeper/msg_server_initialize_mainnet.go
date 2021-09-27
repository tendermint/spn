package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) InitializeMainnet(goCtx context.Context, msg *types.MsgInitializeMainnet) (*types.MsgInitializeMainnetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgInitializeMainnetResponse{}, nil
}
