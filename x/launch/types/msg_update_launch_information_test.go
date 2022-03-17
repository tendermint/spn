package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgUpdateLaunchInformation_ValidateBasic(t *testing.T) {
	launchID := uint64(0)

	msgInvalidGenesisHash := sample.MsgUpdateLaunchInformation(
		sample.Address(),
		launchID,
		false,
		true,
		false,
		false,
	)
	genesisURL := types.NewGenesisURL("foo.com", "NoHash")
	msgInvalidGenesisHash.InitialGenesis = &genesisURL

	msgInvalidGenesisChainID := sample.MsgUpdateLaunchInformation(
		sample.Address(),
		launchID,
		false,
		true,
		false,
		false,
	)
	msgInvalidGenesisChainID.GenesisChainID = "invalid"

	for _, tc := range []struct {
		desc  string
		msg   types.MsgUpdateLaunchInformation
		valid bool
	}{
		{
			desc: "valid message",
			msg: sample.MsgUpdateLaunchInformation(
				sample.Address(),
				launchID,
				true,
				true,
				true,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new genesis chain ID",
			msg: sample.MsgUpdateLaunchInformation(
				sample.Address(),
				launchID,
				true,
				false,
				false,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new source",
			msg: sample.MsgUpdateLaunchInformation(
				sample.Address(),
				launchID,
				false,
				true,
				false,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new genesis",
			msg: sample.MsgUpdateLaunchInformation(
				sample.Address(),
				launchID,
				false,
				false,
				true,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new genesis with a custom genesis url",
			msg: sample.MsgUpdateLaunchInformation(
				sample.Address(),
				launchID,
				false,
				false,
				true,
				true,
			),
			valid: true,
		},
		{
			desc: "invalid coordinator address",
			msg: sample.MsgUpdateLaunchInformation(
				"invalid",
				launchID,
				false,
				true,
				true,
				false,
			),
			valid: false,
		},
		{
			desc: "no value to edit",
			msg: sample.MsgUpdateLaunchInformation(
				sample.Address(),
				launchID,
				false,
				false,
				false,
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
