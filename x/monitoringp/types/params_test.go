package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
)

var (
	chainID        = "foo-1"
	consensusState = spntypes.NewConsensusState(
		"2022-01-12T12:25:19.523109Z",
		"48C4C20AC5A7BD99A45AEBAB92E61F5667253A2C51CCCD84D20327D3CB8737C9",
		"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	)
	invalidConsensusState = spntypes.NewConsensusState(
		"foo",
		"DD388ED4B9DED48DEDF7C4A781AB656DD5C56D50655A662A92B516B33EA97EA2",
		"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	)
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
				LastBlockHeight:         1000,
				ConsumerChainID:         chainID,
				ConsumerConsensusState:  consensusState,
				ConsumerUnbondingPeriod: spntypes.DefaultUnbondingPeriod,
				ConsumerRevisionHeight:  spntypes.DefaultRevisionHeight,
			},
		},
		{
			name: "should prevent invalid last block height",
			params: Params{
				LastBlockHeight:         0,
				ConsumerChainID:         chainID,
				ConsumerConsensusState:  consensusState,
				ConsumerUnbondingPeriod: spntypes.DefaultUnbondingPeriod,
				ConsumerRevisionHeight:  spntypes.DefaultRevisionHeight,
			},
			wantErr: true,
		},
		{
			name: "should prevent invalid consumer chain ID",
			params: Params{
				LastBlockHeight:         1000,
				ConsumerChainID:         "foo",
				ConsumerConsensusState:  consensusState,
				ConsumerUnbondingPeriod: spntypes.DefaultUnbondingPeriod,
				ConsumerRevisionHeight:  spntypes.DefaultRevisionHeight,
			},
			wantErr: true,
		},
		{
			name: "should prevent invalid consumer consensus state",
			params: Params{
				LastBlockHeight:         1000,
				ConsumerChainID:         chainID,
				ConsumerConsensusState:  invalidConsensusState,
				ConsumerUnbondingPeriod: spntypes.DefaultUnbondingPeriod,
				ConsumerRevisionHeight:  spntypes.DefaultRevisionHeight,
			},
			wantErr: true,
		},
		{
			name: "should prevent invalid consumer unbonding period",
			params: Params{
				LastBlockHeight:         1000,
				ConsumerChainID:         chainID,
				ConsumerConsensusState:  consensusState,
				ConsumerUnbondingPeriod: spntypes.MinimalUnbondingPeriod - 1,
				ConsumerRevisionHeight:  spntypes.DefaultRevisionHeight,
			},
			wantErr: true,
		},
		{
			name: "should prevent invalid consumer revision height",
			params: Params{
				LastBlockHeight:         1000,
				ConsumerChainID:         chainID,
				ConsumerConsensusState:  consensusState,
				ConsumerUnbondingPeriod: spntypes.DefaultUnbondingPeriod,
				ConsumerRevisionHeight:  0,
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

	t.Run("validateConsumerChainID expect a string", func(t *testing.T) {
		require.Error(t, validateConsumerConsensusState(100))
	})
}

func TestValidateLastBlockHeight(t *testing.T) {
	require.Error(t, validateLastBlockHeight("foo"), "should expect a int64")
	require.Error(t, validateLastBlockHeight(int64(0)), "prevent using 0")
	require.NoError(t, validateLastBlockHeight(int64(1)))
}

func TestValidateConsumerConsensusState(t *testing.T) {
	require.Error(t, validateConsumerConsensusState(100), "should expect a ConsensusState")
	require.Error(t, validateConsumerConsensusState(invalidConsensusState), "should prevent invalid ConsensusState")
	require.NoError(t, validateConsumerConsensusState(spntypes.ConsensusState{}), "empty is valid")
	require.NoError(t, validateConsumerConsensusState(consensusState))
}

func TestValidateConsumerChainID(t *testing.T) {
	require.Error(t, validateConsumerChainID(100), "should expect a string")
	require.Error(t, validateConsumerChainID("invalid-id"), "should prevent invalid chain ID")
	require.NoError(t, validateConsumerChainID(chainID))
}

func TestValidateConsumerUnbondingPeriod(t *testing.T) {
	require.Error(t, validateConsumerUnbondingPeriod("foo"), "should expect a int64")
	require.Error(t, validateConsumerUnbondingPeriod(int64(spntypes.MinimalUnbondingPeriod-1)), "should prevent below minimal value")
	require.NoError(t, validateConsumerUnbondingPeriod(int64(spntypes.MinimalUnbondingPeriod)))
	require.NoError(t, validateConsumerUnbondingPeriod(int64(spntypes.DefaultUnbondingPeriod)))
}

func TestValidateConsumerRevisionHeight(t *testing.T) {
	require.Error(t, validateConsumerRevisionHeight("foo"), "should expect a uint64")
	require.Error(t, validateConsumerRevisionHeight(uint64(0)), "should prevent using 0")
	require.NoError(t, validateConsumerRevisionHeight(uint64(1)))
}

func TestValidateDebugMode(t *testing.T) {
	require.Error(t, validateDebugMode(1))
	require.NoError(t, validateDebugMode(false))
}
