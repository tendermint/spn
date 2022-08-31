package types

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	// DefaultMinLaunchTime ...
	// TODO: set back this value to the default one
	// time.Hour * 24
	DefaultMinLaunchTime = time.Hour * 5
	DefaultMaxLaunchTime = time.Hour * 24 * 7

	// DefaultRevertDelay is the delay after the launch time when it is possible to revert the launch of the chain
	// Chain launch can be reverted on-chain when the actual chain launch failed (incorrect gentx, etc...)
	// This delay must be small be big enough to ensure nodes had the time to bootstrap\
	// This currently corresponds to 1 hour
	DefaultRevertDelay = time.Hour

	DefaultChainCreationFee = sdk.Coins(nil) // EmptyCoins

	MaxParametrableLaunchTime  = time.Hour * 24 * 31
	MaxParametrableRevertDelay = time.Hour * 24

	KeyLaunchTimeRange  = []byte("LaunchTimeRange")
	KeyRevertDelay      = []byte("RevertDelay")
	KeyChainCreationFee = []byte("ChainCreationFee")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewLaunchTimeRange creates a new LaunchTimeRange instance
func NewLaunchTimeRange(minLaunchTime, maxLaunchTime time.Duration) LaunchTimeRange {
	return LaunchTimeRange{
		MinLaunchTime: minLaunchTime,
		MaxLaunchTime: maxLaunchTime,
	}
}

// NewParams creates a new Params instance
func NewParams(minLaunchTime, maxLaunchTime, revertDelay time.Duration, chainCreationFee sdk.Coins) Params {
	return Params{
		LaunchTimeRange:  NewLaunchTimeRange(minLaunchTime, maxLaunchTime),
		RevertDelay:      revertDelay,
		ChainCreationFee: chainCreationFee,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinLaunchTime,
		DefaultMaxLaunchTime,
		DefaultRevertDelay,
		DefaultChainCreationFee,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLaunchTimeRange, &p.LaunchTimeRange, validateLaunchTimeRange),
		paramtypes.NewParamSetPair(KeyRevertDelay, &p.RevertDelay, validateRevertDelay),
		paramtypes.NewParamSetPair(KeyChainCreationFee, &p.ChainCreationFee, validateChainCreationFee),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateLaunchTimeRange(p.LaunchTimeRange); err != nil {
		return err
	}
	if err := validateRevertDelay(p.RevertDelay); err != nil {
		return err
	}
	return p.ChainCreationFee.Validate()
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
	v, ok := i.(time.Duration)
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

func validateChainCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return v.Validate()
}
