package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	claim "github.com/tendermint/spn/x/claim/types"
)

func TestClaimRecord_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc        string
		claimRecord claim.ClaimRecord
		valid       bool
	}{
		{
			desc:        "valid claim record",
			claimRecord: sample.ClaimRecord(r),
			valid:       true,
		},
		{
			desc: "claim record with no completed mission is valid",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdk.OneInt(),
				CompletedMissions: []uint64{},
			},
			valid: true,
		},
		{
			desc: "should prevent invalid address",
			claimRecord: claim.ClaimRecord{
				Address:           "invalid",
				Claimable:         sdk.OneInt(),
				CompletedMissions: []uint64{0, 1, 2},
			},
			valid: false,
		},
		{
			desc: "should prevent zero claimable amount",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdk.ZeroInt(),
				CompletedMissions: []uint64{0, 1, 2},
			},
			valid: false,
		},
		{
			desc: "should prevent negative claimable amount",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdk.NewInt(-1),
				CompletedMissions: []uint64{0, 1, 2},
			},
			valid: false,
		},
		{
			desc: "should prevent duplicate completed mission IDs",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdk.OneInt(),
				CompletedMissions: []uint64{0, 1, 2, 0},
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, tc.claimRecord.Validate() == nil)
		})
	}
}

func TestClaimRecord_IsMissionCompleted(t *testing.T) {
	require.False(t, claim.ClaimRecord{
		Address:           sample.Address(r),
		Claimable:         sdk.OneInt(),
		CompletedMissions: []uint64{},
	}.IsMissionCompleted(0))

	require.False(t, claim.ClaimRecord{
		Address:           sample.Address(r),
		Claimable:         sdk.OneInt(),
		CompletedMissions: []uint64{1, 2, 3},
	}.IsMissionCompleted(0))

	require.True(t, claim.ClaimRecord{
		Address:           sample.Address(r),
		Claimable:         sdk.OneInt(),
		CompletedMissions: []uint64{0, 1, 2, 3},
	}.IsMissionCompleted(0))

	require.True(t, claim.ClaimRecord{
		Address:           sample.Address(r),
		Claimable:         sdk.OneInt(),
		CompletedMissions: []uint64{0, 1, 2, 3},
	}.IsMissionCompleted(3))
}
