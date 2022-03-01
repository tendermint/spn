package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	sdk.Context,
	testkeeper.TestKeepers,
	types.MsgServer,
	profiletypes.MsgServer,
) {
	ctx, tk := testkeeper.NewTestKeepers(t)

	return ctx,
		tk,
		keeper.NewMsgServerImpl(*tk.CampaignKeeper),
		profilekeeper.NewMsgServerImpl(*tk.ProfileKeeper)
}
