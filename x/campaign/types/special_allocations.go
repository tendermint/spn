package types

import (
	"github.com/pkg/errors"
)

// NewSpecialAllocations returns new special allocations
func NewSpecialAllocations(genesisDistribution, claimableAirdrop Shares) SpecialAllocations {
	return SpecialAllocations{
		GenesisDistribution: genesisDistribution,
		ClaimableAirdrop:    claimableAirdrop,
	}
}

// EmptySpecialAllocations returns special allocation with empty shares
func EmptySpecialAllocations() SpecialAllocations {
	return NewSpecialAllocations(EmptyShares(), EmptyShares())
}

// Validate validates the special allocation structure
func (sa SpecialAllocations) Validate() error {
	err := CheckShares(sa.GenesisDistribution)
	if err != nil {
		return errors.Wrap(err, "invalid genesis distribution")
	}
	err = CheckShares(sa.ClaimableAirdrop)
	if err != nil {
		return errors.Wrap(err, "invalid claimable airdrop")
	}
	return nil
}

// TotalShares returns the total amount of shares for the special allocations
func (sa SpecialAllocations) TotalShares() Shares {
	return IncreaseShares(sa.GenesisDistribution, sa.ClaimableAirdrop)
}
