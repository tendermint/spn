package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgCreateCampaign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateCampaign
		err  error
	}{
		{
			name: "valid message",
			msg: types.MsgCreateCampaign{
				Coordinator:  sample.Address(),
				CampaignName: sample.CampaignName(),
				TotalSupply:  sample.Coins(),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgCreateCampaign{
				Coordinator:  "invalid_address",
				CampaignName: sample.CampaignName(),
				TotalSupply:  sample.Coins(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid campaign name",
			msg: types.MsgCreateCampaign{
				Coordinator:  sample.Address(),
				CampaignName: invalidCampaignName,
				TotalSupply:  sample.Coins(),
			},
			err: types.ErrInvalidCampaignName,
		},
		{
			name: "invalid total supply",
			msg: types.MsgCreateCampaign{
				Coordinator:  sample.Address(),
				CampaignName: sample.CampaignName(),
				TotalSupply:  invalidCoins,
			},
			err: types.ErrInvalidTotalSupply,
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
