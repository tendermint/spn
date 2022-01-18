package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Monitoringc(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
