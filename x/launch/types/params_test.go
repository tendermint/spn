package types

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	tests := []struct {
		name   string
		params Params
		err    error
	}{
		{
			name: "invalid range",
			params: Params{
				MinLaunchTime: DefaultMaxLaunchTime,
				MaxLaunchTime: DefaultMinLaunchTime,
			},
			err: errors.New("MinLaunchTime can't be higher than MaxLaunchTime"),
		},
		{
			name: "valid range",
			params: Params{
				MinLaunchTime: DefaultMinLaunchTime,
				MaxLaunchTime: DefaultMaxLaunchTime,
			},
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

func TestValidateLaunchTime(t *testing.T) {
	tests := []struct {
		name       string
		launchTime interface{}
		err        error
	}{
		{
			name:       "invalid interface",
			launchTime: "test",
			err:        fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:       "invalid interface",
			launchTime: MaxParametrableLaunchTime + 1,
			err:        errors.New("max parametrable launch time reached"),
		},
		{
			name:       "max launch time",
			launchTime: MaxParametrableLaunchTime,
		},
		{
			name:       "valid launch time",
			launchTime: uint64(time.Hour.Seconds() * 24),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateLaunchTime(tt.launchTime)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
