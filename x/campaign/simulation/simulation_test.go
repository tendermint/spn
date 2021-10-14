package simulation_test

import (
	"testing"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	campaignsim "github.com/tendermint/spn/x/campaign/simulation"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestGetCoordSimAccount(t *testing.T) {
	pk, ctx := testkeeper.Profile(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	// No coordinator account
	_, _, found := campaignsim.GetCoordSimAccount(r, ctx, pk, accs)
	require.False(t, found)

	// Create a coordinator from an account
	acc, accPos := simtypes.RandomAcc(r, accs)
	coordID := pk.AppendCoordinator(ctx, profiletypes.Coordinator{
		Address:     acc.Address.String(),
		Description: sample.CoordinatorDescription(),
	})
	acc, foundCoordID, found := campaignsim.GetCoordSimAccount(r, ctx, pk, accs)
	require.True(t, found)
	require.EqualValues(t, accs[accPos], acc)
	require.EqualValues(t, coordID, foundCoordID)
}
