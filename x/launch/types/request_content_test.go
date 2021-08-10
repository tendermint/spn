package types_test

import (
	"testing"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestAccountRemovalCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	request := sample.Request("foo")
	content := &types.AccountRemoval{
		Address: sample.AccAddress(),
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackAccountRemoval(cdc)
	require.NoError(t, err)
	require.EqualValues(t, content, result)

	invalidContent := &types.Request{}
	request.Content, err = codec.NewAnyWithValue(invalidContent)
	require.NoError(t, err)
	_, err = request.UnpackAccountRemoval(cdc)
	require.Error(t, err)
	require.Equal(t, "not a accountRemoval request", err.Error())
}

func TestGenesisValidatorCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	request := sample.Request("foo")
	chainID, _ := sample.ChainID(1)
	content := &types.GenesisValidator{
		Address: sample.AccAddress(),
		ChainID: chainID,
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackGenesisValidator(cdc)

	require.NoError(t, err)
	require.EqualValues(t, content, result)

	invalidContent := &types.Request{}
	request.Content, err = codec.NewAnyWithValue(invalidContent)
	require.NoError(t, err)

	_, err = request.UnpackGenesisValidator(cdc)
	require.Error(t, err)
	require.Equal(t, "not a genesisValidator request", err.Error())
}

func TestValidatorRemovalCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	request := sample.Request("foo")
	content := &types.ValidatorRemoval{
		ValAddress: sample.AccAddress(),
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackValidatorRemoval(cdc)

	require.NoError(t, err)
	require.EqualValues(t, content, result)

	invalidContent := &types.Request{}
	request.Content, err = codec.NewAnyWithValue(invalidContent)
	require.NoError(t, err)

	_, err = request.UnpackValidatorRemoval(cdc)
	require.Error(t, err)
	require.Equal(t, err.Error(), "not a validatorRemoval request")

	_, err = request.UnpackGenesisValidator(cdc)
	require.Error(t, err)
	require.Equal(t, "not a genesisValidator request", err.Error())
	_, err = request.UnpackValidatorRemoval(cdc)
	require.Error(t, err)
	require.Equal(t, err.Error(), "not a validatorRemoval request")
}

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
