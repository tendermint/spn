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
	// int64(time.Hour.Seconds() * 24)
	DefaultMinLaunchTime = int64(5)
	DefaultMaxLaunchTime = int64(time.Hour.Seconds() * 24 * 7)

	// DefaultRevertDelay is the delay after the launch time when it is possible to revert the launch of the chain
	// Chain launch can be reverted on-chain when the actual chain launch failed (incorrect gentx, etc...)
	// This delay must be small be big enough to ensure nodes had the time to bootstrap\
	// This currently corresponds to 1 hour
	DefaultRevertDelay = int64(60 * 60)

	MaxParametrableLaunchTime  = int64(time.Hour.Seconds() * 24 * 31)
	MaxParametrableRevertDelay = int64(time.Hour.Seconds() * 24)

	KeyLaunchTimeRange = []byte("LaunchTimeRange")
	KeyRevertDelay     = []byte("RevertDelay")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewLaunchTimeRange creates a new LaunchTimeRange instance
func NewLaunchTimeRange(minLaunchTime, maxLaunchTime int64) LaunchTimeRange {
	return LaunchTimeRange{
		MinLaunchTime: minLaunchTime,
		MaxLaunchTime: maxLaunchTime,
	}
}

// NewParams creates a new Params instance
func NewParams(minLaunchTime, maxLaunchTime, revertDelay int64) Params {
	return Params{
		LaunchTimeRange: NewLaunchTimeRange(minLaunchTime, maxLaunchTime),
		RevertDelay:     revertDelay,
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
		paramtypes.NewParamSetPair(KeyLaunchTimeRange, &p.LaunchTimeRange, validateLaunchTimeRange),
		paramtypes.NewParamSetPair(KeyRevertDelay, &p.RevertDelay, validateRevertDelay),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateLaunchTimeRange(p.LaunchTimeRange); err != nil {
		return err
	}

	return validateRevertDelay(p.RevertDelay)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateLaunchTimeRange(i interface{}) error {
	v, ok := i.(LaunchTimeRange)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// it is enough to check that minLaunchTime is positive since it must be that minLaunchTime < maxLaunchTime
	if v.MinLaunchTime < 0 {
		return errors.New("MinLaunchTime can't be negative")
	}

	if v.MinLaunchTime > v.MaxLaunchTime {
		return errors.New("MinLaunchTime can't be higher than MaxLaunchTime")
	}

	// just need to check max launch time due to check above that guarantees correctness of the range
	if v.MaxLaunchTime > MaxParametrableLaunchTime {
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
