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
			previousCoins: tc.CoinsFromString(t, "1000foo,1000bar"),
			updatedCoins:  sdk.NewCoins(),
			wantedCoins:   tc.CoinsFromString(t,"1000foo,1000bar"),
		},
		{
			name:          "update from empty set",
			previousCoins: sdk.NewCoins(),
			updatedCoins:  tc.CoinsFromString(t,"1000foo,1000bar"),
			wantedCoins:   tc.CoinsFromString(t,"1000foo,1000bar"),
		},
		{
			name:          "update existing",
			previousCoins: tc.CoinsFromString(t,"3000foo,4000bar"),
			updatedCoins:  tc.CoinsFromString(t,"1000foo,2000bar"),
			wantedCoins:   tc.CoinsFromString(t,"1000foo,2000bar"),
		},
		{
			name:          "disjoint set",
			previousCoins: tc.CoinsFromString(t,"3000toto,4000tata"),
			updatedCoins:  tc.CoinsFromString(t,"1000foo,2000bar"),
			wantedCoins:   tc.CoinsFromString(t,"3000toto,4000tata,1000foo,2000bar"),
		},
		{
			name:          "new values",
			previousCoins: tc.CoinsFromString(t,"3000toto,4000tata"),
			updatedCoins:  tc.CoinsFromString(t,"1000foo,2000bar,5000toto,6000tata"),
			wantedCoins:   tc.CoinsFromString(t,"5000toto,6000tata,1000foo,2000bar"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newCoins := campaign.UpdateTotalSupply(tt.previousCoins, tt.updatedCoins)
			require.True(t, newCoins.IsEqual(tt.wantedCoins))
		})
	}
}
