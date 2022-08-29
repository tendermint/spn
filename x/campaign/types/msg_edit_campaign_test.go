package types_test

import (
	"testing"

	spntypes "github.com/tendermint/spn/pkg/types"

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
				Coordinator: sample.Address(r),
				Name:        invalidCampaignName,
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrInvalidCampaignName,
		},
		{
			name: "invalid coordinator address",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: "invalid_address",
				Name:        sample.CampaignName(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: sdkerrortypes.ErrInvalidAddress,
		},
		{
			name: "valid message - both modified",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(r),
				Name:        sample.CampaignName(r),
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "valid message - name modified",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(r),
				Name:        sample.CampaignName(r),
				Metadata:    []byte{},
			},
		},
		{
			name: "valid message - metadata modified",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(r),
				Name:        "",
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "invalid metadata length",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(r),
				Name:        sample.CampaignName(r),
				Metadata:    sample.Metadata(r, spntypes.MaxMetadataLength+1),
			},
			err: types.ErrInvalidMetadataLength,
		},
		{
			name: "no fields modified",
			msg: types.MsgEditCampaign{
				CampaignID:  0,
				Coordinator: sample.Address(r),
				Name:        "",
				Metadata:    []byte{},
			},
			err: sdkerrortypes.ErrInvalidRequest,
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
