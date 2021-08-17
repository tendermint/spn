package keeper_test

import (
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/profile/keeper"
)

func setupMsgServer(t testing.TB) (sdk.Context, *keeper.Keeper, types.MsgServer) {
	k, ctx := testkeeper.Profile(t)
	return ctx, k, keeper.NewMsgServerImpl(*k)
}
