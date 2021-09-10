package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestNewDelayedVesting(t *testing.T) {
	vesting := sample.Shares()
	endTime := time.Now().Unix()

	vestingOptions := types.NewShareDelayedVesting(vesting, endTime)

	delayedVesting := vestingOptions.GetDelayedVesting()
	require.NotNil(t, delayedVesting)
	require.True(t, sdk.Coins(vesting).IsEqual(sdk.Coins(delayedVesting.Vesting)))
	require.EqualValues(t, endTime, delayedVesting.EndTime)
}

func TestDelayedVesting_Validate(t *testing.T) {
	tests := []struct {
		name   string
		option types.ShareVestingOptions
		valid  bool
	}{
		{
			name: "vesting without shares",
			option: *types.NewShareDelayedVesting(
				nil,
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting with empty shares",
			option: *types.NewShareDelayedVesting(
				types.EmptyShares(),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting with invalid coins",
			option: *types.NewShareDelayedVesting(
				types.NewSharesFromCoins(sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}}),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting with invalid timestamp",
			option: *types.NewShareDelayedVesting(
				sample.Shares(),
				0,
			),
			valid: false,
		},
		{
			name: "valid account vesting",
			option: *types.NewShareDelayedVesting(
				sample.Shares(),
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
