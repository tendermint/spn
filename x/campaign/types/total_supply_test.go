package types_test

import (
	"testing"

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
			name:          "no update",
			previousCoins: tc.Coins(t, "1000foo,1000bar"),
			updatedCoins:  sdk.NewCoins(),
			wantedCoins:   tc.Coins(t, "1000foo,1000bar"),
		},
		{
			name:          "update from empty set",
			previousCoins: sdk.NewCoins(),
			updatedCoins:  tc.Coins(t, "1000foo,1000bar"),
			wantedCoins:   tc.Coins(t, "1000foo,1000bar"),
		},
		{
			name:          "update existing",
			previousCoins: tc.Coins(t, "3000foo,4000bar"),
			updatedCoins:  tc.Coins(t, "1000foo,2000bar"),
			wantedCoins:   tc.Coins(t, "1000foo,2000bar"),
		},
		{
			name:          "disjoint set",
			previousCoins: tc.Coins(t, "3000toto,4000tata"),
			updatedCoins:  tc.Coins(t, "1000foo,2000bar"),
			wantedCoins:   tc.Coins(t, "3000toto,4000tata,1000foo,2000bar"),
		},
		{
			name:          "new values",
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
	newCoins := func(coinsStr string) sdk.Coins {
		coins, err := sdk.ParseCoinsNormalized(coinsStr)
		require.NoError(t, err)
		return coins
	}

	tests := []struct {
		name        string
		coins       sdk.Coins
		supplyRange campaign.TotalSupplyRange
		valid       bool
	}{
		{
			name:        "invalid supply range",
			coins:       newCoins("1000foo,1000bar"),
			supplyRange: campaign.NewTotalSupplyRange(sdk.NewInt(1_000), sdk.NewInt(100)),
			valid:       false,
		},
		{
			name:        "total supply less than min",
			coins:       newCoins("100foo,1000bar"),
			supplyRange: campaign.NewTotalSupplyRange(sdk.NewInt(1000), sdk.NewInt(10_000)),
			valid:       false,
		},
		{
			name:        "total supply more than max",
			coins:       newCoins("1000foo,10000bar"),
			supplyRange: campaign.NewTotalSupplyRange(sdk.NewInt(1000), sdk.NewInt(1000)),
			valid:       false,
		},
		{
			name:        "valid supply",
			coins:       newCoins("1000foo,1000bar"),
			supplyRange: campaign.NewTotalSupplyRange(sdk.NewInt(100), sdk.NewInt(1000)),
			valid:       true,
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
			name:        "min total supply lower than one",
			supplyRange: campaign.NewTotalSupplyRange(sdk.ZeroInt(), sdk.OneInt()),
			valid:       false,
		},
		{
			name:        "min total supply greater than max total supply",
			supplyRange: campaign.NewTotalSupplyRange(sdk.NewInt(2), sdk.OneInt()),
			valid:       false,
		},
		{
			name:        "valid total supply range",
			supplyRange: campaign.NewTotalSupplyRange(sdk.OneInt(), sdk.OneInt()),
			valid:       true,
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
