package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateDebugMode(t *testing.T) {
	require.Error(t, validateDebugMode(1))
	require.NoError(t, validateDebugMode(false))
}
