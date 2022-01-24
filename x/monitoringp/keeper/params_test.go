package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/pkg/ibctypes"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestGetParams(t *testing.T) {
	k, _, ctx := testkeeper.MonitoringpKeeper(t)
	params := types.DefaultParams()
	k.SetParams(ctx, params)
	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, ibctypes.ConsensusState{}, k.ConsumerConsensusState(ctx))
	require.EqualValues(t, types.DefautConsumerChainID, k.ConsumerChainID(ctx))

	chainID := sample.GenesisChainID()
	cs := sample.ConsensusState(0)
	params = types.NewParams(chainID, cs)
	k.SetParams(ctx, params)
	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, cs, k.ConsumerConsensusState(ctx))
	require.EqualValues(t, chainID, k.ConsumerChainID(ctx))
}
