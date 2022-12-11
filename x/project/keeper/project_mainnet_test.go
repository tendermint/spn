package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
)

func TestIsProjectMainnetLaunchTriggered(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should show project has mainnet with launch triggered", func(t *testing.T) {
		projectMainnetLaunched := sample.Project(r, 0)
		projectMainnetLaunched.MainnetInitialized = true
		chainLaunched := sample.Chain(r, 0, 0)
		chainLaunched.LaunchTriggered = true
		chainLaunched.IsMainnet = true
		projectMainnetLaunched.MainnetID = tk.LaunchKeeper.AppendChain(ctx, chainLaunched)
		projectMainnetLaunched.ProjectID = tk.ProjectKeeper.AppendProject(ctx, projectMainnetLaunched)
		res, err := tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, projectMainnetLaunched.ProjectID)
		require.NoError(t, err)
		require.True(t, res)
	})

	t.Run("should show project has mainnet with launch not triggered", func(t *testing.T) {
		projectMainnetInitialized := sample.Project(r, 1)
		projectMainnetInitialized.MainnetInitialized = true
		chain := sample.Chain(r, 0, 0)
		chain.LaunchTriggered = false
		chain.IsMainnet = true
		projectMainnetInitialized.MainnetID = tk.LaunchKeeper.AppendChain(ctx, chain)
		projectMainnetInitialized.ProjectID = tk.ProjectKeeper.AppendProject(ctx, projectMainnetInitialized)
		res, err := tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, projectMainnetInitialized.ProjectID)
		require.NoError(t, err)
		require.False(t, res)
	})

	t.Run("should show project with mainnnet not initialized", func(t *testing.T) {
		projectMainnetNotInitialized := sample.Project(r, 2)
		projectMainnetNotInitialized.MainnetInitialized = false
		projectMainnetNotInitialized.ProjectID = tk.ProjectKeeper.AppendProject(ctx, projectMainnetNotInitialized)
		res, err := tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, projectMainnetNotInitialized.ProjectID)
		require.NoError(t, err)
		require.False(t, res)
	})

	t.Run("should show mainnet not found", func(t *testing.T) {
		projectMainnetNotFound := sample.Project(r, 3)
		projectMainnetNotFound.MainnetInitialized = true
		projectMainnetNotFound.MainnetID = 1000
		projectMainnetNotFound.ProjectID = tk.ProjectKeeper.AppendProject(ctx, projectMainnetNotFound)
		_, err := tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, projectMainnetNotFound.ProjectID)
		require.Error(t, err)
	})

	t.Run("should show associated mainnet chain is not mainnet", func(t *testing.T) {
		projectInvalidMainnet := sample.Project(r, 4)
		projectInvalidMainnet.MainnetInitialized = true
		chainNoMainnet := sample.Chain(r, 0, 0)
		chainNoMainnet.LaunchTriggered = false
		chainNoMainnet.IsMainnet = false
		projectInvalidMainnet.MainnetID = tk.LaunchKeeper.AppendChain(ctx, chainNoMainnet)
		projectInvalidMainnet.ProjectID = tk.ProjectKeeper.AppendProject(ctx, projectInvalidMainnet)
		_, err := tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, projectInvalidMainnet.ProjectID)
		require.Error(t, err)
	})

	t.Run("should show project not found", func(t *testing.T) {
		_, err := tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, 1000)
		require.Error(t, err)
	})
}
