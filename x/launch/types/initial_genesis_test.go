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

func TestGenesisURLHash(t *testing.T) {
	require.EqualValues(t,
		"2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae",
		types.GenesisURLHash("foo"),
	)
}
