package types

import (
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	// MinTimeframeSize defines the minimum allowed value for either `RegistrationPeriod` or `WithdrawalDelay`
	MinTimeframeSize = int64(time.Hour.Seconds() * 24) // One day
	// MaxTimeframeSize defines the maximum allowed value for either `RegistrationPeriod` or `WithdrawalDelay`
	MaxTimeframeSize = int64(time.Hour.Seconds() * 24 * 21) // 21 days

	KeyAllocationPrice       = []byte("AllocationPrice")
	KeyParticipationTierList = []byte("ParticipationTierList")
	KeyRegistrationPeriod    = []byte("RegistrationPeriod")
	KeyWithdrawalDelay       = []byte("WithdrawalDelay")

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
	DefaultRegistrationPeriod = int64(time.Hour.Seconds() * 24 * 7)  // One week
	DefaultWithdrawalDelay    = int64(time.Hour.Seconds() * 24 * 14) // Two weeks
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	allocationPrice AllocationPrice,
	participationTierList []Tier,
	registrationPeriod,
	withdrawalDelay int64,
) Params {
	return Params{
		AllocationPrice:       allocationPrice,
		ParticipationTierList: participationTierList,
		RegistrationPeriod:    registrationPeriod,
		WithdrawalDelay:       withdrawalDelay,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultAllocationPrice,
		DefaultParticipationTierList,
		DefaultRegistrationPeriod,
		DefaultWithdrawalDelay,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAllocationPrice, &p.AllocationPrice, validateAllocationPrice),
		paramtypes.NewParamSetPair(KeyParticipationTierList, &p.ParticipationTierList, validateParticipationTierList),
		paramtypes.NewParamSetPair(KeyRegistrationPeriod, &p.RegistrationPeriod, validateAllocationsUsageTimeFrame),
		paramtypes.NewParamSetPair(KeyWithdrawalDelay, &p.WithdrawalDelay, validateAllocationsUsageTimeFrame),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateAllocationPrice(p.AllocationPrice); err != nil {
		return err
	}

	return validateParticipationTierList(p.ParticipationTierList)
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
	for _, tier := range participationTierList {
		// check IDs are unique
		if _, ok = tiersIndexMap[tier.TierID]; ok {
			return fmt.Errorf("duplicated tier ID: %v", tier.TierID)
		}
		tiersIndexMap[tier.TierID] = struct{}{}

		if tier.RequiredAllocations <= 0 {
			return errors.New("required allocations must be greater than zero")
		}

		if err := validateTierBenefits(tier.Benefits); err != nil {
			return err
		}
	}

	return nil
}

func validateTierBenefits(b TierBenefits) error {
	if b.MaxBidAmount.IsNil() {
		return errors.New("max bid amount should be set")
	}

	if !b.MaxBidAmount.IsPositive() {
		return fmt.Errorf("max bid amount must be greater than zero")
	}

	return nil
}

// validateAllocationsUsageTimeFrame validates a time frame parameter related to allocations usage, specifically both
// RegistrationPeriod and WithdrawalDelay
func validateAllocationsUsageTimeFrame(v interface{}) error {
	timeFrame, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if timeFrame <= 0 {
		return errors.New("time frame must be positive")
	}

	if timeFrame < MinTimeframeSize || timeFrame > MaxTimeframeSize {
		return fmt.Errorf("time frame value must be in the range [%v, %v]", MinTimeframeSize, MaxTimeframeSize)
	}

	return nil
}
