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
		msg, broken := keeper.DuplicatedAccountInvariant(*tk.LaunchKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("invalid case", func(t *testing.T) {
		addr := sample.Address(r)
		tk.LaunchKeeper.SetVestingAccount(ctx, sample.VestingAccount(r, 0, addr))
		tk.LaunchKeeper.SetGenesisAccount(ctx, sample.GenesisAccount(r, 0, addr))
		msg, broken := keeper.DuplicatedAccountInvariant(*tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestInvalidChainInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		chain := sample.Chain(r, 0, 0)
		campaign := sample.Campaign(r, 0)
		chain.LaunchID = tk.LaunchKeeper.AppendChain(ctx, chain)
		chain.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaign)
		chain.HasCampaign = true
		msg, broken := keeper.InvalidChainInvariant(*tk.LaunchKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("invalid case", func(t *testing.T) {
		chain1 := sample.Chain(r, 0, 0)
		chain1.LaunchTriggered = true
		chain1.LaunchID = tk.LaunchKeeper.AppendChain(ctx, chain1)
		chain2 := sample.Chain(r, 1, 0)
		chain2.HasCampaign = true
		chain2.CampaignID = 1000
		msg, broken := keeper.InvalidChainInvariant(*tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestUnknownRequestTypeInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		tk.LaunchKeeper.AppendRequest(ctx, sample.Request(r, 0, sample.Address(r)))
		msg, broken := keeper.UnknownRequestTypeInvariant(*tk.LaunchKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("invalid case", func(t *testing.T) {
		tk.LaunchKeeper.AppendRequest(ctx, sample.RequestWithContent(r, 0,
			sample.GenesisAccountContent(r, 0, "invalid"),
		))
		msg, broken := keeper.UnknownRequestTypeInvariant(*tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
