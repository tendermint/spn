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
	k, ctx := testkeeper.MonitoringpKeeper(t)
	params := types.DefaultParams()
	k.SetParams(ctx, params)
	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, ibctypes.ConsensusState{}, k.ConsumerConsensusState(ctx))

	cs := sample.ConsensusState(0)
	params = types.NewParams(cs)
	k.SetParams(ctx, params)
	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, cs, k.ConsumerConsensusState(ctx))
}