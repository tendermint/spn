package keeper_test

import (
	"testing"

	"github.com/tendermint/spn/testutil/sample"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
)

func TestGetParams(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow params get", func(t *testing.T) {
		params := sample.ParticipationParams(r)
		tk.ParticipationKeeper.SetParams(ctx, params)
		require.EqualValues(t, params, tk.ParticipationKeeper.GetParams(ctx))
		require.EqualValues(t, params.AllocationPrice, tk.ParticipationKeeper.AllocationPrice(ctx))
		require.EqualValues(t, params.ParticipationTierList, tk.ParticipationKeeper.ParticipationTierList(ctx))
		require.EqualValues(t, params.RegistrationPeriod, tk.ParticipationKeeper.RegistrationPeriod(ctx))
		require.EqualValues(t, params.WithdrawalDelay, tk.ParticipationKeeper.WithdrawalDelay(ctx))
	})
}
