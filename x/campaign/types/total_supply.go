package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// UpdateTotalSupply returns updated total supply by adding new denoms and replacing existing ones
// This method doesn't check coins format
func UpdateTotalSupply(coins, updatedCoins sdk.Coins) sdk.Coins {
	// List denoms that are updated
	updateDenoms := make(map[string]struct{})
	for i := range updatedCoins {
		updateDenoms[updatedCoins.GetDenomByIndex(i)] = struct{}{}
	}

	// List coins that remains not updated
	notUpdated := sdk.NewCoins()
	for _, previousCoin := range coins {
		if _, ok := updateDenoms[previousCoin.Denom]; !ok {
			notUpdated = append(notUpdated, previousCoin)
		}
	}

	return append(updatedCoins, notUpdated...).Sort()
}