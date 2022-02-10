package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgCreateChain(t *testing.T) {
	k, _, campaignKeeper, srv, profileSrv, campaignSrv, sdkCtx := setupMsgServer(t)
	ctx := sdk.WrapSDKContext(sdkCtx)

	// Create an invalid coordinator
	invalidCoordAddress := sample.Address()
	msgCreateInvalidCoordinator := sample.MsgCreateCoordinator(invalidCoordAddress)
	_, err := profileSrv.CreateCoordinator(ctx, &msgCreateInvalidCoordinator)
	require.NoError(t, err)

	// Create a coordinator
	coordAddress := sample.Address()
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	resCoord, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordID := resCoord.CoordinatorID

	// Create coordinator and disable
	disableCoordAddress := sample.Address()
	msgCreateCoordinator = sample.MsgCreateCoordinator(disableCoordAddress)
	_, err = profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	msgDisableCoord := sample.MsgDisableCoordinator(disableCoordAddress)
	_, err = profileSrv.DisableCoordinator(ctx, &msgDisableCoord)
	require.NoError(t, err)

	// Create a campaign with disabled
	msgCreateCampaign := sample.MsgCreateCampaign(disableCoordAddress)
	_, err = campaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.ErrorIs(t, err, profiletypes.ErrCoordAddressNotFound)

	// Create a campaign
	msgCreateCampaign = sample.MsgCreateCampaign(coordAddress)
	resCampaign, err := campaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	campaignID := resCampaign.CampaignID

	for _, tc := range []struct {
		name          string
		msg           types.MsgCreateChain
		wantedChainID uint64
		err           error
	}{
		{
			name:          "valid message",
			msg:           sample.MsgCreateChain(coordAddress, "", false, campaignID),
			wantedChainID: 0,
		},
		{
			name:          "creates a unique chain ID",
			msg:           sample.MsgCreateChain(coordAddress, "", false, campaignID),
			wantedChainID: 1,
		},
		{
			name:          "valid message with genesis url",
			msg:           sample.MsgCreateChain(coordAddress, "foo.com", false, campaignID),
			wantedChainID: 2,
		},
		{
			name:          "creates message with campaign",
			msg:           sample.MsgCreateChain(coordAddress, "", true, campaignID),
			wantedChainID: 3,
		},
		{
			name: "coordinator doesn't exist for the chain",
			msg:  sample.MsgCreateChain(sample.Address(), "", false, 0),
			err:  profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid campaign id",
			msg:  sample.MsgCreateChain(coordAddress, "", true, 1000),
			err:  types.ErrCreateChainFail,
		},
		{
			name: "invalid coordinator address",
			msg:  sample.MsgCreateChain(invalidCoordAddress, "", true, 1000),
			err:  types.ErrCreateChainFail,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := srv.CreateChain(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			require.EqualValues(t, tc.wantedChainID, got.LaunchID)

			// The chain must exist in the store
			chain, found := k.GetChain(sdkCtx, got.LaunchID)
			require.True(t, found)
			require.EqualValues(t, coordID, chain.CoordinatorID)
			require.EqualValues(t, got.LaunchID, chain.LaunchID)
			require.EqualValues(t, tc.msg.GenesisChainID, chain.GenesisChainID)
			require.EqualValues(t, tc.msg.SourceURL, chain.SourceURL)
			require.EqualValues(t, tc.msg.SourceHash, chain.SourceHash)

			// Compare initial genesis
			if tc.msg.GenesisURL == "" {
				require.Equal(t, types.NewDefaultInitialGenesis(), chain.InitialGenesis)
			} else {
				require.Equal(
					t,
					types.NewGenesisURL(tc.msg.GenesisURL, tc.msg.GenesisHash),
					chain.InitialGenesis,
				)
			}

			// Chain created from MsgCreateChain is never a mainnet
			require.False(t, chain.IsMainnet)

			require.Equal(t, tc.msg.HasCampaign, chain.HasCampaign)

			if tc.msg.HasCampaign {
				require.Equal(t, tc.msg.CampaignID, chain.CampaignID)
				campaignChains, found := campaignKeeper.GetCampaignChains(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)
				require.Contains(t, campaignChains.Chains, chain.LaunchID)
			}
		})
	}
}
