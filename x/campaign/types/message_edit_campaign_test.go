package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgEditCampaign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgEditCampaign
		err  error
	}{
		{
			name: "invalid campaign name",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(),
				Name:        invalidCampaignName,
				Metadata:    sample.Metadata(20),
			},
			err: types.ErrInvalidCampaignName,
		},
		{
			name: "invalid coordinator address",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: "invalid_address",
				Name:        "newName",
				Metadata:    sample.Metadata(20),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid coordinator message",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(),
				Name:        "newName",
				Metadata:    sample.Metadata(20),
			},
		},
		// TODO add test for metadata
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
