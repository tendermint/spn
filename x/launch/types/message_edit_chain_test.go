package types_test

import (
	codec "github.com/cosmos/cosmos-sdk/codec/types"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgEditChain_ValidateBasic(t *testing.T) {
	var err error
	chainID, _ := sample.ChainID(0)

	msgInvalidInitialGenesis := sample.MsgEditChain(
		sample.AccAddress(),
		chainID,
		true,
		false,
		false,
	)
	msgInvalidInitialGenesis.InitialGenesis, err = codec.NewAnyWithValue(&types.Chain{})
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		desc  string
		msg   types.MsgEditChain
		valid bool
	}{
		{
			desc: "valid message with new source and genesis",
			msg: sample.MsgEditChain(
				sample.AccAddress(),
				chainID,
				true,
				true,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new source",
			msg: sample.MsgEditChain(
				sample.AccAddress(),
				chainID,
				true,
				false,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new genesis",
			msg: sample.MsgEditChain(
				sample.AccAddress(),
				chainID,
				false,
				true,
				false,
			),
			valid: true,
		},
		{
			desc: "valid message with new genesis with a custom genesis url",
			msg: sample.MsgEditChain(
				sample.AccAddress(),
				chainID,
				false,
				true,
				true,
			),
			valid: true,
		},
		{
			desc: "invalid coordinator address",
			msg: sample.MsgEditChain(
				"invalid",
				chainID,
				true,
				true,
				false,
			),
			valid: false,
		},
		{
			desc: "invalid chain id",
			msg: sample.MsgEditChain(
				sample.AccAddress(),
				"invalid",
				true,
				true,
				false,
			),
			valid: false,
		},
		{
			desc: "no value to edit",
			msg: sample.MsgEditChain(
				sample.AccAddress(),
				chainID,
				false,
				false,
				false,
			),
			valid: false,
		},
		{
			desc: "invalid initial genesis",
			msg: msgInvalidInitialGenesis,
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
