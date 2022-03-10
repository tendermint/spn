package participation_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation"
)

func TestGenesis(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	genesisState := sample.ParticipationGenesisState()
	participation.InitGenesis(ctx, *tk.ParticipationKeeper, genesisState)
	got := participation.ExportGenesis(ctx, *tk.ParticipationKeeper)

	require.Equal(t, genesisState.Params, got.Params)

	require.ElementsMatch(t, genesisState.UsedAllocationsList, got.UsedAllocationsList)
	require.ElementsMatch(t, genesisState.AuctionUsedAllocationsList, got.AuctionUsedAllocationsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
