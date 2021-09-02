package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*profilekeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	sdk.Context,
	codec.Marshaler,
) {
	k, profileKeeper, ctx, cdc := testkeeper.Launch(t)

	return k,
		profileKeeper,
		keeper.NewMsgServerImpl(*k),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		ctx,
		cdc
}
