package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	claimkeeper "github.com/ignite/modules/x/claim/keeper"

	launchtypes "github.com/tendermint/spn/x/launch/types"
)

type MissionSendRequestHooks struct {
	k         claimkeeper.Keeper
	missionID uint64
}

// NewMissionSendRequestHooks returns a launch hook that triggers mission completion on sending a new request
func NewMissionSendRequestHooks(k claimkeeper.Keeper, missionID uint64) MissionSendRequestHooks {
	return MissionSendRequestHooks{k, missionID}
}

var _ launchtypes.LaunchHooks = MissionSendRequestHooks{}

// RequestCreated completes mission when a request is created
func (h MissionSendRequestHooks) RequestCreated(
	ctx sdk.Context,
	creator string,
	_,
	_ uint64,
	_ launchtypes.RequestContent,
) {
	h.k.CompleteMission(ctx, h.missionID, creator)
}
