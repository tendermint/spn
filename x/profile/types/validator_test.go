package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestValidator_AddValidatorOperatorAddress(t *testing.T) {
	var (
		operatorAddress   = sample.Address(r)
		operatorAddresses = []string{
			sample.Address(r),
			sample.Address(r),
			sample.Address(r),
		}
	)
	tests := []struct {
		name              string
		operatorAddress   string
		operatorAddresses []string
		want              []string
	}{
		{
			name:              "should allow adding an operator address in an existing list",
			operatorAddress:   operatorAddress,
			operatorAddresses: operatorAddresses,
			want:              append(operatorAddresses, operatorAddress),
		},
		{
			name:              "should allow add an operator address in an empty list",
			operatorAddress:   operatorAddress,
			operatorAddresses: []string{},
			want:              []string{operatorAddress},
		},
		{
			name:              "should prevent adding duplicated operator addresses",
			operatorAddress:   operatorAddress,
			operatorAddresses: append(operatorAddresses, operatorAddress),
			want:              append(operatorAddresses, operatorAddress),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := types.Validator{OperatorAddresses: tt.operatorAddresses}
			got := validator.AddValidatorOperatorAddress(tt.operatorAddress)
			require.Equal(t, tt.want, got.OperatorAddresses)
		})
	}
}

func TestValidator_HasOperatorAddress(t *testing.T) {
	var (
		operatorAddress   = sample.Address(r)
		operatorAddresses = []string{
			sample.Address(r),
			sample.Address(r),
			sample.Address(r),
		}
	)
	tests := []struct {
		name              string
		operatorAddress   string
		operatorAddresses []string
		want              bool
	}{
		{
			name:              "should return false if address not found",
			operatorAddress:   operatorAddress,
			operatorAddresses: operatorAddresses,
			want:              false,
		},
		{
			name:              "should return true if only this address present",
			operatorAddress:   operatorAddress,
			operatorAddresses: []string{operatorAddress},
			want:              true,
		},
		{
			name:              "should return false for empty list",
			operatorAddress:   operatorAddress,
			operatorAddresses: []string{},
			want:              false,
		},
		{
			name:              "should return false for nil",
			operatorAddress:   operatorAddress,
			operatorAddresses: nil,
			want:              false,
		},
		{
			name:              "should return true if address found in a list",
			operatorAddress:   operatorAddress,
			operatorAddresses: append(operatorAddresses, operatorAddress),
			want:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := types.Validator{OperatorAddresses: tt.operatorAddresses}
			got := validator.HasOperatorAddress(tt.operatorAddress)
			require.Equal(t, tt.want, got)
		})
	}
}
