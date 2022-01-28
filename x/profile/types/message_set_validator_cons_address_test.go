package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	valtypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

const validatorKey = `{
  "address": "B4AAC35ED4E14C09E530B10AF4DD604FAAC597C0",
  "pub_key": {
    "type": "tendermint/PubKeyEd25519",
    "value": "sYTsd7W1+SBtjD3BN/aTEDFvfRbZ9zdfpQH2Lk3MRK4="
  },
  "priv_key": {
    "type": "tendermint/PrivKeyEd25519",
    "value": "j45JhnCflEk3T6FC8LLuJqg9tPfCzJH+UYZY88xn+0exhOx3tbX5IG2MPcE39pMQMW99Ftn3N1+lAfYuTcxErg=="
  }
}`

func TestMsgSetValidatorConsAddress_ValidateBasic(t *testing.T) {
	valKey, err := valtypes.LoadValidatorKey([]byte(validatorKey))
	require.NoError(t, err)
	signature, err := valKey.Sign(0, "spn-1")
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  types.MsgSetValidatorConsAddress
		err  error
	}{
		{
			name: "invalid validator address",
			msg: types.MsgSetValidatorConsAddress{
				ValidatorAddress:    "invalid_address",
				ValidatorConsPubKey: valKey.PubKey.Bytes(),
				ValidatorKeyType:    valKey.PubKey.Type(),
				Signature:           signature,
				Nonce:               0,
				ChainID:             "spn-1",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid validator consensus key",
			msg: types.MsgSetValidatorConsAddress{
				ValidatorAddress:    sample.Address(),
				ValidatorConsPubKey: sample.Bytes(10),
				ValidatorKeyType:    "invalid_key_type",
				Signature:           signature,
				Nonce:               0,
				ChainID:             "spn-1",
			},
			err: types.ErrInvalidValidatorKey,
		},
		{
			name: "invalid signature",
			msg: types.MsgSetValidatorConsAddress{
				ValidatorAddress:    sample.Address(),
				ValidatorConsPubKey: valKey.PubKey.Bytes(),
				ValidatorKeyType:    valKey.PubKey.Type(),
				Signature:           signature,
				Nonce:               99,
				ChainID:             "invalid_chain_id",
			},
			err: types.ErrInvalidValidatorSignature,
		},
		{
			name: "valid message",
			msg: types.MsgSetValidatorConsAddress{
				ValidatorAddress:    sample.Address(),
				ValidatorConsPubKey: valKey.PubKey.Bytes(),
				ValidatorKeyType:    valKey.PubKey.Type(),
				Signature:           signature,
				Nonce:               0,
				ChainID:             "spn-1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
