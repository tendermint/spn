package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func coinsStr(t *testing.T, str string) sdk.Coins {
	coins, err := sdk.ParseCoinsNormalized(str)
	require.NoError(t, err)
	return coins
}

func TestNewDelayedVesting(t *testing.T) {
	totalBalance := coinsStr(t, "1000foo500bar2000toto")
	vesting := coinsStr(t, "500foo500bar")
	endTime := time.Now().Unix()

	vestingOptions := types.NewDelayedVesting(totalBalance, vesting, endTime)

	delayedVesting := vestingOptions.GetDelayedVesting()
	require.NotNil(t, delayedVesting)
	require.True(t, vesting.IsEqual(delayedVesting.Vesting))
	require.True(t, totalBalance.IsEqual(delayedVesting.TotalBalance))
	require.EqualValues(t, endTime, delayedVesting.EndTime)
}

func TestDelayedVesting_Validate(t *testing.T) {
	sampleTotalBalance := coinsStr(t, "1000foo500bar1000toto")
	sampleVesting := coinsStr(t, "500foo500bar")

	tests := []struct {
		name   string
		option types.VestingOptions
		valid  bool
	}{
		{
			name: "no total balance",
			option: *types.NewDelayedVesting(
				nil,
				sample.Coins(),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "no vesting",
			option: *types.NewDelayedVesting(
				sample.Coins(),
				nil,
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting with invalid total balance",
			option: *types.NewDelayedVesting(
				sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				sample.Coins(),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting with invalid vesting",
			option: *types.NewDelayedVesting(
				sample.Coins(),
				sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting with total balance smaller than vesting",
			option: *types.NewDelayedVesting(
				coinsStr(t,"1000foo500bar2000toto"),
				coinsStr(t,"1000foo501bar2000toto"),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting contains coins not present in total balance",
			option: *types.NewDelayedVesting(
				coinsStr(t,"1000foo500bar"),
				coinsStr(t,"1000foo500bar2000toto"),
				time.Now().Unix(),
			),
			valid: false,
		},
		{
			name: "vesting with invalid timestamp",
			option: *types.NewDelayedVesting(
				sampleTotalBalance,
				sampleVesting,
				0,
			),
			valid: false,
		},
		{
			name: "valid account vesting",
			option: *types.NewDelayedVesting(
				sampleTotalBalance,
				sampleVesting,
				time.Now().Unix(),
			),
			valid: true,
		},
		{
			name: "vesting is equal to total balance",
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
