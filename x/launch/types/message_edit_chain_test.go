package types_test

import (
	"testing"
)

func TestMsgEditChain_ValidateBasic(t *testing.T) {
	//validMsg := sample.MsgCreateChain(sample.AccAddress(), "valid", "")
	//invalidAddress := sample.MsgCreateChain("invalid", "valid", "")
	//invalidName := sample.MsgCreateChain(sample.AccAddress(), "invalid-name", "")
	//
	//for _, tc := range []struct {
	//	desc  string
	//	msg   types.MsgCreateChain
	//	valid bool
	//}{
	//	{
	//		desc:  "valid message",
	//		msg:   validMsg,
	//		valid: true,
	//	},
	//	{
	//		desc:  "invalid address",
	//		msg:   invalidAddress,
	//		valid: false,
	//	},
	//	{
	//		desc:  "invalid chain name",
	//		msg:   invalidName,
	//		valid: false,
	//	},
	//} {
	//	tc := tc
	//	t.Run(tc.desc, func(t *testing.T) {
	//		err := tc.msg.ValidateBasic()
	//		if tc.valid {
	//			require.NoError(t, err)
	//		} else {
	//			require.Error(t, err)
	//		}
	//	})
	//}
}