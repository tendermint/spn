package simulation_test

import (
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/profile/types"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/tendermint/spn/testutil/sample"
	profilesim "github.com/tendermint/spn/x/profile/simulation"
)

func TestFindCoordinatorAccount(t *testing.T) {
	var (
		k, ctx = testkeeper.Profile(t)
		r = sample.Rand()
		accs = simulation.RandomAccounts(r, 20)
	)

	t.Run("false for empty set", func(t *testing.T) {
		_, found := profilesim.FindCoordinatorAccount(r, ctx, *k, []simulation.Account{}, false)
		require.False(t, found)
		_, found = profilesim.FindCoordinatorAccount(r, ctx, *k, []simulation.Account{}, true)
		require.False(t, found)
	})

	t.Run("no existing coordinator account", func(t *testing.T) {
		_, found := profilesim.FindCoordinatorAccount(r, ctx, *k, accs, true)
		require.False(t, found)

		acc, found := profilesim.FindCoordinatorAccount(r, ctx, *k, accs, false)
		require.True(t, found)
		require.Contains(t, accs, acc)
	})

	// Set ccordinator to a random account
	acc, accIndex := simulation.RandomAcc(r, accs)
	k.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address: acc.Address.String(),
		CoordinatorId: sample.Uint64(),
	})

	t.Run("one coordinator account", func(t *testing.T) {
		acc, found := profilesim.FindCoordinatorAccount(r, ctx, *k, accs, true)
		require.True(t, found)
		require.Contains(t, accs, acc)

		acc, found = profilesim.FindCoordinatorAccount(r, ctx, *k, accs, false)
		require.True(t, found)
		require.Contains(t, accs, acc)
	})

	t.Run("all coordinator account", func(t *testing.T) {
		_, found := profilesim.FindCoordinatorAccount(r, ctx, *k, accs[accIndex:accIndex+1], false)
		require.False(t, found)

		acc, found := profilesim.FindCoordinatorAccount(r, ctx, *k, accs[accIndex:accIndex+1], true)
		require.True(t, found)
		require.Contains(t, accs, acc)
	})
}