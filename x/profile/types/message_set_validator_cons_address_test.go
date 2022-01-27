package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgSetValidatorConsAddress_ValidateBasic(t *testing.T) {
	var (
		valKey = `{
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
	)
	tests := []struct {
		name string
		msg  types.MsgSetValidatorConsAddress
		err  error
	}{
		{
			name: "invalid validator key",
			msg: types.MsgSetValidatorConsAddress{
				ValidatorKey: []byte("invalid_key"),
			},
			err: types.ErrInvalidValidatorKey,
		},
		{
			name: "valid message",
			msg: types.MsgSetValidatorConsAddress{
				ValidatorKey: []byte(valKey),
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
