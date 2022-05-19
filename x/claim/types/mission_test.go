package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	claim "github.com/tendermint/spn/x/claim/types"
)

func TestMission_Validate(t *testing.T) {
	invalidWeightGreaterThanOne := sample.Mission(r)
	amt, err := sdk.NewDecFromStr("1.0000001")
	require.NoError(t, err)
	invalidWeightGreaterThanOne.Weight = amt

	invalidWeightLessThanZero := sample.Mission(r)
	amt, err = sdk.NewDecFromStr("-0.0000001")
	require.NoError(t, err)
	invalidWeightLessThanZero.Weight = amt

	for _, tc := range []struct {
		desc    string
		mission claim.Mission
		valid   bool
	}{
		{
			desc:    "valid mission",
			mission: sample.Mission(r),
			valid:   true,
		},
		{
			desc:    "invalid weight - greater than 1",
			mission: invalidWeightGreaterThanOne,
			valid:   false,
		},
		{
			desc:    "invalid weight - less than 0",
			mission: invalidWeightLessThanZero,
			valid:   false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, tc.mission.Validate() == nil)
		})
	}
}
