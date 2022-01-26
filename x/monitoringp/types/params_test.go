package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
)

func TestParamsValidate(t *testing.T) {
	tests := []struct {
		name    string
		params  Params
		wantErr bool
	}{
		{
			name:   "default is valid",
			params: DefaultParams(),
		},
		{
			name: "valid consumer consensus state",
			params: Params{
				ConsumerChainID:        sample.GenesisChainID(),
				ConsumerConsensusState: sample.ConsensusState(0),
			},
		},
		{
			name: "invalid consumer consensus state",
			params: Params{
				ConsumerChainID: sample.GenesisChainID(),
				ConsumerConsensusState: spntypes.NewConsensusState(
					"foo",
					"DD388ED4B9DED48DEDF7C4A781AB656DD5C56D50655A662A92B516B33EA97EA2",
					"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
				),
			},
			wantErr: true,
		},
		{
			name: "invalid consumer chain ID",
			params: Params{
				ConsumerChainID:        "foo",
				ConsumerConsensusState: sample.ConsensusState(0),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}

	t.Run("validateConsumerConsensusState expect a ConsensusState pointer", func(t *testing.T) {
		require.Error(t, validateConsumerConsensusState(100))
	})

	t.Run("validateConsumerChainID expect a string", func(t *testing.T) {
		require.Error(t, validateConsumerConsensusState(100))
	})
}
