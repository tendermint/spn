package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRevertLaunch_ValidateBasic(t *testing.T) {
	addr := sample.AccAddress()
	chainID := uint64(0)

	for _, tc := range []struct {
		desc  string
		msg   types.MsgRevertLaunch
		valid bool
	}{
		{
			desc:  "valid message",
			msg:   *types.NewMsgRevertLaunch(addr, chainID),
			valid: true,
		},
		{
			desc:  "invalid coordinator address",
			msg:   *types.NewMsgRevertLaunch("invalid", chainID),
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
