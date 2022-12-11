package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
)

func TestMsgRedeemVouchers_ValidateBasic(t *testing.T) {
	invalidCoins := sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	addr := sample.Address(r)
	tests := []struct {
		name string
		msg  types.MsgRedeemVouchers
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgRedeemVouchers{
				Sender:     sample.Address(r),
				Account:    sample.Address(r),
				Vouchers:   sample.Vouchers(r, 0),
				CampaignID: 0,
			},
		},
		{
			name: "should allow validation of valid msg with same account and sender",
			msg: types.MsgRedeemVouchers{
				Sender:     addr,
				Account:    addr,
				Vouchers:   sample.Vouchers(r, 0),
				CampaignID: 0,
			},
		},
		{
			name: "should prevent validation of msg with invalid sender address",
			msg: types.MsgRedeemVouchers{
				Sender:     "invalid_address",
				Account:    sample.Address(r),
				Vouchers:   sample.Vouchers(r, 0),
				CampaignID: 0,
			},
			err: types.ErrInvalidVoucherAddress,
		},
		{
			name: "should prevent validation of msg with invalid account address",
			msg: types.MsgRedeemVouchers{
				Sender:     sample.Address(r),
				Account:    "invalid_address",
				Vouchers:   sample.Vouchers(r, 0),
				CampaignID: 0,
			},
			err: types.ErrInvalidVoucherAddress,
		},
		{
			name: "should prevent validation of msg with invalid coin voucher",
			msg: types.MsgRedeemVouchers{
				Sender:     sample.Address(r),
				Account:    sample.Address(r),
				Vouchers:   invalidCoins,
				CampaignID: 0,
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "should prevent validation of msg with vouchers not matching campaign",
			msg: types.MsgRedeemVouchers{
				Sender:     sample.Address(r),
				Account:    sample.Address(r),
				Vouchers:   sample.Vouchers(r, 10),
				CampaignID: 0,
			},
			err: types.ErrNoMatchVouchers,
		},
		{
			name: "should prevent validation of msg with invalid voucher prefix",
			msg: types.MsgRedeemVouchers{
				Sender:  sample.Address(r),
				Account: sample.Address(r),
				Vouchers: sdk.NewCoins(
					sdk.NewCoin("invalid/foo", sdkmath.NewInt(100)),
				),
				CampaignID: 0,
			},
			err: types.ErrNoMatchVouchers,
		},
		{
			name: "should prevent validation of msg with empty vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:     sample.Address(r),
				Account:    sample.Address(r),
				Vouchers:   sdk.Coins{},
				CampaignID: 0,
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
