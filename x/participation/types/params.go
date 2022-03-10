package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyAllocationPrice           = []byte("AllocationPrice")
	KeyParticipationTierList     = []byte("ParticipationTierList")
	KeyWithdrawalAllocationDelay = []byte("WithdrawalAllocationDelay")

	DefaultAllocationPrice = AllocationPrice{
		Bonded: sdk.NewInt(1000),
	}
	DefaultParticipationTierList = []Tier{
		{
			TierID:              1,
			RequiredAllocations: 1,
			Benefits: TierBenefits{
				MaxBidAmount: sdk.NewInt(1000),
			},
		},
		{
			TierID:              2,
			RequiredAllocations: 2,
			Benefits: TierBenefits{
				MaxBidAmount: sdk.NewInt(2000),
			},
		},
		{
			TierID:              3,
			RequiredAllocations: 5,
			Benefits: TierBenefits{
				MaxBidAmount: sdk.NewInt(10000),
			},
		},
		{
			TierID:              4,
			RequiredAllocations: 10,
			Benefits: TierBenefits{
				MaxBidAmount: sdk.NewInt(30000),
			},
		},
	}

	// TODO: Determine the default values for this param
	DefaultWithdrawalAllocationDelay = uint64(0)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	allocationPrice AllocationPrice,
	participationTierList []Tier,
	withdrawalAllocationDelay uint64,
) Params {
	return Params{
		AllocationPrice:           allocationPrice,
		ParticipationTierList:     participationTierList,
		WithdrawalAllocationDelay: withdrawalAllocationDelay,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultAllocationPrice,
		DefaultParticipationTierList,
		DefaultWithdrawalAllocationDelay,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAllocationPrice, &p.AllocationPrice, validateAllocationPrice),
		paramtypes.NewParamSetPair(KeyParticipationTierList, &p.ParticipationTierList, validateParticipationTierList),
		paramtypes.NewParamSetPair(KeyWithdrawalAllocationDelay, &p.WithdrawalAllocationDelay, validateWithdrawalAllocationDelay),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateAllocationPrice(p.AllocationPrice); err != nil {
		return err
	}

	if err := validateParticipationTierList(p.ParticipationTierList); err != nil {
		return err
	}

	return validateWithdrawalAllocationDelay(p.WithdrawalAllocationDelay)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateAllocationPrice validates the AllocationPrice param
func validateAllocationPrice(v interface{}) error {
	allocationPrice, ok := v.(AllocationPrice)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if allocationPrice.Bonded.IsNil() {
		return errors.New("value for 'bonded' should be set")
	}

	if !allocationPrice.Bonded.IsPositive() {
		return errors.New("value for 'bonded' must be greater than zero")
	}

	return nil
}

// validateParticipationTierList validates the ParticipationTierList param
func validateParticipationTierList(v interface{}) error {
	participationTierList, ok := v.([]Tier)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	tiersIndexMap := make(map[uint64]struct{})
	lastTierID := uint64(0)
	lastRequiredAllocations := uint64(0)
	lastTierBenefits := TierBenefits{MaxBidAmount: sdk.ZeroInt()}
	for _, tier := range participationTierList {
		// check IDs are unique
		if _, ok = tiersIndexMap[tier.TierID]; ok {
			return fmt.Errorf("duplicated tier ID: %v", tier.TierID)
		}
		tiersIndexMap[tier.TierID] = struct{}{}

		// check IDs are sorted
		if lastTierID > tier.TierID {
			return fmt.Errorf("tiers must be sorted by ID")
		}
		lastTierID = tier.TierID

		if tier.RequiredAllocations <= lastRequiredAllocations {
			return errors.New("required allocations must be strictly increasing and greater than zero")
		}
		lastRequiredAllocations = tier.RequiredAllocations

		if err := validateNextTierBenefits(tier.Benefits, lastTierBenefits); err != nil {
			return err
		}
		lastTierBenefits = tier.Benefits
	}

	return nil
}

func validateNextTierBenefits(next, last TierBenefits) error {
	if next.MaxBidAmount.IsNil() || last.MaxBidAmount.IsNil() {
		return errors.New("max bid amount should be set")
	}

	if next.MaxBidAmount.LTE(last.MaxBidAmount) || !next.MaxBidAmount.IsPositive() {
		return fmt.Errorf("max bid amount must be strictly increasing and greater than zero")
	}

	return nil
}

// validateWithdrawalAllocationDelay validates the WithdrawalAllocationDelay param
func validateWithdrawalAllocationDelay(v interface{}) error {
	withdrawalAllocationDelay, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = withdrawalAllocationDelay

	return nil
}
