package types_test

import (
	"fmt"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
	project "github.com/tendermint/spn/x/project/types"
)

var (
	prefixedShareFoo    = project.SharePrefix + "foo"
	prefixedShareBar    = project.SharePrefix + "bar"
	prefixedShareFoobar = project.SharePrefix + "foobar"
)

func TestEmptyShares(t *testing.T) {
	t.Run("should allow creation of empty shares", func(t *testing.T) {
		shares := project.EmptyShares()
		require.Equal(t, project.Shares(nil), shares)
	})
}

func TestNewShares(t *testing.T) {
	t.Run("should prevent creation of invalid shares", func(t *testing.T) {
		_, err := project.NewShares("invalid")
		require.Error(t, err)
	})

	t.Run("should allow creation of valid shares", func(t *testing.T) {
		shares, err := project.NewShares("100foo,200bar")
		require.NoError(t, err)
		require.Equal(t, shares, tc.Shares(t, "100foo,200bar"))
	})
}

func TestNewSharesFromCoins(t *testing.T) {
	t.Run("should allow creation of valid shares from coins", func(t *testing.T) {
		shares := project.NewSharesFromCoins(sdk.NewCoins(
			sdk.NewCoin("foo", sdkmath.NewInt(100)),
			sdk.NewCoin("bar", sdkmath.NewInt(200)),
		))
		require.Equal(t, shares, tc.Shares(t, "100foo,200bar"))
	})
}

func TestCheckShares(t *testing.T) {
	t.Run("should allow check of valid shares", func(t *testing.T) {
		require.NoError(t, project.CheckShares(tc.Shares(t, "100foo,200bar")))
	})

	t.Run("should prevent check of invalid shares", func(t *testing.T) {
		require.Error(t, project.CheckShares(project.Shares(sdk.NewCoins(
			sdk.NewCoin("foo", sdkmath.NewInt(100)),
			sdk.NewCoin("s/bar", sdkmath.NewInt(200)),
		))))
	})
}

func TestIncreaseShares(t *testing.T) {
	for _, tc := range []struct {
		name      string
		shares    project.Shares
		newShares project.Shares
		expected  project.Shares
	}{
		{
			name:      "should increase shares",
			shares:    tc.Shares(t, "100foo,100bar"),
			newShares: tc.Shares(t, "50foo,50bar,50foobar"),
			expected:  tc.Shares(t, "150foo,150bar,50foobar"),
		},
		{
			name:      "should perform nothing with two empty sets",
			shares:    project.EmptyShares(),
			newShares: project.EmptyShares(),
			expected:  project.EmptyShares(),
		},
		{
			name:      "should increase empty set",
			shares:    project.EmptyShares(),
			newShares: tc.Shares(t, "100foo,200bar"),
			expected:  tc.Shares(t, "100foo,200bar"),
		},
		{
			name:      "should create no new shares",
			shares:    tc.Shares(t, "100foo,100bar"),
			newShares: project.EmptyShares(),
			expected:  tc.Shares(t, "100foo,100bar"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, project.IncreaseShares(tc.shares, tc.newShares))
		})
	}
}

func TestDecreaseShares(t *testing.T) {
	for _, tc := range []struct {
		desc       string
		shares     project.Shares
		toDecrease project.Shares
		expected   project.Shares
		isError    bool
	}{
		{
			desc:       "should decrease empty set",
			shares:     project.EmptyShares(),
			toDecrease: tc.Shares(t, "100foo,100bar"),
			isError:    true,
		},
		{
			desc:       "should decrease from empty set",
			shares:     tc.Shares(t, "100foo,100bar"),
			toDecrease: project.EmptyShares(),
			expected:   tc.Shares(t, "100foo,100bar"),
		},
		{
			desc:       "should decrease to negative shares",
			shares:     tc.Shares(t, "100foo,50bar"),
			toDecrease: tc.Shares(t, "100foo,100bar"),
			isError:    true,
		},
		{
			desc:       "should decrease normal set",
			shares:     tc.Shares(t, "100foo,100bar,50foobar"),
			toDecrease: tc.Shares(t, "30foo,100bar"),
			expected:   tc.Shares(t, "70foo,50foobar"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			decreased, err := project.DecreaseShares(tc.shares, tc.toDecrease)
			require.Equal(t, tc.isError, err != nil)
			if !tc.isError {
				require.Equal(t, tc.expected, decreased)
			}
		})
	}
}

func TestShareIsTotalReached(t *testing.T) {
	for _, tc := range []struct {
		desc           string
		shares         project.Shares
		maxTotalShares uint64
		reached        bool
		isValid        bool
	}{
		{
			desc:           "should return false with empty shares",
			shares:         project.EmptyShares(),
			maxTotalShares: 0,
			reached:        false,
			isValid:        true,
		},
		{
			desc:           "should return false when total not reached",
			shares:         tc.Shares(t, "1000foo,100bar"),
			maxTotalShares: 1000,
			reached:        false,
			isValid:        true,
		},
		{
			desc:           "should return true when total reached",
			shares:         tc.Shares(t, "1001foo,100bar"),
			maxTotalShares: 1000,
			reached:        true,
			isValid:        true,
		},
		{
			desc: "should return error if shares are invalid",
			shares: project.NewSharesFromCoins(sdk.Coins{
				sdk.NewCoin("foo", sdkmath.NewIntFromUint64(500)),
				sdk.NewCoin("foo", sdkmath.NewIntFromUint64(500)),
			}),
			maxTotalShares: 1000,
			reached:        false,
			isValid:        false,
		},
		{
			desc: "should return error if shares have invalid format",
			shares: project.Shares(sdk.Coins{
				sdk.NewCoin("foo", sdkmath.NewIntFromUint64(500)),
			}),
			maxTotalShares: 1000,
			reached:        false,
			isValid:        false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			reached, err := project.IsTotalSharesReached(tc.shares, tc.maxTotalShares)
			if tc.isValid {
				require.NoError(t, err)
				require.True(t, tc.reached == reached)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestIsEqualShares(t *testing.T) {
	type args struct {
		share1 project.Shares
		share2 project.Shares
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should be equal shares",
			args: args{
				share1: tc.Shares(t, fmt.Sprintf("%dfoo,101bar", spntypes.TotalShareNumber)),
				share2: tc.Shares(t, fmt.Sprintf("%dfoo,101bar", spntypes.TotalShareNumber)),
			},
			want: true,
		},
		{
			name: "should be not equal values",
			args: args{
				share1: tc.Shares(t, fmt.Sprintf("%dfoo,10bar", spntypes.TotalShareNumber)),
				share2: tc.Shares(t, fmt.Sprintf("%dfoo,101bar", spntypes.TotalShareNumber)),
			},
			want: false,
		},
		{
			name: "should be false with invalid coin number between sets",
			args: args{
				share1: tc.Shares(t, fmt.Sprintf("%dfoo,10bar", spntypes.TotalShareNumber)),
				share2: tc.Shares(t, fmt.Sprintf("%dfoo", spntypes.TotalShareNumber)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := project.IsEqualShares(tt.args.share1, tt.args.share2)
			require.True(t, got == tt.want)
		})
	}
}

func TestSharesAmountOf(t *testing.T) {
	tests := []struct {
		name   string
		shares project.Shares
		want   int64
	}{
		{
			name:   "should return positive amount",
			shares: tc.Shares(t, "200foo,205bar,50foobar"),
			want:   50,
		},
		{
			name:   "should return zero amount",
			shares: tc.Shares(t, "100foo,100bar,0foobar"),
			want:   0,
		},
		{
			name:   "should return zero for absent denom",
			shares: tc.Shares(t, "100foo,100bar"),
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shares.AmountOf(project.SharePrefix + "foobar")
			require.True(t, got == tt.want)
		})
	}
}

func TestSharesIsAllLTE(t *testing.T) {
	type args struct {
		share1 project.Shares
		share2 project.Shares
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return all less than",
			args: args{
				share1: tc.Shares(t, "100foo,100bar"),
				share2: tc.Shares(t, "200foo,205bar"),
			},
			want: true,
		},
		{
			name: "should return equal",
			args: args{
				share1: tc.Shares(t, "200foo,100bar"),
				share2: tc.Shares(t, "200foo,100bar"),
			},
			want: true,
		},
		{
			name: "should return true for different set less",
			args: args{
				share1: tc.Shares(t, "5foo"),
				share2: tc.Shares(t, "50foo,10bar"),
			},
			want: true,
		},
		{
			name: "should return false for not all denom less than",
			args: args{
				share1: tc.Shares(t, "200foo,100bar"),
				share2: tc.Shares(t, "100foo,105bar"),
			},
			want: false,
		},
		{
			name: "should return false for no denom less than",
			args: args{
				share1: tc.Shares(t, "200foo,200bar"),
				share2: tc.Shares(t, "100foo,105bar"),
			},
			want: false,
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

func TestSharesEmpty(t *testing.T) {
	for _, tc := range []struct {
		desc   string
		shares project.Shares
		empty  bool
	}{
		{
			desc:   "should be empty",
			shares: project.EmptyShares(),
			empty:  true,
		},
		{
			desc:   "should be empty for nil shares",
			shares: project.Shares(nil),
			empty:  true,
		},
		{
			desc:   "should be not empty",
			shares: tc.Shares(t, "100foo"),
			empty:  false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.Equal(t, tc.empty, tc.shares.Empty())
		})
	}
}

func TestSharesString(t *testing.T) {
	for _, tc := range []struct {
		desc   string
		shares project.Shares
		str    string
	}{
		{
			desc:   "should return empty string for empty shares",
			shares: project.EmptyShares(),
			str:    "",
		},
		{
			desc:   "should return one denom",
			shares: tc.Shares(t, "100foo"),
			str:    fmt.Sprintf("100%sfoo", project.SharePrefix),
		},
		{
			desc:   "should return list of denoms",
			shares: tc.Shares(t, "100foo,100bar"),
			str:    fmt.Sprintf("100%sbar,100%sfoo", project.SharePrefix, project.SharePrefix),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.Equal(t, tc.str, tc.shares.String())
		})
	}
}

func TestShares_CoinsFromTotalSupply(t *testing.T) {
	tests := []struct {
		name             string
		shares           project.Shares
		totalSupply      sdk.Coins
		totalShareNumber uint64
		expected         sdk.Coins
		wantErr          bool
	}{
		{
			name:             "should return empty for empty total supply",
			shares:           sample.Shares(r),
			totalSupply:      sdk.NewCoins(),
			totalShareNumber: 10000,
			expected:         sdk.NewCoins(),
			wantErr:          false,
		},
		{
			name:             "should return empty for empty shares",
			shares:           project.EmptyShares(),
			totalSupply:      sample.Coins(r),
			totalShareNumber: 10000,
			expected:         sdk.NewCoins(),
			wantErr:          false,
		},
		{
			name:             "should return total supply if all share 100%",
			shares:           tc.Shares(t, "100foo,100bar,100baz"),
			totalSupply:      tc.Coins(t, "1000foo,500bar,200baz"),
			totalShareNumber: 100,
			expected:         tc.Coins(t, "1000foo,500bar,200baz"),
			wantErr:          false,
		},
		{
			name:             "should omit coins with no share",
			shares:           tc.Shares(t, "100foo,100baz"),
			totalSupply:      tc.Coins(t, "1000foo,500bar,200baz"),
			totalShareNumber: 100,
			expected:         tc.Coins(t, "1000foo,200baz"),
			wantErr:          false,
		},
		{
			name:             "should omit coins with with share but no total supply",
			shares:           tc.Shares(t, "100foo,100bar,100baz"),
			totalSupply:      tc.Coins(t, "1000foo,200baz"),
			totalShareNumber: 100,
			expected:         tc.Coins(t, "1000foo,200baz"),
			wantErr:          false,
		},
		{
			name:             "should return coins total supply relative to the share",
			shares:           tc.Shares(t, "5000foo,3000bar,8000baz"),
			totalSupply:      tc.Coins(t, "100000foo,100000bar,100000baz"),
			totalShareNumber: 10000,
			expected:         tc.Coins(t, "50000foo,30000bar,80000baz"),
			wantErr:          false,
		},
		{
			name:             "should return cut decimal from coins",
			shares:           tc.Shares(t, "5000foo"),
			totalSupply:      tc.Coins(t, "11foo"),
			totalShareNumber: 10000,
			expected:         tc.Coins(t, "5foo"),
			wantErr:          false,
		},
		{
			name:             "should return no share if less than 1",
			shares:           tc.Shares(t, "9999foo"),
			totalSupply:      tc.Coins(t, "1foo"),
			totalShareNumber: 10000,
			expected:         sdk.NewCoins(),
			wantErr:          false,
		},
		{
			name:             "should prevent using total share number 0",
			shares:           sample.Shares(r),
			totalSupply:      sample.Coins(r),
			totalShareNumber: 0,
			expected:         sdk.NewCoins(),
			wantErr:          true,
		},
		{
			name:             "should prevent shares with amount greater than total share number",
			shares:           tc.Shares(t, "5foo,100bar,5baz"),
			totalSupply:      sample.Coins(r),
			totalShareNumber: 10,
			expected:         sdk.NewCoins(),
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coins, err := tt.shares.CoinsFromTotalSupply(tt.totalSupply, tt.totalShareNumber)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.True(t, coins.IsEqual(tt.expected),
					"%s should be %s",
					coins.String(),
					tt.expected.String(),
				)
			}
		})
	}
}
