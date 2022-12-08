package types

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	tests := []struct {
		name   string
		params Params
		err    error
	}{
		{
			name: "should prevent validate params with invalid launch time range",
			params: NewParams(
				DefaultMaxLaunchTime,
				DefaultMinLaunchTime,
				DefaultRevertDelay,
				DefaultFee,
				DefaultFee,
				DefaultMaxMetadataLength,
			),
			err: errors.New("MinLaunchTime can't be higher than MaxLaunchTime"),
		},
		{
			name: "should validate valid params",
			params: NewParams(
				DefaultMinLaunchTime,
				DefaultMaxLaunchTime,
				DefaultRevertDelay,
				DefaultFee,
				DefaultFee,
				DefaultMaxMetadataLength,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateLaunchTimeRange(t *testing.T) {
	tests := []struct {
		name            string
		launchTimeRange interface{}
		err             error
	}{
		{
			name:            "should prevent validate launch time range with invalid interface",
			launchTimeRange: "test",
			err:             fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:            "should prevent validate launch time range where min is negative",
			launchTimeRange: NewLaunchTimeRange(-1, 1),
			err:             errors.New("MinLaunchTime can't be negative"),
		},
		{
			name:            "should prevent validate launch time range where max is too high",
			launchTimeRange: NewLaunchTimeRange(1, MaxParametrableLaunchTime+1),
			err:             errors.New("max parametrable launch time reached"),
		},
		{
			name:            "should prevent validate launch time range where max lower than min",
			launchTimeRange: NewLaunchTimeRange(10, 1),
			err:             errors.New("MinLaunchTime can't be higher than MaxLaunchTime"),
		},
		{
			name:            "should validate launch time range with max launch time",
			launchTimeRange: NewLaunchTimeRange(1, MaxParametrableLaunchTime),
		},
		{
			name:            "should validate valid launch time range",
			launchTimeRange: NewLaunchTimeRange(0, time.Hour*24),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateLaunchTimeRange(tt.launchTimeRange)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateRevertDelay(t *testing.T) {
	tests := []struct {
		name        string
		revertDelay interface{}
		err         error
	}{
		{
			name:        "should prevent validate revert delay with invalid interface",
			revertDelay: "test",
			err:         fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:        "should prevent validate revert delay too high",
			revertDelay: MaxParametrableRevertDelay + 1,
			err:         errors.New("max parametrable revert delay reached"),
		},
		{
			name:        "should prevent validate revert delay not positive",
			revertDelay: time.Duration(0),
			err:         errors.New("revert delay parameter must be positive"),
		},
		{
			name:        "should validate max revert delay",
			revertDelay: MaxParametrableRevertDelay,
		},
		{
			name:        "should validate valid revert delay",
			revertDelay: time.Minute,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRevertDelay(tt.revertDelay)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateFee(t *testing.T) {
	tests := []struct {
		name string
		fee  interface{}
		err  error
	}{
		{
			name: "should prevent validate creation fee with invalid interface",
			fee:  "test",
			err:  fmt.Errorf("invalid parameter type: string"),
		},
		{
			name: "should prevent validate creation fee with invalid coin",
			fee:  sdk.Coins{sdk.Coin{Denom: "foo", Amount: sdkmath.NewInt(-1)}},
			err:  errors.New("coin -1foo amount is not positive"),
		},
		{
			name: "should validate empty fee",
			fee:  DefaultFee,
		},
		{
			name: "should validate valid fee",
			fee: sdk.NewCoins(
				sdk.NewInt64Coin("foo", rand.Int63n(1000)+1),
				sdk.NewInt64Coin("bar", rand.Int63n(1000)+1),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateFee(tt.fee)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateMaxMetadataLength(t *testing.T) {
	tests := []struct {
		name              string
		maxMetadataLength interface{}
		err               error
	}{
		{
			name:              "invalid interface",
			maxMetadataLength: "test",
			err:               fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:              "invalid number type",
			maxMetadataLength: 1000,
			err:               fmt.Errorf("invalid parameter type: int"),
		},
		{
			name:              "valid param",
			maxMetadataLength: uint64(1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateMaxMetadataLength(tt.maxMetadataLength)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
