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

	msgInvalidGenesisHash := sample.MsgEditChain(
		sample.Address(),
		launchID,
		false,
		true,
		false,
		false,
		false,
		0,
		false,
	)
	genesisURL := types.NewGenesisURL("foo.com", "NoHash")
	msgInvalidGenesisHash.InitialGenesis = &genesisURL

	msgInvalidGenesisChainID := sample.MsgEditChain(
		sample.Address(),
		launchID,
		false,
		true,
		false,
		false,
		false,
		0,
		false,
	)
	msgInvalidGenesisChainID.GenesisChainID = "invalid"

	msgInvalidMetadataLen := sample.MsgEditChain(
		sample.Address(),
		launchID,
		false,
		false,
		false,
		false,
		false,
		0,
		false,
	)
	msgInvalidMetadataLen.Metadata = sample.Bytes(spntypes.MaxMetadataLength + 1)

	for _, tc := range []struct {
		desc  string
		msg   types.MsgEditChain
		valid bool
	}{
		{
			desc: "valid message",
			msg: sample.MsgEditChain(
				sample.Address(),
				launchID,
				true,
				true,
				true,
				false,
				false,
				0,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new genesis chain ID",
			msg: sample.MsgEditChain(
				sample.Address(),
				launchID,
				true,
				false,
				false,
				false,
				false,
				0,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new source",
			msg: sample.MsgEditChain(
				sample.Address(),
				launchID,
				false,
				true,
				false,
				false,
				false,
				0,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new genesis",
			msg: sample.MsgEditChain(
				sample.Address(),
				launchID,
				false,
				false,
				true,
				false,
				false,
				0,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new genesis with a custom genesis url",
			msg: sample.MsgEditChain(
				sample.Address(),
				launchID,
				false,
				false,
				true,
				true,
				false,
				0,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new metadata",
			msg: sample.MsgEditChain(
				sample.Address(),
				launchID,
				false,
				false,
				false,
				false,
				false,
				0,
				true,
			),
			valid: true,
		},
		{
			desc: "invalid coordinator address",
			msg: sample.MsgEditChain(
				"invalid",
				launchID,
				false,
				true,
				true,
				false,
				false,
				0,
				false,
			),
			valid: false,
		},
		{
			desc: "no value to edit",
			msg: sample.MsgEditChain(
				sample.Address(),
				launchID,
				false,
				false,
				false,
				false,
				false,
				0,
				false,
			),
			valid: false,
		},
		{
			desc:  "invalid initial genesis hash",
			msg:   msgInvalidGenesisHash,
			valid: false,
		},
		{
			desc:  "invalid initial genesis chain ID",
			msg:   msgInvalidGenesisChainID,
			valid: false,
		},
		{
			desc:  "invalid metadata length",
			msg:   msgInvalidMetadataLen,
			valid: false,
		},
	} {
		tc := tc
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
