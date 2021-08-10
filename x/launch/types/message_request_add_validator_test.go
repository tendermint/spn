package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddValidator_ValidateBasic(t *testing.T) {
	chainID, _ := sample.ChainID(0)

	validMsg := sample.MsgRequestAddValidator(sample.AccAddress(), chainID)
	emptyConsPubKey := validMsg
	emptyConsPubKey.ConsPubKey = []byte{}
	emptyGentx := validMsg
	emptyGentx.GenTx = []byte{}
	emptyPeer := validMsg
	emptyPeer.Peer = ""
	invalidSelfDelegation := validMsg
	invalidSelfDelegation.SelfDelegation.Denom = ""
	zeroDelegation := validMsg
	zeroDelegation.SelfDelegation.Amount = sdk.NewInt(0)

	for _, tc := range []struct {
		name  string
		msg   types.MsgRequestAddValidator
		valid bool
	}{
		{
			name:  "valid message",
			msg:   validMsg,
			valid: true,
		},
		{
			name:  "invalid address",
			msg:   sample.MsgRequestAddValidator("invalid", chainID),
			valid: false,
		},
		{
			name:  "invalid chain ID",
			msg:   sample.MsgRequestAddValidator(sample.AccAddress(), "invalid"),
			valid: false,
		},
		{
			name:  "empty consensus public key",
			msg:   emptyConsPubKey,
			valid: false,
		},
		{
			name:  "empty gentx",
			msg:   emptyGentx,
			valid: false,
		},
		{
			name:  "empty peer",
			msg:   emptyPeer,
			valid: false,
		},
		{
			name:  "invalid self delegation",
			msg:   invalidSelfDelegation,
			valid: false,
		},
		{
			name:  "zero self delegation",
			msg:   zeroDelegation,
			valid: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}

}
