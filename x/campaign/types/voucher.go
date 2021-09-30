package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// VoucherPrefix is the prefix used to represent a voucher denomination
	VoucherPrefix = "v/"
)

// SharesToVouchers returns new Coins vouchers from the Shares representation
func SharesToVouchers(shares Shares, campaignID uint64) (sdk.Coins, error) {
	if err := CheckShares(shares); err != nil {
		return nil, err
	}
	vouchers := make(sdk.Coins, len(shares))
	for i, coin := range shares {
		denom := strings.TrimPrefix(coin.Denom, SharePrefix)
		coin.Denom = VoucherDenom(campaignID, denom)
		vouchers[i] = coin
	}
	return vouchers, nil
}

// CheckVouchers checks if given Vouchers are valid
func CheckVouchers(vouchers sdk.Coins, campaignID uint64) error {
	for _, voucher := range vouchers {
		prefix := VoucherDenom(campaignID, "")
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

// VouchersToShares returns new Shares from the Coins vouchers representation
func VouchersToShares(vouchers sdk.Coins, campaignID uint64) (Shares, error) {
	if err := CheckVouchers(vouchers, campaignID); err != nil {
		return nil, err
	}
	shares := make(Shares, len(vouchers))
	for i, coin := range vouchers {
		coin.Denom = VoucherToShareDenom(campaignID, coin.Denom)
		shares[i] = coin
	}
	return shares, nil
}

// VoucherDenom returns the Voucher name with prefix
func VoucherDenom(campaignID uint64, denom string) string {
	return fmt.Sprintf("%s%d/%s", VoucherPrefix, campaignID, denom)
}

// VoucherToShareDenom remove the voucher prefix and add the share prefix
func VoucherToShareDenom(campaignID uint64, denom string) string {
	prefix := VoucherDenom(campaignID, "")
	shareDenom := strings.TrimPrefix(denom, prefix)
	return SharePrefix + shareDenom
}
