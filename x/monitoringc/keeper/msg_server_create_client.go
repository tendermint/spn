package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func (k msgServer) CreateClient(goCtx context.Context, msg *types.MsgCreateClient) (*types.MsgCreateClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	
	return &types.MsgCreateClientResponse{}, nil
}
