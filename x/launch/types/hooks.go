package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// LaunchHooks event hooks for launch module
type LaunchHooks interface {
	RequestCreated(
		ctx sdk.Context,
		creator string,
		launchID,
		requestID uint64,
		content RequestContent,
	)
}

// MultiLaunchHooks combines multiple launch hooks
type MultiLaunchHooks []LaunchHooks

func NewMultiLaunchHooks(hooks ...LaunchHooks) MultiLaunchHooks {
	return hooks
}

func (h MultiLaunchHooks) RequestCreated(
	ctx sdk.Context,
	creator string,
	launchID,
	requestID uint64,
	content RequestContent,
) {
	for i := range h {
		h[i].RequestCreated(
			ctx,
			creator,
			launchID,
			requestID,
			content,
		)
	}
}
