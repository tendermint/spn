package ibctypes_test

import (
	"encoding/base64"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/pkg/ibctypes"
)

func TestConsensusStateFile_RootHash(t *testing.T) {
	csf := ibctypes.ConsensusState{
		NextValHash: "foo",
		Root: ibctypes.MerkeRool{
			Hash: "bar",
		},
		Timestamp: "foobar",
	}
	require.EqualValues(t, "bar", csf.RootHash())
}

func TestNewConsensusState(t *testing.T) {
	tests := []struct {
		name           string
		timestamp      string
		nextValSetHash string
		rootHash       string
		wantErr        bool
	}{
		{
			name:           "returns a new consensus state",
			timestamp:      "2022-01-12T07:56:35.394367Z",
			nextValSetHash: "DD388ED4B9DED48DEDF7C4A781AB656DD5C56D50655A662A92B516B33EA97EA2",
			rootHash:       "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
		},
		{
			name:           "invalid timestamp",
			timestamp:      "foo",
			nextValSetHash: "DD388ED4B9DED48DEDF7C4A781AB656DD5C56D50655A662A92B516B33EA97EA2",
			rootHash:       "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
			wantErr:        true,
		},
		{
			name:           "invalid next validator set hash",
			timestamp:      "2022-01-12T07:56:35.394367Z",
			nextValSetHash: "foo",
			rootHash:       "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
			wantErr:        true,
		},
		{
			name:           "invalid root hash",
			timestamp:      "2022-01-12T07:56:35.394367Z",
			nextValSetHash: "DD388ED4B9DED48DEDF7C4A781AB656DD5C56D50655A662A92B516B33EA97EA2",
			rootHash:       "foo",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ibctypes.NewConsensusState(tt.timestamp, tt.nextValSetHash, tt.rootHash)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.EqualValues(t, tt.timestamp, got.Timestamp.Format(time.RFC3339Nano))
			require.EqualValues(t, tt.nextValSetHash, got.NextValidatorsHash.String())
			require.EqualValues(t, tt.rootHash, base64.StdEncoding.EncodeToString(got.Root.Hash))
		})
	}
}

func TestParseConsensusStateFile(t *testing.T) {
	t.Run("parse a dumped consensus state", func(t *testing.T) {
		consensusStateYAML := `next_validators_hash: DD388ED4B9DED48DEDF7C4A781AB656DD5C56D50655A662A92B516B33EA97EA2
root:
  hash: 47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=
timestamp: "2022-01-12T07:56:35.394367Z"
`
		f, err := os.CreateTemp("", "spn_consensus_state_test")
		require.NoError(t, err)
		t.Cleanup(func() {
			f.Close()
			os.Remove(f.Name())
		})
		_, err = f.WriteString(consensusStateYAML)
		require.NoError(t, err)

		csf, err := ibctypes.ParseConsensusStateFile(f.Name())
		require.NoError(t, err)
		require.EqualValues(t, "2022-01-12T07:56:35.394367Z", csf.Timestamp)
		require.EqualValues(t, "DD388ED4B9DED48DEDF7C4A781AB656DD5C56D50655A662A92B516B33EA97EA2", csf.NextValHash)
		require.EqualValues(t, "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=", csf.RootHash())

	})

	t.Run("non-existent file", func(t *testing.T) {
		_, err := ibctypes.ParseConsensusStateFile("/foo/bar/foobar")
		require.Error(t, err)
	})

	t.Run("invalid file", func(t *testing.T) {
		consensusStateYAML := `foo`
		f, err := os.CreateTemp("", "spn_consensus_state_test")
		require.NoError(t, err)
		t.Cleanup(func() {
			f.Close()
			os.Remove(f.Name())
		})
		_, err = f.WriteString(consensusStateYAML)
		require.NoError(t, err)

		_, err = ibctypes.ParseConsensusStateFile(f.Name())
		require.Error(t, err)
	})
}
