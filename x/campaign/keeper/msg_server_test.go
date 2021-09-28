package keeper_test

import (
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*profilekeeper.Keeper,
	*launchkeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	sdk.Context,
) {
	campaignKeeper, launchKeeper, profileKeeper, ctx := testkeeper.AllKeepers(t)

	return campaignKeeper,
		profileKeeper,
		launchKeeper,
		keeper.NewMsgServerImpl(*campaignKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		ctx
}
