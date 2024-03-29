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
func SharesToVouchers(shares Shares, projectID uint64) (sdk.Coins, error) {
	if err := CheckShares(shares); err != nil {
		return nil, err
	}
	vouchers := make(sdk.Coins, len(shares))
	for i, coin := range shares {
		denom := strings.TrimPrefix(coin.Denom, SharePrefix)
		coin.Denom = VoucherDenom(projectID, denom)
		vouchers[i] = coin
	}
	return vouchers, nil
}

// CheckVouchers checks if given Vouchers are valid
func CheckVouchers(vouchers sdk.Coins, projectID uint64) error {
	for _, voucher := range vouchers {
		prefix := VoucherDenom(projectID, "")
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
func VouchersToShares(vouchers sdk.Coins, projectID uint64) (Shares, error) {
	if err := CheckVouchers(vouchers, projectID); err != nil {
		return nil, err
	}
	shares := make(Shares, len(vouchers))
	for i, coin := range vouchers {
		coin.Denom = VoucherToShareDenom(projectID, coin.Denom)
		shares[i] = coin
	}
	return shares, nil
}

// VoucherDenom returns the Voucher name with prefix
func VoucherDenom(projectID uint64, denom string) string {
	return fmt.Sprintf("%s%d%s%s", VoucherPrefix, projectID, VoucherSeparator, denom)
}

// VoucherToShareDenom remove the voucher prefix and add the share prefix
func VoucherToShareDenom(projectID uint64, denom string) string {
	prefix := VoucherDenom(projectID, "")
	shareDenom := strings.TrimPrefix(denom, prefix)
	return SharePrefix + shareDenom
}

// VoucherProject returns the project associated to a voucher denom
func VoucherProject(denom string) (uint64, error) {
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
