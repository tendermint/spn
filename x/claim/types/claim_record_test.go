package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
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
				Claimable:         sdkmath.OneInt(),
				CompletedMissions: []uint64{},
			},
			valid: true,
		},
		{
			desc: "should prevent invalid address",
			claimRecord: claim.ClaimRecord{
				Address:           "invalid",
				Claimable:         sdkmath.OneInt(),
				CompletedMissions: []uint64{0, 1, 2},
			},
			valid: false,
		},
		{
			desc: "should prevent zero claimable amount",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.ZeroInt(),
				CompletedMissions: []uint64{0, 1, 2},
			},
			valid: false,
		},
		{
			desc: "should prevent negative claimable amount",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.NewInt(-1),
				CompletedMissions: []uint64{0, 1, 2},
			},
			valid: false,
		},
		{
			desc: "should prevent duplicate completed mission IDs",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.OneInt(),
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
		Claimable:         sdkmath.OneInt(),
		CompletedMissions: []uint64{},
	}.IsMissionCompleted(0))

	require.False(t, claim.ClaimRecord{
		Address:           sample.Address(r),
		Claimable:         sdkmath.OneInt(),
		CompletedMissions: []uint64{1, 2, 3},
	}.IsMissionCompleted(0))

	require.True(t, claim.ClaimRecord{
		Address:           sample.Address(r),
		Claimable:         sdkmath.OneInt(),
		CompletedMissions: []uint64{0, 1, 2, 3},
	}.IsMissionCompleted(0))

	require.True(t, claim.ClaimRecord{
		Address:           sample.Address(r),
		Claimable:         sdkmath.OneInt(),
		CompletedMissions: []uint64{0, 1, 2, 3},
	}.IsMissionCompleted(3))
}

func TestClaimRecord_ClaimableFromMission(t *testing.T) {
	for _, tt := range []struct {
		desc        string
		claimRecord claim.ClaimRecord
		mission     claim.Mission
		expected    sdkmath.Int
	}{
		{
			desc: "should allow get claimable from mission with full weight",
			claimRecord: claim.ClaimRecord{
				Claimable: sdkmath.NewIntFromUint64(100),
			},
			mission: claim.Mission{
				Weight: sdk.OneDec(),
			},
			expected: sdkmath.NewIntFromUint64(100),
		},
		{
			desc: "should allow get claimable from mission with empty weight",
			claimRecord: claim.ClaimRecord{
				Claimable: sdkmath.NewIntFromUint64(100),
			},
			mission: claim.Mission{
				Weight: sdk.ZeroDec(),
			},
			expected: sdkmath.ZeroInt(),
		},
		{
			desc: "should allow get claimable from mission with half weight",
			claimRecord: claim.ClaimRecord{
				Claimable: sdkmath.NewIntFromUint64(100),
			},
			mission: claim.Mission{
				Weight: tc.Dec(t, "0.5"),
			},
			expected: sdkmath.NewIntFromUint64(50),
		},
		{
			desc: "should allow get claimable and cut decimal",
			claimRecord: claim.ClaimRecord{
				Claimable: sdkmath.NewIntFromUint64(201),
			},
			mission: claim.Mission{
				Weight: tc.Dec(t, "0.5"),
			},
			expected: sdkmath.NewIntFromUint64(100),
		},
		{
			desc: "should return no claimable if decimal cut",
			claimRecord: claim.ClaimRecord{
				Claimable: sdkmath.NewIntFromUint64(1),
			},
			mission: claim.Mission{
				Weight: tc.Dec(t, "0.99"),
			},
			expected: sdkmath.NewIntFromUint64(0),
		},
	} {
		t.Run(tt.desc, func(t *testing.T) {
			got := tt.claimRecord.ClaimableFromMission(tt.mission)
			require.True(t, got.Equal(tt.expected),
				"expected: %s, got %s",
				tt.expected.String(),
				got.String(),
			)
		})
	}
}
