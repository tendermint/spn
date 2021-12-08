package types

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewDelayedVesting(totalBalance, vesting sdk.Coins, endTime int64) *VestingOptions {
	return &VestingOptions{
		Options: &VestingOptions_DelayedVesting{
			DelayedVesting: &DelayedVesting{
				TotalBalance: totalBalance,
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
		dv := vestionOptions.DelayedVesting
		if dv.Vesting.Empty() {
			return errors.New("empty vesting coins for DelayedVesting")
		}
		if !dv.Vesting.IsValid() {
			return fmt.Errorf("invalid vesting coins for DelayedVesting: %s", dv.Vesting.String())
		}

		if !dv.TotalBalance.IsValid() {
			return fmt.Errorf("invalid total balance for DelayedVesting: %s", dv.TotalBalance.String())
		}

		if !dv.Vesting.IsAllLTE(dv.TotalBalance) {
			return errors.New("vesting denoms is not a subset of the total balance denoms")
		}

		if dv.EndTime == 0 {
			return errors.New("end time for DelayedVesting cannot be 0")
		}
	default:
		return errors.New("unrecognized vesting options")
	}
	return nil
}
