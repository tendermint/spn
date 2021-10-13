package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// VoucherSeparator is used in voucher denom to separate the denom component
	VoucherSeparator = "/"

	// VoucherPrefix is the prefix used to represent a voucher denomination
	VoucherPrefix = "v" + VoucherSeparator
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
	return fmt.Sprintf("%s%d%s%s", VoucherPrefix, campaignID, VoucherSeparator, denom)
}

// VoucherToShareDenom remove the voucher prefix and add the share prefix
func VoucherToShareDenom(campaignID uint64, denom string) string {
	prefix := VoucherDenom(campaignID, "")
	shareDenom := strings.TrimPrefix(denom, prefix)
	return SharePrefix + shareDenom
}

// VoucherCampaign returns the campaign associated to a voucher denom
func VoucherCampaign(denom string) (uint64, error) {
	if !strings.HasPrefix(denom, VoucherPrefix) {
		return 0, errors.New("no voucher prefix")
	}
	denom = strings.TrimPrefix(denom, VoucherPrefix)

	parsed := strings.Split(denom, VoucherSeparator)
	if len(parsed) != 2 {
		return 0, errors.New("invalid format")
	}
	if parsed[1] == "" {
		return 0, errors.New("actual denom is empty")
	}
	return strconv.ParseUint(parsed[0], 10, 64)
}
