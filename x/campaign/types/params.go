package types

import (
	"fmt"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	ParamStoreKeyMinTotalSupply = []byte("mintotalsupply")
	ParamStoreKeyMaxTotalSupply = []byte("maxtotalsupply")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return Params{
		MinTotalSupply: sdk.NewInt(1_000),                 // A thousand
		MaxTotalSupply: sdk.NewInt(1_000_000_000_000_000), // One Quadrillion
	}
}

// String implements stringer interface
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyMinTotalSupply, &p.MinTotalSupply, validateMinTotalSupply),
		paramtypes.NewParamSetPair(ParamStoreKeyMaxTotalSupply, &p.MaxTotalSupply, validateMaxTotalSupply),
	}
}

// ValidateBasic performs basic validation on campaign parameters.
func (p Params) ValidateBasic() error {
	if p.MinTotalSupply.IsNegative() || p.MinTotalSupply.IsZero() {
		return fmt.Errorf(
			"minimum total supply should be greater than one: %s", p.MinTotalSupply,
		)
	}
	if p.MaxTotalSupply.LT(p.MinTotalSupply) || p.MaxTotalSupply.IsNegative() || p.MaxTotalSupply.IsZero() {
		return fmt.Errorf(
			"maximum total supply should be greater or equal than minimum total supply: %s", p.MaxTotalSupply,
		)
	}

	return nil
}

func validateMinTotalSupply(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("parameter cannot be negative")
	}

	if v.IsZero() {
		return fmt.Errorf("parameter cannot be zero")
	}

	return nil
}

func validateMaxTotalSupply(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("parameter cannot be negative")
	}

	if v.IsZero() {
		return fmt.Errorf("parameter cannot be zero")
	}

	return nil
}
