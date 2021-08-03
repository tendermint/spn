package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) TriggerLaunch(goCtx context.Context, msg *types.MsgTriggerLaunch) (*types.MsgTriggerLaunchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTriggerLaunchResponse{}, nil
}
