package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
	"time"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyDecayInformation = []byte("DecayInformation")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(di DecayInformation) Params {
	return Params{
		DecayInformation: di,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DecayInformation{
		Enabled:    false,
		DecayStart: time.UnixMilli(0),
		DecayEnd:   time.UnixMilli(0),
	})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDecayInformation, &p.DecayInformation, validateDecayInformation),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return validateDecayInformation(p.DecayInformation)
}

// validateDecayInformation validates the DecayInformation param
func validateDecayInformation(v interface{}) error {
	decayInfo, ok := v.(DecayInformation)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	return decayInfo.Validate()
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
