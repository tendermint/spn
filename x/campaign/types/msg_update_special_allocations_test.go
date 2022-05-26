package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgUpdateSpecialAllocations_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdateSpecialAllocations
		err  error
	}{
		{
			name: "valid message",
			msg: types.MsgUpdateSpecialAllocations{
				Coordinator:        sample.Address(r),
				CampaignID:         1,
				SpecialAllocations: sample.SpecialAllocations(r),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgUpdateSpecialAllocations{
				Coordinator:        "invalid_address",
				CampaignID:         1,
				SpecialAllocations: sample.SpecialAllocations(r),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid special allocations",
			msg: types.MsgUpdateSpecialAllocations{
				Coordinator:        sample.Address(r),
				CampaignID:         1,
				SpecialAllocations: types.NewSpecialAllocations(invalidShares, sample.Shares(r)),
			},
			err: types.ErrInvalidSpecialAllocations,
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
