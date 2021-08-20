package profile_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile"
)

func TestGenesis(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)

	original := sample.ProfileGenesisState()
	profile.InitGenesis(ctx, *keeper, original)
	got := profile.ExportGenesis(ctx, *keeper)

	// Compare lists
	require.Len(t, got.ValidatorList, len(original.ValidatorList))
	require.Subset(t, original.ValidatorList, got.ValidatorList)

	require.Len(t, got.CoordinatorList, len(original.CoordinatorList))
	require.Subset(t, original.CoordinatorList, got.CoordinatorList)

	require.Len(t, got.CoordinatorByAddressList, len(original.CoordinatorByAddressList))
	require.Subset(t, original.CoordinatorByAddressList, got.CoordinatorByAddressList)

	require.Equal(t, original.CoordinatorCount, got.CoordinatorCount)
}
