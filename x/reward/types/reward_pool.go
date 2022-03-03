package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewRewardPool returns a new RewardPool object
func NewRewardPool(launchID uint64, currentRewardHeight int64) RewardPool {
	return RewardPool{
		LaunchID:            launchID,
		CurrentRewardHeight: currentRewardHeight,
		Closed:              false,
	}
}

// Validate check the RewardPool object
func (m RewardPool) Validate() error {
	if m.InitialCoins.Empty() {
		return errors.New("empty reward pool coins")
	}
	if err := m.InitialCoins.Validate(); err != nil {
		return fmt.Errorf("invalid reward pool coins: %s", err)
	}
	if m.RemainingCoins.Empty() {
		return errors.New("empty reward pool coins")
	}
	if err := m.RemainingCoins.Validate(); err != nil {
		return fmt.Errorf("invalid reward pool coins: %s", err)
	}

	// check that coins have same denom set
	if !m.RemainingCoins.DenomsSubsetOf(m.InitialCoins) || m.RemainingCoins.Len() != m.InitialCoins.Len() {
		return fmt.Errorf("initial coins and remaining coins must be of the same denom set")
	}

	if m.RemainingCoins.IsAnyGTE(m.InitialCoins) {
		return errors.New("current coin cannot be greater than initial coin")
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
