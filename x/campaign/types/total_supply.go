package types

import (
	"fmt"


	sdkmath "cosmossdk.io/math"
	sdkerrors "cosmossdk.io/errors"
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
func ValidateTotalSupply(totalSupply sdk.Coins, supplyRange TotalSupplyRange) error {
	if err := supplyRange.ValidateBasic(); err != nil {
		return err
	}

	for _, coin := range totalSupply {
		totalSupply := coin.Amount
		if totalSupply.LT(supplyRange.MinTotalSupply) {
			return fmt.Errorf(
				"total supply for %s cannot be lower than minTotalSupply %s: %s",
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

// ValidateBasic performs basic validation on an instance of TotalSupplyRange
func (sr TotalSupplyRange) ValidateBasic() error {
	if sr.MinTotalSupply.IsNil() {
		return sdkerrors.Wrap(ErrInvalidSupplyRange, "minimum total supply should be set")
	}

	if sr.MaxTotalSupply.IsNil() {
		return sdkerrors.Wrap(ErrInvalidSupplyRange, "maximum total supply should be set")
	}

	if sr.MinTotalSupply.LT(sdkmath.OneInt()) {
		return sdkerrors.Wrap(ErrInvalidSupplyRange, "minimum total supply should be greater than one")
	}

	if sr.MaxTotalSupply.LT(sr.MinTotalSupply) {
		return sdkerrors.Wrap(ErrInvalidSupplyRange, "maximum total supply should be greater or equal than minimum total supply")
	}

	return nil
}
