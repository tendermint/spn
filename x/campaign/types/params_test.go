package types

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	sdkmath "cosmossdk.io/math"
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
			name: "should allow validate valid params",
			params: NewParams(
				DefaultMinTotalSupply,
				DefaultMaxTotalSupply,
				DefaultCampaignCreationFee,
				DefaultMaxMetadataLength,
			),
		},
		{
			name: "should prevent validate invalid min total supply",
			params: NewParams(
				sdkmath.ZeroInt(),
				DefaultMaxTotalSupply,
				DefaultCampaignCreationFee,
				DefaultMaxMetadataLength,
			),
			err: errors.New("minimum total supply should be greater than one: invalid total supply range"),
		},
		{
			name: "should prevent validate min total supply greater than max",
			params: NewParams(
				DefaultMaxTotalSupply,
				DefaultMinTotalSupply,
				DefaultCampaignCreationFee,
				DefaultMaxMetadataLength,
			),
			err: errors.New("maximum total supply should be greater or equal than minimum total supply: invalid total supply range"),
		},
		{
			name: "should prevent validate invalid coins for campaign creation fee",
			params: NewParams(
				DefaultMinTotalSupply,
				DefaultMaxTotalSupply,
				sdk.Coins{sdk.Coin{Denom: "foo", Amount: sdkmath.NewInt(-1)}},
				DefaultMaxMetadataLength,
			),
			err: errors.New("coin -1foo amount is not positive"),
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
				MaxTotalSupply: DefaultMinTotalSupply.Add(sdkmath.OneInt()),
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

func TestValidateCampaignCreationFee(t *testing.T) {
	tests := []struct {
		name        string
		creationFee interface{}
		err         error
	}{
		{
			name:        "invalid interface",
			creationFee: "test",
			err:         fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:        "invalid coin",
			creationFee: sdk.Coins{sdk.Coin{Denom: "foo", Amount: sdkmath.NewInt(-1)}},
			err:         errors.New("coin -1foo amount is not positive"),
		},
		{
			name:        "valid empty param",
			creationFee: DefaultCampaignCreationFee,
		},
		{
			name: "valid param",
			creationFee: sdk.NewCoins(
				sdk.NewInt64Coin("foo", rand.Int63n(1000)+1),
				sdk.NewInt64Coin("bar", rand.Int63n(1000)+1),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCampaignCreationFee(tt.creationFee)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateMaxMetadataLength(t *testing.T) {
	tests := []struct {
		name              string
		maxMetadataLength interface{}
		err               error
	}{
		{
			name:              "invalid interface",
			maxMetadataLength: "test",
			err:               fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:              "invalid number type",
			maxMetadataLength: 1000,
			err:               fmt.Errorf("invalid parameter type: int"),
		},
		{
			name:              "valid param",
			maxMetadataLength: uint64(1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateMaxMetadataLength(tt.maxMetadataLength)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
