package types

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultMinTotalSupply = sdk.NewInt(100)                   // One hundred
	DefaultMaxTotalSupply = sdk.NewInt(1_000_000_000_000_000) // One Quadrillion

	ParamStoreKeyTotalSupplyRange = []byte("totalsupplyrange")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewTotalSupplyRange creates a new TotalSupplyRange instance
func NewTotalSupplyRange(minTotalSupply, maxTotalSupply sdk.Int) TotalSupplyRange {
	return TotalSupplyRange{
		MinTotalSupply: minTotalSupply,
		MaxTotalSupply: maxTotalSupply,
	}
}

// NewParams creates a new Params instance
func NewParams(minTotalSupply, maxTotalSupply sdk.Int) Params {
	return Params{
		TotalSupplyRange: NewTotalSupplyRange(minTotalSupply, maxTotalSupply),
	}
}

// DefaultParams returns default campaign parameters
func DefaultParams() Params {
	return NewParams(DefaultMinTotalSupply, DefaultMaxTotalSupply)
}

// String implements stringer interface
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyTotalSupplyRange, &p.TotalSupplyRange, validateTotalSupplyRange),
	}
}

// ValidateBasic performs basic validation on campaign parameters.
func (p Params) ValidateBasic() error {
	return validateTotalSupplyRange(p.TotalSupplyRange)
}

func validateTotalSupplyRange(i interface{}) error {
	v, ok := i.(TotalSupplyRange)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.MinTotalSupply.LT(sdk.OneInt()) {
		return errors.New("minimum total supply should be greater than one")
	}

	if v.MaxTotalSupply.LT(v.MinTotalSupply) {
		return errors.New("maximum total supply should be greater or equal than minimum total supply")
	}

	return nil
}
