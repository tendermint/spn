package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type MissionVoteHooks struct {
	k         Keeper
	missionID uint64
}

// NewMissionVoteHooks returns a GovHooks that triggers mission completion on voting for a proposal
func (k Keeper) NewMissionVoteHooks(missionID uint64) MissionVoteHooks {
	return MissionVoteHooks{k, missionID}
}

var _ govtypes.GovHooks = MissionVoteHooks{}

// AfterProposalVote completes mission when a vote is cast
func (h MissionVoteHooks) AfterProposalVote(ctx sdk.Context, _ uint64, voterAddr sdk.AccAddress) {
	// TODO: error handling
	_ = h.k.CompleteMission(ctx, h.missionID, voterAddr.String())
}

// AfterProposalSubmission implements GovHooks
func (h MissionVoteHooks) AfterProposalSubmission(_ sdk.Context, _ uint64) {
}

// AfterProposalDeposit implements GovHooks
func (h MissionVoteHooks) AfterProposalDeposit(_ sdk.Context, _ uint64, _ sdk.AccAddress) {
}

// AfterProposalFailedMinDeposit implements GovHooks
func (h MissionVoteHooks) AfterProposalFailedMinDeposit(_ sdk.Context, _ uint64) {
}

// AfterProposalVotingPeriodEnded implements GovHooks
func (h MissionVoteHooks) AfterProposalVotingPeriodEnded(_ sdk.Context, _ uint64) {
}
