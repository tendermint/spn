package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewShareDelayedVesting(vesting Shares, endTime int64) *ShareVestingOptions {
	return &ShareVestingOptions{
		Options: &ShareVestingOptions_DelayedVesting{
			DelayedVesting: &ShareDelayedVesting{
				Vesting: vesting,
				EndTime: endTime,
			},
		},
	}
}

// GetDelayedVestingShare return the vesting share for delayed vesting options
func (m ShareVestingOptions) GetDelayedVestingShare() (Shares, error) {
	switch vestionOptions := m.Options.(type) {
	case *ShareVestingOptions_DelayedVesting:
		return vestionOptions.DelayedVesting.Vesting, nil
	default:
		return nil, errors.New("invalid vesting options type")
	}
}

// Validate check the ShareDelayedVesting object
func (m ShareVestingOptions) Validate() error {
	switch vestionOptions := m.Options.(type) {
	case *ShareVestingOptions_DelayedVesting:
		if sdk.Coins(vestionOptions.DelayedVesting.Vesting).Empty() {
			return errors.New("empty vesting shares for ShareDelayedVesting")
		}
		if !sdk.Coins(vestionOptions.DelayedVesting.Vesting).IsValid() {
			return fmt.Errorf(
				"invalid vesting shares for DelayedVesting: %s",
				sdk.Coins(vestionOptions.DelayedVesting.Vesting).String(),
			)
		}
		if vestionOptions.DelayedVesting.EndTime == 0 {
			return errors.New("end time for DelayedVesting cannot be 0")
		}
	default:
		return errors.New("unrecognized vesting options")
	}
	return nil
}
