package keeper_test

import (
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
	types.MsgServer,
	profiletypes.MsgServer,
	sdk.Context,
) {
	k, ctx := testkeeper.Launch(t)

	return k,
		keeper.NewMsgServerImpl(*k),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		ctx
}
