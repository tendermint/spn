package types

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Shares represents the portion of a supply
type Shares sdk.Coins

const (
	// DefaultTotalShareNumber is the default number of total share for an underlying supply asset
	DefaultTotalShareNumber = 100000

	// SharePrefix is the prefix used to represent a share denomination
	// A sdk.Coin containing this prefix must never be represented in a balance in the bank module
	SharePrefix = "s/"
)

// EmptyShares returns shares object that contains no share
func EmptyShares() Shares {
	return Shares{}
}

// NewShares returns new shares from comma-separated coins (100atom,200iris...)
func NewShares(str string) (Shares, error) {
	coins, err := sdk.ParseCoinsNormalized(str)
	if err != nil {
		return nil, err
	}
	return NewSharesFromCoins(coins), nil
}

// NewSharesFromCoins returns new shares from the coins representation
func NewSharesFromCoins(coins sdk.Coins) Shares {
	for i := range coins {
		coins[i].Denom = SharePrefix + coins[i].Denom
	}

	return Shares(coins)
}

// CheckShares checks if given shares are valid shares
func CheckShares(shares Shares) error {
	for _, coin := range shares {
		if !strings.HasPrefix(coin.Denom, SharePrefix) {
			return fmt.Errorf("%s doesn't contain the share prefix %s", coin.Denom, SharePrefix)
		}
	}

	return nil
}

// IncreaseShares increases the number of shares
func IncreaseShares(shares, newShares Shares) Shares {
	return Shares(sdk.Coins(shares).Add(sdk.Coins(newShares)...))
}

// DecreaseShares decreases the number of shares or returns a error if shares can't be decreased
func DecreaseShares(shares, toDecrease Shares) (Shares, error) {
	decreasedCoins, negative := sdk.Coins(shares).SafeSub(sdk.Coins(toDecrease))
	if negative {
		return nil, errors.New("shares cannot be decreased to negative")
	}

	return Shares(decreasedCoins), nil
}

// IsTotalSharesReached checks if the provided shares overflow the total number of shares
// Denoms not specified in totalShares uses DefaultTotalShareNumber as the default number of total shares
func IsTotalSharesReached(shares, totalShares Shares) bool {
	// Check the explicitly defined total shares
	totalMap := make(map[string]uint64)
	for _, coin := range totalShares {
		totalMap[coin.Denom] = coin.Amount.Uint64()
	}

	for _, coin := range shares {
		// If the denom is not specified in total share, we compare the default total share number
		total, ok := totalMap[coin.Denom]
		if ok {
			if coin.Amount.Uint64() > total {
				return true
			}
		} else {
			if coin.Amount.Uint64() > DefaultTotalShareNumber {
				return true
			}
		}
	}

	// denoms defined in totalShares but not in shares are not checked
	// the number of shares for an undefined denom is 0 by default therefore the total is never reached
	return false
}
