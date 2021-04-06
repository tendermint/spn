package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"testing"
)

func TestSetUsername(t *testing.T) {
	ctx, k := spnmocks.MockIdentityContext()
	address, _ := sdk.AccAddressFromBech32(spnmocks.MockAccAddress())

	// The username should be the address if it is not set
	username, _ := k.GetUsernameFromAddress(ctx, address)
	require.Equal(t, address.String(), username)

	// Prevent setting an invalid username
	err := k.SetUsername(ctx, address, "foo!")
	require.Error(t, err)

	// Can set a username
	err = k.SetUsername(ctx, address, "foo")
	require.NoError(t, err)

	// Username can be retrieve
	username, _ = k.GetUsernameFromAddress(ctx, address)
	require.Equal(t, "foo", username)

	// Username can be retrieved from the identifier
	id, _ := k.GetIdentifier(ctx, address.String())
	username, _ = k.GetUsername(ctx, id)
	require.Equal(t, "foo", username)

	// Can set a new username
	err = k.SetUsername(ctx, address, "bar")
	require.NoError(t, err)
	username, _ = k.GetUsernameFromAddress(ctx, address)
	require.Equal(t, "bar", username)
}

func TestGetIdentifier(t *testing.T) {
	ctx, k := spnmocks.MockIdentityContext()
	address, _ := sdk.AccAddressFromBech32(spnmocks.MockAccAddress())

	// Return the address
	identifier, _ := k.GetIdentifier(ctx, address.String())
	require.Equal(t, address.String(), identifier)
}

func TestGetAddresses(t *testing.T) {
	ctx, k := spnmocks.MockIdentityContext()
	address, _ := sdk.AccAddressFromBech32(spnmocks.MockAccAddress())

	// Return only the address provided
	addresses, _ := k.GetAddresses(ctx, address.String())
	require.Equal(t, 1, len(addresses))
	require.True(t, address.Equals(addresses[0]))
}

func TestIdentityExists(t *testing.T) {
	ctx, k := spnmocks.MockIdentityContext()

	// Return true if the identifier is an address
	address, _ := sdk.AccAddressFromBech32(spnmocks.MockAccAddress())
	exists, _ := k.IdentityExists(ctx, address.String())
	require.True(t, exists)

	// Return false if not an address
	exists, _ = k.IdentityExists(ctx, "foo")
	require.False(t, exists)
}
