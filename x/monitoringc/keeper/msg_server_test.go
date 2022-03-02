package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	sdk.Context,
	testkeeper.TestKeepers,
	types.MsgServer,
	profiletypes.MsgServer,
	launchtypes.MsgServer,
) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	return ctx,
		tk,
		keeper.NewMsgServerImpl(*tk.MonitoringConsumerKeeper),
		profilekeeper.NewMsgServerImpl(*tk.ProfileKeeper),
		launchkeeper.NewMsgServerImpl(*tk.LaunchKeeper)
}
