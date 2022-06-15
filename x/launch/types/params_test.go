package types

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

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
			name:   "should prevent validate params with invalid launch time range",
			params: NewParams(DefaultMaxLaunchTime, DefaultMinLaunchTime, DefaultRevertDelay, DefaultChainCreationFee),
			err:    errors.New("MinLaunchTime can't be higher than MaxLaunchTime"),
		},
		{
			name:   "should validate valid params",
			params: NewParams(DefaultMinLaunchTime, DefaultMaxLaunchTime, DefaultRevertDelay, DefaultChainCreationFee),
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
			launchTimeRange: NewLaunchTimeRange(0, int64(time.Hour.Seconds()*24)),
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
			revertDelay: int64(0),
			err:         errors.New("revert delay parameter must be positive"),
		},
		{
			name:        "should validate max revert delay",
			revertDelay: MaxParametrableRevertDelay,
		},
		{
			name:        "should validate valid revert delay",
			revertDelay: int64(time.Minute.Seconds() * 1),
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

func TestValidateChainCreationFee(t *testing.T) {
	tests := []struct {
		name        string
		creationFee interface{}
		err         error
	}{
		{
			name:        "should prevent validate creation fee with invalid interface",
			creationFee: "test",
			err:         fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:        "should prevent validate creation fee with invalid coin",
			creationFee: sdk.Coins{sdk.Coin{Denom: "foo", Amount: sdk.NewInt(-1)}},
			err:         errors.New("coin -1foo amount is not positive"),
		},
		{
			name:        "should validate empty fee",
			creationFee: DefaultChainCreationFee,
		},
		{
			name: "should validate valid fee",
			creationFee: sdk.NewCoins(
				sdk.NewInt64Coin("foo", rand.Int63n(1000)+1),
				sdk.NewInt64Coin("bar", rand.Int63n(1000)+1),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateChainCreationFee(tt.creationFee)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
