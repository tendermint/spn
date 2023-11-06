package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	projecttypes "github.com/tendermint/spn/x/project/types"
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
	// using `project` module account for minting as `launch` does not have one
	for _, addr := range addrs {
		accAddr, err := sdk.AccAddressFromBech32(addr)
		require.NoError(t, err)
		err = bk.MintCoins(sdkCtx, projecttypes.ModuleName, coins)
		require.NoError(t, err)
		err = bk.SendCoinsFromModuleToAccount(sdkCtx, projecttypes.ModuleName, accAddr, coins)
		require.NoError(t, err)
	}
}

func TestMsgCreateChain(t *testing.T) {
	var (
		coordAddrs       = make([]string, 5)
		coordMap         = make(map[string]uint64)
		prjtMap          = make(map[string]uint64)
		sdkCtx, tk, ts   = testkeeper.NewTestSetup(t)
		ctx              = sdk.WrapSDKContext(sdkCtx)
		chainCreationFee = sample.Coins(r)
	)

	// Create an invalid coordinator
	_, invalidCoordAddr := ts.CreateCoordinator(ctx, r)
	invalidCoordAddress := invalidCoordAddr.String()

	// Create coordinators
	for i := range coordAddrs {
		addr := sample.Address(r)
		coordAddrs[i] = addr
		coordMap[addr], _ = ts.CreateCoordinator(ctx, r, addr)
	}

	// Create a project for each valid coordinator
	for i := range coordAddrs {
		addr := coordAddrs[i]
		prjtMap[addr] = ts.CreateProject(ctx, r, addr)
	}

	// assign random sdk.Coins to `chainCreationFee` param and provide balance to coordinators
	// coordAddrs[4] is not funded
	initCreationFeeAndFundCoordAccounts(t, tk.LaunchKeeper, tk.BankKeeper, sdkCtx, chainCreationFee, 1, coordAddrs[:4]...)

	// create message with an invalid metadata length
	msgCreateChainInvalidMetadata := sample.MsgCreateChain(
		r,
		coordAddrs[0],
		"",
		false,
		prjtMap[coordAddrs[0]],
	)
	maxMetadataLength := tk.LaunchKeeper.MaxMetadataLength(sdkCtx)
	msgCreateChainInvalidMetadata.Metadata = sample.Metadata(r, maxMetadataLength+1)

	for _, tc := range []struct {
		name          string
		msg           types.MsgCreateChain
		wantedChainID uint64
		err           error
	}{
		{
			name:          "should create a chain",
			msg:           sample.MsgCreateChain(r, coordAddrs[0], "", false, prjtMap[coordAddrs[0]]),
			wantedChainID: 0,
		},
		{
			name:          "should allow creating a chain with a unique chain ID",
			msg:           sample.MsgCreateChain(r, coordAddrs[1], "", false, prjtMap[coordAddrs[1]]),
			wantedChainID: 1,
		},
		{
			name:          "should allow creating a chain with genesis url",
			msg:           sample.MsgCreateChain(r, coordAddrs[2], "foo.com", false, prjtMap[coordAddrs[2]]),
			wantedChainID: 2,
		},
		{
			name:          "should allow creating a chain with project",
			msg:           sample.MsgCreateChain(r, coordAddrs[3], "", true, prjtMap[coordAddrs[3]]),
			wantedChainID: 3,
		},
		{
			name: "should prevent creating a chain where coordinator doesn't exist for the chain",
			msg:  sample.MsgCreateChain(r, sample.Address(r), "", false, 0),
			err:  profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should prevent creating a chain with invalid project id",
			msg:  sample.MsgCreateChain(r, coordAddrs[0], "", true, 1000),
			err:  types.ErrCreateChainFail,
		},
		{
			name: "should prevent creating a chain with invalid coordinator address",
			msg:  sample.MsgCreateChain(r, invalidCoordAddress, "", true, 1000),
			err:  types.ErrCreateChainFail,
		},
		{
			name: "should prevent creating a chain with insufficient balance to cover creation fee",
			msg:  sample.MsgCreateChain(r, coordAddrs[4], "", false, prjtMap[coordAddrs[4]]),
			err:  types.ErrFundCommunityPool,
		},
		{
			name: "should prevent a chain with invalid metadata length",
			msg:  msgCreateChainInvalidMetadata,
			err:  types.ErrInvalidMetadataLength,
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
			require.EqualValues(t, tc.msg.InitialGenesis, chain.InitialGenesis)

			// Chain created from MsgCreateChain is never a mainnet
			require.False(t, chain.IsMainnet)

			require.Equal(t, tc.msg.HasProject, chain.HasProject)

			if tc.msg.HasProject {
				require.Equal(t, tc.msg.ProjectID, chain.ProjectID)
				projectChains, found := tk.ProjectKeeper.GetProjectChains(sdkCtx, tc.msg.ProjectID)
				require.True(t, found)
				require.Contains(t, projectChains.Chains, chain.LaunchID)
			}

			// check fee deduction
			postBalance := tk.BankKeeper.SpendableCoins(sdkCtx, accAddr)
			require.True(t, preBalance.Sub(chainCreationFee...).IsEqual(postBalance))
		})
	}
}
