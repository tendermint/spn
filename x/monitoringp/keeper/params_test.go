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
	k, _, ctx := testkeeper.MonitoringpKeeper(t)
	params := types.DefaultParams()
	k.SetParams(ctx, params)
	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, spntypes.ConsensusState{}, k.ConsumerConsensusState(ctx))
	require.EqualValues(t, types.DefautConsumerChainID, k.ConsumerChainID(ctx))
	require.EqualValues(t, false, k.DebugMode(ctx))

	chainID := sample.GenesisChainID()
	cs := sample.ConsensusState(0)
	params = types.NewParams(1000, chainID, cs, true)
	k.SetParams(ctx, params)
	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, cs, k.ConsumerConsensusState(ctx))
	require.EqualValues(t, chainID, k.ConsumerChainID(ctx))
	require.EqualValues(t, true, k.DebugMode(ctx))
	require.EqualValues(t, 1000, k.LastBlockHeight(ctx))
}
