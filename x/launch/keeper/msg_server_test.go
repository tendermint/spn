package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
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
