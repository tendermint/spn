package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	claim "github.com/tendermint/spn/x/claim/types"
)

func TestClaimRecord_Validate(t *testing.T) {
	invalidAddress := sample.ClaimRecord(r)
	invalidAddress.Address = "invalid"

	invalidClaimable := sample.ClaimRecord(r)
	invalidClaimable.Claimable = sdk.NewInt(-1)

	invalidCompletedMissions := sample.ClaimRecord(r)
	invalidCompletedMissions.CompletedMissions = []uint64{0, 1, 2, 0}

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
			desc:        "invalid address",
			claimRecord: invalidAddress,
			valid:       false,
		},
		{
			desc:        "invalid claimable amount",
			claimRecord: invalidClaimable,
			valid:       false,
		},
		{
			desc:        "duplicate completed mission IDs",
			claimRecord: invalidCompletedMissions,
			valid:       false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, tc.claimRecord.Validate() == nil)
		})
	}
}
