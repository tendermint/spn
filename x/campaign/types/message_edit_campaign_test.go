package types_test

import (
	spntypes "github.com/tendermint/spn/pkg/types"
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
				Name:        sample.CampaignName(),
				Metadata:    sample.Metadata(20),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid message - both modified",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(),
				Name:        sample.CampaignName(),
				Metadata:    sample.Metadata(20),
			},
		},
		{
			name: "valid message - name modified",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(),
				Name:        sample.CampaignName(),
				Metadata:    []byte{},
			},
		},
		{
			name: "valid message - metadata modified",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(),
				Name:        "",
				Metadata:    sample.Metadata(20),
			},
		},
		{
			name: "invalid metadata length",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(),
				Name:        sample.CampaignName(),
				Metadata:    sample.Metadata(spntypes.MaxMetadataLength + 1),
			},
			err: types.ErrInvalidMetadataLength,
		},
		{
			name: "no fields modified",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(),
				Name:        "",
				Metadata:    []byte{},
			},
			err: sdkerrors.ErrInvalidRequest,
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
