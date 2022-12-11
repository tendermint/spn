package types_test

import (
	"errors"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	project "github.com/tendermint/spn/x/project/types"
)

var (
	voucherProjectID  = uint64(10)
	prefixedVoucherFoo = project.VoucherDenom(voucherProjectID, "foo")
	prefixedVoucherBar = project.VoucherDenom(voucherProjectID, "bar")
)

func TestCheckVouchers(t *testing.T) {
	tests := []struct {
		name       string
		projectID uint64
		vouchers   sdk.Coins
		err        error
	}{
		{
			name:       "should allow check with one valid coin",
			projectID: voucherProjectID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdkmath.NewInt(100)),
			),
		},
		{
			name:       "should allow check with two valid coins",
			projectID: voucherProjectID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdkmath.NewInt(100)),
				sdk.NewCoin(prefixedVoucherBar, sdkmath.NewInt(200)),
			),
		},
		{
			name:       "should prevent check with one valid and one invalid coins",
			projectID: voucherProjectID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdkmath.NewInt(100)),
				sdk.NewCoin("foo", sdkmath.NewInt(200)),
			),
			err: errors.New("foo doesn't contain the voucher prefix v/10/"),
		},
		{
			name:       "should prevent check with one invalid coin",
			projectID: voucherProjectID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin("foo", sdkmath.NewInt(200)),
			),
			err: errors.New("foo doesn't contain the voucher prefix v/10/"),
		},
		{
			name:       "should prevent check with invalid project id",
			projectID: 1000,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdkmath.NewInt(200)),
			),
			err: errors.New("v/10/foo doesn't contain the voucher prefix v/1000/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := project.CheckVouchers(tt.vouchers, tt.projectID)
			if tt.err != nil {
				require.Error(t, err)
				require.Equal(t, tt.err, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestSharesToVouchers(t *testing.T) {
	tests := []struct {
		name       string
		projectID uint64
		shares     project.Shares
		want       sdk.Coins
		err        error
	}{
		{
			name:       "should validate with one valid share",
			projectID: voucherProjectID,
			shares:     tc.Shares(t, "10foo"),
			want:       tc.Vouchers(t, "10foo", voucherProjectID),
		},
		{
			name:       "should validate with two valid shares",
			projectID: voucherProjectID,
			shares:     tc.Shares(t, "10foo,11bar"),
			want:       tc.Vouchers(t, "10foo,11bar", voucherProjectID),
		},
		{
			name:       "should prevent validation with invalid share prefix",
			projectID: 1000,
			shares:     project.Shares(tc.Coins(t, "10t/foo")),
			err:        errors.New("t/foo doesn't contain the share prefix s/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := project.SharesToVouchers(tt.shares, tt.projectID)
			if tt.err != nil {
				require.Error(t, err)
				require.Equal(t, tt.err, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestVoucherName(t *testing.T) {
	tests := []struct {
		name       string
		projectID uint64
		coin       string
		want       string
	}{
		{
			name:       "should prepend to 10/foo",
			projectID: 10,
			coin:       "foo",
			want:       "v/10/foo",
		},
		{
			name:       "should prepend to 0/foo",
			projectID: 0,
			coin:       "foo",
			want:       "v/0/foo",
		},
		{
			name:       "should prepend to empty denom",
			projectID: 10,
			coin:       "",
			want:       "v/10/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := project.VoucherDenom(tt.projectID, tt.coin)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestVouchersToShares(t *testing.T) {
	tests := []struct {
		name       string
		projectID uint64
		vouchers   sdk.Coins
		want       project.Shares
		err        error
	}{
		{
			name:       "should convert one voucher",
			projectID: voucherProjectID,
			vouchers:   tc.Vouchers(t, "10foo", voucherProjectID),
			want:       tc.Shares(t, "10foo"),
		},
		{
			name:       "should convert two vouchers",
			projectID: voucherProjectID,
			vouchers:   tc.Vouchers(t, "10foo,11bar", voucherProjectID),
			want:       tc.Shares(t, "10foo,11bar"),
		},
		{
			name:       "should fail with wrong project id",
			projectID: 1000,
			// use old coin syntax to write incorrect coins
			vouchers: tc.Coins(t, "10v/10/bar,11v/10/foo"),
			err:      errors.New("v/10/bar doesn't contain the voucher prefix v/1000/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := project.VouchersToShares(tt.vouchers, tt.projectID)
			if tt.err != nil {
				require.Error(t, err)
				require.Equal(t, tt.err, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestVoucherToShareDenom(t *testing.T) {
	tests := []struct {
		name       string
		projectID uint64
		denom      string
		want       string
	}{
		{
			name:       "should convert foo voucher",
			projectID: 10,
			denom:      prefixedVoucherFoo,
			want:       prefixedShareFoo,
		},
		{
			name:       "should convert bar voucher",
			projectID: 10,
			denom:      prefixedVoucherBar,
			want:       prefixedShareBar,
		},
		{
			name:       "should prepend to invalid voucher",
			projectID: 10,
			denom:      "t/bar",
			want:       "s/t/bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := project.VoucherToShareDenom(tt.projectID, tt.denom)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestVoucherProject(t *testing.T) {
	tests := []struct {
		name       string
		denom      string
		projectID uint64
		valid      bool
	}{
		{
			name:       "should allow with project 0",
			denom:      "v/0/foo",
			projectID: uint64(0),
			valid:      true,
		},
		{
			name:       "should allow with project 50",
			denom:      "v/50/bar",
			projectID: uint64(50),
			valid:      true,
		},
		{
			name:  "should fail with no voucher prefix",
			denom: "0/foo",
			valid: false,
		},
		{
			name:  "should fail with no invalid format",
			denom: "v/0/foo/bar",
			valid: false,
		},
		{
			name:  "should fail when project ID is not a number",
			denom: "v/foo/foo",
			valid: false,
		},
		{
			name:  "should fail with empty project ID",
			denom: "v//foo",
			valid: false,
		},
		{
			name:  "should fail when actual denom is empty",
			denom: "v/0/",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projectID, err := project.VoucherProject(tt.denom)
			if !tt.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.EqualValues(t, tt.projectID, projectID)
		})
	}
}
