package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgCreateChain_ValidateBasic(t *testing.T) {
	invalidGenesisHash := sample.MsgCreateChain(sample.AccAddress(), "foo.com")
	invalidGenesisHash.GenesisHash = "NoHash"

	for _, tc := range []struct {
		desc  string
		msg   types.MsgCreateChain
		valid bool
	}{
		{
			desc:  "valid message",
			msg:   sample.MsgCreateChain(sample.AccAddress(), ""),
			valid: true,
		},
		{
			desc:  "valid message with genesis URL",
			msg:   sample.MsgCreateChain(sample.AccAddress(), "foo.com"),
			valid: true,
		},
		{
			desc:  "invalid address",
			msg:   sample.MsgCreateChain("invalid", ""),
			valid: false,
		},
		{
			desc:  "invalid genesis hash for custom genesis",
			msg:   invalidGenesisHash,
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
