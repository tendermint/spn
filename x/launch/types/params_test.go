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
			name:   "invalid launch time range",
			params: NewParams(DefaultMaxLaunchTime, DefaultMinLaunchTime, DefaultRevertDelay, DefaultCampaignCreationFee),
			err:    errors.New("MinLaunchTime can't be higher than MaxLaunchTime"),
		},
		{
			name:   "valid params",
			params: NewParams(DefaultMinLaunchTime, DefaultMaxLaunchTime, DefaultRevertDelay, DefaultCampaignCreationFee),
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
			name:            "invalid interface",
			launchTimeRange: "test",
			err:             fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:            "invalid range - min is negative",
			launchTimeRange: NewLaunchTimeRange(-1, 1),
			err:             errors.New("MinLaunchTime can't be negative"),
		},
		{
			name:            "invalid range - too high",
			launchTimeRange: NewLaunchTimeRange(1, MaxParametrableLaunchTime+1),
			err:             errors.New("max parametrable launch time reached"),
		},
		{
			name:            "invalid range - max lower than min",
			launchTimeRange: NewLaunchTimeRange(10, 1),
			err:             errors.New("MinLaunchTime can't be higher than MaxLaunchTime"),
		},
		{
			name:            "max launch time",
			launchTimeRange: NewLaunchTimeRange(1, MaxParametrableLaunchTime),
		},
		{
			name:            "valid launch time",
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
			name:        "invalid interface",
			revertDelay: "test",
			err:         fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:        "invalid interface - too high",
			revertDelay: MaxParametrableRevertDelay + 1,
			err:         errors.New("max parametrable revert delay reached"),
		},
		{
			name:        "max revert delay",
			revertDelay: MaxParametrableRevertDelay,
		},
		{
			name:        "valid revert delay",
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
			name:        "invalid interface",
			creationFee: "test",
			err:         fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:        "invalid coin",
			creationFee: sdk.Coins{sdk.Coin{Denom: "foo", Amount: sdk.NewInt(-1)}},
			err:         errors.New("coin -1foo amount is not positive"),
		},
		{
			name:        "valid empty param",
			creationFee: DefaultChainCreationFee,
		},
		{
			name: "valid param",
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
