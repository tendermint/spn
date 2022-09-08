package types

import (
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewShareDelayedVesting return the ShareVestingOptions
func NewShareDelayedVesting(totalShare, vesting Shares, endTime time.Time) *ShareVestingOptions {
	return &ShareVestingOptions{
		Options: &ShareVestingOptions_DelayedVesting{
			DelayedVesting: &ShareDelayedVesting{
				TotalShares: totalShare,
				Vesting:     vesting,
				EndTime:     endTime,
			},
		},
	}
}

// Validate check the ShareDelayedVesting object
func (m ShareVestingOptions) Validate() error {
	switch vestionOptions := m.Options.(type) {
	case *ShareVestingOptions_DelayedVesting:
		dv := vestionOptions.DelayedVesting

		if sdk.Coins(dv.Vesting).Empty() {
			return errors.New("empty vesting shares for ShareDelayedVesting")
		}
		if !sdk.Coins(dv.Vesting).IsValid() {
			return fmt.Errorf(
				"invalid vesting shares for DelayedVesting: %s",
				sdk.Coins(dv.Vesting).String(),
			)
		}

		if !sdk.Coins(dv.TotalShares).IsValid() {
			return fmt.Errorf(
				"invalid total balance for DelayedVesting: %s",
				sdk.Coins(dv.TotalShares).String(),
			)
		}
		if !dv.Vesting.IsAllLTE(dv.TotalShares) {
			return errors.New("vesting is not a subset of the total shares")
		}

		if vestionOptions.DelayedVesting.EndTime.IsZero() {
			return errors.New("end time for DelayedVesting cannot be 0")
		}
	default:
		return errors.New("unrecognized vesting options")
	}
	return nil
}
