package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type MissionDelegationHooks struct {
	k         Keeper
	missionID uint64
}

// NewMissionDelegationHooks returns a StakingHooks that triggers mission completion on delegation for an account
func (k Keeper) NewMissionDelegationHooks(missionID uint64) MissionDelegationHooks {
	return MissionDelegationHooks{k, missionID}
}

var _ stakingtypes.StakingHooks = MissionDelegationHooks{}

// BeforeDelegationCreated completes mission when a delegation is performed
func (h MissionDelegationHooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, _ sdk.ValAddress) {
	// TODO: error handling
	_ = h.k.CompleteMission(ctx, h.missionID, delAddr.String())
}

// AfterValidatorCreated implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorCreated(_ sdk.Context, _ sdk.ValAddress) {
}

// AfterValidatorRemoved implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorRemoved(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
}

// BeforeDelegationSharesModified implements StakingHooks
func (h MissionDelegationHooks) BeforeDelegationSharesModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
}

// AfterDelegationModified implements StakingHooks
func (h MissionDelegationHooks) AfterDelegationModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
}

// BeforeValidatorSlashed implements StakingHooks
func (h MissionDelegationHooks) BeforeValidatorSlashed(_ sdk.Context, _ sdk.ValAddress, _ sdk.Dec) {
}

// BeforeValidatorModified implements StakingHooks
func (h MissionDelegationHooks) BeforeValidatorModified(_ sdk.Context, _ sdk.ValAddress) {
}

// AfterValidatorBonded implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorBonded(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
}

// AfterValidatorBeginUnbonding implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorBeginUnbonding(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
}

// BeforeDelegationRemoved implements StakingHooks
func (h MissionDelegationHooks) BeforeDelegationRemoved(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
}
