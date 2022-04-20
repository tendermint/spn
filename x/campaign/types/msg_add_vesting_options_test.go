package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgAddVestingOptions_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgAddVestingOptions
		err  error
	}{
		{
			name: "valid message",
			msg: types.MsgAddVestingOptions{
				Coordinator:    sample.Address(r),
				CampaignID:     0,
				VestingOptions: sample.ShareVestingOptions(r),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgAddVestingOptions{
				Coordinator:    "invalid_address",
				CampaignID:     0,
				VestingOptions: sample.ShareVestingOptions(r),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid vesting options",
			msg: types.MsgAddVestingOptions{
				Coordinator:    "invalid_address",
				CampaignID:     0,
				VestingOptions: *types.NewShareDelayedVesting(sample.Shares(r), sample.Shares(r), 0),
			},
			err: sdkerrors.ErrInvalidAddress,
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
