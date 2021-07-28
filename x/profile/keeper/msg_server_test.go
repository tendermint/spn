package keeper

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	keeper, ctx := setupKeeper(t)
	return NewMsgServerImpl(*keeper), sdk.WrapSDKContext(ctx)
}

// TODO remove it after merge the PR#205
func setupMsgServerAndKeeper(t testing.TB) (sdk.Context, *Keeper, types.MsgServer) {
	keeper, ctx := setupKeeper(t)
	return ctx, keeper, NewMsgServerImpl(*keeper)
}
