package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	campaigntypes "github.com/tendermint/spn/x/campaign/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgEditChain(t *testing.T) {
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)
	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.Address()
	coordAddress2 := sample.Address()
	coordNoExist := sample.Address()
	launchIDNoExist := uint64(1000)

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddress2)
	_, err = ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	// Create a chain
	msgCreateChain := sample.MsgCreateChain(coordAddress, "", false, 0)
	res, err := ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	launchID := res.LaunchID

	// create a campaign
	msgCreateCampaign := sample.MsgCreateCampaign(coordAddress)
	resCampaign, err := ts.CampaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)

	// create a chain with an existing campaign
	msgCreateChain = sample.MsgCreateChain(coordAddress, "", true, resCampaign.CampaignID)
	res, err = ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	launchIDHasCampaign := res.LaunchID

	// create a campaign
	msgCreateCampaign = sample.MsgCreateCampaign(coordAddress)
	resCampaign, err = ts.CampaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	validCampaignID := resCampaign.CampaignID

	// create a campaign from a different address
	msgCreateCampaign = sample.MsgCreateCampaign(coordAddress2)
	resCampaign, err = ts.CampaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	campaignDifferentCoordinator := resCampaign.CampaignID

	// Create a new chain for more tests
	msgCreateChain = sample.MsgCreateChain(coordAddress, "", false, 0)
	res, err = ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	launchID2 := res.LaunchID

	// create a new campaign and add a chainCampaigns entry to it
	msgCreateCampaign = sample.MsgCreateCampaign(coordAddress)
	resCampaign, err = ts.CampaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	campaignDuplicateChain := resCampaign.CampaignID

	err = tk.CampaignKeeper.AddChainToCampaign(sdkCtx, campaignDuplicateChain, launchID2)
	require.NoError(t, err)

	for _, tc := range []struct {
		name string
		msg  types.MsgEditChain
		err  error
	}{
		{
			name: "edit genesis chain ID",
			msg: sample.MsgEditChain(coordAddress, launchID,
				true,
				false,
				false,
				false,
				false,
				0,
				false,
			),
		},
		{
			name: "edit source",
			msg: sample.MsgEditChain(coordAddress, launchID,
				false,
				true,
				false,
				false,
				false,
				0,
				false,
			),
		},
		{
			name: "edit initial genesis with default genesis",
			msg: sample.MsgEditChain(coordAddress, launchID,
				false,
				false,
				true,
				false,
				false,
				0,
				false,
			),
		},
		{
			name: "edit initial genesis with genesis url",
			msg: sample.MsgEditChain(coordAddress, launchID,
				false,
				false,
				true,
				true,
				false,
				0,
				false,
			),
		},
		{
			name: "edit source and initial genesis",
			msg: sample.MsgEditChain(coordAddress, launchID,
				false,
				true,
				true,
				true,
				false,
				0,
				false,
			),
		},
		{
			name: "set campaign ID",
			msg: sample.MsgEditChain(coordAddress, launchID,
				false,
				false,
				false,
				false,
				true,
				validCampaignID,
				false,
			),
		},
		{
			name: "edit metadata",
			msg: sample.MsgEditChain(coordAddress, launchID,
				false,
				false,
				false,
				false,
				false,
				0,
				true,
			),
		},
		{
			name: "non existent launch id",
			msg: sample.MsgEditChain(coordAddress, launchIDNoExist,
				false,
				true,
				false,
				false,
				false,
				0,
				false,
			),
			err: types.ErrChainNotFound,
		},
		{
			name: "non existent coordinator",
			msg: sample.MsgEditChain(coordNoExist, launchID,
				false,
				true,
				false,
				false,
				false,
				0,
				false,
			),
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator",
			msg: sample.MsgEditChain(coordAddress2, launchID,
				false,
				true,
				false,
				false,
				false,
				0,
				false,
			),
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "chain already has campaign",
			msg: sample.MsgEditChain(coordAddress, launchIDHasCampaign,
				false,
				false,
				false,
				false,
				true,
				0,
				false,
			),
			err: types.ErrChainHasCampaign,
		},
		{
			name: "campaign does not exist",
			msg: sample.MsgEditChain(coordAddress, launchID2,
				false,
				false,
				false,
				false,
				true,
				999,
				false,
			),
			err: campaigntypes.ErrCampaignNotFound,
		},
		{
			name: "campaign has a different coordinator",
			msg: sample.MsgEditChain(coordAddress, launchID2,
				false,
				false,
				false,
				false,
				true,
				campaignDifferentCoordinator,
				false,
			),
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "campaign chain entry is duplicated",
			msg: sample.MsgEditChain(coordAddress, launchID2,
				false,
				false,
				false,
				false,
				true,
				campaignDuplicateChain,
				false,
			),
			err: types.ErrAddChainToCampaign,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Fetch the previous state of the chain to perform checks
			var previousChain types.Chain
			var found bool
			if tc.err == nil {
				previousChain, found = tk.LaunchKeeper.GetChain(sdkCtx, tc.msg.LaunchID)
				require.True(t, found)
			}

			// Send the message
			_, err := ts.LaunchSrv.EditChain(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			// The chain must continue to exist in the store
			chain, found := tk.LaunchKeeper.GetChain(sdkCtx, tc.msg.LaunchID)
			require.True(t, found)

			// Unchanged values
			require.EqualValues(t, previousChain.CoordinatorID, chain.CoordinatorID)
			require.EqualValues(t, previousChain.CreatedAt, chain.CreatedAt)
			require.EqualValues(t, previousChain.LaunchTimestamp, chain.LaunchTimestamp)
			require.EqualValues(t, previousChain.LaunchTriggered, chain.LaunchTriggered)

			// Compare changed values
			if tc.msg.GenesisChainID != "" {
				require.EqualValues(t, tc.msg.GenesisChainID, chain.GenesisChainID)
			} else {
				require.EqualValues(t, previousChain.GenesisChainID, chain.GenesisChainID)
			}
			if tc.msg.SourceURL != "" {
				require.EqualValues(t, tc.msg.SourceURL, chain.SourceURL)
				require.EqualValues(t, tc.msg.SourceHash, chain.SourceHash)
			} else {
				require.EqualValues(t, previousChain.SourceURL, chain.SourceURL)
				require.EqualValues(t, previousChain.SourceHash, chain.SourceHash)
			}

			if tc.msg.InitialGenesis != nil {
				require.EqualValues(t, *tc.msg.InitialGenesis, chain.InitialGenesis)
			} else {
				require.EqualValues(t, previousChain.InitialGenesis, chain.InitialGenesis)
			}

			if len(tc.msg.Metadata) > 0 {
				require.EqualValues(t, tc.msg.Metadata, chain.Metadata)
			} else {
				require.EqualValues(t, previousChain.Metadata, chain.Metadata)
			}

			if tc.msg.SetCampaignID {
				require.True(t, chain.HasCampaign)
				require.EqualValues(t, tc.msg.CampaignID, chain.CampaignID)
				// ensure campaign exist
				_, found := tk.CampaignKeeper.GetCampaign(sdkCtx, chain.CampaignID)
				require.True(t, found)
				// ensure campaign chains exist
				campaignChains, found := tk.CampaignKeeper.GetCampaignChains(sdkCtx, chain.CampaignID)
				require.True(t, found)

				// check that the chain launch ID is in the campaign chains
				found = false
				for _, chainID := range campaignChains.Chains {
					if chainID == chain.LaunchID {
						found = true
						break
					}
				}

				require.True(t, found)
			}
		})
	}
}
