package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgRedeemVouchers_ValidateBasic(t *testing.T) {
	addr := sample.Address(r)
	tests := []struct {
		name string
		msg  types.MsgRedeemVouchers
		err  error
	}{
		{
			name: "invalid sender address",
			msg: types.MsgRedeemVouchers{
				Sender:     "invalid_address",
				Account:    sample.Address(r),
				Vouchers:   sample.Vouchers(r, 0),
				CampaignID: 0,
			},
			err: sdkerrortypes.ErrInvalidAddress,
		},
		{
			name: "invalid account address",
			msg: types.MsgRedeemVouchers{
				Sender:     sample.Address(r),
				Account:    "invalid_address",
				Vouchers:   sample.Vouchers(r, 0),
				CampaignID: 0,
			},
			err: sdkerrortypes.ErrInvalidAddress,
		},
		{
			name: "invalid coin voucher",
			msg: types.MsgRedeemVouchers{
				Sender:     sample.Address(r),
				Account:    sample.Address(r),
				Vouchers:   invalidCoins,
				CampaignID: 0,
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "vouchers don't match to campaign",
			msg: types.MsgRedeemVouchers{
				Sender:     sample.Address(r),
				Account:    sample.Address(r),
				Vouchers:   sample.Vouchers(r, 10),
				CampaignID: 0,
			},
			err: types.ErrNoMatchVouchers,
		},
		{
			name: "invalid voucher prefix",
			msg: types.MsgRedeemVouchers{
				Sender:  sample.Address(r),
				Account: sample.Address(r),
				Vouchers: sdk.NewCoins(
					sdk.NewCoin("invalid/foo", sdk.NewInt(100)),
				),
				CampaignID: 0,
			},
			err: types.ErrNoMatchVouchers,
		},
		{
			name: "empty vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:     sample.Address(r),
				Account:    sample.Address(r),
				Vouchers:   sdk.Coins{},
				CampaignID: 0,
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "valid message",
			msg: types.MsgRedeemVouchers{
				Sender:     sample.Address(r),
				Account:    sample.Address(r),
				Vouchers:   sample.Vouchers(r, 0),
				CampaignID: 0,
			},
		},
		{
			name: "valid for same account and sender",
			msg: types.MsgRedeemVouchers{
				Sender:     addr,
				Account:    addr,
				Vouchers:   sample.Vouchers(r, 0),
				CampaignID: 0,
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
