package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	campaign "github.com/tendermint/spn/x/campaign/types"
)

func TestUpdateTotalSupply(t *testing.T) {
	newCoins := func (coinsStr string) sdk.Coins {
		coins, err := sdk.ParseCoinsNormalized(coinsStr)
		require.NoError(t, err)
		return coins
	}

	tests := []struct {
		name       string
		previousCoins sdk.Coins
		updatedCoins sdk.Coins
		wantedCoins sdk.Coins
	}{
		{
			name: "no update",
			previousCoins: newCoins("1000foo,1000bar"),
			updatedCoins: sdk.NewCoins(),
			wantedCoins: newCoins("1000foo,1000bar"),
		},
		{
			name: "update from empty set",
			previousCoins: sdk.NewCoins(),
			updatedCoins: newCoins("1000foo,1000bar"),
			wantedCoins: newCoins("1000foo,1000bar"),
		},
		{
			name: "update existing",
			previousCoins: newCoins("3000foo,4000bar"),
			updatedCoins: newCoins("1000foo,2000bar"),
			wantedCoins: newCoins("1000foo,2000bar"),
		},
		{
			name: "disjoint set",
			previousCoins: newCoins("3000toto,4000tata"),
			updatedCoins: newCoins("1000foo,2000bar"),
			wantedCoins: newCoins("3000toto,4000tata,1000foo,2000bar"),
		},
		{
			name: "new values",
			previousCoins: newCoins("3000toto,4000tata"),
			updatedCoins: newCoins("1000foo,2000bar,5000toto,6000tata"),
			wantedCoins: newCoins("5000toto,6000tata,1000foo,2000bar"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newCoins := campaign.UpdateTotalSupply(tt.previousCoins, tt.updatedCoins)
			require.True(t, newCoins.IsEqual(tt.wantedCoins))
		})
	}
}