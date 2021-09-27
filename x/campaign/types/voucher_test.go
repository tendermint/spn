package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	campaign "github.com/tendermint/spn/x/campaign/types"
)

var (
	prefixedVoucherFoo = campaign.VoucherPrefix + "foo"
	prefixedVoucherBar = campaign.VoucherPrefix + "bar"
)

func TestEmptyVoucher(t *testing.T) {
	voucher := campaign.EmptyVoucher()
	require.Equal(t, voucher, campaign.Voucher{})
}

func TestNewVoucher(t *testing.T) {
	_, err := campaign.NewVoucher("invalid")
	require.Error(t, err)

	voucher, err := campaign.NewVoucher("100foo")
	require.NoError(t, err)
	require.Equal(t, voucher, campaign.Voucher(sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(100))))
}

func TestNewVoucherFromCoin(t *testing.T) {
	voucher := campaign.NewVoucherFromCoin(sdk.NewCoin("bar", sdk.NewInt(200)))
	require.Equal(t, voucher, campaign.Voucher(
		sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(200)),
	))
}

func TestCheckVoucher(t *testing.T) {
	require.NoError(t, campaign.CheckVoucher(campaign.Voucher(
		sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(200)),
	)))
	require.Error(t, campaign.CheckVoucher(campaign.Voucher(
		sdk.NewCoin("foo", sdk.NewInt(100)),
	)))
}

func TestAddVoucher(t *testing.T) {
	for _, tc := range []struct {
		desc       string
		voucher    campaign.Voucher
		newVoucher campaign.Voucher
		expected   campaign.Voucher
		isError    bool
	}{
		{
			desc:    "increase empty set",
			voucher: campaign.EmptyVoucher(),
			newVoucher: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(100)),
			),
			isError: true,
		},
		{
			desc: "no new voucher",
			voucher: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(100)),
			),
			newVoucher: campaign.EmptyVoucher(),
			isError:    true,
		},
		{
			desc: "invalid coin denom",
			voucher: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(100)),
			),
			newVoucher: campaign.EmptyVoucher(),
			expected: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(100)),
			),
			isError: true,
		},
		{
			desc: "increase voucher",
			voucher: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(100)),
			),
			newVoucher: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(50)),
			),
			expected: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(150)),
			),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := campaign.AddVoucher(tc.voucher, tc.newVoucher)
			if tc.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestDecreaseVoucher(t *testing.T) {
	for _, tc := range []struct {
		desc       string
		voucher    campaign.Voucher
		toDecrease campaign.Voucher
		expected   campaign.Voucher
		isError    bool
	}{
		{
			desc:    "decrease empty set",
			voucher: campaign.EmptyVoucher(),
			toDecrease: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(100)),
			),
			isError: true,
		},
		{
			desc: "decrease from empty set",
			voucher: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(100)),
			),
			toDecrease: campaign.EmptyVoucher(),
			isError:    true,
		},
		{
			desc: "decrease to negative",
			voucher: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(100)),
			),
			toDecrease: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(100)),
			),
			isError: true,
		},
		{
			desc: "decrease normal set",
			voucher: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(100)),
			),
			toDecrease: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(30)),
			),
			expected: campaign.NewVoucherFromCoin(
				sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(70)),
			),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := campaign.DecreaseVoucher(tc.voucher, tc.toDecrease)
			if tc.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestIsEqualVoucher(t *testing.T) {
	type args struct {
		voucher1 campaign.Voucher
		voucher2 campaign.Voucher
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal voucher",
			args: args{
				voucher1: campaign.NewVoucherFromCoin(
					sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(1000)),
				),
				voucher2: campaign.NewVoucherFromCoin(
					sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(1000)),
				),
			},
			want: true,
		},
		{
			name: "not equal values",
			args: args{
				voucher1: campaign.NewVoucherFromCoin(
					sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(10)),
				),
				voucher2: campaign.NewVoucherFromCoin(
					sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(101)),
				),
			},
			want: false,
		},
		{
			name: "not equal denom",
			args: args{
				voucher1: campaign.NewVoucherFromCoin(
					sdk.NewCoin(prefixedVoucherFoo, sdk.NewInt(10)),
				),
				voucher2: campaign.NewVoucherFromCoin(
					sdk.NewCoin(prefixedVoucherBar, sdk.NewInt(101)),
				),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := campaign.IsEqualVoucher(tt.args.voucher1, tt.args.voucher2)
			require.True(t, got == tt.want)
		})
	}
}
