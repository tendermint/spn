package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewDelayedVesting(vesting sdk.Coins, endTime int64) *VestingOptions {
	return &VestingOptions{
		Options: &VestingOptions_DelayedVesting{
			DelayedVesting: &DelayedVesting{
				Vesting: vesting,
				EndTime: endTime,
			},
		},
	}
}

// Validate check the DelayedVesting object
func (m VestingOptions) Validate() error {
	switch vestionOptions := m.Options.(type) {
	case *VestingOptions_DelayedVesting:
		if vestionOptions.DelayedVesting.Vesting.Empty() {
			return errors.New("empty vesting coins for DelayedVesting")
		}
		if !vestionOptions.DelayedVesting.Vesting.IsValid() {
			return fmt.Errorf("invalid vesting coins for DelayedVesting: %s", vestionOptions.DelayedVesting.Vesting.String())
		}
		if vestionOptions.DelayedVesting.EndTime == 0 {
			return errors.New("end time for DelayedVesting cannot be 0")
		}
	default:
		return errors.New("unrecognized vesting options")
	}
	return nil
}
