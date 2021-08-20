package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestNewDelayedVesting(t *testing.T) {
	vesting := sample.Coins()
	endTime := time.Now().Unix()

	vestingOptions := types.NewDelayedVesting(vesting, endTime)

	delayedVesting := vestingOptions.GetDelayedVesting()
	require.NotNil(t, delayedVesting)
	require.True(t, vesting.IsEqual(delayedVesting.Vesting))
	require.EqualValues(t, endTime, delayedVesting.EndTime)
}

func TestDelayedVesting_Validate(t *testing.T) {
	tests := []struct {
		name   string
		option types.VestingOptions
		valid    bool
	}{
		{
			name: "vesting without coins",
			option: *types.NewDelayedVesting(
				nil,
				time.Now().Unix(),
			),
			valid: false,
		}, {
			name: "vesting with invalid coins",
			option: *types.NewDelayedVesting(
				sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				time.Now().Unix(),
			),
			valid: false,
		}, {
			name: "vesting with invalid timestamp",
			option: *types.NewDelayedVesting(
				sample.Coins(),
				0,
			),
			valid: false,
		}, {
			name: "valid account vesting",
			option: *types.NewDelayedVesting(
				sample.Coins(),
				time.Now().Unix(),
			),
			valid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.option.Validate()
			if tt.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
