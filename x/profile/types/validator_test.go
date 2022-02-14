package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestValidator_AddValidatorConsensusAddress(t *testing.T) {
	var (
		consensusAddress   = sample.PubKey().Bytes()
		consensusAddresses = [][]byte{
			sample.PubKey().Bytes(),
			sample.PubKey().Bytes(),
			sample.PubKey().Bytes(),
		}
	)
	tests := []struct {
		name               string
		consensusAddress   []byte
		consensusAddresses [][]byte
		want               [][]byte
	}{
		{
			name:               "valid case",
			consensusAddress:   consensusAddress,
			consensusAddresses: consensusAddresses,
			want:               append(consensusAddresses, consensusAddress),
		},
		{
			name:               "empty consensus address list",
			consensusAddress:   consensusAddress,
			consensusAddresses: [][]byte{},
			want:               [][]byte{consensusAddress},
		},
		{
			name:               "already existing consensus address case",
			consensusAddress:   consensusAddress,
			consensusAddresses: append(consensusAddresses, consensusAddress),
			want:               append(consensusAddresses, consensusAddress),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := types.Validator{ConsensusAddresses: tt.consensusAddresses}
			got := validator.AddValidatorConsensusAddress(tt.consensusAddress)
			require.Equal(t, tt.want, got.ConsensusAddresses)
		})
	}
}

func TestValidator_RemoveValidatorConsensusAddress(t *testing.T) {
	var (
		consensusAddress   = sample.PubKey().Bytes()
		consensusAddresses = [][]byte{
			sample.PubKey().Bytes(),
			sample.PubKey().Bytes(),
			sample.PubKey().Bytes(),
		}
	)
	tests := []struct {
		name               string
		consensusAddress   []byte
		consensusAddresses [][]byte
		want               [][]byte
	}{
		{
			name:               "valid case",
			consensusAddress:   consensusAddress,
			consensusAddresses: append(consensusAddresses, consensusAddress),
			want:               consensusAddresses,
		},
		{
			name:               "empty consensus address list",
			consensusAddress:   consensusAddress,
			consensusAddresses: [][]byte{consensusAddress},
			want:               [][]byte{},
		},
		{
			name:               "non existing consensus address case",
			consensusAddress:   consensusAddress,
			consensusAddresses: consensusAddresses,
			want:               consensusAddresses,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := types.Validator{ConsensusAddresses: tt.consensusAddresses}
			got := validator.RemoveValidatorConsensusAddress(tt.consensusAddress)
			require.Equal(t, tt.want, got.ConsensusAddresses)
		})
	}
}

func TestValidator_HasConsensusAddress(t *testing.T) {
	var (
		consensusAddress   = sample.PubKey().Bytes()
		consensusAddresses = [][]byte{
			sample.PubKey().Bytes(),
			sample.PubKey().Bytes(),
			sample.PubKey().Bytes(),
		}
	)
	tests := []struct {
		name               string
		consensusAddress   []byte
		consensusAddresses [][]byte
		want               bool
	}{
		{
			name:               "hasn't the consensus address",
			consensusAddress:   consensusAddress,
			consensusAddresses: consensusAddresses,
			want:               false,
		},
		{
			name:               "only the consensus address",
			consensusAddress:   consensusAddress,
			consensusAddresses: [][]byte{consensusAddress},
			want:               true,
		},
		{
			name:               "empty consensus address list",
			consensusAddress:   consensusAddress,
			consensusAddresses: [][]byte{},
			want:               false,
		},
		{
			name:               "nil consensus address list",
			consensusAddress:   consensusAddress,
			consensusAddresses: nil,
			want:               false,
		},
		{
			name:               "has the consensus address",
			consensusAddress:   consensusAddress,
			consensusAddresses: append(consensusAddresses, consensusAddress),
			want:               true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := types.Validator{ConsensusAddresses: tt.consensusAddresses}
			got := validator.HasConsensusAddress(tt.consensusAddress)
			require.Equal(t, tt.want, got)
		})
	}
}
