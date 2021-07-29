package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) EditChain(goCtx context.Context, msg *types.MsgEditChain) (*types.MsgEditChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEditChainResponse{}, nil
}
