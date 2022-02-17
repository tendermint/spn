package types

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	// DefaultMinLaunchTime ...
	// TODO: set back this value to the default one
	// uint64(time.Hour.Seconds() * 24)
	DefaultMinLaunchTime = uint64(5)
	DefaultMaxLaunchTime = uint64(time.Hour.Seconds() * 24 * 7)

	// DefaultRevertDelay is the delay after the launch time when it is possible to revert the launch of the chain
	// Chain launch can be reverted on-chain when the actual chain launch failed (incorrect gentx, etc...)
	// This delay must be small be big enough to ensure nodes had the time to bootstrap\
	// This currently corresponds to 1 hour
	DefaultRevertDelay = int64(60 * 60)

	MaxParametrableLaunchTime  = uint64(time.Hour.Seconds() * 24 * 31)
	MaxParametrableRevertDelay = int64(time.Hour.Seconds() * 24)

	KeyMinLaunchTime = []byte("MinLaunchTime")
	KeyMaxLaunchTime = []byte("MaxLaunchTime")
	KeyRevertDelay   = []byte("RevertDelay")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(minLaunchTime, maxLaunchTime uint64, revertDelay int64) Params {
	return Params{
		MinLaunchTime: minLaunchTime,
		MaxLaunchTime: maxLaunchTime,
		RevertDelay:   revertDelay,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinLaunchTime,
		DefaultMaxLaunchTime,
		DefaultRevertDelay,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinLaunchTime, &p.MinLaunchTime, validateLaunchTime),
		paramtypes.NewParamSetPair(KeyMaxLaunchTime, &p.MaxLaunchTime, validateLaunchTime),
		paramtypes.NewParamSetPair(KeyRevertDelay, &p.RevertDelay, validateRevertDelay),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.MinLaunchTime > p.MaxLaunchTime {
		return errors.New("MinLaunchTime can't be higher than MaxLaunchTime")
	}

	if err := validateLaunchTime(p.MinLaunchTime); err != nil {
		return err
	}

	if err := validateLaunchTime(p.MaxLaunchTime); err != nil {
		return err
	}

	return validateRevertDelay(p.RevertDelay)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateLaunchTime(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v > MaxParametrableLaunchTime {
		return errors.New("max parametrable launch time reached")
	}
	return nil
}

func validateRevertDelay(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v > MaxParametrableRevertDelay {
		return errors.New("max parametrable revert delay reached")
	}

	if v <= 0 {
		return errors.New("revert delay parameter must be positive")
	}

	return nil
}
