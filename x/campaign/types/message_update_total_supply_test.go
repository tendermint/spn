package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgUpdateTotalSupply_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdateTotalSupply
		err  error
	}{
		{
			name: "valid address",
			msg: types.MsgUpdateTotalSupply{
				Coordinator: sample.Address(),
				CampaignID:  0,
				TotalSupplyUpdate: sample.Coins(),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgUpdateTotalSupply{
				Coordinator: "invalid_address",
				CampaignID:  0,
				TotalSupplyUpdate: sample.Coins(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid total supply",
			msg: types.MsgUpdateTotalSupply{
				Coordinator: sample.Address(),
				CampaignID:  0,
				TotalSupplyUpdate: invalidCoins,
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "empty total supply",
			msg: types.MsgUpdateTotalSupply{
				Coordinator: sample.Address(),
				CampaignID:  0,
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
