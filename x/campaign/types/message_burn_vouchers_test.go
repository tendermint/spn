package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgBurnVouchers_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgBurnVouchers
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgBurnVouchers{
				Sender:     "invalid_address",
				CampaignID: 0,
				Vouchers:   sample.Coins(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid message",
			msg: types.MsgBurnVouchers{
				Sender:     sample.AccAddress(),
				CampaignID: 0,
				Vouchers:   sample.Coins(),
			},
		},
		{
			name: "invalid shares",
			msg: types.MsgBurnVouchers{
				Sender:     sample.AccAddress(),
				CampaignID: 0,
				Vouchers:   invalidCoins,
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "empty shares",
			msg: types.MsgBurnVouchers{
				Sender:     sample.AccAddress(),
				CampaignID: 0,
				Vouchers:   sdk.Coins{},
			},
			err: types.ErrInvalidVouchers,
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
