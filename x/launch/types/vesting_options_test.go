package types_test

import (
	"errors"
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
		err    error
	}{
		{
			name: "vesting without coins",
			option: *types.NewDelayedVesting(
				nil,
				time.Now().Unix(),
			),
			err: errors.New("invalid vesting coins for DelayedVesting"),
		}, {
			name: "vesting with invalid coins",
			option: *types.NewDelayedVesting(
				sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				time.Now().Unix(),
			),
			err: errors.New("invalid vesting coins for DelayedVesting: 10: the coin list is invalid"),
		}, {
			name: "vesting with invalid timestamp",
			option: *types.NewDelayedVesting(
				sample.Coins(),
				0,
			),
			err: errors.New("end time for DelayedVesting cannot be 0"),
		}, {
			name: "valid account vesting",
			option: *types.NewDelayedVesting(
				sample.Coins(),
				time.Now().Unix(),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.option.Validate()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
