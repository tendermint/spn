package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewRewardPool returns a new RewardPool object
func NewRewardPool(
	launchID,
	lastRewardHeight,
	currentRewardHeight uint64,
	provider string,
	coins sdk.Coins,
) *RewardPool {
	return &RewardPool{
		LaunchID:            launchID,
		LastRewardHeight:    lastRewardHeight,
		CurrentRewardHeight: currentRewardHeight,
		Provider:            provider,
		Coins:               coins,
	}
}

// Validate check the RewardPool object
func (m RewardPool) Validate() error {
	if m.Coins.Empty() {
		return errors.New("empty reward pool coins")
	}
	if err := m.Coins.Validate(); err != nil {
		return fmt.Errorf("invalid reward pool coins: %s", err)
	}
	if _, err := sdk.AccAddressFromBech32(m.Provider); err != nil {
		return fmt.Errorf("invalid provider address: %s", err)
	}
	if m.CurrentRewardHeight < m.LastRewardHeight {
		return fmt.Errorf(
			"current reward height (%d) is lower than the last reward height (%d)",
			m.CurrentRewardHeight,
			m.LastRewardHeight,
		)
	}
	return nil
}
