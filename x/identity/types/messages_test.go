package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/identity/types"

	"testing"
)

func TestMsgSetUsername(t *testing.T) {
	// Can create a message with a valid username
	addr, _ := sdk.AccAddressFromBech32(spnmocks.MockAccAddress())

	_, err := types.NewMsgSetUsername(addr, "foo-bar_40foo_01")
	require.NoError(t, err)
	_, err = types.NewMsgSetUsername(addr, spnmocks.MockRandomString(types.UsernameMaxLength))
	require.NoError(t, err)

	// Prevent to create message with invalid name
	_, err = types.NewMsgSetUsername(addr, "")
	require.Error(t, err)
	_, err = types.NewMsgSetUsername(addr, "foo!")
	require.Error(t, err)
	_, err = types.NewMsgSetUsername(addr, spnmocks.MockRandomString(types.UsernameMaxLength+1))
	require.Error(t, err)

}
