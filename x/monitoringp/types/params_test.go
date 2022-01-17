package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/pkg/ibctypes"
	"github.com/tendermint/spn/testutil/sample"
)

func TestParamsValidate(t *testing.T) {
	validConsensusState := sample.ConsensusState(0)
	invalidConsensusState := ibctypes.NewConsensusState(
		"foo",
		"DD388ED4B9DED48DEDF7C4A781AB656DD5C56D50655A662A92B516B33EA97EA2",
		"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	)

	tests := []struct {
		name    string
		params  Params
		wantErr bool
	}{
		{
			name:   "default is valid",
			params: NewParams(),
		},
		{
			name: "valid consumer consensus state",
			params: Params{
				ConsumerConsensusState: &validConsensusState,
			},
		},
		{
			name: "invalid consumer consensus state",
			params: Params{
				ConsumerConsensusState: &invalidConsensusState,
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
		require.Error(t, validateConsumerConsensusState(sample.ConsensusState(0)))
	})
}
