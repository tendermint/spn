package errors_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	spnerrors "github.com/tendermint/spn/pkg/errors"
)

func TestCritical(t *testing.T) {
	require.ErrorIs(t, spnerrors.ErrCritical, spnerrors.Critical("foo"))
}

func TestCriticalf(t *testing.T) {
	require.ErrorIs(t, spnerrors.ErrCritical, spnerrors.Criticalf("foo %v", "bar"))
}