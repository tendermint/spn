package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*launchkeeper.Keeper,
	*profilekeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	launchtypes.MsgServer,
	sdk.Context,
) {
	_, launchKeeper, profileKeeper, rewardKeeper, _, _, ctx := testkeeper.AllKeepers(t)

	return rewardKeeper,
		launchKeeper,
		profileKeeper,
		keeper.NewMsgServerImpl(*rewardKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		launchkeeper.NewMsgServerImpl(*launchKeeper),
		ctx
}
