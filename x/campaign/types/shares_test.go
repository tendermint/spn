package types_test

import (
	"fmt"
	"testing"

	tc "github.com/tendermint/spn/testutil/constructor"

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
	require.Equal(t, campaign.Shares(nil), shares)
}

func TestNewShares(t *testing.T) {
	_, err := campaign.NewShares("invalid")
	require.Error(t, err)

	shares, err := campaign.NewShares("100foo,200bar")
	require.NoError(t, err)
	require.Equal(t, shares, tc.Shares(t, "100foo,200bar"))
}

func TestNewSharesFromCoins(t *testing.T) {
	shares := campaign.NewSharesFromCoins(sdk.NewCoins(
		sdk.NewCoin("foo", sdk.NewInt(100)),
		sdk.NewCoin("bar", sdk.NewInt(200)),
	))
	require.Equal(t, shares, tc.Shares(t, "100foo,200bar"))
}

func TestCheckShares(t *testing.T) {
	require.NoError(t, campaign.CheckShares(tc.Shares(t, "100foo,200bar")))
	require.Error(t, campaign.CheckShares(campaign.Shares(sdk.NewCoins(
		sdk.NewCoin("foo", sdk.NewInt(100)),
		sdk.NewCoin("s/bar", sdk.NewInt(200)),
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
			desc:      "increase empty set",
			shares:    campaign.EmptyShares(),
			newShares: tc.Shares(t, "100foo,200bar"),
			expected:  tc.Shares(t, "100foo,200bar"),
		},
		{
			desc:      "no new shares",
			shares:    tc.Shares(t, "100foo,100bar"),
			newShares: campaign.EmptyShares(),
			expected:  tc.Shares(t, "100foo,100bar"),
		},
		{
			desc:      "increase shares",
			shares:    tc.Shares(t, "100foo,100bar"),
			newShares: tc.Shares(t, "50foo,50bar,50foobar"),
			expected:  tc.Shares(t, "150foo,150bar,50foobar"),
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
			desc:       "decrease empty set",
			shares:     campaign.EmptyShares(),
			toDecrease: tc.Shares(t, "100foo,100bar"),
			isError:    true,
		},
		{
			desc:       "decrease from empty set",
			shares:     tc.Shares(t, "100foo,100bar"),
			toDecrease: campaign.EmptyShares(),
			expected:   tc.Shares(t, "100foo,100bar"),
		},
		{
			desc:       "decrease to negative",
			shares:     tc.Shares(t, "100foo,50bar"),
			toDecrease: tc.Shares(t, "100foo,100bar"),
			isError:    true,
		},
		{
			desc:       "decrease normal set",
			shares:     tc.Shares(t, "100foo,100bar,50foobar"),
			toDecrease: tc.Shares(t, "30foo,100bar"),
			expected:   tc.Shares(t, "70foo,50foobar"),
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
		desc    string
		shares  campaign.Shares
		reached bool
	}{
		{
			desc:    "empty is false",
			shares:  campaign.EmptyShares(),
			reached: false,
		},
		{
			desc:    "no default total is reached",
			shares:  tc.Shares(t, fmt.Sprintf("%dfoo,100bar", campaign.DefaultTotalShareNumber)),
			reached: false,
		},
		{
			desc:    "a default total is reached",
			shares:  tc.Shares(t, fmt.Sprintf("%dfoo,100bar", campaign.DefaultTotalShareNumber+1)),
			reached: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.Equal(t, tc.reached, campaign.IsTotalSharesReached(tc.shares))
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
				share1: tc.Shares(t, fmt.Sprintf("%dfoo,101bar", campaign.DefaultTotalShareNumber)),
				share2: tc.Shares(t, fmt.Sprintf("%dfoo,101bar", campaign.DefaultTotalShareNumber)),
			},
			want: true,
		},
		{
			name: "not equal values",
			args: args{
				share1: tc.Shares(t, fmt.Sprintf("%dfoo,10bar", campaign.DefaultTotalShareNumber)),
				share2: tc.Shares(t, fmt.Sprintf("%dfoo,101bar", campaign.DefaultTotalShareNumber)),
			},
			want: false,
		},
		{
			name: "invalid coin number",
			args: args{
				share1: tc.Shares(t, fmt.Sprintf("%dfoo,10bar", campaign.DefaultTotalShareNumber)),
				share2: tc.Shares(t, fmt.Sprintf("%dfoo", campaign.DefaultTotalShareNumber)),
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

func TestSharesAmountOf(t *testing.T) {
	tests := []struct {
		name   string
		shares campaign.Shares
		want   int64
	}{
		{
			name:   "present positive",
			shares: tc.Shares(t, "200foo,205bar,50foobar"),
			want:   50,
		},
		{
			name:   "present zero",
			shares: tc.Shares(t, "100foo,100bar,0foobar"),
			want:   0,
		},
		{
			name:   "absent",
			shares: tc.Shares(t, "100foo,100bar"),
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shares.AmountOf(campaign.SharePrefix + "foobar")
			require.True(t, got == tt.want)
		})
	}
}

func TestSharesIsAllLTE(t *testing.T) {
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
			name: "all less",
			args: args{
				share1: tc.Shares(t, "100foo,100bar"),
				share2: tc.Shares(t, "200foo,205bar"),
			},
			want: true,
		},
		{
			name: "not everyone less",
			args: args{
				share1: tc.Shares(t, "200foo,100bar"),
				share2: tc.Shares(t, "100foo,105bar"),
			},
			want: false,
		},
		{
			name: "no one less",
			args: args{
				share1: tc.Shares(t, "200foo,200bar"),
				share2: tc.Shares(t, "100foo,105bar"),
			},
			want: false,
		},
		{
			name: "equal",
			args: args{
				share1: tc.Shares(t, "200foo,100bar"),
				share2: tc.Shares(t, "200foo,100bar"),
			},
			want: true,
		},
		{
			name: "different set less",
			args: args{
				share1: tc.Shares(t, "5foo"),
				share2: tc.Shares(t, "50foo,10bar"),
			},
			want: true,
		},
		{
			name: "different set greater",
			args: args{
				share1: tc.Shares(t, "100foo"),
				share2: tc.Shares(t, "50foo,10bar"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.share1.IsAllLTE(tt.args.share2)
			require.True(t, got == tt.want)
		})
	}
}
