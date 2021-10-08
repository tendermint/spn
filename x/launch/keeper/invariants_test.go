package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
)

func TestDuplicatedAccountInvariant(t *testing.T) {
	k, _, _, _, _, _, ctx := setupMsgServer(t) //nolint
	t.Run("valid case", func(t *testing.T) {
		k.SetVestingAccount(ctx, sample.VestingAccount(0, sample.AccAddress()))
		k.SetGenesisAccount(ctx, sample.GenesisAccount(0, sample.AccAddress()))
		_, isValid := keeper.DuplicatedAccountInvariant(*k)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		addr := sample.AccAddress()
		k.SetVestingAccount(ctx, sample.VestingAccount(0, addr))
		k.SetGenesisAccount(ctx, sample.GenesisAccount(0, addr))
		_, isValid := keeper.DuplicatedAccountInvariant(*k)(ctx)
		require.Equal(t, true, isValid)
	})
}

func TestZeroLaunchTimestampInvariant(t *testing.T) {
	k, _, _, _, _, _, ctx := setupMsgServer(t) //nolint
	t.Run("valid case", func(t *testing.T) {
		chain := sample.Chain(0, 0)
		chain.LaunchTimestamp = 1000
		chain.Id = k.AppendChain(ctx, chain)
		_, isValid := keeper.ZeroLaunchTimestampInvariant(*k)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		chain := sample.Chain(0, 0)
		chain.LaunchTimestamp = 0
		chain.Id = k.AppendChain(ctx, chain)
		_, isValid := keeper.ZeroLaunchTimestampInvariant(*k)(ctx)
		require.Equal(t, true, isValid)
	})
}

func TestUnknownRequestTypeInvariant(t *testing.T) {
	k, _, _, _, _, _, ctx := setupMsgServer(t) //nolint
	t.Run("valid case", func(t *testing.T) {
		k.AppendRequest(ctx, sample.Request(0))
		_, isValid := keeper.UnknownRequestTypeInvariant(*k)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		k.AppendRequest(ctx, sample.RequestWithContent(0,
			sample.GenesisAccountContent(0, "invalid"),
		))
		_, isValid := keeper.UnknownRequestTypeInvariant(*k)(ctx)
		require.Equal(t, true, isValid)
	})
}
