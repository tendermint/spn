package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddValidator_ValidateBasic(t *testing.T) {
	launchID := uint64(0)

	addr := sample.Address(r)
	validMsg := sample.MsgRequestAddValidator(r, sample.Address(r), sample.Address(r), launchID)
	emptyConsPubKey := validMsg
	emptyConsPubKey.ConsPubKey = []byte{}
	emptyGentx := validMsg
	emptyGentx.GenTx = []byte{}
	emptyPeer := validMsg
	emptyPeer.Peer = types.Peer{}
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
			name:  "same creator and validator",
			msg:   sample.MsgRequestAddValidator(r, addr, addr, launchID),
			valid: true,
		},
		{
			name:  "invalid creator address",
			msg:   sample.MsgRequestAddValidator(r, "invalid", sample.Address(r), launchID),
			valid: false,
		},
		{
			name:  "invalid validator address",
			msg:   sample.MsgRequestAddValidator(r, sample.Address(r), "invalid", launchID),
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
