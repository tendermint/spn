package keeper_test

import (
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*profilekeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	sdk.Context,
) {
	campaignKeeper, _, profileKeeper, ctx := testkeeper.AllKeepers(t)

	return campaignKeeper,
		profileKeeper,
		keeper.NewMsgServerImpl(*campaignKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		ctx
}