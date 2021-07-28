package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

func setupMsgServer(t testing.TB) (
	*Keeper,
	*profilekeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	sdk.Context,
	codec.Marshaler,
) {
	keeper, profileKeeper, ctx, cdc := setupKeeper(t)

	return keeper,
		profileKeeper,
		NewMsgServerImpl(*keeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		ctx,
		cdc
}
