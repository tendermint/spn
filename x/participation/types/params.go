package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyAllocationPrice           = []byte("AllocationPrice")
	KeyParticipationTierList     = []byte("ParticipationTierList")
	KeyWithdrawalAllocationDelay = []byte("WithdrawalAllocationDelay")

	// TODO: Determine the default values for these params
	DefaultAllocationPrice = AllocationPrice{
		Staking: sdk.OneInt(),
	}
	DefaultParticipationTierList = []Tier{
		{
			TierID:              0,
			RequiredAllocations: 0,
			Benefits: TierBenefits{
				MaxBidAmount: sdk.ZeroInt(),
			},
		},
	}
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

	// TODO implement validation
	_ = allocationPrice

	return nil
}

// validateParticipationTierList validates the ParticipationTierList param
func validateParticipationTierList(v interface{}) error {
	participationTierList, ok := v.([]Tier)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = participationTierList

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
