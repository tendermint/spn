package types

import (
	"errors"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		params Params
		err    error
	}{
		{
			name:   "invalid min total supply",
			params: NewParams(sdk.ZeroInt(), DefaultMaxTotalSupply),
			err:    errors.New("minimum total supply should be greater than one"),
		},
		{
			name:   "min total supply greater than max",
			params: NewParams(DefaultMaxTotalSupply, DefaultMinTotalSupply),
			err:    errors.New("maximum total supply should be greater or equal than minimum total supply"),
		},
		{
			name:   "valid range",
			params: NewParams(DefaultMinTotalSupply, DefaultMaxTotalSupply),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.ValidateBasic()
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateTotalSupplyRange(t *testing.T) {
	tests := []struct {
		name        string
		supplyRange interface{}
		err         error
	}{
		{
			name:        "invalid interface",
			supplyRange: "test",
			err:         fmt.Errorf("invalid parameter type: string"),
		},
		{
			name: "invalid minTotalSupply",
			supplyRange: TotalSupplyRange{
				MinTotalSupply: sdk.ZeroInt(),
				MaxTotalSupply: DefaultMaxTotalSupply,
			},
			err: errors.New("minimum total supply should be greater than one"),
		},
		{
			name: "invalid range",
			supplyRange: TotalSupplyRange{
				MinTotalSupply: DefaultMaxTotalSupply,
				MaxTotalSupply: DefaultMinTotalSupply,
			},
			err: errors.New("maximum total supply should be greater or equal than minimum total supply"),
		},
		{
			name: "valid range",
			supplyRange: TotalSupplyRange{
				MinTotalSupply: DefaultMinTotalSupply,
				MaxTotalSupply: DefaultMinTotalSupply.Add(sdk.OneInt()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTotalSupplyRange(tt.supplyRange)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
