package keeper_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"testing"
)

func TestSetValidator(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chainID := spnmocks.MockRandomAlphaString(5)
	_, address := spnmocks.MockValAddress()

	// IsValidatorSet returns false for not set validator
	isSet := k.IsValidatorSet(ctx, chainID, address)
	require.False(t, isSet)

	// SetValidator set the validator
	k.SetValidator(ctx, chainID, address)
	isSet = k.IsValidatorSet(ctx, chainID, address)
	require.True(t, isSet)

	// The IsValidatorSet is not set for a different chain
	isSet = k.IsValidatorSet(ctx, spnmocks.MockRandomAlphaString(6), address)
	require.False(t, isSet)

	// RemoveValidator removes the validator
	k.RemoveValidator(ctx, chainID, address)
	isSet = k.IsValidatorSet(ctx, chainID, address)
	require.False(t, isSet)
}
