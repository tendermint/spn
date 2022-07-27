package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/x/claim/types"
	"testing"
	"time"
)

func TestDecayInformation_Validate(t *testing.T) {
	tests := []struct {
		name      string
		decayInfo types.DecayInformation
		wantErr   bool
	}{
		{
			name: "should validate decay information with disabled",
			decayInfo: types.DecayInformation{
				Enabled:    false,
				DecayStart: time.UnixMilli(2000),
				DecayEnd:   time.UnixMilli(1000),
			},
		},
		{
			name: "should validate decay information with enabled and start equals to end",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.UnixMilli(1000),
				DecayEnd:   time.UnixMilli(1000),
			},
		},
		{
			name: "should validate decay information with enabled and end greater than start",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.UnixMilli(1000),
				DecayEnd:   time.UnixMilli(10000),
			},
		},
		{
			name: "should prevent validate decay information with enabled and start greater than end",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.UnixMilli(1001),
				DecayEnd:   time.UnixMilli(1000),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.decayInfo.Validate()

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDecayInformation_ApplyDecayFactor(t *testing.T) {
	tests := []struct {
		name          string
		decayInfo     types.DecayInformation
		coins         sdk.Coins
		currentTime   time.Time
		expectedCoins sdk.Coins
	}{
		{
			name: "should apply no change if decay disabled",
			decayInfo: types.DecayInformation{
				Enabled: false,
			},
			coins:         tc.Coins(t, "100foo,100bar"),
			expectedCoins: tc.Coins(t, "100foo,100bar"),
		},
		{
			name: "should apply no change if decay not started",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.Unix(1000, 0),
				DecayEnd:   time.Unix(10000, 0),
			},
			currentTime:   time.Unix(500, 0),
			coins:         tc.Coins(t, "100foo,100bar"),
			expectedCoins: tc.Coins(t, "100foo,100bar"),
		},
		{
			name: "should return zero coins if end of decay",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.Unix(1000, 0),
				DecayEnd:   time.Unix(10000, 0),
			},
			currentTime:   time.Unix(10000, 0),
			coins:         tc.Coins(t, "100foo,100bar"),
			expectedCoins: sdk.NewCoins(),
		},
		{
			name: "should return zero coins if end of decay without start",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.Unix(10000, 0),
				DecayEnd:   time.Unix(10000, 0),
			},
			currentTime:   time.Unix(10000, 0),
			coins:         tc.Coins(t, "100foo,100bar"),
			expectedCoins: sdk.NewCoins(),
		},
		{
			name: "should return zero coins if decay ended",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.Unix(1000, 0),
				DecayEnd:   time.Unix(10000, 0),
			},
			currentTime:   time.Unix(10001, 0),
			coins:         tc.Coins(t, "100foo,100bar"),
			expectedCoins: sdk.NewCoins(),
		},
		{
			name: "should apply half decay factor",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.Unix(10000, 0),
				DecayEnd:   time.Unix(20000, 0),
			},
			currentTime:   time.Unix(15000, 0),
			coins:         tc.Coins(t, "200000foo,2000000bar"),
			expectedCoins: tc.Coins(t, "100000foo,1000000bar"),
		},
		{
			name: "should apply 0.6 decay factor",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.Unix(10000, 0),
				DecayEnd:   time.Unix(20000, 0),
			},
			currentTime:   time.Unix(14000, 0),
			coins:         tc.Coins(t, "100000foo,1000000bar"),
			expectedCoins: tc.Coins(t, "60000foo,600000bar"),
		},
		{
			name: "should apply 0.2 decay factor",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.Unix(10000, 0),
				DecayEnd:   time.Unix(20000, 0),
			},
			currentTime:   time.Unix(18000, 0),
			coins:         tc.Coins(t, "100000foo,1000000bar"),
			expectedCoins: tc.Coins(t, "20000foo,200000bar"),
		},
		{
			name: "should apply decay factor and truncate decimals",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.Unix(10000, 0),
				DecayEnd:   time.Unix(20000, 0),
			},
			currentTime:   time.Unix(15000, 0),
			coins:         tc.Coins(t, "100000foo,1bar,1000000000003baz"),
			expectedCoins: tc.Coins(t, "50000foo,500000000001baz"),
		},
		{
			name: "should return ze coins if factor applied to zero coins",
			decayInfo: types.DecayInformation{
				Enabled:    true,
				DecayStart: time.Unix(10000, 0),
				DecayEnd:   time.Unix(20000, 0),
			},
			currentTime:   time.Unix(15000, 0),
			coins:         sdk.NewCoins(),
			expectedCoins: sdk.NewCoins(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newCoins := tt.decayInfo.ApplyDecayFactor(tt.coins, tt.currentTime)

			require.True(t, newCoins.IsEqual(tt.expectedCoins),
				"new coins are not equal to expected coins, %s != %s",
				newCoins.String(),
				tt.expectedCoins.String(),
			)
		})
	}
}
