package types_test

import (
	"errors"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	campaign "github.com/tendermint/spn/x/campaign/types"
)

var (
	campaignID         = uint64(10)
	prefixedVoucherFoo = campaign.VoucherName(campaignID, "foo")
	prefixedVoucherBar = campaign.VoucherName(campaignID, "bar")
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
			err: fmt.Errorf("v/10/foo doesn't contain the voucher prefix v/1000/"),
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
			name:       "invalid campaign id",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := campaign.SharesToVouchers(tt.campaignID, tt.shares)
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
			got := campaign.VoucherName(tt.campaignID, tt.coin)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestVouchersToShares(t *testing.T) {
	tests := []struct {
		name     string
		vouchers sdk.Coins
		want     campaign.Shares
	}{
		{
			name: "test one voucher",
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(10)),
			),
			want: campaign.Shares(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(10)),
			)),
		},
		{
			name: "test two vouchers",
			vouchers: sdk.NewCoins(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(10)),
				sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(11)),
			),
			want: campaign.Shares(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(10)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(11)),
			)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := campaign.VouchersToShares(tt.vouchers)
			require.Equal(t, tt.want, got)
		})
	}
}
