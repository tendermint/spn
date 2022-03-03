package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*profilekeeper.Keeper,
	*campaignkeeper.Keeper,
	bankkeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	campaigntypes.MsgServer,
	sdk.Context,
) {
	campaignKeeper, launchLKeeper, profileKeeper, _, _, bankKeeper, _, ctx := testkeeper.AllKeepers(t)

	return launchLKeeper,
		profileKeeper,
		campaignKeeper,
		bankKeeper,
		keeper.NewMsgServerImpl(*launchLKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		campaignkeeper.NewMsgServerImpl(*campaignKeeper),
		ctx
}
