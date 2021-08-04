package types_test

import (
	"testing"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestVestedAccountCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	request := sample.Request("foo")
	content := &types.VestedAccount{
		Address: sample.AccAddress(),
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackVestedAccount(cdc)
	require.NoError(t, err)
	require.EqualValues(t, content, result)
	invalidContent := &types.Request{}
	request.Content, err = codec.NewAnyWithValue(invalidContent)
	require.NoError(t, err)
	_, err = request.UnpackVestedAccount(cdc)
	require.Error(t, err)
	require.Equal(t, "not a vestedAccount request", err.Error())
}
