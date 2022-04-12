package simulation_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	profilesim "github.com/tendermint/spn/x/profile/simulation"
	"github.com/tendermint/spn/x/profile/types"
)

func TestFindCoordinatorAccount(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		r          = sample.Rand()
		accs       = simulation.RandomAccounts(r, 20)
	)

	t.Run("false for empty set", func(t *testing.T) {
		_, found := profilesim.FindCoordinatorAccount(r, ctx, *tk.ProfileKeeper, []simulation.Account{}, false)
		require.False(t, found)
		_, found = profilesim.FindCoordinatorAccount(r, ctx, *tk.ProfileKeeper, []simulation.Account{}, true)
		require.False(t, found)
	})

	t.Run("no existing coordinator account", func(t *testing.T) {
		_, found := profilesim.FindCoordinatorAccount(r, ctx, *tk.ProfileKeeper, accs, true)
		require.False(t, found)

		acc, found := profilesim.FindCoordinatorAccount(r, ctx, *tk.ProfileKeeper, accs, false)
		require.True(t, found)
		require.Contains(t, accs, acc)
	})

	// Set ccordinator to a random account
	acc, _ := simulation.RandomAcc(r, accs)
	tk.ProfileKeeper.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       acc.Address.String(),
		CoordinatorID: sample.Uint64(r),
	})

	t.Run("one coordinator account", func(t *testing.T) {
		acc, found := profilesim.FindCoordinatorAccount(r, ctx, *tk.ProfileKeeper, accs, true)
		require.True(t, found)
		require.Contains(t, accs, acc)

		acc, found = profilesim.FindCoordinatorAccount(r, ctx, *tk.ProfileKeeper, accs, false)
		require.True(t, found)
		require.Contains(t, accs, acc)
	})

	t.Run("all coordinator account", func(t *testing.T) {
		_, found := profilesim.FindCoordinatorAccount(r, ctx, *tk.ProfileKeeper, []simulation.Account{acc}, false)
		require.False(t, found)

		acc, found := profilesim.FindCoordinatorAccount(r, ctx, *tk.ProfileKeeper, []simulation.Account{acc}, true)
		require.True(t, found)
		require.Contains(t, accs, acc)
	})
}
