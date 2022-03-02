package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyDebugMode = []byte("DebugMode")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(debugMode bool) Params {
	return Params{
		DebugMode: debugMode,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(false)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			KeyDebugMode,
			&p.DebugMode,
			validateDebugMode,
		),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return validateDebugMode(p.DebugMode)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateDebugMode checks the param is a boolean
func validateDebugMode(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
