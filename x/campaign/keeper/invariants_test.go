package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/keeper"
)

func TestAccountWithoutCampaignInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		campaign := sample.Campaign(r, 0)
		campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaign)
		tk.CampaignKeeper.SetMainnetAccount(ctx, sample.MainnetAccount(r, campaign.CampaignID, sample.Address(r)))
		_, isValid := keeper.AccountWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		tk.CampaignKeeper.SetMainnetAccount(ctx, sample.MainnetAccount(r, 100, sample.Address(r)))
		_, isValid := keeper.AccountWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.Equal(t, true, isValid)
	})
}

func TestVestingAccountWithoutCampaignInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		campaign := sample.Campaign(r, 0)
		campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaign)
		tk.CampaignKeeper.SetMainnetVestingAccount(ctx, sample.MainnetVestingAccount(r, campaign.CampaignID, sample.Address(r)))
		_, isValid := keeper.VestingAccountWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		tk.CampaignKeeper.SetMainnetVestingAccount(ctx, sample.MainnetVestingAccount(r, 10000, sample.Address(r)))
		_, isValid := keeper.VestingAccountWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.Equal(t, true, isValid)
	})
}
