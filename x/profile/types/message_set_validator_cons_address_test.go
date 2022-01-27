package types_test

import (
	"github.com/tendermint/tendermint/crypto/ed25519"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgSetValidatorConsAddress_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgSetValidatorConsAddress
		err  error
	}{
		{
			name: "invalid validator address",
			msg: types.MsgSetValidatorConsAddress{
				ValidatorAddress:    "invalid_address",
				ValidatorConsPubKey: sample.Bytes(10),
				ValidatorKeyType:    ed25519.KeyType,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid validator consensus key",
			msg: types.MsgSetValidatorConsAddress{
				ValidatorAddress:    sample.Address(),
				ValidatorConsPubKey: sample.Bytes(10),
				ValidatorKeyType:    "invalid_key_type",
			},
			err: types.ErrInvalidValidatorKey,
		},
		{
			name: "valid message",
			msg: types.MsgSetValidatorConsAddress{
				ValidatorAddress:    sample.Address(),
				ValidatorConsPubKey: sample.Bytes(10),
				ValidatorKeyType:    ed25519.KeyType,
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
