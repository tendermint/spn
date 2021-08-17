package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (sdk.Context, *keeper.Keeper, types.MsgServer) {
	k, ctx := testkeeper.Profile(t)
	return ctx, k, keeper.NewMsgServerImpl(*k)
}
