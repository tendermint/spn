package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Voucher represents a coin Voucher
type Voucher sdk.Coin

const (
	// VoucherPrefix is the prefix used to represent a voucher denomination
	// A sdk.Coin containing this prefix must never be represented in a balance in the bank module
	VoucherPrefix = "v/"
)

// EmptyVoucher returns Voucher object that contains an empty voucher
func EmptyVoucher() Voucher {
	return Voucher(sdk.Coin{})
}

// NewVoucher returns new Voucher from coin (100atom,200iris...)
func NewVoucher(str string) (Voucher, error) {
	coin, err := sdk.ParseCoinNormalized(str)
	if err != nil {
		return Voucher{}, err
	}
	return NewVoucherFromCoin(coin), nil
}

// NewVoucherFromCoin returns new Voucher from the coin representation
func NewVoucherFromCoin(coin sdk.Coin) Voucher {
	coin.Denom = VoucherPrefix + coin.Denom
	return Voucher(coin)
}

// CheckVoucher checks if given Voucher are valid Voucher
func CheckVoucher(voucher Voucher) error {
	if !strings.HasPrefix(voucher.Denom, VoucherPrefix) {
		return fmt.Errorf(
			"%s doesn't contain the voucher prefix %s",
			voucher.Denom,
			VoucherPrefix,
		)
	}
	return nil
}

// IsEqualVoucher returns true if the two sets of Voucher have the same value
func IsEqualVoucher(voucher, newVoucher Voucher) bool {
	if voucher.Denom != newVoucher.Denom {
		return false
	}
	return sdk.Coin(voucher).IsEqual(sdk.Coin(newVoucher))
}

// AddVoucher sum Vouchers amount
func AddVoucher(voucher, newVoucher Voucher) (Voucher, error) {
	if voucher.Denom != newVoucher.Denom {
		return Voucher{}, fmt.Errorf(
			"invalid coin denominations: %s != %s",
			voucher.Denom,
			newVoucher.Denom,
		)
	}
	return Voucher(sdk.Coin(voucher).Add(sdk.Coin(newVoucher))), nil
}

// DecreaseVoucher subtracts vouchers amount
func DecreaseVoucher(voucher, toDecrease Voucher) (Voucher, error) {
	if voucher.Denom != toDecrease.Denom {
		return Voucher{}, fmt.Errorf(
			"invalid coin denominations: %s != %s",
			voucher.Denom,
			toDecrease.Denom,
		)
	}
	return Voucher(sdk.Coin(voucher).Sub(sdk.Coin(toDecrease))), nil
}
