package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/monitoringp/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, _, ctx := keepertest.MonitoringpKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
