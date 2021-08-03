package types

import (
	"errors"
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"time"
)

// TODO: Determine default values
var (
	DefaultMinLaunchTime = uint64(time.Hour.Seconds() * 24)
	DefaultMaxLaunchTime = uint64(time.Hour.Seconds() * 24 * 7)

	MaxParametrableLaunchTime = uint64(time.Hour.Seconds() * 24 * 31)

	KeyMinLaunchTime = []byte("MinLaunchTime")
	KeyMaxLaunchTime = []byte("MaxLaunchTime")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// NewParams creates a new Params instance
func NewParams(minLaunchTime, maxLaunchTime uint64) Params {
	return Params{
		MinLaunchTime: minLaunchTime,
		MaxLaunchTime: maxLaunchTime,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinLaunchTime, &p.MinLaunchTime, validateLaunchTime),
		paramtypes.NewParamSetPair(KeyMaxLaunchTime, &p.MaxLaunchTime, validateLaunchTime),
	}
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