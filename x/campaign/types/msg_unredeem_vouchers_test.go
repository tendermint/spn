package types_test

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgUnredeemVouchers_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUnredeemVouchers
		err  error
	}{
		{
			name: "valid message",
			msg: types.MsgUnredeemVouchers{
				Sender:     sample.Address(r),
				CampaignID: 0,
				Shares:     sample.Shares(r),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgUnredeemVouchers{
				Sender:     "invalid_address",
				CampaignID: 0,
				Shares:     sample.Shares(r),
			},
			err: sdkerrortypes.ErrInvalidAddress,
		},
		{
			name: "invalid shares",
			msg: types.MsgUnredeemVouchers{
				Sender:     sample.Address(r),
				CampaignID: 0,
				Shares:     invalidShares,
			},
			err: types.ErrInvalidShares,
		},
		{
			name: "empty shares",
			msg: types.MsgUnredeemVouchers{
				Sender:     sample.Address(r),
				CampaignID: 0,
				Shares:     types.EmptyShares(),
			},
			err: types.ErrInvalidShares,
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
