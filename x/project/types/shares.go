package types

import (
	"fmt"
	"strings"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

// Shares represents the portion of a supply
type Shares sdk.Coins

const (
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
	shares := make(Shares, len(coins))
	for i, coin := range coins {
		coin.Denom = SharePrefix + coins[i].Denom
		shares[i] = coin
	}
	return shares
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

// IsEqualShares returns true if the two sets of Shares have the same value
func IsEqualShares(shares, newShares Shares) bool {
	return sdk.Coins(shares).IsEqual(sdk.Coins(newShares))
}

// IncreaseShares increases the number of shares
func IncreaseShares(shares, newShares Shares) Shares {
	return Shares(sdk.Coins(shares).Add(sdk.Coins(newShares)...))
}

// DecreaseShares decreases the number of shares or returns a error if shares can't be decreased
func DecreaseShares(shares, toDecrease Shares) (Shares, error) {
	decreasedCoins, negative := sdk.Coins(shares).SafeSub(sdk.Coins(toDecrease)...)
	if negative {
		return nil, errors.New("shares cannot be decreased to negative")
	}

	return Shares(decreasedCoins), nil
}

// IsTotalSharesReached checks if the provided shares overflow the total number of shares
func IsTotalSharesReached(shares Shares, maximumTotalShareNumber uint64) (bool, error) {
	if err := sdk.Coins(shares).Validate(); err != nil {
		return false, errors.Wrap(err, "invalid share")
	}
	if err := CheckShares(shares); err != nil {
		return false, errors.Wrap(err, "invalid share format")
	}

	for _, share := range shares {
		if share.Amount.Uint64() > maximumTotalShareNumber {
			return true, nil
		}
	}

	return false, nil
}

// IsAllLTE returns true iff for every denom in shares, the denom is present at
// a smaller or equal amount in sharesB.
func (shares Shares) IsAllLTE(cmpShares Shares) bool {
	return sdk.Coins(shares).IsAllLTE(sdk.Coins(cmpShares))
}

// AmountOf returns the amount of a denom from shares
func (shares Shares) AmountOf(denom string) int64 {
	return sdk.Coins(shares).AmountOf(denom).Int64()
}

// Empty returns true if there are no coins and false otherwise.
func (shares Shares) Empty() bool {
	return sdk.Coins(shares).Empty()
}

// String returns all shares comma separated
func (shares Shares) String() string {
	return sdk.Coins(shares).String()
}

// CoinsFromTotalSupply returns the coins from a total supply reflected by the shares
func (shares Shares) CoinsFromTotalSupply(totalSupply sdk.Coins, totalShareNumber uint64) (coins sdk.Coins, err error) {
	if totalShareNumber == 0 {
		return coins, errors.New("total share number can't be 0")
	}

	// set map for performance
	sharesMap := make(map[string]sdkmath.Int)
	for _, share := range shares {
		if share.Amount.Uint64() > totalShareNumber {
			return coins, fmt.Errorf(
				"share %s amount is greater than total share number %d > %d",
				share.Denom,
				share.Amount.Uint64(),
				totalShareNumber,
			)
		}

		sharesMap[share.Denom] = share.Amount
	}

	// check all coins from total supply
	for _, supply := range totalSupply {
		shareDenom := SharePrefix + supply.Denom
		if amount, ok := sharesMap[shareDenom]; ok {
			// coin balance = (supply * share) / total share
			coinBalance := (supply.Amount.Mul(amount)).Quo(sdkmath.NewIntFromUint64(totalShareNumber))

			if !coinBalance.IsZero() {
				coins = append(coins, sdk.NewCoin(supply.Denom, coinBalance))
			}
		}
	}

	return coins, nil
}
