package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgInitializeMainnet(t *testing.T) {
	var (
		coordID                      uint64
		campaignID                   uint64 = 0
		campaignMainnetInitializedID uint64 = 1
		campaignIncorrectCoordID     uint64 = 2
		campaignEmptySupplyID        uint64 = 3
		coordAddr                           = sample.Address(r)
		coordAddrNoCampaign                 = sample.Address(r)

		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
	)

	t.Run("should allow creation of coordinators", func(t *testing.T) {
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddr,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
		coordID = res.CoordinatorID
		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddrNoCampaign,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
	})

	campaign := sample.Campaign(r, campaignID)
	campaign.CoordinatorID = coordID
	tk.CampaignKeeper.SetCampaign(sdkCtx, campaign)

	campaignMainnetInitialized := sample.Campaign(r, campaignMainnetInitializedID)
	campaignMainnetInitialized.CoordinatorID = coordID
	campaignMainnetInitialized.MainnetInitialized = true
	tk.CampaignKeeper.SetCampaign(sdkCtx, campaignMainnetInitialized)

	campaignEmptySupply := sample.Campaign(r, campaignEmptySupplyID)
	campaignEmptySupply.CoordinatorID = coordID
	campaignEmptySupply.TotalSupply = sdk.NewCoins()
	tk.CampaignKeeper.SetCampaign(sdkCtx, campaignEmptySupply)

	campaignIncorrectCoord := sample.Campaign(r, campaignIncorrectCoordID)
	campaignIncorrectCoord.CoordinatorID = coordID
	tk.CampaignKeeper.SetCampaign(sdkCtx, campaignIncorrectCoord)

	for _, tc := range []struct {
		name string
		msg  types.MsgInitializeMainnet
		err  error
	}{
		{
			name: "should allow initialize mainnet",
			msg: types.MsgInitializeMainnet{
				CampaignID:     campaignID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
		},
		{
			name: "should fail if campaign not found",
			msg: types.MsgInitializeMainnet{
				CampaignID:     1000,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "should fail if mainnet already initialized",
			msg: types.MsgInitializeMainnet{
				CampaignID:     campaignMainnetInitializedID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrMainnetInitialized,
		},
		{
			name: "should fail if campaign has empty supply",
			msg: types.MsgInitializeMainnet{
				CampaignID:     campaignEmptySupplyID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "should fail with non-existent coordinator",
			msg: types.MsgInitializeMainnet{
				CampaignID:     campaignIncorrectCoordID,
				Coordinator:    sample.Address(r),
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should fail with invalid coordinator",
			msg: types.MsgInitializeMainnet{
				CampaignID:     campaignIncorrectCoordID,
				Coordinator:    coordAddrNoCampaign,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: profiletypes.ErrCoordInvalid,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ts.CampaignSrv.InitializeMainnet(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			campaign, found := tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)
			require.True(t, campaign.MainnetInitialized)
			require.EqualValues(t, res.MainnetID, campaign.MainnetID)

			// Chain is in launch module
			chain, found := tk.LaunchKeeper.GetChain(sdkCtx, campaign.MainnetID)
			require.True(t, found)
			require.True(t, chain.HasCampaign)
			require.True(t, chain.IsMainnet)
			require.EqualValues(t, tc.msg.CampaignID, chain.CampaignID)

			// Mainnet ID is listed in campaign chains
			campaignChains, found := tk.CampaignKeeper.GetCampaignChains(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)
			require.Contains(t, campaignChains.Chains, campaign.MainnetID)
		})
	}
}
