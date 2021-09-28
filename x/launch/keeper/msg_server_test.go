package keeper_test

import (
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*profilekeeper.Keeper,
	*campaignkeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	campaigntypes.MsgServer,
	sdk.Context,
) {
	campaignKeeper, launchLKeeper, profileKeeper, ctx := testkeeper.AllKeepers(t)

	return launchLKeeper,
		profileKeeper,
		campaignKeeper,
		keeper.NewMsgServerImpl(*launchLKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		campaignkeeper.NewMsgServerImpl(*campaignKeeper),
		ctx
}
