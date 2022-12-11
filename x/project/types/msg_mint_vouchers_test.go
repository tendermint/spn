package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgMintVouchers_ValidateBasic(t *testing.T) {
	invalidShares := types.Shares{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	tests := []struct {
		name string
		msg  types.MsgMintVouchers
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				ProjectID:  0,
				Shares:      sample.Shares(r),
			},
		},
		{
			name: "should prevent validation of msg with invalid address",
			msg: types.MsgMintVouchers{
				Coordinator: "invalid_address",
				ProjectID:  0,
				Shares:      sample.Shares(r),
			},
			err: profile.ErrInvalidCoordAddress,
		},
		{
			name: "should prevent validation of msg with invalid shares",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				ProjectID:  0,
				Shares:      invalidShares,
			},
			err: types.ErrInvalidShares,
		},
		{
			name: "should prevent validation of msg with empty shares",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				ProjectID:  0,
				Shares:      types.EmptyShares(),
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
