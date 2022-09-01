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
	t.Run("should not break with valid state", func(t *testing.T) {
		tk.LaunchKeeper.SetVestingAccount(ctx, sample.VestingAccount(r, 0, sample.Address(r)))
		tk.LaunchKeeper.SetGenesisAccount(ctx, sample.GenesisAccount(r, 0, sample.Address(r)))
		msg, broken := keeper.DuplicatedAccountInvariant(*tk.LaunchKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should break with duplicated account", func(t *testing.T) {
		addr := sample.Address(r)
		tk.LaunchKeeper.SetVestingAccount(ctx, sample.VestingAccount(r, 0, addr))
		tk.LaunchKeeper.SetGenesisAccount(ctx, sample.GenesisAccount(r, 0, addr))
		msg, broken := keeper.DuplicatedAccountInvariant(*tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestInvalidChainInvariant(t *testing.T) {
	t.Run("should not break with valid state", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		chain := sample.Chain(r, 0, 0)
		campaign := sample.Campaign(r, 0)
		chain.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaign)
		chain.HasCampaign = true
		_ = tk.LaunchKeeper.AppendChain(ctx, chain)
		msg, broken := keeper.InvalidChainInvariant(*tk.LaunchKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should break with an invalid chain", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		chain := sample.Chain(r, 0, 0)
		chain.GenesisChainID = "_invalid_"
		_ = tk.LaunchKeeper.AppendChain(ctx, chain)
		msg, broken := keeper.InvalidChainInvariant(*tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should break with a chain that does not have a valid associated campaign", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		chain := sample.Chain(r, 0, 0)
		chain.HasCampaign = true
		chain.CampaignID = 1000
		_ = tk.LaunchKeeper.AppendChain(ctx, chain)
		msg, broken := keeper.InvalidChainInvariant(*tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestUnknownRequestTypeInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should not break with valid state", func(t *testing.T) {
		tk.LaunchKeeper.AppendRequest(ctx, sample.Request(r, 0, sample.Address(r)))
		msg, broken := keeper.UnknownRequestTypeInvariant(*tk.LaunchKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should break with an invalid request", func(t *testing.T) {
		tk.LaunchKeeper.AppendRequest(ctx, sample.RequestWithContent(r, 0,
			sample.GenesisAccountContent(r, 0, "invalid"),
		))
		msg, broken := keeper.UnknownRequestTypeInvariant(*tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
