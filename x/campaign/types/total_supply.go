package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

// ValidateTotalSupply checks whether the total supply for each denom is within the provided total supply range
func ValidateTotalSupply(coins sdk.Coins, supplyRange TotalSupplyRange) error {
	// extra safety check, this should never happen
	if supplyRange.MaxTotalSupply.LT(supplyRange.MinTotalSupply) {
		return fmt.Errorf(
			"provided total supply range is invalid, min > max: [%s, %s]",
			supplyRange.MinTotalSupply.String(),
			supplyRange.MaxTotalSupply.String())
	}

	for _, coin := range coins {
		totalSupply := coin.Amount
		if totalSupply.LT(supplyRange.MinTotalSupply) {
			return fmt.Errorf(
				"total supply for %s cannot be less than minTotalSupply %s: %s",
				coin.Denom,
				supplyRange.MinTotalSupply.String(),
				coin.Amount.String())
		}
		if totalSupply.GT(supplyRange.MaxTotalSupply) {
			return fmt.Errorf(
				"total supply for %s cannot be greater than maxTotalSupply %s: %s",
				coin.Denom,
				supplyRange.MaxTotalSupply.String(),
				coin.Amount.String())
		}
	}
	return nil
}
