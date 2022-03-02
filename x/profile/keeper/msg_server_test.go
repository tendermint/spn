package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (sdk.Context, testkeeper.TestKeepers, types.MsgServer) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	return ctx, tk, keeper.NewMsgServerImpl(*tk.ProfileKeeper)
}
