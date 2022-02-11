package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgUpdateTotalShares_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdateTotalShares
		err  error
	}{
		{
			name: "valid address",
			msg: types.MsgUpdateTotalShares{
				Coordinator: sample.Address(),
				CampaignID:  0,
				TotalShares: sample.Shares(),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgUpdateTotalShares{
				Coordinator: "invalid_address",
				CampaignID:  0,
				TotalShares: sample.Shares(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid coins for total shares",
			msg: types.MsgUpdateTotalShares{
				Coordinator: sample.Address(),
				CampaignID:  0,
				TotalShares: types.NewSharesFromCoins(invalidCoins),
			},
			err: types.ErrInvalidShares,
		},
		{
			name: "empty total shares",
			msg: types.MsgUpdateTotalShares{
				Coordinator: sample.Address(),
				CampaignID:  0,
				TotalShares: types.EmptyShares(),
			},
			err: types.ErrInvalidShares,
		},
		{
			name: "invalid shares prefix for total shares",
			msg: types.MsgUpdateTotalShares{
				Coordinator: sample.Address(),
				CampaignID:  0,
				TotalShares: types.Shares(sdk.NewCoins(
					sdk.NewCoin("foo", sdk.NewInt(100)),
				)),
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
