package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	campaign "github.com/tendermint/spn/x/campaign/types"
)

func TestUpdateTotalSupply(t *testing.T) {
	tests := []struct {
		name          string
		previousCoins sdk.Coins
		updatedCoins  sdk.Coins
		wantedCoins   sdk.Coins
	}{
		{
			name:          "should perform no update",
			previousCoins: tc.Coins(t, "1000foo,1000bar"),
			updatedCoins:  sdk.NewCoins(),
			wantedCoins:   tc.Coins(t, "1000foo,1000bar"),
		},
		{
			name:          "should update from empty set",
			previousCoins: sdk.NewCoins(),
			updatedCoins:  tc.Coins(t, "1000foo,1000bar"),
			wantedCoins:   tc.Coins(t, "1000foo,1000bar"),
		},
		{
			name:          "should update existing coins",
			previousCoins: tc.Coins(t, "3000foo,4000bar"),
			updatedCoins:  tc.Coins(t, "1000foo,2000bar"),
			wantedCoins:   tc.Coins(t, "1000foo,2000bar"),
		},
		{
			name:          "should update disjoint coin set",
			previousCoins: tc.Coins(t, "3000toto,4000tata"),
			updatedCoins:  tc.Coins(t, "1000foo,2000bar"),
			wantedCoins:   tc.Coins(t, "3000toto,4000tata,1000foo,2000bar"),
		},
		{
			name:          "should add new values",
			previousCoins: tc.Coins(t, "3000toto,4000tata"),
			updatedCoins:  tc.Coins(t, "1000foo,2000bar,5000toto,6000tata"),
			wantedCoins:   tc.Coins(t, "5000toto,6000tata,1000foo,2000bar"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newCoins := campaign.UpdateTotalSupply(tt.previousCoins, tt.updatedCoins)
			require.True(t, newCoins.IsEqual(tt.wantedCoins))
		})
	}
}

func TestValidateTotalSupply(t *testing.T) {
	tests := []struct {
		name        string
		coins       sdk.Coins
		supplyRange campaign.TotalSupplyRange
		valid       bool
	}{
		{
			name:        "should allow validation for valid supply",
			coins:       tc.Coins(t, "1000foo,1000bar"),
			supplyRange: campaign.NewTotalSupplyRange(sdkmath.NewInt(100), sdkmath.NewInt(1000)),
			valid:       true,
		},
		{
			name:        "should prevent validation of invalid supply range",
			coins:       tc.Coins(t, "1000foo,1000bar"),
			supplyRange: campaign.NewTotalSupplyRange(sdkmath.NewInt(1_000), sdkmath.NewInt(100)),
			valid:       false,
		},
		{
			name:        "should prevent validation of total supply less than min",
			coins:       tc.Coins(t, "100foo,1000bar"),
			supplyRange: campaign.NewTotalSupplyRange(sdkmath.NewInt(1000), sdkmath.NewInt(10_000)),
			valid:       false,
		},
		{
			name:        "should prevent validation of total supply greater than max",
			coins:       tc.Coins(t, "1000foo,10000bar"),
			supplyRange: campaign.NewTotalSupplyRange(sdkmath.NewInt(1000), sdkmath.NewInt(1000)),
			valid:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := campaign.ValidateTotalSupply(tt.coins, tt.supplyRange)
			if tt.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestTotalSupplyRange_ValidateBasic(t *testing.T) {
	tests := []struct {
		name        string
		supplyRange campaign.TotalSupplyRange
		valid       bool
	}{
		{
			name:        "should allow validation with valid total supply range",
			supplyRange: campaign.NewTotalSupplyRange(sdkmath.OneInt(), sdkmath.OneInt()),
			valid:       true,
		},
		{
			name:        "should prevent validation with min total supply less than one",
			supplyRange: campaign.NewTotalSupplyRange(sdkmath.ZeroInt(), sdkmath.OneInt()),
			valid:       false,
		},
		{
			name:        "should prevent validation with min total supply greater than max total supply",
			supplyRange: campaign.NewTotalSupplyRange(sdkmath.NewInt(2), sdkmath.OneInt()),
			valid:       false,
		},
		{
			name:        "should prevent validation with uninitialized min total supply",
			supplyRange: campaign.NewTotalSupplyRange(sdkmath.Int{}, sdkmath.OneInt()),
			valid:       false,
		},
		{
			name:        "should prevent validation with prevent uninitialized max total supply",
			supplyRange: campaign.NewTotalSupplyRange(sdkmath.OneInt(), sdkmath.Int{}),
			valid:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.supplyRange.ValidateBasic()
			if tt.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
