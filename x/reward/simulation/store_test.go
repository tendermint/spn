package simulation_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	rewardsimulation "github.com/tendermint/spn/x/reward/simulation"
	"github.com/tendermint/spn/x/reward/types"
)

func TestFindRandomChainWithCoordBalance(t *testing.T) {
	var (
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)

		r              = sample.Rand()
		ctx            = sdk.WrapSDKContext(sdkCtx)
		msgCreateCoord = sample.MsgCreateCoordinator(sample.Address(r))
	)

	t.Run("no chains", func(t *testing.T) {
		_, found := rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, true, false, sdk.NewCoins())
		require.False(t, found)
		_, found = rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, true, true, sample.Coins(r))
		require.False(t, found)
	})

	// Create coordinator
	res, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoord)
	require.NoError(t, err)

	coord, found := tk.ProfileKeeper.GetCoordinator(sdkCtx, res.CoordinatorID)
	require.True(t, found)

	// set balance
	coordBalance := sample.Coins(r)
	tk.Mint(sdkCtx, coord.Address, coordBalance)

	t.Run("chain without coordinator", func(t *testing.T) {
		tk.LaunchKeeper.AppendChain(sdkCtx, launchtypes.Chain{
			CoordinatorID:   1000,
			LaunchTriggered: true,
		})
		_, found := rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, true, false, sdk.NewCoins())
		require.False(t, found)
	})

	t.Run("chain with launch triggered", func(t *testing.T) {
		tk.LaunchKeeper.AppendChain(sdkCtx, launchtypes.Chain{
			CoordinatorID:   res.CoordinatorID,
			LaunchTriggered: true,
		})
		_, found := rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, true, false, sdk.NewCoins())
		require.False(t, found)
		_, found = rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, true, true, coordBalance)
		require.False(t, found)
		_, found = rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, false, false, sdk.NewCoins())
		require.False(t, found)
		_, found = rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, false, true, coordBalance)
		require.False(t, found)
	})

	t.Run("valid chains", func(t *testing.T) {
		norewardPoolID := tk.LaunchKeeper.AppendChain(sdkCtx, launchtypes.Chain{
			CoordinatorID:   res.CoordinatorID,
			LaunchTriggered: false,
		})

		hasRewardPoolID := tk.LaunchKeeper.AppendChain(sdkCtx, launchtypes.Chain{
			CoordinatorID:   res.CoordinatorID,
			LaunchTriggered: false,
		})

		tk.RewardKeeper.SetRewardPool(sdkCtx, types.RewardPool{
			LaunchID: hasRewardPoolID,
			Provider: coord.Address,
		})

		got, found := rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, true, false, sdk.NewCoins())
		require.True(t, found)
		require.Equal(t, res.CoordinatorID, got.CoordinatorID)
		require.Equal(t, hasRewardPoolID, got.LaunchID)
		got, found = rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, true, true, coordBalance)
		require.True(t, found)
		require.Equal(t, res.CoordinatorID, got.CoordinatorID)
		require.Equal(t, hasRewardPoolID, got.LaunchID)

		got, found = rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, false, false, sdk.NewCoins())
		require.True(t, found)
		require.Equal(t, res.CoordinatorID, got.CoordinatorID)
		require.Equal(t, norewardPoolID, got.LaunchID)
		got, found = rewardsimulation.FindRandomChainWithCoordBalance(r, sdkCtx, *tk.RewardKeeper, tk.BankKeeper, false, true, coordBalance)
		require.True(t, found)
		require.Equal(t, res.CoordinatorID, got.CoordinatorID)
		require.Equal(t, norewardPoolID, got.LaunchID)
	})
}
