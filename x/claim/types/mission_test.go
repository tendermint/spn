package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
	claim "github.com/tendermint/spn/x/claim/types"
)

func TestMission_Validate(t *testing.T) {
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
			desc: "should accept weigth 0",
			mission: claim.Mission{
				MissionID:   sample.Uint64(r),
				Description: sample.String(r, 30),
				Weight:      tc.Dec(t, "0"),
			},
			valid: true,
		},
		{
			desc: "should accept weight 1",
			mission: claim.Mission{
				MissionID:   sample.Uint64(r),
				Description: sample.String(r, 30),
				Weight:      tc.Dec(t, "1"),
			},
			valid: true,
		},
		{
			desc: "should prevent weight greater than 1",
			mission: claim.Mission{
				MissionID:   sample.Uint64(r),
				Description: sample.String(r, 30),
				Weight:      tc.Dec(t, "1.0000001"),
			},
			valid: false,
		},
		{
			desc: "should prevent weight less than 0",
			mission: claim.Mission{
				MissionID:   sample.Uint64(r),
				Description: sample.String(r, 30),
				Weight:      tc.Dec(t, "-0.0000001"),
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, tc.mission.Validate() == nil)
		})
	}
}
