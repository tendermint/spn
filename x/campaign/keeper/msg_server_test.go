package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*profilekeeper.Keeper,
	*launchkeeper.Keeper,
	bankkeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	sdk.Context,
) {
	campaignKeeper, launchKeeper, profileKeeper, _, _, bankKeeper, ctx := testkeeper.AllKeepers(t)

	return campaignKeeper,
		profileKeeper,
		launchKeeper,
		bankKeeper,
		keeper.NewMsgServerImpl(*campaignKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		ctx
}
