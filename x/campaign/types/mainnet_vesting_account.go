package types

// GetTotalShares return total shares for account and delayed vesting options
func (m MainnetVestingAccount) GetTotalShares() (Shares, error) {
	vestingShares, err := m.VestingOptions.GetDelayedVestingShare()
	if err != nil {
		return nil, err
	}
	totalShares := IncreaseShares(m.StartingShares, vestingShares)
	return totalShares, nil
}
