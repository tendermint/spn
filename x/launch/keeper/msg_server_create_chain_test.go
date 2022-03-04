package keeper_test

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func initCreationFeeAndFundCoordAccounts(
	t *testing.T,
	keeper *keeper.Keeper,
	bk bankkeeper.Keeper,
	sdkCtx sdk.Context,
	fee sdk.Coins,
	numCreations int64,
	addrs ...string,
) {
	// set fee param to `coins`
	params := keeper.GetParams(sdkCtx)
	params.ChainCreationFee = fee
	keeper.SetParams(sdkCtx, params)

	coins := sdk.NewCoins()
	for _, coin := range fee {
		coin.Amount = coin.Amount.MulRaw(numCreations)
		coins = coins.Add(coin)
	}

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

func TestMsgCreateChain(t *testing.T) {
	var (
		coordAddr1       = sample.Address()
		coordAddr2       = sample.Address()
		coordAddr3       = sample.Address()
		coordAddr4       = sample.Address()
		coordAddr5       = sample.Address()
		sdkCtx, tk, ts   = testkeeper.NewTestSetup(t)
		ctx              = sdk.WrapSDKContext(sdkCtx)
		chainCreationFee = sample.Coins()
	)

	// Create an invalid coordinator
	invalidCoordAddress := sample.Address()
	msgCreateInvalidCoordinator := sample.MsgCreateCoordinator(invalidCoordAddress)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateInvalidCoordinator)
	require.NoError(t, err)

	// Create coordinators
	coordMap := make(map[string]uint64)
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddr1)
	resCoord, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordMap[coordAddr1] = resCoord.CoordinatorID
	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddr2)
	resCoord, err = ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordMap[coordAddr2] = resCoord.CoordinatorID
	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddr3)
	resCoord, err = ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordMap[coordAddr3] = resCoord.CoordinatorID
	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddr4)
	resCoord, err = ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordMap[coordAddr4] = resCoord.CoordinatorID
	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddr5)
	resCoord, err = ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordMap[coordAddr5] = resCoord.CoordinatorID

	// Create a campaign for each valid coordinator
	campMap := make(map[string]uint64)
	msgCreateCampaign := sample.MsgCreateCampaign(coordAddr1)
	resCampaign, err := ts.CampaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	campMap[coordAddr1] = resCampaign.CampaignID
	msgCreateCampaign = sample.MsgCreateCampaign(coordAddr2)
	resCampaign, err = ts.CampaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	campMap[coordAddr2] = resCampaign.CampaignID
	msgCreateCampaign = sample.MsgCreateCampaign(coordAddr3)
	resCampaign, err = ts.CampaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	campMap[coordAddr3] = resCampaign.CampaignID
	msgCreateCampaign = sample.MsgCreateCampaign(coordAddr4)
	resCampaign, err = ts.CampaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	campMap[coordAddr4] = resCampaign.CampaignID
	msgCreateCampaign = sample.MsgCreateCampaign(coordAddr5)
	resCampaign, err = ts.CampaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	campMap[coordAddr5] = resCampaign.CampaignID

	// assign random sdk.Coins to `chainCreationFee` param and provide balance to coordinators
	// coordAddr5 is not funded
	initCreationFeeAndFundCoordAccounts(
		t,
		tk.LaunchKeeper,
		tk.BankKeeper,
		sdkCtx,
		chainCreationFee,
		1,
		coordAddr1,
		coordAddr2,
		coordAddr3,
		coordAddr4,
	)

	for _, tc := range []struct {
		name          string
		msg           types.MsgCreateChain
		wantedChainID uint64
		err           error
	}{
		{
			name:          "valid message",
			msg:           sample.MsgCreateChain(coordAddr1, "", false, campMap[coordAddr1]),
			wantedChainID: 0,
		},
		{
			name:          "creates a unique chain ID",
			msg:           sample.MsgCreateChain(coordAddr2, "", false, campMap[coordAddr2]),
			wantedChainID: 1,
		},
		{
			name:          "valid message with genesis url",
			msg:           sample.MsgCreateChain(coordAddr3, "foo.com", false, campMap[coordAddr3]),
			wantedChainID: 2,
		},
		{
			name:          "creates message with campaign",
			msg:           sample.MsgCreateChain(coordAddr4, "", true, campMap[coordAddr4]),
			wantedChainID: 3,
		},
		{
			name: "coordinator doesn't exist for the chain",
			msg:  sample.MsgCreateChain(sample.Address(), "", false, 0),
			err:  profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid campaign id",
			msg:  sample.MsgCreateChain(coordAddr1, "", true, 1000),
			err:  types.ErrCreateChainFail,
		},
		{
			name: "invalid coordinator address",
			msg:  sample.MsgCreateChain(invalidCoordAddress, "", true, 1000),
			err:  types.ErrCreateChainFail,
		},
		{
			name: "insufficient balance to cover creation fee",
			msg:  sample.MsgCreateChain(coordAddr5, "", false, campMap[coordAddr5]),
			err:  sdkerrors.ErrInsufficientFunds,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// get account initial balance
			accAddr, err := sdk.AccAddressFromBech32(tc.msg.Coordinator)
			require.NoError(t, err)
			preBalance := tk.BankKeeper.SpendableCoins(sdkCtx, accAddr)

			got, err := ts.LaunchSrv.CreateChain(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			require.EqualValues(t, tc.wantedChainID, got.LaunchID)

			// The chain must exist in the store
			chain, found := tk.LaunchKeeper.GetChain(sdkCtx, got.LaunchID)
			require.True(t, found)
			require.EqualValues(t, coordMap[tc.msg.Coordinator], chain.CoordinatorID)
			require.EqualValues(t, got.LaunchID, chain.LaunchID)
			require.EqualValues(t, tc.msg.GenesisChainID, chain.GenesisChainID)
			require.EqualValues(t, tc.msg.SourceURL, chain.SourceURL)
			require.EqualValues(t, tc.msg.SourceHash, chain.SourceHash)
			require.EqualValues(t, tc.msg.Metadata, chain.Metadata)

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
				campaignChains, found := tk.CampaignKeeper.GetCampaignChains(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)
				require.Contains(t, campaignChains.Chains, chain.LaunchID)
			}

			// check fee deduction
			postBalance := tk.BankKeeper.SpendableCoins(sdkCtx, accAddr)
			require.True(t, preBalance.Sub(chainCreationFee).IsEqual(postBalance))
		})
	}
}
