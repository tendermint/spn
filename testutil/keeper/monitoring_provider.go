package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibckeeper "github.com/cosmos/ibc-go/v2/modules/core/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/tendermint/spn/x/monitoringp/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
)

// MonitoringpKeeper returns a keeper of the monitoring provider module for testing purpose
func MonitoringpKeeper(t testing.TB) (*keeper.Keeper, *ibckeeper.Keeper, sdk.Context) {
	return MonitoringpKeeperWithIBCMock(t, []Connection{}, []Channel{})
}

// MonitoringpKeeperWithIBCMock returns a keeper of the monitoring provider module for testing purpose
func MonitoringpKeeperWithIBCMock(
	t testing.TB,
	connectionMock []Connection,
	channelMock []Channel,
) (*keeper.Keeper, *ibckeeper.Keeper, sdk.Context) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	stakingkeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	ibcKeeper := initializer.IBC(paramKeeper, stakingkeeper, *capabilityKeeper)
	monitoringKeeper := initializer.Monitoringp(
		*ibcKeeper,
		*capabilityKeeper,
		paramKeeper,
		connectionMock,
		channelMock,
	)
	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	monitoringKeeper.SetParams(ctx, types.DefaultParams())
	setIBCDefaultParams(ctx, ibcKeeper)

	return monitoringKeeper, ibcKeeper, ctx
}
