package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (sdk.Context, *Keeper, types.MsgServer) {
	keeper, ctx := setupKeeper(t)
	return ctx, keeper, NewMsgServerImpl(*keeper)
}

// TODO remove it after merge the PR#205
func setupMsgServerAndKeeper(t testing.TB) (sdk.Context, *Keeper, types.MsgServer) {
	keeper, ctx := setupKeeper(t)
	return ctx, keeper, NewMsgServerImpl(*keeper)
}

// TODO remove it after merge the PR#205
func setupMsgServerAndKeeper(t testing.TB) (sdk.Context, *Keeper, types.MsgServer) {
	keeper, ctx := setupKeeper(t)
	return ctx, keeper, NewMsgServerImpl(*keeper)
}
