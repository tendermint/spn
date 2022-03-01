package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestAccountWithoutCampaignInvariant(t *testing.T) {
	ctx, tk, _, _ := setupMsgServer(t) //nolint
	t.Run("valid case", func(t *testing.T) {
		campaign := sample.Campaign(0)
		campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaign)
		tk.CampaignKeeper.SetMainnetAccount(ctx, sample.MainnetAccount(campaign.CampaignID, sample.Address()))
		_, isValid := keeper.AccountWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		tk.CampaignKeeper.SetMainnetAccount(ctx, sample.MainnetAccount(100, sample.Address()))
		_, isValid := keeper.AccountWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.Equal(t, true, isValid)
	})
}

func TestVestingAccountWithoutCampaignInvariant(t *testing.T) {
	ctx, tk, _, _ := setupMsgServer(t) //nolint
	t.Run("valid case", func(t *testing.T) {
		campaign := sample.Campaign(0)
		campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaign)
		tk.CampaignKeeper.SetMainnetVestingAccount(ctx, sample.MainnetVestingAccount(campaign.CampaignID, sample.Address()))
		_, isValid := keeper.VestingAccountWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		tk.CampaignKeeper.SetMainnetVestingAccount(ctx, sample.MainnetVestingAccount(10000, sample.Address()))
		_, isValid := keeper.VestingAccountWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.Equal(t, true, isValid)
	})
}

func TestCampaignChainsWithoutCampaignInvariant(t *testing.T) {
	ctx, tk, _, _ := setupMsgServer(t) //nolint
	t.Run("valid case", func(t *testing.T) {
		campaign := sample.Campaign(0)
		campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaign)
		tk.CampaignKeeper.SetCampaignChains(ctx, types.CampaignChains{CampaignID: campaign.CampaignID})
		_, isValid := keeper.CampaignChainsWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		tk.CampaignKeeper.SetCampaignChains(ctx, types.CampaignChains{CampaignID: 1000})
		_, isValid := keeper.CampaignChainsWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.Equal(t, true, isValid)
	})
}
