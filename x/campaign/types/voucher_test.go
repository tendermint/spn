package types_test

import (
	"errors"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	campaign "github.com/tendermint/spn/x/campaign/types"
)

var (
	voucherCampaignID  = uint64(10)
	prefixedVoucherFoo = campaign.VoucherDenom(voucherCampaignID, "foo")
	prefixedVoucherBar = campaign.VoucherDenom(voucherCampaignID, "bar")
)

func TestCheckVouchers(t *testing.T) {
	tests := []struct {
		name       string
		campaignID uint64
		vouchers   sdk.Coins
		err        error
	}{
		{
			name:       "should allow check with one valid coin",
			campaignID: voucherCampaignID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdkmath.NewInt(100)),
			),
		},
		{
			name:       "should allow check with two valid coins",
			campaignID: voucherCampaignID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdkmath.NewInt(100)),
				sdk.NewCoin(prefixedVoucherBar, sdkmath.NewInt(200)),
			),
		},
		{
			name:       "should prevent check with one valid and one invalid coins",
			campaignID: voucherCampaignID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdkmath.NewInt(100)),
				sdk.NewCoin("foo", sdkmath.NewInt(200)),
			),
			err: errors.New("foo doesn't contain the voucher prefix v/10/"),
		},
		{
			name:       "should prevent check with one invalid coin",
			campaignID: voucherCampaignID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin("foo", sdkmath.NewInt(200)),
			),
			err: errors.New("foo doesn't contain the voucher prefix v/10/"),
		},
		{
			name:       "should prevent check with invalid campaign id",
			campaignID: 1000,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdkmath.NewInt(200)),
			),
			err: errors.New("v/10/foo doesn't contain the voucher prefix v/1000/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := campaign.CheckVouchers(tt.vouchers, tt.campaignID)
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
		campaignID uint64
		shares     campaign.Shares
		want       sdk.Coins
		err        error
	}{
		{
			name:       "test one share",
			campaignID: voucherCampaignID,
			shares:     tc.Shares(t, "10foo"),
			want:       tc.Vouchers(t, "10foo", voucherCampaignID),
		},
		{
			name:       "test two shares",
			campaignID: voucherCampaignID,
			shares:     tc.Shares(t, "10foo,11bar"),
			want:       tc.Vouchers(t, "10foo,11bar", voucherCampaignID),
		},
		{
			name:       "another campaign id",
			campaignID: 1000,
			shares:     tc.Shares(t, "10foo,11bar,12foobar"),
			want:       tc.Vouchers(t, "10foo,11bar,12foobar", 1000),
		},
		{
			name:       "invalid share prefix",
			campaignID: 1000,
			shares:     campaign.Shares(tc.Coins(t, "10t/foo")),
			err:        errors.New("t/foo doesn't contain the share prefix s/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := campaign.SharesToVouchers(tt.shares, tt.campaignID)
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
		campaignID uint64
		coin       string
		want       string
	}{
		{
			name:       "test 10/foo",
			campaignID: 10,
			coin:       "foo",
			want:       "v/10/foo",
		},
		{
			name:       "test 0/foo",
			campaignID: 0,
			coin:       "foo",
			want:       "v/0/foo",
		},
		{
			name:       "test empty denom",
			campaignID: 10,
			coin:       "",
			want:       "v/10/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := campaign.VoucherDenom(tt.campaignID, tt.coin)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestVouchersToShares(t *testing.T) {
	tests := []struct {
		name       string
		campaignID uint64
		vouchers   sdk.Coins
		want       campaign.Shares
		err        error
	}{
		{
			name:       "test one voucher",
			campaignID: voucherCampaignID,
			vouchers:   tc.Vouchers(t, "10foo", voucherCampaignID),
			want:       tc.Shares(t, "10foo"),
		},
		{
			name:       "test two vouchers",
			campaignID: voucherCampaignID,
			vouchers:   tc.Vouchers(t, "10foo,11bar", voucherCampaignID),
			want:       tc.Shares(t, "10foo,11bar"),
		},
		{
			name:       "wrong campaign id",
			campaignID: 1000,
			// use old coin syntax to write incorrect coins
			vouchers: tc.Coins(t, "10v/10/bar,11v/10/foo"),
			err:      errors.New("v/10/bar doesn't contain the voucher prefix v/1000/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := campaign.VouchersToShares(tt.vouchers, tt.campaignID)
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
		campaignID uint64
		denom      string
		want       string
	}{
		{
			name:       "foo voucher",
			campaignID: 10,
			denom:      prefixedVoucherFoo,
			want:       prefixedShareFoo,
		},
		{
			name:       "bar voucher",
			campaignID: 10,
			denom:      prefixedVoucherBar,
			want:       prefixedShareBar,
		},
		{
			name:       "invalid voucher",
			campaignID: 10,
			denom:      "t/bar",
			want:       "s/t/bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := campaign.VoucherToShareDenom(tt.campaignID, tt.denom)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestVoucherCampaign(t *testing.T) {
	tests := []struct {
		name       string
		denom      string
		campaignID uint64
		valid      bool
	}{
		{
			name:       "campaign is 0",
			denom:      "v/0/foo",
			campaignID: uint64(0),
			valid:      true,
		},
		{
			name:       "campaign is 50",
			denom:      "v/50/bar",
			campaignID: uint64(50),
			valid:      true,
		},
		{
			name:  "no voucher prefix",
			denom: "0/foo",
			valid: false,
		},
		{
			name:  "invalid format",
			denom: "v/0/foo/bar",
			valid: false,
		},
		{
			name:  "campaign ID is not a number",
			denom: "v/foo/foo",
			valid: false,
		},
		{
			name:  "empty campaign ID",
			denom: "v//foo",
			valid: false,
		},
		{
			name:  "actual denom is empty",
			denom: "v/0/",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			campaignID, err := campaign.VoucherCampaign(tt.denom)
			if !tt.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.EqualValues(t, tt.campaignID, campaignID)
		})
	}
}
