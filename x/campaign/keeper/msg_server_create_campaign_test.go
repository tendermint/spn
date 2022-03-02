package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func initCreationFeeAndCoordAccounts(
	t *testing.T,
	keeper *keeper.Keeper,
	bk bankkeeper.Keeper,
	sdkCtx sdk.Context,
	coins sdk.Coins,
	addrs ...string,
) {
	// set fee param to `coins`
	params := keeper.GetParams(sdkCtx)
	params.CampaignCreationFee = coins
	keeper.SetParams(sdkCtx, params)

	// add `coins` to balance of each coordinator address
	for _, addr := range addrs {
		accAddr, err := sdk.AccAddressFromBech32(addr)
		require.NoError(t, err)
		err = bk.MintCoins(sdkCtx, types.ModuleName, coins)
		require.NoError(t, err)
		err = bk.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, accAddr, coins)
		require.NoError(t, err)
	}
}

func TestMsgCreateCampaign(t *testing.T) {
	var (
		coordAddr1                                                        = sample.Address()
		coordAddr2                                                        = sample.Address()
		campaignKeeper, _, _, bankKeeper, campaignSrv, profileSrv, sdkCtx = setupMsgServer(t)
		ctx                                                               = sdk.WrapSDKContext(sdkCtx)
	)

	// Create coordinators
	coordMap := make(map[string]uint64)
	res, err := profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr1,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	coordMap[coordAddr1] = res.CoordinatorID
	res, err = profileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordAddr2,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	coordMap[coordAddr2] = res.CoordinatorID

	// assign random sdk.Coins to `campaignCreationFee` param and provide balance to coordinators to cover for
	// one campaign creation
	initCreationFeeAndCoordAccounts(t, campaignKeeper, bankKeeper, sdkCtx, sample.Coins(), coordAddr1, coordAddr2)

	for _, tc := range []struct {
		name       string
		msg        types.MsgCreateCampaign
		expectedID uint64
		err        error
	}{
		{
			name: "create a campaign 1",
			msg: types.MsgCreateCampaign{
				CampaignName: sample.CampaignName(),
				Coordinator:  coordAddr1,
				TotalSupply:  sample.TotalSupply(),
				Metadata:     sample.Metadata(20),
			},
			expectedID: uint64(0),
		},
		{
			name: "create a campaign from a different coordinator",
			msg: types.MsgCreateCampaign{
				CampaignName: sample.CampaignName(),
				Coordinator:  coordAddr2,
				TotalSupply:  sample.TotalSupply(),
				Metadata:     sample.Metadata(20),
			},
			expectedID: uint64(1),
		},
		{
			name: "create a campaign from a non existing coordinator",
			msg: types.MsgCreateCampaign{
				CampaignName: sample.CampaignName(),
				Coordinator:  sample.Address(),
				TotalSupply:  sample.TotalSupply(),
				Metadata:     sample.Metadata(20),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "create a campaign with an invalid token supply",
			msg: types.MsgCreateCampaign{
				CampaignName: sample.CampaignName(),
				Coordinator:  coordAddr1,
				TotalSupply:  sample.CoinsWithRange(10, 20),
				Metadata:     sample.Metadata(20),
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "insufficient balance to cover creation fee",
			msg: types.MsgCreateCampaign{
				CampaignName: sample.CampaignName(),
				Coordinator:  coordAddr1,
				TotalSupply:  sample.TotalSupply(),
				Metadata:     sample.Metadata(20),
			},
			err: sdkerrors.ErrInsufficientFunds,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := campaignSrv.CreateCampaign(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectedID, got.CampaignID)
			campaign, found := campaignKeeper.GetCampaign(sdkCtx, got.CampaignID)
			require.True(t, found)
			require.EqualValues(t, got.CampaignID, campaign.CampaignID)
			require.EqualValues(t, tc.msg.CampaignName, campaign.CampaignName)
			require.EqualValues(t, coordMap[tc.msg.Coordinator], campaign.CoordinatorID)
			require.False(t, campaign.MainnetInitialized)
			require.True(t, tc.msg.TotalSupply.IsEqual(campaign.TotalSupply))
			require.EqualValues(t, types.EmptyShares(), campaign.AllocatedShares)
			require.EqualValues(t, types.EmptyShares(), campaign.TotalShares)

			// dynamic share is always disabled
			require.EqualValues(t, false, campaign.DynamicShares)

			// Empty list of campaign chains
			campaignChains, found := campaignKeeper.GetCampaignChains(sdkCtx, got.CampaignID)
			require.True(t, found)
			require.EqualValues(t, got.CampaignID, campaignChains.CampaignID)
			require.Empty(t, campaignChains.Chains)
		})
	}
}
