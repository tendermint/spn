package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibckeeper "github.com/cosmos/ibc-go/v2/modules/core/keeper"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*profilekeeper.Keeper,
	*launchkeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	launchtypes.MsgServer,
	*ibckeeper.Keeper,
	sdk.Context,
) {
	_, launchKeeper, profileKeeper, _, monitoringcKeeper, _, ibcKeeper, ctx := testkeeper.AllKeepers(t)

	return monitoringcKeeper,
		profileKeeper,
		launchKeeper,
		keeper.NewMsgServerImpl(*monitoringcKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		launchkeeper.NewMsgServerImpl(*launchKeeper),
		ibcKeeper,
		ctx
}
