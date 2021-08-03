package types_test

import (
	"fmt"
	"testing"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestGenesisAccountCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	request := sample.Request("foo")
	chainID, _ := sample.ChainID(1)
	content := &types.GenesisAccount{
		Address: sample.AccAddress(),
		ChainID: chainID,
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackGenesisAccount(cdc)
	require.NoError(t, err)
	require.EqualValues(t, content, result)

	invalidContent := &types.Request{}
	request.Content, err = codec.NewAnyWithValue(invalidContent)
	require.NoError(t, err)
	content, err = request.UnpackGenesisAccount(cdc)
	require.ErrorIs(t, err, fmt.Errorf("not a genesisAccount request"))
	require.Nil(t, content)
}
