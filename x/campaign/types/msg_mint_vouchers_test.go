package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgMintVouchers_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgMintVouchers
		err  error
	}{
		{
			name: "valid message",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				CampaignID:  0,
				Shares:      sample.Shares(r),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgMintVouchers{
				Coordinator: "invalid_address",
				CampaignID:  0,
				Shares:      sample.Shares(r),
			},
			err: profile.ErrInvalidCoordAddress,
		},
		{
			name: "invalid shares",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				CampaignID:  0,
				Shares:      invalidShares,
			},
			err: types.ErrInvalidShares,
		},
		{
			name: "empty shares",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				CampaignID:  0,
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
