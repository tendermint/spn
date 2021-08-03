package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgEditChain_ValidateBasic(t *testing.T) {
	addr := sample.AccAddress()
	chainID, _ := sample.ChainID(0)

	for _, tc := range []struct {
		desc  string
		msg   types.MsgTriggerLaunch
		valid bool
	}{
		{
			desc:  "valid message",
			msg:   *types.NewMsgTriggerLaunch(addr, chainID, uint64(1000)),
			valid: true,
		},
		{
			desc:  "invalid coordinator address",
			msg:   *types.NewMsgTriggerLaunch("invalid", chainID, uint64(1000)),
			valid: false,
		},
		{
			desc:  "invalid chain id",
			msg:   *types.NewMsgTriggerLaunch(addr, "invalid", uint64(1000)),
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
