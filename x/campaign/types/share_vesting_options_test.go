package types_test

import (
	"testing"
	"time"

	tc "github.com/tendermint/spn/testutil/constructor"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestNewDelayedVesting(t *testing.T) {
	totalShares := tc.Shares(t, "1000foo,1000bar,500toto")
	vesting := tc.Shares(t, "1000foo,500bar")
	endTime := time.Now().Unix()

	vestingOptions := types.NewShareDelayedVesting(totalShares, vesting, endTime)

	delayedVesting := vestingOptions.GetDelayedVesting()
	require.NotNil(t, delayedVesting)
	require.True(t, sdk.Coins(vesting).IsEqual(sdk.Coins(delayedVesting.Vesting)))
	require.EqualValues(t, endTime, delayedVesting.EndTime)
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
			name: "no total shares",
			option: *types.NewShareDelayedVesting(
				nil,
				sample.Shares(r),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "no vesting",
			option: *types.NewShareDelayedVesting(
				sample.Shares(r),
				nil,
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "empty vesting",
			option: *types.NewShareDelayedVesting(
				sample.Shares(r),
				types.EmptyShares(),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "total shares with invalid coins",
			option: *types.NewShareDelayedVesting(
				types.NewSharesFromCoins(sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}}),
				sample.Shares(r),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting with invalid coins",
			option: *types.NewShareDelayedVesting(
				sample.Shares(r),
				types.NewSharesFromCoins(sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}}),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "total shares smaller than vesting",
			option: *types.NewShareDelayedVesting(
				tc.Shares(t, "1000foo,500bar,2000toto"),
				tc.Shares(t, "1000foo,501bar,2000toto"),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting denoms is not a subset of total shares",
			option: *types.NewShareDelayedVesting(
				tc.Shares(t, "1000foo,500bar"),
				tc.Shares(t, "1000foo,500bar,2000toto"),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting with invalid timestamp",
			option: *types.NewShareDelayedVesting(
				totalShares,
				vesting,
				0,
			),
			valid: false,
		},
		{
			name: "valid account vesting",
			option: *types.NewShareDelayedVesting(
				totalShares,
				vesting,
				time.Now().Unix(),
			),
			valid: true,
		},
		{
			name: "same vesting as total shares",
			option: *types.NewShareDelayedVesting(
				totalShares,
				vesting,
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
