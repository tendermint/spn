package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParams_Validate(t *testing.T) {
	tests := []struct {
		name    string
		params  Params
		wantErr bool
	}{
		{
			name: "should prevent validate params with invalid decay information",
			params: NewParams(DecayInformation{
				Enabled:    true,
				DecayStart: time.UnixMilli(1001),
				DecayEnd:   time.UnixMilli(1000),
			}),
			wantErr: true,
		},
		{
			name:   "should validate valid params",
			params: DefaultParams(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateDecayInformation(t *testing.T) {
	tests := []struct {
		name             string
		decayInformation interface{}
		wantErr          bool
	}{
		{
			name:             "should prevent validate decay information with invalid interface",
			decayInformation: "test",
			wantErr:          true,
		},
		{
			name: "should prevent validate invalid decay information",
			decayInformation: DecayInformation{
				Enabled:    true,
				DecayStart: time.UnixMilli(1001),
				DecayEnd:   time.UnixMilli(1000),
			},
			wantErr: true,
		},
		{
			name: "should validate valid decay information",
			decayInformation: DecayInformation{
				Enabled: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDecayInformation(tt.decayInformation)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
