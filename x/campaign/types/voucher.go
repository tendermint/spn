package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// VoucherPrefix is the prefix used to represent a voucher denomination
	// A sdk.Coin containing this prefix must never be represented in a balance in the bank module
	VoucherPrefix = "v/"
)

// SharesToVouchers returns new Coins vouchers from the Shares representation
func SharesToVouchers(campaignID uint64, shares Shares) sdk.Coins {
	vouchers := make(sdk.Coins, len(shares))
	for i, coin := range shares {
		denom := strings.TrimPrefix(coin.Denom, SharePrefix)
		coin.Denom = VoucherName(campaignID, denom)
		vouchers[i] = coin
	}
	return vouchers
}

// VouchersToShares returns new Shares from the Coins vouchers representation
func VouchersToShares(vouchers sdk.Coins) Shares {
	shares := make(Shares, len(vouchers))
	for i, coin := range vouchers {
		splitDenom := strings.Split(coin.Denom, "/")
		coinName := splitDenom[len(splitDenom)-1]
		coin.Denom = SharePrefix + coinName
		shares[i] = coin
	}
	return shares
}

// VoucherName returns the Voucher name with prefix
func VoucherName(campaignID uint64, coin string) string {
	return fmt.Sprintf("%s%d/%s", VoucherPrefix, campaignID, coin)
}

// CheckVoucher checks if given Voucher are valid Voucher
func CheckVoucher(campaignID uint64, vouchers sdk.Coins) error {
	for _, voucher := range vouchers {
		prefix := VoucherName(campaignID, "")
		if !strings.HasPrefix(voucher.Denom, prefix) {
			return fmt.Errorf(
				"%s doesn't contain the voucher prefix %s",
				voucher.Denom,
				prefix,
			)
		}
	}
	return nil
}
