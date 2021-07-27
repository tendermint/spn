package keeper

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	keeper, ctx := SetupTestKeeper(t)
	return NewMsgServerImpl(*keeper), sdk.WrapSDKContext(ctx)
}
