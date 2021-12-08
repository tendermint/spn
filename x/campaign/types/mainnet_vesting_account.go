package types

import "errors"

// GetTotalShares return total shares for account and delayed vesting options
func (m MainnetVestingAccount) GetTotalShares() (Shares, error) {
	dv := m.VestingOptions.GetDelayedVesting()
	if dv == nil {
		return nil, errors.New("invalid vesting options type")
	}
	return dv.TotalShares, nil
}
