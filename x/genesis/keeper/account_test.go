package keeper_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"testing"
)

func TestSetAccount(t *testing.T) {
	ctx, k := spnmocks.MockGenesisContext()
	chainID := spnmocks.MockRandomAlphaString(5)
	address := spnmocks.MockAccAddress()
	payload := spnmocks.MockProposalAddAccountPayload()

	// IsAccountSet returns false for not set account
	isSet := k.IsAccountSet(ctx, chainID, address)
	require.False(t, isSet)
	_, found := k.GetAccountCoins(ctx, chainID, address)
	require.False(t, found)

	// SetAccount set the account
	k.SetAccount(ctx, chainID, address, payload)
	isSet = k.IsAccountSet(ctx, chainID, address)
	require.True(t, isSet)
	coins, found := k.GetAccountCoins(ctx, chainID, address)
	require.True(t, found)
	require.True(t, payload.Coins.IsEqual(coins))

	// The account is not set for a different chain
	isSet = k.IsAccountSet(ctx, spnmocks.MockRandomAlphaString(6), address)
	require.False(t, isSet)

	// RemoveAccount removes the account
	k.RemoveAccount(ctx, chainID, address)
	isSet = k.IsAccountSet(ctx, chainID, address)
	require.False(t, isSet)
}
