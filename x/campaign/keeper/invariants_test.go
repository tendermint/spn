package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestAccountWithoutCampaignInvariant(t *testing.T) {
	k, _, _, _, _, _, ctx := setupMsgServer(t) //nolint
	t.Run("valid case", func(t *testing.T) {
		campaign := sample.Campaign(0)
		campaign.Id = k.AppendCampaign(ctx, campaign)
		k.SetMainnetAccount(ctx, sample.MainnetAccount(campaign.Id, sample.Address()))
		_, isValid := keeper.AccountWithoutCampaignInvariant(*k)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		k.SetMainnetAccount(ctx, sample.MainnetAccount(100, sample.Address()))
		_, isValid := keeper.AccountWithoutCampaignInvariant(*k)(ctx)
		require.Equal(t, true, isValid)
	})
}

func TestVestingAccountWithoutCampaignInvariant(t *testing.T) {
	k, _, _, _, _, _, ctx := setupMsgServer(t) //nolint
	t.Run("valid case", func(t *testing.T) {
		campaign := sample.Campaign(0)
		campaign.Id = k.AppendCampaign(ctx, campaign)
		k.SetMainnetVestingAccount(ctx, sample.MainnetVestingAccount(campaign.Id, sample.Address()))
		_, isValid := keeper.VestingAccountWithoutCampaignInvariant(*k)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		k.SetMainnetVestingAccount(ctx, sample.MainnetVestingAccount(10000, sample.Address()))
		_, isValid := keeper.VestingAccountWithoutCampaignInvariant(*k)(ctx)
		require.Equal(t, true, isValid)
	})
}

func TestCampaignChainsWithoutCampaignInvariant(t *testing.T) {
	k, _, _, _, _, _, ctx := setupMsgServer(t) //nolint
	t.Run("valid case", func(t *testing.T) {
		campaign := sample.Campaign(0)
		campaign.Id = k.AppendCampaign(ctx, campaign)
		k.SetCampaignChains(ctx, types.CampaignChains{CampaignID: campaign.Id})
		_, isValid := keeper.CampaignChainsWithoutCampaignInvariant(*k)(ctx)
		require.Equal(t, false, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		k.SetCampaignChains(ctx, types.CampaignChains{CampaignID: 1000})
		_, isValid := keeper.CampaignChainsWithoutCampaignInvariant(*k)(ctx)
		require.Equal(t, true, isValid)
	})
}
