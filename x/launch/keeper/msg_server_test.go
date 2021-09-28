package keeper_test

import (
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*campaignkeeper.Keeper,
	*profilekeeper.Keeper,
	types.MsgServer,
	campaigntypes.MsgServer,
	profiletypes.MsgServer,
	sdk.Context,
) {
	campaignKeeper, launchLKeeper, profileKeeper, ctx := testkeeper.AllKeepers(t)

	return launchLKeeper,
		campaignKeeper,
		profileKeeper,
		keeper.NewMsgServerImpl(*launchLKeeper),
		campaignkeeper.NewMsgServerImpl(*campaignKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		ctx
}
