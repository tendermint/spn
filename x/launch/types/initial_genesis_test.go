package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestDefaultInitialGenesis_Validate(t *testing.T) {
	require.NoError(t, types.DefaultInitialGenesis{}.Validate(),
		" DefaultInitialGenesis should always be valid",
	)
}

func TestGenesisURL_Validate(t *testing.T) {
	require.NoError(t, types.GenesisURL{
		Url:  sample.String(30),
		Hash: sample.GenesisHash(),
	}.Validate(),
		" GenesisURL should be valid",
	)

	require.Error(t, types.GenesisURL{
		Url:  "",
		Hash: sample.GenesisHash(),
	}.Validate(),
		" GenesisURL must contain a url",
	)

	require.Error(t, types.GenesisURL{
		Url:  sample.String(30),
		Hash: sample.String(types.HashLength - 1),
	}.Validate(),
		" GenesisURL must contain a valid sha256 hash",
	)
}
