package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgEditChain_ValidateBasic(t *testing.T) {
	launchID := uint64(0)

	msgInvalidMetadataLen := sample.MsgEditChain(r,
		sample.Address(r),
		launchID,
		false,
		0,
		false,
	)
	msgInvalidMetadataLen.Metadata = sample.Bytes(r, spntypes.MaxMetadataLength+1)

	for _, tc := range []struct {
		desc  string
		msg   types.MsgEditChain
		valid bool
	}{
		{
			desc: "valid message",
			msg: sample.MsgEditChain(r,
				sample.Address(r),
				launchID,
				true,
				0,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new metadata",
			msg: sample.MsgEditChain(r,
				sample.Address(r),
				launchID,
				false,
				0,
				true,
			),
			valid: true,
		},
		{
			desc: "valid message with new chain ID",
			msg: sample.MsgEditChain(r,
				sample.Address(r),
				launchID,
				true,
				0,
				false,
			),
			valid: true,
		},
		{
			desc: "invalid coordinator address",
			msg: sample.MsgEditChain(r,
				"invalid",
				launchID,
				true,
				0,
				false,
			),
			valid: false,
		},
		{
			desc: "no value to edit",
			msg: sample.MsgEditChain(r,
				sample.Address(r),
				launchID,
				false,
				0,
				false,
			),
			valid: false,
		},
		{
			desc:  "invalid metadata length",
			msg:   msgInvalidMetadataLen,
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
