package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
)

func TestMsgBurnVouchers_ValidateBasic(t *testing.T) {
	invalidCoins := sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	tests := []struct {
		name string
		msg  types.MsgBurnVouchers
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgBurnVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Vouchers:  sample.Vouchers(r, 0),
			},
		},
		{
			name: "should prevent validation of msg with invalid address",
			msg: types.MsgBurnVouchers{
				Sender:    "invalid_address",
				ProjectID: 0,
				Vouchers:  sample.Coins(r),
			},
			err: types.ErrInvalidVoucherAddress,
		},
		{
			name: "should prevent validation of msg with invalid vouchers",
			msg: types.MsgBurnVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Vouchers:  invalidCoins,
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "should prevent validation of msg with empty vouchers",
			msg: types.MsgBurnVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Vouchers:  sdk.Coins{},
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "should prevent validation of msg with vouchers not matching project",
			msg: types.MsgBurnVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Vouchers: sdk.NewCoins(
					sdk.NewCoin("invalid/foo", sdkmath.NewInt(100)),
				),
			},
			err: types.ErrNoMatchVouchers,
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
