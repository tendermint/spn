package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgInitializeMainnet(t *testing.T) {
	var (
		campaignID                   uint64 = 0
		campaignMainnetInitializedID uint64 = 1
		campaignIncorrectCoordID     uint64 = 2
		campaignEmptySupplyID        uint64 = 3
		coordAddr                           = sample.Address()
		coordAddrNoCampaign                 = sample.Address()

		campaignKeeper, _, launchKeeper, _, campaignSrv, profileSrv, sdkCtx = setupMsgServer(t)
		ctx                                                                 = sdk.WrapSDKContext(sdkCtx)
	)

	// Create coordinators
	res, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	coordID := res.CoordinatorID
	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddrNoCampaign,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)

	// Set different campaigns
	campaign := sample.Campaign(campaignID)
	campaign.CoordinatorID = coordID
	campaignKeeper.SetCampaign(sdkCtx, campaign)

	campaignMainnetInitialized := sample.Campaign(campaignMainnetInitializedID)
	campaignMainnetInitialized.CoordinatorID = coordID
	campaignMainnetInitialized.MainnetInitialized = true
	campaignKeeper.SetCampaign(sdkCtx, campaignMainnetInitialized)

	campaignEmptySupply := sample.Campaign(campaignEmptySupplyID)
	campaignEmptySupply.CoordinatorID = coordID
	campaignEmptySupply.TotalSupply = sdk.NewCoins()
	campaignKeeper.SetCampaign(sdkCtx, campaignEmptySupply)

	campaignIncorrectCoord := sample.Campaign(campaignIncorrectCoordID)
	campaignIncorrectCoord.CoordinatorID = coordID
	campaignKeeper.SetCampaign(sdkCtx, campaignIncorrectCoord)

	for _, tc := range []struct {
		name string
		msg  types.MsgInitializeMainnet
		err  error
	}{
		{
			name: "initialize mainnet",
			msg: types.MsgInitializeMainnet{
				CampaignID:     campaignID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(30),
				SourceURL:      sample.String(20),
				MainnetChainID: sample.GenesisChainID(),
			},
		},
		{
			name: "campaign not found",
			msg: types.MsgInitializeMainnet{
				CampaignID:     1000,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(30),
				SourceURL:      sample.String(20),
				MainnetChainID: sample.GenesisChainID(),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "mainnet already initialized",
			msg: types.MsgInitializeMainnet{
				CampaignID:     campaignMainnetInitializedID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(30),
				SourceURL:      sample.String(20),
				MainnetChainID: sample.GenesisChainID(),
			},
			err: types.ErrMainnetInitialized,
		},
		{
			name: "campaign empty supply",
			msg: types.MsgInitializeMainnet{
				CampaignID:     campaignEmptySupplyID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(30),
				SourceURL:      sample.String(20),
				MainnetChainID: sample.GenesisChainID(),
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "non-existent coordinator",
			msg: types.MsgInitializeMainnet{
				CampaignID:     campaignIncorrectCoordID,
				Coordinator:    sample.Address(),
				SourceHash:     sample.String(30),
				SourceURL:      sample.String(20),
				MainnetChainID: sample.GenesisChainID(),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator",
			msg: types.MsgInitializeMainnet{
				CampaignID:     campaignIncorrectCoordID,
				Coordinator:    coordAddrNoCampaign,
				SourceHash:     sample.String(30),
				SourceURL:      sample.String(20),
				MainnetChainID: sample.GenesisChainID(),
			},
			err: profiletypes.ErrCoordInvalid,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res, err := campaignSrv.InitializeMainnet(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			campaign, found := campaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)
			require.True(t, campaign.MainnetInitialized)
			require.EqualValues(t, res.MainnetID, campaign.MainnetID)

			// Chain is in launch module
			chain, found := launchKeeper.GetChain(sdkCtx, campaign.MainnetID)
			require.True(t, found)
			require.True(t, chain.HasCampaign)
			require.True(t, chain.IsMainnet)
			require.EqualValues(t, tc.msg.CampaignID, chain.CampaignID)

			// Mainnet ID is listed in campaign chains
			campaignChains, found := campaignKeeper.GetCampaignChains(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)
			require.Contains(t, campaignChains.Chains, campaign.MainnetID)
		})
	}
}
