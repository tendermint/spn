package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestGetParams(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)

	t.Run("should allow default params set", func(t *testing.T) {
		params := types.DefaultParams()
		tk.MonitoringProviderKeeper.SetParams(ctx, params)
		require.EqualValues(t, params, tk.MonitoringProviderKeeper.GetParams(ctx))
		require.EqualValues(t, spntypes.ConsensusState{}, tk.MonitoringProviderKeeper.ConsumerConsensusState(ctx))
		require.EqualValues(t, types.DefaultConsumerChainID, tk.MonitoringProviderKeeper.ConsumerChainID(ctx))
		require.EqualValues(t, spntypes.DefaultUnbondingPeriod, tk.MonitoringProviderKeeper.ConsumerUnbondingPeriod(ctx))
		require.EqualValues(t, spntypes.DefaultRevisionHeight, tk.MonitoringProviderKeeper.ConsumerRevisionHeight(ctx))
	})

	t.Run("should allow params set", func(t *testing.T) {
		chainID := sample.GenesisChainID(r)
		cs := sample.ConsensusState(0)
		params := types.NewParams(
			1000,
			chainID,
			cs,
			10,
			20,
		)
		tk.MonitoringProviderKeeper.SetParams(ctx, params)
		require.EqualValues(t, params, tk.MonitoringProviderKeeper.GetParams(ctx))
		require.EqualValues(t, 1000, tk.MonitoringProviderKeeper.LastBlockHeight(ctx))
		require.EqualValues(t, cs, tk.MonitoringProviderKeeper.ConsumerConsensusState(ctx))
		require.EqualValues(t, chainID, tk.MonitoringProviderKeeper.ConsumerChainID(ctx))
		require.EqualValues(t, 10, tk.MonitoringProviderKeeper.ConsumerUnbondingPeriod(ctx))
		require.EqualValues(t, 20, tk.MonitoringProviderKeeper.ConsumerRevisionHeight(ctx))
	})
}
