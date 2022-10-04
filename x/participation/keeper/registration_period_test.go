package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
)

func TestIsRegistrationEnabled(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	registrationPeriod := time.Hour

	params := tk.ParticipationKeeper.GetParams(ctx)
	params.RegistrationPeriod = registrationPeriod
	tk.ParticipationKeeper.SetParams(ctx, params)

	for _, tc := range []struct {
		name             string
		auctionStartTime time.Time
		blockTime        time.Time
		expected         bool
	}{
		{
			name:             "should prevent with registration window not yet started",
			auctionStartTime: ctx.BlockTime().Add(time.Hour * 5),
			blockTime:        ctx.BlockTime(),
			expected:         false,
		},
		{
			name:             "should prevent if auction is started",
			auctionStartTime: ctx.BlockTime(),
			blockTime:        ctx.BlockTime(),
			expected:         false,
		},
		{
			name:             "should allow with registration enabled",
			auctionStartTime: ctx.BlockTime().Add(time.Minute * 30),
			blockTime:        ctx.BlockTime(),
			expected:         true,
		},
		{
			name: "should allow with registration enabled when registration period is longer " +
				"than range between Unix time 0 and auction's start time",
			auctionStartTime: time.Unix(int64((registrationPeriod - time.Minute).Seconds()), 0),
			blockTime:        time.Unix(1, 0),
			expected:         true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tmpCtx := ctx.WithBlockTime(tc.blockTime)
			res := tk.ParticipationKeeper.IsRegistrationEnabled(tmpCtx, tc.auctionStartTime)
			require.Equal(t, tc.expected, res)
		})
	}
}
