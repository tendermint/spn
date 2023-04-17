package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

// Implements LaunchHooks interface
var _ types.LaunchHooks = Keeper{}

// RequestCreated calls associated hook if registered
func (k Keeper) RequestCreated(
	ctx sdk.Context,
	creator string,
	launchID,
	requestID uint64,
	content types.RequestContent,
) {
	if k.hooks != nil {
		k.hooks.RequestCreated(
			ctx,
			creator,
			launchID,
			requestID,
			content,
		)
	}
}
