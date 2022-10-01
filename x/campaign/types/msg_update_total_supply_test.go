package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgUpdateTotalSupply_ValidateBasic(t *testing.T) {
	invalidCoins := sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	tests := []struct {
		name string
		msg  types.MsgUpdateTotalSupply
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgUpdateTotalSupply{
				Coordinator:       sample.Address(r),
				CampaignID:        0,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
		},
		{
			name: "should prevent validation of msg with invalid address",
			msg: types.MsgUpdateTotalSupply{
				Coordinator:       "invalid_address",
				CampaignID:        0,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: sdkerrortypes.ErrInvalidAddress,
		},
		{
			name: "should prevent validation of msg with invalid total supply",
			msg: types.MsgUpdateTotalSupply{
				Coordinator:       sample.Address(r),
				CampaignID:        0,
				TotalSupplyUpdate: invalidCoins,
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "should prevent validation of msg with empty total supply",
			msg: types.MsgUpdateTotalSupply{
				Coordinator:       sample.Address(r),
				CampaignID:        0,
				TotalSupplyUpdate: sdk.NewCoins(),
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
