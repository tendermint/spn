package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestValidator_AddValidatorOperatorAddress(t *testing.T) {
	var (
		operatorAddress   = sample.Address()
		operatorAddresses = []string{
			sample.Address(),
			sample.Address(),
			sample.Address(),
		}
	)
	tests := []struct {
		name               string
		operatorAddress   string
		operatorAddresses []string
		want               []string
	}{
		{
			name:               "valid case",
			operatorAddress:   operatorAddress,
			operatorAddresses: operatorAddresses,
			want:               append(operatorAddresses, operatorAddress),
		},
		{
			name:               "empty operator address list",
			operatorAddress:   operatorAddress,
			operatorAddresses: []string{},
			want:               []string{operatorAddress},
		},
		{
			name:               "already existing operator address case",
			operatorAddress:   operatorAddress,
			operatorAddresses: append(operatorAddresses, operatorAddress),
			want:               append(operatorAddresses, operatorAddress),
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

func TestValidator_RemoveValidatorOperatorAddress(t *testing.T) {
	var (
		operatorAddress   = sample.Address()
		operatorAddresses = []string{
			sample.Address(),
			sample.Address(),
			sample.Address(),
		}
	)
	tests := []struct {
		name               string
		operatorAddress   string
		operatorAddresses []string
		want               []string
	}{
		{
			name:               "valid case",
			operatorAddress:   operatorAddress,
			operatorAddresses: append(operatorAddresses, operatorAddress),
			want:               operatorAddresses,
		},
		{
			name:               "empty operator address list",
			operatorAddress:   operatorAddress,
			operatorAddresses: []string{operatorAddress},
			want:               []string{},
		},
		{
			name:               "non existing operator address case",
			operatorAddress:   operatorAddress,
			operatorAddresses: operatorAddresses,
			want:               operatorAddresses,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := types.Validator{OperatorAddresses: tt.operatorAddresses}
			got := validator.RemoveValidatorOperatorAddress(tt.operatorAddress)
			require.Equal(t, tt.want, got.OperatorAddresses)
		})
	}
}

func TestValidator_HasOperatorAddress(t *testing.T) {
	var (
		operatorAddress   = sample.Address()
		operatorAddresses = []string{
			sample.Address(),
			sample.Address(),
			sample.Address(),
		}
	)
	tests := []struct {
		name               string
		operatorAddress   string
		operatorAddresses []string
		want               bool
	}{
		{
			name:               "hasn't the operator address",
			operatorAddress:   operatorAddress,
			operatorAddresses: operatorAddresses,
			want:               false,
		},
		{
			name:               "only the operator address",
			operatorAddress:   operatorAddress,
			operatorAddresses: []string{operatorAddress},
			want:               true,
		},
		{
			name:               "empty operator address list",
			operatorAddress:   operatorAddress,
			operatorAddresses: []string{},
			want:               false,
		},
		{
			name:               "nil operator address list",
			operatorAddress:   operatorAddress,
			operatorAddresses: nil,
			want:               false,
		},
		{
			name:               "has the operator address",
			operatorAddress:   operatorAddress,
			operatorAddresses: append(operatorAddresses, operatorAddress),
			want:               true,
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
