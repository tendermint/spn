package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func TestKeeper_CreateNewChain(t *testing.T) {
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)
	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.Address(r)
	coordNoProjectAddress := sample.Address(r)

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	res, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordID := res.CoordinatorID

	msgCreateCoordinator = sample.MsgCreateCoordinator(coordNoProjectAddress)
	res, err = ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordNoProjectID := res.CoordinatorID

	// Create a project
	msgCreateProject := sample.MsgCreateProject(r, coordAddress)
	resProject, err := ts.ProjectSrv.CreateProject(ctx, &msgCreateProject)
	require.NoError(t, err)
	projectID := resProject.ProjectID

	for _, tc := range []struct {
		name           string
		coordinatorID  uint64
		genesisChainID string
		sourceURL      string
		sourceHash     string
		initialGenesis types.InitialGenesis
		hasProject     bool
		projectID      uint64
		isMainnet      bool
		balance        sdk.Coins
		metadata       []byte
		wantedID       uint64
		valid          bool
	}{
		{
			name:           "should allow creating a new chain",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(r),
			sourceURL:      sample.String(r, 30),
			sourceHash:     sample.String(r, 20),
			initialGenesis: types.NewDefaultInitialGenesis(),
			hasProject:     false,
			balance:        sample.Coins(r),
			metadata:       sample.Metadata(r, 20),
			wantedID:       0,
			valid:          true,
		},
		{
			name:           "should allow creating a chain associated to a project",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(r),
			sourceURL:      sample.String(r, 30),
			sourceHash:     sample.String(r, 20),
			initialGenesis: types.NewDefaultInitialGenesis(),
			hasProject:     true,
			projectID:      projectID,
			isMainnet:      false,
			balance:        sample.Coins(r),
			metadata:       sample.Metadata(r, 20),
			wantedID:       1,
			valid:          true,
		},
		{
			name:           "should allow creating a mainnet chain",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(r),
			sourceURL:      sample.String(r, 30),
			sourceHash:     sample.String(r, 20),
			initialGenesis: types.NewDefaultInitialGenesis(),
			hasProject:     true,
			projectID:      0,
			isMainnet:      true,
			balance:        sample.Coins(r),
			metadata:       sample.Metadata(r, 20),
			wantedID:       2,
			valid:          true,
		},
		{
			name:           "should allow creating a chain with a custom genesis url",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(r),
			sourceURL:      sample.String(r, 30),
			sourceHash:     sample.String(r, 20),
			initialGenesis: types.NewGenesisURL(sample.String(r, 30), sample.GenesisHash(r)),
			hasProject:     false,
			balance:        sample.Coins(r),
			metadata:       sample.Metadata(r, 20),
			wantedID:       3,
			valid:          true,
		},
		{
			name:           "should allow creating a chain with a custom genesis config file",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(r),
			sourceURL:      sample.String(r, 30),
			sourceHash:     sample.String(r, 20),
			initialGenesis: types.NewGenesisConfig(sample.String(r, 30)),
			hasProject:     false,
			balance:        sample.Coins(r),
			metadata:       sample.Metadata(r, 20),
			wantedID:       4,
			valid:          true,
		},
		{
			name:           "should allow creating a chain with no metadata",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(r),
			sourceURL:      sample.String(r, 30),
			sourceHash:     sample.String(r, 20),
			initialGenesis: types.NewDefaultInitialGenesis(),
			hasProject:     true,
			projectID:      projectID,
			isMainnet:      false,
			balance:        sample.Coins(r),
			wantedID:       5,
			valid:          true,
		},
		{
			name:           "should prevent creating a chain with non-existent coordinator",
			coordinatorID:  100000,
			genesisChainID: sample.GenesisChainID(r),
			sourceURL:      sample.String(r, 30),
			sourceHash:     sample.String(r, 20),
			initialGenesis: types.NewDefaultInitialGenesis(),
			hasProject:     false,
			balance:        sample.Coins(r),
			metadata:       sample.Metadata(r, 20),
			wantedID:       0,
			valid:          false,
		},
		{
			name:           "should prevent creating a chain with non-existent project ID",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(r),
			sourceURL:      sample.String(r, 30),
			sourceHash:     sample.String(r, 20),
			initialGenesis: types.NewDefaultInitialGenesis(),
			hasProject:     true,
			projectID:      1000,
			balance:        sample.Coins(r),
			metadata:       sample.Metadata(r, 20),
			isMainnet:      false,
			valid:          false,
		},
		{
			name:           "should prevent creating a chain with invalid project coordinator",
			coordinatorID:  coordNoProjectID,
			genesisChainID: sample.GenesisChainID(r),
			sourceURL:      sample.String(r, 30),
			sourceHash:     sample.String(r, 20),
			initialGenesis: types.NewDefaultInitialGenesis(),
			hasProject:     true,
			projectID:      projectID,
			isMainnet:      false,
			balance:        sample.Coins(r),
			metadata:       sample.Metadata(r, 20),
			wantedID:       1,
			valid:          false,
		},
		{
			name:           "should prevent creating a chain with invalid chain data (mainnet with project)",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(r),
			sourceURL:      sample.String(r, 30),
			sourceHash:     sample.String(r, 20),
			initialGenesis: types.NewDefaultInitialGenesis(),
			hasProject:     false,
			projectID:      0,
			balance:        sample.Coins(r),
			metadata:       sample.Metadata(r, 20),
			isMainnet:      true,
			valid:          false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tk.LaunchKeeper.CreateNewChain(
				sdkCtx,
				tc.coordinatorID,
				tc.genesisChainID,
				tc.sourceURL,
				tc.sourceHash,
				tc.initialGenesis,
				tc.hasProject,
				tc.projectID,
				tc.isMainnet,
				tc.balance,
				tc.metadata,
			)

			if !tc.valid {
				require.Error(t, err)
				return
			}
			require.EqualValues(t, tc.wantedID, id)

			chain, found := tk.LaunchKeeper.GetChain(sdkCtx, id)
			require.True(t, found)
			require.EqualValues(t, tc.coordinatorID, chain.CoordinatorID)
			require.EqualValues(t, tc.genesisChainID, chain.GenesisChainID)
			require.EqualValues(t, tc.sourceURL, chain.SourceURL)
			require.EqualValues(t, tc.sourceHash, chain.SourceHash)
			require.EqualValues(t, tc.hasProject, chain.HasProject)
			require.EqualValues(t, tc.projectID, chain.ProjectID)
			require.EqualValues(t, tc.isMainnet, chain.IsMainnet)
			require.EqualValues(t, tc.metadata, chain.Metadata)
			require.EqualValues(t, tc.initialGenesis, chain.InitialGenesis)

			// Check chain has been appended in the project
			if tc.hasProject {
				projectChains, found := tk.ProjectKeeper.GetProjectChains(sdkCtx, tc.projectID)
				require.True(t, found)
				require.Contains(t, projectChains.Chains, id)
			}
		})
	}
}

func createNChain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i].LaunchID = keeper.AppendChain(ctx, items[i])
	}
	return items
}

func createNChainForCoordinator(keeper *keeper.Keeper, ctx sdk.Context, coordinatorID uint64, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i].CoordinatorID = coordinatorID
		items[i].LaunchID = keeper.AppendChain(ctx, items[i])
	}
	return items
}

func TestGetChain(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNChain(tk.LaunchKeeper, ctx, 10)

	t.Run("should get a chain", func(t *testing.T) {
		for _, item := range items {
			rst, found := tk.LaunchKeeper.GetChain(ctx, item.LaunchID)
			require.True(t, found)
			require.Equal(t, item, rst)
		}
	})
}

func TestEnableMonitoringConnection(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should enable monitoring connection for a chain", func(t *testing.T) {
		validChain := types.Chain{}
		validChainID := tk.LaunchKeeper.AppendChain(ctx, validChain)
		err := tk.LaunchKeeper.EnableMonitoringConnection(ctx, validChainID)
		require.NoError(t, err)
		rst, found := tk.LaunchKeeper.GetChain(ctx, validChainID)
		require.True(t, found)
		validChain.MonitoringConnected = true
		require.Equal(t, validChain, rst)
	})

	t.Run("should prevent enabling monitoring connection for non existing chain", func(t *testing.T) {
		// if chain does not exist, throw error
		err := tk.LaunchKeeper.EnableMonitoringConnection(ctx, 1)
		require.ErrorIs(t, err, types.ErrChainNotFound)
	})

	t.Run("should prevent enabling monitoring connection if already enabled", func(t *testing.T) {
		chain := types.Chain{}
		chainID := tk.LaunchKeeper.AppendChain(ctx, chain)
		err := tk.LaunchKeeper.EnableMonitoringConnection(ctx, chainID)
		require.NoError(t, err)
		err = tk.LaunchKeeper.EnableMonitoringConnection(ctx, chainID)
		require.ErrorIs(t, err, types.ErrChainMonitoringConnected)
	})
}

func TestGetAllChain(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNChain(tk.LaunchKeeper, ctx, 10)

	t.Run("should get all chains", func(t *testing.T) {
		require.ElementsMatch(t, items, tk.LaunchKeeper.GetAllChain(ctx))
	})
}

func TestChainCounter(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNChain(tk.LaunchKeeper, ctx, 10)

	t.Run("should get chain counter", func(t *testing.T) {
		counter := uint64(len(items))
		require.Equal(t, counter, tk.LaunchKeeper.GetChainCounter(ctx))
	})
}
