package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
)

func TestDuplicatedAccountInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		tk.LaunchKeeper.SetVestingAccount(ctx, sample.VestingAccount(r, 0, sample.Address(r)))
		tk.LaunchKeeper.SetGenesisAccount(ctx, sample.GenesisAccount(r, 0, sample.Address(r)))
		_, isValid := keeper.DuplicatedAccountInvariant(*tk.LaunchKeeper)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		addr := sample.Address(r)
		tk.LaunchKeeper.SetVestingAccount(ctx, sample.VestingAccount(r, 0, addr))
		tk.LaunchKeeper.SetGenesisAccount(ctx, sample.GenesisAccount(r, 0, addr))
		_, isValid := keeper.DuplicatedAccountInvariant(*tk.LaunchKeeper)(ctx)
		require.Equal(t, true, isValid)
	})
}

func TestZeroLaunchTimestampInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		chain := sample.Chain(r, 0, 0)
		chain.LaunchTimestamp = 1000
		chain.LaunchID = tk.LaunchKeeper.AppendChain(ctx, chain)
		_, isValid := keeper.ZeroLaunchTimestampInvariant(*tk.LaunchKeeper)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		chain := sample.Chain(r, 0, 0)
		chain.LaunchTimestamp = 0
		chain.LaunchID = tk.LaunchKeeper.AppendChain(ctx, chain)
		_, isValid := keeper.ZeroLaunchTimestampInvariant(*tk.LaunchKeeper)(ctx)
		require.Equal(t, true, isValid)
	})
}

func TestUnknownRequestTypeInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		tk.LaunchKeeper.AppendRequest(ctx, sample.Request(r, 0, sample.Address(r)))
		_, isValid := keeper.UnknownRequestTypeInvariant(*tk.LaunchKeeper)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		tk.LaunchKeeper.AppendRequest(ctx, sample.RequestWithContent(r, 0,
			sample.GenesisAccountContent(r, 0, "invalid"),
		))
		_, isValid := keeper.UnknownRequestTypeInvariant(*tk.LaunchKeeper)(ctx)
		require.Equal(t, true, isValid)
	})
}
