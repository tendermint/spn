package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
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
	zeroDelegation.SelfDelegation.Amount = sdkmath.ZeroInt()

	for _, tc := range []struct {
		name  string
		msg   types.MsgRequestAddValidator
		valid bool
	}{
		{
			name:  "should validate valid message",
			msg:   validMsg,
			valid: true,
		},
		{
			name:  "should validate valid message with same creator and validator",
			msg:   sample.MsgRequestAddValidator(r, addr, addr, launchID),
			valid: true,
		},
		{
			name:  "should prevent validate message with invalid creator address",
			msg:   sample.MsgRequestAddValidator(r, "invalid", sample.Address(r), launchID),
			valid: false,
		},
		{
			name:  "should prevent validate message with invalid validator address",
			msg:   sample.MsgRequestAddValidator(r, sample.Address(r), "invalid", launchID),
			valid: false,
		},
		{
			name:  "should prevent validate message with empty consensus public key",
			msg:   emptyConsPubKey,
			valid: false,
		},
		{
			name:  "should prevent validate message with empty gentx",
			msg:   emptyGentx,
			valid: false,
		},
		{
			name:  "should prevent validate message with empty peer",
			msg:   emptyPeer,
			valid: false,
		},
		{
			name:  "should prevent validate message with invalid self delegation",
			msg:   invalidSelfDelegation,
			valid: false,
		},
		{
			name:  "should prevent validate message with zero self delegation",
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
