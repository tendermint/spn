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
			name: "invalid campaign name",
			msg: types.MsgUpdateCampaignName{
				CampaignID:  0,
				Coordinator: sample.Address(),
				Name:        invalidCampaignName,
			},
			err: types.ErrInvalidCampaignName,
		},
		{
			name: "invalid coordinator address",
			msg: types.MsgUpdateCampaignName{
				CampaignID:  0,
				Coordinator: "invalid_address",
				Name:        "newName",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid coordinator message",
			msg: types.MsgUpdateCampaignName{
				CampaignID:  0,
				Coordinator: sample.Address(),
				Name:        "newName",
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
