package types_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
)

func TestNewDelayedVesting(t *testing.T) {
	totalShares := tc.Shares(t, "1000foo,1000bar,500toto")
	vesting := tc.Shares(t, "1000foo,500bar")
	endTime := time.Now()

	t.Run("should allow creation of valid delayed vesting", func(t *testing.T) {
		vestingOptions := types.NewShareDelayedVesting(totalShares, vesting, endTime)
		delayedVesting := vestingOptions.GetDelayedVesting()
		require.NotNil(t, delayedVesting)
		require.True(t, sdk.Coins(vesting).IsEqual(sdk.Coins(delayedVesting.Vesting)))
		require.EqualValues(t, endTime, delayedVesting.EndTime)
	})
}

func TestDelayedVesting_Validate(t *testing.T) {
	totalShares := tc.Shares(t, "1000foo,1000bar,500toto")
	vesting := tc.Shares(t, "1000foo,500bar")

	tests := []struct {
		name   string
		option types.ShareVestingOptions
		valid  bool
	}{
		{
			name: "should allow validation for valid account vesting",
			option: *types.NewShareDelayedVesting(
				totalShares,
				vesting,
				time.Now(),
			),
			valid: true,
		},
		{
			name: "should allow validation for same vesting as total shares",
			option: *types.NewShareDelayedVesting(
				totalShares,
				vesting,
				time.Now(),
			),
			valid: true,
		},
		{
			name:   "should prevent validation for invalid share vesting options",
			option: types.ShareVestingOptions{},
			valid:  false,
		},
		{
			name: "should prevent validation for no total shares",
			option: *types.NewShareDelayedVesting(
				nil,
				sample.Shares(r),
				time.Now(),
			),
			valid: false,
		},
		{
			name: "should prevent validation for no vesting",
			option: *types.NewShareDelayedVesting(
				sample.Shares(r),
				nil,
				time.Now(),
			),
			valid: false,
		},
		{
			name: "should prevent validation for empty vesting",
			option: *types.NewShareDelayedVesting(
				sample.Shares(r),
				types.EmptyShares(),
				time.Now(),
			),
			valid: false,
		},
		{
			name: "should prevent validation for total shares with invalid coins",
			option: *types.NewShareDelayedVesting(
				types.NewSharesFromCoins(sdk.Coins{sdk.Coin{Denom: "", Amount: sdkmath.NewInt(10)}}),
				sample.Shares(r),
				time.Now(),
			),
			valid: false,
		},
		{
			name: "should prevent validation for vesting with invalid coins",
			option: *types.NewShareDelayedVesting(
				sample.Shares(r),
				types.NewSharesFromCoins(sdk.Coins{sdk.Coin{Denom: "", Amount: sdkmath.NewInt(10)}}),
				time.Now(),
			),
			valid: false,
		},
		{
			name: "should prevent validation for total shares less than vesting",
			option: *types.NewShareDelayedVesting(
				tc.Shares(t, "1000foo,500bar,2000toto"),
				tc.Shares(t, "1000foo,501bar,2000toto"),
				time.Now(),
			),
			valid: false,
		},
		{
			name: "should prevent validation for vesting denoms not a subset of total shares",
			option: *types.NewShareDelayedVesting(
				tc.Shares(t, "1000foo,500bar"),
				tc.Shares(t, "1000foo,500bar,2000toto"),
				time.Now(),
			),
			valid: false,
		},
		{
			name: "should prevent validation for vesting with invalid timestamp",
			option: *types.NewShareDelayedVesting(
				totalShares,
				vesting,
				time.Time{},
			),
			valid: false,
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
