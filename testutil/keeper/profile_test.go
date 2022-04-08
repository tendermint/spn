package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
)

func TestTestMsgServers_CreateCoordinator(t *testing.T) {
	sdkCtx, tk, tm := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	ctx := sdk.WrapSDKContext(sdkCtx)

	id, addr := tm.CreateCoordinator(ctx, r)
	coord, found := tk.ProfileKeeper.GetCoordinator(sdkCtx, id)
	require.True(t, found)
	require.Equal(t, id, coord.CoordinatorID)
	require.Equal(t, addr.String(), coord.Address)
	coordByAddr, err := tk.ProfileKeeper.GetCoordinatorByAddress(sdkCtx, addr.String())
	require.NoError(t, err)
	require.Equal(t, id, coordByAddr.CoordinatorID)
	require.Equal(t, addr.String(), coordByAddr.Address)
}
