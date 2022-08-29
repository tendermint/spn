package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestNewDelayedVesting(t *testing.T) {
	totalBalance := tc.Coins(t, "1000foo,500bar,2000toto")
	vesting := tc.Coins(t, "500foo,500bar")
	endTime := time.Now().Unix()

	t.Run("should return valid delayed vesting", func(t *testing.T) {
		vestingOptions := types.NewDelayedVesting(totalBalance, vesting, endTime)

		delayedVesting := vestingOptions.GetDelayedVesting()
		require.NotNil(t, delayedVesting)
		require.True(t, vesting.IsEqual(delayedVesting.Vesting))
		require.True(t, totalBalance.IsEqual(delayedVesting.TotalBalance))
		require.EqualValues(t, endTime, delayedVesting.EndTime)
	})
}

func TestDelayedVesting_Validate(t *testing.T) {
	sampleTotalBalance := tc.Coins(t, "1000foo,500bar,1000toto")
	sampleVesting := tc.Coins(t, "500foo,500bar")

	tests := []struct {
		name   string
		option types.VestingOptions
		valid  bool
	}{
		{
			name: "should prevent validate delayed vesting with no total balance",
			option: *types.NewDelayedVesting(
				nil,
				sample.Coins(r),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "should prevent validate delayed vesting with no vesting",
			option: *types.NewDelayedVesting(
				sample.Coins(r),
				nil,
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "should prevent validate delayed vesting with invalid total balance",
			option: *types.NewDelayedVesting(
				sdk.Coins{sdk.Coin{Denom: "", Amount: sdkmath.NewInt(10)}},
				sample.Coins(r),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "should prevent validate delayed vesting with invalid vesting",
			option: *types.NewDelayedVesting(
				sample.Coins(r),
				sdk.Coins{sdk.Coin{Denom: "", Amount: sdkmath.NewInt(10)}},
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "should prevent validate delayed vesting with total balance smaller than vesting",
			option: *types.NewDelayedVesting(
				tc.Coins(t, "1000foo,500bar,2000toto"),
				tc.Coins(t, "1000foo,501bar,2000toto"),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "should prevent validate delayed vesting with vesting denoms not being a subset of total balance",
			option: *types.NewDelayedVesting(
				tc.Coins(t, "1000foo,500bar"),
				tc.Coins(t, "1000foo,500bar,2000toto"),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "should prevent validate delayed vesting with invalid timestamp",
			option: *types.NewDelayedVesting(
				sampleTotalBalance,
				sampleVesting,
				0,
			),
			valid: false,
		},
		{
			name: "should validate valid delayed vesting",
			option: *types.NewDelayedVesting(
				sampleTotalBalance,
				sampleVesting,
				time.Now().Unix(),
			),
			valid: true,
		},
		{
			name: "should validate delayed vesting with vesting equal to total balance",
			option: *types.NewDelayedVesting(
				sampleVesting,
				sampleVesting,
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
