package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// VestingOptions defines the interface for vesting options
type VestingOptions interface {
	Validate() error
}

var _ VestingOptions = &DelayedVesting{}

// Validate check the DelayedVesting object
func (g DelayedVesting) Validate() error {
	if g.Vesting.Empty() || !g.Vesting.IsValid() {
		return sdkerrors.Wrapf(ErrInvalidCoins, "invalid vesting coins for DelayedVesting: %s", g.Vesting.String())
	}
	if g.EndTime == 0 {
		return sdkerrors.Wrap(ErrInvalidTimestamp, "invalid end time for DelayedVesting")
	}
	return nil
}
