package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

// Shares represents the portion of a supply
type Shares sdk.Coins

const (
	// DefaultTotalShare is the default value of total share for an underlying supply asset
	DefaultTotalShare = 100000

	// SharePrefix is the prefix used to represent a share denomination
	// A sdk.Coin containing this prefix must never be represented in a balance in the bank module
	SharePrefix = "s/"
)

// NewShares returns new shares from comma-separated value (100atom,200iris...)
func NewShares(shareStr string) (Shares, error) {
	coins, err := sdk.ParseCoinsNormalized(shareStr)
	if err != nil {
		return nil, err
	}
	return NewSharesFromCoins(coins), nil
}

// NewSharesFromCoins returns new shares from the coins representation
func NewSharesFromCoins(coins sdk.Coins) Shares {
	for _, coin := range coins {
		coin.Denom = SharePrefix + coin.Denom
	}

	return Shares(coins)
}

// CheckShares checks if given shares are valid shares
func CheckShares(shares Shares) error {
	for _, coin := range shares {
		if !strings.HasPrefix(coin.Denom, SharePrefix) {
			fmt.Errorf("%s doesn't contain the share prefix %s", coin.Denom, SharePrefix)
		}
	}

	return nil
}