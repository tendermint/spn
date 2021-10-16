package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgUpdateCampaignName_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdateCampaignName
		err  error
	}{
		{
			name: "invalid name",
			msg: types.MsgUpdateCampaignName{
				Creator: sample.Address(),
				Name:    "",
			},
			err: types.ErrInvalidCampaignName,
		},
		{
			name: "invalid address",
			msg: types.MsgUpdateCampaignName{
				Creator: "invalid_address",
				Name:    "new_name",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address",
			msg: types.MsgUpdateCampaignName{
				Creator: sample.Address(),
				Name:    "new_name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
