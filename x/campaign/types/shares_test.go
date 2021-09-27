package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	campaign "github.com/tendermint/spn/x/campaign/types"
)

var (
	prefixedShareFoo    = campaign.SharePrefix + "foo"
	prefixedShareBar    = campaign.SharePrefix + "bar"
	prefixedShareFoobar = campaign.SharePrefix + "foobar"
)

func TestEmptyShares(t *testing.T) {
	shares := campaign.EmptyShares()
	require.Equal(t, shares, campaign.Shares(nil))
}

func TestNewShares(t *testing.T) {
	_, err := campaign.NewShares("invalid")
	require.Error(t, err)

	shares, err := campaign.NewShares("100foo,200bar")
	require.NoError(t, err)
	require.Equal(t, shares, campaign.Shares(sdk.NewCoins(
		sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
		sdk.NewCoin(prefixedShareBar, sdk.NewInt(200)),
	)))
}

func TestNewSharesFromCoins(t *testing.T) {
	shares := campaign.NewSharesFromCoins(sdk.NewCoins(
		sdk.NewCoin("foo", sdk.NewInt(100)),
		sdk.NewCoin("bar", sdk.NewInt(200)),
	))
	require.Equal(t, shares, campaign.Shares(sdk.NewCoins(
		sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
		sdk.NewCoin(prefixedShareBar, sdk.NewInt(200)),
	)))
}

func TestCheckShares(t *testing.T) {
	require.NoError(t, campaign.CheckShares(campaign.Shares(sdk.NewCoins(
		sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
		sdk.NewCoin(prefixedShareBar, sdk.NewInt(200)),
	))))
	require.Error(t, campaign.CheckShares(campaign.Shares(sdk.NewCoins(
		sdk.NewCoin("foo", sdk.NewInt(100)),
		sdk.NewCoin(prefixedShareBar, sdk.NewInt(200)),
	))))
}

func TestIncreaseShares(t *testing.T) {
	for _, tc := range []struct {
		desc      string
		shares    campaign.Shares
		newShares campaign.Shares
		expected  campaign.Shares
	}{
		{
			desc:   "increase empty set",
			shares: campaign.EmptyShares(),
			newShares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
			expected: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
		},
		{
			desc: "no new shares",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
			newShares: campaign.EmptyShares(),
			expected: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
		},
		{
			desc: "increase shares",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
			newShares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(50)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(50)),
				sdk.NewCoin(prefixedShareFoobar, sdk.NewInt(50)),
			)),
			expected: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(150)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(150)),
				sdk.NewCoin(prefixedShareFoobar, sdk.NewInt(50)),
			)),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.Equal(t, tc.expected, campaign.IncreaseShares(tc.shares, tc.newShares))
		})
	}
}

func TestDecreaseShares(t *testing.T) {
	for _, tc := range []struct {
		desc       string
		shares     campaign.Shares
		toDecrease campaign.Shares
		expected   campaign.Shares
		isError    bool
	}{
		{
			desc:   "decrease empty set",
			shares: campaign.EmptyShares(),
			toDecrease: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
			isError: true,
		},
		{
			desc: "decrease from empty set",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
			toDecrease: campaign.EmptyShares(),
			expected: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
		},
		{
			desc: "decrease to negative",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(50)),
			)),
			toDecrease: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
			isError: true,
		},
		{
			desc: "decrease normal set",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareFoobar, sdk.NewInt(50)),
			)),
			toDecrease: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(30)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
			expected: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(70)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(0)),
				sdk.NewCoin(prefixedShareFoobar, sdk.NewInt(50)),
			)),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			decreased, err := campaign.DecreaseShares(tc.shares, tc.toDecrease)
			require.Equal(t, tc.isError, err != nil)
			if !tc.isError {
				require.Equal(t, tc.expected, decreased)
			}
		})
	}
}

func TestShareIsTotalReached(t *testing.T) {
	for _, tc := range []struct {
		desc        string
		shares      campaign.Shares
		totalShares campaign.Shares
		reached     bool
	}{
		{
			desc:        "empty is false",
			shares:      campaign.EmptyShares(),
			totalShares: campaign.EmptyShares(),
			reached:     false,
		},
		{
			desc: "no default total is reached",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(campaign.DefaultTotalShareNumber)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
			totalShares: campaign.EmptyShares(),
			reached:     false,
		},
		{
			desc: "no custom total is reached",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(50)),
			)),
			totalShares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
				sdk.NewCoin(prefixedShareFoobar, sdk.NewInt(100)),
			)),
			reached: false,
		},
		{
			desc: "a default total is reached",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(campaign.DefaultTotalShareNumber+1)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
			totalShares: campaign.EmptyShares(),
			reached:     true,
		},
		{
			desc: "a custom total is reached",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareFoo, sdk.NewInt(campaign.DefaultTotalShareNumber)),
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(101)),
			)),
			totalShares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin(prefixedShareBar, sdk.NewInt(100)),
			)),
			reached: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.Equal(t, tc.reached, campaign.IsTotalSharesReached(tc.shares, tc.totalShares))
		})
	}
}

func TestIsEqualShares(t *testing.T) {
	type args struct {
		share1 campaign.Shares
		share2 campaign.Shares
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal shares",
			args: args{
				share1: campaign.NewSharesFromCoins(sdk.NewCoins(
					sdk.NewCoin(prefixedShareFoo, sdk.NewInt(campaign.DefaultTotalShareNumber)),
					sdk.NewCoin(prefixedShareBar, sdk.NewInt(101)),
				)),
				share2: campaign.NewSharesFromCoins(sdk.NewCoins(
					sdk.NewCoin(prefixedShareFoo, sdk.NewInt(campaign.DefaultTotalShareNumber)),
					sdk.NewCoin(prefixedShareBar, sdk.NewInt(101)),
				)),
			},
			want: true,
		},
		{
			name: "not equal values",
			args: args{
				share1: campaign.NewSharesFromCoins(sdk.NewCoins(
					sdk.NewCoin(prefixedShareFoo, sdk.NewInt(campaign.DefaultTotalShareNumber)),
					sdk.NewCoin(prefixedShareBar, sdk.NewInt(10)),
				)),
				share2: campaign.NewSharesFromCoins(sdk.NewCoins(
					sdk.NewCoin(prefixedShareFoo, sdk.NewInt(campaign.DefaultTotalShareNumber)),
					sdk.NewCoin(prefixedShareBar, sdk.NewInt(101)),
				)),
			},
			want: false,
		},
		{
			name: "invalid coin number",
			args: args{
				share1: campaign.NewSharesFromCoins(sdk.NewCoins(
					sdk.NewCoin(prefixedShareFoo, sdk.NewInt(campaign.DefaultTotalShareNumber)),
					sdk.NewCoin(prefixedShareBar, sdk.NewInt(10)),
				)),
				share2: campaign.NewSharesFromCoins(sdk.NewCoins(
					sdk.NewCoin(prefixedShareFoo, sdk.NewInt(campaign.DefaultTotalShareNumber)),
				)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := campaign.IsEqualShares(tt.args.share1, tt.args.share2)
			require.True(t, got == tt.want)
		})
	}
}
