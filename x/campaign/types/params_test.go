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
			params: NewParams(sdk.ZeroInt(), DefaultMaxTotalSupply, DefaultCampaignCreationFee),
			err:    errors.New("minimum total supply should be greater than one: invalid total supply range"),
		},
		{
			name:   "min total supply greater than max",
			params: NewParams(DefaultMaxTotalSupply, DefaultMinTotalSupply, DefaultCampaignCreationFee),
			err:    errors.New("maximum total supply should be greater or equal than minimum total supply: invalid total supply range"),
		},
		{
			name:   "invalid coins for campaign creation fee",
			params: NewParams(DefaultMinTotalSupply, DefaultMaxTotalSupply, sdk.Coins{sdk.Coin{Denom: "foo", Amount: sdk.NewInt(-1)}}),
			err:    errors.New("coin -1foo amount is not positive"),
		},
		{
			name:   "valid params",
			params: NewParams(DefaultMinTotalSupply, DefaultMaxTotalSupply, DefaultCampaignCreationFee),
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
			name: "invalid range",
			supplyRange: TotalSupplyRange{
				MinTotalSupply: DefaultMaxTotalSupply,
				MaxTotalSupply: DefaultMinTotalSupply,
			},
			err: errors.New("maximum total supply should be greater or equal than minimum total supply: invalid total supply range"),
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
