package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
)

func TestIsCampaignMainnetLaunchTriggered(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should show campaign has mainnet with launch triggered", func(t *testing.T) {
		campaignMainnetLaunched := sample.Campaign(r, 0)
		campaignMainnetLaunched.MainnetInitialized = true
		chainLaunched := sample.Chain(r, 0, 0)
		chainLaunched.LaunchTriggered = true
		chainLaunched.IsMainnet = true
		campaignMainnetLaunched.MainnetID = tk.LaunchKeeper.AppendChain(ctx, chainLaunched)
		campaignMainnetLaunched.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaignMainnetLaunched)
		res, err := tk.CampaignKeeper.IsCampaignMainnetLaunchTriggered(ctx, campaignMainnetLaunched.CampaignID)
		require.NoError(t, err)
		require.True(t, res)
	})

	t.Run("should show campaign has mainnet with launch not triggered", func(t *testing.T) {
		campaignMainnetInitialized := sample.Campaign(r, 1)
		campaignMainnetInitialized.MainnetInitialized = true
		chain := sample.Chain(r, 0, 0)
		chain.LaunchTriggered = false
		chain.IsMainnet = true
		campaignMainnetInitialized.MainnetID = tk.LaunchKeeper.AppendChain(ctx, chain)
		campaignMainnetInitialized.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaignMainnetInitialized)
		res, err := tk.CampaignKeeper.IsCampaignMainnetLaunchTriggered(ctx, campaignMainnetInitialized.CampaignID)
		require.NoError(t, err)
		require.False(t, res)
	})

	t.Run("should show campaign with mainnnet not initialized", func(t *testing.T) {
		campaignMainnetNotInitialized := sample.Campaign(r, 2)
		campaignMainnetNotInitialized.MainnetInitialized = false
		campaignMainnetNotInitialized.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaignMainnetNotInitialized)
		res, err := tk.CampaignKeeper.IsCampaignMainnetLaunchTriggered(ctx, campaignMainnetNotInitialized.CampaignID)
		require.NoError(t, err)
		require.False(t, res)
	})

	t.Run("should show mainnet not found", func(t *testing.T) {
		campaignMainnetNotFound := sample.Campaign(r, 3)
		campaignMainnetNotFound.MainnetInitialized = true
		campaignMainnetNotFound.MainnetID = 1000
		campaignMainnetNotFound.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaignMainnetNotFound)
		_, err := tk.CampaignKeeper.IsCampaignMainnetLaunchTriggered(ctx, campaignMainnetNotFound.CampaignID)
		require.Error(t, err)
	})

	t.Run("should show associated mainnet chain is not mainnet", func(t *testing.T) {
		campaignInvalidMainnet := sample.Campaign(r, 4)
		campaignInvalidMainnet.MainnetInitialized = true
		chainNoMainnet := sample.Chain(r, 0, 0)
		chainNoMainnet.LaunchTriggered = false
		chainNoMainnet.IsMainnet = false
		campaignInvalidMainnet.MainnetID = tk.LaunchKeeper.AppendChain(ctx, chainNoMainnet)
		campaignInvalidMainnet.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaignInvalidMainnet)
		_, err := tk.CampaignKeeper.IsCampaignMainnetLaunchTriggered(ctx, campaignInvalidMainnet.CampaignID)
		require.Error(t, err)
	})

	t.Run("should show campaign not found", func(t *testing.T) {
		_, err := tk.CampaignKeeper.IsCampaignMainnetLaunchTriggered(ctx, 1000)
		require.Error(t, err)
	})
}
