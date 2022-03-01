package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	sdk.Context,
	testkeeper.TestKeepers,
	types.MsgServer,
	profiletypes.MsgServer,
	campaigntypes.MsgServer,
) {
	ctx, tk := testkeeper.NewTestKeepers(t)

	return ctx,
		tk,
		keeper.NewMsgServerImpl(*tk.LaunchKeeper),
		profilekeeper.NewMsgServerImpl(*tk.ProfileKeeper),
		campaignkeeper.NewMsgServerImpl(*tk.CampaignKeeper)
}
