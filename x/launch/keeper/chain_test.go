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
	k, _, campaignKeeper, _, profileSrv, campaignSrv, sdkCtx := setupMsgServer(t)
	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.Address()
	coordNoCampaignAddress := sample.Address()

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	res, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordID := res.CoordinatorID

	msgCreateCoordinator = sample.MsgCreateCoordinator(coordNoCampaignAddress)
	res, err = profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordNoCampaignID := res.CoordinatorID

	// Create a campaign
	msgCreateCampaign := sample.MsgCreateCampaign(coordAddress)
	resCampaign, err := campaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	campaignID := resCampaign.CampaignID

	for _, tc := range []struct {
		name           string
		coordinatorID  uint64
		genesisChainID string
		sourceURL      string
		sourceHash     string
		genesisURL     string
		genesisHash    string
		hasCampaign    bool
		campaignID     uint64
		isMainnet      bool
		metadata       []byte
		wantedID       uint64
		valid          bool
	}{
		{
			name:           "creating a new chain",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(),
			sourceURL:      sample.String(30),
			sourceHash:     sample.String(20),
			genesisURL:     "",
			hasCampaign:    false,
			metadata:       sample.Metadata(20),
			wantedID:       0,
			valid:          true,
		},
		{
			name:           "creating a chain associated to a campaign",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(),
			sourceURL:      sample.String(30),
			sourceHash:     sample.String(20),
			genesisURL:     "",
			hasCampaign:    true,
			campaignID:     campaignID,
			isMainnet:      false,
			metadata:       sample.Metadata(20),
			wantedID:       1,
			valid:          true,
		},
		{
			name:           "creating a mainnet chain",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(),
			sourceURL:      sample.String(30),
			sourceHash:     sample.String(20),
			genesisURL:     "",
			hasCampaign:    true,
			campaignID:     0,
			isMainnet:      true,
			metadata:       sample.Metadata(20),
			wantedID:       2,
			valid:          true,
		},
		{
			name:           "creating a chain with a custom genesis",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(),
			sourceURL:      sample.String(30),
			sourceHash:     sample.String(20),
			genesisURL:     sample.String(30),
			genesisHash:    sample.GenesisHash(),
			hasCampaign:    false,
			metadata:       sample.Metadata(20),
			wantedID:       3,
			valid:          true,
		},
		{
			name:           "creating a chain with no metadata",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(),
			sourceURL:      sample.String(30),
			sourceHash:     sample.String(20),
			genesisURL:     "",
			hasCampaign:    true,
			campaignID:     campaignID,
			isMainnet:      false,
			wantedID:       4,
			valid:          true,
		},
		{
			name:           "non-existent campaign ID",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(),
			sourceURL:      sample.String(30),
			sourceHash:     sample.String(20),
			genesisURL:     "",
			hasCampaign:    true,
			campaignID:     1000,
			metadata:       sample.Metadata(20),
			isMainnet:      false,
			valid:          false,
		},
		{
			name:           "invalid campaign coordinator",
			coordinatorID:  coordNoCampaignID,
			genesisChainID: sample.GenesisChainID(),
			sourceURL:      sample.String(30),
			sourceHash:     sample.String(20),
			genesisURL:     "",
			hasCampaign:    true,
			campaignID:     campaignID,
			isMainnet:      false,
			metadata:       sample.Metadata(20),
			wantedID:       1,
			valid:          false,
		},
		{
			name:           "invalid chain data (mainnet with campaign)",
			coordinatorID:  coordID,
			genesisChainID: sample.GenesisChainID(),
			sourceURL:      sample.String(30),
			sourceHash:     sample.String(20),
			genesisURL:     "",
			hasCampaign:    false,
			campaignID:     0,
			metadata:       sample.Metadata(20),
			isMainnet:      true,
			valid:          false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			id, err := k.CreateNewChain(
				sdkCtx,
				tc.coordinatorID,
				tc.genesisChainID,
				tc.sourceURL,
				tc.sourceHash,
				tc.genesisURL,
				tc.genesisHash,
				tc.hasCampaign,
				tc.campaignID,
				tc.isMainnet,
				tc.metadata,
			)

			if !tc.valid {
				require.Error(t, err)
				return
			}
			require.EqualValues(t, tc.wantedID, id)

			chain, found := k.GetChain(sdkCtx, id)
			require.True(t, found)
			require.EqualValues(t, tc.coordinatorID, chain.CoordinatorID)
			require.EqualValues(t, tc.genesisChainID, chain.GenesisChainID)
			require.EqualValues(t, tc.sourceURL, chain.SourceURL)
			require.EqualValues(t, tc.sourceHash, chain.SourceHash)
			require.EqualValues(t, tc.hasCampaign, chain.HasCampaign)
			require.EqualValues(t, tc.campaignID, chain.CampaignID)
			require.EqualValues(t, tc.isMainnet, chain.IsMainnet)
			require.EqualValues(t, tc.metadata, chain.Metadata)

			// Compare initial genesis
			if tc.genesisURL == "" {
				require.Equal(t, types.NewDefaultInitialGenesis(), chain.InitialGenesis)
			} else {
				require.Equal(
					t,
					types.NewGenesisURL(tc.genesisURL, tc.genesisHash),
					chain.InitialGenesis,
				)
			}

			// Check chain has been appended in the campaign
			if tc.hasCampaign {
				campaignChains, found := campaignKeeper.GetCampaignChains(sdkCtx, tc.campaignID)
				require.True(t, found)
				require.Contains(t, campaignChains.Chains, id)
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
	keeper, ctx := testkeeper.Launch(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChain(ctx, item.LaunchID)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestGetAllChain(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNChain(keeper, ctx, 10)

	require.ElementsMatch(t, items, keeper.GetAllChain(ctx))
}

func TestChainCounter(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNChain(keeper, ctx, 10)
	counter := uint64(len(items))
	require.Equal(t, counter, keeper.GetChainCounter(ctx))
}
