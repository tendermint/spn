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

func TestMsgUpdateSpecialAllocations_ValidateBasic(t *testing.T) {
	invalidShares := types.Shares{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	tests := []struct {
		name string
		msg  types.MsgUpdateSpecialAllocations
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgUpdateSpecialAllocations{
				Coordinator:        sample.Address(r),
				CampaignID:         1,
				SpecialAllocations: sample.SpecialAllocations(r),
			},
		},
		{
			name: "should prevent validation of msg with invalid address",
			msg: types.MsgUpdateSpecialAllocations{
				Coordinator:        "invalid_address",
				CampaignID:         1,
				SpecialAllocations: sample.SpecialAllocations(r),
			},
			err: profile.ErrInvalidCoordAddress,
		},
		{
			name: "should prevent validation of msg with invalid special allocations",
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
