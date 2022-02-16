package types

import (
	"fmt"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	ParamStoreKeyTotalSupplyRange = []byte("totalsupplyrange")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default campaign parameters
func DefaultParams() Params {
	return Params{
		TotalSupplyRange: TotalSupplyRange{
			MinTotalSupply: sdk.NewInt(1_000),                 // A Thousand
			MaxTotalSupply: sdk.NewInt(1_000_000_000_000_000), // One Quadrillion
		},
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
		paramtypes.NewParamSetPair(ParamStoreKeyTotalSupplyRange, &p.TotalSupplyRange, validateTotalSupplyRange),
	}
}

// ValidateBasic performs basic validation on campaign parameters.
func (p Params) ValidateBasic() error {
	if p.TotalSupplyRange.MinTotalSupply.LT(sdk.ZeroInt()) {
		return fmt.Errorf(
			"minimum total supply should be greater than one: %s", p.TotalSupplyRange.MinTotalSupply,
		)
	}
	if p.TotalSupplyRange.MaxTotalSupply.LT(p.TotalSupplyRange.MinTotalSupply) {
		return fmt.Errorf(
			"maximum total supply should be greater than greater or equal than minimum total supply: %s",
			p.TotalSupplyRange.MaxTotalSupply,
		)
	}

	return nil
}

func validateTotalSupplyRange(i interface{}) error {
	v, ok := i.(TotalSupplyRange)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.MinTotalSupply.LT(sdk.OneInt()) {
		return fmt.Errorf("parameter minTotalSupply cannot be less than one")
	}

	if v.MaxTotalSupply.LT(v.MinTotalSupply) {
		return fmt.Errorf("parameter maxTotalSupply cannot be less than minTotalSupply")
	}

	return nil
}
