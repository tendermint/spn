package keeper_test

import (
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
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)
	ctx := sdk.WrapSDKContext(sdkCtx)
	chainCreationFee := sample.Coins()

	// Create an invalid coordinator
	invalidCoordAddress := sample.Address()
	msgCreateInvalidCoordinator := sample.MsgCreateCoordinator(invalidCoordAddress)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateInvalidCoordinator)
	require.NoError(t, err)

	// Create a coordinator
	coordAddress := sample.Address()
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	resCoord, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordID := resCoord.CoordinatorID

	// Create a campaign
	msgCreateCampaign := sample.MsgCreateCampaign(coordAddress)
	resCampaign, err := ts.CampaignSrv.CreateCampaign(ctx, &msgCreateCampaign)
	require.NoError(t, err)
	campaignID := resCampaign.CampaignID

	// assign random sdk.Coins to `chainCreationFee` param and provide balance to coordinators to cover for
	// one chain creation
	initCreationFeeAndFundCoordAccounts(t, k, bankKeeper, sdkCtx, chainCreationFee, 4, coordAddress)

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
			require.EqualValues(t, coordID, chain.CoordinatorID)
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
