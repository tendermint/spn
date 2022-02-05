package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
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
	bankkeeper.Keeper,
	authkeeper.AccountKeeper,
	types.MsgServer,
	profiletypes.MsgServer,
	launchtypes.MsgServer,
	sdk.Context,
) {
	_, launchKeeper, profileKeeper, rewardKeeper, _, bankKeeper, authKeeper, ctx := testkeeper.AllKeepers(t)

	return rewardKeeper,
		launchKeeper,
		profileKeeper,
		bankKeeper,
		authKeeper,
		keeper.NewMsgServerImpl(*rewardKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		launchkeeper.NewMsgServerImpl(*launchKeeper),
		ctx
}
