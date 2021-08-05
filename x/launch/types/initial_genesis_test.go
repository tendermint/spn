package types

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenesisURLHash(t *testing.T) {
	require.EqualValues(t,
		"2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae",
		GenesisURLHash("foo"),
	)
}
