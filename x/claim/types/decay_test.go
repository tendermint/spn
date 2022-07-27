package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
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
		{},
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
