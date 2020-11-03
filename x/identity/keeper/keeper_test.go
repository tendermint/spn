package keeper_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"testing"
)

func TestSetUsername(t *testing.T) {
	ctx, k := spnmocks.MockIdentityContext()
	address := spnmocks.MockAccAddress()

	// The username should be the address if it is not set
	username, _ := k.GetUsernameFromAddress(ctx, address)
	require.Equal(t, address.String(), username, "GetUsernameFromAddress should return the address if no username")

	// Prevent setting an invalid username
	err := k.SetUsername(ctx, address, "foo!")
	require.Error(t, err, "SetUsername should prevent using an invalid username")

	// Can set a username
	err = k.SetUsername(ctx, address, "foo")
	require.NoError(t, err, "SetUsername allows to set a valid username")

	// Username can be retrieve
	username, _ = k.GetUsernameFromAddress(ctx, address)
	require.Equal(t, "foo", username, "GetUsernameFromAddress should return the new username")

	// Username can be retrieved from the identifier
	id, _ := k.GetIdentifier(ctx, address)
	username, _ = k.GetUsername(ctx, id)
	require.Equal(t, "foo", username, "GetUsername should return the new username")

	// Can set a new username
	err = k.SetUsername(ctx, address, "bar")
	require.NoError(t, err, "SetUsername allows to set a valid username")
	username, _ = k.GetUsernameFromAddress(ctx, address)
	require.Equal(t, "bar", username, "GetUsername should return the new username")
}

func TestGetIdentifier(t *testing.T) {
	ctx, k := spnmocks.MockIdentityContext()
	address := spnmocks.MockAccAddress()

	// Return the address
	identifier, _ := k.GetIdentifier(ctx, address)
	require.Equal(t, address.String(), identifier, "GetIdentifier should return the address")
}

func TestGetAddresses(t *testing.T) {
	ctx, k := spnmocks.MockIdentityContext()
	address := spnmocks.MockAccAddress()

	// Return only the address provided
	addresses, _ := k.GetAddresses(ctx, address.String())
	require.Equal(t, 1, len(addresses), "GetAddresses shoudl only return the address provided")
	require.True(t, address.Equals(addresses[0]), "GetAddresses shoudl only return the address provided")
}

func TestIdentityExists(t *testing.T) {
	ctx, k := spnmocks.MockIdentityContext()

	// Return true if the identifier is an address
	address := spnmocks.MockAccAddress()
	exists, _ := k.IdentityExists(ctx, address.String())
	require.True(t, exists, "Any Bech32 address should be a valid identifier")

	// Return false if not an address
	exists, _ = k.IdentityExists(ctx, "foo")
	require.False(t, exists, "A non Bech32 address should not be a valid identifier")
}
