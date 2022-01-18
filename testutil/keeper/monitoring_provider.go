package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibckeeper "github.com/cosmos/ibc-go/modules/core/keeper"
	"github.com/tendermint/spn/x/monitoringp/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func MonitoringpKeeper(t testing.TB) (*keeper.Keeper, *ibckeeper.Keeper, sdk.Context) {
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
		)

	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	monitoringKeeper.SetParams(ctx, types.DefaultParams())
	setIBCDefaultParams(ctx, ibcKeeper)

	return monitoringKeeper, ibcKeeper, ctx
}
