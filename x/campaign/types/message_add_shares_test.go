package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgAddShares_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgAddShares
		err  error
	}{
		{
			name: "valid address",
			msg: types.MsgAddShares{
				Coordinator: sample.Address(),
				CampaignID:  0,
				Shares:      sample.Shares(),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgAddShares{
				Coordinator: "invalid_address",
				CampaignID:  0,
				Shares:      sample.Shares(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid shates",
			msg: types.MsgAddShares{
				Coordinator: sample.Address(),
				CampaignID:  0,
				Shares:      invalidShares,
			},
			err: types.ErrInvalidShares,
		},
		{
			name: "empty shares",
			msg: types.MsgAddShares{
				Coordinator: sample.Address(),
				CampaignID:  0,
				Shares:      types.Shares{},
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
