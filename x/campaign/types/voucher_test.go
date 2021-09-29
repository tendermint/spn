package types_test

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	campaign "github.com/tendermint/spn/x/campaign/types"
)

var (
	campaignID         = uint64(10)
	prefixedVoucherFoo = campaign.VoucherDenom(campaignID, "foo")
	prefixedVoucherBar = campaign.VoucherDenom(campaignID, "bar")
)

func TestCheckVouchers(t *testing.T) {
	tests := []struct {
		name       string
		campaignID uint64
		vouchers   sdk.Coins
		err        error
	}{
		{
			name:       "one valid coin",
			campaignID: campaignID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(100)),
			),
		},
		{
			name:       "two valid coins",
			campaignID: campaignID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(200)),
			),
		},
		{
			name:       "one valid and one invalid coins",
			campaignID: campaignID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(100)),
				sdk.NewCoin("foo", sdk.NewInt(200)),
			),
			err: errors.New("foo doesn't contain the voucher prefix v/10/"),
		},
		{
			name:       "one invalid coin",
			campaignID: campaignID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin("foo", sdk.NewInt(200)),
			),
			err: errors.New("foo doesn't contain the voucher prefix v/10/"),
		},
		{
			name:       "invalid campaign id",
			campaignID: 1000,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(200)),
			),
			err: errors.New("v/10/foo doesn't contain the voucher prefix v/1000/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := campaign.CheckVouchers(tt.campaignID, tt.vouchers)
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
			campaignID: campaignID,
			shares: campaign.Shares(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(10)),
			)),
			want: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(10)),
			),
		},
		{
			name:       "test two shares",
			campaignID: campaignID,
			shares: campaign.Shares(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(10)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(11)),
			)),
			want: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(10)),
				sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(11)),
			),
		},
		{
			name:       "another campaign id",
			campaignID: 1000,
			shares: campaign.Shares(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(10)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(11)),
				sdk.NewCoin(prefixedShareFoobar, sdk.NewInt(12)),
			)),
			want: sdk.NewCoins(
				sdk.NewCoin("v/1000/foo", sdk.NewInt(10)),
				sdk.NewCoin("v/1000/bar", sdk.NewInt(11)),
				sdk.NewCoin("v/1000/foobar", sdk.NewInt(12)),
			),
		},
		{
			name:       "invalid share prefix",
			campaignID: 1000,
			shares: campaign.Shares(sdk.NewCoins(
				sdk.NewCoin("t/foo", sdk.NewInt(10)),
			)),
			err: errors.New("t/foo doesn't contain the share prefix s/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := campaign.SharesToVouchers(tt.campaignID, tt.shares)
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
			campaignID: campaignID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(10)),
			),
			want: campaign.Shares(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(10)),
			)),
		},
		{
			name:       "test two vouchers",
			campaignID: campaignID,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(10)),
				sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(11)),
			),
			want: campaign.Shares(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(10)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(11)),
			)),
		},
		{
			name:       "wrong campaign id",
			campaignID: 1000,
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(10)),
				sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(11)),
			),
			err: errors.New("v/10/bar doesn't contain the voucher prefix v/1000/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := campaign.VouchersToShares(tt.campaignID, tt.vouchers)
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
